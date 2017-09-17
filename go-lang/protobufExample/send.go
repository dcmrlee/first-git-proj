package main

import (
    "github.com/golang/protobuf/proto"
    "net/http"
    "fmt"
    "bytes"

    pb "github.com/dcmrlee/first-git-proj/go-lang/protobufExample/myproto"
)

func main() {
    myClient := pb.Client{Id: 1234, Name: "Dcmrlee", Email: "lidachao@jd.com", Country: "China"}
    clientInbox := make([]*pb.Client_Mail, 0, 20)
    clientInbox = append(clientInbox,
        &pb.Client_Mail{RemoteEmail: "dcmrlee@qq.com", Body: "Hello QQ Email!"},
        &pb.Client_Mail{RemoteEmail: "dcmrlee@gmail.com", Body: "Hello Gmail!"})
    myClient.Inbox = clientInbox

    data, err := proto.Marshal(&myClient)
    if err != nil {
        fmt.Println(err)
        return
    }

    _, err = http.Post("http://localhost:3000", "", bytes.NewBuffer(data))
    if err != nil {
        fmt.Println(err)
        return
    }
}
