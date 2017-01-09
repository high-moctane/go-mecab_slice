package mecabs

import (
	"reflect"
	"testing"
)

func TestConst(t *testing.T) {
	if BOM.MorphemeString() != BOMS {
		t.Errorf("different in BOM and BOMS")
	}
	if EOM.MorphemeString() != EOMS {
		t.Errorf("different in EOM and EOMS")
	}
}

func TestNew(t *testing.T) {
	m, err := New(map[string]string{})
	if err != nil {
		t.Fatalf("unexpected error")
	}
	defer m.Destroy()
}

func TestNewPhraseString(t *testing.T) {
	m, err := New(map[string]string{})
	if err != nil {
		t.Fatalf("unexpected error")
	}
	defer m.Destroy()

	var ps PhraseString
	var e PhraseString

	ps, err = m.NewPhraseString("")
	if err != nil {
		t.Fatalf("unexpected error")
	}
	e = PhraseString{}
	if !reflect.DeepEqual(ps, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ps)
	}

	ps, err = m.NewPhraseString("こんにちは世界")
	if err != nil {
		t.Fatalf("unexpected error")
	}
	e = PhraseString{
		"こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ",
		"世界\t名詞,一般,,,,,世界,セカイ,セカイ",
	}
	if !reflect.DeepEqual(ps, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ps)
	}
}

func TestNewPhrase(t *testing.T) {
	m, err := New(map[string]string{})
	if err != nil {
		t.Fatalf("unexpected error")
	}
	defer m.Destroy()

	var p Phrase
	var e Phrase

	p, err = m.NewPhrase("")
	if err != nil {
		t.Fatalf("unexpected error")
	}
	e = Phrase{}
	if !reflect.DeepEqual(p, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, p)
	}

	p, err = m.NewPhrase("こんにちは世界")
	if err != nil {
		t.Fatalf("unexpected error")
	}
	e = Phrase{
		{"こんにちは", "感動詞", "", "", "", "", "", "こんにちは", "コンニチハ", "コンニチワ"},
		{"世界", "名詞", "一般", "", "", "", "", "世界", "セカイ", "セカイ"},
	}
	if !reflect.DeepEqual(p, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, p)
	}
}
