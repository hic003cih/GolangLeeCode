# [77. Combinations
](https://leetcode-cn.com/problems/combinations/)


## 題目

Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].
給兩個整數(n,k),返回K所有可能的範圍[1,n]的組合

给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

You may return the answer in any order.
你可以返回答案以任何順序的方式
 

Example 1:

Input: n = 4, k = 2
返回 1 ... 4 中所有可能的 2 个数的组合。
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
Example 2:

Input: n = 1, k = 1
返回 1 ... 1 中所有可能的 1 个数的组合。
Output: [[1]]
 

Constraints:

1 <= n <= 20
1 <= k <= n



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





