require 'test_helper'

class Api::V1::SessionsControllerTest < ActionDispatch::IntegrationTest
  test 'login' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    assert_response :success
    assert_not_equal "", response.cookies['_session_id']
  end

  test 'logout' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    delete api_v1_logout_url
    assert_response :success
  end
end
