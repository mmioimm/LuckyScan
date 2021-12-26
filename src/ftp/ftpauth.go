package ftp

import (
	"port"
	"dic"
	"goftp"

	"fmt"
	"strings"
	"strconv"
	"github.com/fatih/color"
)

func FtpAuth(ip string, port int, user string, pass string)(result bool, err error){
	result = false
	var Lftp *goftp.FTP
	if Lftp, err = goftp.Connect(ip + ":" + strconv.Itoa(port)); err != nil{

	}
	defer Lftp.Close()
	if err = Lftp.Login(user, pass); err == nil{
		result = true
	}
	return result, err
}

func FtpScan(target string, ConfigPort int, ConfigUser string, ConfigPass string){
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
						res, err := FtpAuth(target, ConfigPort, user, password)
						if res == true && err == nil{
							//color.Green("mysql-------" + target + "-------" + strconv.Itoa(ConfigPort) + "-------" + user + "-------" + password)
							fmt.Print("[")
							color.Green(strconv.Itoa(ConfigPort))
							fmt.Print("]")
							fmt.Print("[")
							color.Green("ftp")
							fmt.Print("] host: ")
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
