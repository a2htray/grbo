水稻育种优化算法 Go 实现
==================================

Table of contents

* [介绍](#介绍)
* [安装](#安装)
* [示例](#示例)

## 介绍

水稻育种优化算法是一种受杂交水稻育种技术启发的优化算法，该算法模拟了三系杂交水稻育种的一系列过程。三系分别为保持系、恢复系和不育系，各系之间或系
内部模拟杂交操作或自交操作，从而得到更优的个体。

## 安装

`go get -u github.com/a2htray/grbo`

## 示例

```go
package main

import (
	"fmt"
	"math"
	"github.com/a2htray/grbo"
)

func main() {
	rbo := grbo.New(60, 10, grbo.WithT(3000), grbo.WithC(50), grbo.WithLowerLimit([]float64{
		-100, -100, -100, -100, -100, -100, -100, -100, -100, -100,
    }), grbo.WithUpperLimit([]float64{
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
    }), grbo.WithObjectFunc(func(fs []float64) float64 {
		part := 0.0
		for i := 1; i < len(fs); i++ {
			part += math.Pow(fs[i], 2)
		}
		return math.Pow(fs[0], 2) + math.Pow(10, 6) * part
	}))
	
	rbo.Run()

	for i, rice := range rbo.HistoryBest() {
		fmt.Printf("%iter %d, fitness = %v\n", i, rice.Values(), rice.Fitness())
	}
}
```

