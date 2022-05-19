package main

//#cgo LDFLAGS: -L../client/target/debug/ -lgobinding_contract
//
//#include "go_binding_demo.h"
import "C"
import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage : go run go_bingding_demo.go [program keypair.json]")
		return
	}
	fmt.Println("Test Solana Go Hello World")
	cstr := C.CString(os.Args[1])
	C.test_call_by_go(cstr)
}
