package model

import "golang.org/x/crypto/bcrypt"

// User ...
type User struct {
	ID                 int
	Email              string
	Password           string
	Encrypted_password string
}

// BeforeCreate
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.Encrypted_password = enc
	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
