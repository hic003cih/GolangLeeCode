# 88. Merge Sorted Array
合併排序過的數組
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
預設兩個排序過的整數數組nums1和nums2,合併nums2到nums1中,成為一個排序過的數組
The number of elements initialized in nums1 and nums2 are m and n respectively. 
nums1和nums2初始的元素數量分別為m和n
You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
你可以假設nums1有足夠的容量(大小為m+n)來容納來自nums2的額外元素

```
Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
```
```
Example 2:

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
```

## 解法

1. 用三指針法 ,一個設要總長度的index,一個設對應num1的index,一個設對應num2的index,指針都要減1的長度,要不然會超過index

2. 從後往前比,如果從前往後比的會,會發生多次後移的情況,時間複雜度較高

3. 先比num1和num2最後一個數的大小,把比較大的寫入到總長度的index位置,然後看是哪個數組寫入的,那個數組的指針和總長度指針都減一,往左邊移一位

4. 如果num1或num2其中一個index長度沒了,則改成剩下的數組和總長度的數組比,比到剩下的數組沒有長度為止