package passwords

import "testing"

func TestRandAlphanumericString(t *testing.T) {

	t.Run("Success. 32 bytes", func(t *testing.T) {
		strLen := 32
		randStr := RandAlphanumericString(strLen)
		returnedLen := len(randStr)

		if returnedLen != strLen {
			t.Errorf("Random string returned %s has a different length than requested: %d", randStr, returnedLen)
		}

		t.Log("Random string generated OK:", randStr)
	})

	t.Run("Success. 128 bytes", func(t *testing.T) {
		strLen := 128
		randStr := RandAlphanumericString(strLen)
		returnedLen := len(randStr)

		if returnedLen != strLen {
			t.Errorf("Random string returned %s has a different length than requested: %d", randStr, returnedLen)
		}

		t.Log("Random string generated OK:", randStr)
	})

	t.Run("Success. 128 bytes. 2 calls, different rand strings, properly seeded.", func(t *testing.T) {
		strLen := 128

		randStr1 := RandAlphanumericString(strLen)
		randStr2 := RandAlphanumericString(strLen)

		if randStr1 == randStr2 {
			t.Errorf("Random strings returned %s and %s are equal", randStr1, randStr2)
		}

		t.Log("Random string 1 generated OK:", randStr1)
		t.Log("Random string 2 generated OK:", randStr2)
	})
}

func TestRandASCIIString(t *testing.T) {

	t.Run("Success. 32 bytes", func(t *testing.T) {
		strLen := 32
		randStr := RandASCIIString(strLen)
		returnedLen := len(randStr)

		if returnedLen != strLen {
			t.Errorf("Random string returned %s has a different length than requested: %d", randStr, returnedLen)
		}

		t.Log("Random string generated OK:", randStr)
	})

	t.Run("Success. 128 bytes", func(t *testing.T) {
		strLen := 128
		randStr := RandASCIIString(strLen)
		returnedLen := len(randStr)

		if returnedLen != strLen {
			t.Errorf("Random string returned %s has a different length than requested: %d", randStr, returnedLen)
		}

		t.Log("Random string generated OK:", randStr)
	})

	t.Run("Success. 128 bytes. 2 calls, different rand strings, properly seeded.", func(t *testing.T) {
		strLen := 128

		randStr1 := RandASCIIString(strLen)
		randStr2 := RandASCIIString(strLen)

		if randStr1 == randStr2 {
			t.Errorf("Random strings returned %s and %s are equal", randStr1, randStr2)
		}

		t.Log("Random string 1 generated OK:", randStr1)
		t.Log("Random string 2 generated OK:", randStr2)
	})
}
