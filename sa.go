package main

import (
        "fmt"
        "encoding/csv"
        "io"
        "os"
        "math/rand"
        "math"
        "time"
        "strconv"
)
const DATA_LENGTH int = 425

func prob(e1 float64, e2 float64, t float64)(float64) {
        if e1 >= e2 {
                return 1.0
        } else {
                freq := math.Exp(-math.Abs(e1 - e2)/t)
                return freq
        }
}

func sa(data_set [DATA_LENGTH]float64)(int, float64) {
        rand.Seed(time.Now().Unix())
        start := rand.Intn(DATA_LENGTH)
        max_iter := 0.00001
        step := [...]int{-1,1}
        for iter:=10000.0; iter>max_iter; iter=iter*0.99 {
                rand.Seed(time.Now().Unix())
                if start > 0 && start < (DATA_LENGTH-1) {
                        rand.Seed(time.Now().UnixNano())
                        neigh := step[rand.Intn(2)]
                        next := data_set[start+neigh]
                        freq := prob(next, data_set[start], iter)
                        rand.Seed(time.Now().UnixNano())
                        if rand.Float64() <= freq {
                                start = start+neigh
                        }
                } else if start > 0 {
                        next := data_set[start-1]
                        freq := prob(next, data_set[start], iter)
                        rand.Seed(time.Now().UnixNano())
                        if rand.Float64() <= freq {
                                start -= 1
                        }
                } else {
                        next := data_set[start+1]
                        freq := prob(next, data_set[start], iter)
                        rand.Seed(time.Now().UnixNano())
                        if rand.Float64() <= freq {
                                start += 1 
                        }
                }
        }
        return start, data_set[start]
}

func input_csv()([DATA_LENGTH]float64) {
        var data_set [DATA_LENGTH]float64
        var fp *os.File

        var err error
        if len(os.Args) < 2 {
                fp = os.Stdin
        } else {
                fmt.Printf(">> reading %s\n", os.Args[1])
                fp, err = os.Open(os.Args[1])
                if err != nil {
                        panic(err)
                }
                defer fp.Close()
        }

        read := csv.NewReader(fp)
        read.Comma = ','
        read.LazyQuotes = true
        i := 0
        for {
                data, err := read.Read()
                if err == io.EOF {
                        break
                } else if err != nil {
                        panic(err)
                }
                y, _ := strconv.ParseFloat(data[1], 64)
                data_set[i] = y
                i += 1 
        }
        return data_set
}

func main() {
        data_set := input_csv()
        x, y := sa(data_set)
        fmt.Printf("x=%d, y=%f", x, y)
}
