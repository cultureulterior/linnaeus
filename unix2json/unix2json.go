package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

var out = flag.String("out", "-", "Outfile")

func main() {
	flag.Parse()
	data := make(map[string]int64)
	m_junk, err := ioutil.ReadFile("/proc/meminfo")
	c_junk, err := ioutil.ReadFile("/proc/cpuinfo")
	m_regex := regexp.MustCompile("MemTotal:[[:space:]]+([0-9]+)")
	c_regex := regexp.MustCompile("^processor")
	data["memory"], err = strconv.ParseInt(m_regex.FindStringSubmatch(string(m_junk))[1], 10, 64)
	data["cpus"] = int64(len(c_regex.FindAllIndex(c_junk, -1)))
	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	if out != nil && *out != "-" {
		ioutil.WriteFile(*out, jsonString, 0666)
	} else {
		fmt.Print(string(jsonString))
	}
}
