package main

import (
	"fmt"
	"math/rand"
)

// 剑指 Offer II 030. 插入、删除和随机访问都是 O(1) 的容器
// 设计一个支持在平均时间复杂度 O(1)下，执行以下操作的数据结构：
//
//insert(val)：当元素 val 不存在时返回 true，并向集合中插入该项，否则返回 false 。
//remove(val)：当元素 val 存在时返回 true，并从集合中移除该项，否则返回 false。
//getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。
//
//
//示例 :
//
//输入: inputs = ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
//[[], [1], [2], [2], [], [1], [2], []]
//输出: [null, true, false, true, 2, true, false, 2]
//解释:
//RandomizedSet randomSet = new RandomizedSet();  // 初始化一个空的集合
//randomSet.insert(1); // 向集合中插入 1 ， 返回 true 表示 1 被成功地插入
//
//randomSet.remove(2); // 返回 false，表示集合中不存在 2
//
//randomSet.insert(2); // 向集合中插入 2 返回 true ，集合现在包含 [1,2]
//
//randomSet.getRandom(); // getRandom 应随机返回 1 或 2
//
//randomSet.remove(1); // 从集合中移除 1 返回 true 。集合现在包含 [2]
//
//randomSet.insert(2); // 2 已在集合中，所以返回 false
//
//randomSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2
//
//
//提示：
//
//-231<= val <= 231- 1
//最多进行 2 * 105 次insert ， remove 和 getRandom 方法调用
//当调用getRandom 方法时，集合中至少有一个元素
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/FortPu
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type RandomizedSet struct {
	dict map[int]int
	list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	var list []int
	dict := make(map[int]int)
	return RandomizedSet{
		dict: dict,
		list: list,
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dict[val]; ok {
		return false
	}

	this.dict[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.dict[val]
	if !ok {
		return false
	}

	lastIdx := len(this.list) - 1
	lv := this.list[lastIdx]
	this.list[idx] = lv
	this.dict[lv] = idx
	this.list = this.list[:lastIdx]
	delete(this.dict, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

func main() {
	obj := Constructor()
	p1 := obj.Insert(2)
	_ = obj.Insert(1)
	p2 := obj.Remove(2)
	p3 := obj.GetRandom()

	fmt.Println(p1, p2, p3)
}
