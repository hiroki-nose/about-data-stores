package neptune

import (
	"fmt"

	"github.com/qasaur/gremgo"
)

type NeptuneHandler struct {
	g gremgo.Client
}

func NewNeptuneHandler() (*NeptuneHandler, error) {
	const Address = "wss://{hostname}:8182"
	errs := make(chan error)
	go func(chan error) {
		err := <-errs
		fmt.Println("Lost connection to the database")
		fmt.Println(err.Error())
	}(errs) // Example of connection error handling logic

	dialer := gremgo.NewDialer(Address) // Returns a WebSocket dialer to connect to Gremlin Server
	g, err := gremgo.Dial(dialer, errs)
	if err != nil {
		fmt.Println("Dial fail")
		fmt.Println(err.Error())
		return nil, err
	}
	return &NeptuneHandler{
		g: g,
	}, nil
}
func (c *NeptuneHandler) Execute(query string, bindings, rebindings map[string]string) (interface{}, error) {
	return c.g.Execute(query, bindings, rebindings)
}
