
Sequel.migration do
	up do
		create_table(:projects) do
			primary_key :uuid, type: String, size: 36, auto_increment: false, null: false
			String :name, text: true, null: false
			String :description, text: true, null: false
		end
		create_table(:submissions) do
			primary_key :uuid, type: String, size: 36, auto_increment: false, null: false
			DateTime :date, null: false
			String :message, text: true, null: false
			foreign_key :project_uuid, :projects, key: :uuid, type: String, size: 36, null: false, index: true
		end
		create_table(:submission_updates) do
			primary_key :uuid, type: String, size: 36, auto_increment: false, null: false
			DateTime :date, null: false
			foreign_key :submission_uuid, :submissions, key: :uuid, type: String, size: 36, null: false, index: true
		end
	end
	down do
		drop_table(:submission_updates)
		drop_table(:submissions)
		drop_table(:projects)
	end
end
