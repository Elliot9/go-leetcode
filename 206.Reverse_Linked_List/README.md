# 206. Reverse Linked List 反向鏈
#### tags:
- link

## 題目描述
給定 “單一” 鏈結, 返回反向鏈結

## 可使用的方法：
- link

## 解題思路：
1. 建立新鏈結
2. 保存當前節點的 Next
3. 將當前節點 Next 指向新鏈結
4. 將鏈結移動到 目前節點
5. 將目前節點 移動到 2. 保存的原先 Next

#### 說明：
prev = nil
head = 1 -> 2 -> 3 -> 4 -> 5 -> nil

第一次迭代後：
prev = 1 -> nil
head = 2 -> 3 -> 4 -> 5 -> nil

第二次迭代後：
prev = 2 -> 1 -> nil
head = 3 -> 4 -> 5 -> nil

第三次迭代後：
prev = 3 -> 2 -> 1 -> nil
head = 4 -> 5 -> nil

第四次迭代後：
prev = 4 -> 3 -> 2 -> 1 -> nil
head = 5 -> nil

最後一次迭代後：
prev = 5 -> 4 -> 3 -> 2 -> 1 -> nil
head = nil
