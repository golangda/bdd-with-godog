package stepdefs

import (
    "fmt"
    "github.com/cucumber/godog"
)

type calculator struct {
    inputs []int
    result int
}

func (c *calculator) iHaveEnteredIntoTheCalculator(arg1 int) error {
    c.inputs = append(c.inputs, arg1)
    return nil
}

func (c *calculator) iPressAdd() error {
    c.result = 0
    for _, v := range c.inputs {
        c.result += v
    }
    return nil
}

func (c *calculator) theResultShouldBeOnTheScreen(expected int) error {
    if c.result != expected {
        return fmt.Errorf("expected %d, but got %d", expected, c.result)
    }
    return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
    calc := &calculator{}

    ctx.Step(`^I have entered (\d+) into the calculator$`, calc.iHaveEnteredIntoTheCalculator)
    ctx.Step(`^I press add$`, calc.iPressAdd)
    ctx.Step(`^the result should be (\d+) on the screen$`, calc.theResultShouldBeOnTheScreen)
}