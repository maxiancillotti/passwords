package passwords

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPw(pw, salt string) (string, error) {

	saltedpw := fmt.Sprint(pw, salt)
	hashedpw, err := bcrypt.GenerateFromPassword([]byte(saltedpw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedpw), nil
}

func ValidateHashPw(pw, salt, hash string) error {

	saltedpw := fmt.Sprint(pw, salt)
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(saltedpw))
}
