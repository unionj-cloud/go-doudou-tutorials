# wordcloud

[![wakatime](https://wakatime.com/badge/user/852bcf22-8a37-460a-a8e2-115833174eba/project/07011958-f15b-450a-8678-b52d11936c9b.svg)](https://wakatime.com/badge/user/852bcf22-8a37-460a-a8e2-115833174eba/project/07011958-f15b-450a-8678-b52d11936c9b)

## Step0
Build docker images for all services
```shell
make docker
```

## Step1
Start all services  
```shell
make up
```

## Step2

Open browser and navigate to minio dashboard `http://localhost:9001/`

- Login first
    - Username: minio
    - Password: minio123
- Create bucket `wordcloud` and set it to `public`
- Create access key `testkey` and access secret `testsecret`

## Step3

Open browser and navigate to `http://localhost:3100/`.

- Username: jackchen
- Password: 1234

## Screenshot

![screenshot1](./screencapture1.png)
![screenshot2](./screencapture2.png)