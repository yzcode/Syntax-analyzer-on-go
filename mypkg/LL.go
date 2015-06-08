package mypkg

//
//import (
//	"fmt"
//	"regexp"
//	"strings"
//)
//
//var regs = regexp.MustCompile(`"(.*)"`)
//
//type LL struct {
//	prds  []ProdExp
//	Vn    map[string]bool
//	Vt    map[string]bool
//	X     map[string]int
//	First map[string]map[int32]bool
//}
//
//func (ll *LL) ReadPordExp(ins []string) (bool, string) {
//	ll.prds = make([]ProdExp, 0)
//	ll.Vn = map[string]bool{}
//	ll.Vt = map[string]bool{}
//	for _, val := range ins {
//		if len(val) == 0 {
//			continue
//		}
//		tmp_str := strings.Split(val, " = ")
//		ll.prds = append(ll.prds, ProdExp{tmp_str[0], tmp_str[1]})
//	}
//	for _, val := range ll.prds {
//		if ll.Vn[val.left] == false {
//			ll.Vn[val.left] = true
//		}
//		for _, vals := range strings.Split(val.right, " ") {
//			if vals[0] == '"' {
//				ll.Vt[vals] = true
//				continue
//			} else {
//				if ll.Vn[vals] == false && vals[0] != '$' {
//					ll.Vn[vals] = true
//				}
//			}
//		}
//	}
//	return true, ""
//}
//func (ll *LL) XDealer() (bool, string) {
//	ll.X = make(map[string]int)
//	isok := map[ProdExp]bool{}
//	for _, val := range ll.prds {
//		isok[val] = true
//		tmp_right := strings.Split(val.right, " ")
//		if len(tmp_right) == 1 && tmp_right[0][0] == '$' {
//			isok[val] = false
//			ll.X[val.left] = 1
//		}
//		for _, vals := range tmp_right {
//			if vals[0] == '"' {
//				isok[val] = false
//			}
//		}
//	}
//	tmp_isok := map[string]bool{}
//	for _, val := range ll.prds {
//		if isok[val] == false {
//			continue
//		}
//		tmp_isok[val.left] = true
//	}
//	for key, val := range ll.Vn {
//		if val {
//			if tmp_isok[key] == false {
//				if ll.X[key] == 0 {
//					ll.X[key] = -1
//				}
//			}
//		}
//	}
//	//	for key,val:=range ll.Vn{
//	//		if val{
//	//			fmt.Printf("Vn %s is %v\n",key,ll.X[key])
//	//		}
//	//	}
//	flag_x := true
//	//	for _,val:=range ll.prds  {
//	//		fmt.Printf("%v status is %v\n",val,isok[val])
//	//	}
//	for flag_x {
//		flag_x_tmp := false
//		for _, val := range ll.prds {
//			if isok[val] == false {
//				continue
//			}
//			cnt := 0
//			tmp_right := strings.Split(val.right, " ")
//			for _, vals := range tmp_right {
//				if ll.X[vals] == 0 {
//					cnt++
//				} else if ll.X[vals] == -1 {
//					//fmt.Printf("--------%v\n",val)
//					isok[val] = false
//					tmp_isok := map[string]bool{}
//					for _, valss := range ll.prds {
//						//fmt.Printf("%v status is %v\n",valss,isok[valss])
//						if isok[valss] == false {
//							continue
//						}
//						tmp_isok[valss.left] = true
//					}
//					for key, valss := range ll.Vn {
//						if valss {
//							if tmp_isok[key] == false {
//								if ll.X[key] == 0 {
//									ll.X[key] = -1
//									flag_x_tmp = true
//								}
//							}
//						}
//					}
//					break
//				}
//			}
//			if isok[val] && cnt == 0 {
//				if ll.X[val.left] == 0 {
//					ll.X[val.left] = 1
//				}
//				flag_x_tmp = true
//				for _, valss := range ll.prds {
//					if isok[valss] == false {
//						continue
//					}
//					if strings.EqualFold(valss.left, val.left) {
//						isok[valss] = false
//					}
//				}
//			}
//		}
//		if !flag_x_tmp {
//			flag_x = false
//		}
//	}
//	for key, val := range ll.Vn {
//		if val {
//			fmt.Printf("Vn %s is %v\n", key, ll.X[key])
//		}
//	}
//	return true, ""
//}
//func (ll *LL) FirstDealer() (bool, string) {
//	ll.First = map[string]map[int32]bool{}
//
//	return true, ""
//}
