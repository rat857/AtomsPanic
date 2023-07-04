// 解析文件
package osDo

import (
	"AtomsPanic/templates"
	"bufio"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
	"net/url"
	"os"
)

// 按行读取txt文件，并保存在一个List
func ReadTxtList(fileName string) []string { //第一个参数要读取的文件名字
	var urls = make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		color.Red("%v", err)
	}
	a := bufio.NewReader(file)
	for true {
		data, _, error := a.ReadLine()
		if error != nil {
			break
		} else {
			urls = append(urls, string(data))
		}
	}
	return urls
}

// 读取Yaml为templates的Req结构体
func ReadYamlReq(fileName string) templates.Req { //第一个参数是文件名，返回值是templates.Req
	res, err := os.ReadFile(fileName)
	if err != nil {
		color.Red("!!!!打开poc失败%v", err)
		os.Exit(1)
	}
	var data templates.Req
	if err := yaml.Unmarshal(res, &data); err != nil {
		color.Red("!!!解析poc失败")
		os.Exit(2)
	}
	return data
}

// 写入文件（把List按行写入）
func WriteListTxt(resList []string, fileName string) { //第一个参数是写入的url的list,第二个参数生成的文件名
	resList = RemoveDuplicateIPs(resList) // 给ip去重,因为一个ip上可能有几个端口，几个端口上可能都有漏洞
	//创建文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//创建 bufio.Writer
	writer := bufio.NewWriter(file)

	//循环写入文件
	for _, line := range resList {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	//刷新缓存
	writer.Flush()
}

// 给url去重http //128.0.0.1 890
func RemoveDuplicateIPs(urls []string) []string {
	encountered := make(map[string]bool)
	result := []string{}

	for _, urlStr := range urls {
		u, err := url.Parse(urlStr)
		if err != nil {
			continue
		}

		host := u.Hostname()
		if !encountered[host] {
			encountered[host] = true
			result = append(result, urlStr)
		}
	}

	return result
}

// 给切片去重
func RemoveDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
