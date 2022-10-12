package jwt

import (
	"bds-go-auth-service/common/l"
	"bds-go-auth-service/component/tokenprovider"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
)

var (
	ll = l.New()
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{
		secret: secret,
	}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

func (j *jwtProvider) Generate(payload tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256, myClaims{
			Payload: payload,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: &jwt.NumericDate{
					Time: time.Now().UTC().Add(time.Duration(expiry)),
				},
				IssuedAt: &jwt.NumericDate{
					Time: time.Now().UTC(),
				},
			},
		},
	)

	myToken, err := t.SignedString([]byte(j.secret))

	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:   myToken,
		Created: time.Now().UTC(),
		Expiry:  expiry,
	}, nil

}

func (j *jwtProvider) Validate(tokenString string) (*tokenprovider.TokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if claims, ok := token.Claims.(*myClaims); ok && token.Valid {
		return &claims.Payload, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		ll.Error("token malformed", l.Object("err", err))
		return nil, tokenprovider.ErrInvalidToken
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		ll.Error("token expired", l.Object("err", err))
		return nil, tokenprovider.ErrInvalidToken
	} else {
		ll.Error("token expired", l.Object("err", err))
		return nil, tokenprovider.ErrInvalidToken
	}
}
