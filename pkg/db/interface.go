package db

// API dlfkdsl
type API interface {
	Connect(params ...string) error
}
