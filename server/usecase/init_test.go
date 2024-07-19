package usecase

import (
	"tMinamiii/Tweet/project"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	godotenv.Load(project.Root() + "/.env.test")
	m.Run()
}
