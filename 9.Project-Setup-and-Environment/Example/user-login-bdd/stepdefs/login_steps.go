package stepdefs

import (
    "fmt"
    "github.com/cucumber/godog"
)

type user struct {
    username string
    password string
}

type loginSystem struct {
    registeredUsers map[string]string
    loginSuccess    bool
}

var system loginSystem

func aRegisteredUser(username, password string) error {
    if system.registeredUsers == nil {
        system.registeredUsers = make(map[string]string)
    }
    system.registeredUsers[username] = password
    return nil
}

func userLogsIn(username, password string) error {
    if pass, exists := system.registeredUsers[username]; exists && pass == password {
        system.loginSuccess = true
        return nil
    }
    system.loginSuccess = false
    return fmt.Errorf("invalid login")
}

func loginShouldBeSuccessful() error {
    if !system.loginSuccess {
        return fmt.Errorf("expected login to be successful but it failed")
    }
    return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
    system = loginSystem{} // Reset before each scenario

    ctx.Step(`^a registered user with username "([^"]*)" and password "([^"]*)"$`, aRegisteredUser)
    ctx.Step(`^the user logs in with username "([^"]*)" and password "([^"]*)"$`, userLogsIn)
    ctx.Step(`^the login should be successful$`, loginShouldBeSuccessful)
}