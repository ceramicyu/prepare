package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var strArr = []string{
	` 'host'     => 'mysql-master-3330',
        'username' => 'root',
        'password' => '123456',
        'port'     => '3306',`,
	` 'host'     => 'mysql-master-3340',
        'username' => 'root',
        'password' => '123456',
        'port'     => '3306',`,
	` 'host'     => 'mysql-master-3320',
        'username' => 'root',
        'password' => '123456',
        'port'     => '3306',`,
}
var newStr = []string{
	"", "", "",
}
var Addr = ""
var(
	h string
	sim bool
)
func init() {
	flag.StringVar(&h, "h", "3", "地址")
	flag.BoolVar(&sim, "sim", false, "分支")
}

func main() {

	flag.Parse()
	fmt.Println(h)
	fileName:="/data/www/fm/conf/conf/DBConfig.class.php"
	if (sim){
		fmt.Println("sim 分支")
		fileName="/data/www_sim/fm/conf/conf/DBConfig.class.php"
	}

	fmt.Println("找到数据库配置文件：",fileName)



		Addr = strings.Trim(h, "")
		if Addr == "3" {
			ioutil.WriteFile(fileName, []byte(Config3), 0777)
			fmt.Println("mysql-master-3330",3306)
			fmt.Println("mysql-master-3340",3306)
			fmt.Println("mysql-master-3320",3306)
		}
		if Addr == "4" {
			ioutil.WriteFile(fileName, []byte(Config4), 0777)
			fmt.Println("mysql-master-4330")
			fmt.Println("mysql-master-4340")
			fmt.Println("mysql-master-4320")
		}
		if len(Addr) > 2 {

			b, _ := ioutil.ReadFile(fileName)
			str := string(b)
			for{
				reg := regexp.MustCompile(`mysql-master-(\d+)`)
				matched := reg.FindString(str)
				if len(matched)==0{
					break
				}
				reg2 := regexp.MustCompile(`(\d+)`)
				matched2 := reg2.FindString(matched)

				str = strings.Replace(str, matched, Addr, 2)
				str = strings.Replace(str, "3306", matched2, 2)
				fmt.Println(matched,"=>",Addr,"    ",3306,"=>",matched2)
			}


			ioutil.WriteFile(fileName, []byte(str), 0777)
		}


}




var Config3=`<?php

/**
 * 数据库配置类
 * @date   2015年1月13日 下午4:59:19
 */

namespace fm\conf\conf;

class DBConfig {
    
    public static function getServer($project, $isMaster = false) {
        switch ($project) {
            case 'afu':
                return $isMaster
                    ? self::$MASTER_2
                    : self::$SLAVE_2;
                break;
            
            case 'fm':
            case 'oa':
            case 'sso':
            case 'kf':
            case 'ehelper':
            case 'stat':
            case 'supermarket':
                return $isMaster
                    ? self::$MASTER_1
                    : self::$SLAVE_1;
                break;
            
            case 'bb':
            case 'yunsong':
                return $isMaster
                    ? self::$MASTER_3
                    : self::$SLAVE_3;
                break;
            
            default:
                return $isMaster
                    ? self::$MASTER_2
                    : self::$SLAVE_2;
                break;
        }
    }
    
    public static $MASTER_1 = [
        'host'     => 'mysql-master-3330',
        'username' => 'root',
        'password' => '123456',
        'port'     => '3306',
    ];
    public static $SLAVE_1  = [
        [
            'host'     => 'mysql-master-3330',
            'username' => 'root',
            'password' => '123456',
            'port'     => '3306',
        ]
    ];
    
    public static $MASTER_2 = [
        'host'     => 'mysql-master-3340',
        'username' => 'root',
        'password' => '123456',
        'port'     => '3306',
    ];
    public static $SLAVE_2  = [
        [
            'host'     => 'mysql-master-3340',
            'username' => 'root',
            'password' => '123456',
            'port'     => '3306',
        ]
    ];
    
    public static $MASTER_3 = [
        'host'     => 'mysql-master-3320',
        'username' => 'root',
        'password' => '123456',
        'port'     => '3306',
    ];
    public static $SLAVE_3  = [
        [
            'host'     => 'mysql-master-3320',
            'username' => 'root',
            'password' => '123456',
            'port'     => '3306',
        ]
    ];
    
}`

var Config4=`<?php

/**
 * 数据库配置类
 * @date   2015年1月13日 下午4:59:19
 */

namespace fm\conf\conf;

class DBConfig {
    
    public static function getServer($project, $isMaster = false) {
        switch ($project) {
            case 'afu':
                return $isMaster
                    ? self::$MASTER_2
                    : self::$SLAVE_2;
                break;
            
            case 'fm':
            case 'oa':
            case 'sso':
            case 'kf':
            case 'ehelper':
            case 'stat':
            case 'supermarket':
                return $isMaster
                    ? self::$MASTER_1
                    : self::$SLAVE_1;
                break;
            
            case 'bb':
            case 'yunsong':
                return $isMaster
                    ? self::$MASTER_3
                    : self::$SLAVE_3;
                break;
            
            default:
                return $isMaster
                    ? self::$MASTER_2
                    : self::$SLAVE_2;
                break;
        }
    }
    
    public static $MASTER_1 = [
        'host'     => 'mysql-master-4330',
        'username' => 'root',
        'password' => '123456',
        'port'     => '3306',
    ];
    public static $SLAVE_1  = [
        [
            'host'     => 'mysql-master-4330',
            'username' => 'root',
            'password' => '123456',
            'port'     => '3306',
        ]
    ];
    
    public static $MASTER_2 = [
        'host'     => 'mysql-master-4340',
        'username' => 'root',
        'password' => '123456',
        'port'     => '3306',
    ];
    public static $SLAVE_2  = [
        [
            'host'     => 'mysql-master-4340',
            'username' => 'root',
            'password' => '123456',
            'port'     => '3306',
        ]
    ];
    
    public static $MASTER_3 = [
        'host'     => 'mysql-master-4320',
        'username' => 'root',
        'password' => '123456',
        'port'     => '3306',
    ];
    public static $SLAVE_3  = [
        [
            'host'     => 'mysql-master-4320',
            'username' => 'root',
            'password' => '123456',
            'port'     => '3306',
        ]
    ];
    
}`