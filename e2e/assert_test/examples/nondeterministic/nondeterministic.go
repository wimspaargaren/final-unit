package nondeterministic

import "github.com/gofrs/uuid"

func NonDeterministicFunc() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}

func DeterministicFunc() int {
	return 3
}
