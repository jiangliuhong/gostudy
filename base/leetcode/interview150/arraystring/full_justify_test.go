package arraystring

import (
	"fmt"
	"testing"
)

func TestFullJustify(t *testing.T) {
	//words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	justify := fullJustify(words, 20)
	//fmt.Println(justify)
	for _, item := range justify {
		fmt.Println(item)
	}
}
