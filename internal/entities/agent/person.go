package agent

import (
	i_ "coolsim/internal/commons/identity"
	a_ "coolsim/internal/entities/action"
	t_ "coolsim/internal/types"
	au_ "coolsim/internal/utils/action"

	"go.uber.org/zap"
)

type Person struct {
	*i_.Identity
	home         t_.Environment
	workPlace    t_.Environment
	location     t_.Environment // current location
	capabilities []t_.Action
}

func NewPerson(home t_.Environment, work t_.Environment) *Person {
	// @todo get capabilities from config file
	capabilities := getCapabilities()
	return &Person{
		i_.NewIdentity(),
		home,
		work,
		home, // spawn at home by default
		capabilities,
	}
}

func (person *Person) GetHome() t_.Environment {
	return person.home
}

func (person *Person) GetWorkPlace() t_.Environment {
	return person.workPlace
}

func (person *Person) MoveIn(destination t_.Environment) {
	// person.location.Leave(p) // @todo leave last location ?
	person.location = destination
	// person.location.Come(p) // @todo come to destination ?
}

func (person *Person) Perceive() {
	perception := person.location.GetAPerception()
	zap.S().Infof("The agent %s percepted '%s'",
		person.GetUID(),
		perception.Description(),
	)

	actions := evaluate(
		perception.GetActions(),
		person.capabilities,
	)
	person.Perform(actions)
}

func (person *Person) Perform(actions []t_.Action) {
	for _, action := range actions {
		zap.S().Infof("The agent %s %s",
			person.GetUID(),
			action.Description(),
		)
		action.Do(person)
	}
}

//
// Private methods
//

func evaluate(allActions, capabilities []t_.Action) []t_.Action {
	return au_.Intersection(allActions, capabilities)
}

// @todo pass capabilities file?
func getCapabilities() []t_.Action {
	return []t_.Action{
		a_.NewSleep(),
		a_.NewRest(),
		a_.NewGoOffice(),
		a_.NewBackHome(),
	}
}
