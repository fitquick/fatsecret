/*

Simple test client that shows how to use the FatSecret go library.

The following environment variables must be set in order to use the
FatSecret API...

1) "FATSECRET_CONSUMER_KEY"
2) "FATSECRET_CONSUMER_SECRET"

*/

package main

import (
	"fmt"
	"os"

	"github.com/fitzone/fatsecret"
)

var (
	consumerKey  = os.Getenv("FATSECRET_CONSUMER_KEY")
	sharedSecret = os.Getenv("FATSECRET_CONSUMER_SECRET")
)

func main() {
	// create a fatsecret client
	client, err := fatsecret.NewClient(consumerKey, sharedSecret)
	if err != nil {
		panic(err)
	}

	// search for food by name
	foods, err := client.FoodSearch("coffee")
	if err != nil {
		fmt.Printf("Cannot fetch food from API: err = '%v'", err)
	}
	for _, f := range foods {
		fmt.Printf("FOOD: name = %s\n", f.Name)
	}

	// search for brands by type
	brands, err := client.FoodBrandsByType(fatsecret.BrandTypeManufacturer)
	if err != nil {
		fmt.Printf("Could not fetch brands by type")
	}
	for _, b := range brands {
		fmt.Printf("BRAND MFG: name = %s\n", b)
	}

	// search for brands starting with a letter (use '*' for starting with numbers)
	brands2, err := client.FoodBrandsStartingWith("V")
	if err != nil {
		fmt.Printf("Could not fetch brands by type")
	}
	for _, b := range brands2 {
		fmt.Printf("BRAND STARTS WITH: name = %s\n", b)
	}

	// get the list of food categories
	categories, err := client.FoodCategories()
	if err != nil {
		fmt.Printf("Could not fetch food categories")
	}
	for _, cat := range categories {
		fmt.Printf("CATEGORY: id = '%s', name = '%s'\n", cat.ID, cat.Name)
	}

	// get the list of food sub-categories for a category
	categoryID := "16" // desserts
	subs, err := client.FoodSubCategories(categoryID)
	if err != nil {
		fmt.Printf("Could not fetch food sub-categories")
	}
	for _, s := range subs {
		fmt.Printf("SUB-CATEGORY: name = %s\n", s)
	}

	// find the food id for a barcode
	foodID, err := client.FoodIDForBarcode("0748927052688")
	if err != nil {
		fmt.Printf("Could not fetch brands by type")
	}
	fmt.Printf("\nFOOD ID FOR BARCODE: id = %v\n", foodID)

	// find the detailed food info by id
	foodInfo, err := client.FoodByID("2415647")
	if err != nil {
		fmt.Printf("Could not fetch food info")
	}
	fmt.Printf("\nFOOD INFO: %+v\n", foodInfo)

	// invoke the raw low-level api (used by all higher-level api's)
	body, err := client.InvokeAPI(
		"food_brands.get",
		map[string]string{
			"starts_with": "kraft",
		},
	)
	if err != nil {
		fmt.Printf("Cannot fetch food brands from API: err = '%v'", err)
	}
	fmt.Printf("\nBODY = %v\n", string(body))
}
