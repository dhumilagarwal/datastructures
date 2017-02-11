package main

import (
	"sort"
	"fmt"
)

/*
Given an array A[] of N integers where each value represents number of chocolates in a packet.
Each packet can have variable number of chocolates. There are m students, the task is to
distribute chocolate packets such that :
1. Each student gets one packet.
2. The difference between the number of chocolates given to the students in packet with maximum
chocolates and packet with minimum chocolates is minimum.
 */

/*
Problem Source Url : http://www.geeksforgeeks.org/chocolate-distribution-problem/
Chocolate Distribution Problem - Flipkart - GeeksforGeeks
 */

/*
Problem Analogy :
Objective = Maximum - Minimum should be Minimized.
This is Possible, by Maximizing the Minimum and minimizing the maximum.
This is a problem to minimize or reduce the range of selected Numbers.
Hence all the selected numbers should be as close as possible.
This is possible when the numbers are sorted.
First Step - Sort the Numbers.
then find the range between Ith and I+k th Number where k is the number of students.
Range = I+k th Number minus I th Number
Run this range finding step once for the whole array, and find which case gives the
lowest range.
 */

/*
Time Complexity = O(nlogn) + O(n) = O(nlogn)
O(nlogn) for Sorting.
O(n) for searching in the minimum range in the Sorted Array
 */
func main() {
	chocolatesInPackets := []int{7,3,2,4,9,12,56}
	studentsCount := 3
	fmt.Println(DistributeChocolatesWithMinimumDifference(chocolatesInPackets,studentsCount))
}

// ChocolatesInPackets contains the number of chocolates in each Packet
// countOfStudents contains the number of Students among whom the chocolates are to be distributed
func DistributeChocolatesWithMinimumDifference(chocolatesInPackets []int,countOfStudents int) int{

	// Get the Number of Packets Available.
	countOfPackets := len(chocolatesInPackets)

	// If the number of students is more than the number of packets, then it is impossible to
	// give 1 packet to every student. So we return a Sentinel Value -1
	if countOfStudents>countOfPackets {
		return -1
	}

	// If the Number of Packets is 0 then distribution cannot happen. If the number of Students
	// is negative or zero, then distribution cannot happen. In both cases, the Algorithm returns
	// a Sentinel Value -1
	if (countOfPackets==0 || countOfStudents<=0){
		return -1
	}

	// Sort the Number of Chocolates in each Packet
	// This sorting is done in-place. So the original slice is modified, and hence,
	// we need not store the sorted data in another slice.
	sort.Ints(chocolatesInPackets)

	// The variable to be used for iterating though the chocolatesInPackets slice
	startIndex := 0

	// This variable stores the current difference between maximum and minimum chocolates for
	// a given iteration
	currentRange := 0

	// Maximum number of combinations of countOfStudents number of students possible
	maximumPossibleCombinations := countOfPackets - countOfStudents + 1

	// The minimum range - The answer that has to be returned
	minimumRange := chocolatesInPackets[countOfStudents-1]- chocolatesInPackets[0]

	// Iterate through the slice only till the number of possible combinations
	for startIndex=1;startIndex<maximumPossibleCombinations ;startIndex=startIndex+1  {
		// Get the current Range
		currentRange = chocolatesInPackets[startIndex + countOfStudents-1]- chocolatesInPackets[startIndex]
		// If curent range is less than the minimum, then assign current range to the minimum range
		if (currentRange<minimumRange) {
			minimumRange = currentRange
		}
	}
	// return the final answer minimumRange
	return minimumRange
}
