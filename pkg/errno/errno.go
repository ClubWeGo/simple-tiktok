package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                = 0
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	RecodeAlreadyExistErrCode  = 10003
	AuthorizationFailedErrCode = 10004
	DBErrCode                  = 10005
	RPCErrCode                 = 10006
	RecordNotExistErrCode      = 10007
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	RecordAlreadyExistErr  = NewErrNo(RecodeAlreadyExistErrCode, "Record already exists")
	RecordNotExistErr      = NewErrNo(RecordNotExistErrCode, "Record not exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	DBErr                  = NewErrNo(DBErrCode, "DB error")
	RPCErr                 = NewErrNo(RPCErrCode, "RPC error")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
