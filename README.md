# Social Application - Go

Technology choices:
* [Go](https://golang.org/)
* [Revel](https://revel.github.io/). A chance to experiment with a new framework.
* [PostgreSQL](https://www.postgresql.org/)

### Setup the application

Install Go
```
brew install go
```

You'll also need to ensure your GOPATH is setup correctly.

Install Revel and Command Line Tools.
```
go get github.com/revel/revel
go get github.com/revel/cmd/revel
```

You may also need to ensure the `revel` cmd is in your $PATH. [See Revel guides here](https://revel.github.io/tutorial/gettingstarted.html)

Install Gorm
```
go get github.com/jinzhu/gorm
```

Install PQ
```
go get github.com/lib/pq
```

Clone the repo
```
git@github.com:benhawker/go-json-api.git
```

cd into the directory
```
$ /go-json-api
```

Run the application
```
revel run
```


### Other Approaches

Another approach may include using a Table that makes use of Single Table Inheritance. Example:

```
UserRelationship
====
RelatingUserID
RelatedUserID
Type[friendship, block, notification_subscription etc]
```

For the purposes of the user stories defined here, another approach could have included using a noSQL DB that includes `User` documents with the below approach. 

```
Users
====
Email
Friends []
Subscribers []
etc ...
```


### User Stories

**1. As a user, I need an API to create a friend connection between two email addresses.**

The API should receive the following JSON request:
```
POST localhost:9000/friendships
{
  friends:
    [
      'andy@example.com',
      'john@example.com'
    ]
}
```

The API should return the following JSON response on success:

```
{
  "success": true
}
```

**2. As a user, I need an API to retrieve the friends list for an email address.**

The API should receive the following JSON request:

```
GET localhost:9000/friendships/:email
e.g. GET localhost:9000/friendships/andy@example.com
```

The API should return the following JSON response on success:

```
{
  "success": true,
  "friends" :
    [
      'john@example.com'
    ],
  "count" : 1   
}
```


**3. As a user, I need an API to retrieve the common friends list between two email addresses.**

The API should receive the following JSON request:


```
TODO: STILL TO BE IMPLEMENTED
GET localhost:9000/friendships/common?friends=[andy@example.com,john@example.com]
```

The API should return the following JSON response on success:

```
{
  "success": true,
  "friends" :
    [
      'common@example.com'
    ],
  "count" : 1   
}
```


**4. As a user, I need an API to subscribe to updates from an email address.**

The API should receive the following JSON request:

```
POST localhost:9000/notification_subscriptions

{
  "requestor": "lisa@example.com",
  "target": "john@example.com"
}
```

The API should return the following JSON response on success:

```
{
  "success": true
}
```


**5. As a user, I need an API to block updates from an email address.**

Suppose "andy@example.com" blocks "john@example.com".

- if they are connected as friends, then "andy" will no longer receive notifications from "john"
- if they are not connected as friends, then no new friends connection can be added

The API should receive the following JSON request:

```
POST localhost:9000/blocks

{
  "requestor": "andy@example.com",
  "target": "john@example.com"
}
```

The API should return the following JSON response on success:

```
{
  "success": true
}
```

**6. As a user, I need an API to retrieve all email addresses that can receive updates from an email address.**

Eligibility for receiving updates from i.e. "john@example.com":
- has not blocked updates from "john@example.com", and
- at least one of the following:
  - has a friend connection with "john@example.com"
  - has subscribed to updates from "john@example.com"
  - has been @mentioned in the update

The API should receive the following JSON request:

```
POST localhost:9000/messages

{
  "sender":  "john@example.com",
  "text": "Hello World! kate@example.com"
}
```

The API should return the following JSON response on success:

```
{
  "success": true
  "recipients":
    [
      "lisa@example.com",
      "kate@example.com"
    ]
}
```

