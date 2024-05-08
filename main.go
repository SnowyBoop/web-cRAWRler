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

func main() {
    resp, err := http.Get("https://nudle.ltd/contact")

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }
        defer un(trace("nigg"))
    stringBody := string(body)

        fmt.Println(stringBody)

    if (strings.Contains(stringBody , "discord.gg/")) {
        fmt.Println("contains html")
        }




}

const serverPort = 3333
