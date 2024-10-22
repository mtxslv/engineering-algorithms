package main

import "fmt"

type pirate struct {
	name   string
	bounty float32
}

func main() {
	var mugiwara_bounty_after_water_7 = []pirate{
		pirate{"Monkey D. Luffy", 300000000},
		pirate{"Roronoa Zoro", 120000000},
		pirate{"Nico Robin", 80000000},
		pirate{"Sanji", 77000000},
		pirate{"Franky", 44000000},
		pirate{"Usopp", 30000000},
		pirate{"Nami", 16000000},
		pirate{"Tony Tony Chopper", 50},
	}
	fmt.Println("Strawhats bounties:")
	for _, a_mugiwara := range mugiwara_bounty_after_water_7 {
		fmt.Println("Name: ", a_mugiwara.name, " | Bounty: ", a_mugiwara.bounty)
	}
}
