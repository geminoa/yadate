package main

import (
    "fmt"
    "time"
)

func main() {
    tz := "America/New_York"
    a, err := time.LoadLocation(tz)

    if err != nil {
        panic(err)
    }

    now := time.Now().In(a)
    fmt.Println(now)
}
