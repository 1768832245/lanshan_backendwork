package conf

type MysqlConf struct {
	Host     string `yaml:"host"`   //localhost
	Port     string `yaml:"port"`   //3306
	Config   string `yaml:"config"` //高级配置，几乎固定
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Loglevel string `yaml:"loglevel"` //日志等级
}

type Config struct {
	Mysql MysqlConf `json:"mysql"`
}
