package main

import "fmt"

// A supermarket wants to calculate the value of a customer's basket:

// Prices

//     Coffee - £1
//     Orange - £2
//     Bread - £3

// Basket

// ["coffee", "coffee", "orange", "orange", "orange", "bread"]
// Normally, this basket would cost £11, but today we have a special offer: you can now buy 2
// oranges for £3. This means that, today, the basket costs £10. Please implement a function that
// calculates the total value of the basket, including the special offer, and prints it to the
// console.

type offer struct {
	item     string
	quantity int
	price    int
}

func main() {
	// value in cents
	catalog := make(map[string]int)
	catalog["coffee"] = 100
	catalog["orange"] = 200
	catalog["bread"] = 300

	list := []string{"coffee", "coffee", "orange", "orange", "orange", "bread", "orange"}

	offers := []offer{
		{item: "orange", quantity: 2, price: 300},
	}

	total_price := 0
	list_with_quantities := map[string]int{}
	for _, v := range list {
		list_with_quantities[v] = list_with_quantities[v] + 1
	}

	for prod, quantity := range list_with_quantities {

		// look for an offer for this product
		var currentOffer offer
		for _, o := range offers {
			if o.item == prod {
				currentOffer = o
			}
		}

		for quantity > 0 {
			if currentOffer.item != "" && quantity >= currentOffer.quantity {
				fmt.Println(
					"add price (promo)",
					prod,
					currentOffer.quantity,
					currentOffer.price,
				)
				total_price += currentOffer.price
				quantity -= currentOffer.quantity
			} else {
				price, ok := catalog[prod]
				if !ok {
					panic("not in catalog")
				}
				fmt.Println("add price", prod, 1, price)
				total_price += price

				quantity--
			}
		}
	}

	fmt.Println(list_with_quantities)

	fmt.Println("total for the list:", total_price/100)
}
