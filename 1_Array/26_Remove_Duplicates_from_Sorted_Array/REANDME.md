# 26. Remove Duplicates from Sorted Array
Given a sorted array nums, remove the duplicates in-place such that each element appears only once and returns the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means a modification to the input array will be known to the caller as well.

Internally you can think of this:

Given a sorted array nums, remove the duplicates in-place such that each element appears only once and returns the new length.
預設~~一個數組~~一個已排序的數組nums,將重複的元素移除,並讓每個元素只出現一次,然後重新回傳新的長度

給定一個已排序的數組num，就地刪除重複項，以使每個元素僅出現一次並返回新的長度。

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
~~不允許在其他數組使用~~不要為另一個數組分額外的空間,你必須~~完成這個在只~~通過使用O（1）額外的內存來修改傳入的數組

不要為另一個數組分配額外的空間，必須通過使用O（1）額外的內存就地修改輸入數組來做到這一點。
Clarification:
~~聲明~~澄清
Confused why the returned value is an integer but your answer is an array?
疑惑為什麼回傳的值是一個整數,但你的答案是數組
Note that the input array is passed in by reference, which means a modification to the input array will be known to the caller as well.
注意傳入的數組是傳址,也就是說~~修改過的傳入的數組~~調用者也將知道對輸入數組的修改

Internally you can think of this:
~~你可以想成這樣~~
在內部，您可以想到

// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

Example 1:
```
Input: nums = [1,1,2]
Output: 2, nums = [1,2]
Explanation: Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the returned length.
```
Example 2:
```
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4]
Explanation: Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively. It doesn't matter what values are set beyond the returned length.
```
```
Constraints:

0 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums is sorted in ascending order.
```
