package mecabs

import (
	"reflect"
	"testing"
)

func TestPhraseString(t *testing.T) {
	var o Output
	var e PhraseString
	var ans PhraseString

	o = "EOS\n"
	e = PhraseString{}
	ans = o.PhraseString()
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}

	o = "こんにちは\t感動詞,*,*,*,*,*,こんにちは,コンニチハ,コンニチワ\n世界\t名詞,一般,*,*,*,*,世界,セカイ,セカイ\nEOS\n"
	e = PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"世界\t名詞,一般,,,,,世界,セカイ,セカイ",
	}
	ans = o.PhraseString()
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}
}

func TestPhraseOutput(t *testing.T) {
	var o Output
	var e Phrase
	var ans Phrase

	o = "EOS\n"
	e = Phrase{}
	ans = o.Phrase()
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}

	o = "こんにちは\t感動詞,*,*,*,*,*,こんにちは,コンニチハ,コンニチワ\n世界\t名詞,一般,*,*,*,*,世界,セカイ,セカイ\nEOS\n"
	e = Phrase{
		{"こんにちは", "感動詞", "", "", "", "", "", "こんにちは", "コンニチハ", "コンニチワ"},
		{"世界", "名詞", "一般", "", "", "", "", "世界", "セカイ", "セカイ"},
	}
	ans = o.Phrase()
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}
}
