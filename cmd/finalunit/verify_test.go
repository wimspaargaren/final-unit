// Assert template
package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type VerifySuite struct {
	suite.Suite
}

func (s *VerifySuite) TestCheckGoFiles0() {
	kffANuAYpu := Opts{Verbose: false, Version: true, Dir: "Servantdirt"}
	kffAN := &kffANuAYpu

	klXDI := CheckGoFiles(kffAN)

	s.Error(klXDI)

	_ = klXDI
}

func (s *VerifySuite) TestCheckGoFiles1() {
	lShEdQbOcI := Opts{Verbose: false, Version: false, Dir: "Musemagenta"}
	lShEd := &lShEdQbOcI

	gitqa := CheckGoFiles(lShEd)

	s.Error(gitqa)

	_ = gitqa
}

func (s *VerifySuite) TestCheckGoFiles2() {
	qZkfxCUizm := Opts{Verbose: false, Version: false, Dir: "Wyrmfoil"}
	qZkfx := &qZkfxCUizm

	jgduU := CheckGoFiles(qZkfx)

	s.Error(jgduU)

	_ = jgduU
}

func (s *VerifySuite) TestCheckGoFiles3() {
	siNvepYntU := Opts{Verbose: true, Version: true, Dir: "Roverivory"}
	siNve := &siNvepYntU

	mObhv := CheckGoFiles(siNve)

	s.Error(mObhv)

	_ = mObhv
}

func (s *VerifySuite) TestCheckGoFiles4() {
	bigsTEthCQ := Opts{Verbose: false, Version: true, Dir: "Takerbig"}
	bigsT := &bigsTEthCQ

	fZQBb := CheckGoFiles(bigsT)

	s.Error(fZQBb)

	_ = fZQBb
}

func (s *VerifySuite) TestCheckGoFiles5() {
	vNsLSDjXfv := Opts{Verbose: true, Version: true, Dir: "Wolfmeteor"}
	vNsLS := &vNsLSDjXfv

	wfHvo := CheckGoFiles(vNsLS)

	s.Error(wfHvo)

	_ = wfHvo
}

func (s *VerifySuite) TestCheckGoFiles6() {
	ajwACLLtXs := Opts{Verbose: true, Version: false, Dir: "Whipstellar"}
	ajwAc := &ajwACLLtXs

	jAmzD := CheckGoFiles(ajwAc)

	s.Error(jAmzD)

	_ = jAmzD
}

func (s *VerifySuite) TestCheckGoFiles7() {
	dPCgEzwZxL := Opts{Verbose: true, Version: true, Dir: "Ridersmall"}
	dPCgE := &dPCgEzwZxL

	zLKAN := CheckGoFiles(dPCgE)

	s.Error(zLKAN)

	_ = zLKAN
}

func (s *VerifySuite) TestCheckGoFiles8() {
	rzUGYgVwaw := Opts{Verbose: true, Version: false, Dir: "Rabbitribbon"}
	rzUGY := &rzUGYgVwaw

	aUxDy := CheckGoFiles(rzUGY)

	s.Error(aUxDy)

	_ = aUxDy
}

func (s *VerifySuite) TestCheckGoFiles9() {
	zLSTCakvTJ := Opts{Verbose: false, Version: true, Dir: "Sargentpaper"}
	zLSTC := &zLSTCakvTJ

	lKIhs := CheckGoFiles(zLSTC)

	s.Error(lKIhs)

	_ = lKIhs
}

func (s *VerifySuite) TestGoImportsInstalled0() {
	nqVqh := GoImportsInstalled()

	s.NoError(nqVqh)

	_ = nqVqh
}

func (s *VerifySuite) TestGoImportsInstalled1() {
	pYYmp := GoImportsInstalled()

	s.NoError(pYYmp)

	_ = pYYmp
}

func (s *VerifySuite) TestGoImportsInstalled2() {
	fecRl := GoImportsInstalled()

	s.NoError(fecRl)

	_ = fecRl
}

func (s *VerifySuite) TestGoImportsInstalled3() {
	oXTZu := GoImportsInstalled()

	s.NoError(oXTZu)

	_ = oXTZu
}

func (s *VerifySuite) TestGoImportsInstalled4() {
	tZTah := GoImportsInstalled()

	s.NoError(tZTah)

	_ = tZTah
}

func (s *VerifySuite) TestGoImportsInstalled5() {
	hbQSy := GoImportsInstalled()

	s.NoError(hbQSy)

	_ = hbQSy
}

func (s *VerifySuite) TestGoImportsInstalled6() {
	ebemu := GoImportsInstalled()

	s.NoError(ebemu)

	_ = ebemu
}

func (s *VerifySuite) TestGoImportsInstalled7() {
	pQpzK := GoImportsInstalled()

	s.NoError(pQpzK)

	_ = pQpzK
}

func (s *VerifySuite) TestGoImportsInstalled8() {
	klJMH := GoImportsInstalled()

	s.NoError(klJMH)

	_ = klJMH
}

func (s *VerifySuite) TestGoImportsInstalled9() {
	eXFdB := GoImportsInstalled()

	s.NoError(eXFdB)

	_ = eXFdB
}

func (s *VerifySuite) TestUnreadableDirError0() {
	wcvVWjvMBB := UnreadableDir{Dir: "Scorpionmesquite"}
	wcvVW := &wcvVWjvMBB

	vxJqe := wcvVW.Error()

	s.EqualValues(string(`unable to read files at dir: Scorpionmesquite`), vxJqe)

	_ = vxJqe
}

func (s *VerifySuite) TestUnreadableDirError1() {
	ugrvxOrlBW := UnreadableDir{Dir: "Dutchessancient"}
	ugrvx := &ugrvxOrlBW

	vGVXY := ugrvx.Error()

	s.EqualValues(string(`unable to read files at dir: Dutchessancient`), vGVXY)

	_ = vGVXY
}

func (s *VerifySuite) TestUnreadableDirError2() {
	lRLJQmlPNZ := UnreadableDir{Dir: "Razorpattern"}
	lRLJQ := &lRLJQmlPNZ

	yVOrF := lRLJQ.Error()

	s.EqualValues(string(`unable to read files at dir: Razorpattern`), yVOrF)

	_ = yVOrF
}

func (s *VerifySuite) TestUnreadableDirError3() {
	qANhioJmHw := UnreadableDir{Dir: "Toucancherry"}
	qANhi := &qANhioJmHw

	tiObp := qANhi.Error()

	s.EqualValues(string(`unable to read files at dir: Toucancherry`), tiObp)

	_ = tiObp
}

func (s *VerifySuite) TestUnreadableDirError4() {
	eNoQYJalOw := UnreadableDir{Dir: "Saverlunar"}
	eNoQY := &eNoQYJalOw

	oBdMc := eNoQY.Error()

	s.EqualValues(string(`unable to read files at dir: Saverlunar`), oBdMc)

	_ = oBdMc
}

func (s *VerifySuite) TestUnreadableDirError5() {
	dPrRPqKbEe := UnreadableDir{Dir: "Marecoffee"}
	dPrRP := &dPrRPqKbEe

	wPqsX := dPrRP.Error()

	s.EqualValues(string(`unable to read files at dir: Marecoffee`), wPqsX)

	_ = wPqsX
}

func (s *VerifySuite) TestUnreadableDirError6() {
	rjdxTchmkr := UnreadableDir{Dir: "Fighterdirt"}
	rjdxT := &rjdxTchmkr

	vqEJy := rjdxT.Error()

	s.EqualValues(string(`unable to read files at dir: Fighterdirt`), vqEJy)

	_ = vqEJy
}

func (s *VerifySuite) TestUnreadableDirError7() {
	hInmqzbWtR := UnreadableDir{Dir: "Boashort"}
	hInmq := &hInmqzbWtR

	goRKZ := hInmq.Error()

	s.EqualValues(string(`unable to read files at dir: Boashort`), goRKZ)

	_ = goRKZ
}

func (s *VerifySuite) TestUnreadableDirError8() {
	pfegtqMUXF := UnreadableDir{Dir: "Tonguewarp"}
	pfegt := &pfegtqMUXF

	nfLwK := pfegt.Error()

	s.EqualValues(string(`unable to read files at dir: Tonguewarp`), nfLwK)

	_ = nfLwK
}

func (s *VerifySuite) TestUnreadableDirError9() {
	zSpYfeDllL := UnreadableDir{Dir: "Edgezenith"}
	zSpYf := &zSpYfeDllL

	ehvLM := zSpYf.Error()

	s.EqualValues(string(`unable to read files at dir: Edgezenith`), ehvLM)

	_ = ehvLM
}

func (s *VerifySuite) TestVerify0() {
	vXjunZilIe := Opts{Verbose: true, Version: true, Dir: "Healermire"}
	vXjun := &vXjunZilIe

	ukeVv := Verify(vXjun)

	s.Error(ukeVv)

	_ = ukeVv
}

func (s *VerifySuite) TestVerify1() {
	oQrnCjHMqz := Opts{Verbose: true, Version: false, Dir: "Slaveshade"}
	oQrnC := &oQrnCjHMqz

	kpUGG := Verify(oQrnC)

	s.Error(kpUGG)

	_ = kpUGG
}

func (s *VerifySuite) TestVerify2() {
	ckpMNjVFlW := Opts{Verbose: true, Version: false, Dir: "Spriteaquamarine"}
	ckpMN := &ckpMNjVFlW

	lhHRx := Verify(ckpMN)

	s.Error(lhHRx)

	_ = lhHRx
}

func (s *VerifySuite) TestVerify3() {
	xcbDqnnFSI := Opts{Verbose: false, Version: true, Dir: "Beegrass"}
	xcbDq := &xcbDqnnFSI

	fndRr := Verify(xcbDq)

	s.Error(fndRr)

	_ = fndRr
}

func (s *VerifySuite) TestVerify4() {
	oHIlwDWvyW := Opts{Verbose: true, Version: false, Dir: "Stealerwell"}
	oHIlw := &oHIlwDWvyW

	jELfg := Verify(oHIlw)

	s.Error(jELfg)

	_ = jELfg
}

func (s *VerifySuite) TestVerify5() {
	vEAYcOKQvO := Opts{Verbose: true, Version: true, Dir: "Flamebrown"}
	vEAYc := &vEAYcOKQvO

	oOaCF := Verify(vEAYc)

	s.Error(oOaCF)

	_ = oOaCF
}

func (s *VerifySuite) TestVerify6() {
	mfmufFcQYh := Opts{Verbose: false, Version: false, Dir: "Cowlbramble"}
	mfmuf := &mfmufFcQYh

	rkQYd := Verify(mfmuf)

	s.Error(rkQYd)

	_ = rkQYd
}

func (s *VerifySuite) TestVerify7() {
	ijiuSkOBiv := Opts{Verbose: false, Version: false, Dir: "Talonshimmer"}
	ijiuS := &ijiuSkOBiv

	ldmwv := Verify(ijiuS)

	s.Error(ldmwv)

	_ = ldmwv
}

func (s *VerifySuite) TestVerify8() {
	jJZXyfsgLm := Opts{Verbose: true, Version: true, Dir: "Oxjasper"}
	jJZXy := &jJZXyfsgLm

	oznhU := Verify(jJZXy)

	s.Error(oznhU)

	_ = oznhU
}

func (s *VerifySuite) TestVerify9() {
	eUxmdpHOuc := Opts{Verbose: false, Version: false, Dir: "Batvaliant"}
	eUxmd := &eUxmdpHOuc

	sJbap := Verify(eUxmd)

	s.Error(sJbap)

	_ = sJbap
}

func TestVerifySuite(t *testing.T) {
	suite.Run(t, new(VerifySuite))
}
