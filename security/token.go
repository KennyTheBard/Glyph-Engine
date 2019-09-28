package security

import "crypto/rand"

func GenerateRandomBytes(n int) []byte {
	bs := make([]byte, n)
	rand.Read(bs)
	return bs
}

func GenerateRandomToken(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes, _ := GenerateRandomBytes(n)
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}
