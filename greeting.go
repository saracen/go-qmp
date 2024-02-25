package qmp

import (
	"fmt"
	"strings"
)

type Greeting struct {
	QMP struct {
		Version      VersionInfo `json:"version"`
		Capabilities []QMPCapability
	} `json:"qmp"`
}

type VersionInfo struct {
	QEMU    VersionTriple `json:"qemu"`
	Package string        `json:"package"`
}

type VersionTriple struct {
	Micro int `json:"micro"`
	Minor int `json:"minor"`
	Major int `json:"major"`
}

type QMPCapability string

func (c *Client) negotiate() error {
	// Potentially a bug in the qemu storage daemon, but QMP has two modes of operation
	// as documented here: https://www.qemu.org/docs/master/interop/qmp-spec.html#capabilities-negotiation
	//
	// In negotiation mode, it sounds like asynchronous messages should not be delivered, now this perhaps
	// means those that are sent, rather than events received, but I've witnessed events being received
	// during negotiation, so you cannot assume that the first message you recieve will be a greeting,
	// and that the second message you recieve will be a response to the qmp_capabilities command.
	//
	// To get around this issue, we've configured the json decoder to error on unknown fields, and
	// then we explicitly check for this error message: if received, we keep trying.
	var greeting Greeting
	for {
		err := c.reader.Decode(&greeting)
		if err == nil {
			break
		}
		if strings.Contains(err.Error(), "unknown field") {
			continue
		}
		return fmt.Errorf("reading greeting: %w", err)
	}

	c.greeting = greeting

	if err := c.writer.Encode(map[string]any{
		"execute":   "qmp_capabilities",
		"arguments": map[string]any{"enable": []string{}},
	}); err != nil {
		return fmt.Errorf("sending capabilities: %w", err)
	}

	var resp response
	for {
		err := c.reader.Decode(&resp)
		if err == nil && resp.Error != nil {
			err = resp.Error
			break
		}
		if err == nil {
			break
		}
		if strings.Contains(err.Error(), "unknown field") {
			continue
		}
		return fmt.Errorf("reading capabilities: %w", err)
	}

	return nil
}

func (c *Client) Greeting() Greeting {
	return c.greeting
}
