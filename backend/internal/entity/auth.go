package entity

type AuthManager interface {
	MakeAuthn(userID uint) (string, error)
	FetchAuthn(tokenString string) (*map[string]string, error)
}
