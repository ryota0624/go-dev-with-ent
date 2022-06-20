// Code generated by entc, DO NOT EDIT.

package ogent

import "github.com/ryota0624/go-dev-with-ent/ent"

func NewCarCreate(e *ent.Car) *CarCreate {
	if e == nil {
		return nil
	}
	var ret CarCreate
	ret.ID = e.ID
	ret.Model = e.Model
	ret.RegisteredAt = e.RegisteredAt
	return &ret
}

func NewCarCreates(es []*ent.Car) []CarCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]CarCreate, len(es))
	for i, e := range es {
		r[i] = NewCarCreate(e).Elem()
	}
	return r
}

func (c *CarCreate) Elem() CarCreate {
	if c == nil {
		return CarCreate{}
	}
	return *c
}

func NewCarList(e *ent.Car) *CarList {
	if e == nil {
		return nil
	}
	var ret CarList
	ret.ID = e.ID
	ret.Model = e.Model
	ret.RegisteredAt = e.RegisteredAt
	return &ret
}

func NewCarLists(es []*ent.Car) []CarList {
	if len(es) == 0 {
		return nil
	}
	r := make([]CarList, len(es))
	for i, e := range es {
		r[i] = NewCarList(e).Elem()
	}
	return r
}

func (c *CarList) Elem() CarList {
	if c == nil {
		return CarList{}
	}
	return *c
}

func NewCarRead(e *ent.Car) *CarRead {
	if e == nil {
		return nil
	}
	var ret CarRead
	ret.ID = e.ID
	ret.Model = e.Model
	ret.RegisteredAt = e.RegisteredAt
	return &ret
}

func NewCarReads(es []*ent.Car) []CarRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]CarRead, len(es))
	for i, e := range es {
		r[i] = NewCarRead(e).Elem()
	}
	return r
}

func (c *CarRead) Elem() CarRead {
	if c == nil {
		return CarRead{}
	}
	return *c
}

func NewCarUpdate(e *ent.Car) *CarUpdate {
	if e == nil {
		return nil
	}
	var ret CarUpdate
	ret.ID = e.ID
	ret.Model = e.Model
	ret.RegisteredAt = e.RegisteredAt
	return &ret
}

func NewCarUpdates(es []*ent.Car) []CarUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]CarUpdate, len(es))
	for i, e := range es {
		r[i] = NewCarUpdate(e).Elem()
	}
	return r
}

func (c *CarUpdate) Elem() CarUpdate {
	if c == nil {
		return CarUpdate{}
	}
	return *c
}

func NewCarOwnerRead(e *ent.User) *CarOwnerRead {
	if e == nil {
		return nil
	}
	var ret CarOwnerRead
	ret.ID = e.ID
	ret.Age = e.Age
	ret.Name = e.Name
	return &ret
}

func NewCarOwnerReads(es []*ent.User) []CarOwnerRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]CarOwnerRead, len(es))
	for i, e := range es {
		r[i] = NewCarOwnerRead(e).Elem()
	}
	return r
}

func (u *CarOwnerRead) Elem() CarOwnerRead {
	if u == nil {
		return CarOwnerRead{}
	}
	return *u
}

func NewGroupCreate(e *ent.Group) *GroupCreate {
	if e == nil {
		return nil
	}
	var ret GroupCreate
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewGroupCreates(es []*ent.Group) []GroupCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]GroupCreate, len(es))
	for i, e := range es {
		r[i] = NewGroupCreate(e).Elem()
	}
	return r
}

func (gr *GroupCreate) Elem() GroupCreate {
	if gr == nil {
		return GroupCreate{}
	}
	return *gr
}

func NewGroupList(e *ent.Group) *GroupList {
	if e == nil {
		return nil
	}
	var ret GroupList
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewGroupLists(es []*ent.Group) []GroupList {
	if len(es) == 0 {
		return nil
	}
	r := make([]GroupList, len(es))
	for i, e := range es {
		r[i] = NewGroupList(e).Elem()
	}
	return r
}

func (gr *GroupList) Elem() GroupList {
	if gr == nil {
		return GroupList{}
	}
	return *gr
}

func NewGroupRead(e *ent.Group) *GroupRead {
	if e == nil {
		return nil
	}
	var ret GroupRead
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewGroupReads(es []*ent.Group) []GroupRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]GroupRead, len(es))
	for i, e := range es {
		r[i] = NewGroupRead(e).Elem()
	}
	return r
}

func (gr *GroupRead) Elem() GroupRead {
	if gr == nil {
		return GroupRead{}
	}
	return *gr
}

func NewGroupUpdate(e *ent.Group) *GroupUpdate {
	if e == nil {
		return nil
	}
	var ret GroupUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewGroupUpdates(es []*ent.Group) []GroupUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]GroupUpdate, len(es))
	for i, e := range es {
		r[i] = NewGroupUpdate(e).Elem()
	}
	return r
}

func (gr *GroupUpdate) Elem() GroupUpdate {
	if gr == nil {
		return GroupUpdate{}
	}
	return *gr
}

func NewGroupUsersList(e *ent.User) *GroupUsersList {
	if e == nil {
		return nil
	}
	var ret GroupUsersList
	ret.ID = e.ID
	ret.Age = e.Age
	ret.Name = e.Name
	return &ret
}

func NewGroupUsersLists(es []*ent.User) []GroupUsersList {
	if len(es) == 0 {
		return nil
	}
	r := make([]GroupUsersList, len(es))
	for i, e := range es {
		r[i] = NewGroupUsersList(e).Elem()
	}
	return r
}

func (u *GroupUsersList) Elem() GroupUsersList {
	if u == nil {
		return GroupUsersList{}
	}
	return *u
}

func NewUserCreate(e *ent.User) *UserCreate {
	if e == nil {
		return nil
	}
	var ret UserCreate
	ret.ID = e.ID
	ret.Age = e.Age
	ret.Name = e.Name
	return &ret
}

func NewUserCreates(es []*ent.User) []UserCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserCreate, len(es))
	for i, e := range es {
		r[i] = NewUserCreate(e).Elem()
	}
	return r
}

func (u *UserCreate) Elem() UserCreate {
	if u == nil {
		return UserCreate{}
	}
	return *u
}

func NewUserList(e *ent.User) *UserList {
	if e == nil {
		return nil
	}
	var ret UserList
	ret.ID = e.ID
	ret.Age = e.Age
	ret.Name = e.Name
	return &ret
}

func NewUserLists(es []*ent.User) []UserList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserList, len(es))
	for i, e := range es {
		r[i] = NewUserList(e).Elem()
	}
	return r
}

func (u *UserList) Elem() UserList {
	if u == nil {
		return UserList{}
	}
	return *u
}

func NewUserRead(e *ent.User) *UserRead {
	if e == nil {
		return nil
	}
	var ret UserRead
	ret.ID = e.ID
	ret.Age = e.Age
	ret.Name = e.Name
	return &ret
}

func NewUserReads(es []*ent.User) []UserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserRead, len(es))
	for i, e := range es {
		r[i] = NewUserRead(e).Elem()
	}
	return r
}

func (u *UserRead) Elem() UserRead {
	if u == nil {
		return UserRead{}
	}
	return *u
}

func NewUserUpdate(e *ent.User) *UserUpdate {
	if e == nil {
		return nil
	}
	var ret UserUpdate
	ret.ID = e.ID
	ret.Age = e.Age
	ret.Name = e.Name
	return &ret
}

func NewUserUpdates(es []*ent.User) []UserUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserUpdate, len(es))
	for i, e := range es {
		r[i] = NewUserUpdate(e).Elem()
	}
	return r
}

func (u *UserUpdate) Elem() UserUpdate {
	if u == nil {
		return UserUpdate{}
	}
	return *u
}

func NewUserCarsList(e *ent.Car) *UserCarsList {
	if e == nil {
		return nil
	}
	var ret UserCarsList
	ret.ID = e.ID
	ret.Model = e.Model
	ret.RegisteredAt = e.RegisteredAt
	return &ret
}

func NewUserCarsLists(es []*ent.Car) []UserCarsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserCarsList, len(es))
	for i, e := range es {
		r[i] = NewUserCarsList(e).Elem()
	}
	return r
}

func (c *UserCarsList) Elem() UserCarsList {
	if c == nil {
		return UserCarsList{}
	}
	return *c
}

func NewUserGroupsList(e *ent.Group) *UserGroupsList {
	if e == nil {
		return nil
	}
	var ret UserGroupsList
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewUserGroupsLists(es []*ent.Group) []UserGroupsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserGroupsList, len(es))
	for i, e := range es {
		r[i] = NewUserGroupsList(e).Elem()
	}
	return r
}

func (gr *UserGroupsList) Elem() UserGroupsList {
	if gr == nil {
		return UserGroupsList{}
	}
	return *gr
}
