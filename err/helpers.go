package err

import (
	httpEvent "github.com/taubyte/go-sdk/http/event"
	pubsub "github.com/taubyte/go-sdk/pubsub/node"
)

func Write(h httpEvent.Event, err error) uint32 {
	h.Write(unsafeMarshalError(err))
	h.Return(500)

	return 1
}

func Publish(channel *pubsub.ChannelObject, err error) uint32 {
	channel.Publish(unsafeMarshalError(err))
	return 1
}

func marshalError(err error) ([]byte, error) {
	errStruct := &errStruct{
		Error: err,
	}

	return errStruct.MarshalJSON()
}

func unsafeMarshalError(err error) []byte {
	data, _err := marshalError(err)
	if _err != nil {
		panic(_err)
	}

	return data
}
