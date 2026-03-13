package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
    
    for i := 0; i < len(birdsPerDay); i++{
        count += birdsPerDay[i]
    }
    return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    
    startIndex := (week - 1) * 7
    endIndex := startIndex + 7

    newBirdsSlice := birdsPerDay[startIndex:endIndex]

    return TotalBirdCount(newBirdsSlice)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

    for i:= 0; i < len(birdsPerDay); i += 2{
        birdsPerDay[i] += 1
    }
    return birdsPerDay
}
