package mypkg

type DFANode struct {
	Dfa   *DFA
	Index int
	Next  map[string]*DFANode
	Data  DFAItemSet
}
type DFA struct {
	States   []DFANode
	Start    *DFANode
	AlphaBet map[string]bool
	size     int
	//Final_state map[DFANode]bool
}

func (dfanode *DFANode) AddNextsta(on string, v *DFANode) {
	if dfanode.Next == nil {
		dfanode.Next = map[string]*DFANode{}
	}
	dfanode.Next[on] = v
	return
}
func (dfanode *DFANode) GetNextsta(on string) *DFANode {
	return dfanode.Next[on]
}
func (dfa *DFA) CreateNode() *DFANode {
	dfa.States = append(dfa.States, DFANode{dfa, dfa.size, map[string]*DFANode{}, DFAItemSet{}})
	dfa.size++
	return &dfa.States[dfa.size-1]
}
func (dfa *DFA) AddNext(u *DFANode, on string, v *DFANode) {
	u.AddNextsta(on, v)
	return
}
func (dfa *DFA) GetNext(u *DFANode, on string) *DFANode {
	return u.GetNextsta(on)
}
