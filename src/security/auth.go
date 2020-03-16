package security

type TokenRegister map[string]string

var reg TokenRegister

func Authorizate(username string) string {
	if reg == nil {
		reg = make(TokenRegister)
	}

	token := GenerateRandomToken(50)
	reg[username] = token
	return token
}

func VerifyToken(username, token string) bool {
	if reg == nil {
		reg = make(TokenRegister)
	}

	authToken, ok := reg[username]
	if !ok {
		return false
	}
	return token == authToken
}
