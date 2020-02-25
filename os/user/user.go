// This file has automatically been generated on Wed Feb 26 02:10:10 +05 2020.
// DO NOT EDIT.
package user

import (
	"os/user"
	_ "unsafe"
)

//go:linkname LookupGroupId os/user.LookupGroupId
//go:noescape
func LookupGroupId(gid string) (*user.Group, error)

//go:linkname unknowngrouperrorerror user.sub_unknowngrouperrorerror
func unknowngrouperrorerror(e user.UnknownGroupError) string {
	return e.Error()
}

//go:linkname UnknownGroupErrorError user.sub_unknowngrouperrorerror
//go:noescape
func UnknownGroupErrorError(e user.UnknownGroupError) string

//go:linkname unknowngroupiderrorerror user.sub_unknowngroupiderrorerror
func unknowngroupiderrorerror(e user.UnknownGroupIdError) string {
	return e.Error()
}

//go:linkname UnknownGroupIdErrorError user.sub_unknowngroupiderrorerror
//go:noescape
func UnknownGroupIdErrorError(e user.UnknownGroupIdError) string

//go:linkname unknownuseriderrorerror user.sub_unknownuseriderrorerror
func unknownuseriderrorerror(e user.UnknownUserIdError) string {
	return e.Error()
}

//go:linkname UnknownUserIdErrorError user.sub_unknownuseriderrorerror
//go:noescape
func UnknownUserIdErrorError(e user.UnknownUserIdError) string

//go:linkname Current os/user.Current
//go:noescape
func Current() (*user.User, error)

//go:linkname Lookup os/user.Lookup
//go:noescape
func Lookup(username string) (*user.User, error)

//go:linkname LookupId os/user.LookupId
//go:noescape
func LookupId(uid string) (*user.User, error)

//go:linkname LookupGroup os/user.LookupGroup
//go:noescape
func LookupGroup(name string) (*user.Group, error)

//go:linkname unknownusererrorerror user.sub_unknownusererrorerror
func unknownusererrorerror(e user.UnknownUserError) string {
	return e.Error()
}

//go:linkname UnknownUserErrorError user.sub_unknownusererrorerror
//go:noescape
func UnknownUserErrorError(e user.UnknownUserError) string

//go:linkname usergroupids user.sub_usergroupids
func usergroupids(u *user.User) ([]string, error) {
	return u.GroupIds()
}

//go:linkname UserGroupIds user.sub_usergroupids
//go:noescape
func UserGroupIds(u *user.User) ([]string, error)
