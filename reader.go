package main

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
