package goutils

import (
	"github.com/satori/go.uuid"
)

func UuidFromBytes(input []byte) (string, error) {
	u, err := uuid.FromBytes(input)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
