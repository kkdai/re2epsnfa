package re2epsnfa

import (
	. "github.com/kkdai/e-nfa"
)

type Edge struct {
	SrcState int
	Input    int
	DstState int
}

type Re2EpsNFA struct {
	regexString     string
	nextParentheses []int
	edgeMap         map[Edge]bool
	stateCount      int
}

func NewRe2EpsNFA(str string) *Re2EpsNFA {
	newRe2NFA := &Re2EpsNFA{regexString: str}
	newRe2NFA.edgeMap = make(map[Edge]bool)
	return newRe2NFA
}

func (r *Re2EpsNFA) incCapacity() int {
	r.stateCount = r.stateCount + 1
	return r.stateCount - 1
}

func (r *Re2EpsNFA) addEdge(stateSrc int, cInput int, stateDst int) {
	newEdge = Edge{SrcState: stateSrc, Input: cInput, DstState: stateDst}
	r.edgeMap[newEdge] = true
}

func (r *Re2EpsNFA) union(s1, s2, t1, t2 int) (int, int) {
	newStartState := r.incCapacity()
	newFinalState := r.incCapacity()

	r.addEdge(newStartState, 2, s1)
	r.addEdge(newStartState, 2, s2)

	r.addEdge(t1, 2, newFinalState)
	r.addEdge(t2, 2, newFinalState)

	return newStartState, newFinalState
}

func (r *Re2EpsNFA) concatenation(s1, s2, t1, t2 int) (int, int) {
	r.addEdge(t1, 2, s2)
	return s1, t2
}

func (r *Re2EpsNFA) closure(s, t int) (int, int) {

	newStartState := r.incCapacity()
	newFinalState := r.incCapacity()

	r.addEdge(newStartState, 2, s)
	r.addEdge(t, 2, newFinalState)
	r.addEdge(t, 2, s)
	r.addEdge(newStartState, 2, newFinalState)
	return newStartState, newFinalState
}

func (r *Re2EpsNFA) parse() (int, int) {
}

func (r *Re2EpsNFA) GetEpsNFA() *ENFA {

}
