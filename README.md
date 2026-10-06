# Gator

A command line tool for aggregating RSS feeds and reading posts directly from your terminal!

## Installation

You'll need the latest version of Go as well as a local Postgres database. With those ready, you can install Gator with:

```sh
go install ...
```

## Config

Once installed, create a config file named `.gatorconfig.json` in your home directory. Then use your own Postgres credentials/datbase connection string for the following:

```json
{
    "db_url": "postgres://username:password@localhost:5432/database?sslmode=disable"
}
```

## Commands

First thing you'll want to do is register a new user as many of the commands require a user to be logged in.

```sh
aggre_gator register <name>
```

Once you're registered, you can add a few feeds to aggregate the posts from.

```sh
aggre_gator addfeed <title> <url>
```

Start the aggre_gator with a time frame.

```sh
aggre_gator agg 30s
```

Then view the posts that were saved from any feed you added! If you want to limit how many are shown, add it at the end of the command.

```sh
aggre_gator browse [limit]
```

There are a few extra commands that add a bit more functionality to this core loop:

- `aggre_gator login <name>` : Log in as a previously registered user
- `aggre_gator users` : Lists all users
- `aggre_gator feeds` : Lists all feeds
- `aggre_gator follow <url>` : Follow a feed that was previously added to the database
- `aggre_gator unfollow <url>` : Unfollow a feed that was previously added to the database
- `aggre_gator reset` : This will reset the entire database. All users and feeds will be cleaned out.