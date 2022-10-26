# CardsController
class Api::V1::CardsController < ApplicationController
  before_action :authenticate_user!

  def index
    cards = Card.where(list_id: params['list_id'])
    render json: cards
  rescue ActiveRecord::RecordNotFound => e
    render status: 404, json: { message: e.message }
  end

  def create
    card = Card.create!(title: params['title'], body: params['body'], list_id: params['list_id'])
    render json: card
  rescue ActiveRecord::RecordInvalid => e
    render json: { message: e.message }
  end

  def show
    card = Card.find(params[:id])
    render json: card
  rescue ActiveRecord::RecordNotFound => e
    render status: 404, json: { message: e.message }
  end

  def update
    begin
      card = Card.find(params[:id])
    rescue ActiveRecord::RecordNotFound => e
      render status: 404, json: { message: e.message }
    end
    begin
      card.update!(title: params[:title], body: params[:body], list_id: params[:list_id])
      render json: card
    rescue ActiveRecord::RecordInvalid => e
      render status: 400, json: { message: e.message }
    end
  end

  def destroy
    begin
      card = Card.find(params[:id])
    rescue ActiveRecord::RecordNotFound => e
      render status: 404, json: { message: e.message }
    end
    begin
      card.destroy
      render json: { message: "Card is deleted." }
    rescue ActiveRecord::RecordInvalid => e
      render json: { message: e.message }
    end
  end
end
