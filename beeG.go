package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Options struct {
	AUTHKEY        string
	TEAMID         string
	REVISION       string
	HEARTBEAT      string
	ZOMBEAT        string
	X_ORG_ID       string
	CMDB_SERVICE   string
	ALERTID1       string
	ALERTID2       string
	ALERTTEXT      string
	ALERTENV       string
	BATCH_LOCATION string
	BATCH_COMMENT  string
	BATCH_URL      string
	BATCH_DOC      string
	BEHAVIOR       string
	ALERTMAIL      string
	BEE_URL        string
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

	return conf
}

func main() {

	iniFile := os.Args[1]

	// Get Options from ini file
	config := getOptionsFromIniFile(iniFile)

	// Pretty Print content of "config"
	fmt.Println(prettyPrint(config))
}
