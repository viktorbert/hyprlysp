package reader

import (
	s "strings"
	t "github.com/viktorbert/hyprlysp/implementation/types"
	r "regexp"
	"strconv"
)



func Tokenize(i string) []string {
	i = s.Replace(i,"(", " ( ",-1)
	i = s.Replace(i,")", " ) ",-1)
	i = s.Replace(i,"[", " [ ",-1)
	i = s.Replace(i,"]", " ] ",-1)
	i = s.Replace(i,"{", " { ",-1)
	i = s.Replace(i,"}", " } ",-1)

	return s.Split(i," ")
}

type AstNode struct {
	
	SourcePosition string

	Value t.HLValue
}

var IntLiteral = r.MustCompile("^(-)?[0-9]+")
var OpeningParen = r.MustCompile("(")
var ClosingParen = r.MustCompile(")")

func Parse(tokens []string) (nodes []AstNode) {
	for i,s := range tokens {


		if IntLiteral.MatchString(s) {
			val, err := strconv.Atoi(s)
			if err != nil {
				return nil
			}
			nodes = append(nodes, AstNode {
				strconv.Itoa(i), t.HLSysInt{
					val,
				},
			})
		}
	}
	return
}

