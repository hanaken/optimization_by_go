package main

import (
        "fmt"
        "encoding/csv"
        "io"
        "os"
        "math/rand"
        "time"
        "strconv"
)
const DATA_LENGTH int = 425

func hill_climbing(data_set [DATA_LENGTH]float64)(int, float64) {
        rand.Seed(time.Now().UnixNano())
        start := rand.Intn(DATA_LENGTH)
        right := -1.0
        left := -1.0
        var flg = false
        step := [...]int{-1,1}
        for {
                if flg {
                        break
                }
                if start > 0 && start < (DATA_LENGTH-1) {
                        //issue
                        right = data_set[start+1] - data_set[start]
                        left = data_set[start-1] - data_set[start]
                        if right > 0 && right > left {
                                start += 1
                        } else if left > 0 && left > right {
                                start -= 1
                        } else if right > 0 && right==left {
                                rand.Seed(time.Now().UnixNano())
                                start = start + step[rand.Intn(2)]
                        } else {
                                flg = true
                        }
                } else if start > 0 {
                        left = data_set[start-1] - data_set[start]
                        if left > 0 {
                                start -= 1
                        } else {
                                flg = true
                        }
                } else {
                        right = data_set[start+1] - data_set[start]
                        if right > 0 {
                                start += 1
                        } else {
                                flg = true
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
        x, y := hill_climbing(data_set)
        fmt.Printf("x=%d, y=%f", x, y)
}
