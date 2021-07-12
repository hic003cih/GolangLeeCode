# [35. Search Insert Positions](https://leetcode-cn.com/problems/search-insert-position/)


## 題目

Given a sorted array of distinct integers and a target value, return the index if the target is found. 
給一個已經排過序數字不相同的整數數列,和一個target,並如果target在數列中有發現,返回index
If not, return the index where it would be if it were inserted in order.
如果沒有,返回他應該出現在數列中的index位置
You must write an algorithm with O(log n) runtime complexity.


Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4
Example 4:

Input: nums = [1,3,5,6], target = 0
Output: 0
Example 5:

Input: nums = [1], target = 0
Output: 0


Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104



## 解題
   
1. 二分法
   找到中間值
   如果中間值是錯誤的版本
   代表右邊指針要繼續往左邊找最初的錯誤版本
   如果中間值不是錯誤的版本
   代表左邊指針要繼續往右邊找最初的錯誤版本
   循環檢查,直到右邊大於左邊