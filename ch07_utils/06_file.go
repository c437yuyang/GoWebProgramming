package main

import (
    "fmt"
    "os"
)

func mkdirAndRemoveDemo() {
    err := os.Mkdir("yangxiaopeng.yxp", 0777)
    if err != nil {
        fmt.Println(err)
    }
    err = os.MkdirAll("yangxiaopeng.yxp/test1/test2", 0777)
    if err != nil {
        fmt.Println(err)
    }
    err = os.Remove("yangxiaopeng.yxp") // 删除文件或者文件夹都可以
    if err != nil {
        fmt.Println(err)
    }
    os.RemoveAll("yangxiaopeng.yxp")
}

func openAndWriteDemo() {
    userFile := "yiheng_demo.txt"
    fout, err := os.Create(userFile) // 文件可以存在，直接覆盖
    if err != nil {
        fmt.Println(userFile, err)
        return
    }
    defer fout.Close()
    for i := 0; i < 10; i++ {
        fout.WriteString("Just a test string!\r\n")
        fout.Write([]byte("Just a test byte!\r\n"))
    }
}

func openAndReadDemo() {
    userFile := "yiheng_demo.txt"
    fl, err := os.Open(userFile) // 文件必须存在
    if err != nil {
        fmt.Println(userFile, err)
        return
    }
    defer fl.Close()
    buf := make([]byte, 1024)
    for {
        n, _ := fl.Read(buf)
        if 0 == n {
            break
        }
        os.Stdout.Write(buf[:n])
    }
}
/*
	  os.O_CREATE|os.O_APPEND
	  或者 os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	  os.O_RDONLY // 只读
	  os.O_WRONLY // 只写
	  os.O_RDWR // 读写
	  os.O_APPEND // 追加（Append）
	  os.O_CREATE // 如果文件不存在则先创建
	  os.O_TRUNC // 文件打开时裁剪文件
	  os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	  os.O_SYNC // 以同步I/O的方式打开
	第三个参数：权限(rwx:0-7)
	  0：没有任何权限
	  1：执行权限
	  2：写权限
	  3：写权限和执行权限
	  4：读权限
	  5：读权限和执行权限
	  6：读权限和写权限
	  7：读权限，写权限，执行权限
*/
func openFileDemo() {
    userFile := "yiheng_demo1.txt"
    fl, err := os.OpenFile(userFile, os.O_CREATE|os.O_APPEND, 0x110) // 文件必须存在
    if err != nil {
        fmt.Println(userFile, err)
        return
    }
    defer fl.Close()
    fl.WriteString("abcde")
}

func main() {
    mkdirAndRemoveDemo()
    openAndWriteDemo()
    openAndReadDemo()
    openFileDemo()
}
