package mecabs

import (
	"strings"

	"github.com/shogo82148/go-mecab"
)

type MeCabS struct {
	mecab mecab.MeCab
}

func New(args map[string]string) (MeCabS, error) {
	tagger, err := mecab.New(args)
	if err != nil {
		return MeCabS{}, err
	}
	return MeCabS{mecab: tagger}, nil
}

func (m MeCabS) Destroy() {
	m.mecab.Destroy()
}

func (m MeCabS) Parse(s string) ([][10]string, error) {
	node, err := m.mecab.ParseToNode(s)
	if err != nil {
		return [][10]string{}, err
	}
	ans := make([][10]string, countNodes(node)-2)
	for i := 0; ; node = node.Next() {
		switch node.Stat() {
		case mecab.BOSNode:
		case mecab.EOSNode:
			return ans, nil
		default:
			nodeArr := [10]string{}
			nodeArr[0] = node.Surface()
			feature := strings.Split(node.Feature(), ",")
			for i, v := range feature {
				if v == "*" {
					nodeArr[i+1] = ""
				} else {
					nodeArr[i+1] = feature[i]
				}
			}
			ans[i] = nodeArr
			i++
		}
	}
}

func countNodes(node mecab.Node) (count int) {
	for ; node.Stat() != mecab.EOSNode; node = node.Next() {
		count++
	}
	return count + 1
}
