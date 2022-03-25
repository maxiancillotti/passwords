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
}
