// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/hawa130/serverx/ent/publickey"
	"github.com/hawa130/serverx/ent/user"
	"github.com/rs/xid"
)

// PublicKey is the model entity for the PublicKey schema.
type PublicKey struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// ExpiredAt holds the value of the "expired_at" field.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PublicKeyQuery when eager-loading is set.
	Edges           PublicKeyEdges `json:"edges"`
	public_key_user *xid.ID
	selectValues    sql.SelectValues
}

// PublicKeyEdges holds the relations/edges for other nodes in the graph.
type PublicKeyEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PublicKeyEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PublicKey) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case publickey.FieldKey, publickey.FieldName, publickey.FieldDescription, publickey.FieldType, publickey.FieldStatus:
			values[i] = new(sql.NullString)
		case publickey.FieldCreatedAt, publickey.FieldUpdatedAt, publickey.FieldExpiredAt:
			values[i] = new(sql.NullTime)
		case publickey.FieldID:
			values[i] = new(xid.ID)
		case publickey.ForeignKeys[0]: // public_key_user
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PublicKey fields.
func (pk *PublicKey) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case publickey.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pk.ID = *value
			}
		case publickey.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pk.CreatedAt = value.Time
			}
		case publickey.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pk.UpdatedAt = value.Time
			}
		case publickey.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				pk.Key = value.String
			}
		case publickey.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pk.Name = value.String
			}
		case publickey.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pk.Description = value.String
			}
		case publickey.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pk.Type = value.String
			}
		case publickey.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pk.Status = value.String
			}
		case publickey.FieldExpiredAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expired_at", values[i])
			} else if value.Valid {
				pk.ExpiredAt = value.Time
			}
		case publickey.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field public_key_user", values[i])
			} else if value.Valid {
				pk.public_key_user = new(xid.ID)
				*pk.public_key_user = *value.S.(*xid.ID)
			}
		default:
			pk.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PublicKey.
// This includes values selected through modifiers, order, etc.
func (pk *PublicKey) Value(name string) (ent.Value, error) {
	return pk.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the PublicKey entity.
func (pk *PublicKey) QueryUser() *UserQuery {
	return NewPublicKeyClient(pk.config).QueryUser(pk)
}

// Update returns a builder for updating this PublicKey.
// Note that you need to call PublicKey.Unwrap() before calling this method if this PublicKey
// was returned from a transaction, and the transaction was committed or rolled back.
func (pk *PublicKey) Update() *PublicKeyUpdateOne {
	return NewPublicKeyClient(pk.config).UpdateOne(pk)
}

// Unwrap unwraps the PublicKey entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pk *PublicKey) Unwrap() *PublicKey {
	_tx, ok := pk.config.driver.(*txDriver)
	if !ok {
		panic("ent: PublicKey is not a transactional entity")
	}
	pk.config.driver = _tx.drv
	return pk
}

// String implements the fmt.Stringer.
func (pk *PublicKey) String() string {
	var builder strings.Builder
	builder.WriteString("PublicKey(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pk.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pk.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pk.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(pk.Key)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pk.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pk.Description)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(pk.Type)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(pk.Status)
	builder.WriteString(", ")
	builder.WriteString("expired_at=")
	builder.WriteString(pk.ExpiredAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PublicKeys is a parsable slice of PublicKey.
type PublicKeys []*PublicKey
