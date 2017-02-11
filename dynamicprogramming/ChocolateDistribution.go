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
func main() {
	chocolatesInPackets := []int{7,3,2,4,9,12,56}
	studentsCount := 3
	fmt.Println(DistributeChocolatesWithMinimumDifference(chocolatesInPackets,studentsCount))
}

func DistributeChocolatesWithMinimumDifference(chocoloateInPacket []int,countOfStudents int) int{
	countOfPackets := len(chocoloateInPacket)
	if countOfStudents>countOfPackets {
		return -1
	}
	if (countOfPackets==0 || countOfStudents<=0){
		return -1
	}
	sort.Ints(chocoloateInPacket)
	startIndex := 0
	currentRange := 0
	maximumPossibleCombinations := countOfPackets - countOfStudents + 1
	minimumRange := chocoloateInPacket[countOfStudents-1]-chocoloateInPacket[0]
	for startIndex=1;startIndex<maximumPossibleCombinations ;startIndex=startIndex+1  {
		currentRange = chocoloateInPacket[startIndex + countOfStudents-1]-chocoloateInPacket[startIndex]
		if (currentRange<minimumRange) {
			minimumRange = currentRange
		}
	}
	return minimumRange
}
