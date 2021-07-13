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
   
1. 先排除長度0的情況下,就不會有前綴,直接回傳空
2. 預設array第一個index的string為前綴
3. 計算array長度
4. 寫一個比對prefix的func,比對兩個傳入的字串長度,取短的字串來當prefix
5. 建一個迴圈比對,如果兩個字串的index的字母相同,然後index < length,index加1,繼續比對,直到兩個字串的index的字母不相同
6. 最後跳出迴圈,回傳迴圈比對以後index以前的字母return str1[:index]
7. 主func對array進行迴圈,從第二個開始,將prefix和array中第i個傳入比對prefix的程式
8. 如果比對prefix的程式回傳的prefix長度為0,表示已經沒有prefix,主程式跳出迴圈
