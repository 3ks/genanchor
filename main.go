// 去除 , : + " ? & 等字符并加上空格
// 去除 . / 等字符并加上空格
// 使用**** 标识含有 `` 的标题
// 去除‘ 并恢复完整写法
//
package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	addAnchor()
}

func addAnchor() {
	ab, sp := Load()
	f, err := os.Open("content/index.md")
	if err != nil {
		panic(err)
	}
	new := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		vLine := sc.Text()
		// 标题，且不包含反引号
		if strings.HasPrefix(vLine, "##") && !strings.ContainsRune(vLine, '`') {
			// 去除 # , : + " ? & . /等字符
			values := strings.FieldsFunc(vLine, func(c rune) bool {
				return strings.ContainsRune(`# ,:+"?&`, c)
			})
			for k, v2 := range values {

				if k == 0 {
					vLine += "{#"
				}

				// 非特有名词，转小写
				if !sp[v2] {
					v2 = strings.ToLower(v2)
				}
				// 缩写
				if ab[v2] != "" {
					v2 = ab[v2]
				}

				// 写入当前行
				if k == len(values)-1 {
					vLine = vLine + v2 + "}"
				} else {
					vLine = vLine + v2 + "-"
				}
			}
		}
		new.WriteString(vLine + "\n")
	}
	err = ioutil.WriteFile("content/new.md", new.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
