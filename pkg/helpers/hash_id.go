package helpers

import (
	"errors"
	"github.com/speps/go-hashids/v2"
)

// Interface pake nama generic biar kalau besok ganti library, service lo gak perlu tau
type IDObfuscator interface {
	Encode(id uint) (string, error)
	Decode(hash string) (uint, error)
}

type hashIDManager struct {
	hasher *hashids.HashID
}

func NewHashIDManager(salt string, minLength int) (IDObfuscator, error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength

	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}

	return &hashIDManager{hasher: h}, nil
}

func (m *hashIDManager) Encode(id uint) (string, error) {
	return m.hasher.Encode([]int{int(id)})
}

func (m *hashIDManager) Decode(hash string) (uint, error) {
	numbers, err := m.hasher.DecodeWithError(hash)
	if err != nil || len(numbers) == 0 {
		return 0, errors.New("invalid or tampered id")
	}
	return uint(numbers[0]), nil
}
