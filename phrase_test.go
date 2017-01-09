package mecabs

import (
	"reflect"
	"testing"
)

func TestPhrasePhrase(t *testing.T) {
	var ps PhraseString
	var e Phrase
	var ans Phrase

	ps = PhraseString{}
	e = Phrase{}
	ans = ps.Phrase()
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}

	ps = PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"世界\t名詞,一般,,,,,世界,セカイ,セカイ",
	}
	e = Phrase{
		{"こんにちは", "感動詞", "", "", "", "", "", "こんにちは", "コンニチハ", "コンニチワ"},
		{"世界", "名詞", "一般", "", "", "", "", "世界", "セカイ", "セカイ"},
	}
	ans = ps.Phrase()
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}
}

func TestPhraseStringPhrase(t *testing.T) {
	var p Phrase
	var e PhraseString
	var ans PhraseString

	p = Phrase{}
	e = PhraseString{}
	ans = p.PhraseString()
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}

	p = Phrase{
		{"こんにちは", "感動詞", "", "", "", "", "", "こんにちは", "コンニチハ", "コンニチワ"},
		{"世界", "名詞", "一般", "", "", "", "", "世界", "セカイ", "セカイ"},
	}
	e = PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"世界\t名詞,一般,,,,,世界,セカイ,セカイ",
	}
	ans = p.PhraseString()
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}
}
