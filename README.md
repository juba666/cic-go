# cic.go

## Compound Interest Calculator
Compound interest is one of the most powerful forces in investing. Simple interest simply means a set percentage of the principal for the period, and is rarely used in practice. On the other hand, compound interest is applied to both loans, investments and deposit accounts.
```
f = p × ((r/100)+1) ^ n

f = Future Amount(£)
p = Principal Amount(£)
r = Interest Rate(%)
n = Number of Periods
s = Type of Period (days-weeks-months-years)
```
### An example
```
If the £1000 loan was for a period of 6 months at 10%:

1771.56 = 1000 × (0.1+1) ^ 6 months

£1,000 × 1.10 × 1.10 × 1.10 × 1.10 × 1.10 × 1.10 = £1771.56

f = £1771.56
p = £1000
r = 10% (10/100 for decimal)
n = 6
s = months
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
