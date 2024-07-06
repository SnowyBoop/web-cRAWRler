package main

// test

import (
        "fmt"
        "net/http"
        "log"
        "io/ioutil"
        "strings"
        "time"
        "os"
                "math/rand"
                "net"
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

        f2, err := os.Create("minecraft.ip")
    if err != nil {
        log.Fatal(err)
    }
    defer f2.Close()

        f3, err := os.Create("website.html")
    if err != nil {
        log.Fatal(err)
    }
    defer f3.Close()

        f4, err := os.Create("discord.ip")
    if err != nil {
        log.Fatal(err)
    }
    defer f4.Close()


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
                                if _, err := f.Write([]byte(caller + " is unreachable\n")); err != nil {
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


                f, err := os.OpenFile("scan.ip", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                if err != nil {
                        log.Fatal(err)
                }
                if _, err := f.Write([]byte(caller + " was callable\n")); err != nil {
                        log.Fatal(err)
                }
                if err := f.Close(); err != nil {
                         log.Fatal(err)
                }


                f2, err := os.OpenFile("minecraft.ip", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                if err != nil {
                        log.Fatal(err)
                }
                if _, err := f2.Write([]byte(caller + " has port 25565 open\n")); err != nil {
                        log.Fatal(err)
                }
                if err := f2.Close(); err != nil {
                         log.Fatal(err)
                }


        stringBody := string(body)

        fmt.Println(stringBody)

        if (strings.Contains(stringBody , "discord.gg")) {

                            f, err := os.OpenFile("discord.ip", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                if err != nil {
                        log.Fatal(err)
                }
                if _, err := f.Write([]byte(caller + " contains a discord invite \n")); err != nil {
                        log.Fatal(err)
                        f.Close()
                }
                if err := f.Close(); err != nil {
                         log.Fatal(err)
                }

                        fmt.Println("discord found on" + caller)

        }



        if (strings.Contains(stringBody , "html")) {

                        if (!(strings.Contains(stringBody, "404") || strings.Contains(stringBody, "403") || strings.Contains(stringBody, "Forbidden"))) {

                            f, err := os.OpenFile("website.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                if err != nil {
                        log.Fatal(err)
                }
                if _, err := f.Write([]byte("<a href='" + caller + "'>" + caller + " contains a valid website \n" + "</a></p> <br>")); err != nil {
                        log.Fatal(err)
                        f.Close()
                }
                if err := f.Close(); err != nil {
                         log.Fatal(err)
                }

                        fmt.Println("website found on" + caller)
                        }

        }


        errFunc()
}


func main() {
    fmt.Println("start jobs")

    generateFS()

for true {

                time.Sleep(50 * time.Millisecond)

                ip := make(net.IP, net.IPv4len)
                for i := 0; i < net.IPv4len; i++ {
                        ip[i] = byte(rand.Intn(256))
                }

                ips := ip.String()

        fmt.Println("calling the ip" + ips)

        go checkIP("http://" + ips)


        }

    fmt.Println("all jobs done")

}


