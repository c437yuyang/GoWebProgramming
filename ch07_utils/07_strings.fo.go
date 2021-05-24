package main

import (
    "fmt"
    "strings"
)

func main() {
    {
        fmt.Println(strings.Contains("seafood", "foo"))
        fmt.Println(strings.Contains("seafood", "bar"))
        fmt.Println(strings.Contains("seafood", ""))
        fmt.Println(strings.Contains("", ""))
        //Output:
        //true
        //false
        //true
        //true
    }
    {
        s := []string{"foo", "bar", "baz"}
        fmt.Println(strings.Join(s, ", "))
        //Output:foo, bar, baz
    }
    {
        fmt.Println(strings.Index("chicken", "ken"))
        fmt.Println(strings.Index("chicken", "dmr"))
        //Output:4
        //-1
    }
    {
        fmt.Println("ba" + strings.Repeat("na", 2))
        //Output:banana
    }
    {
        fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) // n为替换次数，-1全部替换
        fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
        //Output:oinky oinky oink
        //moo moo moo
    }
    {
        // split 返回slice
        fmt.Printf("%q\n", strings.Split("a,b,c", ","))
        fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
        fmt.Printf("%q\n", strings.Split(" xyz ", ""))
        fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
        //Output:["a" "b" "c"]
        //["" "man " "plan " "canal panama"]
        //[" " "x" "y" "z" " "]
        //[""]
    }
    {
        fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", " !")) // trim是按照给定的cutset里面每个字符来的，那问题来了，如果想要字符串trim怎么玩
        //Output:["Achtung"]
    }
    {
        // 去除s字符串的空格符，并且按照空格分割返回slice
        fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
        //Output:Fields are: ["foo" "bar" "baz"]
    }
}
