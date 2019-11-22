class AddPrivateToRoom < ActiveRecord::Migration 
  def change
    add_column :room, :private, :boolean
  end
end
