// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"msp/biz_server/gpt/internal/infra/mysql/model/ent/chatmessage"
	"msp/biz_server/gpt/internal/infra/mysql/model/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChatMessageUpdate is the builder for updating ChatMessage entities.
type ChatMessageUpdate struct {
	config
	hooks    []Hook
	mutation *ChatMessageMutation
}

// Where appends a list predicates to the ChatMessageUpdate builder.
func (cmu *ChatMessageUpdate) Where(ps ...predicate.ChatMessage) *ChatMessageUpdate {
	cmu.mutation.Where(ps...)
	return cmu
}

// SetUpdatedAt sets the "updated_at" field.
func (cmu *ChatMessageUpdate) SetUpdatedAt(t time.Time) *ChatMessageUpdate {
	cmu.mutation.SetUpdatedAt(t)
	return cmu
}

// SetChatID sets the "chat_id" field.
func (cmu *ChatMessageUpdate) SetChatID(i int32) *ChatMessageUpdate {
	cmu.mutation.ResetChatID()
	cmu.mutation.SetChatID(i)
	return cmu
}

// SetNillableChatID sets the "chat_id" field if the given value is not nil.
func (cmu *ChatMessageUpdate) SetNillableChatID(i *int32) *ChatMessageUpdate {
	if i != nil {
		cmu.SetChatID(*i)
	}
	return cmu
}

// AddChatID adds i to the "chat_id" field.
func (cmu *ChatMessageUpdate) AddChatID(i int32) *ChatMessageUpdate {
	cmu.mutation.AddChatID(i)
	return cmu
}

// SetRequestID sets the "request_id" field.
func (cmu *ChatMessageUpdate) SetRequestID(s string) *ChatMessageUpdate {
	cmu.mutation.SetRequestID(s)
	return cmu
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (cmu *ChatMessageUpdate) SetNillableRequestID(s *string) *ChatMessageUpdate {
	if s != nil {
		cmu.SetRequestID(*s)
	}
	return cmu
}

// SetText sets the "text" field.
func (cmu *ChatMessageUpdate) SetText(s string) *ChatMessageUpdate {
	cmu.mutation.SetText(s)
	return cmu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (cmu *ChatMessageUpdate) SetNillableText(s *string) *ChatMessageUpdate {
	if s != nil {
		cmu.SetText(*s)
	}
	return cmu
}

// SetVersion sets the "version" field.
func (cmu *ChatMessageUpdate) SetVersion(i int8) *ChatMessageUpdate {
	cmu.mutation.ResetVersion()
	cmu.mutation.SetVersion(i)
	return cmu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (cmu *ChatMessageUpdate) SetNillableVersion(i *int8) *ChatMessageUpdate {
	if i != nil {
		cmu.SetVersion(*i)
	}
	return cmu
}

// AddVersion adds i to the "version" field.
func (cmu *ChatMessageUpdate) AddVersion(i int8) *ChatMessageUpdate {
	cmu.mutation.AddVersion(i)
	return cmu
}

// Mutation returns the ChatMessageMutation object of the builder.
func (cmu *ChatMessageUpdate) Mutation() *ChatMessageMutation {
	return cmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cmu *ChatMessageUpdate) Save(ctx context.Context) (int, error) {
	cmu.defaults()
	return withHooks(ctx, cmu.sqlSave, cmu.mutation, cmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cmu *ChatMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := cmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cmu *ChatMessageUpdate) Exec(ctx context.Context) error {
	_, err := cmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmu *ChatMessageUpdate) ExecX(ctx context.Context) {
	if err := cmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cmu *ChatMessageUpdate) defaults() {
	if _, ok := cmu.mutation.UpdatedAt(); !ok {
		v := chatmessage.UpdateDefaultUpdatedAt()
		cmu.mutation.SetUpdatedAt(v)
	}
}

func (cmu *ChatMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(chatmessage.Table, chatmessage.Columns, sqlgraph.NewFieldSpec(chatmessage.FieldID, field.TypeInt32))
	if ps := cmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmu.mutation.UpdatedAt(); ok {
		_spec.SetField(chatmessage.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cmu.mutation.ChatID(); ok {
		_spec.SetField(chatmessage.FieldChatID, field.TypeInt32, value)
	}
	if value, ok := cmu.mutation.AddedChatID(); ok {
		_spec.AddField(chatmessage.FieldChatID, field.TypeInt32, value)
	}
	if value, ok := cmu.mutation.RequestID(); ok {
		_spec.SetField(chatmessage.FieldRequestID, field.TypeString, value)
	}
	if value, ok := cmu.mutation.Text(); ok {
		_spec.SetField(chatmessage.FieldText, field.TypeString, value)
	}
	if value, ok := cmu.mutation.Version(); ok {
		_spec.SetField(chatmessage.FieldVersion, field.TypeInt8, value)
	}
	if value, ok := cmu.mutation.AddedVersion(); ok {
		_spec.AddField(chatmessage.FieldVersion, field.TypeInt8, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chatmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cmu.mutation.done = true
	return n, nil
}

// ChatMessageUpdateOne is the builder for updating a single ChatMessage entity.
type ChatMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChatMessageMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cmuo *ChatMessageUpdateOne) SetUpdatedAt(t time.Time) *ChatMessageUpdateOne {
	cmuo.mutation.SetUpdatedAt(t)
	return cmuo
}

// SetChatID sets the "chat_id" field.
func (cmuo *ChatMessageUpdateOne) SetChatID(i int32) *ChatMessageUpdateOne {
	cmuo.mutation.ResetChatID()
	cmuo.mutation.SetChatID(i)
	return cmuo
}

// SetNillableChatID sets the "chat_id" field if the given value is not nil.
func (cmuo *ChatMessageUpdateOne) SetNillableChatID(i *int32) *ChatMessageUpdateOne {
	if i != nil {
		cmuo.SetChatID(*i)
	}
	return cmuo
}

// AddChatID adds i to the "chat_id" field.
func (cmuo *ChatMessageUpdateOne) AddChatID(i int32) *ChatMessageUpdateOne {
	cmuo.mutation.AddChatID(i)
	return cmuo
}

// SetRequestID sets the "request_id" field.
func (cmuo *ChatMessageUpdateOne) SetRequestID(s string) *ChatMessageUpdateOne {
	cmuo.mutation.SetRequestID(s)
	return cmuo
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (cmuo *ChatMessageUpdateOne) SetNillableRequestID(s *string) *ChatMessageUpdateOne {
	if s != nil {
		cmuo.SetRequestID(*s)
	}
	return cmuo
}

// SetText sets the "text" field.
func (cmuo *ChatMessageUpdateOne) SetText(s string) *ChatMessageUpdateOne {
	cmuo.mutation.SetText(s)
	return cmuo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (cmuo *ChatMessageUpdateOne) SetNillableText(s *string) *ChatMessageUpdateOne {
	if s != nil {
		cmuo.SetText(*s)
	}
	return cmuo
}

// SetVersion sets the "version" field.
func (cmuo *ChatMessageUpdateOne) SetVersion(i int8) *ChatMessageUpdateOne {
	cmuo.mutation.ResetVersion()
	cmuo.mutation.SetVersion(i)
	return cmuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (cmuo *ChatMessageUpdateOne) SetNillableVersion(i *int8) *ChatMessageUpdateOne {
	if i != nil {
		cmuo.SetVersion(*i)
	}
	return cmuo
}

// AddVersion adds i to the "version" field.
func (cmuo *ChatMessageUpdateOne) AddVersion(i int8) *ChatMessageUpdateOne {
	cmuo.mutation.AddVersion(i)
	return cmuo
}

// Mutation returns the ChatMessageMutation object of the builder.
func (cmuo *ChatMessageUpdateOne) Mutation() *ChatMessageMutation {
	return cmuo.mutation
}

// Where appends a list predicates to the ChatMessageUpdate builder.
func (cmuo *ChatMessageUpdateOne) Where(ps ...predicate.ChatMessage) *ChatMessageUpdateOne {
	cmuo.mutation.Where(ps...)
	return cmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cmuo *ChatMessageUpdateOne) Select(field string, fields ...string) *ChatMessageUpdateOne {
	cmuo.fields = append([]string{field}, fields...)
	return cmuo
}

// Save executes the query and returns the updated ChatMessage entity.
func (cmuo *ChatMessageUpdateOne) Save(ctx context.Context) (*ChatMessage, error) {
	cmuo.defaults()
	return withHooks(ctx, cmuo.sqlSave, cmuo.mutation, cmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cmuo *ChatMessageUpdateOne) SaveX(ctx context.Context) *ChatMessage {
	node, err := cmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cmuo *ChatMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := cmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmuo *ChatMessageUpdateOne) ExecX(ctx context.Context) {
	if err := cmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cmuo *ChatMessageUpdateOne) defaults() {
	if _, ok := cmuo.mutation.UpdatedAt(); !ok {
		v := chatmessage.UpdateDefaultUpdatedAt()
		cmuo.mutation.SetUpdatedAt(v)
	}
}

func (cmuo *ChatMessageUpdateOne) sqlSave(ctx context.Context) (_node *ChatMessage, err error) {
	_spec := sqlgraph.NewUpdateSpec(chatmessage.Table, chatmessage.Columns, sqlgraph.NewFieldSpec(chatmessage.FieldID, field.TypeInt32))
	id, ok := cmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ChatMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chatmessage.FieldID)
		for _, f := range fields {
			if !chatmessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chatmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(chatmessage.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cmuo.mutation.ChatID(); ok {
		_spec.SetField(chatmessage.FieldChatID, field.TypeInt32, value)
	}
	if value, ok := cmuo.mutation.AddedChatID(); ok {
		_spec.AddField(chatmessage.FieldChatID, field.TypeInt32, value)
	}
	if value, ok := cmuo.mutation.RequestID(); ok {
		_spec.SetField(chatmessage.FieldRequestID, field.TypeString, value)
	}
	if value, ok := cmuo.mutation.Text(); ok {
		_spec.SetField(chatmessage.FieldText, field.TypeString, value)
	}
	if value, ok := cmuo.mutation.Version(); ok {
		_spec.SetField(chatmessage.FieldVersion, field.TypeInt8, value)
	}
	if value, ok := cmuo.mutation.AddedVersion(); ok {
		_spec.AddField(chatmessage.FieldVersion, field.TypeInt8, value)
	}
	_node = &ChatMessage{config: cmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chatmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cmuo.mutation.done = true
	return _node, nil
}
