package happyconsumer

// NOTE: you don't have to follow this single method interface, this will be useful if your codebase become humongous
//		 and you want to group some method into a mixer interface (dependency)

type TokenPusher interface {
	AddToken(name string)
}

type TokenGetter interface {
	GetTokenCount(name string) int
}

type Dependency interface {
	TokenPusher
	TokenGetter
}
