package context

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type SomeMock struct {
	mock.Mock
}

func (m *SomeMock) Get(ctx context.Context, myPseudo, targetNamespace string) (string, error) {
	return "", nil
}

type Arguments []interface{}

type X struct {
	RunFn func(Arguments)
}

func Local(x X) {
}
