package main

import (
	"fmt"
	"github.com/etherscan-api"
)

//failed, no value
func main (){
	accountAddr := "0x50352B904445576242444bc1924e93e61090738C"
	ethClient := etherscan.New(etherscan.Ropsten, "EMYG8PZAFAU966YXHV2KIATYMYP9Q9VFPR")
	balance, err := ethClient.AccountBalance(accountAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("balance: ", balance)
}
