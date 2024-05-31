package hyperledger

import (
	"fmt"
	"testing"
)

func Test_ConnectionTest_Success(t *testing.T) {

	clients := NewClientMap(
		"test-network",
		"/home/fusion/Projects/blockchain/hyperledger-fabric-kubernetes/s5-connecting-everything/back-end/hyperledger/config.yaml",
	)

	_, err := clients.AddClient(
		"Admin",
		"enspy",
		"mainchannel",
	)
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := clients.Query("enspy", "mainchannel", "organizations", "index", [][]byte{
		[]byte(""),
		[]byte(""),
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(string(res))
}
