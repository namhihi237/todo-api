package errors

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	INVALID_PARAM: "invalid parameter",
	NOT_FOUND:     "not found",
	UNAUTHORIZED:  "unauthorized",
	FORBIDDEN:     "forbidden",
	SERVER_ERROR:  "server error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[SERVER_ERROR]
}
