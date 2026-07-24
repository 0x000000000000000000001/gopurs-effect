

func PureE(a any) func() any {
	return func() any {
		return a
	}
}

func BindE(a func() any, f func(any) func() any) func() any {
	return func() any {
		resA := a()
		return f(resA)()
	}
}

func UntilE(f func() bool) func() any {
	return func() any {
		for {
			if f() {
				break
			}
		}
		return nil
	}
}

func WhileE(f func() bool, a func() any) func() any {
	return func() any {
		for {
			if !f() {
				break
			}
			a()
		}
		return nil
	}
}

func ForE(lo int, hi int, f func(int) func() any) func() any {
	return func() any {
		for i := lo; i < hi; i++ {
			f(i)()
		}
		return nil
	}
}

func ForeachE(as []any, f func(any) func() any) func() any {
	return func() any {
		for _, v := range as {
			f(v)()
		}
		return nil
	}
}
