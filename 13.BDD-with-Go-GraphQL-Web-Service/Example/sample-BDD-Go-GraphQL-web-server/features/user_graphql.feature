Feature: GraphQL User Management

  Scenario: Create a new user
    When I send a GraphQL mutation to create a user with name "Alice" and email "alice@example.com"
    Then the response should contain user with name "Alice"

  Scenario: Fetch all users
    When I send a GraphQL query to fetch all users
    Then the response should include at least one user

  Scenario: Update existing user
    When I send a GraphQL mutation to update user with ID 1 to name "Alicia" and email "alicia@example.com"
    Then the response should contain user with name "Alicia"

  Scenario: Delete user by ID
    When I send a GraphQL mutation to delete user with ID 1
    Then the response status should be success