package main

import (
    "net"
    "net/http"
    "io/ioutil"
    "os/exec"
	"bytes"
    "log"
    "os"
    "fmt"
    "os/signal"
    "syscall"
    "strings"
    "time"
    //"strconv"
    "github.com/joho/godotenv"
)

var err = godotenv.Load(".env")

var (
    Version = os.Getenv("VERSION")
    BASE_URL = os.Getenv("BASE_URL")
    PORT = os.Getenv("PORT")
 )

func main() {

    if err != nil {
        log.Fatalf("Error loading .env file")
    }
   
    // 创建监听退出 chan
    c := make(chan os.Signal)
    // 监听指定信号 ctrl+c kill ...
    signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

    // 开启协程监听信号
    go func() {
        for s := range c {
            // 简单点，不判断信号类型了，收到信号直接退出
            switch s {
            default:
                ExitFunc()
            }
        }
    }()
    fmt.Println(BASE_URL)
    http.HandleFunc(BASE_URL+"/nslookup",nsLookup)
	http.HandleFunc(BASE_URL+"/baidu",sendBaidu)
    http.HandleFunc(BASE_URL+"/http", sendHttp)
    http.HandleFunc(BASE_URL+"/header", getRemoteHeaders)
    http.HandleFunc(BASE_URL+"/delay", delay)
    http.HandleFunc(BASE_URL+"/curl-testing", curlTesting)
    http.HandleFunc("/", rootData)
    fmt.Println("Start server successfully!\nNow listen  0.0.0.0:"+PORT)
    fmt.Println("Routes:\n [GET] /nslookup?host=baidu.com\n [GET] /baidu\n [GET] /http?url=http://baidu.com\n [GET] /header\n [GET] /delay?s=10" )
    fmt.Println(" [GET] /curl-testing?url=https://baidu.com[&method=GET] \n" )
    http.ListenAndServe(":"+PORT, nil)
}

 func getTpl(req *http.Request) string {
    content := req.Header.Get("User-Agent")  
    path :="./index.html"
   
    if path=="" || strings.Contains(content,"curl") {
        path = "./index.txt"
    }
    c, _ := ioutil.ReadFile(path)
    return string(c)
 }


// 将当前目录写的主页文件写入访问
func rootData(w http.ResponseWriter, req *http.Request) {
    s := getTpl(req)

    hostname, err := os.Hostname()
    if err != nil {
        panic(err)
    }
    s = strings.Replace(s,"{{hostname}}",hostname,1)
    s = strings.Replace(s,"{{version}}",Version,1)
    w.Write([]byte(s))
    w.Write([]byte("\n"))
}

func sendBaidu(w http.ResponseWriter, req *http.Request) {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
    defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	w.Write(body)
}


func nsLookup(w http.ResponseWriter, req *http.Request) {
    query := req.URL.Query()
    host :=query.Get("host")
    ns,err := net.LookupHost(host)
	if err !=nil {
		panic(err)
	}
    sb := strings.Builder{}
    sb.WriteString("Name: ")
    sb.WriteString(host)
    sb.WriteString("\n")
	for _, n:=range ns{
		fmt.Println(n)
        sb.WriteString("Address: ")
        sb.WriteString(n)
        sb.WriteString("\n")
	}
    w.Write([]byte(sb.String()))
}

func sendHttp(w http.ResponseWriter, req *http.Request){
    query := req.URL.Query()
    url :=query.Get("url")
    response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
    defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	w.Write(body)
}

func getRemoteHeaders(w http.ResponseWriter, req *http.Request){
    sb := strings.Builder{}
    sb.WriteString("Headers:\n")
    if len(req.Header) > 0 {
        for k,v := range req.Header {
            sb.WriteString("    ")
            sb.WriteString(k)
            sb.WriteString(": ")
            sb.WriteString(v[0])
            sb.WriteString("\n")
            
        }
    }
    w.Write([]byte(sb.String()))
}

func delay(w http.ResponseWriter,req *http.Request){
    query := req.URL.Query()
    second := query.Get("s")
    duration,_ := time.ParseDuration(second+"s")
    time.Sleep(duration)
    w.Write([]byte(second+"s后\nok\n")) 
}

func curlTesting(w http.ResponseWriter,req *http.Request){
    w.Header().Set("Content-Type","text/plain; charset=utf-8")
    
    query := req.URL.Query()
    url := query.Get("url")
    method := query.Get("method") 
    if method == "" {
        method = "GET"
    }
    // alpine 镜像使用 sh
    out, err := exec.Command("sh","-c","curl -s -o /dev/null -w \"@curl-testing-formatter.txt\" -X "+method + " "+ url ,).Output()
	if err != nil {
	    w.Write([]byte(err.Error()+"\n"))
	}
    w.Write(out)
    w.Write([]byte("\n"))

}

// 捕获到退出信号后，执行的退出流程
func ExitFunc()  {
    fmt.Println("\nThe web server is shutting down")
    os.Exit(0)
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