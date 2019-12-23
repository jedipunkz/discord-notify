# discord-notify

## Description

Discord bot that notify when friend join voice chat channel.

## Installation

```bash
go get github.com/jedipunkz/discord-notify
```

## setup .discord-notify.yaml

Setup $HOME/discord-notify.yaml file with such variables.

```bash
cat << EOF > $HOME/.discord-notify.yaml
token: <your bot token id>
notify_channel_id: <notify channel id>
me: <your own discord name>
EOF
```

## Start app

```bash
nohup /some/path/to/discord-notify
```

## Reference

- [Discord で Voice チャンネルへの入室を検知する bot (ついでに Google Home で通知](https://qiita.com/tyoro/items/abf9dce0e0020573298c)
