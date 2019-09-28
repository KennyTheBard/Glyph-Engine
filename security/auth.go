package security

type TokenRegister map[string]string

func (reg *TokenRegister) Authorizate(username string) string {
	token := GenerateRandomToken(50)
	reg[username] = token
	return token
}

func (reg TokenRegister) VerifyToken(username, token string) bool {
	authToken, err := reg[username]
	return token == authToken
}
