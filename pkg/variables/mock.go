package variables

// nolint: gochecknoglobals
var (
	randomVarNames = []string{
		"alfa",
		"bravo",
		"charlie",
		"delta",
		"echo",
		"foxtrot",
		"golf",
		"hotel",
		"india",
		"juliett",
		"kilo",
		"lima",
		"mike",
		"november",
		"oscar",
		"papa",
		"quebec",
		"romeo",
		"sierra",
		"tango",
		"uniform",
		"victor",
		"whiskey",
		"xray",
		"yankee",
		"zulu",
	}
)

// NewMock creates new mock implementation of var name generator
func NewMock() IGen {
	return &GenMock{
		Index: 0,
	}
}

// GenMock mocks the var name generator
type GenMock struct {
	Index int
}

// Generate returns a mocked var name
func (g *GenMock) Generate() string {
	return randomVarNames[0]
}
