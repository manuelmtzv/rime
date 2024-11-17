package constants

const (
	AuthRegistered            = "Auth.Registered"
	AuthLoggedIn              = "Auth.LoggedIn"
	AuthWrongCredentials      = "Auth.WrongCredentials"
	AuthUserExists            = "Auth.UserExists"
	AuthUserNotFound          = "Auth.UserNotFound"
	AuthTokenExpired          = "Auth.TokenExpired"
	AuthTokenInvalid          = "Auth.TokenInvalid"
	AuthCannotRetrieveSession = "Auth.CannotRetrieveSession"
	AuthTokenCreationFailed   = "Auth.TokenCreationFailed"
	AuthCannotCreateUser      = "Auth.CannotCreateUser"

	UserUpdated  = "User.Updated"
	UserFollowed = "User.Followed"
	UserNotFound = "User.NotFound"

	WritingCreated   = "Writing.Created"
	WritingPublished = "Writing.Published"
	WritingNotFound  = "Writing.NotFound"

	TagCreated  = "Tag.Created"
	TagUpdated  = "Tag.Updated"
	TagDeleted  = "Tag.Deleted"
	TagNotFound = "Tag.NotFound"

	CommentNotFound = "Comment.NotFound"

	HealthOk = "Health.Ok"
)
