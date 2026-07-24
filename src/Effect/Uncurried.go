import "gopurs/output/gopurs_runtime"

func MkEffectFn1(f gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(f, a)
		}
	}
}

func MkEffectFn2(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b)
		}
	}
}

func MkEffectFn3(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c)
		}
	}
}

func MkEffectFn4(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d)
		}
	}
}

func MkEffectFn5(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e)
		}
	}
}

func MkEffectFn6(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e, g gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e), g)
		}
	}
}

func MkEffectFn7(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e, g, h gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e), g), h)
		}
	}
}

func MkEffectFn8(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e, g, h, i gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e), g), h), i)
		}
	}
}

func MkEffectFn9(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e, g, h, i, j gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e), g), h), i), j)
		}
	}
}

func MkEffectFn10(f gopurs_runtime.Value) func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a, b, c, d, e, g, h, i, j, k gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a), b), c), d), e), g), h), i), j), k)
		}
	}
}

func RunEffectFn1(f gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(a)
		}
	}
}

func RunEffectFn2(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(a, b)
			}
		}
	}
}

func RunEffectFn3(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(a, b, c)
				}
			}
		}
	}
}

func RunEffectFn4(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(a, b, c, d)
					}
				}
			}
		}
	}
}

func RunEffectFn5(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func() gopurs_runtime.Value {
							return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(a, b, c, d, e)
						}
					}
				}
			}
		}
	}
}

func RunEffectFn6(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func(g gopurs_runtime.Value) func() gopurs_runtime.Value {
							return func() gopurs_runtime.Value {
								var args []gopurs_runtime.Value
								args = append(args, a, b, c, d, e, g)
								return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(args)
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn7(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func(g gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
							return func(h gopurs_runtime.Value) func() gopurs_runtime.Value {
								return func() gopurs_runtime.Value {
									var args []gopurs_runtime.Value
									args = append(args, a, b, c, d, e, g, h)
									return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(args)
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn8(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func(g gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
							return func(h gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
								return func(i gopurs_runtime.Value) func() gopurs_runtime.Value {
									return func() gopurs_runtime.Value {
										var args []gopurs_runtime.Value
										args = append(args, a, b, c, d, e, g, h, i)
										return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(args)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn9(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func(g gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
							return func(h gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
								return func(i gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
									return func(j gopurs_runtime.Value) func() gopurs_runtime.Value {
										return func() gopurs_runtime.Value {
											var args []gopurs_runtime.Value
											args = append(args, a, b, c, d, e, g, h, i, j)
											return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(args)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func RunEffectFn10(f gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
	return func(a gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
		return func(b gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
			return func(c gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
				return func(d gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
					return func(e gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
						return func(g gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
							return func(h gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
								return func(i gopurs_runtime.Value) func(gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
									return func(j gopurs_runtime.Value) func(gopurs_runtime.Value) func() gopurs_runtime.Value {
										return func(k gopurs_runtime.Value) func() gopurs_runtime.Value {
											return func() gopurs_runtime.Value {
												var args []gopurs_runtime.Value
												args = append(args, a, b, c, d, e, g, h, i, j, k)
												return f.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(args)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
