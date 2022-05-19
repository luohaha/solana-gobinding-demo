# solana-gobinding-demo
Demo for solana go-binding

# How

Solana doesn't official support Go SDK to interact with onchain solana program. So I want to find out another way to interact with solana program in go program.

Solana support Rust SDK, and Rust ffi(foreign function interface) allow calling rust code from c. We can add some hint to rust code:
```
#[no_mangle]
pub unsafe extern  "C" fn test_call_by_go(...) { ... }
```
The `extern "C"` makes this function adhere to the C calling convention, and `no_mangle` attribute turns off Rust's name mangling, so easy to link. Then we can add `crate-type` to cargo.toml to build cdylib. Now cargo build can generate corresponding cdylib.

```
[lib]
crate-type = ["cdylib"]
```

Then we can use [cbindgen](https://github.com/eqrion/cbindgen) tool to generate c header files.

Now we have c header file and cdylib. we can use cgo to use these in go languange:

```
//#cgo LDFLAGS: -L../client/target/debug/ -lgobinding_contract
//
//#include "go_binding_demo.h"
import "C"
```

Be aware of that we have to use libc type in `extern C` rust function's arguments. For example, if I want to pass a string to `extern C` rust function, I can use `*libc::c_char` instead of `String/Str`.

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