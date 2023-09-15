// Code generated by suave/gen. DO NOT EDIT.
// Hash: 3560bf1169d2bb413a9a0303c4043027fe6111a2f3cbfa25bf264612b2694ee4
package types

import "github.com/ethereum/go-ethereum/common"

type BidId [16]byte

// Structs

type Bid struct {
	Id                  BidId
	DecryptionCondition uint64
	AllowedPeekers      []common.Address
	Version             string
}

type BuildBlockArgs struct {
	Slot           uint64
	ProposerPubkey []byte
	Parent         common.Hash
	Timestamp      uint64
	FeeRecipient   common.Address
	GasLimit       uint64
	Random         common.Hash
	Withdrawals    []*Withdrawal
}

type Withdrawal struct {
	Index     uint64
	Validator uint64
	Address   common.Address
	Amount    uint64
}