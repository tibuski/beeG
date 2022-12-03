package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Options struct {
	AUTHKEY          string
	EFFECTIVE_STATUS string
	TEAMID           string
	REVISION         string
	HEARTBEAT        string
	ZOMBEAT          string
	VALUE            string
	X_ORG_ID         string
	CMDB_SERVICE     string
	ALERTID1         string
	ALERTID2         string
	ALERTTEXT        string
	ALERTENV         string
	BATCH_LOCATION   string
	BATCH_COMMENT    string
	BATCH_URL        string
	BATCH_DOC        string
	BEHAVIOR         string
	ALERT_MAIL       string
	BEE_URL          string
}

func prettyPrint(i interface{}) string {
	// Function : Pretty print config array
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func getOptionsFromIniFile(f string) Options {
	// Function : read and format ini file
	conf := Options{}

	// Read iniFile and print error if any
	jsonContent, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	// Read JSON data in variable "content" in struct "config"
	err = json.Unmarshal(jsonContent, &conf)
	if err != nil {
		log.Fatalf("JSON file error :\n %v", err)
	}

	// Create Batch Location string
	hostName, _ := os.Hostname()
	scriptPath, _ := os.Executable()
	conf.BATCH_LOCATION = hostName + " - " + scriptPath

	// Get status from ARG2
	conf.EFFECTIVE_STATUS = os.Args[2]

	// Get Value from ARG3
	conf.VALUE = os.Args[3]

	return conf
}

func postToBee(c Options) string {
	url := c.BEE_URL
	msg, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("Convert to JSON error :\n %v", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(msg))
	if err != nil {
		log.Fatalf("HTTP POST error :\n %v", err)
	}
	fmt.Printf("%v\n", resp)
	return ("ok")
}

func main() {

	iniFile := os.Args[1]

	// Get Options from ini file
	config := getOptionsFromIniFile(iniFile)

	// Send Data to BEE URL
	postToBee(config)

	// Pretty Print content of "config"
	fmt.Println(prettyPrint(config))
}
