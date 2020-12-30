# 35. Search Insert Position
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

給一個排序過包含不同的整數的數組和一個target的value變數,如果target的值在其中,返回其index.如果不在其中,~~找到他應該在的index~~返回他將會按照順序插入的index值


```
Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4
Example 4:

Input: nums = [1,3,5,6], target = 0
Output: 0
Example 5:

Input: nums = [1], target = 0
Output: 0
```
```
Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104
```

## 解題
可以用2分法來拆解

1. 先找出中位數(最前index加長度),然後除已2
2. 各訂一個left和right point,如果target值小於中位數,看左邊那段,把right指針移到中位數的位置
3. 如果target值大於中位數,看又邊那段,把right指針移到(中位數+1)的位置
4. 然後持續比大小,比到left和right指針一致就返回