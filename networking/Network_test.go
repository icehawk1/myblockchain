package networking

import (
	"awesomeProject/blockchain"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	PeerList = append(PeerList, CreatePeer("http://localhost:8000"))
	PeerList = append(PeerList, CreatePeer("http://localhost:63975"))

	os.Exit(m.Run())
}

func TestCreatePeer(t *testing.T) {
	validPeer1 := CreatePeer("http://heise.de/")
	assert.True(t, validPeer1.Validate())

	validPeer2 := CreatePeer("https://heise:7654.de")
	assert.True(t, validPeer2.Validate())

	invalidPeer := CreatePeer("ftp://heise.de/")
	assert.False(t, invalidPeer.Validate())
}

func TestBroadcastTransaction(t *testing.T) {
	tx := blockchain.Transaction{Message: "Tx nummer 1"}
	statusCodes := BroadcastTransaction(tx)
	assert.Equal(t, []int{200,200}, statusCodes)
}

func TestBroadcastBlock(t *testing.T) {
	block := blockchain.Block{Prev: "", Hash: "00009DB3B9F2ACD62E2AE8725EB5AF49438D2AF0B79F4FC196BA2F5BDB1C1F36"}
	statusCodes := BroadcastBlock(block)
	assert.Equal(t, []int{400, 400}, statusCodes)
}
