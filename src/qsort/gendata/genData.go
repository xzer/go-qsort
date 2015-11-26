package main 

import (
    "fmt"
    "qsort/data"
    "os"
    "strconv"
)

func main() {
    var size int = 10 * 10000
    if len(os.Args) > 1 {
        arg := os.Args[1]
        size, _ = strconv.Atoi(arg)
    }
    
    data.GenerateData(size)
    fmt.Printf("Genreated %d numbers.\n", size)
}

