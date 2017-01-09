package mecabs

import (
	"reflect"
	"testing"
)

func TestMakeMorphemeString(t *testing.T) {
	var s string
	var e MorphemeString
	var ans MorphemeString

	s = "こんにちは\t感動詞,*,*,*,*,*,こんにちは"
	e = MorphemeString("こんにちは\t感動詞,,,,,,こんにちは,,")
	ans = MakeMorphemeString(s)
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}

	s = "こんにちは\t感動詞,*,*,*,*,*,こんにちは,コンニチハ,コンニチワ"
	e = MorphemeString("こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ")
	ans = MakeMorphemeString(s)
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}
}

func TestMorpheme(t *testing.T) {
	var ms MorphemeString
	var e Morpheme
	var ans Morpheme

	ms = MakeMorphemeString("こんにちは\t感動詞,,,,,,こんにちは")
	e = Morpheme{"こんにちは", "感動詞", "", "", "", "", "", "こんにちは", "", ""}
	ans = ms.Morpheme()
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}

	ms = MakeMorphemeString("こんにちは\t感動詞,,,,,,こんにちは,コンニチハ,コンニチワ")
	e = Morpheme{"こんにちは", "感動詞", "", "", "", "", "", "こんにちは", "コンニチハ", "コンニチワ"}
	ans = ms.Morpheme()
	if !reflect.DeepEqual(ans, e) {
		t.Errorf("expected\n%v\nbut\n%v", e, ans)
	}
}
