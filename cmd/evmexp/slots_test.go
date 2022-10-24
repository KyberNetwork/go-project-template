package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	usdcBalance   = 9
	usdcAllowance = 10
)

func computeAddrIntHash(addr common.Address, slot int64) common.Hash {
	addr32 := common.BytesToHash(addr[:])
	slot32 := common.LeftPadBytes(big.NewInt(slot).Bytes(), 32)
	return crypto.Keccak256Hash(addr32[:], slot32)
}

func computeAddrHashHash(addr common.Address, hash common.Hash) common.Hash {
	addr32 := common.BytesToHash(addr[:])
	return crypto.Keccak256Hash(addr32[:], hash[:])
}

func TestEncodeSlot(t *testing.T) {
	addr := myWallet
	t.Log(computeAddrIntHash(addr, usdcBalance))   // balance slot
	hh1 := computeAddrIntHash(addr, usdcAllowance) // allowances slot
	t.Log(computeAddrHashHash(univ2routerAddress, hh1))
}

func TestMap(t *testing.T) {
	x := map[string][]byte{}
	v := []byte{0, 1}
	x["key"] = v
	v[0] = 10
	t.Log(x["key"])
}
