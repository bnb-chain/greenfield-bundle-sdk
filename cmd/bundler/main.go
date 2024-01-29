package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"

	"github.com/bnb-chain/greenfield-bundle-sdk/bundle"
	"github.com/bnb-chain/greenfield-bundle-sdk/types"
)

func visit(root string, b *bundle.Bundle) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !f.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			relativePath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			ext := filepath.Ext(path)
			contentType := mime.TypeByExtension(ext)

			content, err := io.ReadAll(file)
			if err != nil {
				return err
			}

			hash := sha256.Sum256(content)
			options := &types.AppendObjectOptions{
				ContentType: contentType,
				HashAlgo:    types.HashAlgo_SHA256,
				Hash:        hash[:],
			}

			_, err = file.Seek(0, io.SeekStart)
			if err != nil {
				return err
			}

			_, err = b.AppendObject(relativePath, file, options)
			if err != nil {
				return err
			}
			fmt.Printf("%s %s\n", relativePath, hex.EncodeToString(hash[:]))
		}
		return nil
	}
}

func bundleCmd(inputDirectory, outputBundleFile string) {
	bundleObject, _, err := bundleDirectory(inputDirectory)
	if err != nil {
		fmt.Printf("Failed to pack directory into bundle: %v\n", err)
		os.Exit(1)
	}

	err = saveBundleToFile(bundleObject, outputBundleFile)
	if err != nil {
		fmt.Printf("Failed to save bundle to file: %v\n", err)
		os.Exit(1)
	}
}

func verifyCmd(inputBundleFile string) {
	_, err := bundle.NewBundleFromFile(inputBundleFile)
	if err != nil {
		fmt.Printf("invalid bundle, err=%s\n", err.Error())
		return
	}
	println("bundle is valid")
}

func bundleDirectory(dir string) (io.ReadSeekCloser, int64, error) {
	b, err := bundle.NewBundle()
	if err != nil {
		return nil, 0, err
	}

	err = filepath.Walk(dir, visit(dir, b))
	if err != nil {
		return nil, 0, err
	}

	return b.FinalizeBundle()
}

func saveBundleToFile(bundle io.ReadSeekCloser, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, bundle)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	bundleCommand := flag.NewFlagSet("bundle", flag.ExitOnError)
	verifyCommand := flag.NewFlagSet("verify", flag.ExitOnError)

	bundleInputDir := bundleCommand.String("input", "", "Input directory to bundle")
	bundleOutputFile := bundleCommand.String("output", "", "Output bundle file")

	verifyInputFile := verifyCommand.String("input", "", "Input bundle file to verify")

	bundleCommand.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  bundle -input <input_directory> -output <output_bundle_file>\n")
		bundleCommand.PrintDefaults()
	}

	verifyCommand.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  verify -input <input_bundle_file>\n")
		verifyCommand.PrintDefaults()
	}

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: Invalid subcommand provided. The program expects either 'bundle' or 'verify' as a subcommand.\n")
		bundleCommand.Usage()
		verifyCommand.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "bundle":
		bundleCommand.Parse(os.Args[2:])
		if *bundleInputDir == "" || *bundleOutputFile == "" {
			bundleCommand.Usage()
			os.Exit(1)
		}
		bundleCmd(*bundleInputDir, *bundleOutputFile)
	case "verify":
		verifyCommand.Parse(os.Args[2:])
		if *verifyInputFile == "" {
			verifyCommand.Usage()
			os.Exit(1)
		}
		verifyCmd(*verifyInputFile)
	default:
		fmt.Fprintf(os.Stderr, "Error: Invalid subcommand provided. The program expects either 'bundle' or 'verify' as a subcommand.\n")
		bundleCommand.Usage()
		verifyCommand.Usage()
		os.Exit(1)
	}
}
