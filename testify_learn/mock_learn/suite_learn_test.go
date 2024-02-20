package mock_learn

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMyTestSuit(t *testing.T) {
	suite.Run(t, new(MyTestSuit))
}
