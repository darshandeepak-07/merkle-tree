package builder

import (
	"log"

	"github.com/darshandeepak-07/merkle-tree/model"
	"github.com/darshandeepak-07/merkle-tree/utils"
)

func BuildMerkleTree(data []string) *model.Node {
	var nodes []model.Node

	for _, transaction := range data {
		hash := utils.GenerateHash(transaction)
		nodes = append(nodes, *&model.Node{Left: nil, Right: nil, Hash: hash})
	}
	log.Println("Node length = ", len(nodes))
	return generateTree(nodes)
}

func generateTree(nodes []model.Node) *model.Node {

	if len(nodes) == 1 {
		return &nodes[0]
	}

	var newLevel []model.Node
	for i := 0; i < len(nodes); i += 2 {
		if i+1 < len(nodes) {
			parentHash := utils.GenerateHash(nodes[i].Hash + nodes[i+1].Hash)
			newLevel = append(newLevel, model.Node{Left: &nodes[i], Right: &nodes[i+1], Hash: parentHash})
		} else {
			parentHash := utils.GenerateHash(nodes[i].Hash + nodes[i].Hash)
			newLevel = append(newLevel, model.Node{Left: &nodes[i], Right: &nodes[i], Hash: parentHash})
		}
	}
	return generateTree(newLevel)
}
