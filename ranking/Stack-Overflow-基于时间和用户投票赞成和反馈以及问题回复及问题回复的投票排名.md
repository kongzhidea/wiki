

### [基于时间和用户投票赞成和反馈以及问题回复及问题回复的投票排名](http://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_stack_overflow.html)

#### [what-formula-should-be-used-to-determine-hot-questions](https://meta.stackexchange.com/questions/11602/what-formula-should-be-used-to-determine-hot-questions)
## Stack-Overflow 采用此排名

```
(log(Qviews)*4) + ((Qanswers * Qscore)/5) + sum(Ascores)
--------------------------------------------------------
((QageInHours+1) - ((QageInHours - Qupdated)/2)) ^ 1.5
```
```
（1）Qviews（问题的浏览次数）
　　
某个问题的浏览次数越多，就代表越受关注，得分也就越高。
这里使用了以10为底的对数，用意是当访问量越来越大，它对得分的影响将不断变小。

（2）Qscore（问题得分）和Qanswers（回答的数量）
首先，Qscore（问题得分）= 赞成票-反对票。
如果某个问题越受到好评，排名自然应该越靠前。
Qanswers表示回答的数量，代表有多少人参与这个问题。这个值越大，得分将成倍放大。
这里需要注意的是，如果无人回答，Qanswers就等于0，这时Qscore再高也没用，
意味着再好的问题，也必须有人回答，否则进不了热点问题排行榜。

（3）Ascores（回答得分）
一般来说，"回答"比"问题"更有意义。这一项的得分越高，就代表回答的质量越高。

（4）Qage（距离问题发表的时间）和Qupdated（距离最后一个回答的时间）
Qage和Qupdated的单位都是秒。如果一个问题的存在时间越久，或者距离上一次回答的时间越久，Qage和Qupdated的值就相应增大。
也就是说，随着时间流逝，这两个值都会越变越大，导致分母增大，因此总得分会越来越小。

（５）总结
Stack Overflow热点问题的排名，与参与度（Qviews和Qanswers）和质量（Qscore和Ascores）成正比，
与时间（Qage和Qupdated）成反比。
```