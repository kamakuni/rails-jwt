class List < ApplicationRecord
    has_many :cards, dependent: :destroy
end
