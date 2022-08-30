模範解答のポイント
1. dummy headを作る。自分の実装だとafter nodeをつくっていたが、それだとafter nodeを作成させるときに、条件分岐が発生する。もし繰り上げがない場合 & loopが終了していない場合。これをなくすために、dummy headを作成する。
2. 三項演算子を使う。今回の場合は値l1, l2がnullの場合valに加算しないという形だった。条件文が単純なので、三項演算子を使用しても良い気がする。

https://leetcode.com/problems/add-two-numbers/discuss/1340/A-summary-about-how-to-solve-Linked-List-problem-C%2B%2B
