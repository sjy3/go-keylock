# go-keylock
一个细粒度的锁工具

## keylock工具包的特性：

* 1、更细的粒度、更高的并发性能
* 2、使用简单
* 3、对于常驻程序来说，无需手动删除keyLock对象。
* 4、基于引用计数的方式回收锁资源。
* 5、被动式清理，无法额外的定时器。
* 6、通过sync.pool实现子锁对象的复用，减少内存分配

## 使用方法

keylock工具包的使用非常简单，只需要下载该包并在代码中导入就可以直接使用。下面是一个使用范例

`go get "github.com/sjy3/go-keylock"`

```golang
import lock "github.com/sjy3/go-keylock"

var keyLock *lock.KeyLock
keyLock = lock.NewKeyLock()

keylock.Lock("shi")
defer keyLock.Unlock("shi")

//需要保证线程安全的操作

```

## 性能及并发测试

我们对keylock工具包进行了并发和基准测试，使用的电脑平台硬件配置如下：

> 型号名称：	MacBook Pro   
> 型号标识符：	MacBookPro16,1   
> 处理器名称：	6-Core Intel Core i7   
> 处理器速度：	2.6 GHz   
> 处理器数目：	1   
> 核总数：	6   
> L2缓存（每个核）：	256 KB   
> L3缓存：	12 MB   
> 超线程技术：	已启用   
> 内存：	16 GB   

测试结果如下：

> goos: darwin    
> goarch: amd64    
> pkg: go-keylock    
> BenchmarkCombinationParallel    
> BenchmarkCombinationParallel-12    	 2004168	       578 ns/op    
> PASS
