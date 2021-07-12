# [704. Binary Search](https://leetcode-cn.com/problems/binary-search/)


## 題目

Given an array of integers nums which is sorted in ascending order, and an integer target,
給予一個降冪排序的整數數組nums和一個整數目標
write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
寫一個func去搜尋target數字是否在nums,如果存在,返回index,如果沒有返回-1
You must write an algorithm with O(log n) runtime complexity.


Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4
Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1


Constraints:

1 <= nums.length <= 104
-104 < nums[i], target < 104
All the integers in nums are unique.
nums is sorted in ascending order.


## 解題
   
1. 二分法
先找出中間值的index,然後比較target的值和中間值哪個大
   如果target大,就比對右半邊的
   如果target小,就比對左半邊的
   這樣就可以少比對一半的值
