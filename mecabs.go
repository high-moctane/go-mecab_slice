// mecabs は MeCab の結果をスライス・配列にして返します。
// github.com/shogo82148/go-mecab のラッパーなので
// パッケージ mecab の関数をそのまま使えます。
package mecabs

import (
	"errors"
	"strings"

	"github.com/shogo82148/go-mecab"
)

// MeCabS は mecab.MeCab を拡張したものです。
type MeCabS struct {
	mecab.MeCab
}

// New は初期化された MeCabS 構造体を返します。
// args は mecab.New() に渡すパラメータです
// 終了時には mecab.Destroy() を呼び出すようにしてください。
func New(args map[string]string) (MeCabS, error) {
	tagger, err := mecab.New(args)
	if err != nil {
		return MeCabS{}, errors.New("MeCabS cannot be created.")
	}
	return MeCabS{MeCab: tagger}, nil
}

// ParseToPhrase は string を Phrase にパーズします。
func (m *MeCabS) ParseToPhrase(s string) (Phrase, error) {
	node, err := m.ParseToNode(s)
	if err != nil {
		return nil, err
	}

	ph := make(Phrase, 0, len([]rune(s)))
	// 最初のノードは mecab.BOSNode なので飛ばす
	for node = node.Next(); node.Stat() != mecab.EOSNode; node = node.Next() {
		var m Morpheme
		m[0] = node.Surface()
		for i, v := range strings.Split(node.Feature(), ",") {
			if v == "*" {
				continue
			}
			m[i+1] = v
		}
		ph = append(ph, m)
	}
	return ph, nil
}
