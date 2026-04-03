package differenceofsquares

func SquareOfSum(n int) int {
	//panic("Please implement the SquareOfSum function")
	count := 0
    
    for i := 1; i <= n; i++{
        count += i
    }
    counts := count * count
    return counts
}

func SumOfSquares(n int) int {
	//panic("Please implement the SumOfSquares function")

    squares := []int{}
    
    for i := 1; i <= n; i++{
    	square := i * i
        squares = append(squares,square)
    }

    squaresCount := 0

    for _, square := range squares{
        squaresCount += square
    }

    return squaresCount
}

func Difference(n int) int {
	//panic("Please implement the Difference function")

    return SquareOfSum(n) - SumOfSquares(n)
}
