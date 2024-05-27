package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
)

func main() {

        fmt.Print("Enter the file path containing the IP addresses: ")
        reader := bufio.NewReader(os.Stdin)
        filePath, _ := reader.ReadString('\n')
        filePath = filePath[:len(filePath)-1]

        file, err := os.Open(filePath)
        if err != nil {
                fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filePath, err)
                return
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                ips := strings.Split(line, ",")
                for _, ip := range ips {
                        ip = strings.TrimSpace(ip)
                  
                        names, err := net.LookupAddr(ip)
                        if err != nil {
                                fmt.Fprintf(os.Stderr, "Error looking up rDNS for %s: %v\n", ip, err)
                                continue
                        }

                        for _, name := range names {
                                addrs, err := net.LookupHost(name)
                                if err != nil {
                                        fmt.Fprintf(os.Stderr, "Error looking up IP for domain %s: %v\n", name, err)
                                        continue
                                }

                                fmt.Printf("%s -> %v\n", name, addrs)
                        }
                }
        }

        if err := scanner.Err(); err != nil {
                fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filePath, err)
                return
        }
}



