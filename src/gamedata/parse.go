package gamedata

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

import (
	. "logger"
)

type field struct {
	name  string
	value string
}

type table struct {
	fields map[int32][]field
}

var tables map[string]table
var reg *regexp.Regexp

func init() {
	tables = make(map[string]table)
	reg = regexp.MustCompile("\"")
	parse(os.Getenv("GOPATH") + "/data")
}

func parse(path string) {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		ERROR("get files", err)
		os.Exit(-1)
	}

	for _, file := range files {
		if file.IsDir() {
			parse(path + file.Name())
		}
		if !strings.HasSuffix(file.Name(), ".csv") {
			continue
		}
		f, err := os.Open(path + file.Name())
		if err != nil {
			ERROR("open file", err)
			os.Exit(-1)
		}

		fileName := strings.TrimRight(file.Name(), ".csv")

		t := table{
			fields: make(map[int32][]field),
		}

		scanner := bufio.NewScanner(f)
		row := 0
		row2head := make(map[int]string)
		for scanner.Scan() {
			tmp := strings.Split(reg.ReplaceAllString(scanner.Text(), ""), ",")
			for k := range tmp {
				if row == 0 && k != 0 { // table head
					if filter(tmp[k]) {
						continue
					}
					row2head[k] = tmp[k]
				}
				if _, ok := row2head[k]; ok && row != 0 && k != 0 {
					id, _ := strconv.Atoi(tmp[0])
					t.fields[int32(id)] = append(t.fields[int32(id)], field{row2head[k], tmp[k]})
				}
			}
			row++
		}
		tables[fileName] = t
		f.Close()
	}
}

func filter(s string) bool {
	if strings.HasPrefix(s, "INT_") {
		return false
	}
	if strings.HasPrefix(s, "STR_") {
		return false
	}
	if strings.HasPrefix(s, "FLOAT_") {
		return false
	}
	return true
}
