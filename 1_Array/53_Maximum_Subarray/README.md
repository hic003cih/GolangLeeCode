# 53. Maximum Subarray

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
設定一個整數的數組,發現一個~~contiguous~~ 連續的 ~subarray~ 子陣列 (至少有一個數字),有一個最大的和,並返回該數


Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
~~跟隨著~~跟進:如果你有解出O(n)的解答,試著解出另一個使用分解和~~conquer~~克服 方法的解答,更加~~subtle~~ 微妙

```
Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

```
Example 2:

Input: nums = [1]
Output: 1
Example 3:

Input: nums = [0]
Output: 0
Example 4:

Input: nums = [-1]
Output: -1
Example 5:

Input: nums = [-2147483647]
Output: -2147483647
```

```
Constraints:

1 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
```
## 解法
使用動態分配法
1. 先取得數組的長度
2. 先訂一個max值,把index0的先存進去
3. 再來用迴圈,從index1開始,並取目前index前面的一個值(有可能是已經加上去的)
  3.1. 如果前面的一個值大於0,代表加上去數字會變大,有幫助,可以加上去
  3.2. 如果是負數,加上去數字越來越小,所以跳過
4. 如果大於max的值,就放到max中