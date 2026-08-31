package tcp

import (
	"gnalloy.org/gnalloy/channel"
	"gnalloy.org/transport-zerocopy"
)

func newFileRegionWriter() channel.FileRegionWriter {
	writer, err := zerocopy.NewChannelWriter(0)
	if err != nil {
		return nil
	}
	return writer
}
