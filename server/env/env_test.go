package env

import (
	"testing"
)

func TestEnv(t *testing.T) {
	want := "test"
	t.Setenv("ENV", want)
	if got := Env(); got != want {
		t.Errorf("Env() = %v, want %v", got, want)
	}
}

func TestDBHost(t *testing.T) {
	want := "dbhost"
	t.Setenv("DB_HOST", want)
	if got := DBHost(); got != want {
		t.Errorf("DBHost() = %v, want %v", got, want)
	}
}

func TestDBUser(t *testing.T) {
	want := "dbuser"
	t.Setenv("DB_USER", want)
	if got := DBUser(); got != want {
		t.Errorf("DBUser() = %v, want %v", got, want)
	}
}

func TestDBPort(t *testing.T) {
	want := "11111"
	t.Setenv("DB_PORT", want)
	if got := DBPort(); got != want {
		t.Errorf("DBPort() = %v, want %v", got, want)
	}
}
func TestDBPassword(t *testing.T) {
	want := "passwd"
	t.Setenv("DB_PASSWORD", want)
	if got := DBPassword(); got != want {
		t.Errorf("DBPassword() = %v, want %v", got, want)
	}
}

func TestDBMaxConnections(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		want    int
		wantErr bool
	}{
		{
			name: "環境変数の値が数字に変換できる場合",
			env:  "12345",
			want: 12345,
		},
		{
			name:    "環境変数の値が数字に変換できない場合",
			env:     "1234abc",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("DB_MAX_CONNECTIONS", tt.env)
			got, err := DBMaxConnections()
			if (err != nil) != tt.wantErr {
				t.Errorf("DBMaxConnections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DBMaxConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIKey(t *testing.T) {
	want := "apikey"
	t.Setenv("API_KEY", want)
	if got := APIKey(); got != want {
		t.Errorf("APIKey() = %v, want %v", got, want)
	}
}

func TestSessionName(t *testing.T) {
	want := "session_name"
	t.Setenv("SESSION_NAME", want)
	if got := SessionName(); got != want {
		t.Errorf("SessionName() = %v, want %v", got, want)
	}
}

func TestDummySessionAccountID(t *testing.T) {
	want := "account_id"
	t.Setenv("DUMMY_SESSION_ACCOUNT_ID", want)
	if got := DummySessionAccountID(); got != want {
		t.Errorf("DummySessionAccountID() = %v, want %v", got, want)
	}
}
