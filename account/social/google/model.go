package google

type Error struct {
	ErrorDescription string `json:"error_description"`
}

func (e Error) Error() string {
	return e.ErrorDescription
}

type TokenInfo struct {
	Aud string `json:"aud"`
	Exp string `json:"exp"`
	Iat string `json:"iat"`
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Email string `json:"email"`
}
