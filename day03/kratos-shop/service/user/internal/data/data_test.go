package data

import (
	"context"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"testing"
	"user/internal/conf"
)

func TestData(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "test biz data")
}

var cleaner func()
var Db *Data
var ctx context.Context

func initialize(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
	)
	return errors.WithStack(err)
}

var _ = ginkgo.BeforeSuite(func() {
	con, f := DockerTestMysql("mysql", "latest")
	cleaner = f
	config := &conf.Data{Database: &conf.Data_Database{Driver: "mysql", Source: con}}
	db := NewDB(config)
	mySqlDb, _, err := NewData(config, nil, db, nil)
	if err != nil {
		return
	}

	Db = mySqlDb
	err = initialize(db)
	if err != nil {
		return
	}
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
})

var _ = ginkgo.AfterSuite(func() {
	cleaner()
})
