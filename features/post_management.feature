Feature: Post Management
  To manage posts, a user should be able to create posts.

  Background:
    Given an account exists with username "sahar"
    And user is logged in with username "sahar"

  @create_post_successfully_cleaninsert_seed
#  @create_post_successfully2_insert_seed
  Scenario: Create a post successfully
    When the user creates a post with title "My First Post" and content "This is my post"
    Then post should be created successfully with title "My First Post" and content "This is my post"
    And user should be redirected to home page


#  @Create
#  Scenario Outline: temp
#    When the user creates post with title <title> and content <content>
#    Then post should be created successfully with title <title> and content <content>
#    And user should be directed to landing page
#
#    Examples:
#      | title         | content         |
#      | "My First Post" | "This is my post" |