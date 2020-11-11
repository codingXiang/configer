package configer

import (
	"bytes"
	"github.com/spf13/viper"
	"strings"
)

type Core struct {
	data *bytes.Buffer
	core *viper.Viper
}

// 參數依序為：
/// 1. 設定檔類型 (支援 yaml、yml、json、properties、ini、hcl、toml)
/// 2. 檔案名稱 (例如檔名為 config.yaml 就輸入 config)
/// 3. 後續皆為檔案路徑，可以支援多個路徑尋找檔案
func NewCore(configType FileType, configName string, paths ...string) *Core {
	handler := NewCoreWithData(configType, nil)
	if configName != "" {
		handler.SetConfigName(configName)
		for _, path := range paths {
			handler.AddConfigPath(path)
		}
	}

	return handler
}

func NewCoreWithData(configType FileType, data []byte) *Core {
	c := new(Core)
	c.core = viper.New()
	c.SetConfigType(configType.String())
	if data != nil {
		c.data = bytes.NewBuffer(data)
	}
	return c
}

func (c *Core) SetAutomaticEnv(prefix, oldStr, newStr string) {
	if prefix != "" {
		c.GetConfig().SetEnvPrefix(prefix)
	}
	c.GetConfig().SetEnvKeyReplacer(strings.NewReplacer(oldStr, newStr))
	c.GetConfig().AutomaticEnv()
}

func (c *Core) SetConfigType(in string) {
	c.GetConfig().SetConfigType(in)
}

func (c *Core) SetConfigName(in string) {
	c.GetConfig().SetConfigName(in)
}

func (c *Core) AddConfigPath(in string) {
	c.GetConfig().AddConfigPath(in)
}

func (c *Core) GetConfig() *viper.Viper {
	return c.core
}

func (c *Core) SetDefault(key string, value interface{}) *Core {
	c.GetConfig().SetDefault(key, value)
	return c
}

func (c *Core) WriteConfig() error {
	return c.core.SafeWriteConfig()
}

func (c *Core) WriteConfigAs(path string) error {
	return c.core.SafeWriteConfigAs(path)
}

func (c *Core) ReadConfig() (*viper.Viper, error) {
	if c.data != nil {
		return c.GetConfig(), c.GetConfig().ReadConfig(c.data)
	} else {
		return c.GetConfig(), c.GetConfig().ReadInConfig()
	}
}
