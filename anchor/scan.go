package anchor

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

func Start(base string) {
	abbreviation, special, spell = loadAbbreviation(), loadSpecial(), loadSpelling()
	if base == "" {
		base = "content/en"
	}
	files = make([]string, 1000)
	fmt.Println("加载 markdown 文件列表中……")
	err := GetAllFile(base)
	if err != nil {
		panic(err)
	}
	fmt.Println("开始生成锚点……")
	wg:=&sync.WaitGroup{}
	for k := range files {
		if k > count || files[k] == "" {
			continue
		}
		wg.Add(1)
		go handle(files[k],wg)
	}
	wg.Wait()
	fmt.Println("任务完成！")
}

func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			err = os.Mkdir(path.Join(strings.ReplaceAll(pathname, "content/en", "content/zh"), fi.Name()), os.ModeDir)
			if err != nil && !os.IsExist(err) {
				panic(err)
			}
			err = GetAllFile(path.Join(pathname, fi.Name()))
			if err != nil {
				fmt.Println(err)
			}
		} else {
			if path.Ext(fi.Name()) == ".md" {
				files[count] = path.Join(pathname, fi.Name())
				count++
				if count+100 > len(files) {
					files = append(files, make([]string, len(files))...)
				}
			}
		}
	}
	return err
}

func loadAbbreviation() (special map[string]string) {
	fmt.Println("加载缩写单词中……")
	special = make(map[string]string)
	f, err := os.Open(".abbreviation")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		v, _, e := rd.ReadLine()
		if e == io.EOF {
			break
		}
		if len(v) == 0 || v[0] == '#' {
			continue
		}
		values := strings.Split(string(v), "=")
		if len(values) < 2 {
			continue
		}
		special[values[0]] = values[1]
	}
	return
}

func loadSpecial() (special map[string]string) {
	fmt.Println("加载特殊单词中……")
	special = make(map[string]string)
	f, err := os.Open(".special")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		v, _, e := rd.ReadLine()
		if e == io.EOF {
			break
		}
		if len(v) == 0 || v[0] == '#' {
			continue
		}
		values := strings.Split(string(v), "=")
		if len(values) < 2 {
			continue
		}
		special[values[0]] = values[1]
	}
	return
}

func loadSpelling() (spell map[string]bool) {
	fmt.Println("加载特有单词中……")
	spell = make(map[string]bool)
	f, err := os.Open(".spelling")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		v, _, e := rd.ReadLine()
		if e == io.EOF {
			break
		}
		if len(v) == 0 || v[0] == '#' {
			continue
		}
		spell[string(v)] = true
	}
	return
}
