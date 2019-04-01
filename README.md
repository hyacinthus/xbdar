# 雪豹商情数据分析报告(xuebao dashboard)
组织商情数据成完整的报表形式展现给用户。

## 数据库设计
![数据表结构](/files/imgs/db.png)

### 数据源(datasource)
数据域 | 类型 | 说明 | 参数
----- | --- | ---- | ---
file | json | json文件 | `{"path": "path/to/ds.json"}`
file | yaml | yaml文件 | `{"path": "path/to/ds.yaml"}`
db | sqlite3 | sqlite3数据库 | `{"dsn": "path/to/test.db"}`
db | mysql | mysql数据库 | `{"dsn": "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"}`
db | postgres | postgres数据库 | `{"dsn": "host=myhost port=myport user=gorm dbname=gorm password=mypassword"}`

### 图表(chart)
每个图表数据来自一个特定数据源，由**数据参数**指定获取数据的方式，**图表参数**描述具体的绘图方式。

#### **数据参数**(目前)
对于文件数据源， 使用`key_path`指定数据层次，如果为空，相当于取整个文件数据；当文件内容是一个map形式时，`key_path`以点分隔每一层的字段。

``` javascript
// file.{json,yaml}
{"key_path": "a.b"}

```
对于数据库数据源，使用`sql`来进行查询。
``` javascript
// db.{sqlite3,mysql,postgres}
{"sql": "select key, val from xx_key_val"}
```

#### **图表参数**(目前)
在没有和前端协商更细致的图表样式抽象格式之前，图表参数目前只有一个`plot_js`，代表一段js代码，eval执行之后返回一个function，带两个参数mountNode和data，如下为基础柱状图:
``` javascript
{"plot_js":"(function(node, data) { var chart = new G2.Chart({ container: node, forceFit: true }); chart.source(data); chart.scale('sales', { tickInterval: 20 }); chart.interval().position('year*sales'); chart.render(); }) "}
```

### 报表(dashboard)
每一个报表可以有0个或多个子报表(父ID关联，order排序)，以此组织报表，可以对报表划分章节。key可以对报表分组，比如：同一类报表的不同期应该是相同的key。如何以后有对周期报表的更多需求，再重构。

布局json(layout_json)字段组织报表内容的排版，形式为一个组件树。

#### 报表布局
组件属性: id, type, meta, children
组件|name|说明
---|----|----
标题|Header|
分隔线|Divider|
行容器|Row|
列容器|Col|
图表|Chart|各种图表
文本|Text|支持各种格式:markdown等
标签页|Tabs,Tab|使用标签切换内容
TODO: 各个组件meta属性和组合规则


## API
TODO: 使用Swagger查看和测试接口
### 展示接口
### Dashboard
* 分页查询报表信息
* 查询单个报表信息（递归包括子报表）

### Chart
* 分页查询图表信息
* 查询单个图表信息
* 获取图表数据

### 管理接口
管理后台和其它应用对接使用的接口。
TODO
