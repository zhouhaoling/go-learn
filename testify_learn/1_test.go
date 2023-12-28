package testify_learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {
	var a = 100
	var b = 200
	var c = 100
	assert.Equal(t, a, b, "%d != %d", a, b)
	assert.NotEqual(t, a, c, "%d == %d", a, c)
}

func TestContains(t *testing.T) {
	ct := assert.New(t)
	//是否包含对应字符串
	var e = "123456"
	var f = "1"
	ct.Contains(e, f, "%s 中不包含 %s", e, f)
	//是否包含对应数组
	var one = [5]int{1, 2, 3, 4, 5}
	var two = [2]int{3, 6}
	ct.Contains(one, two, "one does not contain two")
}

func TestDirExists(t *testing.T) {
	de := assert.New(t)
	s := "1d/"
	de.DirExists(s, "%s 不是目录", s)
	d := "1d/d/"
	de.DirExists(d, "%s 不是目录", d)
	h := "1d/h.txt"
	de.DirExists(h, "%s 不是目录", h)
}

func TestElementsMatch(t *testing.T) {
	//包含相同的元素是指这两个数组或切片里面的元素种类和个数都是一样的，只可以顺序不同
	em := assert.New(t)
	one := [5]int{1, 2, 3, 4, 5}
	two := [5]int{5, 4, 2, 3, 1}
	em.ElementsMatch(one, two, "不包含相同元素")
	//元素种类不一样
	three := [4]int{1, 2, 3, 4}
	em.ElementsMatch(one, three, "one 和 three不包含相同元素")
	four := [5]int{1, 2, 3, 4, 8}
	em.ElementsMatch(one, four, "one 和 four不包含相同元素")
}

func TestElementMatch2(t *testing.T) {
	//重复元素的情况下比较
	//第一种
	one := [5]int{3, 3, 9, 0, 10}
	two := [5]int{3, 9, 0, 3, 10}
	assert.ElementsMatch(t, one, two, "不包含相同元素")
	//第二种
	three := [4]int{3, 9, 0, 10}
	assert.ElementsMatch(t, one, three, "one 和 three 不包含相同元素")
}

func TestEmpty(t *testing.T) {
	//字符串类型
	em := assert.New(t)
	s := "hello"
	d := ""
	em.Empty(d, "d 不为空")
	em.Empty(s, "s 不为空")
	//整数
	a := 1
	b := 0
	em.Empty(a, "a 不为0")
	em.Empty(b, "b 不为0")
}

type MyInt int

func TestEqualValues(t *testing.T) {
	var a = 100
	var b MyInt = 100
	assert.Equal(t, a, b, "a 和 MyInt b不相等")
	assert.EqualValues(t, a, b, "a 和 b不相等")
}
