package mypkg

import (
	"fmt"
	"strings"
)

type Syntax struct {
	Prods []ProdExp
	Vn    map[string]bool
	Vt    map[string]bool
	Final string
}
type ProdExp struct {
	Ele []string
}

func (s *Syntax) Index(p ProdExp) int {
	for i, val := range s.Prods {
		if strings.EqualFold(val.String(), p.String()) {
			return i
		}
	}
	return -1
}
func (p ProdExp) String() (ret string) {
	ret = ""
	for _, val := range p.Ele {
		ret = fmt.Sprintf("%v %v", ret, val)
	}
	return
}
func (prodexps *Syntax) Build(in []string) {
	lineNumber := 0
	prodexps.Final = ""
	prodexps.Vn = map[string]bool{}
	prodexps.Vt = map[string]bool{}
	prodexps.Prods = []ProdExp{}
	for _, val := range in {
		lineNumber++
		if val[0] == '#' {
			continue
		}
		if lineNumber == 1 {
			for _, word := range strings.Split(val, " ") {
				prodexps.Vt[strings.Split(word, ":")[0]] = true
			}
		} else if lineNumber == 2 {
			prodexps.Final = val
			prodexps.Prods = append(prodexps.Prods, ProdExp{[]string{val + "'", val}})
		} else {
			tmp_exp := ProdExp{[]string{}}
			for i, exp_ele := range strings.Split(val, " ") {
				if exp_ele[0] != '$' {
					tmp_exp.Ele = append(tmp_exp.Ele, exp_ele)
				}
				if i == 0 {
					prodexps.Vn[exp_ele] = true
				}
			}
			prodexps.Prods = append(prodexps.Prods, tmp_exp)
		}
	}
	//fmt.Printf("%v\n",prodexps.Final)
	//	for key,_:=range prodexps.Vn{
	//		fmt.Printf("%s\n",key)
	//	}
	return
}
