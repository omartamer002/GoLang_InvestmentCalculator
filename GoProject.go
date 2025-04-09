/*
In this application we are going to create an investment calculator that calculates the future value of an investment based on the principal amount, the interest rate, and the time period.
Also it calculates the real future value according to the inflation rate of the year...
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64 
	fmt.Print("The investment amount is : ")
	fmt.Scan(&investmentAmount)
	fmt.Print("The expected return rate is : ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("The number of years is : ")
	fmt.Scan(&years)
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var realFutureValue = futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(realFutureValue)
}
