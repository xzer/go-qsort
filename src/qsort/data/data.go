package data

import (
    "math/rand"
    "fmt"
    "time"
    "os"
    "qsort/common"
    "bufio"
    "strconv"
)

var filepath string = "inout/nums"
var numList []int

func readData(){
    f, err := os.Open(filepath)
    common.CkErr(err)
    scanner := bufio.NewScanner(f)
    var s string
    for scanner.Scan() {
        s = scanner.Text()
        if len(s) == 0 {
            continue
        }
        i, err := strconv.Atoi(s)
        common.CkErr(err)
        //fmt.Println(i)
        numList = append(numList, i)
        //fmt.Println(numList)
    }
    fmt.Println("data read")
    //fmt.Println(numList)
}

func GetData() []int{
    if numList == nil {
        readData()
    }
    rtn := make([]int, len(numList))
    copy(rtn, numList)
    return rtn
}

func GenerateData(size int){
    source := rand.NewSource(time.Now().UnixNano())
    rd := rand.New(source)
    f, err := os.Create(filepath)
    common.CkErr(err)
    defer f.Close()
    
    var n int
    for i:=0;i<size;i++{
        if i % 10000 == 0 {
            fmt.Printf("Generating %d ...\n", i+1)
        }
        n = rd.Intn(size*10)
        f.WriteString(fmt.Sprintf("%d\n", n))
    }
    f.Sync()
}

func StoreData(list []int, fp string){
    f, err := os.Create(fmt.Sprintf("inout/%s", fp))
    common.CkErr(err)
    defer f.Close()
    
    for _, i := range list{
        f.WriteString(fmt.Sprintf("%d\n", i))
    }
    f.Sync()
}