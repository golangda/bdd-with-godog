package main

import (
    "os"
    "testing"

    "github.com/cucumber/godog"
    "github.com/cucumber/godog/colors"
    "github.com/golangda/bdd-with-godog/12.BDD-with-Go-REST-Web-Service/Example/sample-BDD-Go-REST-web-server/stepdefs"

)
func TestFeatures(t *testing.T) {
    opts := godog.Options{
        Format:   "pretty",
        Paths:    []string{"features"},
        Output:   colors.Colored(os.Stdout),
        TestingT: t,
    }

    suite := godog.TestSuite{
        Name:                "user-api",
        ScenarioInitializer: stepdefs.InitializeScenario,
        Options:             &opts,
    }

    if suite.Run() != 0 {
        t.Fatal("godog tests failed")
    }
}