Feature: User Login

  Scenario: Successful login with valid credentials
    Given a registered user with username "john" and password "secret"
    When the user logs in with username "john" and password "secret"
    Then the login should be successful