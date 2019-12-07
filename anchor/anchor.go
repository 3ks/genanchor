package anchor

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
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
func handle(pathname string,wg *sync.WaitGroup) {
	defer wg.Done()
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
				return strings.ContainsRune(`* #,:+"?&/()`, c)
			})

			anchor := make([]string, len(values))
			count := 0
			for _, word := range values {
				v := handleWord(word)
				if v != "" {
					anchor[count] = v
					count++
				}
			}
			vLine = joinLine(vLine, anchor[:count])
		}
		new.WriteString(vLine+"\n")

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
	line += "{#"
	isFirst :=true
	isVar := false
	for _, v := range words {
		if v == "{{<" {
			isVar = true
			line+="-"+v+" "
			continue
		}
		if v == ">}}" {
			isFirst = false
			isVar = false
			line+=" "+v
			continue
		}
		if isVar{
			line+=v
			continue
		}
		v=strings.ReplaceAll(v,"_","")
		if isFirst{
			line+=v
			isFirst=false
		}else{
			line+="-"+v
		}
	}
	return line+"}"
}

func handleWord(word string) string {
	if spell[word] {
		// 特有名词，不处理，来源于 Istio
	} else{
		if _,ok:=special[word] ;ok {
			// 特殊名词，自行维护
			word = special[word]
		} else {
			// 非特有名词、非特殊名词，转小写
			word = strings.ToLower(word)
		}
	}


	// 缩写
	if special[word] != "" {
		return special[word]
	}
	return word
}
