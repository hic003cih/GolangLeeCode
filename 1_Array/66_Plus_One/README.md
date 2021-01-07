# 66. Plus One
Given a non-empty array of decimal digits representing a non-negative integer, increment one to the integer.
預設一個不為空的整數數組,~~裡面有非負數的整數的話~~所表示的非负整数,加一到整數中

给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.
~~數字被~~ 把~~最明顯~~最高位 的數字 存在開頭的列表中,~~還有個別的~~數組 每個元素只 包含一個數字

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
You may assume the integer does not contain any leading zero, except the number 0 itself.
你也許會~~懷疑~~ 假設 整數~~是否~~ 不會包含任何0,除了數字0本身

您可以假設整數不包含任何前導零，數字0本身除外。
 
```
Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
```
```
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
```
```
Example 3:

Input: digits = [0]
Output: [1]
```
```
Constraints:

1 <= digits.length <= 100
0 <= digits[i] <= 9
```