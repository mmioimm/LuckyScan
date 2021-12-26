package smb

import (
	"github.com/stacktitan/smb/smb"
	"port"
	"dic"
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

func SmbAuth(ip string, port int, user string, pass string)(result bool, err error){
	result = false
	options := smb.Options{
		Host: ip,
		Port: port,
		User: user,
		Password: pass,
		Domain: "",
		Workstation: "",
	}
	session, err := smb.NewSession(options, false)
	if err == nil{
		session.Close()
		if session.IsAuthenticated{
			result = true
		}
	}
	return result, err
}


func SmbScan(target string, ConfigPort int, ConfigUser string, ConfigPass string){
	if port.PortCheck(target, ConfigPort){
		if dic.FileIsExit(ConfigUser){
			if dic.FileIsExit(ConfigPass){
				fmt.Println("Attacking... " + target)
				Loop:
				for _, user := range dic.UserDic(ConfigUser){
					for _, password := range dic.UserDic(ConfigPass){
						if strings.Contains(password, "%user%"){
							password = strings.Replace(password, "%user%", user, -1)
						}
						//fmt.Println("Check..." + target + " " + user + " " + password)
						res, err := SmbAuth(target, ConfigPort, user, password)
						if res == true && err == nil{
							//color.Green("mysql-------" + target + "-------" + strconv.Itoa(ConfigPort) + "-------" + user + "-------" + password)
							fmt.Print("[")
							color.Green(strconv.Itoa(ConfigPort))
							fmt.Print("]")
							fmt.Print("[")
							color.Green("smb")
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