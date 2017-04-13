## [基于时间和用户投票赞同和反馈排名](http://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_reddit.html)

## Reddit 使用此排名
## 用户可以投赞成票和反对票

```
/**
 * y的作用是产生加分或减分。保证了得到大量净赞成票的文章，会排在前列；赞成票与反对票接近或相等的文章，会排在后面；
 * 得到净反对票的文章，会排在最后（因为得分是负值）。 如果要是考虑极限情况，可以将y取-1的范围缩小一点。
 * <p>
 * t越大，得分越高，即新帖子的得分会高于老帖子。它起到自动将老帖子的排名往下拉的作用。
 * 分母的45000秒，等于12.5个小时，也就是说，后一天的帖子会比前一天的帖子多得2分
 * <p>
 * 这种算法的一个问题是，对于那些有争议的文章（赞成票和反对票非常接近），它们不可能排到前列
 * <p>
 * Reddit的排名，基本上由发帖时间决定，超级受欢迎的文章会排在最前面，一般性受欢迎的文章、有争议的文章都不会很靠前
 *
 * @return
 */
private static double getScore(int vote, int oppose, Date createTime) {
    int x = vote - oppose; // 赞成票与反对票的差x
    int y = 0; // 投票方向y

    if (x > 0) {
        y = 1;
    } else if (x < 0) {
        y = -1;
    } else {
        y = 0;
    }

    int z = 1; // 帖子的受肯定（否定）的程度z
    if (x != 0) {
        z = Math.abs(x);
    }

    // 帖子的新旧程度t，单位为秒
    Date date = DateUtil.parse("2017-04-01 00:00:00"); // 公司成立时间
    long t = (createTime.getTime() - date.getTime()) / 1000l;

    return Math.log(z) + y * t / 45000;
}
```