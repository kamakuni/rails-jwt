Rails.application.routes.draw do
  namespace :api do
    namespace :v1 do
      get '/users' => 'users#index'
      post '/signup' => 'users#create'
      get '/login' => 'sessions#index' 
      post '/login' => 'sessions#create'
      delete '/logout' => 'sessions#destory'
      resources :lists#, only: %i[index create update destory]
      resources :cards
    end
  end
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
end
