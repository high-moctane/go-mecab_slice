package mecabs

type PhraseString []MorphemeString

type Phrase []Morpheme

func (p PhraseString) Phrase() Phrase {
	ans := make(Phrase, len(p))
	for i, v := range p {
		ans[i] = v.Morpheme()
	}
	return ans
}

func (p Phrase) PhraseString() PhraseString {
	ans := make(PhraseString, len(p))
	for i, v := range p {
		ans[i] = v.MorphemeString()
	}
	return ans
}

func (p Phrase) OriginalForm() (ans string) {
	for _, v := range p {
		ans += v.OriginalForm
	}
	return
}

func (p Phrase) Reading() (string, bool) {
	var ans string
	ok := true
	for _, v := range p {
		if v.Reading == "" {
			ok = false
		}
		ans += v.Reading
	}
	return ans, ok
}

func (p Phrase) Pronounciation() (string, bool) {
	var ans string
	ok := true
	for _, v := range p {
		if v.Pronounciation == "" {
			ok = false
		}
		ans = v.Pronounciation
	}
	return ans, ok
}
