package main

type AnsibleModule interface {
	HasPossibility() []string
	Generate() []string
}

type Module struct {
	Possibility []string
}

