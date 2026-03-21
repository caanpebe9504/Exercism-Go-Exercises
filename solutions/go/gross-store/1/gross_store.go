package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    
    units := map[string]int{
		"quarter_of_a_dozen": 3,
        "half_of_a_dozen" : 6,
        "dozen" : 12,
        "small_gross" : 120,
        "gross" : 144,
        "great_gross": 1728, 
    }
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	customerBill := make(map[string]int)
    return customerBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

    quantity, isUnit := units[unit]
    
    /*for key, value := range units{
        if(key == unit){
            if _,ok := bill[item];ok{
                bill[item] += value
            }else{
                bill[item] = value
            }
            return true
        }
    }
	return false*/
    if !isUnit {
        return false
    }
    
    bill[item] += quantity

    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    
	currentItem, itemValidation := bill[item]
    currentUnit, unitValidation := units[unit]

    if !itemValidation || !unitValidation {
        return false
    }

    newAmount := currentItem - currentUnit

    if newAmount < 0{
        return false
    }else if newAmount == 0{
        delete(bill,item)
    }else{
        bill[item] = newAmount
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {

    qty,ok := bill[item]

    if !ok{
        return 0, false
    }else{
        return qty, true
    }
    
    
    
}
