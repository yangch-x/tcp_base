package zpack

import (
	"myTcpBase/ziface"
	"sync"
)

var packOnce sync.Once

type PackFactory struct{}

var factoryInstance *PackFactory

/*
生成不同封包解包的方式，单例
*/
func Factory() *PackFactory {
	packOnce.Do(func() {
		factoryInstance = new(PackFactory)
	})

	return factoryInstance
}

// NewPack 创建一个具体的拆包解包对象
func (f *PackFactory) NewPack(kind string) ziface.IDataPack {
	var dataPack ziface.IDataPack

	switch kind {
	// 标准默认封包拆包方式
	case ziface.DefaultPackDataPack:
		dataPack = NewDataPack()
		break

		//case 自定义封包拆包方式case

	default:
		dataPack = NewDataPack()
	}

	return dataPack
}
