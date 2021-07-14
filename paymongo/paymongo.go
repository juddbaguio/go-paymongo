package paymongo

import (
	"encoding/base64"
	"errors"
)

type PaymongoInstance struct {
	secret string
}

func NewProvider(secret_key string) (*PaymongoInstance, error) {
	if secret_key != "" {
		return &PaymongoInstance{
			secret: base64.StdEncoding.EncodeToString([]byte(secret_key)),
		}, nil
	}

	return &PaymongoInstance{
		secret: "",
	}, errors.New("secret key is required")
}
