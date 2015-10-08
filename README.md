RE2EPSNFA: A tranform function to translate RE to Epsilon-NFA (Epsilon-Nondeterministic finite automaton)
==============

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/e-nfa/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/re2epsnfa?status.svg)](https://godoc.org/github.com/kkdai/re2epsnfa)  [![Build Status](https://travis-ci.org/kkdai/re2epsnfa.svg?branch=master)](https://travis-ci.org/kkdai/re2epsnfa)



![image](images/nfa_dfa_re_convert.png)

What is this package 
=============

`re2epsnfa` (Regular Expression to Eplison-NFA) will help you to convert regular expression string to [epsilon-NFA](http://github.com/kkdai/enfa) object.

####The convert rule base on three simple rule:

![image](images/RENFA1.png)

`union`: ex: A union B



![image](images/RENFA2.png)

`concatenation`:  ex: A concatenation B
    
    
    
![image](images/RENFA3.png)    

`closure`: ex: Closure(A)

 
What is Epsilon-Nondeterministic finite automaton
=============

`ε-NFA`: Epsilon-Nondeterministic finite automaton (so call:Nondeterministic finite automaton with ε-moves)

In the automata theory, a nondeterministic finite automaton with ε-moves (NFA-ε)(also known as NFA-λ) is an extension of nondeterministic finite automaton(NFA), which allows a transformation to a new state without consuming any input symbols. The transitions without consuming an input symbol are called ε-transitions or λ-transitions. In the state diagrams, they are usually labeled with the Greek letter ε or λ.

(sited from [here](https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton))



Installation and Usage
=============


Install
---------------

    go get github.com/kkdai/re2epsnfa


Usage
---------------

Following is sample code to implement a epsilon-NFA automata diagram as follow:


```go

package main

import (
    "github.com/kkdai/re2epsnfa"
)

func main() {

	trans := NewRe2EpsNFA("(0+1.0)*.(e+1)")
	trans.StartParse()
	enfa := trans.GetEpsNFA()
	enfa.PrintTransitionTable()

}

```

Inspired By
=============

- [ε-NFA: Wiki](https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton_with_%CE%B5-moves)
- [Coursera: Automata](https://class.coursera.org/automata-004/)
- [PDA: Constructing -NFA’s From Regular Expressions](https://people.cs.umass.edu/~immerman/cs250/regExpToNFA.pdf)
- [Youtube: Regular Expression to NFA](https://www.youtube.com/watch?v=RYNN-tb9WxI)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.
