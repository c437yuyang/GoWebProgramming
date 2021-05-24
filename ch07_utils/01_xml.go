package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)

type Recurlyservers struct {
    XMLName     xml.Name `xml:"servers"`      // struct tag, 用来辅助反射
    Version     string   `xml:"version,attr"` // Go语言的反射机制，可以利用这些tag信息来将来自XML文件中的数据反射成对应的struct对象，关于反射如何利用struct tag的更多内容请参阅reflect中的相关内容。
    Svs         []server `xml:"server"`
    Description string   `xml:",innerxml"`
}

type server struct {
    XMLName    xml.Name `xml:"server"`
    ServerName string   `xml:"serverName"`
    ServerIP   string   `xml:"serverIP"`
}

type Servers struct {
    XMLName xml.Name `xml:"servers"`
    Version string   `xml:"version,attr"`
    Svs     []server `xml:"server"`
}

func parseXml() {
    file, err := os.Open("I:\\Documents\\Desktop\\codes\\golang\\src\\GoWebProgramming\\ch07_utils\\servers.xml") // For read access.
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    defer file.Close()
    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    v := Recurlyservers{}
    err = xml.Unmarshal(data, &v)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }

    fmt.Println(v)
}

func buildXml() {
    v := &Servers{Version: "1"}
    v.Svs = append(v.Svs, server{xml.Name{Local: "local1", Space: "space1"}, "Shanghai_VPN", "127.0.0.1"})
    v.Svs = append(v.Svs, server{xml.Name{Local: "local2", Space: "space2"}, "Beijing_VPN", "127.0.0.2"})
    output, err := xml.MarshalIndent(v, "  ", "    ")
    if err != nil {
        fmt.Printf("error: %v\n", err)
    }
    os.Stdout.Write([]byte(xml.Header))

    os.Stdout.Write(output)
}

func main() {
    //parseXml()
    buildXml()

}
