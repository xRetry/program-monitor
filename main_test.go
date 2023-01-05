package main

import (
    "testing"
)


func createTestProgram(cmd string) *Program {
    program := Program{
        name: "Test Program",
        cmdStart: cmd,
        cmdStatus: cmd,
    }
    return &program
}

func TestCheckStatus(t *testing.T) {
    program := createTestProgram("ls")
    ok := program.CheckStatus()
    if !ok {
        t.Error("One word command resulted in an error!")
    }

    program = createTestProgram("echo test")
    ok = program.CheckStatus()
    if !ok {
        t.Error("Two word command resulted an error!")
    }

    program = createTestProgram("dfasdfasdf")
    ok = program.CheckStatus()
    if ok {
        t.Error("Wrong command did not result in an error!")
    }
}

func TestStartProgram(t *testing.T) {
    program := createTestProgram("ls")
    ok := program.Start()
    if !ok {
        t.Error("One word command resulted in an error!")
    }

    program = createTestProgram("echo test")
    ok = program.Start()
    if !ok {
        t.Error("Two word command resulted an error!")
    }

    program = createTestProgram("dfasdfasdf")
    ok = program.Start()
    if ok {
        t.Error("Wrong command did not result in an error!")
    }
}
