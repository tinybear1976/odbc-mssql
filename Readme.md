---
module: github.com/tinybear1976/odbc-mssql
function: 简单封装mssqldb连接的基本内容
version: 0.1.0
path: github.com/tinybear1976/odbc-mssql
---

目录

[TOC]



# 引用

## go.mod

replace 的物理磁盘位置要根据物理目录的实际位置给定。其中  github.com/jmoiron/sqlx  可以在引用github.com/tinybear1976/odbc-mssql后，使用go mod  tidy，让系统自动添加

```go
module test

go 1.15

require (
	github.com/jmoiron/sqlx v1.2.0 // indirect
	hhyt/database/mssqldb v0.1.0
)

replace hhyt/database/mssqldb => ../../hhyt/database/mssqldb@v0.1.0
```

## main.go

连接和基本操作。由例可见，数据库连接初始化可以单独进行，并将其连接保存至包内，直到使用Destroy()，手动销毁所有的连接指针。在进行后续的操作时，只需要使用连接标识来表明调用哪个OracleDB连接即可。

```go
package main

import (
	"fmt"
	"github.com/tinybear1976/odbc-mssql"
)

func main() {

	dbInit()
	DbOptTest()
	mssqldb.Destroy()
}

func dbInit() {
	err := odbcmssqldb.New("local", "172.16.1.3", "sa", "123", "test", 1433)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DbOptTest() {
	db, err := odbcmssqldb.Connect("local")
	if err != nil {
		fmt.Println(err)
		return
	}
	var vals []string
	db.Select(&vals, "select username from users")
	fmt.Println(len(vals), vals)
}

```



# 函数

## New

建立并保存一个以传统odbc方式访问的mssqldb连接，在之后的使用过程中，通过连接标识创建连接并进行数据库操作。在New的过程中模块会尝试连接数据库，如果发现服务器无法连接，则不会保存该数据库连接。

```go
func New(serverTag, server, userid, password, db string, port int) error
```

入口参数：
| 参数名    | 类型   | 描述                     |
| --------- | ------ | ------------------------ |
| serverTag | string | OracleDB数据库连接标识   |
| server    | string | SqlServer数据库服务器地址 |
| userid    | string | 数据库用户名             |
| password  | string | 数据库用户密码           |
| db        | string | 数据库名称               |
| port      | int    | 数据库服务端口           |

返回值：正确返回nil，否则返回错误信息



## Destroy

销毁所有模块内保存的（由New创建的）连接池指针。

```go
func Destroy()
```

入口参数：无

返回值：无



## Connect

获得一个数据库连接。

```go
func Connect(serverTag string) (*sqlx.DB, error) 
```

入口参数：

| 参数名    | 类型   | 描述                  |
| --------- | ------ | --------------------- |
| serverTag | string | mssqldb数据库连接标识 |

返回值：

| 返回变量 | 类型     | 描述                                      |
| -------- | -------- | ----------------------------------------- |
|          | *sqlx.DB | 数据库连接指针                            |
|          | error    | 返回操作结果的错误信息，如果正确则返回nil |

