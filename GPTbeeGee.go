package main

import (
    "bytes"
    "crypto/tls"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "strings"
    "time"
)

// getOptionsFromIniFile reads the INI file and returns the contents as a map of strings.
// It returns an error if there is any problem reading or parsing the INI file
func getOptionsFromIniFile(f string) (map[string]string, error) {
    var conf map[string]string

    // Read iniFile and return error if any
    jsonContent, err := ioutil.ReadFile(f)
    if err != nil {
        return nil, fmt.Errorf("Failed to read ini file: %v", err) //line 17
    }
    if err := json.Unmarshal(jsonContent, &conf); err != nil {
        return nil, fmt.Errorf("Failed to parse JSON: %v", err) //line 21
    }

    // Create Batch Location string
    hostName, _ := os.Hostname()
    scriptPath, _ := os.Executable()
    conf["BATCH_LOCATION"] = hostName + " - " + scriptPath

    // Get status from ARG2
    conf["EFFECTIVE_STATUS"] = os.Args[2]

    // Test if third argument is provided, if not, replace by current date/time
    if len(os.Args) == 4 {
        conf["VALUE"] = os.Args[3]
    } else {
        conf["VALUE"] = time.Now().Format("Monday 2006-01-02 15:04:05") //line 33
    }

    return conf, nil
}

// postToBee makes an HTTP POST request to the specified URL with the JSON data
// It returns an error if there is any problem with the request or response
func postToBee(c map[string]string) error {
    // Disable certiciate as we use port 443 without encryption
    tr := &http.Transport{
        // This is the insecure setting, it should be set to false.
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}

    url := c["BEE_URL"]
    // Remove BEE_URL option from Bee Array
    delete(c, "BEE_URL")

    // Create BEE-DATA message
    beeData := map[string]map[string]string{}
    beeData["bee-data"] = c
    msg, err := json.Marshal(beeData)

    if err != nil {
        return fmt.Errorf("Failed to convert to JSON: %v", err) //line 50
    }

    // Send beeData to url
    resp, err := client.Post(url, "application/json", bytes.NewBuffer(msg))
    if err != nil {
        return fmt.Errorf("Failed to make HTTP POST request: %v", err) //line 56
    }
    defer resp.Body.Close()