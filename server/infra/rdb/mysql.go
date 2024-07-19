package rdb

import (
	"fmt"
	"log"
	"sync"
	"tMinamiii/Tweet/env"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

var (
	session  *dbr.Session
	doOnce   sync.Once
	sessLock sync.RWMutex
)

func InitSession() {
	doOnce.Do(func() {
		conn, err := dbr.Open("mysql", DSN(), nil)
		if err != nil {
			log.Fatalf("failed to create session. err=%v", err)
		}

		num, err := env.DBMaxConnections()
		if err != nil {
			log.Fatalf("failed to create session. err=%v", err)
		}

		conn.SetMaxOpenConns(num)

		s := conn.NewSession(nil)

		session = s
	})
}

func GetTweetSession() *dbr.Session {
	sessLock.Lock()
	defer sessLock.Unlock()

	if session == nil {
		InitSession()
	}

	return session
}

func DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/tweet?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", env.DBUser(), env.DBPassword(), env.DBHost(), env.DBPort())
}
