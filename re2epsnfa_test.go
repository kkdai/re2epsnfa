package re2epsnfa

import "testing"

func TestSimpleRE(t *testing.T) {
	trans := NewRe2EpsNFA("1.0.1")
	trans.StartParse()
	enfa := trans.GetEpsNFA()
	enfa.PrintTransitionTable()
}

func TestConcateRE(t *testing.T) {
	trans := NewRe2EpsNFA("0+1.0.1")
	trans.StartParse()
	enfa := trans.GetEpsNFA()
	enfa.PrintTransitionTable()
}

func TestKleeneStarRE(t *testing.T) {
	trans := NewRe2EpsNFA("(1.0+0.0+1.1)*")
	trans.StartParse()
	enfa := trans.GetEpsNFA()
	enfa.PrintTransitionTable()
}

func TestComplicateRE(t *testing.T) {
	trans := NewRe2EpsNFA("(0+1.0)*.(e+1)")
	trans.StartParse()
	enfa := trans.GetEpsNFA()
	enfa.PrintTransitionTable()
}
