package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    
	return kind == "truck" || kind == "car"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    var choice string
    
    if option1 < option2 {
        choice = option1
    } else {
        choice = option2
    }
    
    return choice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

    var percentage float64
    
    if age < 3{
        percentage = 0.80
    }else if age < 10{
        percentage = 0.70
    }else{
        percentage = 0.50
    }

    return originalPrice * percentage
}
