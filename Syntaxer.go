package main

import (
	"bufio"
	"fmt"
	"github.com/yzcode/SAOG/mypkg"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	syn_file := "syntax.txt"
	token_tabel_file := "lex_table.txt"
	run_syn := mypkg.Syntax{}
	run_token_tabel := mypkg.TokenTable{}
	for i, val := range os.Args {
		if i == 1 {
			syn_file = val
		} else if i == 2 {
			token_tabel_file = val
		}
	}
	prd_str := make([]string, 0)
	if fin, err := os.Open(syn_file); err == nil {
		defer fin.Close()
		fin := bufio.NewReader(fin)
		for {
			if instr, inerr := fin.ReadString('\n'); inerr != io.EOF {
				instr = strings.Replace(instr, "\r\n", "", 1)
				prd_str = append(prd_str, instr)
			} else {
				break
			}
		}
		run_syn.Build(prd_str)
		prd_str = []string{}
		if ttin, err := os.Open(token_tabel_file); err == nil {
			defer ttin.Close()
			ttin := bufio.NewReader(ttin)
			for {
				if instr, inerr := ttin.ReadString('\n'); inerr != io.EOF {
					instr = strings.Replace(instr, "\r\n", "", 1)
					prd_str = append(prd_str, instr)
				} else {
					break
				}
			}
			run_token_tabel.Build(prd_str)
			run_dfa := mypkg.CreateLRDFA(run_syn)
			lrtable := [200]map[string][]string{}
			for i := 0; i < len(run_dfa.States); i++ {
				node := run_dfa.States[i]
				for key, _ := range mypkg.Alphabet {
					if lrtable[i] == nil {
						lrtable[i] = map[string][]string{}
					}
					lrtable[i][key] = []string{}
					t := node.GetNextsta(key)
					if t != nil {
						if run_syn.Vt[key] {
							lrtable[i][key] = append(lrtable[i][key], fmt.Sprintf("S%d", t.Index))
						} else {
							lrtable[i][key] = append(lrtable[i][key], fmt.Sprintf("%d", t.Index))
						}
					}
				}
				for _, val := range node.Data.Items {
					pos := val.Step
					production := val.Pord
					ahead := val.Ahead
					if pos == len(production.Ele)-1 {
						for _, t := range ahead.Ele {
							lrtable[i][t] = append(lrtable[i][t], fmt.Sprintf("r%d", run_syn.Index(production)))
						}
						if len(production.Ele) == 2 && strings.EqualFold("S'", production.Ele[0]) && strings.EqualFold("S", production.Ele[1]) {
							flag_tmp := false
							for _, t := range ahead.Ele {
								if strings.EqualFold("#", t) {
									flag_tmp = true
									break
								}
							}
							if flag_tmp {
								lrtable[i]["#"] = append(lrtable[i]["#"], "acc")
							}
						}
					}
				}
			}
			////lr table output
			//			fmt.Print("\t")
			//			for _,_val:=range mypkg.StaticAlpha{
			//				fmt.Printf("%s\t",_val)
			//			}
			//			fmt.Printf("\n");
			//			for i,_:=range run_dfa.States{
			//				fmt.Printf("%v\t",i)
			//				for _,__val:=range mypkg.StaticAlpha{
			//					fmt.Printf("%v\t",lrtable[i][__val])
			//				}
			//				fmt.Printf("\n")
			//			}
			////lr table output end
			stateStack := []int{0}
			symbolStack := []string{}
			ptr := 0
			flag_suc := false
			fmt.Println("----syntax start ----")
			for !flag_suc {
				fmt.Println("-------")
				fmt.Println(symbolStack)
				now_token := run_token_tabel.Tokens[ptr]
				fmt.Println(now_token)
				head := stateStack[len(stateStack)-1]
				fmt.Println(head)
				action := lrtable[head][now_token]
				fmt.Printf("%v\n", action)
				if len(action) == 0 {
					fmt.Printf("syntax error\nunexpected symbol %s \n", now_token)
					break
				}
				for _, val := range action {
					if strings.EqualFold(val, "acc") {
						fmt.Println("accepted")
						flag_suc = true
						break
					}
				}
				if flag_suc {
					break
				}
				now_action := action[0]
				if now_action[0] == 'S' {
					anss, _ := strconv.Atoi(now_action[1:])
					stateStack = append(stateStack, anss)
					symbolStack = append(symbolStack, now_token)
					ptr++
				} else if now_action[0] == 'r' {
					tmp_ptr, _ := strconv.Atoi(now_action[1:])
					tmp_len := len(run_syn.Prods[tmp_ptr].Ele) - 1
					fmt.Println(tmp_ptr)
					fmt.Printf("len of prod : %d len of stateStack : %d symbolStack : %d\n", tmp_len, len(stateStack), len(symbolStack))
					if len(run_syn.Prods[tmp_ptr].Ele)-1 > 0 {
						stateStack = stateStack[:len(stateStack)-tmp_len]
						symbolStack = symbolStack[:len(symbolStack)-tmp_len]
					}
					tmp_next, _ := strconv.Atoi(lrtable[stateStack[len(stateStack)-1]][run_syn.Prods[tmp_ptr].Ele[0]][0])
					stateStack = append(stateStack, tmp_next)
					symbolStack = append(symbolStack, run_syn.Prods[tmp_ptr].Ele[0])
				}
			}

		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
	}
}
