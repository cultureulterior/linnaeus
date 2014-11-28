package main

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
	"fmt"
	"os"
	"strings"
	"regexp"
	"log"
	"encoding/json"
	"io/ioutil"
	"flag"
)

var out = flag.String("out", "-", "Outfile")

func main() {
	flag.Parse()
	auth, err := aws.GetAuth("","")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]string)
	info := []string{"instance-id","ami-id","hostname","instance-type","kernel-id","public-ipv4","security-groups","reservation-id","placement/availability-zone"}
	re := regexp.MustCompile("[[:^alnum:]]")
	for _,key:= range info {
		if val,err := aws.GetMetaData(key); err==nil {
			data[re.ReplaceAllString(key,"_")] = string(val)
		}
	}
	data["region"]=strings.TrimRight(data["placement_availability_zone"],"abcdefghijklm")
	region := aws.Regions[data["region"]]
	client := ec2.New(auth, region)
	filter := ec2.NewFilter()
	filter.Add("resource-id", data["instance_id"])
	resp, err := client.Tags(filter)

	if err == nil {		
		for _,rtag:= range resp.Tags {
			data["tags_"+re.ReplaceAllString(rtag.Tag.Key,"_")]=rtag.Tag.Value
		}
	        fmt.Fprintf(os.Stderr, "Read %d tags\n", len(resp.Tags))
	} else {
	     fmt.Fprintf(os.Stderr, "Error reading tags '%q'\n", err)
	}
	jsonString, err := json.Marshal(data)
	if out != nil && *out != "-" {
		ioutil.WriteFile(*out, jsonString, 0666)
	} else {
		fmt.Print(string(jsonString))
	}
}
