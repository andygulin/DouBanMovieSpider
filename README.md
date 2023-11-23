### 爬取豆瓣电影/电视剧信息

#### 爬取地址

- https://movie.douban.com/subject/1292226/
- https://movie.douban.com/subject/1294638/
- https://movie.douban.com/subject/1291843/
- ......

#### 使用方法

```shell
go build

# 打印爬取信息-主页
./DouBanMovieSpider subject info {subjectId}
# 打印爬取信息-短评
./DouBanMovieSpider comment info {subjectId}
# 打印爬取信息-影评
./DouBanMovieSpider review info {subjectId}
# 打印爬取信息-图片
./DouBanMovieSpider photo info {subjectId}
......

# 输出爬取信息至文件-主页
./DouBanMovieSpider subject file {subjectId}
# 输出爬取信息至文件-短评
./DouBanMovieSpider comment file {subjectId}
# 输出爬取信息至文件-影评
./DouBanMovieSpider review file {subjectId}
# 输出爬取信息至文件-图片
./DouBanMovieSpider photo file {subjectId}
......

# 存储信息至MongoDB-主页
./DouBanMovieSpider subject store {subjectId}
# 存储信息至MongoDB-短评
./DouBanMovieSpider comment store {subjectId}
# 存储信息至MongoDB-影评
./DouBanMovieSpider review store {subjectId}
# 存储信息至MongoDB-图片
./DouBanMovieSpider photo store {subjectId}
......
```

#### 配置MongoDB

- store/mongo.go