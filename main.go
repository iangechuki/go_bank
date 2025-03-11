package main

import (
	"fmt"

	"github.com/iangechuki/go_bank/util"
)

func main() {
	owner := util.RandomOwner()
	money := util.RandomMoney()
	str := fmt.Sprintf("Owner: %s Money: %d", owner, money)
	fmt.Println(str)
}
