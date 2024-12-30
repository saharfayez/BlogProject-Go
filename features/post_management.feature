Feature: Post management

  To manage posts, a user should be able to create and delete posts.
  Background:
    Given an account exists with username "sahar"

  @create
  Scenario Outline: Create a post successfully
    When the user creates post with title <title> and content <content>
    Then post should be created successfully with title <title> and content <content>
    And user should be directed to home page

    Examples:
      | title         | content         |
      | "My First Post" | "This is my post" |


#  @delete
#  Scenario: Delete a post successfully
#    Given a user is logged in with username "sahar" and password "0000"
#    And the user has created a post with title "My First Post" and content "This is my post."
#    When the user sends a DELETE request to the post's endpoint
#    Then the response status should be 200
#    And the response body should be "post deleted successfully"


