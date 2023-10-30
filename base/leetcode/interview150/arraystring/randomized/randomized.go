package randomized

import "math/rand"

// 380. O(1) 时间插入、删除和获取随机元素

type RandomizedSet struct {
	vals     map[int]int
	valindex []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{vals: map[int]int{}, valindex: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.vals[val]; !ok {
		this.valindex = append(this.valindex, val)
		this.vals[val] = len(this.valindex) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.vals[val]; ok {
		this.valindex[this.vals[val]] = this.valindex[len(this.valindex)-1]
		this.vals[this.valindex[len(this.valindex)-1]] = this.vals[val]
		this.valindex = this.valindex[:len(this.valindex)-1]
		delete(this.vals, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	intn := rand.Intn(len(this.valindex))
	return this.valindex[intn]
}
