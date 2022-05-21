package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/phplaber/finder/pkg/utils"
)

func usage() {
	fmt.Print("Finder v0.1 - A tool for collect Mac installed Apps\n\n")
	fmt.Print("only use three commands: \n- ls\n- whoami\n- hostname\n")
	os.Exit(0)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// 应用列表
	// /Applications
	cmd := "ls -l /Applications | awk '{$1=$2=$3=$4=$5=$6=$7=$8=\"\"; print $0}'"
	apps := utils.ExecCmd(cmd)
	gapplist, _ := utils.String2Lines(apps)

	// ~/Applications
	yacmd := "find ~/Applications -iname  \"*.app\""
	yaapps := utils.ExecCmd(yacmd)
	uapplisttemp, _ := utils.String2Lines(yaapps)
	var uapplist []string
	for _, uapptemp := range uapplisttemp {
		// 截取 Applications 后面的内容
		uapp := strings.Split(uapptemp, "Applications")[1]
		uapplist = append(uapplist, uapp)
	}
	applist := append(gapplist, uapplist...)
	// 账号
	username := utils.ExecCmd("whoami")
	// 计算机名
	hostname := utils.ExecCmd("hostname")

	save(username, hostname, applist)
	fmt.Print("Successful :) \n")
}
