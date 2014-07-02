package irc

import (
	"strings"
	"unicode/utf8"
)

type message struct {
	Prefix  string
	Command string
	Params  []string
}

/*
	c.f. http://tools.ietf.org/html/rfc1459.html
	<message>  ::= [':' <prefix> <SPACE> ] <command> <params> <crlf>
	<prefix>   ::= <servername> | <nick> [ '!' <user> ] [ '@' <host> ]
	<command>  ::= <letter> { <letter> } | <number> <number> <number>
	<SPACE>    ::= ' ' { ' ' }
	<params>   ::= <SPACE> [ ':' <trailing> | <middle> <params> ]

	<middle>   ::= <Any *non-empty* sequence of octets not including SPACE or NUL or CR or LF, the first of which may not be ':'>
	<trailing> ::= <Any, possibly *empty*, sequence of octets not including NUL or CR or LF>

	<crlf>     ::= CR LF
*/

// ParseMessage parses irc message to struct `irc.message`
// WARNING: This doen't check validity of message
func ParseMessage(msg string) *message {
	rest := msg
	m := &message{}

	// "[':' <prefix> <SPACE> ] <command> <params> <crlf>" -> "[':' <prefix> <SPACE> ] <command> <params>"
	rest = strings.TrimSuffix(rest, "\r\n")

	// "[':' <prefix> <SPACE> ] <command> <params>" -> "<prefix>", "<command> <params>"
	if strings.HasPrefix(rest, ":") {
		s := strings.SplitN(rest, " ", 2)
		m.Prefix = strings.TrimPrefix(s[0], ":")

		rest = strings.TrimPrefix(s[1], " ")
	}

	// "<command> <params>" -> "<command>", "<params>"
	s := strings.SplitN(rest, " ", 2)
	m.Command = s[0]
	params := " " + s[1]

	// "<params>" -> []string
	tokens := []rune(params)
	for len(tokens) > 0 {
		t := tokens[0]
		tokens = tokens[1:]

		switch t {
		case ' ':
			// if previous rune space, we do nop.
			if len(m.Params) > 0 {
				if m.Params[len(m.Params)-1] == "" {
					continue
				}
			}

			m.Params = append(m.Params, "")
		case ':':
			trailing := string(tokens[utf8.RuneLen(t)-1:])
			m.Params[len(m.Params)-1] += trailing
			tokens = nil
		default:
			m.Params[len(m.Params)-1] += string(t)
		}
	}

	return m
}
