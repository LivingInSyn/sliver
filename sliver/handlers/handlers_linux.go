package handlers

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"github.com/bishopfox/sliver/protobuf/sliverpb"
)

var (
	linuxHandlers = map[uint32]RPCHandler{
		sliverpb.MsgPsReq:         psHandler,
		sliverpb.MsgTerminate:     terminateHandler,
		sliverpb.MsgPing:          pingHandler,
		sliverpb.MsgLsReq:         dirListHandler,
		sliverpb.MsgDownloadReq:   downloadHandler,
		sliverpb.MsgUploadReq:     uploadHandler,
		sliverpb.MsgCdReq:         cdHandler,
		sliverpb.MsgPwdReq:        pwdHandler,
		sliverpb.MsgRmReq:         rmHandler,
		sliverpb.MsgMkdirReq:      mkdirHandler,
		sliverpb.MsgTaskReq:       taskHandler,
		sliverpb.MsgRemoteTaskReq: remoteTaskHandler,
		sliverpb.MsgIfconfigReq:   ifconfigHandler,
		sliverpb.MsgExecuteReq:    executeHandler,

		sliverpb.MsgScreenshotReq: screenshotHandler,

		sliverpb.MsgNetstatReq:  netstatHandler,
		sliverpb.MsgSideloadReq: sideloadHandler,
	}
)

func GetSystemHandlers() map[uint32]RPCHandler {
	return linuxHandlers
}