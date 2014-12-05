package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

var out = flag.String("out", "-", "Outfile")

func main() {
	flag.Parse()
	data := make(map[string]map[string][]string)
	if ints,err := net.Interfaces(); err==nil {
		for _,inte := range ints {
			data[inte.Name]=make(map[string][]string)
			if adds,err:= inte.Addrs(); err==nil {
				for _,add:=range adds {
					switch ifa := add.(type) {
						case *net.IPAddr:
						if add4:=ifa.IP.To4();add4 == nil{
							data[inte.Name]["ipv6"] = append(data[inte.Name]["ipv6"],add.String())
						} else {
							data[inte.Name]["ipv4"] = append(data[inte.Name]["ipv4"],add.String())
						}
						case *net.IPNet:
						if add4:=ifa.IP.To4();add4 == nil{
							data[inte.Name]["ipv6"]=append(data[inte.Name]["ipv6"],add.String())
						} else {
							data[inte.Name]["ipv4"] = append(data[inte.Name]["ipv4"],add.String())
						}
					}
				}
			} else { log.Fatal(err) }
		}
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
