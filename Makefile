V := @
LDFLAGS :=

cmds = cmd_load_data 

all: app $(cmds)
cmds: $(cmds)

app:
	$(V)echo + $@
	$(V)go build -ldflags "$(LDFLAGS) $(ldflags)" -o $@ .

$(cmds): cmd_%:
	$(V)echo + $@
	$(V)go build -ldflags "$(LDFLAGS) $(ldflags)" -o $@ ./cmds/$@


.PHONY: app $(cmds) clean

clean:
	rm -f app $(cmds)


