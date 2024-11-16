package redisrepo

import (
	"fmt"
	"time"
)

//three keys
// we need to store for unique identity
// session key
// client string -> accepts from client argument as string
// returns as string

// if we are defining two return types we need to mention both in the retirn cases
func sessionKey(client string) (string, error) {
	if client == "" {
		return "", fmt.Errorf("client cannot be empty")
	}
	return "sessions#" + client, nil
}

// we ise # cause if wants to combine anything

// user key
func userKey() string {
	return "users#"
}

// chatkey

func chatKey() string {
	return fmt.Sprintf("chat#%d", time.Now().UnixMilli())
}

func chatIndex() string {
	return "idx#chats"
}

func ContactListZkey(username string) string {
	return "contacts#" + username
}
