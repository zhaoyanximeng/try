package lib

import "context"

type Auth struct {
	Token string `json:"token"`
}

func NewAuth(token string) *Auth {
	return &Auth{Token: token}
}

func (z *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"token": z.Token,
	}, nil
}

func (z *Auth) RequireTransportSecurity() bool {
	return false
}
