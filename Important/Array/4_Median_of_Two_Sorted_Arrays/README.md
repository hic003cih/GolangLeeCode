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

# 暴力法
1. 先合併array
2. 再排序合併的array
3. 總長度除2如果餘0,代表示偶數,中位數會有兩個,totalLength除出來的數和除出來的數減1
4. 總長度除2如果餘1,代表示奇數,中位數就是totalLength除以2


# 递归消除法
找出中位數,並且迴圈執行从nums1和nums2中移除任意n个元素（不移除nums1最大元素和nums2最小元素）
1. 先取出兩個nums的長度
2. 保證nums1比nums2短,如果nums2比較長,將nums2放在前面
3. 取出nums1和nums2的中位數
4. 當
