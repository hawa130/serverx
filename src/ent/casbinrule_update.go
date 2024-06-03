// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hawa130/computility-cloud/ent/casbinrule"
	"github.com/hawa130/computility-cloud/ent/predicate"
)

// CasbinRuleUpdate is the builder for updating CasbinRule entities.
type CasbinRuleUpdate struct {
	config
	hooks    []Hook
	mutation *CasbinRuleMutation
}

// Where appends a list predicates to the CasbinRuleUpdate builder.
func (cru *CasbinRuleUpdate) Where(ps ...predicate.CasbinRule) *CasbinRuleUpdate {
	cru.mutation.Where(ps...)
	return cru
}

// SetPtype sets the "Ptype" field.
func (cru *CasbinRuleUpdate) SetPtype(s string) *CasbinRuleUpdate {
	cru.mutation.SetPtype(s)
	return cru
}

// SetNillablePtype sets the "Ptype" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillablePtype(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetPtype(*s)
	}
	return cru
}

// SetV0 sets the "V0" field.
func (cru *CasbinRuleUpdate) SetV0(s string) *CasbinRuleUpdate {
	cru.mutation.SetV0(s)
	return cru
}

// SetNillableV0 sets the "V0" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV0(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV0(*s)
	}
	return cru
}

// SetV1 sets the "V1" field.
func (cru *CasbinRuleUpdate) SetV1(s string) *CasbinRuleUpdate {
	cru.mutation.SetV1(s)
	return cru
}

// SetNillableV1 sets the "V1" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV1(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV1(*s)
	}
	return cru
}

// SetV2 sets the "V2" field.
func (cru *CasbinRuleUpdate) SetV2(s string) *CasbinRuleUpdate {
	cru.mutation.SetV2(s)
	return cru
}

// SetNillableV2 sets the "V2" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV2(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV2(*s)
	}
	return cru
}

// SetV3 sets the "V3" field.
func (cru *CasbinRuleUpdate) SetV3(s string) *CasbinRuleUpdate {
	cru.mutation.SetV3(s)
	return cru
}

// SetNillableV3 sets the "V3" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV3(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV3(*s)
	}
	return cru
}

// SetV4 sets the "V4" field.
func (cru *CasbinRuleUpdate) SetV4(s string) *CasbinRuleUpdate {
	cru.mutation.SetV4(s)
	return cru
}

// SetNillableV4 sets the "V4" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV4(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV4(*s)
	}
	return cru
}

// SetV5 sets the "V5" field.
func (cru *CasbinRuleUpdate) SetV5(s string) *CasbinRuleUpdate {
	cru.mutation.SetV5(s)
	return cru
}

// SetNillableV5 sets the "V5" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV5(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV5(*s)
	}
	return cru
}

// Mutation returns the CasbinRuleMutation object of the builder.
func (cru *CasbinRuleUpdate) Mutation() *CasbinRuleMutation {
	return cru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *CasbinRuleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cru.sqlSave, cru.mutation, cru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cru *CasbinRuleUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *CasbinRuleUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *CasbinRuleUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cru *CasbinRuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(casbinrule.Table, casbinrule.Columns, sqlgraph.NewFieldSpec(casbinrule.FieldID, field.TypeString))
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.Ptype(); ok {
		_spec.SetField(casbinrule.FieldPtype, field.TypeString, value)
	}
	if value, ok := cru.mutation.V0(); ok {
		_spec.SetField(casbinrule.FieldV0, field.TypeString, value)
	}
	if value, ok := cru.mutation.V1(); ok {
		_spec.SetField(casbinrule.FieldV1, field.TypeString, value)
	}
	if value, ok := cru.mutation.V2(); ok {
		_spec.SetField(casbinrule.FieldV2, field.TypeString, value)
	}
	if value, ok := cru.mutation.V3(); ok {
		_spec.SetField(casbinrule.FieldV3, field.TypeString, value)
	}
	if value, ok := cru.mutation.V4(); ok {
		_spec.SetField(casbinrule.FieldV4, field.TypeString, value)
	}
	if value, ok := cru.mutation.V5(); ok {
		_spec.SetField(casbinrule.FieldV5, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{casbinrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cru.mutation.done = true
	return n, nil
}

// CasbinRuleUpdateOne is the builder for updating a single CasbinRule entity.
type CasbinRuleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CasbinRuleMutation
}

// SetPtype sets the "Ptype" field.
func (cruo *CasbinRuleUpdateOne) SetPtype(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetPtype(s)
	return cruo
}

// SetNillablePtype sets the "Ptype" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillablePtype(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetPtype(*s)
	}
	return cruo
}

// SetV0 sets the "V0" field.
func (cruo *CasbinRuleUpdateOne) SetV0(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV0(s)
	return cruo
}

// SetNillableV0 sets the "V0" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV0(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV0(*s)
	}
	return cruo
}

// SetV1 sets the "V1" field.
func (cruo *CasbinRuleUpdateOne) SetV1(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV1(s)
	return cruo
}

// SetNillableV1 sets the "V1" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV1(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV1(*s)
	}
	return cruo
}

// SetV2 sets the "V2" field.
func (cruo *CasbinRuleUpdateOne) SetV2(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV2(s)
	return cruo
}

// SetNillableV2 sets the "V2" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV2(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV2(*s)
	}
	return cruo
}

// SetV3 sets the "V3" field.
func (cruo *CasbinRuleUpdateOne) SetV3(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV3(s)
	return cruo
}

// SetNillableV3 sets the "V3" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV3(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV3(*s)
	}
	return cruo
}

// SetV4 sets the "V4" field.
func (cruo *CasbinRuleUpdateOne) SetV4(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV4(s)
	return cruo
}

// SetNillableV4 sets the "V4" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV4(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV4(*s)
	}
	return cruo
}

// SetV5 sets the "V5" field.
func (cruo *CasbinRuleUpdateOne) SetV5(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV5(s)
	return cruo
}

// SetNillableV5 sets the "V5" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV5(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV5(*s)
	}
	return cruo
}

// Mutation returns the CasbinRuleMutation object of the builder.
func (cruo *CasbinRuleUpdateOne) Mutation() *CasbinRuleMutation {
	return cruo.mutation
}

// Where appends a list predicates to the CasbinRuleUpdate builder.
func (cruo *CasbinRuleUpdateOne) Where(ps ...predicate.CasbinRule) *CasbinRuleUpdateOne {
	cruo.mutation.Where(ps...)
	return cruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cruo *CasbinRuleUpdateOne) Select(field string, fields ...string) *CasbinRuleUpdateOne {
	cruo.fields = append([]string{field}, fields...)
	return cruo
}

// Save executes the query and returns the updated CasbinRule entity.
func (cruo *CasbinRuleUpdateOne) Save(ctx context.Context) (*CasbinRule, error) {
	return withHooks(ctx, cruo.sqlSave, cruo.mutation, cruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *CasbinRuleUpdateOne) SaveX(ctx context.Context) *CasbinRule {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *CasbinRuleUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *CasbinRuleUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cruo *CasbinRuleUpdateOne) sqlSave(ctx context.Context) (_node *CasbinRule, err error) {
	_spec := sqlgraph.NewUpdateSpec(casbinrule.Table, casbinrule.Columns, sqlgraph.NewFieldSpec(casbinrule.FieldID, field.TypeString))
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CasbinRule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, casbinrule.FieldID)
		for _, f := range fields {
			if !casbinrule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != casbinrule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.Ptype(); ok {
		_spec.SetField(casbinrule.FieldPtype, field.TypeString, value)
	}
	if value, ok := cruo.mutation.V0(); ok {
		_spec.SetField(casbinrule.FieldV0, field.TypeString, value)
	}
	if value, ok := cruo.mutation.V1(); ok {
		_spec.SetField(casbinrule.FieldV1, field.TypeString, value)
	}
	if value, ok := cruo.mutation.V2(); ok {
		_spec.SetField(casbinrule.FieldV2, field.TypeString, value)
	}
	if value, ok := cruo.mutation.V3(); ok {
		_spec.SetField(casbinrule.FieldV3, field.TypeString, value)
	}
	if value, ok := cruo.mutation.V4(); ok {
		_spec.SetField(casbinrule.FieldV4, field.TypeString, value)
	}
	if value, ok := cruo.mutation.V5(); ok {
		_spec.SetField(casbinrule.FieldV5, field.TypeString, value)
	}
	_node = &CasbinRule{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{casbinrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cruo.mutation.done = true
	return _node, nil
}
