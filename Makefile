addercli: cmd/addercli/main.go
	go build -o $@ $^

.PHONY: clean
clean:
	rm -f addercli