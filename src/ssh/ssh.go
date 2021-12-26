package ssh

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"port"
	"dic"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
)

func SshAuth(ip string, port int, user string, pass string)(result bool, err error){
	result = false
	authMethods := []ssh.AuthMethod{}

    keyboardInteractiveChallenge := func(
        user,
        instruction string,
        questions []string,
        echos []bool,
    ) (answers []string, err error) {
        if len(questions) == 0 {
            return []string{}, nil
        }
        return []string{pass}, nil
    }

    authMethods = append(authMethods, ssh.KeyboardInteractive(keyboardInteractiveChallenge))
    authMethods = append(authMethods, ssh.Password(pass))

    sshConfig := &ssh.ClientConfig{
        User: user,
        Auth: authMethods,
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
		},
    }

	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", ip, port), sshConfig)
	if err == nil {
		defer client.Close()
		session, err := client.NewSession()
		errRet := session.Run("echo ISOK")
		if err == nil && errRet == nil {
			defer session.Close()
			result = true
		}
	}
	return result,err
}


func SshScan(target string, ConfigPort int, ConfigUser string, ConfigPass string){
	if port.PortCheck(target, ConfigPort){
		if dic.FileIsExit(ConfigUser){
			if dic.FileIsExit(ConfigPass){
				fmt.Println("Attacking... " + target)
				Loop:
				for _, user := range dic.UserDic(ConfigUser){
					for _, password := range dic.UserDic(ConfigPass){
						//fmt.Println("Check..." + target + " " + user + " " + password)
						if strings.Contains(password, "%user%"){
							password = strings.Replace(password, "%user%", user, -1)
						}
						res, err := SshAuth(target, ConfigPort, user, password)
						if res == true && err == nil{
							//color.Green("mysql-------" + target + "-------" + strconv.Itoa(ConfigPort) + "-------" + user + "-------" + password)
							fmt.Print("[")
							color.Green(strconv.Itoa(ConfigPort))
							fmt.Print("]")
							fmt.Print("[")
							color.Green("ssh")
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