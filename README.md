![Compound Interest Calculator Logo](https://raw.githubusercontent.com/juba666/cic/master/repository-open-graph-template.png)
# Compound Interest Calculator
Compound interest is one of the most powerful forces in investing. Simple interest simply means a set percentage of the principal for the period, and is rarely used in practice. On the other hand, compound interest is applied to both loans, investments and deposit accounts.

- Calculate future amount for loan or investment

      f = p × ((r/100)+1) ^ n

- Calculate principal amount for loan or investment

      p = f / ((r/100)+1) ^ n

- Calculate interest rate for loan or investment

      r = (f/p) ^ 1/n - 1

- Calculate number of periods for loan or investment

      n = ln(f/p) / ln(1+r)
      
- Calculate periodic compounding future amount for loan or investment

      f = p × (((r/100)/n)+1) ^ n
      
- Calculate effective annual interest rate for loan or investment

      e = (1+((r/100)/n)) ^ n − 1
      
- Calculate present value of annuity 

      v = p × (1 − (1+r) ^ −n) / r

- Calculate value of each payment for annuity 

      p = v × r / (1 − (1+r) ^ −n)

## 1. Future Amount Calculator
```
f = p × ((r/100)+1) ^ n

f = Future Amount (£)
p = Principal Amount (£)
r = Interest Rate (%)
n = Number of Periods
s = Type of Period (days-weeks-months-years)
```
### An example
```
If the £1000 loan was for a period of 6 months at 10% interest rate, what is the future amount?:

1771.56 = 1000 × (0.1+1) ^ 6 months

£1,000 × (1.10 × 1.10 × 1.10 × 1.10 × 1.10 × 1.10) = £1771.56

f = £1771.56
p = £1000
r = 10% (10/100 for decimal)
n = 6
s = months
```
## 2. Principal Amount Calculator
```
p = f / ((r/100)+1) ^ n

f = Future Amount (£)
p = Principal Amount (£)
r = Interest Rate (%)
n = Number of Periods
s = Type of Period (days-weeks-months-years)
```
### An example
```
What principal amount do you need to invest, to get £2000 in 5 years at 10% interest rate?:

1241.84 = 2000 / (0.1+1) ^ 5 years

£2,000 / (1.10 × 1.10 × 1.10 × 1.10 × 1.10) = £1241.84

f = £2000
p = £1241.84
r = 10% (10/100 for decimal)
n = 5
s = years
```
## 3. Interest Rate Calculator
```
r = (f/p) ^ 1/n - 1

f = Future Amount (£)
p = Principal Amount (£)
r = Interest Rate (%)(x100 for percentage)
n = Number of Periods
s = Type of Period (days-weeks-months-years)
```
### An example
```
You have £1,000, and want it to grow to £2,000 in 5 Years, what interest rate do you need?:

0.1487 = (2000/1000) ^ 1/5 years - 1

(£2000/£1000) ^ 1/5 - 1 = 0.1487 x 100 = 14.87%

f = £2000
p = £1000
r = 14.87% 
n = 5
s = years
```
## 4. Number of Periods Calculator
```
(note: it uses the natural logarithm function ln)

n = ln(f/p) / ln(1+r)

f = Future Amount (£)
p = Principal Amount (£)
r = Interest Rate (%)
n = Number of Periods
s = Type of Period (days-weeks-months-years)
```
### An example
```
How many years will it take to turn £1,000 into £2,000 at 10% interest rate?:

7.27 = ln(2000/1000) / ln(1+(r/100))

ln(£2000/£1000) / ln(1+(10/100)) = 7.27 years

f = £2000
p = £1000
r = 10% (10/100 for decimal)
n = 7.27
s = years
```
## 5. Periodic Compounding Future Amount Calculator
```
f = p × (((r/100)/n)+1) ^ n

f = Future Amount (£)
p = Principal Amount (£)
r = Interest Rate (%)
n = Number of Periods within the year (Semiannually=2,Quarterly=4,Monthly=12,Daily=365,Continuously=2.71828182845904523536028747135266249775724709369995957
49669676277240766303535475945713821785251664274)
```
### An example
```
If £1000 was invested at 10% interest rate, "Compounded Semiannually"; what is the future amount?:

1102.5 = 1000 × (((10/100)/2)+1) ^ 2

£1000 × (((10/100)/2)+1) ^ 2 = £1102.50

f = £1102.50
p = £1000
r = 10% (10/100 for decimal)
n = 2 
```
## 6. Effective Annual Interest Rate Calculator
```
e = (1+((r/100)/n)) ^ n − 1

e = Effective Annual Interest Rate (%)(x100 for percentage)
r = Nominal Interest Rate (%)
n = Number of Periods within the year (Semiannually=2,Quarterly=4,Monthly=12,Daily=365,Continuously=2.71828182845904523536028747135266249775724709369995957
49669676277240766303535475945713821785251664274)
```
### An example
```
What effective annual interest rate do you get when the ad says "6% compounded monthly"?:

0.06168 = (1+((6/100)/12)) ^ 12 − 1

0.06168 x 100 = 6.168%

e = 6.168%
r = 6% (6/100 for decimal)
n = 12 
```
## 7. Present Value of Annuity Calculator
```
v = p × (1 − (1+r) ^ −n) / r

v = Present Value of Annuity 
p = Value of Each Payment for Annuity
n = Number of Periods
r = Interest Rate per Period
```
### An example
```
What is the present value for annuity of £400 a month for 5 years?, use a monthly interest rate of 1%:

17982.02 = 400 x (1 - (1+(1/100)) ^ -60) / 0.01

£400 x (1 - (1+(1/100)) ^ -60) / 0.01 = £17982.02

v = £17982.02 
p = £400
n = 60 (12 months x 5 years)
r = 1% (1/100 for decimal)
```
## 8. Value of Each Payment for Annuity Calculator
```
p = v × r / (1 − (1+r) ^ −n)

p = Value of Each Payment for Annuity
v = Present Value of Annuity 
n = Number of Periods
r = Interest Rate per Period
```
### An example
```
Say you have £10,000 and want to get a monthly income for 6 years out of it, 
how much could you get each month?, assume a monthly interest rate of 0.5%:

165.73 = 10000 × (0.5/100) / (1 − (1+(0.5/100)) ^ −72)

£10000 × (0.5/100) / (1 − (1+(0.5/100)) ^ −72) = £165.73

p = £165.73
v = £10000
n = 72 (12 months x 6 years)
r = 0.5% (0.5/100 for decimal)
```
## Built With 

Go - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

https://golang.org/

https://play.golang.org/

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use SemVer for versioning. For the versions available, see the tags on this repository.

## Author
```
Ganesh Niruban
```
See also the list of contributors who participated in this project.

## License

This project is licensed under the GNU General Public License v3.0 - see the LICENSE.md file for details.

## Acknowledgments

https://www.mathsisfun.com/money/compound-interest.html

https://en.wikipedia.org/wiki/Compound_interest
