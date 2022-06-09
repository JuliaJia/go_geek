package data

import (
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
})
