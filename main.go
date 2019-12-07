// 去除 , : + " ? & . / 等字符并加上空格
// 去除' 并恢复完整写法
// 处理特殊单词
package main

import (
	"fmt"
	"genanchor/anchor"
	"time"
)

func main() {
	t1:=time.Now()
	anchor.Start("")
	fmt.Printf("共计耗时：%v ms\n",time.Now().Sub(t1).Milliseconds())
}