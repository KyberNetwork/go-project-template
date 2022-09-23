package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	t.Log(computeAddrIntHash(addr, 101)) // balance slot
	hh1 := computeAddrIntHash(addr, 102) // allowances slot
	t.Log(computeAddrHashHash(kyberswapRouterAddress, hh1))
}
