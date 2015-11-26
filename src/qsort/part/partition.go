package part

import (
    //"fmt"
)

func Partition(list []int, start int, end int) int{
    
    //fmt.Printf("parting %d -> %d", start, end)
    
    var index = (start + end) / 2
    var pivot = list[index]
   
    swap(list, index, end)
    index = start
    
    for i:=start; i <= end; i++{
        if list[i] < pivot {
            swap(list, i, index)
            index++
        }
    }
    swap(list, index, end)
    list[index] = pivot
    return index
}

func swap(list []int, a int, b int){
    t := list[a]
    list[a] = list[b]
    list[b] = t
}

