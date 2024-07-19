package env

import (
	"log"
	"os"
	"strconv"
)

func Env() string {
	return os.Getenv("ENV")
}

func DBHost() string {
	return os.Getenv("DB_HOST")
}

func DBUser() string {
	return os.Getenv("DB_USER")
}

func DBPort() string {
	return os.Getenv("DB_PORT")
}

func DBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func DBMaxConnections() (int, error) {
	v := os.Getenv("DB_MAX_CONNECTIONS")
	conn, err := strconv.Atoi(v)
	if err != nil {
		log.Printf("failed to atoi DB_MAX_CONNECTIONS. err=%v\n", err)
		return 0, err
	}
	return conn, nil
}

func APIKey() string {
	return os.Getenv("API_KEY")
}

func SessionName() string {
	return os.Getenv("SESSION_NAME")
}

func DummySessionAccountID() string {
	return os.Getenv("DUMMY_SESSION_ACCOUNT_ID")
}
