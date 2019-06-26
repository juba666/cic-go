package main
import(
"fmt"
"strings"
"math"
"os"
)
func main() {
fmt.Println("CHOOSE FROM THE FOLLOWING OPTIONS:\n"+
"                           \n"+
"1.Future Amount Calculator\n"+
"2.Principal Amount Calculator\n"+
"3.Interest Rate Calculator\n"+
"4.Number of Periods Calculator\n"+
"5.Periodic Compounding Future Amount Calculator\n"+
"6.Effective Annual Interest Rate Calculator\n"+
"7.Exit\n")                         
const FA = 1
const PA = 2
const IR = 3
const NP = 4
const PCFA = 5
const EAIR = 6
const EX = 7
var CH float64
fmt.Print("Enter a choice(1-7): ")
fmt.Scanf("%f\n", &CH)
if CH == FA {
        fmt.Println("You have chosen Future Amount Calculator")
		calcFA()
} else if CH == PA {
	fmt.Println("You have chosen Principal Amount Calculator")
		calcPA()
} else if CH == IR {
	fmt.Println("You have chosen Interest Rate Calculator")
		calcIR()
} else if CH == NP {
	fmt.Println("You have chosen Number of Periods Calculator")
		calcNP()
} else if CH == PCFA {
	fmt.Println("You have chosen Periodic Compounding Future Amount Calculator")
		calcPCFA()
} else if CH == EAIR {
	fmt.Println("You have chosen Effective Annual Interest Rate Calculator")
		calcEAIR()
} else if CH == EX {
	fmt.Println("Thank You for using Compound Interest Calculator!")
		os.Exit(0)
} else {
	error := "ERROR: you did not enter a number between 1 & 7"
        fmt.Println("\n",error)
		main()
}
}
func calcFA() {
var p float64
var s string
var n float64
var r float64
fmt.Print("Enter a Principal Amount(£): ")
fmt.Scanf("%f\n", &p)
tp := fmt.Sprintf("%f", p)
if !(p > 0) || ((strings.ContainsRune(tp, '\u00a3')) || ((strings.ContainsRune(tp, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a currency symbol")
		calcFA()
} else {
fmt.Print("Enter a Type of Period(days-weeks-months-years): ")
fmt.Scanf("%s\n", &s)
fmt.Print("Enter a Number of Periods("+s+"): ")}
fmt.Scanf("%f\n", &n)
if !(n > 0) {
        fmt.Println("ERROR: you did not enter a number greater than zero")
		calcFA()
} else {
fmt.Print("Enter a Interest Rate(%): ")}
fmt.Scanf("%f\n", &r)
if !(r > 0) {
        fmt.Println("ERROR: you did not enter a number greater than zero")
		calcFA()
} else {
f := math.Pow(float64(r)/100+1, float64(n))*p
F := fmt.Sprintf("%.2f", f)
defer fmt.Println("Future Amount:£",F)
}
}
func calcPA() {
var f float64
var s string
var n float64
var r float64
fmt.Print("Enter a Future Amount(£): ")
fmt.Scanf("%f\n", &f)
fmt.Print("Enter a Type of Period(days-weeks-months-years): ")
fmt.Scanf("%s\n", &s)
fmt.Print("Enter a Number of Periods("+s+"): ") 
fmt.Scanf("%f\n", &n) 
fmt.Print("Enter a Interest Rate(%): ")
fmt.Scanf("%f\n", &r)
p := math.Pow(float64(r)/100+1, float64(n)) 
fp := f/p
P := fmt.Sprintf("%.2f", fp) 
fmt.Println("Principal Amount:£",P)
}
func calcIR() {
var f float64
var s string
var n float64
var p float64
fmt.Print("Enter a Principal Amount(£): ")
fmt.Scanf("%f\n", &p)
fmt.Print("Enter a Future Amount(£): ")
fmt.Scanf("%f\n", &f)
fmt.Print("Enter a Type of Period(days-weeks-months-years): ")
fmt.Scanf("%s\n", &s)
fmt.Print("Enter a Number of Periods("+s+"): ") 
fmt.Scanf("%f\n", &n) 
fp := f/p
nd := 1/n
r := (math.Pow(float64(fp), float64(nd))-1)*100
R := fmt.Sprintf("%.2f", r) 
fmt.Println("Interest Rate:",R,"%")
}
func calcNP() {
var f float64
var s string
var r float64
var p float64
fmt.Print("Enter a Principal Amount(£): ")
fmt.Scanf("%f\n", &p)
fmt.Print("Enter a Future Amount(£): ")
fmt.Scanf("%f\n", &f)
fmt.Print("Enter a Type of Period(days-weeks-months-years): ")
fmt.Scanf("%s\n", &s)
fmt.Print("Enter a Interest Rate(%): ")
fmt.Scanf("%f\n", &r)
fp := f/p
rd := (r/100)+1
n := math.Log(fp)/math.Log(rd)
N := fmt.Sprintf("%.2f", n) 
fmt.Println("Number of",s,":",N)
}
func calcPCFA() {
var p float64
var n float64
var r float64
fmt.Print("Enter a Principal Amount(£): ")
fmt.Scanf("%f\n", &p)
fmt.Print("Enter a Number of Periods within the year: ") 
fmt.Scanf("%f\n", &n) 
fmt.Print("Enter a Interest Rate(%): ")
fmt.Scanf("%f\n", &r)
rn := (r/100)/n
f := math.Pow(float64(rn)+1, float64(n))*p
F := fmt.Sprintf("%.2f", f) 
fmt.Println("Periodic Compounding Future Amount:£",F)
}
func calcEAIR() {
var r float64
var n float64
fmt.Print("Enter a Number of Periods within the year: ") 
fmt.Scanf("%f\n", &n) 
fmt.Print("Enter a Nominal Interest Rate(%): ")
fmt.Scanf("%f\n", &r)
rn := (r/100)/n
R := math.Pow(float64(rn)+1, float64(n))-1
e := R*100
E := fmt.Sprintf("%.3f", e) 
fmt.Println("Effective Annual Interest Rate:",E,"%")
}
