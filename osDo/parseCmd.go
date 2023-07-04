package osDo

import (
	"strings"
)

func split(cmd string) []string {
	return strings.Split(cmd, ",")
}
func ParseCmd() (urlList, pocYamlList []string, proxy, urlFile string) { //解析cmd命令
	url, urlFile, proxy, pocYaml := Getcmd()
	urlList = split(url)
	pocYamlList = split(pocYaml)
	return urlList, pocYamlList, proxy, urlFile
}
