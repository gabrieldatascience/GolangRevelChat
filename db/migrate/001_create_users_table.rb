class CreateUsersTable < ActiveRecord::Migration 
  def change
    create_table :user do |t|
      t.string   :name, :null => false
      t.string   :email, :null => false
      t.string   :salt, :null => false
      t.string   :encryptpasswd, :null => false
      t.datetime :created
      t.datetime :updated
    end
  end
end
