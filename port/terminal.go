package port

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Port interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	Close() error
}

type Terminal struct {
	port   Port
	logger *log.Logger
}

func NewTerminal(logger *log.Logger, port Port) Terminal {
	return Terminal{logger: logger, port: port}
}

func (p Terminal) Close() {
	if p.port != nil {
		if err := p.port.Close(); err != nil {
			p.logger.Printf("close port: %v", err)
		}
	}
}

func (p Terminal) Write(cmd string, args ...string) string {
	out, err := p.WriteE(cmd, args...)
	if err != nil {
		fmt.Printf("error: write %s command: %v\n", cmd, err)
		p.Close()
		os.Exit(1)
	}
	return out
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
	buff := make([]byte, 10)
	for {
		n, err := p.port.Read(buff)
		p.logger.Printf("read: %d bytes", n)
		if err != nil {
			if err == io.EOF {
				p.logger.Print("read: EOF")
				break
			}
			return "", fmt.Errorf("error: read: %w", err)
		}
		if n == 0 {
			p.logger.Print("read: EOF")
			break
		}
	}
	return strings.TrimSpace(string(buff)), nil
}
