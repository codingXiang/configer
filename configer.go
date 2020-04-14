package configer

import "github.com/spf13/viper"

type (
	ConfigerInterface interface {
		GetCore(key string) CoreInterface
		AddCore(key string, handler CoreInterface) ConfigerInterface
	}
	CoreInterface interface {
		SetAutomaticEnv()
		SetConfigType(in string)
		SetConfigName(in string)
		AddConfigPath(in string)
		ReadConfig() (*viper.Viper, error)
	}

	//Configer : 整體設定檔
	Configer struct {
		handler map[string]CoreInterface
	}

	Core struct {
		core *viper.Viper
	}
)

var (
	//Config : 設定檔變數
	Config ConfigerInterface
)

// 參數依序為：
/// 1. 設定檔類型 (支援 yaml、yml、json、properties、ini、hcl、toml)
/// 2. 檔案名稱 (例如檔名為 config.yaml 就輸入 config)
/// 3. 後續皆為檔案路徑，可以支援多個路徑尋找檔案
func NewConfigerCore(configType string, configName string, paths ...string) CoreInterface {
	var handler = &Core{
		core: viper.New(),
	}
	handler.SetConfigType(configType)
	handler.SetConfigName(configName)
	for _, path := range paths {
		handler.AddConfigPath(path)
	}
	return handler
}

//NewConfiger 初始化
func NewConfiger() ConfigerInterface {
	var config = &Configer{}
	config.handler = map[string]CoreInterface{}
	return config
}

//GetCore : 取得組態控制器
func (this *Configer) GetCore(key string) CoreInterface {
	// 檢查 key 是否存在
	if value, ok := this.handler[key]; ok {
		return value
	}
	return nil
}

//AddCore : 加入組態控制器
func (this *Configer) AddCore(key string, handler CoreInterface) ConfigerInterface {
	this.handler[key] = handler
	return this
}
func (c *Core) SetAutomaticEnv() {
	c.getCore().AutomaticEnv()
}

func (c *Core) SetConfigType(in string) {
	c.getCore().SetConfigType(in)
}

func (c *Core) SetConfigName(in string) {
	c.getCore().SetConfigName(in)
}

func (c *Core) AddConfigPath(in string) {
	c.getCore().AddConfigPath(in)
}

func (c *Core) getCore() *viper.Viper {
	return c.core
}

func (c *Core) ReadConfig() (*viper.Viper, error) {
	if err := c.getCore().ReadInConfig(); err == nil {
		return c.getCore(), nil
	} else {
		panic("read config error")
		return nil, err
	}
}
