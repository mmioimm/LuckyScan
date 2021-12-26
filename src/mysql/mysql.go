package mysql

import (
	"fmt"
	"database/sql"
	"strconv"

	"port"
	"dic"

	"github.com/fatih/color"

	
	_"github.com/go-sql-driver/mysql"
)

func MysqlAuth(ip string, port int, user string, pass string)(result bool, err error){
	result = false
	db, err := sql.Open("mysql", user + ":" + pass + "@tcp(" + ip + ":" + strconv.Itoa(port) + ")/mysql?charset=utf8")
	if err != nil{
		
	}
	if db.Ping() == nil{
		result = true
	}
	return result, err
}

func MysqlScan(target string, ConfigPort int, ConfigUser string, ConfigPass string) {
	if port.PortCheck(target, ConfigPort){
		if dic.FileIsExit(ConfigUser){
			if dic.FileIsExit(ConfigPass){
				fmt.Println("Attacking... " + target)
				Loop:
				for _, user := range dic.UserDic(ConfigUser){
					for _, password := range dic.UserDic(ConfigPass){
						//fmt.Println("Check..." + target + " " + user + " " + password)
						res, err := MysqlAuth(target, ConfigPort, user, password)
						if res == true && err == nil{
							//color.Green("mysql-------" + target + "-------" + strconv.Itoa(ConfigPort) + "-------" + user + "-------" + password)
							fmt.Print("[")
							color.Green(strconv.Itoa(ConfigPort))
							fmt.Print("]")
							fmt.Print("[")
							color.Green("mysql")
							fmt.Print("]\thost: ")
							color.Green(target)
							fmt.Print("\tuser: ")
							color.Green(user)
							fmt.Print("\tpass: ")
							color.Green(password + "\n")
							break Loop
						}
					}
				}
			} else {
				color.Red("pass file not exit!")
			}
		} else {
			color.Red("user file not exit!")
		}
	}
}
