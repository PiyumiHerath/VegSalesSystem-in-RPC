package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)


func main() {
	var st []string
	var rep []string
	var repl []string
	var s []string
	var pick int
	var veg, amount, price string


	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	

	fmt.Println(`
		**---- Welcome to GreenGrocers ---- **
		`)

	fmt.Println(`
		GreenGrocers: Store Dashboard
		-----------------------------
		1.Visit the store stocks 
		2.Add new Vegetable item
		3.Search for details of a Vegetable
		4.Edit Vegetable Stock Amount  
		5.Edit Vegetable Stock Price  
		6.Exit
		`)

	fmt.Println("Enter your choice Number :")
	fmt.Scan(&pick)

	
	for ok := true; ok; ok = (pick != 0) {
		switch pick{

		case 1:
			client.Call("API.GetVegData","",&st)

			fmt.Println(`
				Currently in Store Stock 
				--------------------------
				`)	
			for _, str := range st {
          		fmt.Printf(" %s\n", str)
        	}
        	

        case 2:
			fmt.Println("Enter New Veg Details: Name |  Amount(Kg) | Price(Rs)")
			fmt.Scan(&veg,&amount, &price )
		   	client.Call("API.AddNewVeg","\n" + veg +" "+ amount +" "+ price,"")
		   	

		case 3:
			fmt.Println("Enter vegetable to search :")
			fmt.Scan(&veg)
			client.Call("API.GetByName", veg, &repl)
			fmt.Println(" \n Vegetable Details: ",repl)
			
			
		case 4:
			fmt.Println("Enter vegetable to change Amount :")
			fmt.Scan(&veg)
			fmt.Println("Enter Details: Amount ")
			fmt.Scan(&amount)
			client.Call("API.EditVegData", veg, &rep)
			client.Call("API.AddNewVeg","\n" + veg +" "+ amount +" "+ rep[2],"")
			

		case 5:
			fmt.Println("Enter vegetable to change Price :")
			fmt.Scan(&veg)
			fmt.Println("Enter Details: Price ")
			fmt.Scan(&price)
			client.Call("API.EditVegData", veg, &s)
			client.Call("API.AddNewVeg","\n" + veg +" "+ s[1] +" "+ price,"")
			

		case 0:
			os.Exit(2)

		}
		fmt.Println(
				` 

		Do you want to proceed? 
				If Yes - Enter the choice number
				If No  - Press 0`)
		fmt.Scan(&pick)
	}
}
