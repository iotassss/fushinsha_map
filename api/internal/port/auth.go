package port

import (
	"context"
)

type Claims struct {
	// Issuer          string // iss
	// AuthorizedParty string // azp
	// Audience        string // aud

	Subject string // sub (Googleの一意ID)

	// AtHash          string // at_hash
	// HostedDomain    string // hd
	// Email           string // email
	// EmailVerified   bool   // email_verified
	// IssuedAt        int64  // iat (UNIX秒)
	// ExpiresAt       int64  // exp (UNIX秒)
	// Nonce           string // nonce
}

// https://github.com/golang-jwt/jwt/blob/main/token.goを参考に必要なフィールドを定義
type Token struct {
	// 	Raw       string         // Raw contains the raw token.  Populated when you [Parse] a token
	// 	Method    SigningMethod  // Method is the signing method used or to be used
	// 	Header    map[string]any // Header is the first segment of the token in decoded form

	Claims Claims // Claims is the second segment of the token in decoded form

	// Signature []byte         // Signature is the third segment of the token in decoded form.  Populated when you Parse a token
	// Valid     bool           // Valid specifies if the token is valid.  Populated when you Parse/Verify a token
}

type AuthVerifier interface {
	Verify(ctx context.Context, idToken string) error
}

type JWTParser interface {
	Parse(idToken string) (Token, error)
}

type GoogleAuthVerifier interface {
	VerifyAndParse(ctx context.Context, idToken string) (Token, error)
}
