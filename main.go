package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    buf := make([]byte, 0, 4*1024)
    scanner.Buffer(buf, 4*1024)

    for scanner.Scan() {
        line := scanner.Text()

        if strings.HasPrefix(line, "mr status:") {
            line = line[len("mr status:")+1:]
            printRepo(line)
        } else if strings.HasPrefix(line, "mr update:") {
            line = line[len("mr update:")+1:]
            printRepo(line)
        } else if line != "" && line != "Already up to date." {
            fmt.Println("  " + line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func printRepo(line string) {
    split := strings.Split(line, "/")
    alias := split[len(split)-1]

    if strings.HasPrefix(alias, "finished (") {
        fmt.Println("\033[0;34m" + alias + "\033[0m")
    } else {
        fmt.Println("\033[0;35m" + alias + "\033[0m" + " (" + line + ")")
    }
}