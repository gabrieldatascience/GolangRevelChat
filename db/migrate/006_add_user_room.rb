class AddUserRoom < ActiveRecord::Migration
  def change 
    create_table :user_room, :id => false do |t|
      t.integer :user_id, :null => false
      t.integer :room_id, :null => false

      t.timestamps
    end
  end
end
