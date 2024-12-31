Feature: Post Management
  To manage posts, a user should be able to create and delete posts.
  Background:
    Given an account exists with username "sahar"
    And user is logged in with username "sahar"

  Scenario Outline: Create a post successfully
    When the user creates a post with title <title> and content <content>
    Then post should be created successfully with title <title> and content <content>
    And user should be redirected to home page

    Examples:
      | title         | content         |
      | "My First Post" | "This is my post" |

#  @Create
#  Scenario Outline: temp
#    When the user creates post with title <title> and content <content>
#    Then post should be created successfully with title <title> and content <content>
#    And user should be directed to landing page
#
#    Examples:
#      | title         | content         |
#      | "My First Post" | "This is my post" |