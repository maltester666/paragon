// Code generated by entc, DO NOT EDIT.

package target

import (
	"github.com/kcarretto/paragon/ent/schema"
)

const (
	// Label holds the string label denoting the target type in the database.
	Label = "target"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldPrimaryIP holds the string denoting the primaryip vertex property in the database.
	FieldPrimaryIP = "primary_ip"
	// FieldMachineUUID holds the string denoting the machineuuid vertex property in the database.
	FieldMachineUUID = "machine_uuid"
	// FieldPublicIP holds the string denoting the publicip vertex property in the database.
	FieldPublicIP = "public_ip"
	// FieldPrimaryMAC holds the string denoting the primarymac vertex property in the database.
	FieldPrimaryMAC = "primary_mac"
	// FieldHostname holds the string denoting the hostname vertex property in the database.
	FieldHostname = "hostname"
	// FieldLastSeen holds the string denoting the lastseen vertex property in the database.
	FieldLastSeen = "last_seen"

	// Table holds the table name of the target in the database.
	Table = "targets"
	// TasksTable is the table the holds the tasks relation/edge.
	TasksTable = "tasks"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "tasks"
	// TasksColumn is the table column denoting the tasks relation/edge.
	TasksColumn = "target_id"
	// TagsTable is the table the holds the tags relation/edge. The primary key declared below.
	TagsTable = "target_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// CredentialsTable is the table the holds the credentials relation/edge.
	CredentialsTable = "credentials"
	// CredentialsInverseTable is the table name for the Credential entity.
	// It exists in this package in order to avoid circular dependency with the "credential" package.
	CredentialsInverseTable = "credentials"
	// CredentialsColumn is the table column denoting the credentials relation/edge.
	CredentialsColumn = "target_id"
)

// Columns holds all SQL columns for target fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPrimaryIP,
	FieldMachineUUID,
	FieldPublicIP,
	FieldPrimaryMAC,
	FieldHostname,
	FieldLastSeen,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"target_id", "tag_id"}
)

var (
	fields = schema.Target{}.Fields()

	// descMachineUUID is the schema descriptor for MachineUUID field.
	descMachineUUID = fields[2].Descriptor()
	// MachineUUIDValidator is a validator for the "MachineUUID" field. It is called by the builders before save.
	MachineUUIDValidator = descMachineUUID.Validators[0].(func(string) error)
)
