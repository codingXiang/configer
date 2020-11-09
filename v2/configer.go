package configer

type Configer struct {
	handler map[string]*Core
}

var (
	//Config : 設定檔變數
	Config *Configer
)

func init() {
	Config = NewConfiger()
}

//NewConfiger 初始化
func NewConfiger() *Configer {
	var config = &Configer{}
	config.handler = make(map[string]*Core)
	return config
}

//GetCore : 取得組態控制器
func (c *Configer) GetCore(key string) *Core {
	// 檢查 key 是否存在
	if value, ok := c.handler[key]; ok {
		return value
	}
	return nil
}

//AddCore : 加入組態控制器
func (c *Configer) AddCore(key string, handler *Core) *Configer {
	c.handler[key] = handler
	return c
}
