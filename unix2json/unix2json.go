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
	if m_junk, err := ioutil.ReadFile("/proc/meminfo"); err==nil {
		m_regex := regexp.MustCompile("MemTotal:[[:space:]]+([0-9]+)")		
		data["memory"], err = strconv.ParseInt(m_regex.FindStringSubmatch(string(m_junk))[1], 10, 64)
	} else { log.Fatal(err) }
	if c_junk, err := ioutil.ReadFile("/proc/cpuinfo"); err==nil {
		c_regex := regexp.MustCompile("^processor")	
		data["cpus"] = int64(len(c_regex.FindAllIndex(c_junk, -1)))
	} else { log.Fatal(err) }
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
