package biz_test

import (
	"crypto/sha512"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"strings"
	"user/internal/biz"
	"user/internal/mocks/mrepo"
)

var _ = ginkgo.Describe("UserUsecase", func() {
	var userCase *biz.UserUsecase
	var mUserRepo *mrepo.MockUserRepo
	ginkgo.BeforeEach(func() {
		mUserRepo = mrepo.NewMockUserRepo(ctl)
		userCase = biz.NewUserUsecase(mUserRepo, nil)
	})

	ginkgo.It("CreateUser", func() {
		info := &biz.User{
			ID:       1,
			Mobile:   "18043838438",
			Password: "geek123456",
			NickName: "geeker",
			Role:     1,
			Birthday: 693629981,
		}
		mUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(info, nil)
		l, err := userCase.CreateUser(ctx, info)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(err).ToNot(gomega.HaveOccurred())
		gomega.Ω(l.ID).To(gomega.Equal(int64(1)))
		gomega.Ω(l.Mobile).To(gomega.Equal("18043838438"))
	})

	ginkgo.It("ListUser", func() {
		info := make([]*biz.User, 0)
		info = append(info, &biz.User{
			ID:       1,
			Mobile:   "18043838438",
			Password: "geek123456",
			NickName: "geeker",
			Role:     1,
			Birthday: 693629981,
		})
		mUserRepo.EXPECT().ListUser(ctx, gomock.Any(), gomock.Any()).Return(info, 1, nil)
		l, total, err := userCase.ListUser(ctx, 1, 10)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(err).ToNot(gomega.HaveOccurred())
		gomega.Ω(l).ToNot(gomega.BeEmpty())
		gomega.Ω(len(l)).To(gomega.Equal(1))
		gomega.Ω(total).To(gomega.Equal(1))
		gomega.Ω(l[0].Mobile).To(gomega.Equal("18043838438"))
	})
	ginkgo.It("UpdateUser", func() {
		updateInfo := &biz.User{
			ID:       1,
			Mobile:   "18543838438",
			Password: "geek123456",
			NickName: "geeker",
			Role:     1,
			Birthday: 693629981,
		}

		mUserRepo.EXPECT().UpdateUser(ctx, gomock.Any()).Return(true, nil)
		l, err := userCase.UpdateUser(ctx, updateInfo)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(err).ToNot(gomega.HaveOccurred())
		gomega.Ω(l).To(gomega.Equal(true))
	})
	ginkgo.It("GetUserByMobile", func() {
		info := &biz.User{
			ID:       1,
			Mobile:   "18043838438",
			Password: "geek123456",
			NickName: "geeker",
			Role:     1,
			Birthday: 693629981,
		}
		mUserRepo.EXPECT().GetUserByMobile(ctx, gomock.Any()).Return(info, nil)
		l, err := userCase.GetUserByMobile(ctx, "18043838438")
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(err).ToNot(gomega.HaveOccurred())
		gomega.Ω(l.ID).To(gomega.Equal(int64(1)))
		gomega.Ω(l.Mobile).To(gomega.Equal("18043838438"))
	})
	ginkgo.It("CheckPassword", func() {
		p := "geek123456"
		options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
		salt, encodedPwd := password.Encode(p, options)
		ep := "$pbkdf2-sha512$" + salt + "$" + encodedPwd
		epl := strings.Split(ep, "$")
		c := password.Verify(p, epl[2], epl[3], options)
		mUserRepo.EXPECT().CheckPassword(ctx, p, ep).Return(c, nil)
		check, err := userCase.CheckPassword(ctx, p, ep)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(err).ToNot(gomega.HaveOccurred())
		gomega.Ω(check).To(gomega.BeTrue())
	})
	ginkgo.It("GetUserById", func() {
		info := &biz.User{
			ID:       1,
			Mobile:   "18043838438",
			Password: "geek123456",
			NickName: "geeker",
			Role:     1,
			Birthday: 693629981,
		}
		mUserRepo.EXPECT().GetUserById(ctx, gomock.Any()).Return(info, nil)
		l, err := userCase.GetUserById(ctx, 1)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(err).ToNot(gomega.HaveOccurred())
		gomega.Ω(l.ID).To(gomega.Equal(int64(1)))
	})
})
