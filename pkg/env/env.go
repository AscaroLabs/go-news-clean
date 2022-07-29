package env

import (
	"fmt"
	"os"
)

// All variables for project
var (
	GrpcHost = Getter("GRPC_HOST", "localhost")
	GrpcPort = fmt.Sprintf(":%s", Getter("GRPC_PORT", "50051"))
	HttpHost = Getter("HTTP_HOST", "localhost")
	Port     = fmt.Sprintf(":%s", Getter("SERVER_PORT", "8080"))
	// ----------------------------------- DB -----------------------------------
	Host     = Getter("DB_HOST", "")
	Password = Getter("DB_PASSWORD", "")
	User     = Getter("DB_USER", "")
	Dbname   = Getter("DB_NAME", "")
	DbPort   = Getter("DB_PORT", "")
	SslMode  = Getter("DB_SSL_MODE", "")

	// ----------------------------------- S3 -----------------------------------
	S3Endpoint         = Getter("S3_ENDPOINT", "")
	S3AccessKey        = Getter("S3_ACCESS_KEY", "")
	S3Secret           = Getter("S3_SECRET", "")
	S3SecureConnection = Getter("S3_SECURE_MODE", "false")
	S3Region           = Getter("S3_REGION", "")
	S3TraceON          = Getter("S3_LOGMODE", "")
)

// Getter -
func Getter(key, defaultValue string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}
	return defaultValue
}
