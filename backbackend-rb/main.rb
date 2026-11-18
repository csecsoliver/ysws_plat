require 'grpc'
require 'protobufschema/greeter_services_pb'
require 'sequel'

class ReviewServer < Protobufschema::ReviewQueue::Service
  def new()
    self.DB = Sequel.sqlite('review.db')
  end
  def enqueue(req, _call)

    Protobufschema::QueueTicket.new(uuid: req.name+req.author+req.description)
  end
end

server = GRPC::RpcServer.new

server.add_http2_port('0.0.0.0:50051', :this_port_is_insecure)
server.handle(ReviewServer)
server.run_till_terminated_or_interrupted(%w[INT TERM])
