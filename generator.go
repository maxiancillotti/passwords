package passwords

import (
	"math/rand"
	"time"
)

const (
	alphaNumericChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	// Note if you copy this: Double quotation mark ( " ) and back slash ( \ ) are escaped.
	printableASCIIChars = " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
)

var (
	randSeed int64
)

// Generates a random alphanumeric string of the given length.
// Call to generate a password or salt.
func RandAlphanumericString(slen int) string {

	refreshSeed()

	b := make([]byte, slen)
	for i := range b {
		b[i] = alphaNumericChars[rand.Intn(len(alphaNumericChars))]
	}
	return string(b)
}

// Generates a random ASCII (codes 32-126) string of the given length.
// Call to generate a password or salt.
func RandASCIIString(slen int) string {

	refreshSeed()

	b := make([]byte, slen)
	for i := range b {
		b[i] = printableASCIIChars[rand.Intn(len(printableASCIIChars))]
	}
	return string(b)
}

func refreshSeed() {

	previousSeed := randSeed

	for {
		randSeed = time.Now().UnixNano()

		if previousSeed != randSeed {
			break
		}
	}
	rand.Seed(randSeed)
}
