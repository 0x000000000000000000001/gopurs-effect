package Effect_Uncurried

func MkEffectFn1(f interface{}) interface{} {
	return func(a interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(nil)
	}
}

func MkEffectFn2(f interface{}) interface{} {
	return func(a, b interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(nil)
	}
}

func MkEffectFn3(f interface{}) interface{} {
	return func(a, b, c interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(nil)
	}
}

func MkEffectFn4(f interface{}) interface{} {
	return func(a, b, c, d interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(nil)
	}
}

func MkEffectFn5(f interface{}) interface{} {
	return func(a, b, c, d, e interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(nil)
	}
}

func MkEffectFn6(f interface{}) interface{} {
	return func(a, b, c, d, e, argf interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(argf).(func(interface{}) interface{})(nil)
	}
}

func MkEffectFn7(f interface{}) interface{} {
	return func(a, b, c, d, e, argf, g interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(argf).(func(interface{}) interface{})(g).(func(interface{}) interface{})(nil)
	}
}

func MkEffectFn8(f interface{}) interface{} {
	return func(a, b, c, d, e, argf, g, h interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(argf).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(nil)
	}
}

func MkEffectFn9(f interface{}) interface{} {
	return func(a, b, c, d, e, argf, g, h, i interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(argf).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i).(func(interface{}) interface{})(nil)
	}
}

func MkEffectFn10(f interface{}) interface{} {
	return func(a, b, c, d, e, argf, g, h, i, j interface{}) interface{} {
		return f.(func(interface{}) interface{})(a).(func(interface{}) interface{})(b).(func(interface{}) interface{})(c).(func(interface{}) interface{})(d).(func(interface{}) interface{})(e).(func(interface{}) interface{})(argf).(func(interface{}) interface{})(g).(func(interface{}) interface{})(h).(func(interface{}) interface{})(i).(func(interface{}) interface{})(j).(func(interface{}) interface{})(nil)
	}
}

func RunEffectFn1(f interface{}, a interface{}) interface{} {
	return func(_ interface{}) interface{} {
		return f.(func(interface{}) interface{})(a)
	}
}

func RunEffectFn2(f interface{}, a, b interface{}) interface{} {
	return func(_ interface{}) interface{} {
		return f.(func(interface{}, interface{}) interface{})(a, b)
	}
}

func RunEffectFn3(f interface{}, a, b, c interface{}) interface{} {
	return func(_ interface{}) interface{} {
		return f.(func(interface{}, interface{}, interface{}) interface{})(a, b, c)
	}
}

func RunEffectFn4(f interface{}, a, b, c, d interface{}) interface{} {
	return func(_ interface{}) interface{} {
		return f.(func(interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d)
	}
}

func RunEffectFn5(f interface{}, a, b, c, d, e interface{}) interface{} {
	return func(_ interface{}) interface{} {
		return f.(func(interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e)
	}
}

func RunEffectFn6(f interface{}, a, b, c, d, e, argf interface{}) interface{} {
	return func(_ interface{}) interface{} {
		return f.(func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e, argf)
	}
}

func RunEffectFn7(f interface{}, a, b, c, d, e, argf, g interface{}) interface{} {
	return func(_ interface{}) interface{} {
		return f.(func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e, argf, g)
	}
}

func RunEffectFn8(f interface{}, a, b, c, d, e, argf, g, h interface{}) interface{} {
	return func(_ interface{}) interface{} {
		return f.(func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e, argf, g, h)
	}
}

func RunEffectFn9(f interface{}, a, b, c, d, e, argf, g, h, i interface{}) interface{} {
	return func(_ interface{}) interface{} {
		return f.(func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e, argf, g, h, i)
	}
}

func RunEffectFn10(f interface{}, a, b, c, d, e, argf, g, h, i, j interface{}) interface{} {
	return func(_ interface{}) interface{} {
		return f.(func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{})(a, b, c, d, e, argf, g, h, i, j)
	}
}
