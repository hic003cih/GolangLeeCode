# [329. Spiral Matrix](https://leetcode-cn.com/problems/spiral-matrix/)


## 題目

Given an m x n matrix, return all elements of the matrix in spiral order.
給一個m乘n的矩陣.返回所有矩陣元素使用
 

Example 1:

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

```

Example 2:

```
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

```

Constraints:
```
2 <= nums.length <= 103
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
```

## 解題
1. 暴力法
用雙層for把nums拆解,第一層用來固定第一個index
第二層用來執行後面的每個index,然後在最裡面的迴圈加起來,並返回index
   
2. 哈希表
用一個哈希表來將數字存入表中,然後用target-x去尋找這個數值是否等於target的值
