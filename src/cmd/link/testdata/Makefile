ALL=\
	autosection.6\
	autoweak.6\
	dead.6\
	hello.6\
	layout.6\
	pclntab.6\

all: $(ALL)

%.6: %.s
	go tool 6a $*.s

pclntab.s: genpcln.go
	go run genpcln.go >pclntab.s
