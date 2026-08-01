package Effect_Uncurried

import (
    "gopurs/output/gopurs_runtime"
)

func box(v interface{}) gopurs_runtime.Value {
    if val, ok := v.(gopurs_runtime.Value); ok {
        return val
    }
    return gopurs_runtime.Box(v)
}

func MkEffectFn1(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply(fn, a)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn2(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func2(func(a gopurs_runtime.Value, b gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply2(fn, a, b)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn3(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func3(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply3(fn, a, b, c)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn4(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func4(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply4(fn, a, b, c, d)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn5(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func5(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply5(fn, a, b, c, d, e)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn6(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func6(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, argf gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply6(fn, a, b, c, d, e, argf)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn7(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func7(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, argf gopurs_runtime.Value, g gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply7(fn, a, b, c, d, e, argf, g)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn8(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func8(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, argf gopurs_runtime.Value, g gopurs_runtime.Value, h gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply8(fn, a, b, c, d, e, argf, g, h)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn9(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func9(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, argf gopurs_runtime.Value, g gopurs_runtime.Value, h gopurs_runtime.Value, i gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply9(fn, a, b, c, d, e, argf, g, h, i)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func MkEffectFn10(f interface{}) interface{} {
    fn := box(f)
    return gopurs_runtime.Func10(func(a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, argf gopurs_runtime.Value, g gopurs_runtime.Value, h gopurs_runtime.Value, i gopurs_runtime.Value, j gopurs_runtime.Value) gopurs_runtime.Value {
        eff := gopurs_runtime.Apply10(fn, a, b, c, d, e, argf, g, h, i, j)
        return gopurs_runtime.Apply(eff, gopurs_runtime.Value{})
    })
}

func RunEffectFn1(f interface{}, a interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.Apply(box(f), box(a)).UnsafePtr
    }
}

func RunEffectFn2(f interface{}, a interface{}, b interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp2(box(f), box(a), box(b)).UnsafePtr
    }
}

func RunEffectFn3(f interface{}, a interface{}, b interface{}, c interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp3(box(f), box(a), box(b), box(c)).UnsafePtr
    }
}

func RunEffectFn4(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp4(box(f), box(a), box(b), box(c), box(d)).UnsafePtr
    }
}

func RunEffectFn5(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp5(box(f), box(a), box(b), box(c), box(d), box(e)).UnsafePtr
    }
}

func RunEffectFn6(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, argf interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp6(box(f), box(a), box(b), box(c), box(d), box(e), box(argf)).UnsafePtr
    }
}

func RunEffectFn7(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, argf interface{}, g interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp7(box(f), box(a), box(b), box(c), box(d), box(e), box(argf), box(g)).UnsafePtr
    }
}

func RunEffectFn8(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, argf interface{}, g interface{}, h interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp8(box(f), box(a), box(b), box(c), box(d), box(e), box(argf), box(g), box(h)).UnsafePtr
    }
}

func RunEffectFn9(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, argf interface{}, g interface{}, h interface{}, i interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp9(box(f), box(a), box(b), box(c), box(d), box(e), box(argf), box(g), box(h), box(i)).UnsafePtr
    }
}

func RunEffectFn10(f interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, argf interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
    return func(_ interface{}) interface{} {
        return gopurs_runtime.UncurriedApp10(box(f), box(a), box(b), box(c), box(d), box(e), box(argf), box(g), box(h), box(i), box(j)).UnsafePtr
    }
}
