1. Install protobuf using homebrew
brew install protobuf@3
2. If installed proto2 by mistake: 
brew tap beeftornado/rmtree
brew rmtree protobuf
3. Run this command to know if protobuf, and protoc-gen-go is in your path
echo export PATH="/opt/homebrew/opt/protobuf@3/bin:$PATH" >> ~/.zshrc
and add bin:$HOME/go/bin in the path too
which protoc-gen-go
4. Compile via the following command
protoc -I=/Users/john/grpc --go-grpc_out=/Users/john/grpc/chat /Users/john/grpc/chat.proto