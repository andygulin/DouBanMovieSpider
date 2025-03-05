### 爬取豆瓣电影/电视剧信息

#### 爬取地址

- https://movie.douban.com/subject/1292226/
- https://movie.douban.com/subject/1294638/
- https://movie.douban.com/subject/1291843/
- ......

#### 使用方法

```shell
go build
```

```shell
# 打印爬取信息
# 主页
./DouBanMovieSpider subject info {subjectId}
# 短评
./DouBanMovieSpider comment info {subjectId}
# 影评
./DouBanMovieSpider review info {subjectId}
# 图片
./DouBanMovieSpider photo info {subjectId}
```

```shell
# 输出爬取信息至文件
# 主页
./DouBanMovieSpider subject file {subjectId}
# 短评
./DouBanMovieSpider comment file {subjectId}
# 影评
./DouBanMovieSpider review file {subjectId}
# 图片
./DouBanMovieSpider photo file {subjectId}
```

```shell
# 存储信息至MongoDB
# 主页
./DouBanMovieSpider subject store {subjectId}
# 短评
./DouBanMovieSpider comment store {subjectId}
# 影评
./DouBanMovieSpider review store {subjectId}
# 图片
./DouBanMovieSpider photo store {subjectId}
```

```shell
# 从MongoDB中查询信息
# 主页
./DouBanMovieSpider subject query {subjectId}
# 短评
./DouBanMovieSpider comment query {subjectId} {pageNo} {pageSize}
# 影评
./DouBanMovieSpider review query {subjectId} {pageNo} {pageSize}
# 图片
./DouBanMovieSpider photo query {subjectId} {pageNo} {pageSize}
```

#### 配置MongoDB
- [conf.yaml](conf/conf.yaml)