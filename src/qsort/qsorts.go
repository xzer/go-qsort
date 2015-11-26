package main 

import (
    "fmt"
    "qsort/part"
    "qsort/data"
    "qsort/common"
    "runtime"
)

func fuck(){
    var _ = fmt.Print
}

func main() {
    var list []int = data.GetData() //first time read data from disk
    common.MeasureTime("qsort_classical", func(){
            list = data.GetData()
            qsort_classical(list, 0, len(list)-1)
    })
    data.StoreData(list, "qsort_classical_result")
    
    for i:=1; i <= runtime.NumCPU(); i++{
        measureGoRoutine(i)
    }
    
    for i:=1; i <= runtime.NumCPU(); i++{
        measureGoRoutine_single_channel(i)
    }
    
    for i:=1; i <= runtime.NumCPU(); i++{
        measureGoRoutine_ultimate_optimization(i)
    }
    
    fmt.Println("finished")
}

func measureGoRoutine(core int){
    runtime.GOMAXPROCS(core)
    var list []int 
    log := fmt.Sprintf("qsort_goroutine(%d)", core)
    common.MeasureTime(log, func(){
            list = data.GetData()
            semaphoreChannel := make(chan int)
            go qsort_goroutine(list, 0, len(list)-1, semaphoreChannel)
            for{
                <- semaphoreChannel
                break
            }
    })
    data.StoreData(list, fmt.Sprintf("qsort_goroutine(%d)_result", core))
}

func measureGoRoutine_single_channel(core int){
    runtime.GOMAXPROCS(core)
    var list []int 
    log := fmt.Sprintf("qsort_goroutine_single_channel(%d)", core)
    common.MeasureTime(log, func(){
            list = data.GetData()
            count := 1
            var op int
            semaphoreChannel := make(chan int, len(list))
            go qsort_goroutine_single_channel(list, 0, len(list)-1, semaphoreChannel)
            for{
                op = <- semaphoreChannel
                count += op
                if count == 0 {
                    break
                }
            }
    })
    data.StoreData(list, fmt.Sprintf("qsort_goroutine_single_channel(%d)_result", core))
}

func measureGoRoutine_ultimate_optimization(core int){
    runtime.GOMAXPROCS(core)
    var list []int 
    log := fmt.Sprintf("qsort_goroutine_ultimate_optimization(%d)", core)
    common.MeasureTime(log, func(){
            list = data.GetData()
            count := 1
            var op int
            semaphoreChannel := make(chan int, len(list))
            go qsort_goroutine_ultimate_optimization(list, 0, len(list)-1, semaphoreChannel)
            for{
                op = <- semaphoreChannel
                count += op
                if count == 0 {
                    break
                }
            }
    })
    data.StoreData(list, fmt.Sprintf("qsort_goroutine_ultimate_optimization(%d)_result", core))
}

func qsort_classical(list []int, start int, end int){
    if start >= end {
         return
    }
    
    index := part.Partition(list, start, end)
    if start < index-1 {
        qsort_classical(list, start, index-1)
    }
    
    if index+1 < end {
        qsort_classical(list, index+1, end)
    }
}

func qsort_goroutine(list []int, start int, end int, semaphoreChannel chan<- int){
    if start >= end {
        semaphoreChannel <- 0
        return
    }
    
    index := part.Partition(list, start, end)
    
    subChannel := make(chan int)
    var subCount int = 0
    
    if start < index-1 {
        go qsort_goroutine(list, start, index-1, subChannel)
        subCount++
    }
    if index+1 < end {
        go qsort_goroutine(list, index+1, end, subChannel)
        subCount++
    }
    if subCount>0 {
        for i:= 0; i<subCount; i++ {
            <- subChannel
        }
    }
    semaphoreChannel <- 0
}

func qsort_goroutine_single_channel(list []int, start int, end int, semaphoreChannel chan<- int){
    if start >= end {
        semaphoreChannel <- -1
        return
    }
    
    index := part.Partition(list, start, end)
    
    if start < index-1 {
        semaphoreChannel <- 1
        go qsort_goroutine_single_channel(list, start, index-1, semaphoreChannel)
    }
    if index+1 < end {
        semaphoreChannel <- 1
        go qsort_goroutine_single_channel(list, index+1, end, semaphoreChannel)
    }

    semaphoreChannel <- -1
}

func qsort_goroutine_ultimate_optimization(list []int, start int, end int, semaphoreChannel chan<- int){
    if start >= end {
        if semaphoreChannel != nil {
            semaphoreChannel <- -1
        }
        return
    }
    
    index := part.Partition(list, start, end)
    
    if end-start > 10000 {
        if start < index-1 {
            if semaphoreChannel != nil {
                semaphoreChannel <- 1
            }
            go qsort_goroutine_ultimate_optimization(list, start, index-1, semaphoreChannel)
        }
        if index+1 < end {
            if semaphoreChannel != nil {
                semaphoreChannel <- 1
            }
            go qsort_goroutine_ultimate_optimization(list, index+1, end, semaphoreChannel)
        }
    }else{
        if start < index-1 {
            qsort_goroutine_ultimate_optimization(list, start, index-1, nil)
        }
        if index+1 < end {
            qsort_goroutine_ultimate_optimization(list, index+1, end, nil)
        }
    }

    if semaphoreChannel != nil {
        semaphoreChannel <- -1
    }
}


