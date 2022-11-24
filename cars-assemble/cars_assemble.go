package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	fltProductionRate := float64(productionRate)
	result := fltProductionRate * successRate / 6000
	return int(result)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	remainder := carsCount % 10
	tens := (carsCount - remainder) / 10
	return uint(tens*95000) + uint(remainder*10000)
}
