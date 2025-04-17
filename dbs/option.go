package dbs

type option struct {
	query []Query
	Order any
	limit int
}

// 定义选项函数类型
type fnOption func(*option)

// 定义选项函数
func WithQuery(query []Query) fnOption {
	return func(o *option) {
		o.query = query
	}
}

func WithOrder(order any) fnOption {
	return func(o *option) {
		o.Order = order
	}
}

func WithLimit(limit int) fnOption {
	return func(o *option) {
		o.limit = limit
	}
}

// 构造函数中应用选项
func NewOption(opts ...fnOption) *option {
	opt := &option{
		query: []Query{},
		Order: "id",
		limit: 0,
	}

	for _, optFn := range opts {
		optFn(opt)
	}

	return opt
}
