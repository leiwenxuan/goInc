package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	//f, err3 := os.Create("./output3.txt") //创建文件
	//defer f.Close()
	//fmt.Println(err3)
	//n2, err3 := f.Write([]byte("雷文轩")) //写入文件(字节数组)
	//fmt.Printf("写入 %d 个字节n", n2)
	//n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
	//fmt.Printf("写入 %d 个字节n", n3)
	//f.Sync()
	start := time.Now()

	b, err := ioutil.ReadFile("./output3.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	cost := time.Since(start)
	fmt.Printf("执行inception 运行时间 =[%s]", cost)
	fmt.Println(str)

}
