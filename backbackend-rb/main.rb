$LOAD_PATH.unshift(File.join(__dir__, 'protobufschema'))
require 'grpc'
require 'greeter_services_pb'

class GreeterServer < Protobufschema::Greeter::Service
  def say_hello(req, _call)
    Protobufschema::HelloReply.new(message: "Hello, #{req.name}")
  end
end

server = GRPC::RpcServer.new

server.add_http2_port('0.0.0.0:50051', :this_port_is_insecure)
server.handle(GreeterServer)
server.run_till_terminated_or_interrupted(%w[INT TERM])
