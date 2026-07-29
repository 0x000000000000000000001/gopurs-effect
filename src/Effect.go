package Effect

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
func UntilE(f func() any) func() any {
	return func() any {
		for {
			if f().(bool) {
				break
			}
		}
		return nil
	}
}
func WhileE(f func() any, a func() any) func() any {
	return func() any {
		for {
			if !f().(bool) {
				break
			}
			a()
		}
		return nil
	}
}
func ForE(lo int64, hi int64, f func(any) func() any) func() any {
	return func() any {
		for i := lo; i < hi; i++ {
			f(i)()
		}
		return nil
	}
}
func ForeachE(as []any, f func(any) func() any) func() any {
	return func() any {
		for _, a := range as {
			f(a)()
		}
		return nil
	}
}
