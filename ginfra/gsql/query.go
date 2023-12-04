package gsql

type Filter struct {
	Column      string
	Value       any
	Comparasion ComparisonOperator
	Logical     LogicalOperator
	Type        ClausaWhereSql
}

// Filters is a map that associates a name of a column (as string) with a Filter struct.
// It is used to generate SQL queries with named arguments.
// Check pgx.NamedArgs for more information.
type Filters []Filter
