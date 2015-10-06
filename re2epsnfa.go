package re2epsnfa

import (
	"fmt"
	"strconv"

	. "github.com/kkdai/e-nfa"
)

const epsilon = 2 //total symbol, 0, 1, 2(epsilon)

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
	newEdge := Edge{SrcState: stateSrc, Input: cInput, DstState: stateDst}
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

func (r *Re2EpsNFA) calculateNext(re string) {
	reLength := len(re)
	for i := 0; i <= reLength; i++ {
		if re[i] == '(' {
			k := 0
			j := i

			for {
				if re[j] == '(' {
					k = k + 1
				}

				if re[j] == ')' {
					k = k - 1
				}

				if k == 0 {
					break
				}
				j = j + 1
			}
		} else {
			r.nextParentheses[i] = i
		}
	}
}

func (r *Re2EpsNFA) parse(re string, s, t int) (int, int) {

	//single symbol
	if s == t {
		newStart := r.incCapacity()
		newFinal := r.incCapacity()

		if re[s] == 'e' {
			r.addEdge(newStart, epsilon, newFinal)
		} else {
			num, _ := strconv.Atoi(string(re[s]))
			r.addEdge(newStart, num, newFinal)
		}
	}

	//(...)
	if re[s] == '(' && re[t] == ')' {
		if r.nextParentheses[s] == t {
			return r.parse(re, s+1, t-1)
		}
	}

	//RE1+RE2
	i := s
	for {
		i = r.nextParentheses[i]

		if i <= t && re[i] == '+' {
			s1, t1 := r.parse(re, s, i-1)
			s2, t2 := r.parse(re, i+1, t)
			retS, retF := r.union(s1, s2, t1, t2)
			return retS, retF
		}
		i = i + 1
	}

	//RE1.RE2
	i = s
	for {
		i = r.nextParentheses[i]

		if i <= t && re[i] == '.' {
			s1, t1 := r.parse(re, s, i-1)
			s2, t2 := r.parse(re, i+1, t)

			retS, retF := r.concatenation(s1, s2, t1, t2)
			return retS, retF

		}
		i = i + 1
	}

	//(RE)*
	s1, t1 := r.parse(re, s, t-1)
	retS, retF := r.closure(s1, t1)
	return retS, retF
}

func (r *Re2EpsNFA) StartParse() string {
	var result []byte
	nfaStart, nfaFinal := r.parse(r.regexString, 0, len(r.regexString)-1)
	fmt.Printf("new NFA s=%d, f=%d\n", nfaStart, nfaFinal)

}

func (r *Re2EpsNFA) GetEpsNFA() *ENFA {
	return nil
}
