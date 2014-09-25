// THIS FILE WAS AUTOGENERATED - ANY EDITS TO THIS WILL BE LOST WHEN IT IS REGENERATED

package sqlc



type varcharField struct {
	name string
	table TableLike
}

type VarcharField interface {
	TableField
	Eq(value string) Condition
	IsEq(value VarcharField) JoinCondition
}

func (c *varcharField) Name() string {
	return c.name
}

func (c *varcharField) Table() string {
	return c.table.Name()
}

func (c *varcharField) Eq(pred string) Condition {
	return Condition{Binding: FieldBinding{Value: pred, Field: c}}
}

func (c *varcharField) IsEq(pred VarcharField) JoinCondition {
	return JoinCondition{Lhs: c, Rhs: pred, Predicate: EqPredicate}
}

func Varchar(table TableLike, name string) VarcharField {
	return &varcharField{name: name, table:table}
}

