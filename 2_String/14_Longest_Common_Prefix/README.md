# [14. Longest Common Prefix](https://leetcode-cn.com/problems/longest-common-prefix/)


## 題目

Write a function to find the longest common prefix string amongst an array of strings.
寫一個function去查到最長的common prefix

編寫一個函數來查找字符串數組中最長的公共前綴字符串。

If there is no common prefix, return an empty string "".
如果沒有公共前綴，則返回空字符串“”。

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
 

Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lower-case English letters.

## 解題
   
1. 循環去比對array的每個
