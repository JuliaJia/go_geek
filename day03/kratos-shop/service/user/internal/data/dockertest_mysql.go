package data

import (
	"database/sql"
	"fmt"
	"github.com/ory/dockertest/v3"
	"log"
	"time"
)

func DockerTestMysql(img, version string) (string, func()) {
	return innerDockerTestMysql(img, version)
}

func innerDockerTestMysql(img, version string) (string, func()) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Count not connect to docker: %s", err)
	}
	resource, err := pool.Run(img, version, []string{"MYSQL_ROOT_PASSWORD=secret", "MYSQL_ROOT_HOST=%"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	conStr := fmt.Sprintf("root:secret@(127.0.0.1:%s)/mysql?parseTime=True", resource.GetPort("3306/tcp"))

	if err := pool.Retry(func() error {
		var err error
		db, err := sql.Open("mysql", conStr)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return conStr, func() {
		if err = pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}
}
