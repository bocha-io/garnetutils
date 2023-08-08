# Garnetutils

This CLI app will autogenerate the files required by the `Go` backend to interact with the `Garnet` indexer.

## Requirements

- Go >= 1.20

## Installation

```sh
go install github.com/bocha-io/garnetutils@latest
```

## Usage

```sh
garnetutils generate -i ./mud.config.ts -o /tmp/garnetgenerated/
```

Params:

- `-i`: mud config path
- `-o`: destination folder

### Using the .go files

The autogenerated files must be in a folder named `garnethelpers` inside your go project.

- The `getters` file has all the helpers needed to read the blockchain information using the `Garnet` lib.
- The `events` file has all the functions to create `MudEvents`, events used to predict the result of a transaction.
- The `types` file has the common types, you need to call `NewGameObject` in your go project with a reference to the `Garnet` db, so the `getters` file has all the dependencies set to read the game status.

## Build from source

```sh
git clone https://github.com/bocha-io/garnetutils
cd garnetutils
make build
```

The `garnetutils` binary will be located at `garnetutils/build`.

## TODOs:

- Add all the possible mud types. (Right now it only supports the types that we are using in our games).
- Add support for MUD modules.
- WIP: Add a solidity transpiler for the validation and prediction of the systems implemented by the user in MUD.

## Alpha command:

It generates helpers and predictions for your MUD contracts

There are some limitations:

- You can not use Conditionals in solidity, you must use If statements.
- You can not use init an array in the same line, you need to declare it and set position by position.
- Your functions must have unique names.
- Your structs must be declared inside the contract definition.
- The first character of every function that you want to optimistically predict must be uppercased.

### Usage

```sh
garnetutils alpha -i ../eternal-legends-garnet/contracts-builder/contracts -o /tmp/garnetgenerated/
```
