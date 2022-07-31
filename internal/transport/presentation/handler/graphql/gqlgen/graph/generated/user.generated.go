// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"strconv"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/icaroribeiro/new-go-code-challenge-template-2/internal/transport/presentation/handler/graphql/gqlgen/graph/entity"
	"github.com/satori/go.uuid"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _User_id(ctx context.Context, field graphql.CollectedField, obj *entity.User) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_User_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(uuid.UUID)
	fc.Result = res
	return ec.marshalNUUID2githubᚗcomᚋsatoriᚋgoᚗuuidᚐUUID(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_User_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "User",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type UUID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _User_username(ctx context.Context, field graphql.CollectedField, obj *entity.User) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_User_username(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Username, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_User_username(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "User",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var userImplementors = []string{"User"}

func (ec *executionContext) _User(ctx context.Context, sel ast.SelectionSet, obj *entity.User) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, userImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("User")
		case "id":

			out.Values[i] = ec._User_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "username":

			out.Values[i] = ec._User_username(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNUser2ᚕᚖgithubᚗcomᚋicaroribeiroᚋnewᚑgoᚑcodeᚑchallengeᚑtemplateᚑ2ᚋinternalᚋtransportᚋpresentationᚋhandlerᚋgraphqlᚋgqlgenᚋgraphᚋentityᚐUserᚄ(ctx context.Context, sel ast.SelectionSet, v []*entity.User) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNUser2ᚖgithubᚗcomᚋicaroribeiroᚋnewᚑgoᚑcodeᚑchallengeᚑtemplateᚑ2ᚋinternalᚋtransportᚋpresentationᚋhandlerᚋgraphqlᚋgqlgenᚋgraphᚋentityᚐUser(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNUser2ᚖgithubᚗcomᚋicaroribeiroᚋnewᚑgoᚑcodeᚑchallengeᚑtemplateᚑ2ᚋinternalᚋtransportᚋpresentationᚋhandlerᚋgraphqlᚋgqlgenᚋgraphᚋentityᚐUser(ctx context.Context, sel ast.SelectionSet, v *entity.User) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._User(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
