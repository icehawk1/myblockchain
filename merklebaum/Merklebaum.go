package merklebaum

import (
	"awesomeProject/blockchain"
	"encoding/json"
)

type Merklebaum struct {
	Hash  string
	left  *Merklebaum
	right *Merklebaum
	elem  *blockchain.Hashable
}

func (self Merklebaum) ComputeHash() string {
	if self.elem != nil {
		return (*self.elem).ComputeHash()
	} else {

		// Prepend 0 to prevent second preimage attack
		input := "0"
		if self.left != nil {
			input += self.left.Hash
		}

		if self.right != nil {
			input += self.right.Hash
		}

		return blockchain.ComputeSha256Hex(input)
	}
}

func (self Merklebaum) ComputeHashByte() []byte {
	panic("implement me")
}

func CreateMerklebaum(content []blockchain.Hashable) Merklebaum {
	if len(content) == 0 {
		return Merklebaum{}
	}

	var leafs = make([]blockchain.Hashable, 0, len(content))
	for i := 0; i < len(content); i++ {
		leafs = append(leafs, Merklebaum{Hash: content[i].ComputeHash(), elem: &content[i]})
	}

	var bäume = make([]*Merklebaum, 0, len(leafs))
	for len(bäume) > 1 {
		bäume = createMerkleLevel(bäume)
	}

	return *bäume[0]
}
func createMerkleLevel(bäume []*Merklebaum) []*Merklebaum {

	var result = make([]*Merklebaum, 0, len(bäume)/2+1)

	// In case of odd number of trees, skip last tree for later
	for i := 0; i+1 < len(bäume); i += 2 {
		neuerbaum := Merklebaum{left: bäume[i], right: bäume[i+1]}
		neuerbaum.Hash = neuerbaum.ComputeHash()
		result = append(result, &neuerbaum)
	}

	if len(bäume)%2 == 1 {
		result = append(result, &Merklebaum{left: bäume[len(bäume)-1]})
	}
	return result
}

func (self Merklebaum) IsValid() bool {
	return false
}

func (self Merklebaum) IsLeaf() bool {
	return self.left == nil && self.right == nil
}

func (self Merklebaum) HasNode(path []string) bool {
	return false
}

func (self Merklebaum) GetElements() []blockchain.Hashable {
	return []blockchain.Hashable{}
}

func (self Merklebaum) CreateSpvProof(leaf blockchain.Hashable) (proof []string, ok bool) {
	return []string{}, true
}

func (self Merklebaum) Contains(leaf blockchain.Hashable) bool {
	return false
}

func (self *Merklebaum) MarshalJSON() ([]byte, error) {
	return []byte(""), nil
}

func (self *Merklebaum) UnmarshalJSON(receivedData []byte) error {
	return json.Unmarshal(receivedData, &self)
}
