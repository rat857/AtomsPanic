package scan

import (
	"AtomsPanic/osDo"
	"AtomsPanic/templates"
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Result struct {
	URL   string
	Error error
}

func Scan() []string {
	var bodList = make([]string, 0)
	//获取url列表
	urlList, pocYaml, proxy, urlFile := osDo.ParseCmd()
	var GetUrlList = make([]string, 0)
	if urlList[0] != "nul" && urlFile == "nul" {
		GetUrlList = urlList
	} else if urlList[0] == "nul" && urlFile != "nul" {
		GetUrlList = osDo.ReadTxtList(urlFile)
	} else if urlList[0] != "nul" && urlFile != "nul" {
		color.Red("!请只使用一种方式-u或-f")
	} else {
		color.Red("!!!没有检测到任何目标")
	}
	GetUrlList = osDo.RemoveDuplicates(GetUrlList) //url去重
	fmt.Println(GetUrlList)
	//获取url列表终点
	var tr *http.Transport
	proxList := strings.Split(proxy, "://")
	if proxy != "nul" { //判断proxy是否有值
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},                           //防止https证书错误的问题
			Proxy:           http.ProxyURL(&url.URL{Scheme: proxList[0], Host: proxList[1]}), //代理
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //防止https证书错误的问题
		}
	}
	client := &http.Client{Transport: tr}
	var wg sync.WaitGroup
	//fmt.Println(proxy)
	//fmt.Println(tr)
	results := make(chan Result)

	for _, url := range GetUrlList {
		wg.Add(1)
		//go worker(ur, pocYaml, client, results)
		ur := url
		go func() {
			defer wg.Done()
			var loList = make([]string, 0)
			for _, pocYamlName := range pocYaml {
				pocYamlInfo := osDo.ReadYamlReq(pocYamlName)
				for _, pocInfo := range pocYamlInfo.Http.Packets {
					data := pocInfo.Body
					boby := strings.NewReader(data)
					req, _ := http.NewRequest(pocInfo.Method, ur+pocInfo.Path, boby)
					for key, vule := range pocInfo.Head {
						req.Header.Set(key, vule)
					}
					//fmt.Println(pocInfo.Body)
					respChan := make(chan *http.Response)
					errChan := make(chan error)

					go func() {
						resp, err := client.Do(req)
						if err != nil {
							errChan <- err
						} else {
							respChan <- resp
						}
					}()
					select {
					case resp := <-respChan:
						// 处理响应
						// ...
						var mu sync.Mutex
						mu.Lock()

						/*if try(resp, pocYamlInfo.Http) {
							color.Red(ur + "have value")

						}*/

						lo := d(resp, pocInfo)
						loList = append(loList, lo)
						//fmt.Println(loList)
						mu.Unlock()
						resp.Body.Close()
					case err := <-errChan:
						results <- Result{URL: ur, Error: err}
					case <-time.After(10 * time.Second):
						results <- Result{URL: ur, Error: fmt.Errorf("request timeout")}
					}
				}
				//fmt.Println(loList)
				var mu sync.Mutex
				mu.Lock()
				switch pocYamlInfo.Http.Logic {
				case "or":
					if Contains(loList, "01") {
						color.Red(ur + "=====or=====" + "have value")
						bodList = append(bodList, ur)
					} else {
						color.Green(ur + "=====and=====" + "no value")
					}
				default:
					if Contains(loList, "01") && !Contains(loList, "02") {
						color.Red(ur + "=====and=====" + "have value")
						bodList = append(bodList, ur)
					} else {
						color.Green(ur + "=====or=====" + "no value")
					}
				}
				mu.Unlock()
			}
		}()
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	// 处理结果
	for result := range results {
		if result.Error != nil {
			color.Yellow("Error for URL %s: %v\n", result.URL, result.Error)
		}
	}
	return bodList
}
func Contains(slice []string, value string) bool { //判断一个切片里是否有某个值
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func d(resp *http.Response, info templates.Packets) string {
	respBodyByte, _ := io.ReadAll(resp.Body)
	//color.Blue(string(respBodyByte))
	var logic = "0"

	var regexlogic = "0"
	for _, regex := range info.Regex {
		if strings.Contains(string(respBodyByte), regex) {
			//fmt.Println(info.Path + regex)
			//fmt.Println("======")
			regexlogic = regexlogic + "1"
			//fmt.Println(info.Path + "----" + "regexlogic:" + regexlogic)
		} else {
			//color.Red(info.Path + regex)
			//color.Red("======")
			regexlogic = regexlogic + "2"
			//color.Blue(info.Path + "----" + "regexlogic:" + regexlogic)
			//fmt.Println(string(respBodyByte))
			//color.Blue(string(respBodyByte))
		}

	}
	switch info.RegexLogic {
	case "or":
		if strings.Contains(regexlogic, "1") {
			logic = logic + "1"
			//color.Green("logic:" + logic)
		} else {
			logic = logic + "2"
			//color.Green("logic:" + logic)
		}
	default:
		if strings.Contains(regexlogic, "1") && !strings.Contains(regexlogic, "2") {
			logic = logic + "1"
			//fmt.Println("logic:" + logic)
		} else {
			logic = logic + "2"
			//color.Green("logic:" + logic)
		}
	}

	return logic
}
