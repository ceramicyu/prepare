package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
)
var  Controller string
func init() {
	flag.StringVar(&Controller,"c","","控制器")
}
func main()  {
	//Controller="Bed"
	flag.Parse()
	fileName:=`C:\Users\Administrator\Desktop\star.bat`
	str:=""
	fmt.Println(Controller)
	if Controller=="" {
		Controller="Bed"
	}
	str = fmt.Sprintf(`start "D:\Program Files\PhpStorm 2017.2.4\bin\phpstorm64.exe" 
	"D:\afu\data\www\gj\appserv\modules\v1\%sController.class.php" `,Controller)
	err:=ioutil.WriteFile(fileName,[]byte(str),0777)
	fmt.Println(err)
	exec.Command("cmd.exe", "/c",fileName ).Output()
}
