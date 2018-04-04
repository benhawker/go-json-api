package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type ResponseBody struct {
	Success bool `json:"success"`
}

func (c App) Index() revel.Result {
	available_routes := []string{
		"GET     /users",
		"GET     /friendships",
		"GET     /friendships/:email",
		"POST    /friendships",
		"POST    /messages",
		"POST    /blocks",
		"GET     /notification_subscriptions",
		"POST    /notification_subscriptions",
	}

	routesHelp := map[string][]string{"Available routes:": available_routes}
	return c.RenderJSON(routesHelp)
}
