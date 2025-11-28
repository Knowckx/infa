package infa

import (
	"fmt"
	"math"
)

// 坐标点
type Int2 struct {
	X int
	Y int
}

func (s Int2) String() string {
	out := fmt.Sprintf("[%d, %d]", s.X, s.Y)
	return out
}

// 两点的距离
func (s *Int2) GetDistance(in *Int2) float64 {
	xGap := in.X - s.X
	yGap := in.Y - s.Y
	res := math.Sqrt(float64(xGap*xGap + yGap*yGap))
	return res
}

// 两点的中间位置
func (s *Int2) GetMidPos(in *Int2) *Int2 {
	out := &Int2{}
	out.X = (in.X + s.X) / 2
	out.Y = (in.Y + s.Y) / 2
	return out
}

// 放和计算有关的东西

// 平均数计算器
type Averager struct {
	sum   float64
	count int
	Avg   float64
}

// 定义一个方法来添加一个新的数并重新计算平均值
func (ac *Averager) AddNumber(num any) {
	useValue := 0.0
	if v, ok := num.(int); ok {
		useValue = float64(v)
	} else if v, ok := num.(float64); ok {
		useValue = v
	}
	ac.sum += useValue
	ac.count++
	ac.Avg = ac.sum / float64(ac.count)
}

// 整数求绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
