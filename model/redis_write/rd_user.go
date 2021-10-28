package redis_write

/*
  用户信息存入redis:  hashmap
  user: username: userInfo
*/

var (
	UserKey = "user"
)

// CreateUser Redis create a user
func CreateUser(username string, newUserBytes []byte) (int, error) {
	flag, err := rs.HPut(UserKey, username, newUserBytes)
	return flag, err
}
