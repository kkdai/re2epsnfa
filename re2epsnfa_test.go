package re2epsnfa

import "testing"

func TestSingle(t *testing.T) {
	trans := NewRe2EpsNFA("1.0.1")
	trans.StartParse()
}
