package track

import (
	"context"
	"fmt"
	pb "github.com/Namdar1Ibrakhim/student-track-system/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/keepalive"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type GRPCClient struct {
	Conn   *grpc.ClientConn
	Client pb.PredictionServiceClient
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewGRPCClient(address string) (*GRPCClient, error) {
	keepaliveParams := keepalive.ClientParameters{
		Time:    10 * time.Minute, // отправка пинга каждые 10 минут
		Timeout: 20 * time.Second, // максимальное время ожидания пинга
	}

	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
		grpc.WithKeepaliveParams(keepaliveParams),
		grpc.WithDefaultCallOptions(grpc.UseCompressor("gzip")))

	if err != nil {
		return nil, fmt.Errorf("error connecting to ML service: %w", err)
	}

	client := pb.NewPredictionServiceClient(conn)
	return &GRPCClient{Conn: conn, Client: client}, nil
}

func (g *GRPCClient) Close() {
	if err := g.Conn.Close(); err != nil {
		logrus.Warnf("Error closing gRPC connection: %v", err)
	}
}
