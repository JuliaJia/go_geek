package data

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strings"
	"time"
	"user/internal/biz"
)

type User struct {
	ID          int64      `gorm:"primarykey"`
	Mobile      string     `gorm:"index:idx_mobile;unique;type:varchar(11) comment '手机号码,唯一标识';not null"`
	Password    string     `gorm:"type:varchar(100) comment '密码需要加密';not null"`
	NickName    string     `gorm:"type:varchar(25) comment '用户昵称'"`
	Birthday    *time.Time `gorm:"type:datetime comment '用户生日'"`
	Gender      string     `gorm:"column:gender;default:male;type:varchar(16) comment '性别'"`
	Role        int        `gorm:"column:role;default:1;type:int comment '1 - 普通用户，2 - 管理员'"`
	CreatedAt   time.Time  `gorm:"column:create_time"`
	UpdatedAt   time.Time  `gorm:"column:update_time"`
	DeletedAt   gorm.DeletedAt
	isDeletedAt bool
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	var user User
	result := r.data.db.Where(&biz.User{Mobile: u.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Mobile = u.Mobile
	user.NickName = u.NickName
	user.Password = encrypt(u.Password)
	res := r.data.db.Create(&user)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return &biz.User{
		ID:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     user.Role,
	}, nil
}

func encrypt(psd string) string {
	options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(psd, options)
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

//用户列表
func (ur *userRepo) ListUser(ctx context.Context, pageNum, pageSize int) ([]*biz.User, int, error) {
	var users []User
	result := ur.data.db.Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	total := int(result.RowsAffected)
	ur.data.db.Scopes(paginate(pageNum, pageSize)).Find(&users)
	urv := make([]*biz.User, 0)
	for _, u := range users {
		urv = append(urv, &biz.User{
			ID:       u.ID,
			Mobile:   u.Mobile,
			Password: u.Password,
			NickName: u.NickName,
			Gender:   u.Gender,
			Role:     u.Role,
		})
	}
	return urv, total, nil
}

func paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (ur *userRepo) GetUserByMobile(ctx context.Context, mobile string) (*biz.User, error) {
	var user User
	result := ur.data.db.Where(&User{Mobile: mobile}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	re := modelToResponse(user)
	return &re, nil
}

func modelToResponse(user User) biz.User {
	userInfoRsp := biz.User{
		ID:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     user.Role,
	}
	return userInfoRsp
}

func (ur *userRepo) UpdateUser(ctx context.Context, u *biz.User) (bool, error) {
	var userInfo User
	birthday := time.Unix(u.Birthday, 0)
	result := ur.data.db.Where(&User{ID: u.ID}).First(&userInfo)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, status.Error(codes.NotFound, "用户不存在")
	}
	userInfo.NickName = u.NickName
	userInfo.Gender = u.Gender
	userInfo.Birthday = &birthday
	res := ur.data.db.Save(&userInfo)
	if res.Error != nil {
		return false, status.Error(codes.Internal, res.Error.Error())
	}
	return true, nil
}

func (ur *userRepo) CheckPassword(ctx context.Context, p, ep string) (bool, error) {
	op := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
	pi := strings.Split(ep, "$")
	check := password.Verify(p, pi[2], pi[3], op)
	return check, nil
}

func (ur *userRepo) GetUserById(ctx context.Context, id int64) (*biz.User, error) {
	var user User
	result := ur.data.db.Where(&User{ID: id}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	re := modelToResponse(user)
	return &re, nil
}
