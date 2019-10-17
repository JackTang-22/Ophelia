package base

import (
	"fmt"
	"sync"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/17 17:12
 * @Description:
 */

func TestMap(t *testing.T) {
	// 常见的创建
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(m, m["a"]) // map[a:1 b:2] 1
	m1 := make(map[string]int)
	m1["ab"] = 3
	fmt.Println(m1, m1["ab"]) // map[ab:3] 3
	delete(m1, "c")           // 删除一个k-v 如果k不存在 则什么也不执行

	// ********************************************
	// 并发读写的map
	// 编译错误 fatal error: concurrent map read and map write
	// 出现了并发读写，因为用两个并发程序不断的对map进行读和写，产生了竞态问题。map内部会对这种错误进行检查并提前发现
	/*m2 := make(map[string]int)
	    // 无限读
	    go func() {
			for {
				m2["1"] = 1
			}
		}()
	    // 无限取
	    go func() {
			for {
				_ = m2["1"]
			}
		}()

		for {
			_ = 1
		}*/

	// ***************************************
	// sync.Map
	var syncMap sync.Map
	// 写入
	syncMap.Store("a", 1)
	syncMap.Store("s", "b")
	syncMap.Store("s1", "b1")
	// 读取
	fmt.Println(syncMap) // {{0 0} {{map[] true}} map[a:0xc000086020 s:0xc000086028 s1:0xc000086030] 0}
	fmt.Println(syncMap.Load("a")) // 1 true
	// 删除
	syncMap.Delete("s1")
	// 遍历
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, "->", value)
		return true
	})
}
