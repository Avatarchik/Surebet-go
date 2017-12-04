package dataBase

import "surebetSearch/dataBase/types"

func getUnique(collectedPairs, newPairs []types.EventPair) []types.EventPair {
	var uniquePairs []types.EventPair
loop:
	for _, newPair := range newPairs {
		for _, collectedPair := range collectedPairs {
			if collectedPair == newPair {
				continue loop
			}
		}
		uniquePairs = append(uniquePairs, newPair)
	}
	return uniquePairs
}

func UniqueAndDiff(collectedPairs *[]types.EventPair, newPairs []types.EventPair) int {
	//Check if slice is sorted, then sort if necessary
	uniquePairs := getUnique(*collectedPairs, newPairs)
	*collectedPairs = append(*collectedPairs, uniquePairs...)
	//Sort after appending
	return len(uniquePairs)
}
