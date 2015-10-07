package re2epsnfa

import "testing"

func TestSingle(t *testing.T) {
	trans := NewRe2EpsNFA("(0+1.0)*.(e+1)")
	trans.StartParse()

	enfa := trans.GetEpsNFA()

	enfa.PrintTransitionTable()
}
