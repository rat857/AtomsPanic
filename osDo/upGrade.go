package osDo

import (
	"github.com/fatih/color"
	"os/exec"
)

func getBashCommand() string {
	userName, userUid := Getuser()
	switch userUid {
	case 0:
		return "cd /" + userName + "/.local/ && rm -rf nuclei-templates && git clone https://github.com/projectdiscovery/nuclei-templates.git && cd -"
	default:
		return "cd /home/" + userName + "/.local/ && rm -rf nuclei-templates&& git clone https://github.com/projectdiscovery/nuclei-templates.git && cd -"
	}
}
func UpGrade() { //更新
	cmd := exec.Command("bash", "-c", getBashCommand())
	err := cmd.Run()
	if err != nil {
		color.Red("Failed to execute command: %v,%s", err, "出错")
	}
}
