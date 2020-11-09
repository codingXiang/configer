package configer_test

import (
	"github.com/codingXiang/configer/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	configType = "yaml"
	configName = "test"
	configPath = []string{".", "./config"}
)

var (
)

//Suite struct
type Suite struct {
	suite.Suite
	configer *configer.Configer
}

//初始化 Suite
func (s *Suite) SetupSuite() {

}

//TestStart 為測試程式進入點
func TestStartSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) BeforeTest(suiteName, testName string) {
	s.configer = configer.NewConfiger()
}

func (s *Suite) TestNewConfiger() {
	c := configer.NewConfiger()
	assert.IsType(s.T(), new(configer.Configer), c)
}

func (s *Suite) TestAddCore() {
	core := configer.NewCore(configer.YAML, configName, configPath...)

	const name = "core"
	s.configer.AddCore(name, core)

	assert.Equal(s.T(), core, s.configer.GetCore(name))
}
