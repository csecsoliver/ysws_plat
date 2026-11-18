$LOAD_PATH.unshift(File.join(__dir__, 'protobufschema'))
require 'grpc'
require 'greeter_services_pb'

class ReviewServer < Protobufschema::ReviewQueue::Service
  def Enqueue(req, _call)
    Protobufschema::QueueTicket.new(uuid: "asd")
  end
end

server = GRPC::RpcServer.new

server.add_http2_port('0.0.0.0:50051', :this_port_is_insecure)
server.handle(GreeterServer)
server.run_till_terminated_or_interrupted(%w[INT TERM])
