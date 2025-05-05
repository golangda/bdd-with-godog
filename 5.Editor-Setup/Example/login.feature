Feature: User Login

  Scenario: Successful login
    Given the user is on the login page
    When the user enters valid credentials
    Then they should be redirected to the dashboard