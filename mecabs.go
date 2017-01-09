package mecabs

import "github.com/shogo82148/go-mecab"

const (
	BOMS = "\tBOS,,,,,,,,"
	EOMS = "\tEOS,,,,,,,,"
)

var (
	BOM Morpheme = Morpheme{PartOfSpeech: "BOS"}
	EOM Morpheme = Morpheme{PartOfSpeech: "EOS"}
)

type MeCabS struct {
	mecab.MeCab
}

func New(args map[string]string) (MeCabS, error) {
	tagger, err := mecab.New(args)
	if err != nil {
		return MeCabS{}, err
	}
	return MeCabS{MeCab: tagger}, nil
}

func (m *MeCabS) NewPhraseString(s string) (PhraseString, error) {
	parsed, err := m.Parse(s)
	if err != nil {
		return PhraseString{}, err
	}
	return Output(parsed).PhraseString(), nil
}

func (m *MeCabS) NewPhrase(s string) (Phrase, error) {
	parsed, err := m.Parse(s)
	if err != nil {
		return Phrase{}, err
	}
	return Output(parsed).Phrase(), nil
}
