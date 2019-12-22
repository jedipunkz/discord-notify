# discord-notify

## Description

Discord bot that notify us when our friend join voice chat channel.

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
EOF
```

## Start app

```bash
nohup /some/path/discord-notify
