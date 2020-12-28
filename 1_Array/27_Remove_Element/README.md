# 27. Remove Element
Given an array nums and a value val, remove all instances of that value in-place and return the new length.

設定一個nums的數組和一個名為val的value,~~移除所有的在value中的實例~~需要移除所有數值等於val的元素,並返回新的長度


Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

不要給予額外的空間給其他的數組,你必須在使用O(1)的額外記憶體的情況下~~完成這個~~修改輸入的數組

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

元素的排序可以改變,~~這個不影響你留新的長度在後面~~你不需要考慮數組中後面的元素超出新的長度

Clarification:

澄清

Confused why the returned value is an integer but your answer is an array?

困惑返回的值為什麼是一個整數,但你的答案是數組嗎?

Note that the input array is passed in by reference, which means a modification to the input array will be known to the caller as well.

因為傳入的數組是以傳址,這個代表~~一個對傳入的數組的修改最好通知呼叫者~~調用者將知道對輸入數組的修改

Internally you can think of this:
在內部，您可以想到：

// nums is passed in by reference. (i.e., without making a copy)
int len = removeElement(nums, val);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
 
```
Example 1:

Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2]
Explanation: Your function should return length = 2, with the first two elements of nums being 2.
It doesn't matter what you leave beyond the returned length. For example if you return 2 with nums = [2,2,3,3] or nums = [2,2,0,0], your answer will be accepted.
```
```
Example 2:

Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3]
Explanation: Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4. Note that the order of those five elements can be arbitrary. It doesn't matter what values are set beyond the returned length.
 ```
```
Constraints:

0 <= nums.length <= 100
0 <= nums[i] <= 50
0 <= val <= 100
```

## 解題想法


使用雙指針法,分別設定一個right和一個left

如果right的值不等於val,則把right的值寫到left指針的位置,並把left指針加一向右移

相等得則跳過,left指針位置不動

最後返回left指針的位置,left指針的位置表示不同於val的值寫入了幾次
