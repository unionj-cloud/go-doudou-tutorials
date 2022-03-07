# wordcloud

[![wakatime](https://wakatime.com/badge/user/852bcf22-8a37-460a-a8e2-115833174eba/project/07011958-f15b-450a-8678-b52d11936c9b.svg)](https://wakatime.com/badge/user/852bcf22-8a37-460a-a8e2-115833174eba/project/07011958-f15b-450a-8678-b52d11936c9b)

## Step0

- `cd wordcloud-bff && docker build -t wordcloud_wordcloudbff .`
- `cd wordcloud-maker && docker build -t wordcloud_wordcloudmaker .`
- `cd wordcloud-seg && docker build -t wordcloud_wordcloudseg .`
- `cd wordcloud-task && docker build -t wordcloud_wordcloudtask .`
- `cd wordcloud-user && docker build -t wordcloud_usersvc .`
- `cd wordcloud-ui && docker build -t wordcloud-ui .`

## Step1

`docker-compose -f docker-compose.yml up -d`

## Step2

Open Chrome or firefox browser and navigate to `http://localhost:3100/`.

- Username: jackchen
- Password: 1234

## Screenshot

![screenshot1](./screencapture1.png)
![screenshot2](./screencapture2.png)