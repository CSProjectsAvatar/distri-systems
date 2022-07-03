# python
python -m grpc_tools.protoc -I ./proto --python_out=. --grpc_python_out=. .\proto\middleware.proto

# go
protoc middleware.proto --go_out=. --go-grpc_out=.