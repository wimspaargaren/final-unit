package chance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ChanceTestSuite struct {
	suite.Suite
}

func (s *ChanceTestSuite) TestIndex() {
	for i := 0; i < 100; i++ {
		res := GetIndex(4)
		s.True(res >= 0 && res <= 3)
	}
}

func (s *ChanceTestSuite) TestLengthZero() {
	for i := 0; i < 100; i++ {
		res := GetIndex(1)
		s.Equal(0, res)
	}
}

func TestChanceTestSuite(t *testing.T) {
	suite.Run(t, new(ChanceTestSuite))
}
