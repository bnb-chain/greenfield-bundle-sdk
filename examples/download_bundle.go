package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bnb-chain/greenfield-bundle-sdk/bundle"
	bundletypes "github.com/bnb-chain/greenfield-bundle-sdk/types"
	"github.com/bnb-chain/greenfield-go-sdk/client"
	"github.com/bnb-chain/greenfield-go-sdk/types"
)

// This is an example of how to parse the bundled object uploaded in the upload_bundle.go example.
// Please ensure the configurations in the common.go file are updated correctly before running this example.
func main() {
	// Prepare accounts and initialize Greenfield client
	account, err := types.NewAccountFromPrivateKey("test", privateKey)
	handleErr(err, "NewAccountFromPrivateKey")
	cli, err := client.New(chainId, rpcAddr, client.Option{DefaultAccount: account})
	handleErr(err, "new Greenfield client")
	ctx := context.Background()

	// Query bucket info
	bucketInfo, err := cli.HeadBucket(ctx, bucketName)
	handleErr(err, "HeadBucket")
	log.Println("bucket info:", bucketInfo.String())

	// Get bundle object from Greenfield
	bundledObject, info, err := cli.GetObject(ctx, bucketName, bundleObjectName, types.GetObjectOptions{})
	handleErr(err, "GetObject")
	log.Printf("get object %s successfully, size %d \n", info.ObjectName, info.Size)

	// Write bundle object into temp file
	bundleFile, err := os.CreateTemp("", bundletypes.TempBundleFilePrefix)
	handleErr(err, "CreateTemp")
	defer bundleFile.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(bundledObject)
	handleErr(err, "ReadFrom")
	_, err = bundleFile.Write(buf.Bytes())
	handleErr(err, "Write")

	// Extract objects from bundled object
	bundle, err := bundle.NewBundleFromFile(bundleFile.Name())
	handleErr(err, "NewBundleFromFile")
	defer bundle.Close()
	obj1, size, err := bundle.GetObject("object1")
	if err != nil || obj1 == nil || size != singleObjectSize {
		handleErr(fmt.Errorf("parse object1 in bundled object failed: %v", err), "GetObject")
	}
	obj2, size, err := bundle.GetObject("object2")
	if err != nil || obj2 == nil || size != singleObjectSize {
		handleErr(fmt.Errorf("parse object2 in bundled object failed: %v", err), "GetObject")
	}

	log.Printf("Congratulations, everything going well!")
}
