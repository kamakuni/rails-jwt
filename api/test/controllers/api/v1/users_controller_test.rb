require 'test_helper'

class Api::V1::UsersControllerTest < ActionDispatch::IntegrationTest
  test 'signup' do
    #name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { email: email, password: password }
    assert_response :success
    user = response.parsed_body
#    assert_equal name, user['name']
    assert_equal email, user['email']
  end

  test 'signup with name which is maximum length' do
    name = 'a' * 50
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    assert_response :success
  end

  test 'signup without email' do
    name = 'test1'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, password: password }
    assert_response 400
  end

  test 'signup without password' do
    email = 'test1@test.com'
    name = 'test1'
    post api_v1_signup_url, params: { email: email, name: name }
    assert_response 400
  end
end
