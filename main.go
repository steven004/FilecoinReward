package main

import (
	"fmt"
	"math/big"
)

// /////
// Devnet settings

const TotalFilecoin = 2000000000
const MiningRewardTotal = 1400000000

const InitialRewardStr = "153856861913558700202"

var InitialReward *big.Int
var miningRewardTotal *big.Int

const FilecoinPrecision = 1000000000000000000

// six years
// Blocks
const HalvingPeriodBlocks = 6 * 365 * 24 * 60 * 2

// Blocks - Does not use any more
const AdjustmentPeriod = 7 * 24 * 60 * 2

// TODO: Move other important consts here
const MaxRound = 428215690

// Init is to init the initial reward
func Init() {
	InitialReward = new(big.Int)

	var ok bool
	InitialReward, ok = InitialReward.
		SetString(InitialRewardStr, 10)
	if !ok {
		panic("could not parse InitialRewardStr")
	}
}

// MiningReward returns correct mining reward
//   coffer is amount of FIL in NetworkAddress
func MiningReward(remainingReward *big.Int) *big.Int {
	ci := big.NewInt(0).Set(remainingReward)
	res := ci.Mul(ci, InitialReward)
	res = res.Div(res, miningRewardTotal)
	return res
}

// FromFil transfer FIL to attoFIL
func FromFil(i uint64) *big.Int {
	res := new(big.Int)
	res.SetUint64(i)
	return res.Mul(res, big.NewInt(FilecoinPrecision))
}

func main() {
	// the mining total in attoFIL
	miningRewardTotal = FromFil(MiningRewardTotal)

	// To initiate initalReward to attoFIL
	Init()
	remainingReward := new(big.Int)
	remainingReward = remainingReward.Set(miningRewardTotal)
	var miningReward *big.Int
	days := 1

	for i := 0; i < MaxRound; i++ {
		miningReward = MiningReward(remainingReward)
		remainingReward = remainingReward.Sub(remainingReward, miningReward)
		if i%2880 == 0 {
			fmt.Println(days, i+1, miningReward, remainingReward)
			days++
		}
		if miningReward.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("-------- in this round to 0 --------\n", days, i+1, miningReward, remainingReward)
			break
		}
	}
	fmt.Println("\nlast years and days:", days/365, days%365)
}
