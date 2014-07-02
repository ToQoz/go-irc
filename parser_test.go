package irc

import (
	"fmt"
	"testing"
)

func TestParseMessage_PRIVMSG(t *testing.T) {
	m := ParseMessage(":example!~example@example.com PRIVMSG #channel :body hogehoge fooo\r\n")

	if m.Prefix != "example!~example@example.com" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "example!~example@example.com", m.Prefix))
	}

	if m.Command != "PRIVMSG" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "PRIVMSG", m.Command))
	}

	if m.Params[0] != "#channel" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "#channel", m.Params[0]))
	}

	if m.Params[1] != "body hogehoge fooo" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "body hogehoge fooo", m.Params[1]))
	}

	m = ParseMessage(":example!~example@example.com  PRIVMSG  #channel  :body hogehoge fooo\r\n")

	if m.Prefix != "example!~example@example.com" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "example!~example@example.com", m.Prefix))
	}

	// test for multiple <space>
	if m.Command != "PRIVMSG" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "PRIVMSG", m.Command))
	}

	if m.Params[0] != "#channel" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "#channel", m.Params[0]))
	}

	if m.Params[1] != "body hogehoge fooo" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "body hogehoge fooo", m.Params[1]))
	}
}

func TestParseMessage_PRIVMSG_WithoutPrefix(t *testing.T) {
	m := ParseMessage("PRIVMSG #channel :body hogehoge fooo\r\n")

	if m.Prefix != "" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "", m.Prefix))
	}

	if m.Command != "PRIVMSG" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "PRIVMSG", m.Command))
	}

	if m.Params[0] != "#channel" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "#channel", m.Params[0]))
	}

	if m.Params[1] != "body hogehoge fooo" {
		t.Error(fmt.Sprintf("(expected)=%#v, but (got)=%#v", "body hogehoge fooo", m.Params[1]))
	}
}
