// Assert template
package main

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type MainSuite struct {
	suite.Suite
}

func (s *MainSuite) TestsetLogger0() {
	xvPbA := log.Level(uint32(72))

	setLogger(xvPbA)
}

func (s *MainSuite) TestsetLogger1() {
	oEcHq := log.Level(uint32(72))

	setLogger(oEcHq)
}

func (s *MainSuite) TestsetLogger2() {
	ezvmf := log.Level(uint32(61))

	setLogger(ezvmf)
}

func (s *MainSuite) TestsetLogger3() {
	yRIXE := log.Level(uint32(24))

	setLogger(yRIXE)
}

func (s *MainSuite) TestsetLogger4() {
	pLMlY := log.Level(uint32(0))

	setLogger(pLMlY)
}

func (s *MainSuite) TestsetLogger5() {
	slhdS := log.Level(uint32(15))

	setLogger(slhdS)
}

func (s *MainSuite) TestsetLogger6() {
	ppCUP := log.Level(uint32(51))

	setLogger(ppCUP)
}

func (s *MainSuite) TestsetLogger7() {
	gwpli := log.Level(uint32(92))

	setLogger(gwpli)
}

func (s *MainSuite) TestsetLogger8() {
	rZBrO := log.Level(uint32(51))

	setLogger(rZBrO)
}

func (s *MainSuite) TestsetLogger9() {
	rzGjt := log.Level(uint32(97))

	setLogger(rzGjt)
}

func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}
