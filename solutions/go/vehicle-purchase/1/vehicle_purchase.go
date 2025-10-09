package purchase

import (
    "fmt"
    "strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
        return true
    } else {
        return false 
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if strings.Compare(option1, option2) <= 0 {
        return option1 + " is clearly the better choice."
    } 

    return option2 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    
	if age < 3 { 
        resellPrice := originalPrice * 0.8
    	fmt.Println(resellPrice)
        return resellPrice
    } else if age >= 10 {
        resellPrice := originalPrice * 0.5
    	fmt.Println(resellPrice)
        return resellPrice
    } else {
        resellPrice := originalPrice * 0.7
        return resellPrice  
    }
}
