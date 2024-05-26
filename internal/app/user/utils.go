package user

import "context"

type userKey string

const user = userKey("user")

func ContextWithUser(c context.Context, u string) context.Context {
	return context.WithValue(c, user, u)
}

func UserFromContext(ctx context.Context) (string, bool) {
	u, ok := ctx.Value(user).(string)
	return u, ok
}
