package wordLadder

import "testing"

/*           1
           /   \
          2     3
         /      / \
	    4      2   4
	    	  /
	         4
*/

func TestWordLadder(t *testing.T) {

	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "lot", "cog"}
	result := ladderLength(beginWord, endWord, wordList)

	if result != 1 {
		t.Errorf("Expected 3, but it was %d instead.", result)
	}
}
