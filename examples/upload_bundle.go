package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/bnb-chain/greenfield-bundle-sdk/bundle"
	"github.com/bnb-chain/greenfield-go-sdk/client"
	"github.com/bnb-chain/greenfield-go-sdk/types"
)

// TThis is an example of how to assemble objects into a bundle and upload them to Greenfield.
// Please ensure the configurations in the common.go file are updated correctly before running this example.
func main() {
	// Prepare accounts and initialize Greenfield client
	account, err := types.NewAccountFromPrivateKey("test", privateKey)
	handleErr(err, "NewAccountFromPrivateKey")
	cli, err := client.New(chainId, rpcAddr, client.Option{DefaultAccount: account})
	handleErr(err, "new Greenfield client")
	ctx := context.Background()

	// get storage providers list and choose a primary SP
	spLists, err := cli.ListStorageProviders(ctx, true)
	handleErr(err, "ListStorageProviders")
	primarySP := spLists[0].GetOperatorAddress()

	// create bucket and query bucket info
	bucketInfo, err := cli.HeadBucket(ctx, bucketName)
	if err != nil {
		_, err = cli.CreateBucket(ctx, bucketName, primarySP, types.CreateBucketOptions{})
		handleErr(err, "CreateBucket")
		log.Printf("create bucket %s on SP: %s successfully \n", bucketName, spLists[0].Endpoint)
		bucketInfo, err = cli.HeadBucket(ctx, bucketName)
		handleErr(err, "HeadBucket")
	}

	log.Println("bucket info:", bucketInfo.String())

	// Prepare two objects for aggregating into a bundle
	var buffer1 bytes.Buffer
	line := `0123456789`
	for i := 0; i < singleObjectSize/10; i++ {
		buffer1.WriteString(fmt.Sprintf("%s", line))
	}
	var buffer2 bytes.Buffer
	line = `9876543210`
	for i := 0; i < singleObjectSize/10; i++ {
		buffer2.WriteString(fmt.Sprintf("%s", line))
	}

	// Assemble above two objects into a bundle object
	bundle, err := bundle.NewBundle()
	handleErr(err, "NewBundle")
	defer bundle.Close()
	_, err = bundle.AppendObject("object1", bytes.NewReader(buffer1.Bytes()), nil)
	handleErr(err, "AppendObject")
	_, err = bundle.AppendObject("object2", bytes.NewReader(buffer2.Bytes()), nil)
	handleErr(err, "AppendObject")
	bundledObject, totalSize, err := bundle.FinalizeBundle()
	handleErr(err, "FinalizeBundle")

	// create and put bundle object onto Greenfield
	_, err = cli.CreateObject(ctx, bucketName, bundleObjectName, bundledObject, types.CreateObjectOptions{})
	handleErr(err, "CreateObject")
	err = cli.PutObject(ctx, bucketName, bundleObjectName, totalSize, bundledObject, types.PutObjectOptions{})
	handleErr(err, "PutObject")

	log.Printf("Congratulations, bundle object: %s has been uploaded to SP\n", bundleObjectName)
}
