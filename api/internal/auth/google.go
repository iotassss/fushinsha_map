package auth

import (
	"context"

	"github.com/iotassss/fushinsha-map-api/internal/port"
	googleidtoken "google.golang.org/api/idtoken"
)

type GoogleAuthVerifier struct {
	Audience string // GoogleのクライアントID
}

func NewGoogleAuthVerifier(audience string) *GoogleAuthVerifier {
	return &GoogleAuthVerifier{
		Audience: audience,
	}
}

func (g *GoogleAuthVerifier) VerifyAndParse(
	ctx context.Context,
	idToken string,
) (port.Token, error) {
	payload, err := googleidtoken.Validate(ctx, idToken, g.Audience)
	if err != nil {
		return port.Token{}, err
	}

	sub, ok := payload.Claims["sub"].(string)
	if !ok {
		return port.Token{}, err
	}

	return port.Token{
		Claims: port.Claims{
			Subject: sub,
		},
	}, nil
}
