package utils

// API jdfklsf
type API interface {
	Login(email, password string) (string, error)
}
