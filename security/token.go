package security

import "crypto/rand"

func GenerateRandomBytes(n int) []byte {
	bs := make([]byte, n)
	if n, err := rand.Read(bs); n == 0 || err != nil {
		return nil
	}
	return bs
}

func GenerateRandomToken(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := GenerateRandomBytes(n)
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}
