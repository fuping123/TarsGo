//
// This file war generated by FastTars2go 1.0
// Generated from AdminF.tars
// Tencent.

package adminf

import (
	"fmt"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
)

type AdminF struct {
	s m.Servant
}

func (_obj *AdminF) Shutdown(_opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(0, "shutdown", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func (_obj *AdminF) Notify(Command string, _opt ...map[string]string) (ret string, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Command, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(0, "notify", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(_resp.SBuffer)
	err = _is.Read_string(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	_ = length
	_ = have
	_ = ty
	return ret, nil
}

func (_obj *AdminF) SetServant(s m.Servant) {
	_obj.s = s
}

func (_obj *AdminF) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

type _impAdminF interface {
	Shutdown() (err error)
	Notify(Command string) (ret string, err error)
}

func (_obj *AdminF) Dispatch(_val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(req.SBuffer)
	_os := codec.NewBuffer()
	_imp := _val.(_impAdminF)
	switch req.SFuncName {
	case "shutdown":
		err := _imp.Shutdown()
		if err != nil {
			return err
		}
	case "notify":
		var Command string
		err = _is.Read_string(&Command, 1, true)
		if err != nil {
			return err
		}
		ret, err := _imp.Notify(Command)
		if err != nil {
			return err
		}

		err = _os.Write_string(ret, 0)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var status map[string]string
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      _os.ToBytes(),
		Status:       status,
		SResultDesc:  "",
		Context:      req.Context,
	}
	_ = length
	_ = have
	_ = ty
	return nil
}
