// Code generated by ent, DO NOT EDIT.

package ogent

import (
	"context"
	"net/http"

	"github.com/go-faster/jx"
	"github.com/ryota0624/go-dev-with-ent/ent"
	"github.com/ryota0624/go-dev-with-ent/ent/user"
)

// OgentHandler implements the ogen generated Handler interface and uses Ent as data layer.
type OgentHandler struct {
	client *ent.Client
}

// NewOgentHandler returns a new OgentHandler.
func NewOgentHandler(c *ent.Client) *OgentHandler { return &OgentHandler{c} }

// rawError renders err as json string.
func rawError(err error) jx.Raw {
	var e jx.Encoder
	e.Str(err.Error())
	return e.Bytes()
}

// CreateUser handles POST /users requests.
func (h *OgentHandler) CreateUser(ctx context.Context, req CreateUserReq) (CreateUserRes, error) {
	b := h.client.User.Create()
	// Add all fields.
	b.SetAge(req.Age)
	b.SetName(req.Name)
	// Add all edges.
	b.AddCarIDs(req.Cars...)
	b.AddGroupIDs(req.Groups...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.User.Query().Where(user.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewUserCreate(e), nil
}

// ListUser handles GET /users requests.
func (h *OgentHandler) ListUser(ctx context.Context, params ListUserParams) (ListUserRes, error) {
	q := h.client.User.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewUserLists(es)
	return (*ListUserOKApplicationJSON)(&r), nil
}

// ReadUser handles GET /users/{id} requests.
func (h *OgentHandler) ReadUser(ctx context.Context, params ReadUserParams) (ReadUserRes, error) {
	q := h.client.User.Query().Where(user.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewUserRead(e), nil
}

// UpdateUser handles PATCH /users/{id} requests.
func (h *OgentHandler) UpdateUser(ctx context.Context, req UpdateUserReq, params UpdateUserParams) (UpdateUserRes, error) {
	b := h.client.User.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Age.Get(); ok {
		b.SetAge(v)
	}
	if v, ok := req.Name.Get(); ok {
		b.SetName(v)
	}
	// Add all edges.
	b.ClearCars().AddCarIDs(req.Cars...)
	b.ClearGroups().AddGroupIDs(req.Groups...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.User.Query().Where(user.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewUserUpdate(e), nil
}
