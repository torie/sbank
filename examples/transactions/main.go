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
	aid := flag.String("account-id", "", "The account to list")
	flag.Parse()

	if *uid == "" || *cid == "" || *secret == "" || *aid == "" {
		flag.Usage()
		os.Exit(-1)
	}

	bank := sbank.New(*uid, *cid, *secret)

	res, err := bank.Transactions(*aid, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nTransactions:")
	for _, item := range res.Items {
		fmt.Println(item, "\n------")
	}
}
