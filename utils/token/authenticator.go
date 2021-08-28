package token

import "github.com/duyledat197/netfix-backend/utils/requestinfo"

type Authenticator interface {
	Generate(payload *requestinfo.Info) (*Token, error)
	Verify(token string) (*Payload, error)
}
