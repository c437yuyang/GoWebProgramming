package main
import (
    "encoding/json"
    "fmt"
)

type Server1 struct {
    ServerName string `json:"serverName"`  // 默认输出字段和fieldName一致，通过tag修改
    ServerIP   string `json:"serverIP"`
}

type Serverslice1 struct {
    Servers []Server1
}


func main() {
    var s Serverslice1
    s.Servers = append(s.Servers, Server1{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
    s.Servers = append(s.Servers, Server1{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
    b, err := json.Marshal(s)
    if err != nil {
        fmt.Println("json err:", err)
    }
    fmt.Println(string(b)) // {"Servers":[{"ServerName":"Shanghai_VPN","ServerIP":"127.0.0.1"},{"ServerName":"Beijing_VPN","ServerIP":"127.0.0.2"}]}
}