
Sequel.migration do
	up do
		create_table(:projects) do
			primary_key :id, null: false
			String :uuid, size: 36, unique: true, null: false
			String :name, text: true, null: false
			String :description, text: true, null: false
		end	
		create_table(:submissions) do
			primary_key :id, null: false
			String :uuid, size: 36, unique: true, null: false
			DateTime :date, null: false
			String :message, text: true, null: false
			foreign_key :project_id, :projects, null: false, index: true
		end
		create_table(:submission_updates) do 
			primary_key :id, null: false
			String :uuid, size: 36, unique: true, null: false
			DateTime :date, null: false
			foreign_key :submission_id, :submissions, null: false, index: true
		end
	end
	down do
		drop_table(:submission_updates)
		drop_table(:submissions)
		drop_table(:projects)
	end
end
