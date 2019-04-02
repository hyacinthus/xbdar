V := @
LDFLAGS :=

cmds = cmd_load_data cmd_start_swagger

all: xbdar docs $(cmds)
cmds: $(cmds)

xbdar:
	$(V)echo + $@
	$(V)go build -ldflags "$(LDFLAGS) $(ldflags)" -o $@ .

docs:
	$(V)echo + $@
	$(V)swag init

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
	$(V)echo - docs/
	$(V)rm -rf docs/

