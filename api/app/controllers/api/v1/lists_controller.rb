# ListsController
class Api::V1::ListsController < ApplicationController
  before_action :authenticate_user!

  def index
    render json: List.where(user_id: session[:user_id])
  end

  def create
    list = List.create!(title: params[:title], user_id: session[:user_id])
    render json: list
  rescue ActiveRecord::RecordInvalid => e
    render status: 400, json: { message: e.message }
  end

  def show
    list = List.find(params[:id])
    render json: list
  rescue ActiveRecord::RecordNotFound => e
    render status: 404, json: { message: e.message }
  end

  def update
    begin
      list = List.find(params[:id])
    rescue ActiveRecord::RecordNotFound => e
      render status: 404, json: { message: e.message }
    end
    begin
      list.update!(title: params[:title])
      render json: list
    rescue ActiveRecord::RecordInvalid => e
      render status: 400, json: { message: e.message }
    end
  end

  def destroy
    begin
      list = List.find(params[:id])
    rescue ActiveRecord::RecordNotFound => e
      render status: 404, json: { message: e.message }
    end
    begin
      list.destroy
      render json: { message: "List is deleted." }
    rescue ActiveRecord::RecordInvalid => e
      render json: { message: e.message }
    end
  end
end
