# Go-gator-feed

A multi-player command line tool for agregating RSS feed and viewing posts using GO

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postres database. You can then install `go-gator-feed` with:

````bash
go install github.com/lucasgjanot/go-gator-feed```
````

## Configuration

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disble"
}
```

Replace values with your database connection string.

## Usage

|  Command   |                   Usage                    |     Description      |
| :--------: | :----------------------------------------: | :------------------: |
| `register` |    `go-gator-feed register <username>`     |  Creates a new user  |
| `addfeed`  |    `go-gator-feed addfeed <name> <Url>`    |      Add a feed      |
|   `agg`    | `go-gator-feed agg <time_interval_string>` | Start the aggregator |
|  `browse`  |       `go-gator-feed browse [limit]`       |    View the posts    |

The are a few other commands you'll need as well:

- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database
