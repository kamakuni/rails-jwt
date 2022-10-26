class CreateCards < ActiveRecord::Migration[6.1]
  def change
    create_table :cards do |t|
      t.string :title, null: false, default: ''
      t.string :body, null: false, default: ''
      t.references :list, foreign_key: true

      t.timestamps
    end
  end
end
