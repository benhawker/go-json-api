package services

import "github.com/revel/revel"

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*Database).Open, revel.BEFORE)
	revel.InterceptMethod((*Database).Close, revel.FINALLY)
}
