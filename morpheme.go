package mecabs

import "strings"

type MorphemeString string

type Morpheme struct {
	OriginalForm         string
	PartOfSpeech         string
	PartOfSpeechSection1 string
	PartOfSpeechSection2 string
	PartOfSpeechSection3 string
	ConjugatedForm1      string
	ConjugatedForm2      string
	Inflection           string
	Reading              string
	Pronounciation       string
}

func MakeMorphemeString(s string) MorphemeString {
	tab := strings.Split(s, "\t")
	tab[1] = strings.Replace(tab[1], "*", "", -1)
	ans := strings.Join(tab, "\t")
	nComma := strings.Count(ans, ",")
	for i := 0; i+nComma < 8; i++ {
		ans += ","
	}
	return MorphemeString(ans)
}

func (m MorphemeString) Morpheme() Morpheme {
	tab := strings.Split(string(m), "\t")
	prop := make([]string, 9)
	for i, s := range strings.Split(tab[1], ",") {
		prop[i] = s
	}
	return Morpheme{
		OriginalForm:         tab[0],
		PartOfSpeech:         prop[0],
		PartOfSpeechSection1: prop[1],
		PartOfSpeechSection2: prop[2],
		PartOfSpeechSection3: prop[3],
		ConjugatedForm1:      prop[4],
		ConjugatedForm2:      prop[5],
		Inflection:           prop[6],
		Reading:              prop[7],
		Pronounciation:       prop[8],
	}
}

func (m *Morpheme) MorphemeString() MorphemeString {
	str := m.OriginalForm + "\t" +
		m.PartOfSpeech + "," +
		m.PartOfSpeechSection1 + "," +
		m.PartOfSpeechSection2 + "," +
		m.PartOfSpeechSection3 + "," +
		m.ConjugatedForm1 + "," +
		m.ConjugatedForm2 + "," +
		m.Inflection + "," +
		m.Reading + "," +
		m.Pronounciation
	return MorphemeString(str)
}
