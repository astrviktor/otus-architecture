package softstop

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw07/object"
	"testing"
)

func TestSoftStopOk(t *testing.T) {
	obj := object.New()

	softStopChannel := make(chan struct{}, 1)
	testStopChannel := make(chan struct{}, 1)

	go func() {
		<-softStopChannel
		close(testStopChannel)
	}()

	obj.SetProperty("SoftStopChannel", softStopChannel)

	softStopAdapter := NewSoftStopAdapter(obj)
	softStopCommand := NewSoftStopCommand(softStopAdapter)

	err := softStopCommand.Execute()
	require.Equal(t, nil, err)

	<-testStopChannel
}
