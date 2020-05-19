package index

const(
	ValueBiggerRelationship 	= iota +1  // 大于关系
	ValueSmallerRelationship			   // 小于关系
	ValueEqualRelationship				   // 等于关系
	ValueNotEqualRelationship 			   // 不等于关系
)

// 索引接口
type Index interface {
	// 通过查询条件获取到数据存储位置
	GetPosition(conditionFieldList []string, conditionList []int, conditionValueList []interface{}) string

	// 设置索引
	SetPosition(conditionFieldList []string, conditionValueList []interface{}, position string) error
}