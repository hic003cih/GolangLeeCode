# [1. Two Sum](https://leetcode.com/problems/two-sum/)


## 題目

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

 

Example 1:

```Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

Example 2:

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```
Example 3:
```
Input: nums = [3,3], target = 6
Output: [0,1]
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