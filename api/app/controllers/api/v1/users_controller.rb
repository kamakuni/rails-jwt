# UsersController
class Api::V1::UsersController < ApplicationController
  def create
    user = User.create!(email: params[:email], password: params[:password])
    render json: { name: user.name, email: user.email }
  rescue ActiveRecord::RecordInvalid => e
    render status: 400, json: { message: e.message }
  end
end
