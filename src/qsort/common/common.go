package common

import (
    "time"
    "fmt"
)

type Runner func()

func CkErr(e error) {
    if e != nil {
        panic(e)
    }
}

func MeasureTime(log string, runner Runner){
    var loop int64 = 100
    var sum int64
    for i:=int64(0); i<loop; i++{
        start := time.Now()
        runner()
        end := time.Now()
        sum += end.UnixNano() - start.UnixNano()
    }
    
    avg := time.Duration(sum / loop)
    
    fmt.Printf("%s takes %s\n", log, avg)
}