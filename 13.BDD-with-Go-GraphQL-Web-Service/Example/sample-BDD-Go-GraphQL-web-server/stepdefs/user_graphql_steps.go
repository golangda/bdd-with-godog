package stepdefs

import (
    "bytes"
    "fmt"
    "io"
    "net/http"

    "github.com/cucumber/godog"
)

type userSuite struct {
    responseBody []byte
    statusCode   int
}

func (s *userSuite) iSendGraphqlMutationToCreateUser(name, email string) error {
    query := fmt.Sprintf(`{"query": "mutation { createUser(name: \"%s\", email: \"%s\") { name } }"}`, name, email)
    resp, err := http.Post("http://localhost:8080/graphql", "application/json", bytes.NewBuffer([]byte(query)))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    s.statusCode = resp.StatusCode
    s.responseBody, _ = io.ReadAll(resp.Body)
    return nil
}

func (s *userSuite) iSendGraphqlQueryToFetchAllUsers() error {
    query := `{"query": "{ users { name email } }"}`
    resp, err := http.Post("http://localhost:8080/graphql", "application/json", bytes.NewBuffer([]byte(query)))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    s.statusCode = resp.StatusCode
    s.responseBody, _ = io.ReadAll(resp.Body)
    return nil
}

func (s *userSuite) theResponseShouldContainUserWithName(name string) error {
    return assertBodyContains(s.responseBody, name)
}

func (s *userSuite) theResponseShouldIncludeAtLeastOneUser() error {
    return assertBodyContains(s.responseBody, "name")
}

func (s *userSuite) theResponseStatusShouldBeSuccess() error {
    if s.statusCode < 200 || s.statusCode > 299 {
        return fmt.Errorf("expected 2xx status code, got %d", s.statusCode)
    }
    return nil
}

func assertBodyContains(body []byte, expected string) error {
    if !bytes.Contains(body, []byte(expected)) {
        return fmt.Errorf("expected response to contain %q, got: %s", expected, string(body))
    }
    return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
    s := &userSuite{}
    ctx.Step(`^I send a GraphQL mutation to create a user with name "([^"]*)" and email "([^"]*)"$`, s.iSendGraphqlMutationToCreateUser)
    ctx.Step(`^I send a GraphQL query to fetch all users$`, s.iSendGraphqlQueryToFetchAllUsers)
    ctx.Step(`^the response should contain user with name "([^"]*)"$`, s.theResponseShouldContainUserWithName)
    ctx.Step(`^the response should include at least one user$`, s.theResponseShouldIncludeAtLeastOneUser)
    ctx.Step(`^the response status should be success$`, s.theResponseStatusShouldBeSuccess)
}