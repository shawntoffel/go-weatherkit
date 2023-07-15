package weatherkit

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestGenerateSignedJwt(t *testing.T) {
	pk, err := createPrivateKeyPEM()
	if err != nil {
		t.Error(err)
	}

	g := Credentials{
		PrivateKey: pk,
		KeyID:      "key",
		TeamID:     "team",
		ServiceID:  "service",
	}

	signed, _, err := g.SignedJWT(time.Minute * 10)
	if err != nil {
		t.Error(err)
	}

	t.Log(signed)

	claims := jwt.MapClaims{}

	parsed, _ := jwt.ParseWithClaims(signed, &claims, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	alg := parsed.Header["alg"]
	expectedAlg := "ES256"
	if alg != "ES256" {
		t.Errorf("expected alg: %s, got: %s", expectedAlg, alg)
	}

	iss := claims["iss"]
	if iss != g.TeamID {
		t.Errorf("expected team ID : %s, got: %s", g.TeamID, iss)
	}

	sub := claims["sub"]
	if sub != g.ServiceID {
		t.Errorf("expected service ID : %s, got: %s", g.ServiceID, sub)
	}

	kid := parsed.Header["kid"]
	if kid != g.KeyID {
		t.Errorf("expected kid: %s, got: %s", g.KeyID, kid)
	}

	id := parsed.Header["id"]
	expectedID := g.TeamID + "." + g.ServiceID
	if id != g.TeamID+"."+g.ServiceID {
		t.Errorf("expected team ID: %s, got: %s", expectedID, id)
	}
}

func createPrivateKeyPEM() ([]byte, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return []byte{}, err
	}

	marshalled, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return []byte{}, err
	}

	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: marshalled,
	}

	return pem.EncodeToMemory(block), nil
}
