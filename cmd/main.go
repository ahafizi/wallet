package main

import (
	"log"

	"github.com/ahafizi/wallet/pkg/wallet"
)

func main() {
	s := &wallet.Service{}

	// _, err := s.RegisterAccount("+992880806776")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// err = s.ExportToFile("../data/accounts.txt")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	err := s.ImportFromFile("data/export.txt")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(s.FindAccountByID(1))
}
