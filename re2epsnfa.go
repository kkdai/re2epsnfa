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
}

func NewRe2EpsNFA(str string) *Re2EpsNFA {
	newRe2NFA := &Re2EpsNFA{regexString: str}
	newRe2NFA.edgeMap = make(map[Edge]bool)
	return newRe2NFA
}

func (r *Re2EpsNFA) union(s1, s2, t1, t2 int) (int, int) {

}

func (r *Re2EpsNFA) concatenation(s1, s2, t1, t2 int) (int, int) {
}

func (r *Re2EpsNFA) closure() (int, int) {
}

func (r *Re2EpsNFA) parse() (int, int) {
}

func (r *Re2EpsNFA) GetEpsNFA() *ENFA {

}
