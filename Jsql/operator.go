package Jsql

// ComparisonOperator is a string type used to denote SQL comparison operators.
type ComparisonOperator string

// LogicalOperator is a string type used to denote SQL logical operators.
type LogicalOperator string

// Constants for various SQL comparison operators.
const (
	// Equals represents the SQL "=" operator.
	Equals ComparisonOperator = "="

	// NotEquals represents the SQL "!=" operator.
	NotEquals ComparisonOperator = "!="

	// Like represents the SQL "LIKE" operator.
	Like ComparisonOperator = "LIKE"

	// NotLike represents the SQL "NOT LIKE" operator.
	NotLike ComparisonOperator = "NOT LIKE"

	// GreaterThan represents the SQL ">" operator.
	GreaterThan ComparisonOperator = ">"

	// GreaterOrEqual represents the SQL ">=" operator.
	GreaterOrEqual ComparisonOperator = ">="

	// LessThan represents the SQL "<" operator.
	LessThan ComparisonOperator = "<"

	// LessOrEqual represents the SQL "<=" operator.
	LessOrEqual ComparisonOperator = "<="

	// Before is a semantics alias for LessThan.
	Before = LessThan

	// After is a semantics alias for GreaterThan.
	After = GreaterThan

	// BeforeOrEqual is a semantics alias for LessOrEqual.
	BeforeOrEqual = LessOrEqual

	// AfterOrEqual is a semantics alias for GreaterOrEqual.
	AfterOrEqual = GreaterOrEqual

	// IsNull represents the SQL "IS NULL" operator.
	IsNull ComparisonOperator = "IS NULL"

	// IsNotNull represents the SQL "IS NOT NULL" operator.
	IsNotNull ComparisonOperator = "IS NOT NULL"

	// And represents the SQL "AND" logical operator.
	And LogicalOperator = "AND"

	// Or represents the SQL "OR" logical operator.
	Or LogicalOperator = "OR"
)
