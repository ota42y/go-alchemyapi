package alchemyapi

import ()

// AlchemyAPI is alchemyapi client
type AlchemyAPI struct {
	c          *config
	connection apiConnection
}

type config struct {
	apikey string
}

// New return AlchemyAPI pointer
func New(token string) *AlchemyAPI {
	return &AlchemyAPI{
		c: &config{
			apikey: token,
		},
		connection: &httpImp{},
	}
}
