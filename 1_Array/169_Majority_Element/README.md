# 169. Majority Element
Given an array nums of size n, return the majority element.
設定一個空間為n的數組,返回~~主要的~~多數 元素
給定一個大小為n的數組num，返回多數元素。
The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
~~主要的~~ 多數 元素的定義是出現超過⌊n / 2⌋次,你可以假設majority元素都出現在數組中


Example 1:

Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2
 

Constraints:

n == nums.length
1 <= n <= 5 * 104
-231 <= nums[i] <= 231 - 1