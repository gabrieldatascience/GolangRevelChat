class AddUniqueToUser < ActiveRecord::Migration
  def change
    execute "ALTER TABLE user ADD UNIQUE (name)"
    execute "ALTER TABLE user ADD UNIQUE (email)"
    execute "ALTER TABLE room  ADD UNIQUE (room_key)"
  end
end
