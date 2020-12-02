# Geoos
我们的组织`spatial-go`正式成立，这是我们的第一个开源项目`Geoos`,`Geoos`提供有关空间数据和几何算法,使用`Go`语言包装实现。
欢迎大家使用并提出宝贵意见！

## 内容列表

- [安装](#安装)
- [使用说明](#使用说明)
- [维护者](#维护者)
- [如何贡献](#如何贡献)
- [使用许可](#使用许可)


## 安装

项目依赖 [GEOS](https://github.com/libgeos/geos)（GEOS 是 JTS 的C++版本实现) ,需要首先完成`GEOS`的安装。`GEOS`安装方法如下：

1、OS X系统安装(brew 方式)
```sh
$ brew install geos
```
2、Ubuntu
```sh
$ apt-get install libgeos-dev
```
3、源码安装
```sh
$ wget http://download.osgeo.org/geos/geos-3.9.0.tar.bz2
$ tar xvfj geos-3.9.0.tar.bz2
$ cd geos-3.9.0
$ ./configure
$ make
$ sudo make install
```

## 目录结构
1. `geo` 包下是对`GEOS C`库的引用和调用，以此来实现空间运算。
2. `algorithm` 是对外暴露的空间运算方法定义。
3. `strategy.go` 定义了空间运算底层算法的选择实现。

## 使用说明
以计算面积`Area`为例。
```
package main

import (
	"fmt"
	"github.com/spatial-go/geoos"
)

func main() {
	// First, choose the default algorithm.
	strategy := geoos.NormalStrategy()
	// Secondly, manufacturing test data and convert it to geometry
	const wkt = `POLYGON((-1 -1, 1 -1, 1 1, -1 1, -1 -1))`
	geometry, _ := geoos.UnmarshalString(wkt)
	// Last， call the Area () method and get result.
	area, e := strategy.Area(geometry)
	if e != nil {
		fmt.Printf(e.Error())
	}
	fmt.Printf("%f", area)
	// get result 4.0
}

```

## 维护者

[@spatial-go](https://github.com/spatial-go)。


## 如何贡献

感谢 [OSGeo](https://www.osgeo.org/)，我们也将秉承“开放、共创、共赢”的目标理念在空间计算领域贡献自己的一份力量。

非常欢迎你的加入！[提一个 Issue](https://github.com/spatial-go/geos/issues/new) 

联系邮箱： [os@changjing.ai](os@changjing.ai)


## 使用许可

[LGPL-2.1 ](LICENSE)

