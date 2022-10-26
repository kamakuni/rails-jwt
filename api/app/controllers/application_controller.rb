# ApplicationController
class ApplicationController < ActionController::API
  include ActionController::Cookies
  before_action :has_xhr_header!, except: [:index, :show]

  def authenticate_user!
    render status: 401, json: { message: "Users is unauthorized." } unless session[:user_id]
  end

  private
  def has_xhr_header!
    render status: 403, json: { message: "Forbidden."} unless request.xhr?
  end
end
