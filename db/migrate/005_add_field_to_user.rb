class AddFieldToUser < ActiveRecord::Migration
  def change
    add_column :user, :weibo, :string
    add_column :user, :site, :string
    add_column :user, :Introduction, :text
    add_column :user, :avatar, :string
    add_column :user, :signature, :string
    add_column :user, :github, :string
  end
end
