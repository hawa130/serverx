// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hawa130/computility-cloud/ent/casbinrule"
	"github.com/hawa130/computility-cloud/ent/predicate"
)

// CasbinRuleDelete is the builder for deleting a CasbinRule entity.
type CasbinRuleDelete struct {
	config
	hooks    []Hook
	mutation *CasbinRuleMutation
}

// Where appends a list predicates to the CasbinRuleDelete builder.
func (crd *CasbinRuleDelete) Where(ps ...predicate.CasbinRule) *CasbinRuleDelete {
	crd.mutation.Where(ps...)
	return crd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crd *CasbinRuleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, crd.sqlExec, crd.mutation, crd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (crd *CasbinRuleDelete) ExecX(ctx context.Context) int {
	n, err := crd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crd *CasbinRuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(casbinrule.Table, sqlgraph.NewFieldSpec(casbinrule.FieldID, field.TypeString))
	if ps := crd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	crd.mutation.done = true
	return affected, err
}

// CasbinRuleDeleteOne is the builder for deleting a single CasbinRule entity.
type CasbinRuleDeleteOne struct {
	crd *CasbinRuleDelete
}

// Where appends a list predicates to the CasbinRuleDelete builder.
func (crdo *CasbinRuleDeleteOne) Where(ps ...predicate.CasbinRule) *CasbinRuleDeleteOne {
	crdo.crd.mutation.Where(ps...)
	return crdo
}

// Exec executes the deletion query.
func (crdo *CasbinRuleDeleteOne) Exec(ctx context.Context) error {
	n, err := crdo.crd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{casbinrule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crdo *CasbinRuleDeleteOne) ExecX(ctx context.Context) {
	if err := crdo.Exec(ctx); err != nil {
		panic(err)
	}
}
