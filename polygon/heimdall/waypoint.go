package heimdall

import (
	"fmt"
	"math/big"

	libcommon "github.com/ledgerwatch/erigon-lib/common"
)

// checkpoints and milestones are both hash hashAccumulators as defined
// here https://www.ethportal.net/concepts/hash-accumulators

type Waypoint interface {
	fmt.Stringer
	StartBlock() *big.Int
	EndBlock() *big.Int
	RootHash() libcommon.Hash
	Timestamp() uint64
	Length() int
	CmpRange(n uint64) int
}

type WaypointFields struct {
	Proposer   libcommon.Address `json:"proposer"`
	StartBlock *big.Int          `json:"start_block"`
	EndBlock   *big.Int          `json:"end_block"`
	RootHash   libcommon.Hash    `json:"root_hash"`
	ChainID    string            `json:"bor_chain_id"`
	Timestamp  uint64            `json:"timestamp"`
}

func (a *WaypointFields) Length() int {
	return int(new(big.Int).Sub(a.EndBlock, a.StartBlock).Int64() + 1)
}

func (a *WaypointFields) CmpRange(n uint64) int {
	num := new(big.Int).SetUint64(n)
	if num.Cmp(a.StartBlock) < 0 {
		return -1
	}
	if num.Cmp(a.StartBlock) > 0 {
		return 1
	}
	return 0
}

type Waypoints []Waypoint
