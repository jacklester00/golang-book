package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    // Convert Fahrenheit to Celsius
    fmt.Println("Enter temperature in Fahrenheit:")
    inputStr, _ := reader.ReadString('\n')
    inputStr = strings.TrimSpace(inputStr)
    input, err := strconv.ParseFloat(inputStr, 64)
    if err != nil {
        fmt.Println("Invalid input:", err)
        return
    }
    temp := (input - 32) * 5 / 9
    fmt.Println("Temperature in Celsius:", temp)

    // Convert feet to meters
    fmt.Println("Enter feet to convert to meters:")
    feetStr, _ := reader.ReadString('\n')
    feetStr = strings.TrimSpace(feetStr)
    feet, err := strconv.ParseFloat(feetStr, 64)
    if err != nil {
        fmt.Println("Invalid input:", err)
        return
    }
    meters := feet * 0.3048
    fmt.Println("Meters:", meters)
}

