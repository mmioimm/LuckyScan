package main

import (
	"flag"
	"strings"
	"fmt"
	"strconv"
	"sync"
	"time"
	"net"
	"io/ioutil"
	"gopkg.in/yaml.v2"

	"github.com/fatih/color"

	"mysql"
	"mssql"
	"smb"
	"ssh"
	"ftp"
)

type MySql struct {
	Port     int    `yaml:"Port"`
	User     string `yaml:"UserDic"`
	Password string `yaml:"PassDic"`
}

type MsSql struct {
	Port	 int	`yaml:"Port"`
	User	 string	`yaml:"UserDic"`
	Password string `yaml:"PassDic"`
}

type Smb struct {
	Port	 int	`yaml:"Port"`
	User	 string	`yaml:"UserDic"`
	Password string `yaml:"PassDic"`
}

type Ssh struct {
	Port	 int	`yaml:"Port"`
	User	 string	`yaml:"UserDic"`
	Password string `yaml:"PassDic"`
}

type Ftp struct {
	Port	 int	`yaml:"Port"`
	User	 string	`yaml:"UserDic"`
	Password string `yaml:"PassDic"`
}

type Config struct{
	SetMySql MySql `yaml:"MySql"`
	SetMsSql MsSql `yaml:"MsSql"`
	SetSmb Smb `yaml:"Smb"`
	SetSsh Ssh `yaml:"Ssh"`
	SetFtp Ftp `yaml:"Ftp"`
}

func CommonScan(ScanType string, target string, setting Config){
	ScanType = strings.ToUpper(ScanType)
	if strings.Contains(ScanType, ","){
		Scanlist := strings.Split(ScanType, ",")
		var wg sync.WaitGroup
		for _, i := range Scanlist{
			wg.Add(1)
			go func(i string){
				defer wg.Done()
				PassScan(i, target, setting)
			}(i)
		}
		wg.Wait()
	} else{
		PassScan(ScanType, target, setting)
	}
}



func PassScan(ScanType string, target string, setting Config){
	//if strings.Contains(ScanType, ","){
	//	CommonScan(ScanType, target, setting)
	//}
	//ScanType = strings.ToUpper(ScanType)
	if ScanType == "MYSQL"{
		port := setting.SetMySql.Port
		user := setting.SetMySql.User
		pass := setting.SetMySql.Password
		mysql.MysqlScan(target, port, user, pass)
	} else if ScanType == "MSSQL"{
		port := setting.SetMsSql.Port
		user := setting.SetMsSql.User
		pass := setting.SetMsSql.Password
		mssql.MssqlScan(target, port, user, pass)
	} else if ScanType == "SMB"{
		port := setting.SetSmb.Port
		user := setting.SetSmb.User
		pass := setting.SetSmb.Password
		smb.SmbScan(target, port, user, pass)
	} else if ScanType == "MS17010"{
		smb.MS17010(target, 3)
	} else if ScanType == "SSH"{
		port := setting.SetSsh.Port
		user := setting.SetSsh.User
		pass := setting.SetSsh.Password
		ssh.SshScan(target, port, user, pass)
	} else if ScanType == "FTP"{
		port := setting.SetFtp.Port
		user := setting.SetFtp.User
		pass := setting.SetFtp.Password
		ftp.FtpScan(target, port, user, pass)
	}
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func main(){
	color.Yellow("https://github.com/mmioimm\n")
	target := flag.String("h", "127.0.0.1", "127.0.0.1")
	ScanType := flag.String("s", "mysql", "mysql/mssql/ftp/smb/ms17010/ssh")
	flag.Parse()
	color.Green("\nTarget: " + *target)
	color.Green("\nScanType: " + *ScanType + "\n")
	var setting Config
	config, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		fmt.Println(err)
	}
	yaml.Unmarshal(config, &setting)
	fmt.Println("\nStarting at: "+time.Now().Format("2006-01-02 15:04:05") + "\n")
	fmt.Println("-----------------------------------------------------------------------------")
	//切割ip段
	if strings.Contains(*target, "-") && strings.Contains(*target, "."){
		CRange := strings.Split(*target, "-")
		SplitIp := strings.Split(CRange[0], ".")
		Cip := SplitIp[0] + "." +  SplitIp[1] + "." + SplitIp[2]
		BeginIp := SplitIp[3]
		EndIp := CRange[1]
		IpStart, err := strconv.Atoi(BeginIp)
		IpEnd, err := strconv.Atoi(EndIp)
		if err != nil{
			fmt.Println(err)
		}
		var wg sync.WaitGroup
		for i:=IpStart;i<=IpEnd;i++{
			ip := Cip + "." + strconv.Itoa(i)
			wg.Add(1)
			go func(ip string){
				defer wg.Done()
				CommonScan(*ScanType, ip, setting)
			}(ip)
		}
		wg.Wait()
	} else if strings.Contains(*target, "/"){
		ip, ipNet, err := net.ParseCIDR(*target)
		if err != nil{
			fmt.Println(err)
			return
		}
		var wg sync.WaitGroup
		for ip:=ip.Mask(ipNet.Mask);ipNet.Contains(ip);incIP(ip){
			wg.Add(1)
			go func(ip string){
				defer wg.Done()
				CommonScan(*ScanType, ip, setting)
			}(ip.String())
		}
		wg.Wait()
	} else {
		CommonScan(*ScanType, *target, setting)
	}
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("\nAttack finished: "+time.Now().Format("2006-01-02 15:04:05"))
}