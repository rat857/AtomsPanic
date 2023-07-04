package templates

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Req struct {
	ID   string `yaml:"id"` //类似名字，尽可能是CVE，CNVD编号Eg:CVE-2023-1133
	Info `yaml:"info"`
	Http `yaml:"http"`
}
type Info struct {
	Name         string              `yaml:"name"`          //漏洞名字,Eg:Apache Struts2 S2-008 RCE
	Severity     string              `yaml:"severity"`      //危险程度Eg:critical,high
	Inner_Type   string              `yaml:"type"`          //Eg:SQL注入，文件上传，任意命令执行
	SearchEngine []map[string]string `yaml:"search-engine"` //搜索方法Eg:fofa:app="通达OA网络智能办公系统"
	Link         []string            `yaml:"link"`          //漏洞相关的链接
}
type Http struct {
	Logic   string    `yaml:"logic"`   //逻辑关系,Eg:and,or
	Packets []Packets `yaml:"Packets"` //请求的数据包及验证

}
type Packets struct {
	Method   string            `yaml:"method"` //请求方式，GET，POST
	Path     string            `yaml:"path"`   //数据包的path
	Head     map[string]string `yaml:"head"`   //数据包的head
	Body     string            `yaml:"body"`   //数据包的body
	Matchers `yaml:"matchers"` //验证
}
type Matchers struct {
	//Type       string   `yaml:"type"`        //判断方法，Eg:status,word
	RegexLogic string   `yaml:"regex-logic"` //多个regex时的判断方法
	Regex      []string `yaml:"regex"`
}
type GetHttp struct {
}
type OtherHttp struct {
}

func main() {
	//格式测试写入
	var test Req
	test.ID = "CVE-2023-1133"
	test.Name = "Apache Struts2 S2-008 RCE"
	test.Severity = "high"
	test.Inner_Type = "SQL注入"
	//test.SearchEngine = []string{"fofa:app=\"通达OA网络智能办公系统\"", "fofa: icon_hash=\"1420424513\""}
	test.SearchEngine = []map[string]string{{"fofa": "app=\\\"通达OA网络智能办公系统\\\""}, {"shodan": "windows7"}}
	test.Link = []string{"http://wiki.peiqi.tech/wiki/serverapp/VMware/VMware%20Workspace%20ONE%20Access%20SSTI%E6%BC%8F%E6%B4%9E%20CVE-2022-22954.html", "http://wiki.peiqi.tech/wiki/cms/DocCMS/DocCMS%20keyword%20SQL%E6%B3%A8%E5%85%A5%E6%BC%8F%E6%B4%9E.html"}

	//test.Path = []string{"/?M_id=1%27&type=product"}
	test.Packets = []Packets{{Method: "GET", Path: "/?M_id=1%27&type=product", Head: map[string]string{"Useragent": "justdo", "Type": "notype"}, Matchers: Matchers{RegexLogic: "and", Regex: []string{"success", "status", "200"}}},
		{Method: "POST", Body: "passwd=123.com", Path: "/1.php", Matchers: Matchers{Regex: []string{"200"}}}}
	test.Logic = "and"
	//test.Matchers = {Type: "status", Matchers_type: Matchers_type{Status: []int{200, 301}, Word: []string{"success", "login"}}
	//test.Matchers = []Matchers{{Type: "word", Regex: []string{"success", "status"}}, {Type: "status", Regex: []string{"200", "301"}}}
	data, _ := yaml.Marshal(test)
	os.WriteFile("test.yaml", data, 0666)
	/*	//格式测试读取
		res, _ := os.ReadFile("test.yaml")
		var data Req
		yaml.Unmarshal(res, &data)
		fmt.Println(data.Packets[0].Matchers)*/
}
