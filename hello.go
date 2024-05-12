package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"time"
	"os"
)

func trace(s string) (string, time.Time) {
    log.Println("START:", s)
    return s, time.Now()
}

func un(s string, startTime time.Time) {
    endTime := time.Now()
    log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}

func errFunc() {
	panic("FUCK")
}

func generateFS() {

    f, err := os.Create("scan.ip")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()


}

func checkIP(caller string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from function error:", r)
		}
	}()

        resp, err := http.Get(caller)

	if err != nil {
		fmt.Println("Error:", err)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("IP was unreachable:", r)
			    	f, err := os.OpenFile("scan.ip", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    				if err != nil {
        				log.Fatal(err)
   				}
    				if _, err := f.Write([]byte("something is unreachable\n")); err != nil {
        				log.Fatal(err)
    				}
    				if err := f.Close(); err != nil {
        				log.Fatal(err)
    				}
			}
		}()
		panic(err)
	}

        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)

        stringBody := string(body)

	fmt.Println(stringBody)

	if (strings.Contains(stringBody , "test")) {}
	

	errFunc()
}

func main() {
    fmt.Println("start jobs")

    generateFS()

    checkIP("https://192.142.12.12")
    checkIP("https://192.142.12.20")

    fmt.Println("all jobs done")

}

