
# Gathering App

This App simulation of real Gathering App, Where some user can be invited to an event





## Preparation

This Project can run on docker, so before start, make sure you have docker running in your machine. after that you just need to `compose up`. if the services already run. just check `localhost:1212`

```bash
  docker compose up
```



## API Reference

### Member

#### Register Member

```http
  POST /member/register

  {
	"first_name":"First-XA",
	"last_name": "Last-X",
	"email":"first@last.com"
  }
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `first_name` | `string` | **Required**. |
| `last_name` | `string` | **Required**. |
| `email` | `string` | **Required**. |

#### Find Member By ID

```http
  GET /member/findByID?id=${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


#### Update Member

```http
  POST /member/update

  {
	"id": 1699448427928626125,
	"first_name":"First XX",
	"last_name": "Last",
	"email":"first@last.com"
}
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `int64` | **Required**. |
| `first_name` | `string` | **Required**. |
| `last_name` | `string` | **Required**. |
| `email` | `string` | **Required**. |

#### Delete Member By ID

```http
  POST /member/deleteByID?id=${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


### Gathering

#### Create Gathering

```http
  POST /gathering/create

  {
	"name": "gathering-XA",
	"location": "locc",
	"type": 1,
	"creator": 321
  }
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required**. |
| `location` | `string` | **Required**. |
| `type` | `int` | **Required**. |
| `creator` | `int64` | **Required**. |

#### Find Gathering By ID

```http
  GET /gahering/findByID?id=${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


#### Update Gathering

```http
  POST /gathering/update

  {
	"id": 1699448427928626125,
	"name": "gathering-XA",
	"location": "locc",
	"type": 1,
	"creator": 321
}
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `int64` | **Required**. |
| `name` | `string` | **Required**. |
| `location` | `string` | **Required**. |
| `type` | `int` | **Required**. |
| `creator` | `int64` | **Required**. |

#### Delete Gathering By ID

```http
  POST /gathering/deleteByID?id=${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


### Invitation

#### Invite Member to Gathering

```http
  POST /invitation/invite

  {
	"member_id": 1699505339894172406,
	"gathering_id": 1699505352293158891,
	"status": 1
  }
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `member_id` | `int64` | **Required**. |
| `gathering_id` | `int64` | **Required**. |
| `status` | `int` | **Required**. |

#### Find Invitation By ID

```http
  GET /invitation/findByID?id=${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


#### Update Invitation

```http
  POST /gathering/update

  {
	"id": 1699448427928626125,
	"member_id": 1699505339894172406,
	"gathering_id": 1699505352293158891,
	"status": 1
}
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `int64` | **Required**. |
| `member_id` | `int64` | **Required**. |
| `gathering_id` | `int64` | **Required**. |
| `status` | `int` | **Required**. |

#### Delete Invitation By ID

```http
  POST /invitation/deleteByID?id=${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |
