package mypkg

import (
	"fmt"
	"sort"
)

var maxlen = 0

type DFAItem struct {
	Step  int
	Pord  ProdExp
	Ahead ProdExp
}
type DFAItemSet struct {
	Items     []DFAItem
	items_str map[string]bool
}

func (iset *DFAItemSet) String() (ret string) {
	tmp_con := []string{}
	for _, val := range iset.Items {
		//fmt.Println(fmt.Sprintf("%s",val.String()))
		tmp_con = append(tmp_con, fmt.Sprintf("%s", val.String()))
	}
	//fmt.Println(tmp_con)
	//fmt.Printf("len : %d \n",len(tmp_con))
	sort.Strings(tmp_con)
	ans := 1
	seed := 131313
	for _, val := range tmp_con {
		for _, value := range val {
			ans = ans*seed + int(value)
			ans %= 1000000000 + 7
		}
	}
	ret = fmt.Sprintf("%d", ans)
	return
}
func (item *DFAItem) String() (ret string) {
	ans := item.Step
	seed := 131313
	for _, val := range item.Pord.String() {
		ans = ans*seed + int(val)
		ans %= 1000000000 + 7
	}
	tmp_str := []string{}
	for _, val := range item.Ahead.Ele {
		tmp_str = append(tmp_str, val)
	}
	sort.Strings(tmp_str)
	for _, val := range tmp_str {
		for _, chr := range val {
			ans = ans*seed + int(chr)
			ans %= 1000000000 + 7
		}
	}
	//fmt.Printf("%d\n",maxlen)
	//fmt.Printf("%d\n",len(fmt.Sprintf("%s$%s$%d", item.ahead, item.pord, item.step)))
	return fmt.Sprintf("%d", ans)
}
func (item *DFAItemSet) Print() {
	fmt.Println("----------")
	for _, val := range item.Items {
		fmt.Printf("step: %d production: ", val.Step)
		for _, valitem := range val.Pord.Ele {
			fmt.Printf("%v ", valitem)
		}
		fmt.Printf(" ahead: ")
		for _, valitem := range val.Ahead.Ele {
			fmt.Printf("%v ", valitem)
		}
		fmt.Printf("\n")
	}
	fmt.Println("----------")
}
