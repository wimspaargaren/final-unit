package importcustomtype

import "github.com/sirupsen/logrus"

func Level() logrus.Level {
	return logrus.InfoLevel
}

type Something byte

const SomeSomething Something = Something(byte(4))

func Test() Something {
	return SomeSomething
}
