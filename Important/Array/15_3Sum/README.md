# [15. 3Sum](https://leetcode-cn.com/problems/3sum/)


## 題目

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, 
and nums[i] + nums[j] + nums[k] == 0.
給一個整數array nums,返回三個數 i,j,k 相加以後等於0的值

Notice that the solution set must not contain duplicate triplets.
不能包含重複的值

 

Example 1:

```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
```

Example 2:

```
Input: nums = []
Output: []
```
Example 3:
```
Input: nums = [0]
Output: []
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
