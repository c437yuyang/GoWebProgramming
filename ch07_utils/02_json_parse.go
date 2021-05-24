package main

import (
    "encoding/json"
    "fmt"
)

type Server struct {
    ServerName string
    ServerIP   string
    IsOK       bool
}

type Serverslice struct {
    Servers []Server
}

/**
如果key是foo
首先查找tag含有Foo的可导出的struct字段(首字母大写)
其次查找字段名是Foo的导出字段
最后查找类似FOO或者FoO这样的除了首字母之外其他大小写不敏感的导出字段

能够被赋值的字段必须是可导出字段(即首字母大写）。同时JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略，这样的一个好处是：当你接收到一个很大的JSON数据结构而你却只想获取其中的部分数据的时候，你只需将你想要的数据对应的字段名大写，即可轻松解决这个问题。
*/
func main() {
    {
        var s Serverslice
        str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1","isOK":true},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
        json.Unmarshal([]byte(str), &s)
        fmt.Println(s)
    }

    {
        b := []byte(`{"name":"n1","age":1,"parents":["p1","p2"]}`)
        var f interface{}
        err := json.Unmarshal(b, &f)
        fmt.Println(f, err) // map[age:1 name:n1 parents:[p1 p2]] <nil>

        // 类型断言判断是否map类型
        m := f.(map[string]interface{})
        for k, v := range m {
            switch vv := v.(type) {
            case string:
                fmt.Println(k, "is string", vv)
            case bool:
                fmt.Println(k, "is bool", vv)
            case float64:
                fmt.Println(k, "is number", vv)
            case []interface{}:
                fmt.Println(k, "is an array", vv)
                for idx, ele := range vv {
                    fmt.Printf("arr[%d] is:%s\n", idx, ele)
                }
            default:
                fmt.Println(k, "is of a type I don't know how to handle")
            }
        }
    }

    //{
    //    js, err := NewJson([]byte(`{
    //            "test": {
    //                "array": [1, "2", 3],
    //                "int": 10,
    //                "float": 5.150,
    //                "bignum": 9223372036854775807,
    //                "string": "simplejson",
    //                "bool": true
    //            }
    //        }`))
    //
    //    arr, _ := js.Get("test").Get("array").Array()
    //    i, _ := js.Get("test").Get("int").Int()
    //    ms := js.Get("test").Get("string").MustString()
    //}

}
