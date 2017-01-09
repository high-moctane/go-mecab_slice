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
