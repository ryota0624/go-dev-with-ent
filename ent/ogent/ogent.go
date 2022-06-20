// Code generated by entc, DO NOT EDIT.

package ogent

import (
	"context"
	"net/http"

	"github.com/go-faster/jx"
	"github.com/ryota0624/go-dev-with-ent/ent"
	"github.com/ryota0624/go-dev-with-ent/ent/car"
	"github.com/ryota0624/go-dev-with-ent/ent/group"
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

// CreateCar handles POST /cars requests.
func (h *OgentHandler) CreateCar(ctx context.Context, req CreateCarReq) (CreateCarRes, error) {
	b := h.client.Car.Create()
	// Add all fields.
	b.SetModel(req.Model)
	b.SetRegisteredAt(req.RegisteredAt)
	// Add all edges.
	if v, ok := req.Owner.Get(); ok {
		b.SetOwnerID(v)
	}
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
	q := h.client.Car.Query().Where(car.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewCarCreate(e), nil
}

// ReadCar handles GET /cars/{id} requests.
func (h *OgentHandler) ReadCar(ctx context.Context, params ReadCarParams) (ReadCarRes, error) {
	q := h.client.Car.Query().Where(car.IDEQ(params.ID))
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
	return NewCarRead(e), nil
}

// UpdateCar handles PATCH /cars/{id} requests.
func (h *OgentHandler) UpdateCar(ctx context.Context, req UpdateCarReq, params UpdateCarParams) (UpdateCarRes, error) {
	b := h.client.Car.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Model.Get(); ok {
		b.SetModel(v)
	}
	if v, ok := req.RegisteredAt.Get(); ok {
		b.SetRegisteredAt(v)
	}
	// Add all edges.
	if v, ok := req.Owner.Get(); ok {
		b.SetOwnerID(v)
	}
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
	q := h.client.Car.Query().Where(car.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewCarUpdate(e), nil
}

// DeleteCar handles DELETE /cars/{id} requests.
func (h *OgentHandler) DeleteCar(ctx context.Context, params DeleteCarParams) (DeleteCarRes, error) {
	err := h.client.Car.DeleteOneID(params.ID).Exec(ctx)
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
	return new(DeleteCarNoContent), nil

}

// ListCar handles GET /cars requests.
func (h *OgentHandler) ListCar(ctx context.Context, params ListCarParams) (ListCarRes, error) {
	q := h.client.Car.Query()
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
	r := NewCarLists(es)
	return (*ListCarOKApplicationJSON)(&r), nil
}

// ReadCarOwner handles GET /cars/{id}/owner requests.
func (h *OgentHandler) ReadCarOwner(ctx context.Context, params ReadCarOwnerParams) (ReadCarOwnerRes, error) {
	q := h.client.Car.Query().Where(car.IDEQ(params.ID)).QueryOwner()
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
	return NewCarOwnerRead(e), nil
}

// CreateGroup handles POST /groups requests.
func (h *OgentHandler) CreateGroup(ctx context.Context, req CreateGroupReq) (CreateGroupRes, error) {
	b := h.client.Group.Create()
	// Add all fields.
	b.SetName(req.Name)
	// Add all edges.
	b.AddUserIDs(req.Users...)
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
	q := h.client.Group.Query().Where(group.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewGroupCreate(e), nil
}

// ReadGroup handles GET /groups/{id} requests.
func (h *OgentHandler) ReadGroup(ctx context.Context, params ReadGroupParams) (ReadGroupRes, error) {
	q := h.client.Group.Query().Where(group.IDEQ(params.ID))
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
	return NewGroupRead(e), nil
}

// UpdateGroup handles PATCH /groups/{id} requests.
func (h *OgentHandler) UpdateGroup(ctx context.Context, req UpdateGroupReq, params UpdateGroupParams) (UpdateGroupRes, error) {
	b := h.client.Group.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Name.Get(); ok {
		b.SetName(v)
	}
	// Add all edges.
	b.ClearUsers().AddUserIDs(req.Users...)
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
	q := h.client.Group.Query().Where(group.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewGroupUpdate(e), nil
}

// DeleteGroup handles DELETE /groups/{id} requests.
func (h *OgentHandler) DeleteGroup(ctx context.Context, params DeleteGroupParams) (DeleteGroupRes, error) {
	err := h.client.Group.DeleteOneID(params.ID).Exec(ctx)
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
	return new(DeleteGroupNoContent), nil

}

// ListGroup handles GET /groups requests.
func (h *OgentHandler) ListGroup(ctx context.Context, params ListGroupParams) (ListGroupRes, error) {
	q := h.client.Group.Query()
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
	r := NewGroupLists(es)
	return (*ListGroupOKApplicationJSON)(&r), nil
}

// ListGroupUsers handles GET /groups/{id}/users requests.
func (h *OgentHandler) ListGroupUsers(ctx context.Context, params ListGroupUsersParams) (ListGroupUsersRes, error) {
	q := h.client.Group.Query().Where(group.IDEQ(params.ID)).QueryUsers()
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
	r := NewGroupUsersLists(es)
	return (*ListGroupUsersOKApplicationJSON)(&r), nil
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

// DeleteUser handles DELETE /users/{id} requests.
func (h *OgentHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error) {
	err := h.client.User.DeleteOneID(params.ID).Exec(ctx)
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
	return new(DeleteUserNoContent), nil

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

// ListUserCars handles GET /users/{id}/cars requests.
func (h *OgentHandler) ListUserCars(ctx context.Context, params ListUserCarsParams) (ListUserCarsRes, error) {
	q := h.client.User.Query().Where(user.IDEQ(params.ID)).QueryCars()
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
	r := NewUserCarsLists(es)
	return (*ListUserCarsOKApplicationJSON)(&r), nil
}

// ListUserGroups handles GET /users/{id}/groups requests.
func (h *OgentHandler) ListUserGroups(ctx context.Context, params ListUserGroupsParams) (ListUserGroupsRes, error) {
	q := h.client.User.Query().Where(user.IDEQ(params.ID)).QueryGroups()
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
	r := NewUserGroupsLists(es)
	return (*ListUserGroupsOKApplicationJSON)(&r), nil
}
