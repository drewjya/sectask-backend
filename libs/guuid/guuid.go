package guuid

import "github.com/google/uuid"

func GenerateNewUUIDV4() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
