package mypkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var Alphabet = map[string]bool{}
var productions = map[string][]ProdExp{}
var nullable = map[string]int{}
var tmp_nullable = map[string]bool{}
var first = map[string]map[string]bool{}
var tmp_first = map[string]bool{}
var itemSetMap = map[string]*DFANode{}
var initItem = DFAItem{0, ProdExp{[]string{"S'", "S"}}, ProdExp{[]string{"#"}}}
var initSet = DFAItemSet{[]DFAItem{initItem}, map[string]bool{}}
var StaticAlpha = []string{}

func CreateLRDFA(in Syntax) (dfa DFA) {
	buildAlpha()
	for key, _ := range in.Vn {
		Alphabet[key] = true
	}
	for key, _ := range in.Vt {
		Alphabet[key] = true
	}
	for _, val := range in.Prods {
		productions[val.Ele[0]] = append(productions[val.Ele[0]], val)
	}
	initSet.items_str[initSet.Items[0].String()] = true
	closure(initItem, &initSet, in)
	//	for _,item:=range initSet.items{
	//		for _,val:=range item.ahead.ele {
	//			fmt.Printf("%s ",val)
	//		}
	//		fmt.Println()
	//	}
	dfa = DFA{size: 0}
	dfa.AlphaBet = Alphabet
	dfa.Start = dfa.CreateNode()
	dfa.Start.Data = initSet
	itemSetMap[initSet.String()] = dfa.Start
	queue := []*DFANode{dfa.Start}
	for len(queue) > 0 {
		tmpstate := queue[0]
		queue = queue[1:]
		for _, key := range StaticAlpha {
			new_item_set := DFAItemSet{items_str: map[string]bool{}}
			for _, item := range tmpstate.Data.Items {
				pos := item.Step
				production := item.Pord
				ahead := item.Ahead
				if pos+1 < len(production.Ele) && strings.EqualFold(production.Ele[pos+1], key) {
					new_item := DFAItem{pos + 1, production, ahead}
					new_item_set.Items = append(new_item_set.Items, new_item)
					new_item_set.items_str[new_item.String()] = true
					closure(new_item, &new_item_set, in)
				}
			}
			if len(new_item_set.Items) == 0 {
				continue
			}
			//var node *DFANode
			if itemSetMap[new_item_set.String()] == nil {
				node := dfa.CreateNode()
				node.Data = new_item_set
				itemSetMap[new_item_set.String()] = node
				queue = append(queue, node)
				if node == nil {
					fmt.Printf("nil err\n")
				}
			} else {
				//				fmt.Println("new -----:")
				//				new_item_set.Print()
				//				fmt.Println("old -----:")
				//				itemSetMap[new_item_set.String()].Data.Print()
			}
			//fmt.Println("--------------")
			//			if strings.EqualFold("function",key) && dfa.size<100{
			//				for _,val_item:=range new_item_set.Items{
			//					fmt.Printf("set step: %d\n",val_item.Step)
			//					for _,val_prd:=range val_item.Pord.Ele{
			//						fmt.Printf("%s ",val_prd)
			//					}
			//					fmt.Println()
			//				}
			//			}
			//fmt.Println(new_item_set.String())
			//fmt.Println(new_item_set.String())
			//fmt.Println(new_item_set.String())
			//fmt.Println("--------------")
			dfa.AddNext(tmpstate, key, itemSetMap[new_item_set.String()])
			if itemSetMap[new_item_set.String()] == nil {
				//fmt.Printf("node error %d\n",dfa.size);
			}
			fmt.Printf("%v -------%s------->%v\n", tmpstate.Index, key, itemSetMap[new_item_set.String()].Index)
		}
	}
	return
}
func buildNullable(item string, in Syntax) (ret bool) {
	if nullable[item] != 0 {
		return nullable[item] == 1
	}
	if tmp_nullable[item] {
		nullable[item] = -1
		return false
	}
	tmp_nullable[item] = true
	nullable[item] = -1
	for _, val := range productions[item] {
		tmp_flag := true
		for i, word := range val.Ele {
			if i == 0 {
				continue
			}
			if in.Vt[word] {
				tmp_flag = false
			} else {
				tmp_flag = tmp_flag && buildNullable(word, in)
			}
		}
		if (nullable[item] == 1) || tmp_flag {
			nullable[item] = 1
		}
	}
	tmp_nullable[item] = false
	return nullable[item] == 1
}

func buildFirst(item []string, in Syntax) (ret map[string]bool) {
	tmp := ProdExp{item}
	if len(first[tmp.String()]) > 0 {
		return first[tmp.String()]
	}
	tmp_first[tmp.String()] = true
	first[tmp.String()] = map[string]bool{}
	for _, val := range item {
		if in.Vt[val] {
			first[tmp.String()][val] = true
			break
		} else {
			for _, exp := range productions[val] {
				if in.Vt[exp.Ele[0]] {
					first[tmp.String()][exp.Ele[0]] = true
				} else if tmp_first[ProdExp{exp.Ele[1:]}.String()] == false {
					for key, _ := range buildFirst(exp.Ele[1:], in) {
						first[tmp.String()][key] = true
					}
				}
			}
		}
		if buildNullable(val, in) == false {
			break
		}
	}
	tmp_first[tmp.String()] = false
	return first[tmp.String()]
}

func closure(item DFAItem, itemset *DFAItemSet, in Syntax) {
	pos := item.Step
	production := item.Pord
	ahead := item.Ahead
	var rightpart []string
	if len(production.Ele) > pos+1 {
		rightpart = production.Ele[pos+1:]
	} else {
		rightpart = []string{}
	}
	if len(rightpart) == 0 || in.Vn[rightpart[0]] == false {
		return
	}
	for _, prod := range productions[rightpart[0]] {
		tmpset := map[string]bool{}
		for _, val := range ahead.Ele {
			for word, _ := range buildFirst(append(rightpart[1:], val), in) {
				tmpset[word] = true
			}
		}
		tmpstr := []string{}
		for key, _ := range tmpset {
			tmpstr = append(tmpstr, key)
		}
		newitem := DFAItem{0, prod, ProdExp{tmpstr}}
		if itemset.items_str[newitem.String()] == false {
			itemset.Items = append(itemset.Items, newitem)
			itemset.items_str[newitem.String()] = true
			closure(newitem, itemset, in)
		}
	}
	return
}
func buildAlpha() {
	symbol_file := "symbolset.txt"
	fin, err := os.Open(symbol_file)
	defer fin.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fin := bufio.NewReader(fin)
		for {
			if instr, inerr := fin.ReadString('\n'); inerr != io.EOF {
				instr = strings.Replace(instr, "\r\n", "", 1)
				StaticAlpha = append(StaticAlpha, instr)
			} else {
				break
			}
		}
	}
	fmt.Println("------")
	//	for _,val:=range StaticAlpha{
	//		fmt.Println(val)
	//	}
}
