package JOsql

import (
	"strings"
)

// FilterNamedArg struct is an abstraction that represents named argument filters for SQL queries.
// Each FilterNamedArg has a Value, NamedArg, ComparisonOperator, and a LogicalOperator.
type FilterNamedArg struct {
	Column      string
	Value       any
	NamedArg    string
	Comparasion ComparisonOperator
	Logical     LogicalOperator
	Type        ClausaWhereSql
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
func (q QFilterNamedArgs) ToQuery(firstWhere bool, prefix PrefixNamedArgPG) (query string, namedArgs map[string]any) {
	if firstWhere && len(q) > 0 {
		query += "WHERE "
	}

	namedArgs = make(map[string]any)
	for _, filter := range q {

		switch filter.Type {
		case FullTextSearch:
			namedArgs["full_text_search"] = filter.Value
		case In, NotIn:
			if filter.Value == nil {
				continue
			}
			val, ok := filter.Value.([]string)
			if ok {
				filter.Value = strings.Join(val, ", ")
			}

			if filter.Logical != "" {
				query += string(filter.Logical) + " "
			}
			query += filter.Column + " " + string(filter.Type) + " (" + string(prefix) + filter.NamedArg + ") "
			namedArgs[filter.NamedArg] = filter.Value
		case IsNull, IsNotNull:
			if filter.Logical != "" {
				query += string(filter.Logical) + " "
			}
			query += filter.Column + " " + string(filter.Type) + " "
		case Like, NotLike:
			if filter.Value == nil {
				continue
			}

			if filter.Logical != "" {
				query += string(filter.Logical) + " "
			}
			query += filter.Column + " " + string(filter.Type) + " %" + string(prefix) + filter.NamedArg + "% "
			namedArgs[filter.NamedArg] = filter.Value
		default:
			if filter.Comparasion == "" {
				filter.Comparasion = Equals
			}

			if filter.Value == nil {
				continue
			}

			if filter.Logical != "" {
				query += string(filter.Logical) + " "
			}
			query += filter.Column + " " + string(filter.Comparasion) + " " + string(prefix) + filter.NamedArg + " "
		}
		// if filter.Comparasion == IsNotNull || filter.Comparasion == IsNull {
		// 	query += filter.Column + " " + string(filter.Comparasion) + " "
		// } else if filter.Value != nil {
		// 	if filter.NamedArg == "" {
		// 		filter.NamedArg = filter.Column
		// 	}
		// 	query += filter.Column + " " + string(filter.Comparasion) + " " + string(prefix) + filter.NamedArg + " "
		// 	namedArgs[filter.NamedArg] = filter.Value
		// }

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
