package main

import (
    "encoding/json"
    "os"
)

/**

字段的tag是"-"，那么这个字段不会输出到JSON
tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串

*/

type Server2 struct {
    // ID 不会导出到JSON中
    ID int `json:"-"`

    // ServerName2 的值会进行二次JSON编码
    ServerName  string `json:"serverName"`
    ServerName2 string `json:"serverName2,string"`

    // 如果 ServerIP 为空，则不输出到JSON串中
    ServerIP   string `json:"serverIP,omitempty"`
}

/**
JSON对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型(T是Go语言中任意的类型)
Channel, complex和function是不能被编码成JSON的
嵌套的数据是不能编码的，不然会让JSON编码进入死循环
指针在编码的时候会输出指针指向的内容，而空指针会输出null
 */

func main () {
    s := Server2 {
        ID:         3,
        ServerName:  `Go "1.0" `,
        ServerName2: `Go "1.0" `,
        ServerIP:   ``,
    }
    b, _ := json.Marshal(s)
    os.Stdout.Write(b)
}