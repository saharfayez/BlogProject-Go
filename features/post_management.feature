Feature: Post management

  To manage posts, a user should be able to create and delete posts.

  Scenario: Create a post successfully
    Given a user is logged in with username "sahar" and password "0000"
    When the user sends a POST request to "/api/posts" with title "My First Post" and content "This is my post."
    Then the response status should be 201
    And the response body should be "post created successfully"

