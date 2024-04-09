package main

import (
	"log"
)

// The config information is consistent with the testnet of greenfield
// You need to set the privateKey, bucketName, bundleObjectName to make the basic examples work well
const (
	rpcAddr          = "https://gnfd-testnet-fullnode-tendermint-us.bnbchain.org:443"
	chainId          = "greenfield_5600-1"
	singleObjectSize = 1000

	// Please update your configuration here before running the examples.
	privateKey       = ""
	bucketName       = "bundle-sdk-example"
	bundleObjectName = "bundle-object"
)

func handleErr(err error, funcName string) {
	if err != nil {
		log.Fatalln("fail to " + funcName + ": " + err.Error())
	}
}
