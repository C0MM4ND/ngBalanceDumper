package main

import (
	"bytes"
	"fmt"
	"github.com/c0mm4nd/ngsheetdumper/common"
	"github.com/dgraph-io/badger/v2"
	"github.com/mr-tron/base58"
	"github.com/ngchain/ngcore/ngtypes"
	"math/big"
)

func main() {
	opt := badger.DefaultOptions("./ngdb")
	opt.Truncate = true
	db, err := badger.Open(opt)
	if err != nil {
		panic(err)
	}

	//txn := db.NewTransaction(false)
	//// txs.Get()
	//hash, err := ngblocks.GetLatestHash(txn)
	//if err != nil {
	//	panic(err)
	//}
	//
	//block, err := ngblocks.GetBlockByHash(txn, hash)
	//if err != nil {
	//	panic(err)
	//}
	//print(block.Height)

	//state := ngstate.State{db: db}
	balances := make(map[string]string)
	err = db.View(func(txn *badger.Txn) error {
		ki := txn.NewIterator(badger.DefaultIteratorOptions)
		defer ki.Close()
		addrTobBalancePrefix := []byte("ab:")
		for ki.Seek(addrTobBalancePrefix);ki.ValidForPrefix(addrTobBalancePrefix);ki.Next() {
			kv := ki.Item()
			v, err := kv.ValueCopy(nil)
			if err != nil {
				panic(err)
			}
			num := new(big.Int).SetBytes(v)
			k := kv.Key()
			//fmt.Println(len(kv.Key()))
			if len(k) == 38 {
				k = k[3:]
			}
			//if len(kv.Key()) == 51 {
			//	fmt.Println(len(k))
			//	fmt.Println(hex.EncodeToString(utils.Bytes2PublicKey(kv.Key()[2:]).Serialize()))
			//}

			//if num.String()[0:4] == "1673" {
			//	fmt.Println("1673: "+ hex.EncodeToString(kv.Key()))
			//	base , _ := base58.FastBase58Decoding("Lo4odLdP1MWpmvq7yKwJ5cLEeRqsNf8d2i5QVwWRNUYV92Ua")
			//	//base = append(addrTobBalancePrefix, base...)
			//	str := base58.FastBase58Encoding(base)
			//	base = []byte(str)
			//	fmt.Println("loop", i, "->", len(base))
			//	i++
			//	base = append(addrTobBalancePrefix, base...)
			//	if base != key
			//}
			if len(k) == 51 {
				k = k[3:]
				k, _ =base58.FastBase58Decoding(string(k))
				//k = k[3:]
				//k, _ =base58.FastBase58Decoding(string(k))
			}



			if num.String()[0:4] == "1673" {
				base , _ := base58.FastBase58Decoding("Lo4odLdP1MWpmvq7yKwJ5cLEeRqsNf8d2i5QVwWRNUYV92Ua")
				if !bytes.Equal(base, k){
					panic(len(kv.Key()))
				}
			}

			if len(k) == 50 {// this means we need to pad
				k = k[3:]
				k, _ = base58.FastBase58Decoding(string(k))
				//k = k[3:]
				//k, _ =base58.FastBase58Decoding(string(k))
			}


			if num.String()[0:10] == "6335115345" {
				base , _ := base58.FastBase58Decoding("BeaDaq2xVa697RGBu5hZ48EYMdgbNVd1kr9DgZoTFsymY4Y")
				if !bytes.Equal(base, k){
					panic(len(kv.Key()))
				}
			}

			//
			if len(k)!=35{
				fmt.Println(num.String())
				panic(len(kv.Key()))
			}

			if balances[ngtypes.Address(k).BS58()] != "" {
				fmt.Println(ngtypes.Address(k).BS58(), balances[ngtypes.Address(k).BS58()], ",", num.String())
				continue
			}

			balances[ngtypes.Address(k).BS58()] = num.String()
				//new(big.Float).Quo(
				//new(big.Float).SetInt(num),
				//new(big.Float).SetInt(ngtypes.NG),
				//).Float64()
			//fmt.Println(ngtypes.Address(k), num)
		}

		//base , _ := base58.FastBase58Decoding("BeaDaq2xVa697RGBu5hZ48EYMdgbNVd1kr9DgZoTFsymY4Y")
		//base = append(base, []byte{0}...)
		//if base58.FastBase58Encoding(base) !="BeaDaq2xVa697RGBu5hZ48EYMdgbNVd1kr9DgZoTFsymY4Y" {
		//	panic(base58.FastBase58Encoding(base))
		//}

		return nil
	})

	common.SaveMap(balances, "testnet_v1_ending_balances.json")
}
