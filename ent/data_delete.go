// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/choonhong/hotel-data-merge/ent/data"
	"github.com/choonhong/hotel-data-merge/ent/predicate"
)

// DataDelete is the builder for deleting a Data entity.
type DataDelete struct {
	config
	hooks    []Hook
	mutation *DataMutation
}

// Where appends a list predicates to the DataDelete builder.
func (dd *DataDelete) Where(ps ...predicate.Data) *DataDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DataDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(data.Table, sqlgraph.NewFieldSpec(data.FieldID, field.TypeString))
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DataDeleteOne is the builder for deleting a single Data entity.
type DataDeleteOne struct {
	dd *DataDelete
}

// Where appends a list predicates to the DataDelete builder.
func (ddo *DataDeleteOne) Where(ps ...predicate.Data) *DataDeleteOne {
	ddo.dd.mutation.Where(ps...)
	return ddo
}

// Exec executes the deletion query.
func (ddo *DataDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{data.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DataDeleteOne) ExecX(ctx context.Context) {
	if err := ddo.Exec(ctx); err != nil {
		panic(err)
	}
}