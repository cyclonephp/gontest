package main

import(
    "fmt"
    "os"
    "gontest/input"
    "net/http"
    "io/ioutil"
    "gontest/runner"
)

var exitCodes = map[string]int{
   "io": 1,
   "xml": 2, 
}

func main(){
	var in, ioErr = ioutil.ReadAll(os.Stdin)
	if ioErr != nil {
	   fmt.Println(ioErr)
	   os.Exit(exitCodes["io"])
	   return
	}
	var req, xmlErr = input.ReadRawInput(in);
	if xmlErr != nil {
	   fmt.Println("failed to parse input XML") 
	   fmt.Println(xmlErr)
	   os.Exit(exitCodes["xml"])
	   return
	}
	fmt.Println(req)
	var client = http.Client{}
	runner.Run(req, client)
	/*fmt.Println("-----------")
	resp, err := http.Get("http://cyclonephp.org")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))*/
}
