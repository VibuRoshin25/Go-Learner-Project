package payload

type SignInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GenerateTokenPayload struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}

type ValidateTokenPayload struct {
	Token string `json:"token"`
}
