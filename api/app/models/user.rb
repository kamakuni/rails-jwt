class User < ApplicationRecord
  validates :email, :password, presence: true
  validates :name, length: { maximum: 50, too_long: "%{count} characters is the maximum allowed" }
  has_secure_password
end
