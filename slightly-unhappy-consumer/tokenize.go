package slightlyunhappyconsumer

import "whatev.com/interface/database"

type TokenService struct {
	dep database.TokenStorage
}

func NewTokenService(dep database.TokenStorage) *TokenService {
	return &TokenService{
		dep: dep,
	}
}

func (s TokenService) Tokenize(tokens []string) {
	for _, v := range tokens {
		s.dep.AddToken(v)
	}
}

func (s TokenService) GetTokenCount(name string) int {
	return s.dep.GetTokenCount(name)
}
