package model

type DBConf struct {
	//驱动
	Driver string
	//主机地址
	Host string
	//Port 主机端口
	Port string
	//用户名
	User string
	//密码
	Password string
	//DbName
	DbName string
	//Charset
	Charset string
}
