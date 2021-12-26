package dic

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

//判断文件是否存在
func FileIsExit(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

//打开user字典
func UserDic(file string) []string{
	var userlist []string
	userfile, err := os.Open(file)
	if err != nil{
		fmt.Println(err)
	}
	defer userfile.Close()
	scanner := bufio.NewScanner(userfile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan(){
		user := strings.TrimSpace(scanner.Text())
		if user != ""{
			userlist = append(userlist, user)
		}
	}
	return userlist
}