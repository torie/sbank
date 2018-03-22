package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/torie/sbank"
)

func main() {
	uid := flag.String("user-id", "", "customer id (social security number)")
	cid := flag.String("client-id", "", "clientId obtained from Sbanken API Beta / utviklerportalen")
	secret := flag.String("client-secret", "", "password obtained from Sbanken API Beta / utviklerportalen")
	flag.Parse()

	if *uid == "" || *cid == "" || *secret == "" {
		flag.Usage()
		os.Exit(-1)
	}

	bank := sbank.New(*cid, *secret)

	cresp, err := bank.Customer(*uid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nCustomer:")
	fmt.Println(cresp.Customer)

	aresp, err := bank.Accounts(*uid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nAccounts:")
	for _, account := range aresp.Accounts {
		fmt.Println(account, "\n------")
	}

	if len(aresp.Accounts) == 0 {
		return
	}

	tresp, err := bank.Transactions(*uid, aresp.Accounts[0].AccountNumber, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nTransactions:")
	for _, transaction := range tresp.Transactions {
		fmt.Println(transaction, "\n------")
	}
}
