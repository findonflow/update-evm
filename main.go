package main

import (
	"fmt"

	"github.com/bjartek/overflow/v2"
)

func main() {
	o := overflow.Overflow()

	fmt.Println("Look for a TODO comment at the start of the following output")
	fmt.Println("======")
	o.Script(`access(all) fun main(address:Address) : String {

       let account=getAccount(address)!
       if let dc=account.contracts.get(name:"EVM") {
          return String.fromUTF8(dc.code) ?? ""
       }
       return ""

    }`, overflow.WithArg("address", "account"))
}
