package main

import (
        "fmt"
        "net/http"
        "log"
        "io/ioutil"
        "strings"
        "time"
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
                                fmt.Println("Recovered from panic:", r)
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
    fmt.Println("main exec")
    defer func() {
        if r := recover(); r != nil {
                fmt.Println("Function error:", r)
                }
        }()

    checkIP("https://192.142.12.12")
    checkIP("https://192.142.12.20")

    fmt.Println("all jobs done")

}
