package mecabs

import (
	"testing"

	"reflect"
)

func TestNewMeCabS(t *testing.T) {
	mecabs, err := New(map[string]string{"output-format-type": "wakati", "all-morphs": ""})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	defer mecabs.Destroy()
}

func TestParse(t *testing.T) {
	mecabs, err := New(map[string]string{})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	defer mecabs.Destroy()

	result, err := mecabs.Parse("こんにちは世界")
	if err != nil {
		t.Errorf("parse error: %v", err)
		return
	}
	expected := [][10]string{
		{"こんにちは", "感動詞", "", "", "", "", "", "こんにちは", "コンニチハ", "コンニチワ"},
		{"世界", "名詞", "一般", "", "", "", "", "世界", "セカイ", "セカイ"},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("want %v, but %v", expected, result)
	}
}

func BenchmarkParse(b *testing.B) {
	mecabs, _ := New(map[string]string{})
	defer mecabs.Destroy()

	for i := 0; i < b.N; i++ {
		mecabs.Parse("ある日の暮方の事である。一人の下人げにんが、羅生門らしょうもんの下で雨やみを待っていた。　広い門の下には、この男のほかに誰もいない。ただ、所々丹塗にぬりの剥はげた、大きな円柱まるばしらに、蟋蟀きりぎりすが一匹とまっている。羅生門が、朱雀大路すざくおおじにある以上は、この男のほかにも、雨やみをする市女笠いちめがさや揉烏帽子もみえぼしが、もう二三人はありそうなものである。それが、この男のほかには誰もいない。")
	}
}
