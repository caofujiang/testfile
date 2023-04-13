package main

import (
	"fmt"
	"log"
	"os/exec"
	"path"
	"strings"
)

func main() {

	//var out bytes.Buffer

	filePath := "C:\\test-script/2.bat"
	//filePath := "C:\\test-script/hello.ps1"
	params := []string{"1111111", "fs2222", "333333", "44455"}

	paramsStr := strings.Join(params, ",")
	surfix := path.Ext(filePath)

	//要写绝对路径  D:\51cto_course\51ctowkxz\51cto_video_download_20210110\video_ts_file\decode.bat
	//cmd := exec.Command("cmd.exe", "/C", `C:\test-script/echo.bat`, "qqq1111111")
	//cmd := exec.Command("cmd.exe", "/C", `C:\test-script/input.bat`)
	//cmd := exec.Command("cmd.exe", "/C", `C:\test-script/1.bat`, "qqq1111111", "dfsdffs")
	//cmd := exec.Command("cmd.exe", "/C", `C:\test-script/2.bat`, "qqq1111111", "dfsdffs", "eferewf", "212323")
	//cmd := exec.Command("powershell.exe", "/C", `C:\test-script/hello.ps1`, "qqq1111111", "dfdsfsd")

	var cmd *exec.Cmd
	if surfix == ".bat" {
		cmd = exec.Command("cmd.exe", "/C", filePath, paramsStr)
	} else if surfix == ".ps1" {
		fileParams := make([]string, 0)
		fileParams = append(fileParams, filePath)
		fileParams = append(fileParams, params...)
		cmd = exec.Command("powershell.exe", fileParams...)
	} else {
		fmt.Println("file format is wrong,please check")
	}

	//方式1：
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//fmt.Printf("%s", out.String())
	//方式2
	result, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s", result)

}
