package redisrepo

import(
	"fmt"
	"encoding/json"
	"time"
	"strings"
	"log"
	"context"

	"github.com/go-redis/redis/v8"

	"gochatapp/model"

)

// function registerNewUser
