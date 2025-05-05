Feature: User Login

  Scenario: Entering credentials
    Given I enter username "john"
    And I enter password "secret" for username "john"
    Then the login should be successful

  Scenario: Users loaded from table
    Given the following users exist:
      | username | password |
      | alice    | 12345    |
      | bob      | pass123  |