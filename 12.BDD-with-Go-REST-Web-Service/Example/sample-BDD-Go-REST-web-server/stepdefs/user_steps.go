package stepdefs

import (
    "bytes"
    "fmt"
    "io"
    "net/http"

    "github.com/cucumber/godog"
)

type userSuite struct {
    response *http.Response
}

func (s *userSuite) iCreateUser(name, email string) error {
    body := fmt.Sprintf(`{"name":"%s","email":"%s"}`, name, email)
    resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer([]byte(body)))
    s.response = resp
    return err
}

func (s *userSuite) iFetchAllUsers() error {
    resp, err := http.Get("http://localhost:8080/users")
    s.response = resp
    return err
}

func (s *userSuite) iFetchUserById(id int) error {
    resp, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%d", id))
    s.response = resp
    return err
}

func (s *userSuite) iUpdateUser(id int, name, email string) error {
    body := fmt.Sprintf(`{"name":"%s","email":"%s"}`, name, email)
    client := &http.Client{}
    req, _ := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/users/%d", id), bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    s.response = resp
    return err
}

func (s *userSuite) iDeleteUser(id int) error {
    client := &http.Client{}
    req, _ := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/users/%d", id), nil)
    resp, err := client.Do(req)
    s.response = resp
    return err
}

func (s *userSuite) responseCodeShouldBe(expected int) error {
    if s.response == nil {
        return fmt.Errorf("no response received")
    }
    if s.response.StatusCode != expected {
        b, _ := io.ReadAll(s.response.Body)
        return fmt.Errorf("expected %d, got %d: %s", expected, s.response.StatusCode, string(b))
    }
    return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
    s := &userSuite{}
    ctx.Step(`^I create a user with name "([^"]*)" and email "([^"]*)"$`, s.iCreateUser)
    ctx.Step(`^I fetch all users$`, s.iFetchAllUsers)
    ctx.Step(`^I fetch user with ID (\d+)$`, s.iFetchUserById)
    ctx.Step(`^I update user with ID (\d+) to name "([^"]*)" and email "([^"]*)"$`, s.iUpdateUser)
    ctx.Step(`^I delete user with ID (\d+)$`, s.iDeleteUser)
    ctx.Step(`^the response code should be (\d+)$`, s.responseCodeShouldBe)
}