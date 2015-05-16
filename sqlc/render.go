package sqlc

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

var predicateTypes = map[PredicateType]string{
	EqPredicate: "=",
	GtPredicate: ">",
	GePredicate: ">=",
	LtPredicate: "<",
	LePredicate: "<=",
}

func (u *update) String(d Dialect) string {
	return toString(d, u)
}

func (u *update) Render(d Dialect, w io.Writer) (placeholders []interface{}) {
	fmt.Fprintf(w, "UPDATE %s SET ", u.table.Name())

	setFragments := make([]string, len(u.bindings))
	setValues := make([]interface{}, len(u.bindings))

	for i, binding := range u.bindings {
		col := binding.Field.Name()
		setFragments[i] = fmt.Sprintf("%s = %s", col, d.renderPlaceholder(i+1))
		setValues[i] = binding.Value
	}

	setClause := strings.Join(setFragments, ", ")
	fmt.Fprint(w, setClause)

	fmt.Fprint(w, " ")

	paramCount := len(u.bindings)
	whereValues := renderWhereClause(u.table.Name(), u.predicate, d, paramCount, w)

	placeholders = append(setValues, whereValues...)

	return placeholders
}

func (i *insert) String(d Dialect) string {
	return toString(d, i)
}

func (i *insert) Render(d Dialect, w io.Writer) (placeholders []interface{}) {
	fmt.Fprintf(w, "INSERT INTO %s (", i.table.Name())
	colFragments := make([]string, len(i.bindings))
	for i, binding := range i.bindings {
		col := binding.Field.Name()
		colFragments[i] = col
	}
	colClause := strings.Join(colFragments, ", ")
	fmt.Fprint(w, colClause)

	fmt.Fprint(w, ") VALUES (")

	placeHolderFragments := make([]string, len(i.bindings))
	values := make([]interface{}, len(i.bindings))
	for i, binding := range i.bindings {
		placeHolderFragments[i] = d.renderPlaceholder(i + 1)
		values[i] = binding.Value
	}

	placeHolderClause := strings.Join(placeHolderFragments, ",")
	fmt.Fprint(w, placeHolderClause)
	fmt.Fprint(w, ")")

	if i.returning != nil {
		fmt.Fprint(w, " RETURNING ")
		fmt.Fprint(w, i.returning.Name())
	}

	return values
}

func resolveParentAlias(alias string, col Field) string {
	if alias == "" {
		if tabCol, ok := col.(TableField); ok {
			if tabCol.Parent() != nil {
				return tabCol.Parent().MaybeAlias()
			}
		}
		return ""
	} else {
		return alias
	}
}

func columnClause(alias string, cols []Field) string {
	colFragments := make([]string, len(cols))
	for i, col := range cols {
		al := resolveParentAlias(alias, col)
		aliased := fmt.Sprintf("%s.%s", al, col.Name())

		var f string
		fun := col.Function()

		f = renderFunction(aliased, fun)

		if col.Alias() != "" {
			f = fmt.Sprintf("%s AS %s", f, col.Alias())
		}

		colFragments[i] = f

	}
	return strings.Join(colFragments, ", ")
}

func renderFunction(aliased string, fun FieldFunction) string {

	if &fun == nil {
		return ""
	} else if fun.Child != nil {
		aliased = renderFunction(aliased, *fun.Child)
	}

	var f string

	if fun.Name == "Count" {
		f = fun.Expr
	} else {

		if fun.Expr == "" {
			f = aliased
		} else {
			if len(fun.Args) > 0 {
				args := make([]interface{}, 1)
				args[0] = aliased
				args = append(args, fun.Args...)
				f = fmt.Sprintf(fun.Expr, args...)
			} else {
				f = fmt.Sprintf(fun.Expr, aliased)
			}
		}
	}

	return f
}

func renderWhereClause(alias string, conds []Condition, d Dialect, paramCount int, w io.Writer) []interface{} {
	fmt.Fprint(w, "WHERE ")

	whereFragments := make([]string, len(conds))
	values := make([]interface{}, len(conds))

	for i, condition := range conds {
		field := condition.Binding.Field
		al := resolveParentAlias(alias, field)
		col := field.Name()
		pred := condition.Predicate
		placeHolder := d.renderPlaceholder(i + paramCount + 1)
		whereFragments[i] = fmt.Sprintf("%s.%s %s %s", al, col, predicateTypes[pred], placeHolder)
		values[i] = condition.Binding.Value
	}

	whereClause := strings.Join(whereFragments, " AND ")
	fmt.Fprint(w, whereClause)

	return values
}

func (d Dialect) renderPlaceholder(n int) string {
	switch d {
	case Postgres:
		return fmt.Sprintf("$%d", n)
	case Oracle:
		return fmt.Sprintf(":%d", n)
	default:
		return "?"
	}
}

func toString(d Dialect, r Renderable) string {
	var buf bytes.Buffer
	r.Render(d, &buf)
	return buf.String()
}
