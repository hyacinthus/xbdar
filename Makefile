V := @
LDFLAGS :=

cmds = cmd_load_data cmd_start_swagger

all: app docs $(cmds)
cmds: $(cmds)

app: xbdar

xbdar:
	$(V)echo + $@
	$(V)go build -ldflags "$(LDFLAGS) $(ldflags)" -o $@ .

docs:
	$(V)echo + $@
	$(V)scripts/genDocs.sh

cmd_start_swagger: docs

$(cmds): cmd_%:
	$(V)echo + $@
	$(V)go build -ldflags "$(LDFLAGS) $(ldflags)" -o $@ ./cmds/$@


.PHONY: xbdar docs $(cmds) clean

clean:
	$(V)echo - xbdar
	$(V)rm -f xbdar
	$(V)echo - $(cmds)
	$(V)rm -f $(cmds)

