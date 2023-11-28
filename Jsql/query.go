package Jsql

// FilterNamedArg struct is an abstraction that represents named argument filters for SQL queries.
// Each FilterNamedArg has a Value, NamedArg, ComparisonOperator, and a LogicalOperator.
type FilterNamedArg struct {
	Column      string
	Value       any
	NamedArg    string
	Comparasion ComparisonOperator
	Logical     LogicalOperator
}

// QFilterNamedArgs is a map that associates a name of a column (as string) with a FilterNamedArg struct.
// It is used to generate SQL queries with named arguments.
// Check pgx.NamedArgs for more information.
type QFilterNamedArgs []FilterNamedArg

// ToQuery is a method on the QFilterNamedArgs type that generates an SQL query string by iterating over the QFilterNamedArgs map
// and concatenating column names, comparison operators and logical operators.
// It also constructs a named arguments map from the QFilterNamedArgs with column names as keys and filter values as values.
// Default logical operator is And.
// If namedArg is empty, it will be set to the column name.
// If value is empty, it will not be made into a string.
// firstWhere is a boolean that determines whether the WHERE keyword should be prepended to the query string.
// If QFIltersNamed is empty, it will return an empty string and an empty map.
func (q QFilterNamedArgs) ToQuery(firstWhere bool, prefixNamedArg string) (query string, namedArgs map[string]any) {
	if firstWhere && len(q) > 0 {
		query += "WHERE "
	}

	namedArgs = make(map[string]any)
	i := 0
	totalFilters := len(q)
	for _, filter := range q {
		i++

		if filter.Comparasion == IsNotNull || filter.Comparasion == IsNull {
			query += filter.Column + " " + string(filter.Comparasion) + " "
		} else if filter.Value != nil {
			if filter.NamedArg == "" {
				filter.NamedArg = filter.Column
			}
			query += filter.Column + " " + string(filter.Comparasion) + " " + prefixNamedArg + filter.NamedArg + " "
			namedArgs[filter.NamedArg] = filter.Value
		}

		if i != totalFilters && filter.Value != nil {
			if filter.Logical == "" {
				filter.Logical = And
			}
			query += string(filter.Logical) + " "
		}
	}

	return
}

// GenerateQFilterNamed is a helper function to create an instance of FilterNamedArg.
func GenerateQFilterNamed(value, namedArg string, comparasion ComparisonOperator, logical LogicalOperator) FilterNamedArg {
	return FilterNamedArg{
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}
}

// GenerateQFilterNamedArgByColumn is a helper function to create an instance of FilterNamedArg, setting the NamedArg as the column
func GenerateQFilterNamedArgByColumn(value string, comparasion ComparisonOperator, logical LogicalOperator) FilterNamedArg {
	return FilterNamedArg{
		Value:       value,
		Comparasion: comparasion,
		Logical:     logical,
	}
}
