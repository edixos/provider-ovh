package main

import "testing"

func TestParseStdLogLine(t *testing.T) {
	cases := map[string]struct {
		line      string
		wantLevel string
		wantMsg   string
	}{
		"Debug":    {line: "[DEBUG] Logged in on OVH API\n", wantLevel: "DEBUG", wantMsg: "Logged in on OVH API"},
		"Warn":     {line: "[WARN] something odd", wantLevel: "WARN", wantMsg: "something odd"},
		"NoLevel":  {line: "plain message\n", wantLevel: "", wantMsg: "plain message"},
		"EmptyTag": {line: "[] message", wantLevel: "", wantMsg: "[] message"},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			level, msg := parseStdLogLine(tc.line)
			if level != tc.wantLevel || msg != tc.wantMsg {
				t.Errorf("parseStdLogLine(%q) = (%q, %q), want (%q, %q)", tc.line, level, msg, tc.wantLevel, tc.wantMsg)
			}
		})
	}
}
