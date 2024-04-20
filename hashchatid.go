package hashchatid

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashChatID(chatID int64) string {
	strChatID := fmt.Sprintf("%d", chatID)
	hasher := sha256.New()

	_, err := hasher.Write([]byte(strChatID))
	if err != nil {
		return ""
	}

	hashedChatID := hex.EncodeToString(hasher.Sum(nil))

	return hashedChatID
}
