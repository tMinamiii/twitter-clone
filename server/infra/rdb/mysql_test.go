package rdb

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type testTable struct {
	Username string `db:"username"`
}

func TestGetTweetSession(t *testing.T) {
	got := GetTweetSession()
	if got == nil {
		t.Errorf("GetTweetSession() is nil")
	}
	t.Cleanup(func() { got.Exec("DROP TABLE IF EXISTS test_table") })

	_, err := got.Exec("CREATE TABLE test_table ( username VARCHAR(255) PRIMARY KEY ) ENGINE = InnoDB")
	if err != nil {
		t.Errorf("failed to create test_table %v", err)
	}

	// セッションが使用できるかInsert/Selectして確認
	username := "testname"
	ctx := context.Background()
	if _, err := got.InsertInto("test_table").Pair("username", username).ExecContext(ctx); err != nil {
		t.Errorf("insert error %v", err)
	}

	m := &testTable{}
	err = got.Select("*").From("test_table").Where("username = ?", username).LoadOneContext(ctx, m)
	if err != nil {
		t.Errorf("select error %v", err)
	}

	if m.Username != username {
		t.Error("username is wrong value")
	}
}

func TestDSN(t *testing.T) {
	want := "tweet:passwd@tcp(localhost:23306)/tweet?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"
	if got := DSN(); got != want {
		t.Errorf("DSN() = %v, want %v", got, want)
	}
}
