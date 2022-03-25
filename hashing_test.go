package passwords

import (
	"testing"
)

func TestHashing(t *testing.T) {
	pw := "APIUserPassword"
	salt := "sv8TM3WJcLQ4GXIsCBhUSS0964L4ZA7S"

	hashed, _ := HashPw(pw, salt)
	t.Log("Hashed pw:", hashed)

	validationErr := ValidateHashPw(pw, salt, hashed)
	if validationErr != nil {
		t.Error("Hash validation failed:", validationErr)
		return
	}
	t.Log("Hash validation OK")
}
