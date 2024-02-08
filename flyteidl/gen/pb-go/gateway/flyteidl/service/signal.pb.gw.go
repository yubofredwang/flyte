// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: flyteidl/service/signal.proto

/*
Package service is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package service

import (
	"context"
	"io"
	"net/http"

	extAdmin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
	extService "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_SignalService_ListSignals_0 = &utilities.DoubleArray{Encoding: map[string]int{"workflow_execution_id": 0, "project": 1, "domain": 2, "name": 3}, Base: []int{1, 6, 7, 8, 9, 2, 0, 4, 0, 6, 0, 0, 0, 0}, Check: []int{0, 1, 1, 1, 1, 2, 6, 2, 8, 2, 10, 3, 4, 5}}
)

func request_SignalService_ListSignals_0(ctx context.Context, marshaler runtime.Marshaler, client extService.SignalServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extAdmin.SignalListRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["workflow_execution_id.project"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.project")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.project", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.project", err)
	}

	val, ok = pathParams["workflow_execution_id.domain"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.domain")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.domain", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.domain", err)
	}

	val, ok = pathParams["workflow_execution_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.name", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.name", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_SignalService_ListSignals_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListSignals(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_SignalService_ListSignals_0(ctx context.Context, marshaler runtime.Marshaler, server extService.SignalServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extAdmin.SignalListRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["workflow_execution_id.project"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.project")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.project", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.project", err)
	}

	val, ok = pathParams["workflow_execution_id.domain"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.domain")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.domain", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.domain", err)
	}

	val, ok = pathParams["workflow_execution_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.name", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.name", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_SignalService_ListSignals_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ListSignals(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_SignalService_ListSignals_1 = &utilities.DoubleArray{Encoding: map[string]int{"workflow_execution_id": 0, "org": 1, "project": 2, "domain": 3, "name": 4}, Base: []int{1, 8, 9, 10, 11, 12, 2, 0, 4, 0, 6, 0, 8, 0, 0, 0, 0, 0}, Check: []int{0, 1, 1, 1, 1, 1, 2, 7, 2, 9, 2, 11, 2, 13, 3, 4, 5, 6}}
)

func request_SignalService_ListSignals_1(ctx context.Context, marshaler runtime.Marshaler, client extService.SignalServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extAdmin.SignalListRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["workflow_execution_id.org"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.org")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.org", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.org", err)
	}

	val, ok = pathParams["workflow_execution_id.project"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.project")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.project", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.project", err)
	}

	val, ok = pathParams["workflow_execution_id.domain"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.domain")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.domain", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.domain", err)
	}

	val, ok = pathParams["workflow_execution_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.name", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.name", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_SignalService_ListSignals_1); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListSignals(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_SignalService_ListSignals_1(ctx context.Context, marshaler runtime.Marshaler, server extService.SignalServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extAdmin.SignalListRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["workflow_execution_id.org"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.org")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.org", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.org", err)
	}

	val, ok = pathParams["workflow_execution_id.project"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.project")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.project", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.project", err)
	}

	val, ok = pathParams["workflow_execution_id.domain"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.domain")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.domain", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.domain", err)
	}

	val, ok = pathParams["workflow_execution_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "workflow_execution_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "workflow_execution_id.name", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "workflow_execution_id.name", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_SignalService_ListSignals_1); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ListSignals(ctx, &protoReq)
	return msg, metadata, err

}

func request_SignalService_SetSignal_0(ctx context.Context, marshaler runtime.Marshaler, client extService.SignalServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extAdmin.SignalSetRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.SetSignal(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_SignalService_SetSignal_0(ctx context.Context, marshaler runtime.Marshaler, server extService.SignalServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extAdmin.SignalSetRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.SetSignal(ctx, &protoReq)
	return msg, metadata, err

}

func request_SignalService_SetSignal_1(ctx context.Context, marshaler runtime.Marshaler, client extService.SignalServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extAdmin.SignalSetRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id.execution_id.org"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id.execution_id.org")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "id.execution_id.org", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id.execution_id.org", err)
	}

	msg, err := client.SetSignal(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_SignalService_SetSignal_1(ctx context.Context, marshaler runtime.Marshaler, server extService.SignalServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extAdmin.SignalSetRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id.execution_id.org"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id.execution_id.org")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "id.execution_id.org", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id.execution_id.org", err)
	}

	msg, err := server.SetSignal(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterSignalServiceHandlerServer registers the http handlers for service SignalService to "mux".
// UnaryRPC     :call SignalServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterSignalServiceHandlerFromEndpoint instead.
func RegisterSignalServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server extService.SignalServiceServer) error {

	mux.Handle("GET", pattern_SignalService_ListSignals_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/flyteidl.service.SignalService/ListSignals", runtime.WithHTTPPathPattern("/api/v1/signals/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_SignalService_ListSignals_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SignalService_ListSignals_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_SignalService_ListSignals_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/flyteidl.service.SignalService/ListSignals", runtime.WithHTTPPathPattern("/api/v1/signals/org/{workflow_execution_id.org}/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_SignalService_ListSignals_1(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SignalService_ListSignals_1(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_SignalService_SetSignal_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/flyteidl.service.SignalService/SetSignal", runtime.WithHTTPPathPattern("/api/v1/signals"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_SignalService_SetSignal_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SignalService_SetSignal_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_SignalService_SetSignal_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/flyteidl.service.SignalService/SetSignal", runtime.WithHTTPPathPattern("/api/v1/signals/org/{id.execution_id.org}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_SignalService_SetSignal_1(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SignalService_SetSignal_1(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterSignalServiceHandlerFromEndpoint is same as RegisterSignalServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterSignalServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterSignalServiceHandler(ctx, mux, conn)
}

// RegisterSignalServiceHandler registers the http handlers for service SignalService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterSignalServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterSignalServiceHandlerClient(ctx, mux, extService.NewSignalServiceClient(conn))
}

// RegisterSignalServiceHandlerClient registers the http handlers for service SignalService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "extService.SignalServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "extService.SignalServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "extService.SignalServiceClient" to call the correct interceptors.
func RegisterSignalServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client extService.SignalServiceClient) error {

	mux.Handle("GET", pattern_SignalService_ListSignals_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/flyteidl.service.SignalService/ListSignals", runtime.WithHTTPPathPattern("/api/v1/signals/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SignalService_ListSignals_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SignalService_ListSignals_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_SignalService_ListSignals_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/flyteidl.service.SignalService/ListSignals", runtime.WithHTTPPathPattern("/api/v1/signals/org/{workflow_execution_id.org}/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SignalService_ListSignals_1(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SignalService_ListSignals_1(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_SignalService_SetSignal_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/flyteidl.service.SignalService/SetSignal", runtime.WithHTTPPathPattern("/api/v1/signals"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SignalService_SetSignal_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SignalService_SetSignal_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_SignalService_SetSignal_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/flyteidl.service.SignalService/SetSignal", runtime.WithHTTPPathPattern("/api/v1/signals/org/{id.execution_id.org}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SignalService_SetSignal_1(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SignalService_SetSignal_1(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_SignalService_ListSignals_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5}, []string{"api", "v1", "signals", "workflow_execution_id.project", "workflow_execution_id.domain", "workflow_execution_id.name"}, ""))

	pattern_SignalService_ListSignals_1 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5, 1, 0, 4, 1, 5, 6, 1, 0, 4, 1, 5, 7}, []string{"api", "v1", "signals", "org", "workflow_execution_id.org", "workflow_execution_id.project", "workflow_execution_id.domain", "workflow_execution_id.name"}, ""))

	pattern_SignalService_SetSignal_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v1", "signals"}, ""))

	pattern_SignalService_SetSignal_1 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "v1", "signals", "org", "id.execution_id.org"}, ""))
)

var (
	forward_SignalService_ListSignals_0 = runtime.ForwardResponseMessage

	forward_SignalService_ListSignals_1 = runtime.ForwardResponseMessage

	forward_SignalService_SetSignal_0 = runtime.ForwardResponseMessage

	forward_SignalService_SetSignal_1 = runtime.ForwardResponseMessage
)