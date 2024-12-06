package happyconsumer

type TokenService struct {
	dep Dependency
}

func NewTokenService(dep Dependency) *TokenService {
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
