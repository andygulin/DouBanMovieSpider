### 爬取豆瓣电影/电视剧信息

#### 爬取地址

- https://movie.douban.com/subject/1292226/
- https://movie.douban.com/subject/1294638/
- https://movie.douban.com/subject/1291843/
- ......

#### 使用方法

```shell
go build

# 打印爬取信息
./DouBanMovieSpider info {subjectId}
./DouBanMovieSpider info 1292226
./DouBanMovieSpider info 1294638
......

# 存储信息至MongoDB
./DouBanMovieSpider store {subjectId}
./DouBanMovieSpider store 1292226
./DouBanMovieSpider store 1294638
......
```

#### 配置MongoDB

- store/mongo.go