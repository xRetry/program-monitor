package main

import (
    "log"
    "os"
    "os/exec"
    "strings"
)

type Program struct {
    name string
    cmd_start string
    cmd_status string
}


func (p Program) CheckStatus() bool {
    log.Printf("Checking status: %s\n", p.name)
    args := strings.Split(p.cmd_status, " ")
    if len(args) < 1 {
        log.Printf("Invalid check command for %s: %s\n: ", p.name, p.cmd_status)
    }

    cmd := exec.Command(args[0], args[1:]...)
    _, err := cmd.Output()
    if err != nil {
        log.Println(err)    
        // TODO: Attempt restart
        return false
    }
    log.Printf("Status OK: %s\n", p.name)
    return true
}


func main() {
    // If the file doesn't exist, create it or append to the file
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)
}
