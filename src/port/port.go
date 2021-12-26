package port

import (
	"net"
	"fmt"
	"strconv"
	"time"
)

//检测端口开放情况
func PortCheck(host string, port int)(result bool){
	result = false
	ip := net.ParseIP(host)
	tcpAddr := net.TCPAddr{
		IP: ip,
		Port: port,
	}
	addr := host + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if conn != nil{
		fmt.Println(tcpAddr.IP, tcpAddr.Port, "Open")
		conn.Close()
		result = true
	}
	if err != nil{
		//fmt.Println(err)
	}
	return result

}