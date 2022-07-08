package data

import (
	"crypto/sha512"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"user/internal/biz"
)

var _ = ginkgo.Describe("User", func() {
	var ro biz.UserRepo
	var uD *biz.User
	ginkgo.BeforeEach(func() {
		ro = NewUserRepo(Db, nil)
		uD = &biz.User{
			ID:       1,
			Mobile:   "18043838438",
			Password: "geek123456",
			NickName: "geeker",
			Role:     1,
			Birthday: 693629981,
		}
	})
	ginkgo.It("CreateUser", func() {
		u, err := ro.CreateUser(ctx, uD)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(u.Mobile).Should(gomega.Equal("18043838438"))
	})
	ginkgo.It("ListUser", func() {
		user, total, err := ro.ListUser(ctx, 1, 10)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(user).ShouldNot(gomega.BeEmpty())
		gomega.Ω(total).Should(gomega.Equal(1))
		gomega.Ω(len(user)).Should(gomega.Equal(1))
		gomega.Ω(user[0].Mobile).Should(gomega.Equal("18043838438"))
	})
	ginkgo.It("GetUserByMobile", func() {
		user, err := ro.GetUserByMobile(ctx, "18043838438")
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(user.Mobile).Should(gomega.Equal("18043838438"))
	})
	ginkgo.It("UpdateUser", func() {
		info := &biz.User{
			ID:       1,
			Mobile:   "18543838438",
			Password: "geek123456",
			NickName: "geeker",
			Role:     1,
			Birthday: 693629981,
		}
		user, err := ro.UpdateUser(ctx, info)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(user).Should(gomega.Equal(true))
	})
	ginkgo.It("CheckPassword", func() {
		p := "geek123456"
		options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
		salt, encodedPwd := password.Encode(p, options)
		ep := "$pbkdf2-sha512$" + salt + "$" + encodedPwd
		check, err := ro.CheckPassword(ctx, p, ep)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(check).Should(gomega.BeTrue())
	})
	ginkgo.It("GetUserById", func() {
		user, err := ro.GetUserById(ctx, 1)
		gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
		gomega.Ω(user.ID).Should(gomega.Equal(int64(1)))
	})
})
