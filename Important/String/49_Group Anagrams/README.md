# [49. Group Anagrams](https://leetcode-cn.com/problems/group-anagrams/)


## 題目

Given an array of strings strs, group the anagrams together. 
給於一個 字串ARRAY strs,
You can return the answer in any order.
你可以返回任何順序的答案
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, 
一個Anagram是一個字母,也可能是是一個單字,
typically using all the original letters exactly once.
更可能是使用所有原始的字母

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母都恰好只用一次。


 

Example 1:

```Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

```

Example 2:

```
Input: strs = [""]
Output: [[""]]
```
Example 3:
```
Input: strs = ["a"]
Output: [["a"]]
```

Constraints:
```
1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lower-case English letters.
```

## 解題
1. 暴力法
用雙層for把nums拆解,第一層用來固定第一個index
第二層用來執行後面的每個index,然後在最裡面的迴圈加起來,並返回index
   
2. 哈希表
用一個哈希表來將數字存入表中,然後用target-x去尋找這個數值是否等於target的值
