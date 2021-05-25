package main

import (
	"fmt"
	"github.com/c0mm4nd/ngsheetdumper/common"
	"math/big"
)

func main() {
	genesisMap := common.ReadMap("genesis.json")
	testnetV1EndMap := common.ReadMap("testnet_v1_ending_balances.json")
	deltaMap := make(map[string]string)
	v2Map := make(map[string]string)
	for addr, strBal := range testnetV1EndMap {
		var delta, endBal,genBal *big.Int

		var ok bool
		endBal, ok = new(big.Int).SetString(strBal, 10)
		if !ok {
			panic("")
		}


		strGenBal, exists := genesisMap[addr]
		if !exists {
			genBal = big.NewInt(0)
			delta = endBal
		}else {
			genBal, ok = new(big.Int).SetString(strGenBal, 10)
			if !ok {
				fmt.Println(strGenBal)
				panic("")
			}

			delta = new(big.Int).Sub(endBal, genBal)
		}



		deltaMap[addr] = delta.String()
		delta.Div(delta, big.NewInt(10))
		v2Map[addr] = genBal.Add(genBal, delta).String()
	}

	common.SaveMap(v2Map, "testnet_v2_start_balances.json")
	common.SaveMap(deltaMap, "testnet_v1_delta_balances.json")
}
