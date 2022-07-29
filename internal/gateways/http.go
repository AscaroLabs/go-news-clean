package gateways

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	pb "go-news-clean/internal/proto"
	"go-news-clean/pkg/env"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

type httpServer struct {
	mux *runtime.ServeMux
}

func (s *httpServer) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	pb.RegisterNewsServiceHandlerFromEndpoint(
		ctx,
		s.mux,
		fmt.Sprintf("%s:%s", env.GrpcHost, env.GrpcPort),
		opts,
	)
	pb.RegisterContentCheckServiceHandlerFromEndpoint(
		ctx,
		s.mux,
		fmt.Sprintf("%s:%s", env.GrpcHost, env.GrpcPort),
		opts,
	)
	if err := http.ListenAndServe(
		fmt.Sprintf(":%s", env.Port),
		s.mux,
	); err != nil {
		return err
	}
	return nil
}

func NewHTTPServer() *httpServer {
	mux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(httpResponseStatusCodeModifier),
		runtime.WithIncomingHeaderMatcher(CustomMatcher),
	)
	return &httpServer{
		mux: mux,
	}
}

func CustomMatcher(key string) (string, bool) {
	switch key {
	case "Authorization":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func httpResponseStatusCodeModifier(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		log.Fatal("[REST] Can't get metadata from context")
	}
	if vals := md.HeaderMD.Get("x-http-code"); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}
		delete(md.HeaderMD, "x-http-code")
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")

		switch code {
		case 401:
			w.Header().Set("WWW-Authenticate", "Bearer")
			w.WriteHeader(http.StatusUnauthorized)
			return nil
		default:
			w.WriteHeader(code)
		}

	}
	return nil
}
