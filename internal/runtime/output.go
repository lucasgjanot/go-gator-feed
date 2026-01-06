package runtime
import "github.com/lucasgjanot/go-gator-feed/internal/database"

type Output interface {
	UserCreated(user database.User)
	UserLoggedIn(username string)
	ResetedDatabase()
	ListUsers(s * State, users []database.User)
}
