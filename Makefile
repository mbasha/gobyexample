# Default target to run the program
run:
	go run examples/*.go $(example)

# Clean up any temporary files (if needed in the future)
clean:
	rm -f *.out