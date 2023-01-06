package main

import (
    "testing"
)


func createTestProgram(cmd string) *Program {
    program := Program{
        Name: "Test Program",
        CmdStart: cmd,
        CmdStatus: cmd,
        Status: nil,
    }
    return &program
}


func TestStartProgram(t *testing.T) {
    program := createTestProgram("ls")
    program.Check(true)
    if program.Status != nil {
        t.Error("One word command resulted in an error!")
    }

    program = createTestProgram("echo test")
    program.Start()
    if program.Status != nil {
        t.Error("Two word command resulted an error!")
    }

    program = createTestProgram("dfasdfasdf")
    program.Start()
    if program.Status.Type != RETURN_ERROR {
        t.Error("Wrong command did not result in an return error!")
    }

    program = createTestProgram("")
    program.Start()
    if program.Status.Type != PARSING_ERROR {
        t.Error("Wrong command did not result in an parsing error!")
    }
}

func TestCheckStatus(t *testing.T) {
    program := createTestProgram("ls")
    program.Check(true)
    if program.Status != nil {
        t.Error("One word command resulted in an error!")
    }

    program = createTestProgram("echo test")
    program.Check(true)
    if program.Status != nil {
        t.Error("Two word command resulted an error!")
    }

    program = createTestProgram("dfasdfasdf")
    program.Check(true)
    if program.Status.Type != RETURN_ERROR {
        t.Error("Wrong command did not result in an return error!")
    }

    program = createTestProgram("")
    program.Check(true)
    if program.Status.Type != PARSING_ERROR {
        t.Error("Wrong command did not result in an parsing error!")
    }
}

//func TestCheckAll(t *testing.T) {
//    store :- ProgramStore{}
//}
