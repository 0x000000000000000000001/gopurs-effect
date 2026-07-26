

func PureE(a interface{}) func() interface{} {
	return func() interface{} {
		return a
	}
}

func BindE(a func() interface{}, f func(interface{}) func() interface{}) func() interface{} {
	return func() interface{} {
		resA := a()
		return f(resA)()
	}
}

func UntilE(f func() bool) func() interface{} {
	return func() interface{} {
		for {
			if f() {
				break
			}
		}
		return nil
	}
}

func WhileE(f func() bool, a func() interface{}) func() interface{} {
	return func() interface{} {
		for {
			if !f() {
				break
			}
			a()
		}
		return nil
	}
}

func ForE(lo int, hi int, f func(int) func() interface{}) func() interface{} {
	return func() interface{} {
		for i := lo; i < hi; i++ {
			f(i)()
		}
		return nil
	}
}

func ForeachE(as []interface{}, f func(interface{}) func() interface{}) func() interface{} {
	return func() interface{} {
		for _, v := range as {
			f(v)()
		}
		return nil
	}
}
