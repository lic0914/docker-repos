package main

import (
	"net"
    "os"
	"os/exec"
	"bytes"
    "fmt"
)
func main(){
	name, _ := os.Hostname()
	fmt.Println(name,"\n")
	ns,err := net.LookupHost("baidu.com")
	if err !=nil {
		fmt.Println(err.Error())
		return
	}
	for _, n:=range ns{
		fmt.Println(n)

	}

	fmt.Println("====")
	bytes := execSample("baidu.com")
	fmt.Print(string(bytes))
	fmt.Println("====")
	bytes = curlSample("baidu.com")
	fmt.Print(string(bytes))
	
}

func execSample(host string) []byte{
	cmd := exec.Command("nslookup", host)
		cmdOutput := &bytes.Buffer{}
		cmd.Stdout = cmdOutput
		err1 := cmd.Run()
		if err1 != nil {
			//os.Stderr.WriteString(err1.Error())
            return []byte(err1.Error())
        }
	//	fmt.Print(string(cmdOutput.Bytes()))
        return cmdOutput.Bytes();
}

func curlSample(host string) []byte{
	out, err := exec.Command("bash","-c","curl --help",host).Output()
	if err != nil {
		return []byte(err.Error())
	}
	
	//	fmt.Print(string(cmdOutput.Bytes()))
	return out;
}