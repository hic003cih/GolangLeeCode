# [90. Subsets II
](https://leetcode-cn.com/problems/subsets-ii/)


## 題目

Given an integer array nums that may contain duplicates, return all possible subsets (the power set).
給一個整數數組nums,數組中可能包含重複的數
返回所有可能的子集(降幕)

The solution set must not contain duplicate subsets. 
解答不能包含重複的子集
Return the solution in any order.
解答可以任意順序

 

Example 1:

Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
Example 2:

Input: nums = [0]
Output: [[],[0]]
 

Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10


## 解題
迭代法实现子集枚举
假如是[]int{1, 2, 3} 
這個的結果會有2的N次
N就是這個array有的數字數
會有[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]共8組子集答案
我們可以假設[]int{1, 2, 3} ,有000,100,010....這麼多組不重複的子集組合
[] -> 000
[1] -> 100
[2] -> 010
[1,2] -> 110
[3] -> 001
[1,3] -> 101
[2,3] -> 011
[1,2,3] -> 111

1. 先算出array的長度
2. 建立一個dfs func ,重複呼叫自己本身,一直往下找有可能的子集(ex: 1往下找1,2,在往下找1,2,3)
3. 當dfs func循環的長度等於array的長度時(當是1,2,3時),把子集存到答案中,跳出當前循環,並將子集中的尾數去除(當是1,2,3時,去除3,變回1,2)
4. 去除以後繼續進入dfs func循環中,因為去除尾數就代表回到上一層(從1,2開始),繼續找子集





