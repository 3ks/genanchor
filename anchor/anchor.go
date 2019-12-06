package anchor

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	abbreviation map[string]string
	special      map[string]string
	spell        map[string]bool
	files        []string
	count        int
)

// load:en+path
// output:zh+path
func handle(pathname string) {
	f, err := os.Open(pathname)
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
				return strings.ContainsRune(` #,:+"?&/()`, c)
			})

			anchor := make([]string, len(values))
			count := 0
			for k, word := range values {
				v := handleWord(word)
				if v != "" {
					anchor[count] = v
					count++
				}
			}
			vLine = joinLine(vLine, anchor[:count])
		}
		//vLine = strings.ReplaceAll(strings.ReplaceAll(vLine, "{{<-", "{{< "), "->}}", " >}}")
		new.WriteString(vLine)

	}

	newPath := strings.ReplaceAll(pathname, "/en/", "/zh/")
	err = ioutil.WriteFile(newPath, new.Bytes(), 0644)
	if err != nil {
		_ = os.MkdirAll(path.Dir(newPath), os.ModeDir)
	}
	err = ioutil.WriteFile(newPath, new.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func joinLine(line string, words []string) string {
	anchor := "{#"
	isVar := false
	for k, v := range words {
		if v == "{{<" {
			isVar = true
		}
		if v == ">}}" {
			isVar = false
		}

	}
}

func handleWord(word string) string {
	if spell[word] {
		// 特有名词，不处理，来源于 Istio
	} else if special[word] != "" {
		// 特殊名词，自行维护
		word = special[word]
	} else {
		// 非特有名词、非特俗名词，转小写
		word = strings.ToLower(word)
	}

	// 缩写
	if special[word] != "" {
		word = special[word]
	}
	return word
}
