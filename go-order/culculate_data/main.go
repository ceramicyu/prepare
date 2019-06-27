package main

import "fmt"

func main(){

	//获取项目列表 projectList
	projectList:=GetProjectList()
	for _,projectInfo :=range projectList{
		fmt.Println(projectInfo)
	}

}
