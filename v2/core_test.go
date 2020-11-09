package configer_test

import (
	"github.com/codingXiang/configer/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	testData = []byte(`application:
  port: 8080`)
)

//CoreSuite struct
type CoreSuite struct {
	suite.Suite
	core *configer.Core
}

//初始化 CoreSuite
func (s *CoreSuite) SetupCoreSuite() {

}

//TestStart 為測試程式進入點
func TestStartCoreSuite(t *testing.T) {
	suite.Run(t, new(CoreSuite))
}

func (s *CoreSuite) BeforeTest(suiteName, testName string) {
	//s.core = configer.NewCoreWithData(testData)
}

func (s *CoreSuite) TestNewCoreWithData() {
	c := configer.NewCoreWithData(testData)
	assert.IsType(s.T(), new(configer.Core), c)
}

func (s *CoreSuite) TestNewCore() {
	c := configer.NewCore(configer.YAML, "config", "./config")
	config, err := c.ReadConfig()
	assert.Nil(s.T(), err, "read config has failed")
	assert.Equal(s.T(), 8888, config.GetInt("application.port"))
}
