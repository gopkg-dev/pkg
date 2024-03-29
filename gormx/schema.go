package gormx

// ListResult 响应列表数据
type ListResult struct {
	List       interface{}       `json:"list"`
	Pagination *PaginationResult `json:"pagination,omitempty"`
}

// PaginationResult 分页查询结果
type PaginationResult struct {
	Total    int64 `json:"total"`
	Current  int   `json:"current"`
	PageSize int   `json:"pageSize"`
}

// PaginationParam 分页查询条件
type PaginationParam struct {
	Pagination bool `query:"-"`        // 是否使用分页查询
	OnlyCount  bool `query:"-"`        // 是否仅查询count
	Current    int  `query:"current"`  // 当前页
	PageSize   int  `query:"pageSize"` // 页大小
}

// GetCurrent 获取当前页
func (a PaginationParam) GetCurrent() int {
	if a.Current == 0 {
		return 1
	}
	return a.Current
}

// GetPageSize 获取页大小
func (a PaginationParam) GetPageSize() int {
	pageSize := a.PageSize
	if a.PageSize == 0 {
		pageSize = 100
	}
	return pageSize
}

// QueryOptions 查询可选参数项
type QueryOptions struct {
	OrderFields  []*OrderField // 排序字段
	SelectFields []string      //
}

// OrderDirection 排序方向
type OrderDirection int

const (
	// OrderByASC 升序排序
	OrderByASC OrderDirection = iota + 1
	// OrderByDESC 降序排序
	OrderByDESC
)

// NewOrderFieldWithKeys 创建排序字段(默认升序排序)，可指定不同key的排序规则
// keys 需要排序的key
// directions 排序规则，按照key的索引指定，索引默认从0开始
func NewOrderFieldWithKeys(keys []string, directions ...map[int]OrderDirection) []*OrderField {
	m := make(map[int]OrderDirection)
	if len(directions) > 0 {
		m = directions[0]
	}

	fields := make([]*OrderField, len(keys))
	for i, key := range keys {
		d := OrderByASC
		if v, ok := m[i]; ok {
			d = v
		}

		fields[i] = NewOrderField(key, d)
	}

	return fields
}

// NewOrderFields 创建排序字段列表
func NewOrderFields(orderFields ...*OrderField) []*OrderField {
	return orderFields
}

// NewOrderField 创建排序字段
func NewOrderField(key string, d OrderDirection) *OrderField {
	return &OrderField{
		Key:       key,
		Direction: d,
	}
}

// OrderField 排序字段
type OrderField struct {
	Key       string         // 字段名(字段名约束为小写蛇形)
	Direction OrderDirection // 排序方向
}

// NewIDResult 创建响应唯一标识实例
func NewIDResult(id string) *IDResult {
	return &IDResult{
		ID: id,
	}
}

// IDResult 响应唯一标识
type IDResult struct {
	ID string `json:"id"`
}
