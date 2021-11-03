package topwords

import (
	"sort"
)

func TopWords(s string, n int) []string {
	var words []string
	word := ""
	for _,v := range s{
		if v>='a' && v<='z' || v>='A' && v<='Z'{
			word += string(v)
		}else if word!="" {
			words = append(words, word)
			word = ""
		}
	}
	if word!=""{
		words = append(words, word)
	}
	cntWords := make(map[string]int,len(words))
	for _,v := range words{
		cntWords[v]++
	}
	wordsFreq := make([]string,0,len(cntWords))
	for k := range cntWords{
		wordsFreq = append(wordsFreq,k)
	}
	sort.Slice(wordsFreq, func(i,j int) bool {
		return cntWords[wordsFreq[i]]>cntWords[wordsFreq[j]]
	})
	return wordsFreq[:n]
}
