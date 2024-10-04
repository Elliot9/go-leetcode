# 220. Contains Duplicate III
#### tags:
- HashSet
- Bucket Sort
- Sliding Window

## 題目：
給定一個整數陣列 nums 和兩個整數 indexDiff 和 valueDiff，如果陣列中有兩個不同的索引 i 和 j
使得 i != j, i ！= j，且 abs(nums[i] - nums[j]) <= valueDiff 和 abs(i - j) <= indexDiff，則返回 true。

## 可使用的方法：
- 暴力解法
- 桶排序（Bucket Sort）和滑動窗口(Sliding Window)

## 解題思路： 
### 桶
每個桶會儲存一個固定範圍內的數字（範圍大小為 valueDiff + 1）。這樣可以確保同一個桶內的兩個數字差值一定小於或等於 valueDiff。
依據每個數字的數值將其分配到對應的桶。例如，數字 num 被分配到桶 ID 為 num / (valueDiff + 1) 的桶中，對於負數則需特殊處理。
例如：當 valueDiff = 4,  (0,1,2,3,4) 會是在 bucket(0)，而(5,6,7,8,9) 會在 bucket(1)
當位於相同桶內則符合 valueDiff 內

### 滑動窗口
為了確保索引 ∣i−j∣≤indexDiff 的條件，我們在遍歷數組時保持一個滑動窗口，該窗口只保存最近的 indexDiff + 1 個數字。
當窗口大小超過 indexDiff + 1 時，我們會移除最舊的數字，從而保證索引差距的約束。

### 鄰近桶
檢查相鄰的桶（即桶 ID ±1），這是因為跨桶邊界的數字差值也可能滿足條件

## 解法步驟：

1. 先處理 indexDiff 和 valueDiff 為不合法的情況（例如小於等於零）。
2. 桶 ID 計算：我們定義了一個函數 getBucketId(num) 來計算每個數字的桶 ID。桶的寬度為 valueDiff + 1，以確保同一個桶內的數字差值不會超過 valueDiff。
3. 遍歷數組：對於數組中的每個數字：使用 getBucketId() 函數計算該數字的桶 ID。 檢查當前桶及其相鄰的桶中是否已存在滿足條件的數字。將當前數字插入到其對應的桶中。
4. 移除已超出 indexDiff 範圍的數字（即維護滑動窗口）。

## 時間與空間複雜度：
時間複雜度 O(n) 
- 我們遍歷數組 nums 一次，每次操作（桶插入、刪除和鄰居檢查）都為常數時間。

空間複雜度 O(indexDiff)
- 我們在滑動窗口中最多保存 indexDiff + 1 個元素，因此所需的空間與 indexDiff 成正比。