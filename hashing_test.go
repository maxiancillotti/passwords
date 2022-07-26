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

func TestHashing2(t *testing.T) {

	pw := "rW&r#`@nFRw]2@reVU+~h0q pn5}@`_TQkwf6\"H:f,`24&=gWi~ag%8GC/z#:+Gg"
	salt := "N|pf2dO:CwH\\2hjgH?dc(( ^agdh;4  "

	t.Logf("randpw: '%s'", pw)
	t.Logf("randsalt: '%s'", salt)

	storedHashpw := "$2a$10$eIqiLuuijq/kZo59gdq99eenwxhiXYVW1oTzkoYXzDHhqYM0C.3K2"
	t.Logf("storedHashpw: '%s'", storedHashpw)

	validationErr := ValidateHashPw(pw, salt, storedHashpw)
	if validationErr != nil {
		t.Error("Stored hash validation failed:", validationErr)
		return
	}

	newHashedpw, _ := HashPw(pw, salt)
	t.Logf("storedHashpw: '%s'", newHashedpw)

	validationErr = ValidateHashPw(pw, salt, newHashedpw)
	if validationErr != nil {
		t.Error("New hash validation failed:", validationErr)
		return
	}

	t.Log("Hash validation OK")
}

func TestHashingWithGenerator(t *testing.T) {
	pwLen := 64
	saltLen := 32

	pw := RandASCIIString(pwLen)
	salt := RandASCIIString(saltLen)

	t.Logf("randpw: '%s'", pw)
	t.Logf("randsalt: '%s'", salt)

	hashed, _ := HashPw(pw, salt)
	t.Logf("storedHashpw: '%s'", hashed)

	validationErr := ValidateHashPw(pw, salt, hashed)
	if validationErr != nil {
		t.Error("New hash validation failed:", hashed)
		return
	}

	t.Log("Hash validation OK")
}
