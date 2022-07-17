package middleware

import (
	"context"
)

type Authentication struct {
	clientId     string
	clientSecret string
}

func NewClientAuthentication(clientId, clientSecret string) *Authentication {
	return &Authentication{
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

// WithClientCredentials todo
func (a *Authentication) WithClientCredentials(clientID, clientSecret string) {
	a.clientId = clientID
	a.clientSecret = clientSecret
}

func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		ClientHeaderKey: a.clientId,
		ClientSecretKey: a.clientSecret,
	}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
