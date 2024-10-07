package command

import (
	"net/http"
)

func NewHttpCommand(r *http.Request, b func(*http.Request) Command) *CommandWrapper {
	c := b(r)
	return &CommandWrapper{Command: c}
}

type CommandWrapper struct {
	Command
}

func (cb *CommandWrapper) Auth() *CommandWrapper {
	cb.Command = WithAuth(cb.Command)
	return cb
}

func (cb *CommandWrapper) Log() *CommandWrapper {
	cb.Command = WithLog(cb.Command)
	return cb
}

func (cb *CommandWrapper) Audit(r *http.Request) *CommandWrapper {
	cb.Command = WithAudit(cb.Command, r)
	return cb
}

// func (cb *CommandWrapper) Validator() *CommandWrapper {
// 	cb.Command = WithValid(cb.Command)
// 	return cb
// }

// func (cb *CommandWrapper) Cache(exp time.Duration, clear ...string) *CommandWrapper {
// 	cb.Command = WithCache(cb.Command, exp, clear...)
// 	return cb
// }

// func (cb *CommandWrapper) FillableCard() *CommandWrapper {
// 	cb.Command = WithFillableCard(cb.Command)
// 	return cb
// }

// func (cb *CommandWrapper) RequestWriter(w http.ResponseWriter, r *http.Request) *CommandWrapper {
// 	cb.Command = WithRequestWriterCommand(cb.Command, w, r)
// 	return cb
// }

// func (cb *CommandWrapper) RateLimit(limit int, w http.ResponseWriter, r *http.Request) *CommandWrapper {
// 	cb.Command = WithRateLimit(cb.Command, limit, w, r)
// 	return cb
// }
