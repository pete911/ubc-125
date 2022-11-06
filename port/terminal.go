package port

import (
	"fmt"
	"log"
	"strings"
)

type Port interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	Close() error
}

type Terminal struct {
	logger      *log.Logger
	port        Port
	programMode bool
}

func NewTerminal(logger *log.Logger, port Port, programMode bool) (Terminal, error) {
	t := Terminal{logger: logger, port: port, programMode: programMode}
	if !programMode {
		return t, nil
	}
	if _, err := t.WriteE("PRG"); err != nil {
		t.Close()
		return Terminal{}, err
	}
	logger.Print("entered program mode")
	return t, nil
}

func (p Terminal) Close() {
	if p.port == nil {
		p.logger.Print("close port: port is nil, nothing to close")
		return
	}
	if p.programMode {
		out, err := p.WriteE("EPG")
		if err != nil {
			p.logger.Printf("error: close port: exit program mode: %v", err)
		}
		p.logger.Printf("exit program mode: %s", out)
	}
	if err := p.port.Close(); err != nil {
		p.logger.Printf("error: close port: %v", err)
	}
	p.logger.Print("port closed")
}

func (p Terminal) WriteE(cmd string, args ...string) (string, error) {
	if p.port == nil {
		return "", fmt.Errorf("port is not opened")
	}

	p.logger.Printf("input: %s %v", cmd, args)
	request := cmd
	if len(args) != 0 {
		request = fmt.Sprintf("%s,%s", cmd, strings.Join(args, ","))
	}

	p.logger.Printf("request: %s", request)
	if err := p.write(request); err != nil {
		return "", err
	}

	response, err := p.read()
	if err != nil {
		return "", err
	}
	p.logger.Printf("response: %s", response)

	out := strings.TrimPrefix(response, fmt.Sprintf("%s,", cmd))
	p.logger.Printf("output: %s", out)
	return out, nil
}

func (p Terminal) write(in string) error {
	in = fmt.Sprintf("%s\n\r", in)
	_, err := p.port.Write([]byte(in))
	return err
}

func (p Terminal) read() (string, error) {
	buff := make([]byte, 100)
	for {
		n, err := p.port.Read(buff)
		p.logger.Printf("read: %d bytes", n)
		if err != nil {
			return "", fmt.Errorf("error: read: %w", err)
		}
		if n == 0 || strings.Contains(string(buff[:n]), "\r") {
			return strings.TrimSpace(string(buff[:n])), nil
		}
	}
}
