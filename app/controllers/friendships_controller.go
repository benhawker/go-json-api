package controllers

import (
  "net/http"

  "fmt"
  "github.com/benhawker/go-json-api/app/models"
  "github.com/benhawker/go-json-api/app/services"
  "github.com/revel/revel"

)

type FriendshipsController struct {
  *revel.Controller
  services.Database
}

func (c FriendshipsController) Index() revel.Result {
  f := make([]models.Friendship, 0)

  if err := c.Gorm.Find(&f).Error; err != nil {
    // if err := c.Gorm.Where("requester_id = ? OR receiver_id = ? ", 999, 888).Find(&f).Error; err != nil {
    panic(err)
  }

  json := make([]models.FriendshipJSON, 0)
  for _, friendship := range f {
    json = append(json, models.NewFriendshipJSON(friendship))
  }

  return c.RenderJSON(json)
}

// GET http://localhost:9000/friendships/test@email.com

// Expected response body:
//
// {
//   "success": true,
//   "friends" :
//     [
//       'john@example.com'
//     ],
//   "count" : 1   
// }

// type Response struct {
//   Code    int    `json:"code"`
//   Body    ShowBody `json:"message"`
// }


type Response struct {
  Success bool `json:"success"`
  Friends []string `json:"friends"`
  Count int `json:"count"`
}


func (c FriendshipsController) Show(email string) revel.Result {
  user := models.User{}

  if err := c.Gorm.Where("email = ?", email).First(&user).Error; err != nil {
    c.Response.Status = http.StatusNotFound
    return c.RenderJSON(fmt.Sprintf("We have no user by this email: %s", email))
  }

  friendships := make([]models.Friendship, 0)
  if err := c.Gorm.Where("requester_id = ? OR receiver_id = ? ", user.Id, user.Id).Find(&friendships).Error; err != nil {
    panic(err)
  }

  userIds := []int{}
  for _, f := range friendships {
    if user.Id != uint(f.ReceiverId) {
      userIds = append(userIds, f.ReceiverId)
    }

    if user.Id != uint(f.ReceiverId) {
      userIds = append(userIds, f.RequesterId)
    }
  }

  users := make([]models.User, 0)
  if err := c.Gorm.Where("id in (?)", userIds).Find(&users).Error; err != nil {
    c.RenderJSON(err)
  }

  emails := []string{}
  for _, u := range users {
    emails = append(emails, u.Email)
  }


  response := Response{Success: true, Friends: emails, Count: len(users)}
  return c.RenderJSON(response)
}



type RequestBody struct {
  Friends []string `json:"friends"`
}

// Expected request body:
// {
//   friends:
//     [
//       'andy@example.com',
//       'john@example.com'
//     ]
// }
func (c FriendshipsController) Create() revel.Result {
  var requestBody RequestBody
  c.Params.BindJSON(&requestBody)
  user := models.User{}

  // Check requester exists
  if err := c.Gorm.Where("email = ?", requestBody.Friends[0]).First(&user).Error; err != nil {
    c.Response.Status = http.StatusBadRequest
    return c.RenderJSON(fmt.Sprintf("We don't recognise the requesting user: %s", requestBody.Friends[0]))
  }

  // Check receiver exists
  if err := c.Gorm.Where("email = ?", requestBody.Friends[1]).First(&user).Error; err != nil {
    c.Response.Status = http.StatusBadRequest
    return c.RenderJSON(fmt.Sprintf("We don't recognise the receiving user: %s", requestBody.Friends[1]))
  }

  //Create the Friendship
  // TODO

  // Render 200
  success := map[string]bool{"success": true}
  c.Response.Status = http.StatusOK
  return c.RenderJSON(success)
}
