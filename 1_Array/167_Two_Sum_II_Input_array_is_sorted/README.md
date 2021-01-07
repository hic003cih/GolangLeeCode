# 167. Two Sum II - Input array is sorted
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
給定一個已經升冪排序好的數組整數,找到兩個數相加以後具體(specific)等於target的數量
The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.
這個函數twoSum應該返回兩個數的指標(indices)位置,index1的位置必須小於index2

Note:

Your returned answers (both index1 and index2) are not zero-based.
你返回的答案不能等於0
You may assume that each input would have exactly one solution and you may not use the same element twice.
你可以假設每個輸入都有一個解答,你不能使用相同的元素兩次
```
Example 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
```
```
Example 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
```
```
Example 3:

Input: numbers = [-1,0], target = -1
Output: [1,2]
 ```

Constraints:

2 <= nums.length <= 3 * 104
-1000 <= nums[i] <= 1000
nums is sorted in increasing order.
-1000 <= target <= 1000