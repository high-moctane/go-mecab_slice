package mecabs

import "testing"

func TestMorphemeString(t *testing.T) {
	morph := Morpheme{"あ", "い", "う", "え", "お", "か", "き", "く", "け", "こ"}
	if morph.String() != "あ\tい,う,え,お,か,き,く,け,こ" {
		t.Errorf("expected \"あ\\tい,う,え,お,か,き,く,け,こ\", but %v", morph.String())
	}
}
