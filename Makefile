fbx-to-tileset:cmd/fbx-to-tileset.go COPY
	go build -o build/fbx-to-tileset cmd/fbx-to-tileset.go

COPY:
	chmod 755 third_party/bin/*
	cp third_party/bin/* build/

.PHONY:clean
clean:
	-rm build/fbx-to-tileset