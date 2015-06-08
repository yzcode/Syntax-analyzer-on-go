package mypkg

import (
	"strings"
)

type TokenTable struct {
	Tokens []string
}

func (tt *TokenTable) Build(in []string) {
	for _, val := range in {
		tt.Tokens = append(tt.Tokens, strings.Split(val, "\t")[0])
	}
	//	for _,val:=range tt.Tokens{
	//		fmt.Println(val)
	//	}
	return
}
