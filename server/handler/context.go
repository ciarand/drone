package handler

import (
	"github.com/drone/drone/shared/model"
	"github.com/zenazn/goji/web"
)

// ToUser returns the User from the current
// request context. If the User does not exist
// a nil value is returned.
func ToUser(c web.C) *model.User {
	var v = c.Env["user"]
	if v == nil {
		return nil
	}
	u, ok := v.(*model.User)
	if !ok {
		return nil
	}
	return u
}

// ToRepo returns the Repo from the current
// request context. If the Repo does not exist
// a nil value is returned.
func ToRepo(c web.C) *model.Repo {
	var v = c.Env["repo"]
	if v == nil {
		return nil
	}
	r, ok := v.(*model.Repo)
	if !ok {
		return nil
	}
	return r
}

// ToRole returns the Role from the current
// request context. If the Role does not exist
// a nil value is returned.
func ToRole(c web.C) *model.Perm {
	var v = c.Env["role"]
	if v == nil {
		return nil
	}
	p, ok := v.(*model.Perm)
	if !ok {
		return nil
	}
	return p
}
