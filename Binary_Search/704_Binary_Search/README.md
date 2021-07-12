# [278. First Bad Version](https://leetcode-cn.com/problems/first-bad-version/)


## 題目

You are a product manager and currently leading a team to develop a new product.
你是一個產品經理,並正在領導一個團隊開發新產品
Unfortunately, the latest version of your product fails the quality check. 
但不幸的,最新版本的品質測試失敗了
Since each version is developed based on the previous version, all the versions after a bad version are also bad.
因為每個版本都是開發在前一版本的基礎上,所有在壞版本以後的都是壞的
Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.
你有n個版本,然後你想要找出壞掉造成後面版本損壞的版本
You are given an API bool isBadVersion(version) which returns whether version is bad. 
你給予一個API bool isBadVersion,來回傳這個版本是壞的
Implement a function to find the first bad version. 
執行一個function來找到最開始壞掉的本版
You should minimize the number of calls to the API.
你應該使用最小呼叫API的次數

Example 1:

Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5)-> true
call isBadVersion(4)-> true
Then 4 is the first bad version.
Example 2:

Input: n = 1, bad = 1
Output: 1


Constraints:

1 <= bad <= n <= 231 - 1



## 解題
   
1. 二分法
   找到中間值
   如果中間值是錯誤的版本
   代表右邊指針要繼續往左邊找最初的錯誤版本
   如果中間值不是錯誤的版本
   代表左邊指針要繼續往右邊找最初的錯誤版本
   循環檢查,直到右邊大於左邊