class CreateRoomTable < ActiveRecord::Migration 
  def change
    create_table :room do |t|
      t.integer   :user_id, :null => false
      t.string    :room_key, :null => false
      t.string    :title, :null => false
      t.string    :description

      t.timestamps
    end
  end
end
