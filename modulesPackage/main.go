package main

import (
    "fmt"

    qt "github.com/satyabansahoo2000/modulesPackage/quotes"
    str "github.com/satyabansahoo2000/modulesPackage/strings"
)

func main() {
    o, e := str.CountOddEven("12345")
    fmt.Println(o, e) // 3 2

    fmt.Println(qt.GetEmoji("turtle"))
}