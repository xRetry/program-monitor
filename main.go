package main

import (
    "log"
    "os"
    "os/exec"
    "strings"
    "errors"
)

type Program struct {
    name string
    cmdStart string
    cmdStatus string
}

     
type ExecError struct {
    Err error
    Type int
}

const (
    PARSING_ERROR = iota
    RETURN_ERROR = iota
)


func (p Program) CheckStatus() bool {
    log.Printf("[%s] Checking status\n", p.name)
    err := execConsoleCommand(p.cmdStatus)
    if err == nil {
        log.Printf("[%s] Status OK\n", p.name)
        return true
    }

    log.Printf("[%s] %s\n", p.name, err.Err)

    if err.Type == RETURN_ERROR {
        return p.Start() 
    }

    return false
}

func (p Program) Start() bool {
    log.Printf("[%s] Starting program\n", p.name)
    err := execConsoleCommand(p.cmdStart)
    if err != nil {
        log.Printf("[%s] %s\n", p.name, err.Err)
        return false
    }
    log.Printf("[%s] Program started\n", p.name)
    return true
}

func execConsoleCommand(command string) *ExecError {
    args := strings.Split(command, " ")
    if len(args) < 1 {
        return &ExecError{
            errors.New("The command has to be at least one word!"),
            PARSING_ERROR,
        }
    }

    cmd := exec.Command(args[0], args[1:]...)
    _, err := cmd.Output()
    if err != nil {
        return &ExecError{
            err,
            RETURN_ERROR,
        }
    }

    return nil
}

func main() {
    // If the file doesn't exist, create it or append to the file
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)
}
