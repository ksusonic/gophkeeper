package server

import (
	"context"
	"net"
	"runtime/debug"
	"strings"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	grpclogging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/ksusonic/gophkeeper/internal/logging"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcController interface {
	ServiceName() string
	RegisterService(srv *grpc.Server)
}

type GrpcServer struct {
	srv    *grpc.Server
	logger logging.Logger
	cfg    *config.ServerConfig
}

type AuthInterceptor interface {
	selector.Matcher
	AuthFunc(ctx context.Context) (context.Context, error)
}

func NewGrpcServer(
	cfg *config.ServerConfig,
	logger logging.Logger,
	authInterceptor AuthInterceptor,
) *GrpcServer {
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)
	return &GrpcServer{
		srv: grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				otelgrpc.UnaryServerInterceptor(),
				grpclogging.UnaryServerInterceptor(&InterceptorLogger{logger}, grpclogging.WithFieldsFromContext(traceIDFunction)),
				selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authInterceptor.AuthFunc), authInterceptor),
				recovery.UnaryServerInterceptor(
					recovery.WithRecoveryHandler(panicRecoveryHandler(logger)),
				),
			),
			grpc.ChainStreamInterceptor(
				otelgrpc.StreamServerInterceptor(),
				grpclogging.StreamServerInterceptor(&InterceptorLogger{logger}, grpclogging.WithFieldsFromContext(traceIDFunction)),
				selector.StreamServerInterceptor(auth.StreamServerInterceptor(authInterceptor.AuthFunc), authInterceptor),
				recovery.StreamServerInterceptor(
					recovery.WithRecoveryHandler(panicRecoveryHandler(logger)),
				),
			)),
		cfg:    cfg,
		logger: logger,
	}
}

func (s *GrpcServer) RegisterController(controller GrpcController) {
	s.logger.Info("Registered %s", controller.ServiceName())
	controller.RegisterService(s.srv)
}

func (s *GrpcServer) ListenAndServe(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		s.logger.Fatal("Cannot listen: %v", err)
	}

	s.logger.Info("GRPC server Listening at %v", lis.Addr())
	if err := s.srv.Serve(lis); err != nil {
		s.logger.Fatal("Failed to serve: %v", err)
	}
}

func (s *GrpcServer) GracefulStop() {
	s.srv.GracefulStop()
}

func panicRecoveryHandler(logger logging.Logger) func(p any) error {
	return func(p any) error {
		errorID := uuid.New()
		logger.Error("errorID: %s got panic: %v, stack: %s", errorID, p, debug.Stack())
		return status.Errorf(codes.Internal, "internal error id: %s", errorID)
	}
}

type InterceptorLogger struct {
	logger logging.Logger
}

func (i *InterceptorLogger) Log(ctx context.Context, level grpclogging.Level, msg string, fields ...any) {
	messageBuilder := func(msg string, fields ...any) (format string) {
		s := strings.Builder{}
		s.Grow(len(msg) + len(fields)*10) // grow on average message len

		s.WriteString(msg + " {")
		fieldIsKey := true
		for i := range fields {
			if fieldIsKey {
				s.WriteString(`"%s"=`)
			} else {
				s.WriteString(`"%v"`)
			}
			if i != len(fields)-1 {
				s.WriteRune(' ')
			}
			fieldIsKey = !fieldIsKey
		}
		s.WriteRune('}')
		return s.String()
	}

	switch level {
	case grpclogging.LevelDebug:
		i.logger.Debug(messageBuilder(msg, fields...), fields...)
	case grpclogging.LevelInfo:
		i.logger.Info(messageBuilder(msg, fields...), fields...)
	case grpclogging.LevelWarn:
		i.logger.Warn(messageBuilder(msg, fields...), fields...)
	case grpclogging.LevelError:
		i.logger.Error(messageBuilder(msg, fields...), fields...)
	default:
		i.logger.Info(messageBuilder(msg, fields...), fields...)
	}
}

func traceIDFunction(ctx context.Context) grpclogging.Fields {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return grpclogging.Fields{"traceID", span.TraceID().String()}
	}
	return nil
}
