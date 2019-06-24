package main

import (

"fmt"
"math"

)

func main() {

var p float64
var s string
var n float64
var r float64

  fmt.Print("Enter a Principal Amount(£): ")
    fmt.Scanf("%f\n", &p)

fmt.Print("Enter a Type of Period(days-wks-mnths-yrs): ")
    fmt.Scanf("%s\n", &s)

  fmt.Print("Enter a Number of Periods("+s+"): ")
    fmt.Scanf("%f\n", &n)

  fmt.Print("Enter a Interest Rate(%): ")
    fmt.Scanf("%f\n", &r)

   f := math.Pow(float64(r)/100+1, float64(n))*p

  fmt.Println("Future Amount:£",f)

}