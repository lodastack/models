package models

type Report struct {
	UUID        []string `json:"uuid"`
	NewIPList   []string `json:"newiplist"`
	OldIPList   []string `json:"oldiplist"`
	Ns          []string `json:"ns"`
	Version     string   `json:"version"`
	NewHostname string   `json:"newhostname"`
	OldHostname string   `json:"oldhostname"`
	AgentType   string   `json:"agenttype"`
	Update      bool     `json:"update"`
}
