package main

func startsWith(input, prefix string) bool {
	return len(input) >= len(prefix) && input[:len(prefix)] == prefix
}
