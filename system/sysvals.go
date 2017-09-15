package system

import ()

const (
	SESSIONSTATUS = "SESSIONSTATUS"
	SESSIONINFO   = "username"

	USERINFO = "user"

	RESPONSE = "RESPONSE"

	PERMITIONINFO = "perms"
)

type SessionStatus int

const (
	Session_New SessionStatus = iota
	Session_Exist
	Session_Invalid
	Session_Delete
)