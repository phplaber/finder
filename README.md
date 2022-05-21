**Finder** 是一款收集企业内部 Mac 平台安装应用的工具，用于软件版权合规管理。

## 功能
1. 收集安装应用，数据源包括 `/Applications` 和 `~/Applications`
2. 收集计算机名（`hostname`）
3. 收集账号（`whoami`）

## 使用
1. 下载源代码
``` bash
git clone https://github.com/phplaber/finder.git
```
2. 创建 MySQL 库、表和字段
* 创建数据库
* 执行 `db.sql` 中 SQL 语句，创建表和字段
3. 将 `configs/db.go.sample` 重命名为 `configs/db.go`，并配置 `DataSourceName`
4. 编译程序，生成可执行文件
``` bash
env GOOS=darwin GOARCH=amd64 go build -v ./cmd/finder/
```
5. 赋予可执行文件执行权限，然后执行
``` bash
chmod +x ./finder
./finder
```