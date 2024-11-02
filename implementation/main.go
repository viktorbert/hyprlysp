package main

import (
    //std imports
    "fmt"
    //repo imports
    "github.com/viktorbert/hyprlysp/implementation/types"
    "github.com/viktorbert/hyprlysp/implementation/reader"


    //other imports
    "github.com/chzyer/readline"

    // temp
    //s "strings"
)

/*
func Tokenize(i string) []string {
	i = s.Replace(i,"(", " ( ",-1)
	i = s.Replace(i,")", " ) ",-1)
	i = s.Replace(i,"[", " [ ",-1)
	i = s.Replace(i,"]", " ] ",-1)
	i = s.Replace(i,"{", " { ",-1)
	i = s.Replace(i,"}", " } ",-1)

	return s.Split(i," ")
}
*/

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
        fmt.Fprintln(rl,(reader.Tokenize(line)))
    }
}