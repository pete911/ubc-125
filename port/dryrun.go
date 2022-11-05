package port

import (
	"fmt"
	"strings"
)

type dryRun struct {
	cmd *strings.Reader
}

func OpenDryRun() Port {
	return &dryRun{}
}

func (s *dryRun) Read(b []byte) (n int, err error) {
	return s.cmd.Read(b)
}

func (s *dryRun) Write(b []byte) (n int, err error) {
	// simulate response, prepare <CMD>,OK response for read operation
	in := strings.TrimSpace(strings.Split(string(b), ",")[0])
	s.cmd = strings.NewReader(fmt.Sprintf("%s,OK\n\r", in))
	return len(b), nil
}

func (s *dryRun) Close() error {
	return nil
}
