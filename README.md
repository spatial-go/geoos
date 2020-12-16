![](https://assets-dev.dituwuyou.com/changjing/static/images/logo_blue.png)

# Geos
我们的组织`spatial-go`正式成立，这是我们的第一个开源项目`geos`,`geos`提供有关空间数据和几何算法,使用`Go`语言包装实现。
欢迎大家使用并提出宝贵意见！

## 内容列表

- [背景](#背景)
- [安装](#安装)
- [使用说明](#使用说明)
- [相关仓库](#相关仓库)
- [维护者](#维护者)
- [如何贡献](#如何贡献)
- [使用许可](#使用许可)

## 背景

`Geos`是面向空间定义和空间计算的基础类库，基于`Go`语言实现。目前已实际用于商业项目中。


## 安装

项目依赖 [geos](https://github.com/libgeos/geos)（geos 是 JTS 的C++版本实现) 请先确保本地已经安装。

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
$ wget http://download.osgeo.org/geos/geos-3.3.8.tar.bz2
$ tar xvfj geos-3.3.8.tar.bz2
$ cd geos-3.3.8
$ ./configure
$ make
$ sudo make install
```

## 目录结构
![](http://ww1.sinaimg.cn/large/007CUp1qly1glmh76nm95j30jc0tq77g.jpg)
1. `geo` 包下是对`GEOS C`库的引用和调用，以此来实现空间运算。
2. `algorith-interface` 是对外暴露的空间运算定义。
3. `strategy.go` 定义了空间运算底层算法的选择实现。
```

func NormalStrategy() Algorithm {
   return AlgorithmStrategy(GEOS)
}

func AlgorithmStrategy(name string) Algorithm {
   switch name {
   case GEOS:
      return new(GEOSAlgorithm)
   default:
      return nil
   }
}
```
默认算法底层实现使用`Geos 的 C库`，后续可自行实现或使用其他算法实现空间运算接口，屏蔽算法具体实现。
    
## 使用说明
以计算面积`area`为例。

```
package main

import (
   "fmt"
   "github.com/spatial-go/geos"
)

func main() {
   strategy := geos.NormalStrategy()
   const wkt = `POLYGON((-1 -1, 1 -1, 1 1, -1 1, -1 -1))`
   geometry, _ := geos.UnmarshalString(wkt)
   area, e := strategy.Area(geometry)
   if e!=nil{
      fmt.Printf(e.Error())
   }

   fmt.Printf("%f",area)
   // 输出4.0

}
```


## 相关仓库

- [geos](https://github.com/libgeos/geos) — 一个C++实现 JTS 标准的空间操作类库。
- [https://github.com/paulmach/orb](https://github.com/paulmach/orb) Orb定义了一组用于在Golang中处理地理和平面/投影几何数据的方法

## 维护者

[@spatial-go](https://github.com/spatial-go)。

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/spatial-go/geos/issues/new) 在GIS领域贡献自己的一份力量。
联系邮箱： [RDC@changjing.ai](RDC@changjing.ai)
公众号：亿景智联


## 使用许可

[LGPL-2.1 ](LICENSE)

