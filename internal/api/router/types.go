package router

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginRequest) CleanParams() error {
	return nil
}

type LoginResponse struct {
	AuthToken string `json:"authToken"`
}
