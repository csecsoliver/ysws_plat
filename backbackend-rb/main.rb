require 'grpc'
require 'greeter_services_pb'
require 'sequel'
require 'securerandom'
class ReviewServer < Protobufschema::ReviewQueue::Service
	def initialize
		@db = Sequel.sqlite('review.db')
	end
	def enqueue(req, _call)
		uuid = SecureRandom.uuid
		@db[:submissions].insert(
			uuid: uuid,
			date: Time.now,
			message: req.description,
			project_id: 1
		)
		Protobufschema::QueueTicket.new(uuid: uuid)
	end
	def init_project(req, _call)
		uuid = SecureRandom.uuid
		@db[:projects].insert(
			uuid: uuid,
			name: req.name,
			description: req.description
		)
		Protobufschema::ProjectTicket.new(uuid: uuid)
	end

end

server = GRPC::RpcServer.new
server.add_http2_port('0.0.0.0:50051', :this_port_is_insecure)
server.handle(ReviewServer)
print "server starting"
server.run_till_terminated_or_interrupted(%w[INT TERM])
