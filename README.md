# discord-notify

[![Build Status](https://travis-ci.org/jedipunkz/discord-notify.svg?branch=master)](https://travis-ci.org/jedipunkz/discord-notify)

:avocado: Discord bot :avocado: that notify when friend join voice chat channel.

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

| Parameter         | Meaning           |
|-------------------|-------------------|
| token             | Token of your bot |
| notify_channel_id | Discord Channel ID which you want to notify. You can get ID on Developer Mode of discord app. |
| me                | Your discord name without (#) numbers. This is required for notifing when friend (*exclude you*) join channel |

## Start app

```bash
nohup /some/path/to/discord-notify
```

## Docker

### docker build

Copy .discord-notify.yaml file from $HOME dir to current dir, and build docker container image.

```bash
cp $HOME/.discord-notify.yaml .
docker build . -t discord-notify
```

### docker run

Run docker container on daemon mode.

```bash
docker run -d slack-ansible
```

## Reference

- [Discord で Voice チャンネルへの入室を検知する bot (ついでに Google Home で通知](https://qiita.com/tyoro/items/abf9dce0e0020573298c)
