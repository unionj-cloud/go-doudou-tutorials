
docker:
	docker build -t wordcloud_wordcloudbff -f wordcloud-bff/Dockerfile .
	docker build -t wordcloud_wordcloudmaker -f wordcloud-maker/Dockerfile .
	docker build -t wordcloud_wordcloudseg -f wordcloud-seg/Dockerfile .
	docker build -t wordcloud_wordcloudtask -f wordcloud-task/Dockerfile .
	docker build -t wordcloud-ui -f wordcloud-ui/Dockerfile .
	docker build -t wordcloud_usersvc -f wordcloud-user/Dockerfile .

up:
	docker-compose -f docker-compose.yml up -d

down:
	docker-compose -f docker-compose.yml down