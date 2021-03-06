// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kcarretto/paragon/ent/credential"
	"github.com/kcarretto/paragon/ent/predicate"
	"github.com/kcarretto/paragon/ent/tag"
	"github.com/kcarretto/paragon/ent/target"
	"github.com/kcarretto/paragon/ent/task"
)

// TargetUpdate is the builder for updating Target entities.
type TargetUpdate struct {
	config
	Name               *string
	PrimaryIP          *string
	MachineUUID        *string
	clearMachineUUID   bool
	PublicIP           *string
	clearPublicIP      bool
	PrimaryMAC         *string
	clearPrimaryMAC    bool
	Hostname           *string
	clearHostname      bool
	LastSeen           *time.Time
	clearLastSeen      bool
	tasks              map[int]struct{}
	tags               map[int]struct{}
	credentials        map[int]struct{}
	removedTasks       map[int]struct{}
	removedTags        map[int]struct{}
	removedCredentials map[int]struct{}
	predicates         []predicate.Target
}

// Where adds a new predicate for the builder.
func (tu *TargetUpdate) Where(ps ...predicate.Target) *TargetUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetName sets the Name field.
func (tu *TargetUpdate) SetName(s string) *TargetUpdate {
	tu.Name = &s
	return tu
}

// SetPrimaryIP sets the PrimaryIP field.
func (tu *TargetUpdate) SetPrimaryIP(s string) *TargetUpdate {
	tu.PrimaryIP = &s
	return tu
}

// SetMachineUUID sets the MachineUUID field.
func (tu *TargetUpdate) SetMachineUUID(s string) *TargetUpdate {
	tu.MachineUUID = &s
	return tu
}

// SetNillableMachineUUID sets the MachineUUID field if the given value is not nil.
func (tu *TargetUpdate) SetNillableMachineUUID(s *string) *TargetUpdate {
	if s != nil {
		tu.SetMachineUUID(*s)
	}
	return tu
}

// ClearMachineUUID clears the value of MachineUUID.
func (tu *TargetUpdate) ClearMachineUUID() *TargetUpdate {
	tu.MachineUUID = nil
	tu.clearMachineUUID = true
	return tu
}

// SetPublicIP sets the PublicIP field.
func (tu *TargetUpdate) SetPublicIP(s string) *TargetUpdate {
	tu.PublicIP = &s
	return tu
}

// SetNillablePublicIP sets the PublicIP field if the given value is not nil.
func (tu *TargetUpdate) SetNillablePublicIP(s *string) *TargetUpdate {
	if s != nil {
		tu.SetPublicIP(*s)
	}
	return tu
}

// ClearPublicIP clears the value of PublicIP.
func (tu *TargetUpdate) ClearPublicIP() *TargetUpdate {
	tu.PublicIP = nil
	tu.clearPublicIP = true
	return tu
}

// SetPrimaryMAC sets the PrimaryMAC field.
func (tu *TargetUpdate) SetPrimaryMAC(s string) *TargetUpdate {
	tu.PrimaryMAC = &s
	return tu
}

// SetNillablePrimaryMAC sets the PrimaryMAC field if the given value is not nil.
func (tu *TargetUpdate) SetNillablePrimaryMAC(s *string) *TargetUpdate {
	if s != nil {
		tu.SetPrimaryMAC(*s)
	}
	return tu
}

// ClearPrimaryMAC clears the value of PrimaryMAC.
func (tu *TargetUpdate) ClearPrimaryMAC() *TargetUpdate {
	tu.PrimaryMAC = nil
	tu.clearPrimaryMAC = true
	return tu
}

// SetHostname sets the Hostname field.
func (tu *TargetUpdate) SetHostname(s string) *TargetUpdate {
	tu.Hostname = &s
	return tu
}

// SetNillableHostname sets the Hostname field if the given value is not nil.
func (tu *TargetUpdate) SetNillableHostname(s *string) *TargetUpdate {
	if s != nil {
		tu.SetHostname(*s)
	}
	return tu
}

// ClearHostname clears the value of Hostname.
func (tu *TargetUpdate) ClearHostname() *TargetUpdate {
	tu.Hostname = nil
	tu.clearHostname = true
	return tu
}

// SetLastSeen sets the LastSeen field.
func (tu *TargetUpdate) SetLastSeen(t time.Time) *TargetUpdate {
	tu.LastSeen = &t
	return tu
}

// SetNillableLastSeen sets the LastSeen field if the given value is not nil.
func (tu *TargetUpdate) SetNillableLastSeen(t *time.Time) *TargetUpdate {
	if t != nil {
		tu.SetLastSeen(*t)
	}
	return tu
}

// ClearLastSeen clears the value of LastSeen.
func (tu *TargetUpdate) ClearLastSeen() *TargetUpdate {
	tu.LastSeen = nil
	tu.clearLastSeen = true
	return tu
}

// AddTaskIDs adds the tasks edge to Task by ids.
func (tu *TargetUpdate) AddTaskIDs(ids ...int) *TargetUpdate {
	if tu.tasks == nil {
		tu.tasks = make(map[int]struct{})
	}
	for i := range ids {
		tu.tasks[ids[i]] = struct{}{}
	}
	return tu
}

// AddTasks adds the tasks edges to Task.
func (tu *TargetUpdate) AddTasks(t ...*Task) *TargetUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTaskIDs(ids...)
}

// AddTagIDs adds the tags edge to Tag by ids.
func (tu *TargetUpdate) AddTagIDs(ids ...int) *TargetUpdate {
	if tu.tags == nil {
		tu.tags = make(map[int]struct{})
	}
	for i := range ids {
		tu.tags[ids[i]] = struct{}{}
	}
	return tu
}

// AddTags adds the tags edges to Tag.
func (tu *TargetUpdate) AddTags(t ...*Tag) *TargetUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTagIDs(ids...)
}

// AddCredentialIDs adds the credentials edge to Credential by ids.
func (tu *TargetUpdate) AddCredentialIDs(ids ...int) *TargetUpdate {
	if tu.credentials == nil {
		tu.credentials = make(map[int]struct{})
	}
	for i := range ids {
		tu.credentials[ids[i]] = struct{}{}
	}
	return tu
}

// AddCredentials adds the credentials edges to Credential.
func (tu *TargetUpdate) AddCredentials(c ...*Credential) *TargetUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.AddCredentialIDs(ids...)
}

// RemoveTaskIDs removes the tasks edge to Task by ids.
func (tu *TargetUpdate) RemoveTaskIDs(ids ...int) *TargetUpdate {
	if tu.removedTasks == nil {
		tu.removedTasks = make(map[int]struct{})
	}
	for i := range ids {
		tu.removedTasks[ids[i]] = struct{}{}
	}
	return tu
}

// RemoveTasks removes tasks edges to Task.
func (tu *TargetUpdate) RemoveTasks(t ...*Task) *TargetUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTaskIDs(ids...)
}

// RemoveTagIDs removes the tags edge to Tag by ids.
func (tu *TargetUpdate) RemoveTagIDs(ids ...int) *TargetUpdate {
	if tu.removedTags == nil {
		tu.removedTags = make(map[int]struct{})
	}
	for i := range ids {
		tu.removedTags[ids[i]] = struct{}{}
	}
	return tu
}

// RemoveTags removes tags edges to Tag.
func (tu *TargetUpdate) RemoveTags(t ...*Tag) *TargetUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTagIDs(ids...)
}

// RemoveCredentialIDs removes the credentials edge to Credential by ids.
func (tu *TargetUpdate) RemoveCredentialIDs(ids ...int) *TargetUpdate {
	if tu.removedCredentials == nil {
		tu.removedCredentials = make(map[int]struct{})
	}
	for i := range ids {
		tu.removedCredentials[ids[i]] = struct{}{}
	}
	return tu
}

// RemoveCredentials removes credentials edges to Credential.
func (tu *TargetUpdate) RemoveCredentials(c ...*Credential) *TargetUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.RemoveCredentialIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TargetUpdate) Save(ctx context.Context) (int, error) {
	if tu.MachineUUID != nil {
		if err := target.MachineUUIDValidator(*tu.MachineUUID); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"MachineUUID\": %v", err)
		}
	}
	return tu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TargetUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TargetUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TargetUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TargetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   target.Table,
			Columns: target.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: target.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := tu.Name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldName,
		})
	}
	if value := tu.PrimaryIP; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldPrimaryIP,
		})
	}
	if value := tu.MachineUUID; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldMachineUUID,
		})
	}
	if tu.clearMachineUUID {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: target.FieldMachineUUID,
		})
	}
	if value := tu.PublicIP; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldPublicIP,
		})
	}
	if tu.clearPublicIP {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: target.FieldPublicIP,
		})
	}
	if value := tu.PrimaryMAC; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldPrimaryMAC,
		})
	}
	if tu.clearPrimaryMAC {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: target.FieldPrimaryMAC,
		})
	}
	if value := tu.Hostname; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldHostname,
		})
	}
	if tu.clearHostname {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: target.FieldHostname,
		})
	}
	if value := tu.LastSeen; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: target.FieldLastSeen,
		})
	}
	if tu.clearLastSeen {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: target.FieldLastSeen,
		})
	}
	if nodes := tu.removedTasks; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   target.TasksTable,
			Columns: []string{target.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.tasks; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   target.TasksTable,
			Columns: []string{target.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := tu.removedTags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   target.TagsTable,
			Columns: target.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.tags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   target.TagsTable,
			Columns: target.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := tu.removedCredentials; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   target.CredentialsTable,
			Columns: []string{target.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: credential.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.credentials; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   target.CredentialsTable,
			Columns: []string{target.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: credential.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TargetUpdateOne is the builder for updating a single Target entity.
type TargetUpdateOne struct {
	config
	id                 int
	Name               *string
	PrimaryIP          *string
	MachineUUID        *string
	clearMachineUUID   bool
	PublicIP           *string
	clearPublicIP      bool
	PrimaryMAC         *string
	clearPrimaryMAC    bool
	Hostname           *string
	clearHostname      bool
	LastSeen           *time.Time
	clearLastSeen      bool
	tasks              map[int]struct{}
	tags               map[int]struct{}
	credentials        map[int]struct{}
	removedTasks       map[int]struct{}
	removedTags        map[int]struct{}
	removedCredentials map[int]struct{}
}

// SetName sets the Name field.
func (tuo *TargetUpdateOne) SetName(s string) *TargetUpdateOne {
	tuo.Name = &s
	return tuo
}

// SetPrimaryIP sets the PrimaryIP field.
func (tuo *TargetUpdateOne) SetPrimaryIP(s string) *TargetUpdateOne {
	tuo.PrimaryIP = &s
	return tuo
}

// SetMachineUUID sets the MachineUUID field.
func (tuo *TargetUpdateOne) SetMachineUUID(s string) *TargetUpdateOne {
	tuo.MachineUUID = &s
	return tuo
}

// SetNillableMachineUUID sets the MachineUUID field if the given value is not nil.
func (tuo *TargetUpdateOne) SetNillableMachineUUID(s *string) *TargetUpdateOne {
	if s != nil {
		tuo.SetMachineUUID(*s)
	}
	return tuo
}

// ClearMachineUUID clears the value of MachineUUID.
func (tuo *TargetUpdateOne) ClearMachineUUID() *TargetUpdateOne {
	tuo.MachineUUID = nil
	tuo.clearMachineUUID = true
	return tuo
}

// SetPublicIP sets the PublicIP field.
func (tuo *TargetUpdateOne) SetPublicIP(s string) *TargetUpdateOne {
	tuo.PublicIP = &s
	return tuo
}

// SetNillablePublicIP sets the PublicIP field if the given value is not nil.
func (tuo *TargetUpdateOne) SetNillablePublicIP(s *string) *TargetUpdateOne {
	if s != nil {
		tuo.SetPublicIP(*s)
	}
	return tuo
}

// ClearPublicIP clears the value of PublicIP.
func (tuo *TargetUpdateOne) ClearPublicIP() *TargetUpdateOne {
	tuo.PublicIP = nil
	tuo.clearPublicIP = true
	return tuo
}

// SetPrimaryMAC sets the PrimaryMAC field.
func (tuo *TargetUpdateOne) SetPrimaryMAC(s string) *TargetUpdateOne {
	tuo.PrimaryMAC = &s
	return tuo
}

// SetNillablePrimaryMAC sets the PrimaryMAC field if the given value is not nil.
func (tuo *TargetUpdateOne) SetNillablePrimaryMAC(s *string) *TargetUpdateOne {
	if s != nil {
		tuo.SetPrimaryMAC(*s)
	}
	return tuo
}

// ClearPrimaryMAC clears the value of PrimaryMAC.
func (tuo *TargetUpdateOne) ClearPrimaryMAC() *TargetUpdateOne {
	tuo.PrimaryMAC = nil
	tuo.clearPrimaryMAC = true
	return tuo
}

// SetHostname sets the Hostname field.
func (tuo *TargetUpdateOne) SetHostname(s string) *TargetUpdateOne {
	tuo.Hostname = &s
	return tuo
}

// SetNillableHostname sets the Hostname field if the given value is not nil.
func (tuo *TargetUpdateOne) SetNillableHostname(s *string) *TargetUpdateOne {
	if s != nil {
		tuo.SetHostname(*s)
	}
	return tuo
}

// ClearHostname clears the value of Hostname.
func (tuo *TargetUpdateOne) ClearHostname() *TargetUpdateOne {
	tuo.Hostname = nil
	tuo.clearHostname = true
	return tuo
}

// SetLastSeen sets the LastSeen field.
func (tuo *TargetUpdateOne) SetLastSeen(t time.Time) *TargetUpdateOne {
	tuo.LastSeen = &t
	return tuo
}

// SetNillableLastSeen sets the LastSeen field if the given value is not nil.
func (tuo *TargetUpdateOne) SetNillableLastSeen(t *time.Time) *TargetUpdateOne {
	if t != nil {
		tuo.SetLastSeen(*t)
	}
	return tuo
}

// ClearLastSeen clears the value of LastSeen.
func (tuo *TargetUpdateOne) ClearLastSeen() *TargetUpdateOne {
	tuo.LastSeen = nil
	tuo.clearLastSeen = true
	return tuo
}

// AddTaskIDs adds the tasks edge to Task by ids.
func (tuo *TargetUpdateOne) AddTaskIDs(ids ...int) *TargetUpdateOne {
	if tuo.tasks == nil {
		tuo.tasks = make(map[int]struct{})
	}
	for i := range ids {
		tuo.tasks[ids[i]] = struct{}{}
	}
	return tuo
}

// AddTasks adds the tasks edges to Task.
func (tuo *TargetUpdateOne) AddTasks(t ...*Task) *TargetUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTaskIDs(ids...)
}

// AddTagIDs adds the tags edge to Tag by ids.
func (tuo *TargetUpdateOne) AddTagIDs(ids ...int) *TargetUpdateOne {
	if tuo.tags == nil {
		tuo.tags = make(map[int]struct{})
	}
	for i := range ids {
		tuo.tags[ids[i]] = struct{}{}
	}
	return tuo
}

// AddTags adds the tags edges to Tag.
func (tuo *TargetUpdateOne) AddTags(t ...*Tag) *TargetUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTagIDs(ids...)
}

// AddCredentialIDs adds the credentials edge to Credential by ids.
func (tuo *TargetUpdateOne) AddCredentialIDs(ids ...int) *TargetUpdateOne {
	if tuo.credentials == nil {
		tuo.credentials = make(map[int]struct{})
	}
	for i := range ids {
		tuo.credentials[ids[i]] = struct{}{}
	}
	return tuo
}

// AddCredentials adds the credentials edges to Credential.
func (tuo *TargetUpdateOne) AddCredentials(c ...*Credential) *TargetUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.AddCredentialIDs(ids...)
}

// RemoveTaskIDs removes the tasks edge to Task by ids.
func (tuo *TargetUpdateOne) RemoveTaskIDs(ids ...int) *TargetUpdateOne {
	if tuo.removedTasks == nil {
		tuo.removedTasks = make(map[int]struct{})
	}
	for i := range ids {
		tuo.removedTasks[ids[i]] = struct{}{}
	}
	return tuo
}

// RemoveTasks removes tasks edges to Task.
func (tuo *TargetUpdateOne) RemoveTasks(t ...*Task) *TargetUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTaskIDs(ids...)
}

// RemoveTagIDs removes the tags edge to Tag by ids.
func (tuo *TargetUpdateOne) RemoveTagIDs(ids ...int) *TargetUpdateOne {
	if tuo.removedTags == nil {
		tuo.removedTags = make(map[int]struct{})
	}
	for i := range ids {
		tuo.removedTags[ids[i]] = struct{}{}
	}
	return tuo
}

// RemoveTags removes tags edges to Tag.
func (tuo *TargetUpdateOne) RemoveTags(t ...*Tag) *TargetUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTagIDs(ids...)
}

// RemoveCredentialIDs removes the credentials edge to Credential by ids.
func (tuo *TargetUpdateOne) RemoveCredentialIDs(ids ...int) *TargetUpdateOne {
	if tuo.removedCredentials == nil {
		tuo.removedCredentials = make(map[int]struct{})
	}
	for i := range ids {
		tuo.removedCredentials[ids[i]] = struct{}{}
	}
	return tuo
}

// RemoveCredentials removes credentials edges to Credential.
func (tuo *TargetUpdateOne) RemoveCredentials(c ...*Credential) *TargetUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.RemoveCredentialIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (tuo *TargetUpdateOne) Save(ctx context.Context) (*Target, error) {
	if tuo.MachineUUID != nil {
		if err := target.MachineUUIDValidator(*tuo.MachineUUID); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"MachineUUID\": %v", err)
		}
	}
	return tuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TargetUpdateOne) SaveX(ctx context.Context) *Target {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TargetUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TargetUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TargetUpdateOne) sqlSave(ctx context.Context) (t *Target, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   target.Table,
			Columns: target.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  tuo.id,
				Type:   field.TypeInt,
				Column: target.FieldID,
			},
		},
	}
	if value := tuo.Name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldName,
		})
	}
	if value := tuo.PrimaryIP; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldPrimaryIP,
		})
	}
	if value := tuo.MachineUUID; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldMachineUUID,
		})
	}
	if tuo.clearMachineUUID {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: target.FieldMachineUUID,
		})
	}
	if value := tuo.PublicIP; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldPublicIP,
		})
	}
	if tuo.clearPublicIP {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: target.FieldPublicIP,
		})
	}
	if value := tuo.PrimaryMAC; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldPrimaryMAC,
		})
	}
	if tuo.clearPrimaryMAC {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: target.FieldPrimaryMAC,
		})
	}
	if value := tuo.Hostname; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: target.FieldHostname,
		})
	}
	if tuo.clearHostname {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: target.FieldHostname,
		})
	}
	if value := tuo.LastSeen; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: target.FieldLastSeen,
		})
	}
	if tuo.clearLastSeen {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: target.FieldLastSeen,
		})
	}
	if nodes := tuo.removedTasks; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   target.TasksTable,
			Columns: []string{target.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.tasks; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   target.TasksTable,
			Columns: []string{target.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := tuo.removedTags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   target.TagsTable,
			Columns: target.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.tags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   target.TagsTable,
			Columns: target.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := tuo.removedCredentials; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   target.CredentialsTable,
			Columns: []string{target.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: credential.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.credentials; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   target.CredentialsTable,
			Columns: []string{target.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: credential.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	t = &Target{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
