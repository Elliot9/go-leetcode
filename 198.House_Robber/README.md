# 198. House Robber
# 198. 打家劫舍

#### Tags / 標籤：
- Dynamic Programming / 動態規劃
- Array / 數組

## Problem / 題目：
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

你是一個專業的小偷，計劃偷竊沿街的房屋。每間房內都藏有一定的現金，影響你偷竊的唯一制約因素就是相鄰的房屋裝有相互連通的防盜系統，如果兩間相鄰的房屋在同一晚上被小偷闖入，系統會自動報警。

給定一個代表每個房屋存放金額的非負整數數組，計算你 不觸動警報裝置的情況下 ，一夜之內能夠偷竊到的最高金額。

## Approach / 解題方法：
- Dynamic Programming / 動態規劃

## Solution / 解題思路： 
1. Create a dynamic programming array to store the maximum amount that can be robbed up to each house.
2. Initialize the first two elements of the DP array:
   - dp[0] = nums[0] (rob the first house)
   - dp[1] = max(nums[0], nums[1]) (rob either the first or second house)
3. For each subsequent house i, calculate:
   - dp[i] = max(dp[i-1], dp[i-2] + nums[i])
   - This means either skip the current house and take the previous max, or rob the current house and add it to the max two houses before.
4. The last element of the DP array will contain the maximum amount that can be robbed.

1. 創建一個動態規劃數組，用於存儲到每個房屋為止可以搶劫的最大金額。
2. 初始化 DP 數組的前兩個元素：
   - dp[0] = nums[0]（搶劫第一間房子）
   - dp[1] = max(nums[0], nums[1])（搶劫第一間或第二間房子中的較大值）
3. 對於每個後續的房子 i，計算：
   - dp[i] = max(dp[i-1], dp[i-2] + nums[i])
   - 這意味著要麼跳過當前房子並取前一個最大值，要麼搶劫當前房子並將其加到兩個房子之前的最大值上。
4. DP 數組的最後一個元素將包含可以搶劫的最大金額。

Time Complexity: O(n), where n is the number of houses.
Space Complexity: O(1), as we are modifying the input array in-place.

時間複雜度：O(n)，其中 n 是房子的數量。
空間複雜度：O(1)，因為我們直接在輸入數組上進行修改。
