package main
import "testing"
 
func TestDay2(t *testing.T) {
	if (!IsPasswordValid2("1-3 b: cbbdefg")) {
		t.Fail()
	}
}