package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64)  float64{
    
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

    carsPerMinute := CalculateWorkingCarsPerHour( productionRate, successRate) / 60
    
	return int(carsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    // Calculamos grupos de 10 y unidades sueltas
    gruposDeDiez := carsCount / 10
    autosRestantes := carsCount % 10

    // Calculamos el costo total
    // 95000 por cada grupo y 10000 por cada auto individual
    costo := (gruposDeDiez * 95000) + (autosRestantes * 10000)

    return uint(costo)
}
