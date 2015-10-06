package re2epsnfa

import (
	"fmt"
	"strconv"

	. "github.com/kkdai/e-nfa"
)

const epsilon = 2 //total symbol, 0, 1, 2(epsilon)
const maxNumState = 200

type Edge struct {
	Src   int
	Input int
	Dst   int
}

type Closure struct {
	Src int
	Dst int
}

type Re2EpsNFA struct {
	regexString     string
	nextParentheses []int
	edgeMap         map[Edge]bool
	stateCount      int
	closureMap      map[Closure]bool
}

func NewRe2EpsNFA(str string) *Re2EpsNFA {
	newRe2NFA := &Re2EpsNFA{regexString: str}
	newRe2NFA.edgeMap = make(map[Edge]bool)
	newRe2NFA.closureMap = make(map[Closure]bool)
	return newRe2NFA
}

func (r *Re2EpsNFA) incCapacity() int {
	r.stateCount = r.stateCount + 1
	return r.stateCount - 1
}

func (r *Re2EpsNFA) addEdge(stateSrc int, cInput int, stateDst int) {
	newEdge := Edge{Src: stateSrc, Input: cInput, Dst: stateDst}
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
	for i := 0; i < reLength; i++ {
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
			r.nextParentheses = append(r.nextParentheses, j)

		} else {
			r.nextParentheses = append(r.nextParentheses, i)
		}
	}
}

func (r *Re2EpsNFA) checkClosureExist(src, target int) bool {
	closureExist, _ := r.closureMap[Closure{Src: src, Dst: target}]
	return closureExist
}

func (r *Re2EpsNFA) checkPathExist(src, input, dst int) bool {

	pathExist, _ := r.edgeMap[Edge{Src: src, Input: input, Dst: dst}]
	return pathExist
}

func (r *Re2EpsNFA) calcClosure() {
	queue := make([]int, 200)

	for i := 0; i <= r.stateCount; i++ {
		for j := 0; j < r.stateCount; j++ {
			r.closureMap[Closure{Src: i, Dst: j}] = true
		}

		head := -1
		tail := 0
		queue[0] = i

		r.closureMap[Closure{Src: i, Dst: i}] = true

		for head < tail {
			head = head + 1
			j := queue[head]

			for k := 0; k < r.stateCount; k++ {
				closureExist := r.checkClosureExist(i, k)
				pathExist := r.checkPathExist(j, epsilon, k)

				if !closureExist && pathExist {
					tail = tail + 1
					queue[tail] = k

					r.closureMap[Closure{Src: i, Dst: k}] = true
				}
			}
		}
	}
}

func (r *Re2EpsNFA) parse(re string, s, t int) (int, int) {
	fmt.Println("Parse  s=", s, " t=", t)

	//single symbol
	if s == t {
		newStart := r.incCapacity()
		newFinal := r.incCapacity()
		fmt.Println("single symbol=", newStart, newFinal)

		if re[s] == 'e' {
			r.addEdge(newStart, epsilon, newFinal)
		} else {
			num, _ := strconv.Atoi(string(re[s]))
			r.addEdge(newStart, num, newFinal)
			fmt.Println("single: addEdge", newStart, num, newFinal)
		}
		return newStart, newFinal
	}

	//(...)
	if re[s] == '(' && re[t] == ')' {
		if r.nextParentheses[s] == t {
			return r.parse(re, s+1, t-1)
		}
	}

	//RE1+RE2
	i := s
	for i <= t {
		i = r.nextParentheses[i]

		if i <= t && re[i] == '+' {
			fmt.Println("RE1+RE2")
			s1, t1 := r.parse(re, s, i-1)
			s2, t2 := r.parse(re, i+1, t)
			retS, retF := r.union(s1, s2, t1, t2)
			return retS, retF
		}
		i = i + 1
	}

	//RE1.RE2
	i = s
	for i <= t {
		i = r.nextParentheses[i]

		if i <= t && re[i] == '.' {
			fmt.Println("RE1.RE2")
			s1, t1 := r.parse(re, s, i-1)
			s2, t2 := r.parse(re, i+1, t)

			fmt.Println("concate: ", s1, s2, t1, t2)
			retS, retF := r.concatenation(s1, s2, t1, t2)
			return retS, retF

		}
		i = i + 1
	}

	//(RE)*
	fmt.Println("go (RE)*")
	s1, t1 := r.parse(re, s, t-1)
	retS, retF := r.closure(s1, t1)
	return retS, retF
}

func (r *Re2EpsNFA) StartParse() string {
	var result []byte
	fmt.Println(r.regexString)
	r.calculateNext(r.regexString)
	fmt.Println("Next Pa=", r.nextParentheses)
	nfaStart, nfaFinal := r.parse(r.regexString, 0, len(r.regexString)-1)
	fmt.Printf("new NFA s=%d, f=%d\n", nfaStart, nfaFinal)

	r.calcClosure()

	for length := 1; length < 7; length++ {
		for num := 0; uint(num) < (uint(1) << uint(length)); num++ {

		}
	}

	return string(result)
}

func (r *Re2EpsNFA) GetEpsNFA() *ENFA {
	return nil
}
