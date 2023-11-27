package Jsql

// FilterNamedQuery struct is an abstraction that represents named argument filters for SQL queries.
// Each FilterNamedQuery has a Value, NamedArg, ComparisonOperator, and a LogicalOperator.
type FilterNamedQuery struct {
	Value       string
	NamedArg    string
	Comparasion ComparisonOperator
	Logical     LogicalOperator
}

// QFilterNamed is a map that associates a name of a column (as string) with a FilterNamedQuery struct.
// It is used to generate SQL queries with named arguments.
// Check pgx.NamedArgs for more information.
type QFilterNamed map[string]FilterNamedQuery

// ToQuery is a method on the QFilterNamed type that generates an SQL query string by iterating over the QFilterNamed map
// and concatenating column names, comparison operators and logical operators.
// It also constructs a named arguments map from the QFilterNamed with column names as keys and filter values as values.
// Default logical operator is And.
// If namedArg is empty, it will be set to the column name.
// If value is empty, it will not be made into a string.
// firstWhere is a boolean that determines whether the WHERE keyword should be prepended to the query string.
// If QFIltersNamed is empty, it will return an empty string and an empty map.
func (q QFilterNamed) ToQuery(firstWhere bool, prefixNamedArg string) (query string, namedArgs map[string]any) {
	if firstWhere && len(q) > 0 {
		query += "WHERE "
	}

	namedArgs = make(map[string]any)
	i := 0
	totalFilters := len(q)
	for column, filter := range q {
		i++

		if filter.Comparasion == IsNotNull || filter.Comparasion == IsNull {
			query += column + " " + string(filter.Comparasion) + " "
		} else if filter.Value != "" {
			if filter.NamedArg == "" {
				filter.NamedArg = column
			}
			query += column + " " + string(filter.Comparasion) + " " + prefixNamedArg + filter.NamedArg + " "
			namedArgs[filter.NamedArg] = filter.Value
		}

		if i != totalFilters && filter.Value != "" {
			if filter.Logical == "" {
				filter.Logical = And
			}
			query += string(filter.Logical) + " "
		}
	}

	return
}

// GenerateQFilterNamed is a helper function to create an instance of FilterNamedQuery.
func GenerateQFilterNamed(value, namedArg string, comparasion ComparisonOperator, logical LogicalOperator) FilterNamedQuery {
	return FilterNamedQuery{
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}
}

// GenerateQFilterNamedArgByColumn is a helper function to create an instance of FilterNamedQuery, setting the NamedArg as the column
func GenerateQFilterNamedArgByColumn(value string, comparasion ComparisonOperator, logical LogicalOperator) FilterNamedQuery {
	return FilterNamedQuery{
		Value:       value,
		Comparasion: comparasion,
		Logical:     logical,
	}
}
