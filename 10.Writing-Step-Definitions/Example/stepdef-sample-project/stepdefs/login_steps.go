package stepdefs

import (
    "fmt"
    "github.com/cucumber/godog"
)

type Suite struct {
    enteredUsername string
    credentials     map[string]string
    loginSuccess    bool
}

func (s *Suite) iEnterUsername(username string) error {
    s.enteredUsername = username
    return nil
}

func (s *Suite) iEnterPasswordForUsername(password, username string) error {
    if s.credentials == nil {
        s.credentials = make(map[string]string)
    }
    s.credentials[username] = password
    return nil
}

func (s *Suite) loginShouldBeSuccessful() error {
    pass, ok := s.credentials[s.enteredUsername]
    if !ok || pass == "" {
        return fmt.Errorf("login failed for user %s", s.enteredUsername)
    }
    s.loginSuccess = true
    return nil
}

func (s *Suite) theFollowingUsersExist(table *godog.Table) error {
    if s.credentials == nil {
        s.credentials = make(map[string]string)
    }

    for _, row := range table.Rows[1:] {
        username := row.Cells[0].Value
        password := row.Cells[1].Value
        s.credentials[username] = password
    }
    return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
    s := &Suite{
        credentials: make(map[string]string),
    }

    ctx.Step(`^I enter username "([^"]*)"$`, s.iEnterUsername)
    ctx.Step(`^I enter password "([^"]*)" for username "([^"]*)"$`, s.iEnterPasswordForUsername)
    ctx.Step(`^the login should be successful$`, s.loginShouldBeSuccessful)
    ctx.Step(`^the following users exist:$`, s.theFollowingUsersExist)
}