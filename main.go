package main

import (
	"AtomsPanic/osDo"
	"AtomsPanic/scan"
)

func main() {
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

	/*	//测试templates可用性
		res, err := os.ReadFile("../test.yaml")
		if err != nil {
			log.Println(err)
		}
		var data templates.Req
		yaml.Unmarshal(res, &data)
		fmt.Println(data.Packets[0].Matchers[0].Regex)*/
	osDo.Title()
	//fmt.Println(osDo.ParseCmd())
	/*for i, s := range scan.GetUrl() {
		fmt.Println(i)
		fmt.Println(s)
	}*/
	//fmt.Println(scan.GetProxy())
	//fmt.Println(len(scan.GetProxy()))
	osDo.WriteListTxt(scan.Scan(), "good.txt")
}
