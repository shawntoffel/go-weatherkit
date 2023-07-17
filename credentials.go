package weatherkit

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const defaultTokenDuration time.Duration = time.Minute * 10

// Credentials holds information required to authenticate against the weatherkit API.
type Credentials struct {
	// PrivateKey is your PEM encoded private key.
	PrivateKey []byte

	// KeyID is the Key identifier from your developer account.
	KeyID string

	// TeamID is the Team ID from your developer account.
	TeamID string

	// ServiceID is the Service ID from your developer account.
	ServiceID string
}

// SignedJWT generates a valid JWT signed with your PEM private key.
// Returns the string JWT along with its expiration time.
func (c *Credentials) SignedJWT(validFor time.Duration) (string, time.Time, error) {
	token, exp, err := c.create(validFor)
	if err != nil {
		return "", time.Time{}, err
	}

	signed, err := c.sign(token)
	if err != nil {
		return "", time.Time{}, err
	}

	return signed, exp, nil
}

func (c *Credentials) create(validFor time.Duration) (*jwt.Token, time.Time, error) {
	err := c.validate(validFor)
	if err != nil {
		return nil, time.Time{}, err
	}

	now := time.Now().UTC()
	exp := now.Add(validFor)

	return &jwt.Token{
		Header: map[string]interface{}{
			"alg": jwt.SigningMethodES256.Alg(),
			"kid": c.KeyID,
			"id":  c.TeamID + "." + c.ServiceID,
		},
		Claims: jwt.MapClaims{
			"iss": c.TeamID,
			"sub": c.ServiceID,
			"iat": now.Unix(),
			"exp": exp.Unix(),
		},
		Method: jwt.SigningMethodES256,
	}, exp, nil
}

func (c *Credentials) sign(token *jwt.Token) (string, error) {
	privateKey, err := jwt.ParseECPrivateKeyFromPEM(c.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key. %s", err)
	}

	signedJwt, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to create signed JWT. %s", err)
	}

	return signedJwt, nil
}

func (c *Credentials) validate(validFor time.Duration) error {
	messages := []string{}

	if len(c.KeyID) < 1 {
		messages = append(messages, "key identifier may not be empty")
	}

	if len(c.TeamID) < 1 {
		messages = append(messages, "team ID may not be empty")
	}

	if len(c.ServiceID) < 1 {
		messages = append(messages, "service ID may not be empty")
	}

	if validFor <= 0 {
		messages = append(messages, "token expiration must be in the future (duration must be greater than 0)")
	}

	if len(messages) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(messages, ", "))
	}

	return nil
}
