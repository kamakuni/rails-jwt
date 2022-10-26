require "test_helper"

class Api::V1::ListsControllerTest < ActionDispatch::IntegrationTest
  test 'create new list' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    post api_v1_lists_url, params: { title: "test1" }
    assert_response :success
    res = @response.parsed_body
    assert_equal "test1", res['title']
  end

  test 'get all of lists' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    post api_v1_lists_url, params: { title: "test1" }
    post api_v1_lists_url, params: { title: "test2" }
    get api_v1_lists_url
    assert_response :success
    lists = @response.parsed_body
    assert_equal "test1", lists[0]['title']
    assert_equal "test2", lists[1]['title']
  end

  test 'get a list' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    post api_v1_lists_url, params: { title: "test1" }
    registered = @response.parsed_body
    id = registered['id']
    get "#{api_v1_lists_url}/#{id}"
    assert_response :success
    showed = @response.parsed_body
    assert_equal registered['title'], showed['title']
  end

  test 'update a list' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    post api_v1_lists_url, params: { title: "test1" }
    before = @response.parsed_body
    id = before['id']
    patch "#{api_v1_lists_url}/#{id}", params: { title: "test1updated" }
    assert_response :success
    after = @response.parsed_body
    assert_equal 'test1updated', after['title']
  end

  test 'delete a list' do
    name = 'test1'
    email = 'test1@test.com'
    password = 'test1'
    post api_v1_signup_url, params: { name: name, email: email, password: password }
    post api_v1_login_url, params: { email: email, password: password }
    post api_v1_lists_url, params: { title: "test1" }
    before = @response.parsed_body
    id = before['id']
    delete "#{api_v1_lists_url}/#{id}"
    assert_response :success
    get "#{api_v1_lists_url}/#{id}"
    assert_response 404
  end
end
