package mecabs

// Phrase は文を Morpheme の列で保存したものです。
type Phrase []Morpheme

// Surface は Morpheme 列の表層形をつなげたものです。
func (ph Phrase) Surface() (ans string) {
	for _, m := range ph {
		ans += m.Surface()
	}
	return
}

// Reading は Morpheme 列の読みをつなげたものです。
// Reading が空文字列の Morpheme が存在する場合は
// ok が false になります。
func (ph Phrase) Reading() (ans string, ok bool) {
	ok = true
	for _, m := range ph {
		ans += m.Reading()
		if m.Reading() == "" {
			ok = false
		}
	}
	return
}

// Pronounciation は Morpheme 列の発音をつなげたものです。
// Pronounciation が空文字列の Morpheme が存在する場合は
// ok が false になります。
func (ph Phrase) Pronounciation() (ans string, ok bool) {
	ok = true
	for _, m := range ph {
		ans += m.Pronounciation()
		if m.Pronounciation() == "" {
			ok = false
		}
	}
	return
}
