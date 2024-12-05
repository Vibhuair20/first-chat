package redisrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// function registerNewUser
// err is used for error handeling

// error while saving your username and password
// in this we check if the username is unique or not if is not unique it will deete it
func RegisterNewuser(username, password string) error {

	err := redisClient.Set(context.Background(), username, password, 0).Err()
	if err != nil {
		log.Println("error while adding in set", err)
		return err

	}
	// redisClient.Del(context.Background(), username)
	// it saves usind sadd value pair and if some occurs while saving it then this displays
	err = redisClient.SAdd(context.Background(), userSetKey(), username).Err()
	if err != nil {
		log.Println("error while adding user in set", err)
		redisClient.Del(context.Background(), username)

		return err

	}
	return nil

}

func isUserExt(username string) bool {
	return redisClient.SIsMember(context.Background(), userSetKey(), username).Val()
}

// it checks if the value of username and password is equal or not in the redis
//it retreives the value from the redis and then checks it with the value in the redis
// these are two steps
//no need to use usersetkey cause we are directly matching it with the password

// func isUserAuth(username string)bool{
// p := redisClient.Get(context.Background(), username).Val()

// if !strings.EqualFold(p, password){
// 		return fmt.Errorf("invalid username or password")
// }
// 	return nil
// }

func isUserAuth(username, password string) error {
	storedPasswordHash := redisClient.Get(context.Background(), username).Val()
	if storedPasswordHash == "" {
		return fmt.Errorf(("invalid username or passwors"))
	}

	if !checkPasswordHash(password, storedPasswordHash) {
		return fmt.Errorf("invalid username or passowrd")
	}
	return nil
}

// compare hash functions
// it will help prevent rainbow attacks by adding salts in between
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


//update the list of contacts for a user and save it in database called redis
func UpdateContactList(username, contact string ) error{   // adding contact -> username  //redis has a fancy list sorted set -> to do list
    zs := &redis.Z{Score: float64(time.Now().Unix()), Member: contact}

    // ZADD add contact to user sorted data

    err := redisClient.ZAdd(context.Background(), ContactListZkey(username), zs).Err()

    if err != nil{
        log.Println("error while updating contact list. username: ", username, "contact:", contact, err)
        return err
    }

    return nil
}

// finction to create a chat key
func createChatKey(c *model.Chat) (string, error){
    // generating a chat key
    chatKey := chatKey()
    fmt.Println("the chatkey is:", chatKey)
    by, err := json.Marshal(c)

    // storing in redis
    res, err

}
