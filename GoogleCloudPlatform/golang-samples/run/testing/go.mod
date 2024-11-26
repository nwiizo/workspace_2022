module github.com/GoogleCloudPlatform/golang-samples/run/testing

go 1.15

require (
	github.com/GoogleCloudPlatform/golang-samples 113bab98e20a
	github.com/GoogleCloudPlatform/golang-samples/run/grpc-ping v0.0.0-20220609084039-bc188447282d
	github.com/GoogleCloudPlatform/golang-samples/run/grpc-server-streaming v0.0.0-20220609084039-bc188447282d
	golang.org/x/net v0.0.0-20220325170049-de3da57026de
	google.golang.org/grpc v1.45.0
)

replace github.com/GoogleCloudPlatform/golang-samples => ../../
