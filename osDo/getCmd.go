package osDo

import (
	"flag"
)

func Getcmd() (string, string, string, string) { //获取cmd命令
	url := flag.String("u", "nul", "指定的url(可以多个url,中间用逗号隔开,Eg:http://www.github.com,http://www.rat857.top)")
	urlFile := flag.String("f", "nul", "指定一个存放url的文件,Eg:/root/Download/target.txt")
	proxy := flag.String("proxy", "nul", "指定代理服务器Eg:socks5:127.0.0.1:7890")
	pocYaml := flag.String("poc", "nul", "指定AtomsPanic格式的poc的yaml的文件进行扫描(可以是多个，同上，用逗号隔开)")
	flag.Parse()
	return *url, *urlFile, *proxy, *pocYaml

}
