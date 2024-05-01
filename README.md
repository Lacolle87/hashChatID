# HashID

HashChatID is a Go package that provides a simple utility function for hashing integer chat IDs using the SHA-256 cryptographic hash function. This package is useful for scenarios where you need to anonymize or obfuscate chat IDs or similar identifiers, ensuring privacy and security in your applications.

## Features

- Hashes integer chat IDs using SHA-256 for secure anonymization.
- Converts integer chat IDs to hexadecimal strings for ease of use and storage.
- Lightweight and easy to integrate into existing Go projects.

## Installation

To use HashID in your Go project:

```go
package main

import (
	"fmt"
	hid "github.com/Lacolle87/hashID"
)

func main() {
	chatID := int64(1234567890)
	hashedID := hid.HashID(chatID)
	fmt.Println("Hashed Chat ID:", hashedID)
}
```
