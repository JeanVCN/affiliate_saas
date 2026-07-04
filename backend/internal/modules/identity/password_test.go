package identity

import "testing"

func TestPasswordHashRoundTrip(t *testing.T) {
	hash, err := hashPassword("CorrectHorse123")
	if err != nil {
		t.Fatalf("hash password: %v", err)
	}
	if hash == "" {
		t.Fatal("hash is empty")
	}

	ok, err := verifyPassword("CorrectHorse123", hash)
	if err != nil {
		t.Fatalf("verify password: %v", err)
	}
	if !ok {
		t.Fatal("password should verify")
	}

	ok, err = verifyPassword("WrongHorse123", hash)
	if err != nil {
		t.Fatalf("verify wrong password: %v", err)
	}
	if ok {
		t.Fatal("wrong password should not verify")
	}
}

func TestValidatePasswordPolicy(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{name: "valid", password: "CorrectHorse123", wantErr: false},
		{name: "short", password: "Short123", wantErr: true},
		{name: "no number", password: "CorrectHorse", wantErr: true},
		{name: "no letter", password: "123456789012", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePassword(tt.password)
			if tt.wantErr && err == nil {
				t.Fatal("expected validation error")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected validation error: %v", err)
			}
		})
	}
}
