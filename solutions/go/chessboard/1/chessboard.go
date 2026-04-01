package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	//panic("Please implement CountInFile()")

    squares,ok := cb[file]
    if !ok{
        return 0
    }

    count := 0
    
    for _, exists := range squares {
       if exists{
           count ++
       } 
    }
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	//panic("Please implement CountInRank()")
     
    if rank < 1 || rank > 8{
        return 0
    }

    index := rank - 1
    count := 0
    
    for _, value := range cb{
        if value[index]{
            count ++
        }
    }
return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	//panic("Please implement CountAll()")

    count := 0
    
    for range cb{
        count += len(cb["A"]) 
    }
    return count 
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	//panic("Please implement CountOccupied()")

   count := 0
    
    // 1. Entramos a cada columna (file) del mapa
    for _, file := range cb {
        // 2. Entramos a cada casilla (occupied) de esa columna
        for _, occupied := range file {
            // 3. Si la casilla es true, sumamos
            if occupied {
                count++
            }
        }
    }
    return count 
}
