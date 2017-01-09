package mecabs

import "strings"

type Output string

func (o Output) PhraseString() PhraseString {
	lines := strings.Split(string(o), "\n")
	lines = lines[:len(lines)-2]
	ans := make(PhraseString, len(lines))
	for i, line := range lines {
		ans[i] = MakeMorphemeString(line)
	}
	return ans
}

func (o Output) Phrase() Phrase {
	return o.PhraseString().Phrase()
}
