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
		nodes = append(nodes, model.Node{Left: nil, Right: nil, Hash: hash})
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
			newLevel = append(newLevel, nodes[i])
		}
	}
	return generateTree(newLevel)
}

func GenerateMerkleProof(transaction string, root *model.Node) ([]string, bool) {
	var proof []string
	return proof, findProof(root, utils.GenerateHash(transaction), &proof)
}

func findProof(node *model.Node, hash string, proof *[]string) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil && node.Hash == hash {
		return true
	}

	if node.Left != nil && findProof(node.Left, hash, proof) {
		if node.Right != nil {
			*proof = append(*proof, node.Right.Hash)
		}
		return true
	}

	if node.Right != nil && findProof(node.Right, hash, proof) {
		*proof = append(*proof, node.Left.Hash)
		return true
	}
	return false
}

func CheckIfTransactionExist(transaction string, root *model.Node) bool {
	proof, isExist := GenerateMerkleProof(transaction, root)
	log.Println("Proof Nodes : ", proof)
	var concatenated []byte

	for _, hash := range proof {
		decodedHash := utils.GenerateHash(hash)
		concatenated = append(concatenated, decodedHash...)
	}
	concatHexString := utils.GenerateHash(string(concatenated))
	log.Println("Concatenated Hash : ", concatHexString)
	log.Println("Root Hash : ", root.Hash)
	log.Println("Is exist : ", isExist)
	return concatHexString == root.Hash
}
