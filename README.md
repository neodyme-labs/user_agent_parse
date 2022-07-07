UserAgent Parser
===============================================

This Caddy Module allows to parse the user agent and splits it into it's components.

You can access the components via these placeholders:
- `{user_agent.name}`
- `{user_agent.version}`
- `{user_agent.os}`
- `{user_agent.os_version}`
- `{user_agent.device}`
- `{user_agent.mobile}`
- `{user_agent.tablet}`
- `{user_agent.desktop}`
- `{user_agent.bot}`
- `{user_agent.url}`

The module is uses the parser from [here](https://github.com/mileusna/useragent).

## Install

First, the [xcaddy](https://github.com/caddyserver/xcaddy) command:

```shell
$ go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
```

Then build Caddy with this Go module plugged in. For example:

```shell
$ xcaddy build --with github.com/neodyme-labs/influx_log
```