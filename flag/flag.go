package flag

import "flag"

var (
	PathConfFile 	= flag.String("c","./config.toml","指定配置文件")
)

//解析命令行
func ParseFlag(){
	flag.Parse() //解析命令行参数
}


