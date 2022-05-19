# solana-gobinding-demo
Demo for solana go-binding

# How



# Usage

### 0. setup environment
Make sure you already install solana, rust, go, cargo, cbindgen

### 1. build solana program and deploy
start local solana testnet
```
solana config set --url http://127.0.0.1:8899
solana-test-validator
```

build program and deploy
```
cd program
cargo build-bpf
solana program deploy target/deploy/helloworld.so
```

### 2. generate cdylib and c header from client

```
Cd client
cargo build
cbindgen --config cbindgen.toml --crate gobinding-contract --output ../go/go_binding_demo.h
```

### 3. run go demo

```
Cd go
go run go_binding_demo.go ../program/target/deploy/helloworld-keypair.json
```

expected output :
```
Test Solana Go Hello World
Connected to remote solana node running version (x).
(x) lamports are required for this transaction.
(x) greetings have been sent.
```

# Reference
Some rust code are copied from `https://github.com/ezekiiel/simple-solana-program`.