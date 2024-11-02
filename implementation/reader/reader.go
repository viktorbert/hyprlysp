package reader

import (
	s "strings"
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

type Parser struct {

}

func Parse(i string) {
	
}
