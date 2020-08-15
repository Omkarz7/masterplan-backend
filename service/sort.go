package service

import (
	"masterplan-backend/models"
	"strconv"
	"strings"
)

//WBSMergeSort sorts masterplan according to Work Breakdown Structure
func WBSMergeSort(items []models.Masterplan) []models.Masterplan {
	num := len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	left := make([]models.Masterplan, middle)
	right := make([]models.Masterplan, num-middle)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return wbsMerge(WBSMergeSort(left), WBSMergeSort(right))
}

func wbsMerge(left, right []models.Masterplan) (result []models.Masterplan) {
	result = make([]models.Masterplan, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if compareSlNo(left[0].SlNo, right[0].SlNo) { //left[0].SlNo < right[0].SlNo
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

func compareSlNo(left, right string) (lesser bool) {
	var strslc []string
	left = strings.TrimSpace(left)
	right = strings.TrimSpace(right)

	strslc = strings.Split(left, ".")
	leftslc := make([]int, len(strslc))
	for i := range strslc {
		//not checking for err in strconv.Atoi under the assumption that
		//data as verified during insertion
		leftslc[i], _ = strconv.Atoi(strings.TrimSpace(strslc[i]))
	}

	strslc = strings.Split(right, ".")
	rightslc := make([]int, len(strslc))
	for i := range strslc {
		//not checking for err in strconv.Atoi under the assumption that
		//data as verified during insertion
		rightslc[i], _ = strconv.Atoi(strings.TrimSpace(strslc[i]))
	}

	//If an Sl No is the child of another then the one with longer length is the child
	//ex. Between 1.2.1 and 1.2, 1.2.1 is the child so it will come after 1.2 in WSB sort
	lesser = len(leftslc) < len(rightslc)

	//iterate over the length of the shortest slice
	//If at values are not equal, find if left<right
	//if it's always equal then we know that the one with
	// biggerlength is child of the other according to the above equation
	var i int
	for i = 0; i < len(leftslc) && i < len(rightslc); i++ {
		if leftslc[i] != rightslc[i] {
			lesser = (leftslc[i] < rightslc[i])
			break
		}
	}

	return lesser
}

//StartDateMergeSort sorts masterplan according to Start Date followed by Work Breakdown Structure
func StartDateMergeSort(items []models.Masterplan) []models.Masterplan {
	num := len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	left := make([]models.Masterplan, middle)
	right := make([]models.Masterplan, num-middle)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return startDateMerge(StartDateMergeSort(left), StartDateMergeSort(right))
}

func startDateMerge(left, right []models.Masterplan) (result []models.Masterplan) {
	result = make([]models.Masterplan, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0].StartDate.Equal(right[0].StartDate) { //left[0].StartDate = right[0].StartDate
			result[i] = left[0]
			if compareSlNo(left[0].SlNo, right[0].SlNo) { //left[0].SlNo < right[0].SlNo
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if left[0].StartDate.Before(right[0].StartDate) { //left[0].StartDate < right[0].StartDate
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
