package adapter

type SummatorInterface interface {
	ExecuteSummator(input, output string) error
}

type SummatorAdapater struct {
	executor *Executor
}

func NewSummatorAdapater(executor *Executor) SummatorAdapater {
	return SummatorAdapater{executor: executor}
}

func (a *SummatorAdapater) ExecuteSummator(input, output string) error {
	app := "./summator"

	return a.executor.Execute(app, input, output)
}
