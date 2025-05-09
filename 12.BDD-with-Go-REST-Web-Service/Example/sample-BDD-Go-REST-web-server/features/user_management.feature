Feature: User Management

  Scenario: Create a new user
    When I create a user with name "John" and email "john@example.com"
    Then the response code should be 201

  Scenario: Get all users
    When I fetch all users
    Then the response code should be 200

  Scenario: Get user by ID
    When I fetch user with ID 1
    Then the response code should be 200

  Scenario: Update user by ID
    When I update user with ID 1 to name "Johnny" and email "johnny@example.com"
    Then the response code should be 200

  Scenario: Delete user by ID
    When I delete user with ID 1
    Then the response code should be 204