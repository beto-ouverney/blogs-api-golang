package errors

// Application error codes.
const (
	ECONFLICT  = "conflict"  // action cannot be performed
	EINTERNAL  = "internal"  // internal error
	EINVALID   = "invalid"   // validation failed
	ENOTFOUND  = "not_found" // entity does not exist
	EFORBIDDEN = "forbidden" //operation forbidden
	EEXPECTED  = "expected"  //expected error that don't need to be logged
	ETIMEOUT   = "timeout"
)

// Error defines a standard application error.
type Error struct {
	Code    string // Machine-readable error code (papel da aplicação)
	Message string // Human-readable message (papel do usuário final)
	Op      string // Logical operation (papel da operação)
	Err     error  // Embedded error  (papel da operação)
	Detail  []byte // JSON encoded data  (papel da operação)
}
