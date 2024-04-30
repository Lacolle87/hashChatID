package hashchatid

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func HashChatID(id int64) string {
	idString := strconv.FormatInt(id, 10)
	hasher := sha256.New()

	_, err := hasher.Write([]byte(idString))
	if err != nil {
		return fmt.Sprintf("hashchatid: %v", err)
	}

	hashedID := hex.EncodeToString(hasher.Sum(nil))

	return hashedID
}
