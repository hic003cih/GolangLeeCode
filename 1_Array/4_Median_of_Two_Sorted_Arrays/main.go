package main

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 7}
	findMedianSortedArrays(nums1, nums2)
}

func findMedianSortedArrays(mis []int, njs []int) float64 {

	lenMis, i := len(mis), 0
	lenNjs, j := len(njs), 0
	res := make([]int, lenMis+lenNjs)

	for k := 0; k < lenMis+lenNjs; k++ {
		//如果i的長度和nums長度一樣
		//或者i<
		if i == lenMis ||
			(i < lenMis && j < lenNjs && mis[i] > njs[j]) {
			res[k] = njs[j]
			j++
			continue
		}

		if j == lenNjs ||
			(i < lenMis && j < lenNjs && mis[i] <= njs[j]) {
			res[k] = mis[i]
			i++
		}
	}

	return 0.0
	// nums1 = append(nums1, nums2...)
	// sum := 0
	// newArray := make([]int, len(nums1))
	// newArray2 := make([]int, len(nums1))
	// //fmt.Println(newArray)
	// //fmt.Println(newArray[len(newArray)-1 : len(newArray)])
	// for i, v := range nums1 {

	// 	/* if i == lenMis ||
	// 		(i < lenMis && j < lenNjs && mis[i] > njs[j]) {
	// 		res[k] = njs[j]
	// 		j++
	// 		continue
	// 	} */

	// 	// if i == 0 {
	// 	// 	newArray = append(newArray, v)
	// 	// }
	// 	// for y, v2 := range newArray {
	// 	// 	fmt.Println(v2)
	// 	// 	if v > v2 {
	// 	// 		//fmt.Println(newArray[:y])
	// 	// 		newArray2 = append(newArray[:y], v)
	// 	// 		fmt.Println(newArray2)
	// 	// 	}
	// 	// }
	// 	// if v > newArray[len(newArray)-1:len(newArray)] {
	// 	// 	newArray = append(newArray, v)
	// 	// }
	// 	sum += v
	// }
	// //fmt.Println(nums1)
	// //fmt.Println(sum)
	// //fmt.Println(len(nums1))
	// //fmt.Println(float64(sum) / float64(len(nums1)))
	// return float64(sum) / float64(len(nums1))

}
