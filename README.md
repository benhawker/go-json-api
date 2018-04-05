# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:

   revel run myapp

### Go to http://localhost:9000/ and you'll see:

    "It works"


# build and run the image
docker build -t sp .
docker run -it -p 9000:9000 sp

### Future Approach

Having worked through this I would simplify the approach in the future 
Refactor using STI

UserRelationship
====
RelatingUserID
RelatedUserID
Type[friend, block, etc]


#### User Stories

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

