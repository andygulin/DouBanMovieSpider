### 爬取豆瓣电影/电视剧信息

#### 爬取地址

- https://movie.douban.com/subject/1292226/
- https://movie.douban.com/subject/1294638/
- https://movie.douban.com/subject/1291843/
- ......

#### 使用方法

```shell
go build -o bin/Spider spider/main.go

# Top250
go build -o bin/Top250 top250/main.go
```

```shell
# 打印爬取信息
# 主页
./Spider subject info {subjectId}
# 短评
./Spider comment info {subjectId}
# 影评
./Spider review info {subjectId}
# 图片
./Spider photo info {subjectId}
```

```shell
# 输出爬取信息至文件
# 主页
./Spider subject file {subjectId}
# 短评
./Spider comment file {subjectId}
# 影评
./Spider review file {subjectId}
# 图片
./Spider photo file {subjectId}
```

```shell
# 存储信息至MongoDB
# 主页
./Spider subject store {subjectId}
# 短评
./Spider comment store {subjectId}
# 影评
./Spider review store {subjectId}
# 图片
./Spider photo store {subjectId}
```

```shell
# 从MongoDB中查询信息
# 主页
./Spider subject query {subjectId}
# 短评
./Spider comment query {subjectId} {pageNo} {pageSize}
# 影评
./Spider review query {subjectId} {pageNo} {pageSize}
# 图片
./Spider photo query {subjectId} {pageNo} {pageSize}
```

#### 配置MongoDB
- [conf.yaml](conf/conf.yaml)