package transport

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
	"github.com/bishopfox/sliver/protobuf/rpcpb"
	"github.com/bishopfox/sliver/server/log"
	"github.com/bishopfox/sliver/server/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var (
	pipeLog = log.NamedLogger("transport", "pipe")
)

// LocalListener - Bind gRPC server to an in-memory listener, which is
//                 typically used for unit testing, but ... it should be fine
func LocalListener() (*grpc.Server, *bufconn.Listener, error) {
	pipeLog.Infof("Binding gRPC to listener ...")
	ln := bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()
	rpcpb.RegisterSliverRPCServer(grpcServer, rpc.NewServer())
	go func() {
		if err := grpcServer.Serve(ln); err != nil {
			pipeLog.Fatalf("gRPC local listener error: %v", err)
		}
	}()
	return grpcServer, ln, nil
}