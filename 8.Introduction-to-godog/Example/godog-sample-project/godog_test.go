package main

import (
    "os"
    "testing"

    "github.com/cucumber/godog"
    "github.com/cucumber/godog/colors"

    "github.com/golangda/bdd-with-godog/8.Introduction-to-godog/Example/godog-sample-project/stepdefs"
)

func TestFeatures(t *testing.T) {
    opts := godog.Options{
        Format:        "pretty",
        Paths:         []string{"features"},
        Output:        colors.Colored(os.Stdout), // ✅ FIXED
        Strict:        true,
        TestingT:      t,
        StopOnFailure: true,
    }

    suite := godog.TestSuite{
        ScenarioInitializer: stepdefs.InitializeScenario,
        Options:             &opts,
    }

    if suite.Run() != 0 {
        t.Fatal("test suite failed")
    }
}
