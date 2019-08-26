package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)
var DB *sqlx.DB

func init(){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"root","123456", "192.168.0.199", 3340, "gj")
	var(err error)
	DB,err=sqlx.Open("mysql",dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	DB.Ping()
}

func main(){
	type Tables struct {
		Tables_in_gj string
	}

	rows,err:= DB.Query("show tables")
	fmt.Println(err)
	//--------遍历放入map----start
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
				table:=string(col.([]byte))
				fmt.Println(string(col.([]byte)))
				BackupMySqlDb("192.168.0.199","3340", "root","123456",  "gj",table,"")
			}
		}
		//fmt.Println(record)
	}


}
func BackupMySqlDb(host, port, user, password, databaseName, tableName, sqlPath string) (error,string)  {
	var cmd *exec.Cmd

	if tableName == "" {
		cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName)
	} else {
		cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName, tableName)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return err,""
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return err,""
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
		return err,""
	}
	now := time.Now().Format("20060102150405")
	var backupPath string
	if tableName == "" {
		backupPath = sqlPath+databaseName+"_"+now+".sql"
	} else {
		backupPath = sqlPath+databaseName+"_"+tableName+"_"+now+".sql"
	}
	err = ioutil.WriteFile(backupPath, bytes, 0644)

	if err != nil {
		panic(err)
		return err,""
	}
	return nil,backupPath
}
