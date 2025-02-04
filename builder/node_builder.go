package builder

import (
	"github.com/darshandeepak-07/merkle-tree/model"
	"github.com/darshandeepak-07/merkle-tree/utils"
)

func CreateNode(left, right *model.Node, data string) *model.Node {
	node := &model.Node{}

	if left == nil && right == nil {
		node.Hash = utils.GenerateHash(data)
	} else {
		mergedHash := append([]byte(left.Hash), []byte(right.Hash)...)
		node.Hash = utils.GenerateHash(string(mergedHash))
	}

	node.Left = left
	node.Right = right

	return node
}
