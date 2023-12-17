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

type ClausaWhereSql string

const (
	Where          ClausaWhereSql = "WHERE"
	FullTextSearch ClausaWhereSql = "full_text_search"
	In             ClausaWhereSql = "IN"
	NotIn          ClausaWhereSql = "NOT IN"
	IsNull         ClausaWhereSql = "IS NULL"
	IsNotNull      ClausaWhereSql = "IS NOT NULL"
	Like           ClausaWhereSql = "LIKE"
	NotLike        ClausaWhereSql = "NOT LIKE"
)

// ComparisonOperator is a string type used to denote SQL comparison operators.
type ComparisonOperator string

// Constants for various SQL comparison operators.
const (
	// Equals represents the SQL "=" operator.
	Equals ComparisonOperator = "="

	// NotEquals represents the SQL "<>" operator.
	NotEquals ComparisonOperator = "<>"

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
)

// LogicalOperator is a string type used to denote SQL logical operators.
type LogicalOperator string

const (
	// And represents the SQL "AND" logical operator.
	And LogicalOperator = "AND"

	// Or represents the SQL "OR" logical operator.
	Or LogicalOperator = "OR"
)

type PrefixNamedArgPG string

const (
	PgxNamedArg  PrefixNamedArgPG = "@"
	SqlxNamedArg PrefixNamedArgPG = ":"
)

// LockingOperator represents a type for database locking operations.
// This type is used to manipulate and control database locking operations,
// like "UPDATE", "INSERT", "DELETE".
type LockingOperator string

const (
	ForUpdate LockingOperator = "FOR UPDATE"
	ForInsert LockingOperator = "FOR INSERT"
	ForDelete LockingOperator = "FOR DELETE"
	NotLock   LockingOperator = ""
)
