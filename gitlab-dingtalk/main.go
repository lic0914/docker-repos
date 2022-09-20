package main

import (
    "net/http"
    "bytes"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "log"
    "fmt"
    "strconv"
    "os"
    "strings"
    "github.com/joho/godotenv"
    "gitlabdingtalk/utils"
    "io/ioutil"
)

type ObjectAttributes struct{
    Id  int       `json:"id"`
    Status string `json:"status"`
    FinishedAt string `json:"finished_at" `
    Duration int `json:"duration"`
    Ref string `json:"ref"`
}
type User struct{
    Name string `json:"name"`
    AvatarUrl string `json:"avatar_url"`
}
type Commit struct{
    Id string `json:"id"`
    Message string `json:"message"`
    Url string `json:"url"`
}
type Project struct{
    Id  int       `json:"id"`
    Name string `json:"name"`
    WebUrl string `json:"web_url"`
}
type GitlabRequest struct {
    Kind string `json:"object_kind"`
    ObjectAttributes   *ObjectAttributes  `json:"object_attributes"` 
    User   *User    `json:"user"` 
    Project *Project `json:"project"` 
    Commit *Commit `json:"commit"`
}


type SendDingtalkRequest struct{
    MsgType string `json:"msgtype"`
    Markdown *DingtalkMd `json:"markdown"`
    At *AtMobiles `json:"at"` 
}
type DingtalkMd struct{
    Title string `json:"title"`
    Text string `json:"text"`
}
type AtMobiles struct{
    AtMobiles []string `json:"atMobiles"`
}

var err = godotenv.Load(".env")

var (
    Version = os.Getenv("VERSION")
    PORT = os.Getenv("PORT")
    SECRET = os.Getenv("SECRET")
    DingtalkUrl = os.Getenv("DingtalkUrl")
)

func SendDingtalk(req SendDingtalkRequest) []byte{

    url := DingtalkUrl
    if SECRET != "" {
        ts := utils.GetTs()
        sign := utils.Sign(ts,SECRET)
        url = url + "&sign="+sign+"&timestamp="+ts
    }
    data,err :=json.Marshal(req) // var buf bytes.Buffer \n json.NewEncoder(&buf).Encode(req)
    if err != nil {
        log.Fatal(err)
    }
    resp, err := http.Post(url,"application/json", bytes.NewReader(data))
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    bodyC, _ := ioutil.ReadAll(resp.Body)
    return bodyC
}

func SendPipeline(m GitlabRequest) []byte{
    
    sb := strings.Builder{}

    sb.WriteString("\n ["+m.Project.Name)
    sb.WriteString("]("+m.Project.WebUrl+") ")
    sb.WriteString(" **"+m.ObjectAttributes.Status+"**")
    sb.WriteString("\n ")

    sb.WriteString(" "+m.Kind)
    sb.WriteString(" [#"+strconv.Itoa(m.ObjectAttributes.Id)+"](")
    sb.WriteString(m.Project.WebUrl+"/-/pipelines/"+strconv.Itoa(m.ObjectAttributes.Id)+")")
   
    sb.WriteString(" to "+m.ObjectAttributes.Ref)
    sb.WriteString("\n\n ")
    sb.WriteString("> ["+string([]byte(m.Commit.Id)[:6]))
    sb.WriteString("]("+m.Commit.Url+") "+m.Commit.Message)
   
    sb.WriteString("\n\n ")
    sb.WriteString(m.User.Name)
    fmt.Println(sb.String())

    text :=  new(DingtalkMd)
    text.Title = m.Kind
    text.Text = sb.String()

    //mobile := os.Getenv("AT_"+m.User.Name)
    req := SendDingtalkRequest{
        MsgType: "markdown",
    }
    req.Markdown=text
    //at :=  new(AtMobiles)
    //at.AtMobiles = []string{""}
    //req.At=at
    data,_ :=json.Marshal(req) 
    fmt.Println(string(data))
    return SendDingtalk(req)
    
}

func Handler(c *gin.Context){

    var kind struct {
        Kind string `json:"object_kind"`
    }
    if err := c.ShouldBindBodyWith(&kind,binding.JSON); err != nil {
        panic(err)
    }

    if kind.Kind == "pipeline" {
        var model  GitlabRequest
        if err := c.ShouldBindBodyWith(&model,binding.JSON); err != nil {
            panic(err)
        }
       
        jsonBytes:=SendPipeline(model)
        c.Data(http.StatusOK,"application/json",jsonBytes)
    }
    

}

func main() {
    
    if err != nil {
        log.Fatalf("Error loading .env file")
        fmt.Println("error")
    }

    r := gin.Default()
    
    r.POST("/webhook",Handler)
    r.Run(":"+PORT)
}



