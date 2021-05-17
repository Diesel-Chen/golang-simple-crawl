# golang-simple-crawl

## 架构图演进

### 单机版爬虫

![单机版 爬虫](https://image.dieselchen.work/uPic/2021/04/09/%E5%8D%95%E6%9C%BA%E7%89%88%20%E7%88%AC%E8%99%AB.png)

### 并发版爬虫(简单版调度器)

- 把fetcher和parser封装成worker，输入request，输出requestresult.
- 新增scheduler，engine拿到request扔给scheduler。worker的输入输出都是channel
- 所有的worker共用一个输入。致命缺点！会产生循环等待

![并发版爬虫（简单版调度器）](https://image.dieselchen.work/uPic/2021/04/09/%E5%B9%B6%E5%8F%91%E7%89%88%E7%88%AC%E8%99%AB%EF%BC%88%E7%AE%80%E5%8D%95%E7%89%88%E8%B0%83%E5%BA%A6%E5%99%A8%EF%BC%89.png)

### 并发版爬虫(并发版调度器)

- engine把request送给scheduler，scheduler给每个request启动一个groutine