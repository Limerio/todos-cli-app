package utils

import (
	"github.com/google/uuid"
)

func GenerateUuid() (uuid.UUID, error) {
	generatedUuid, err := uuid.NewV7()

	if err != nil {
		return uuid.UUID{}, err
	}

	return generatedUuid, nil
}
