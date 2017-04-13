
## [基于时间和用户投票排名算法](http://www.ruanyifeng.com/blog/2012/02/ranking_algorithm_hacker_news.html)

### Hacker News 采用此排名
### 用户只能投赞成票
```
/**
 * P表示帖子的得票数，减去1是为了忽略发帖人的投票。
 * 如果你不想让"高票帖子"与"低票帖子"的差距过大，可以在得票数上加一个小于1的指数，比如(P-1)^0.8
 * <p>
 * T表示距离发帖的时间（单位为小时），加上2是为了防止最新的帖子导致分母过小。
 * <p>
 * G表示"重力因子"（gravityth power），即将帖子排名往下拉的力量，默认值为1.8。
 * G值越大，曲线越陡峭，排名下降得越快，意味着排行榜的更新速度越快。
 *
 * @param createTime 帖子创建时间
 * @return
 */
private static double getScore(int p, Date createTime) {
    double g = 1.8;

    long t = (new Date().getTime() - createTime.getTime()) / (1000l * 60 * 60);

    return (p - 1) / Math.pow(t + 2, g);
}
```