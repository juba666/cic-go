package main
import(
"fmt"
"strings"
"math"
"os"
)
const FA = 1
const PA = 2
const IR = 3
const NP = 4
const PCFA = 5
const EAIR = 6
const PV = 7
const VP = 8
const EX = 9
const HP = 10
var CH float64
var counter int
func main() {
if counter > 0 {
goto choice
} else {
fmt.Println("Welcome to cic! (Compound Interest Calculator)\n"+
"                           \n"+
"CHOOSE FROM THE FOLLOWING OPTIONS:\n"+
"                           \n"+
"1.Future Amount Calculator\n"+
"2.Principal Amount Calculator\n"+
"3.Interest Rate Calculator\n"+
"4.Number of Periods Calculator\n"+
"5.Periodic Compounding Future Amount Calculator\n"+
"6.Effective Annual Interest Rate Calculator\n"+
"7.Present Value of Annuity Calculator\n"+
"8.Value of Each Payment for Annuity Calculator\n")}                   
choice:
fmt.Print("[9=Exit,10=Help]\n"+
"Enter a choice(1-10): ")
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
} else if CH == PV {
	fmt.Println("You have chosen Present Value of Annuity Calculator")
		calcPV()
} else if CH == VP {
	fmt.Println("You have chosen Value of Each Payment for Annuity Calculator")
		calcVP()
} else if CH == EX {
	fmt.Println("Thank You for using cic!")
		os.Exit(0)
} else if CH == HP {
	fmt.Println("CHOOSE FROM THE FOLLOWING OPTIONS:\n"+
"                           \n"+
"1.Future Amount Calculator\n"+
"2.Principal Amount Calculator\n"+
"3.Interest Rate Calculator\n"+
"4.Number of Periods Calculator\n"+
"5.Periodic Compounding Future Amount Calculator\n"+
"6.Effective Annual Interest Rate Calculator\n"+
"7.Present Value of Annuity Calculator\n"+
"8.Value of Each Payment for Annuity Calculator\n")
counter++
main()
} else {
	error := "ERROR: you did not enter a number between 1 & 10"
        fmt.Println(error)
		goto choice
}
}
func calcFA() {
var p float64
var s string
var n float64
var r float64
pa:
fmt.Print("Enter a Principal Amount(£): ")
fmt.Scanf("%f\n", &p)
sp := fmt.Sprintf("%f", p)
if !(p > 0) || ((strings.ContainsRune(sp, '\u00a3')) || ((strings.ContainsRune(sp, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto pa
} else {
fmt.Print("Enter a Type of Period(days-weeks-months-years): ")
fmt.Scanf("%s\n", &s)}
np:
fmt.Print("Enter a Number of Periods("+s+"): ")
fmt.Scanf("%f\n", &n)
if !(n > 0) {
        fmt.Println("ERROR: you did not enter a number greater than zero")
		goto np
} else {
goto ir}
ir:
fmt.Print("Enter a Interest Rate(%): ")
fmt.Scanf("%f\n", &r)
sr := fmt.Sprintf("%f", r)
if !(r > 0) || ((strings.ContainsRune(sr, '\u00a3')) || ((strings.ContainsRune(sr, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto ir
} else {
//Calculate future amount for loan or investment
//f = p × ((r/100)+1) ^ n
f := math.Pow(float64(r)/100+1, float64(n))*p
F := fmt.Sprintf("%.2f", f)
fmt.Println("Future Amount:£",F)
counter++
main()
}
}
func calcPA() {
var f float64
var s string
var n float64
var r float64
fa:
fmt.Print("Enter a Future Amount(£): ")
fmt.Scanf("%f\n", &f)
sf := fmt.Sprintf("%f", f)
if !(f > 0) || ((strings.ContainsRune(sf, '\u00a3')) || ((strings.ContainsRune(sf, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto fa
} else {
fmt.Print("Enter a Type of Period(days-weeks-months-years): ")
fmt.Scanf("%s\n", &s)}
np:
fmt.Print("Enter a Number of Periods("+s+"): ")
fmt.Scanf("%f\n", &n)
if !(n > 0) {
        fmt.Println("ERROR: you did not enter a number greater than zero")
		goto np
} else {
goto ir}
ir:
fmt.Print("Enter a Interest Rate(%): ")
fmt.Scanf("%f\n", &r)
sr := fmt.Sprintf("%f", r)
if !(r > 0) || ((strings.ContainsRune(sr, '\u00a3')) || ((strings.ContainsRune(sr, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto ir
} else {
//Calculate principal amount for loan or investment
//p = f / ((r/100)+1) ^ n
p := math.Pow(float64(r)/100+1, float64(n)) 
fp := f/p
P := fmt.Sprintf("%.2f", fp) 
fmt.Println("Principal Amount:£",P)
counter++
main()
}
}
func calcIR() {
var f float64
var s string
var n float64
var p float64
pa:
fmt.Print("Enter a Principal Amount(£): ")
fmt.Scanf("%f\n", &p)
sp := fmt.Sprintf("%f", p)
if !(p > 0) || ((strings.ContainsRune(sp, '\u00a3')) || ((strings.ContainsRune(sp, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto pa
} else {
goto fa}
fa:
fmt.Print("Enter a Future Amount(£): ")
fmt.Scanf("%f\n", &f)
sf := fmt.Sprintf("%f", f)
if !(f > 0) || ((strings.ContainsRune(sf, '\u00a3')) || ((strings.ContainsRune(sf, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto fa
} else {
fmt.Print("Enter a Type of Period(days-weeks-months-years): ")
fmt.Scanf("%s\n", &s)}
np:
fmt.Print("Enter a Number of Periods("+s+"): ") 
fmt.Scanf("%f\n", &n)
if !(n > 0) {
        fmt.Println("ERROR: you did not enter a number greater than zero")
		goto np
} else {
//Calculate interest rate for loan or investment
//r = (f/p) ^ 1/n - 1
fp := f/p
nd := 1/n
r := (math.Pow(float64(fp), float64(nd))-1)*100
R := fmt.Sprintf("%.2f", r) 
fmt.Println("Interest Rate:",R,"%")
counter++
main()
}
}
func calcNP() {
var f float64
var s string
var r float64
var p float64
pa:
fmt.Print("Enter a Principal Amount(£): ")
fmt.Scanf("%f\n", &p)
sp := fmt.Sprintf("%f", p)
if !(p > 0) || ((strings.ContainsRune(sp, '\u00a3')) || ((strings.ContainsRune(sp, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto pa
} else {
goto fa}
fa:
fmt.Print("Enter a Future Amount(£): ")
fmt.Scanf("%f\n", &f)
sf := fmt.Sprintf("%f", f)
if !(f > 0) || ((strings.ContainsRune(sf, '\u00a3')) || ((strings.ContainsRune(sf, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto fa
} else {
fmt.Print("Enter a Type of Period(days-weeks-months-years): ")
fmt.Scanf("%s\n", &s)}
ir:
fmt.Print("Enter a Interest Rate(%): ")
fmt.Scanf("%f\n", &r)
sr := fmt.Sprintf("%f", r)
if !(r > 0) || ((strings.ContainsRune(sr, '\u00a3')) || ((strings.ContainsRune(sr, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto ir
} else {
//Calculate number of periods for loan or investment
//n = ln(f/p) / ln(1+r)
fp := f/p
rd := (r/100)+1
n := math.Log(fp)/math.Log(rd)
N := fmt.Sprintf("%.2f", n) 
fmt.Println("Number of",s,":",N)
counter++
main()
}
}
func calcPCFA() {
var p float64
var n float64
var r float64
pa:
fmt.Print("Enter a Principal Amount(£): ")
fmt.Scanf("%f\n", &p)
sp := fmt.Sprintf("%f", p)
if !(p > 0) || ((strings.ContainsRune(sp, '\u00a3')) || ((strings.ContainsRune(sp, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto pa
} else {
goto np}
np:
fmt.Print("Enter a Number of Periods within the year: ") 
fmt.Scanf("%f\n", &n)
if !(n > 0) {
        fmt.Println("ERROR: you did not enter a number greater than zero")
		goto np
} else {
goto ir}
ir:
fmt.Print("Enter a Interest Rate(%): ")
fmt.Scanf("%f\n", &r)
sr := fmt.Sprintf("%f", r)
if !(r > 0) || ((strings.ContainsRune(sr, '\u00a3')) || ((strings.ContainsRune(sr, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto ir
} else {
//Calculate periodic compounding future amount for loan or investment
//f = p × (((r/100)/n)+1) ^ n
rn := (r/100)/n
f := math.Pow(float64(rn)+1, float64(n))*p
F := fmt.Sprintf("%.2f", f) 
fmt.Println("Periodic Compounding Future Amount:£",F)
counter++
main()
}
}
func calcEAIR() {
var r float64
var n float64
np:
fmt.Print("Enter a Number of Periods within the year: ") 
fmt.Scanf("%f\n", &n)
if !(n > 0) {
        fmt.Println("ERROR: you did not enter a number greater than zero")
		goto np
} else {
goto ir}
ir:
fmt.Print("Enter a Nominal Interest Rate(%): ")
fmt.Scanf("%f\n", &r)
sr := fmt.Sprintf("%f", r)
if !(r > 0) || ((strings.ContainsRune(sr, '\u00a3')) || ((strings.ContainsRune(sr, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto ir
} else {
//Calculate effective annual interest rate for loan or investment
//e = (1+((r/100)/n)) ^ n − 1
rn := (r/100)/n
R := math.Pow(float64(rn)+1, float64(n))-1
e := R*100
E := fmt.Sprintf("%.3f", e) 
fmt.Println("Effective Annual Interest Rate:",E,"%")
counter++
main()
}
}
func calcPV() {
var p float64
var n float64
var r float64
pv:
fmt.Print("Enter a Value of Each Payment for Annuity(£): ") 
fmt.Scanf("%f\n", &p)
sp := fmt.Sprintf("%f", p)
if !(p > 0) || ((strings.ContainsRune(sp, '\u00a3')) || ((strings.ContainsRune(sp, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero")
		goto pv
} else {
goto np}
np:
fmt.Print("Enter a Number of Periods: ") 
fmt.Scanf("%f\n", &n)
if !(n > 0) {
        fmt.Println("ERROR: you did not enter a number greater than zero")
		goto np
} else {
goto ir}
ir:
fmt.Print("Enter a Interest Rate per Period(%): ")
fmt.Scanf("%f\n", &r)
sr := fmt.Sprintf("%f", r)
if !(r > 0) || ((strings.ContainsRune(sr, '\u00a3')) || ((strings.ContainsRune(sr, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto ir
} else {
//Calculate present value of annuity
//v = p × (1 − (1+r) ^ −n) / r
tr := (r/100)+1
tn := n*-1
tv := math.Pow(float64(tr), float64(tn))
v := ((1-tv)/(r/100))*p
V := fmt.Sprintf("%.2f", v) 
fmt.Println("Present Value of Annuity:£",V)
counter++
main()
}
}
func calcVP() {
var v float64
var n float64
var r float64
vp:
fmt.Print("Enter a Present Value of Annuity(£): ") 
fmt.Scanf("%f\n", &v)
sv := fmt.Sprintf("%f", v)
if !(v > 0) || ((strings.ContainsRune(sv, '\u00a3')) || ((strings.ContainsRune(sv, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero")
		goto vp
} else {
goto np}
np:
fmt.Print("Enter a Number of Periods: ") 
fmt.Scanf("%f\n", &n)
if !(n > 0) {
        fmt.Println("ERROR: you did not enter a number greater than zero")
		goto np
} else {
goto ir}
ir:
fmt.Print("Enter a Interest Rate per Period(%): ")
fmt.Scanf("%f\n", &r)
sr := fmt.Sprintf("%f", r)
if !(r > 0) || ((strings.ContainsRune(sr, '\u00a3')) || ((strings.ContainsRune(sr, '$')))) {
        fmt.Println("ERROR: you did not enter a number greater than zero, or included a symbol")
		goto ir
} else {
//Calculate value of each payment for annuity 
//p = v × r / (1 − (1+r) ^ −n)
tr := (r/100)+1
tn := n*-1
tp := math.Pow(float64(tr), float64(tn))
p := ((r/100)/(1-tp))*v
P := fmt.Sprintf("%.2f", p) 
fmt.Println("Value of Each Payment for Annuity:£",P)
counter++
main()
}
}
