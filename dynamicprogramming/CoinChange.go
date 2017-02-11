package main

import "fmt"

/*
Problem Statement : Given a value N, if we want to make change for N cents, and we have infinite
supply of each of S = { S1, S2, .. , Sm} valued coins, how many ways can we make the change?
The order of coins doesn’t matter.

For example, for N = 4 and S = {1,2,3}, there are four solutions: {1,1,1,1},{1,1,2},{2,2},{1,3}.
So output should be 4. For N = 10 and S = {2, 5, 3, 6}, there are five solutions:
{2,2,2,2,2}, {2,2,3,3}, {2,2,6}, {2,3,5} and {5,5}. So the output should be 5.
 */

/*
Problem Source Url : http://www.geeksforgeeks.org/dynamic-programming-set-7-coin-change/
 */

/*
Problem Analogy : We can find the number of ways we get 1 cents, 2 cents, 3 cents and so on
Let Ways(N) = Number of Ways we want to make change for N cents
Ways(N) = Sum(Ways(N-Si) + 1) , where Si is a valued coin among S1, S2, S3 ...and upto Sm
We make an array Ways[1..N]  where we store Ways[i].
 */

func main() {

	coinsAvailable := []int{2,5,3,6}
	fmt.Println(GetNumberOfWaysToMakeChange(10,coinsAvailable))
}

func GetNumberOfWaysToMakeChange(totalValue int,coinValues []int) int{
	var ways = make([]int,totalValue+1)
	index := 0
	coinIndex := 0
	for index=0;index<=totalValue;index = index+1  {
		ways[index]=0
	}
	var totalWays int
	totalWays = 0
	coinTypesCount := len(coinValues)
	if coinTypesCount==0 {
		return totalWays
	}
	for index=0;index<coinTypesCount ;index=index+1  {
		ways[coinValues[index]] = 1
	}
	for index=1;index<=totalValue;index=index+1  {
		for coinIndex=0;coinIndex<coinTypesCount ;coinIndex=coinIndex+1  {
			if ((index-coinValues[coinIndex])>0) {
				ways[index] = ways[index] + ways[index-coinValues[coinIndex]]
			}
		}
	}
	return ways[totalValue]
}
