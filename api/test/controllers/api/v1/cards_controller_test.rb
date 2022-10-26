require 'test_helper'

class Api::V1::CardsControllerTest < ActionDispatch::IntegrationTest
  test 'get all cards' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    post api_v1_lists_url, params: { title: 'test1' }
    list = @response.parsed_body
    post api_v1_cards_url, params: { title: 'test1', body: 'body1', list_id: list['id'] }
    get "#{api_v1_cards_url}?list_id=#{list['id']}"
    assert_response :success
    cards = @response.parsed_body
    assert_equal 'test1', cards[0]['title']
    assert_equal 'body1', cards[0]['body']
    assert_equal list['id'], cards[0]['list_id']
  end

  test 'create new card' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    post api_v1_lists_url, params: { title: 'test1' }
    list = @response.parsed_body
    post api_v1_cards_url, params: { title: 'test1', body: 'body1', list_id: list['id'] }
    assert_response :success
    card = @response.parsed_body
    assert_equal 'test1', card['title']
    assert_equal 'body1', card['body']
    assert_equal list['id'], card['list_id']
  end

  test 'get a card' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    post api_v1_lists_url, params: { title: 'test1' }
    list = @response.parsed_body
    post api_v1_cards_url, params: { title: 'test1', body: 'body1', list_id: list['id'] }
    registered = @response.parsed_body
    get "#{api_v1_cards_url}/#{registered['id']}"
    assert_response :success
    showed = @response.parsed_body
    assert_equal 'test1', showed['title']
    assert_equal 'body1', showed['body']
    assert_equal list['id'], showed['list_id']
  end

  test 'update a card' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    post api_v1_lists_url, params: { title: 'test1' }
    list = @response.parsed_body
    post api_v1_cards_url, params: { title: 'test1', body: 'body1', list_id: list['id'] }
    before = @response.parsed_body
    patch "#{api_v1_cards_url}/#{before['id']}",
          params: { title: 'test1updated', body: 'body1updated', list_id: list['id'] }
    assert_response :success
    after = @response.parsed_body
    assert_equal 'test1updated', after['title']
    assert_equal 'body1updated', after['body']
    assert_equal list['id'], after['list_id']
  end

  test 'delete a card' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    post api_v1_lists_url, params: { title: 'test1' }
    list = @response.parsed_body
    post api_v1_cards_url, params: { title: 'test1', body: 'body1', list_id: list['id'] }
    before = @response.parsed_body
    id = before['id']
    delete "#{api_v1_cards_url}/#{id}"
    assert_response :success
    get "#{api_v1_cards_url}/#{id}"
    assert_response 404
  end
end
