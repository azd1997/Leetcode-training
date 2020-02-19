package lt804

import "strings"

// 唯一莫尔斯密码词


var codes = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

func uniqueMorseRepresentations(words []string) int {
	set := make(map[string]bool)
	for _, word := range words {
		coded := strings.Builder{}
		for i:=0; i<len(word); i++ {
			coded.WriteString(codes[word[i] - 'a'])
		}
		if !set[coded.String()] {
			set[coded.String()] = true
		}
	}
	return len(set)
}
