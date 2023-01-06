package main

import (
    "log"
    "os"
    "os/exec"
    "strings"
    "errors"
)

type Program struct {
    Name string
    CmdStart string
    CmdStatus string
    Status *ExecError
}

     
type ExecError struct {
    Err error
    Type int
}

const (
    PARSING_ERROR = iota
    RETURN_ERROR = iota
)


func (p *Program) Check(startIfDown bool) {
    log.Printf("[%s] Checking status\n",p.Name)
    err := execConsoleCommand(p.CmdStatus)
    if err == nil {
        log.Printf("[%s] Status OK\n", p.Name)
    } else if startIfDown && err.Type == RETURN_ERROR {
        p.Start()
        return
    } else {
        log.Printf("[%s] %s\n", p.Name, err.Err)
    }

    p.Status = err
}


func (p *Program) Start() {
    log.Printf("[%s] Starting program\n", p.Name)
    err := execConsoleCommand(p.CmdStart)
    if err != nil {
        log.Printf("[%s] %s\n", p.Name, err.Err)
    } else {
        log.Printf("[%s] Program started\n", p.Name)
    }

    p.Status = err
}


func execConsoleCommand(command string) *ExecError {
    args := strings.Split(command, " ")
    if len(args) < 1 || args[0] == "" {
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


type ProgramControl struct {
    programs []Program
}

func (c *ProgramControl) AddProgram(newProgram Program) {
    c.programs = append(c.programs, newProgram)
}

func (c *ProgramControl) CheckAll() {
    for i, _ := range c.programs {
        c.programs[i].Check(false)
    }
}

func (c *ProgramControl) StartAll() {
    for i, _ := range c.programs {
        c.programs[i].Check(true)
    }
}



func main() {
    // If the file doesn't exist, create it or append to the file
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)
}
