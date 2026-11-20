require 'grpc'
require 'greeter_services_pb'
require 'sequel'

class ReviewServer < Protobufschema::ReviewQueue::Service
	def initialize
		@db = Sequel.sqlite('review.db')
	end
	def enqueue(req, _call)
		@db[:submissions].insert(
			uuid: SecureRandom.uuid,
			date: Time.now,
			message: req.description,
			project_id: 1
		)
		Protobufschema::QueueTicket.new(uuid: req.name+req.author+req.description)
	end
end

server = GRPC::RpcServer.new

server.add_http2_port('0.0.0.0:50051', :this_port_is_insecure)
server.handle(ReviewServer)
server.run_till_terminated_or_interrupted(%w[INT TERM])
