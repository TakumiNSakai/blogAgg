# RSSagg

A multi-player command line tool for aggregating RSS feeds and viewing the posts.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `rssagg` with:

```bash
go install ...
```

## Config

Create a `.rssconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Usage

Create a new user:

```bash
rssagg register <name>
```

Add a feed:

```bash
rssagg addfeed <url>
```

Start the aggregator:

```bash
rssagg agg 30s
```

View the posts:

```bash
rssagg browse [limit]
```

There are a few other commands you'll need as well:

- `rssagg login <name>` - Log in as a user that already exists
- `rssagg users` - List all users
- `rssagg feeds` - List all feeds
- `rssagg follow <url>` - Follow a feed that already exists in the database
- `rssagg unfollow <url>` - Unfollow a feed that already exists in the database
