package database

// TokenStorage to store word count
type TokenStorage interface {
	AddToken(name string)
	GetTokenCount(name string) int
	UpdateTokenName(old, new string)
	UpdateTokenCount(name string, count int)
	RemoveToken(name string)
	ResetToken(name string)
	Size() int
}
