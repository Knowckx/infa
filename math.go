package infa

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

