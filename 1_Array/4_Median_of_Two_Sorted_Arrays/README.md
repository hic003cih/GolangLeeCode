# 題目
## 4. Median of Two Sorted Arrays
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

Follow up: The overall run time complexity should be O(log (m+n)).


Given two sorted arrays nums1 and nums2 of size m and n respectively(個別地), return the median(中位數) of the two sorted arrays.
給予兩個排列過的數列 nums1和nums2,個別的大小為m和n,並將兩個排序過的數列結合,返回平均數
Follow up: The overall(整體) run time complexity(複雜性) should be O(log (m+n)).
整體的複雜性應該要是O(log (m+n))
 


```
Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
```
```
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
```
```
Example 3:

Input: nums1 = [0,0], nums2 = [0,0]
Output: 0.00000
```
```
Example 4:

Input: nums1 = [], nums2 = [1]
Output: 1.00000
```
```
Example 5:

Input: nums1 = [2], nums2 = []
Output: 2.00000
 ```
```
Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
```

## 解題

先計算出num合併以後長度,如果不是複數,就表示中位數就只有一個數字,就不用算

先合併

在用for回圈和左右指針的方式比較每一筆的大小(有點像bublesort,一個一個比大小)
如果左指針比較大,將左指針的數移到右指針的位置,然後右指針的數移到左指針的位置,左右指針各加一,比到右指針到最大的len為止
