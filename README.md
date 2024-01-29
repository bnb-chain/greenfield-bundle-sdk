# Greenfield Bundle SDK

Storing small files in Greenfield is inefficient due to the metadata stored on the blockchain being larger than the files
themselves. This leads to higher costs for users. Additionally, Greenfield Blockchain has a capacity limit for processing
files simultaneously.

To address this issue, we have proposed [BEP-323: Bundle Format For Greenfield](https://github.com/bnb-chain/BEPs/pull/323).
This repository contains the Golang version of the bundle format, which guides users on aggregating objects into a bundle
and parsing a bundled object into separate objects.

## Bundle Specification

### Bundle Format
The bundle format specifies the structure and organization of the bundle that users create when packing files.
This format is designed to pack flat files; hierarchical directory structures, or folders, are not supported.

When dealing with a folder, users can simplify its structure by turning it into a series of individual files.
As part of this process, it renames each file to include the folder path. For example, a file originally named
`file.txt` inside the nested folders `dirA` and `dirB` would be renamed to `dirA/dirB/file.txt`.
This approach allows us to maintain the organization of the folder while conforming to the requirement for flat files in the bundle.

There are still constraints for the bundle format. The file names of the files in the bundle should be unique so that
they can be indexed by the file name.

The bundle format is structured into several key components as follows:

*   Data: This portion represents the actual content and is comprised of all the files in bytes.
*   Metadata: This section contains information about the files within the bundle. It facilitates the ability to access files randomly, which means you can jump directly to any file within the bundle without going through all the files.
*   Meta Size: This specifies the size of the bundle's metadata, allowing the construction of the bundle structure without the need to read the entire bundle.
*   Version: This indicates the version number of the bundle protocol being used.

<img width="645" alt="bundle_struct" src="https://github.com/bnb-chain/BEPs/assets/5030187/0bd75d18-9d33-4469-beb4-2b0623a1d48d">

The Meta structure is designed to include essential attributes for each file, outlined as follows:

*   Object Name: This is the name of the file within the bundle.
*   Offset: This attribute marks the starting point of the file's data within the bundle.
*   Size: This details the total length, in bytes, of the file.
*   Hash Algo: This specifies the algorithm used for the file's hash calculation.
*   Hash: This is the cryptographic hash result of the file's content. It serves as a tool for verifying the file's integrity.
*   Content Type: This denotes the MIME type of the file, describing the file's nature and format.
*   Tags: This is a map that holds various additional properties of the file like `owner` .

### Encoding
<img width="645" alt="bundle_encoding" src="https://github.com/bnb-chain/BEPs/assets/5030187/54604c12-b2d7-4984-9e00-d848fe7b4686">

The bundle's encoding format is structured as follows:

*   Data: This consists of the actual file contents, represented as a sequence of bytes.
*   Metadata: Encoded in bytes, this section utilizes Protocol Buffers (protobuf) for serialization.
*   Meta Size: Also an unsigned 64-bit integer, represented using 8 bytes, indicating the size of the metadata section.
*   Version: Serialized as an unsigned 64-bit integer, occupying 8 bytes.

The Meta structure will be serialized with protobuf:

```java
enum HashAlgo {
  Sha256 = 0;
}

message Meta {
  repeated FileMeta meta = 1;
}

message FileMeta {
  string object_name = 1;
  uint64 offset = 2;
  uint64 size = 3;
  HashAlgo status = 4;
  bytes hash = 5;
  string content_type = 6;
  map<string, string> tags = 7;
}
```

## Quick Start

Here is the guide for how to aggregate batch objects as a bundle, and how to parse a bundled object. As for how to
interact with Greenfield, you should refer to [Greenfield GO SDK](https://github.com/bnb-chain/greenfield-go-sdk).

### Aggregate various objects as bundle
Follow the steps below to aggregate multiple objects into a single bundle.

1. Use the `NewBundle` function to create an empty bundle.
2. Use the bundle's `AppendObject` method to add objects to the bundle individually.
3. Use the bundle's `FinalizeBundle` method to seal the bundle, preventing any further objects from being added.
4. To release resources after use, utilize the `Close` method of the bundle.

### Extract objects from bundled object
Follow the steps below to extract various objects from a bundle.

1. Open the bundled object as a bundle instance using `NewBundleFromFile`.
2. Retrieve all the objects' meta within the bundle using the bundle's `GetBundleObjectsMeta` method.
3. Access various objects one by one using the bundle's `GetObject` method.
4. To release resources after use, utilize the `Close` method of the bundle.


## Command line tool

The command line tool `bundler` supports two subcommands: `bundle` and `verify`.

### Build

```shell
make build
```

### Usage

To bundle a directory into a bundle file:

```shell
./build/bundler bundle -input <input_directory> -output <output_bundle_file>
```

To verify a bundle file:

```shell
./build/bundler verify -input <input_bundle_file>
```

If the required arguments are not provided, the tool will output usage information for the respective subcommand.

### Bundle

The `bundle` subcommand takes an input directory and an output bundle file as arguments. It bundles the directory into a bundle file.

### Verify

The `verify` subcommand takes an input bundle file as an argument. It verifies the bundle file whether the bundle is valid.


## Contribution
Thank you for considering helping with the source code! We appreciate contributions from anyone on the internet, no
matter how small the fix may be.

If you would like to contribute to Greenfield, please follow these steps: fork the project, make your changes, commit them,
and send a pull request to the maintainers for review and merge into the main codebase. However, if you plan on submitting
more complex changes, we recommend checking with the core developers first via GitHub issues (we will soon have a Discord channel)
to ensure that your changes align with the project's general philosophy. This can also help reduce the workload of both
parties and streamline the review and merge process.

## Licence

The greenfield-bundle-sdk is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.