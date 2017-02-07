package mecabs

import "strings"

// 各形態素に含まれる情報の長さです。
const MorPropLen = 10

var (
	BOS = Morpheme{"", "BOS", "", "", "", "", "", "", "", ""}
	EOS = Morpheme{"", "EOS", "", "", "", "", "", "", "", ""}
)

// 形態素情報を記録した配列です。
type Morpheme [MorPropLen]string

// 形態素の表層形です。
func (m *Morpheme) Surface() string {
	return m[0]
}

// 形態素の品詞です。
func (m *Morpheme) PartOfSpeech() string {
	return m[1]
}

// 形態素の品詞細分類1です。
func (m *Morpheme) PartOfSpeechSection1() string {
	return m[2]
}

// 形態素の品詞細分類2です。
func (m *Morpheme) PartOfSpeechSection2() string {
	return m[3]
}

// 形態素の品詞細分類3です。
func (m *Morpheme) PartOfSpeechSection3() string {
	return m[4]
}

// 形態素の活用形1です。
func (m *Morpheme) ConjugatedForm1() string {
	return m[5]
}

// 形態素の活用形2です。
func (m *Morpheme) ConjugatedForm2() string {
	return m[6]
}

// 形態素の原型です。
func (m *Morpheme) Inflection() string {
	return m[7]
}

// 形態素の読みです。
func (m *Morpheme) Reading() string {
	return m[8]
}

// 形態素の発音です。
func (m *Morpheme) Pronounciation() string {
	return m[9]
}

// MeCab のフォーマットどおりに出力します
func (m *Morpheme) String() string {
	return m[0] + "\t" + strings.Join(m[1:], ",")
}
