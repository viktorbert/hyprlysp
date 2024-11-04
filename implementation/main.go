package main

import (
    //std imports
    "fmt"

    //repo imports
    "github.com/viktorbert/hyprlysp/implementation/types"
    "github.com/viktorbert/hyprlysp/implementation/reader"


    //other imports
    "github.com/chzyer/readline"

)

func main() {

    

    rl, err := readline.New("> ")
    if err != nil {
        panic(err)
    }
    defer rl.Close()

    fmt.Fprintln(rl,"hyprlysp v0.0.1")

    fmt.Fprintln(rl,types.HLSysInt{2})

    for {
        line, err := rl.Readline()
        if err != nil { // io.EOF
            break
        }
        tokens := reader.Tokenize(line)
        fmt.Fprintln(rl,tokens)
        parsed := reader.Parse(tokens)
        fmt.Fprintln(rl,parsed)
    }
}