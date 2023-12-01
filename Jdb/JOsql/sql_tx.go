package JOsql

// TxOptions are transaction modes within a transaction block
type TxOptions struct {
	IsoLevel       TxIsoLevel
	AccessMode     TxAccessMode
	DeferrableMode TxDeferrableMode

	// BeginQuery is the SQL query that will be executed to begin the transaction. This allows using non-standard syntax
	// such as BEGIN PRIORITY HIGH with CockroachDB. If set this will override the other settings.
	BeginQuery string
}

type TxIsoLevel string

// Transaction isolation levels
const (
	Serializable    TxIsoLevel = "serializable"
	RepeatableRead  TxIsoLevel = "repeatable read"
	ReadCommitted   TxIsoLevel = "read committed"
	ReadUncommitted TxIsoLevel = "read uncommitted"
)

// TxAccessMode is the transaction access mode (read write or read only)
type TxAccessMode string

// Transaction access modes
const (
	ReadWrite TxAccessMode = "read write"
	ReadOnly  TxAccessMode = "read only"
)

// TxDeferrableMode is the transaction deferrable mode (deferrable or not deferrable)
type TxDeferrableMode string

// Transaction deferrable modes
const (
	Deferrable    TxDeferrableMode = "deferrable"
	NotDeferrable TxDeferrableMode = "not deferrable"
)
