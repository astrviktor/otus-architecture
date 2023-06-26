package hardstop

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw12/object"
	"testing"
)

func TestHardStopOk(t *testing.T) {
	obj := object.New()

	hardStopChannel := make(chan struct{}, 1)
	testStopChannel := make(chan struct{}, 1)

	go func() {
		<-hardStopChannel
		close(testStopChannel)
	}()

	obj.SetProperty("HardStopChannel", hardStopChannel)

	hardStopAdapter := NewHardStopAdapter(obj)
	hardStopCommand := NewHardStopCommand(hardStopAdapter)

	err := hardStopCommand.Execute()
	require.Equal(t, nil, err)

	<-testStopChannel
}
