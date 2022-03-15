
docker:
	docker build -t wordcloud_wordcloudbff -f wordcloud-bff/Dockerfile .
	docker build -t wordcloud_wordcloudmaker -f wordcloud-maker/Dockerfile .
	docker build -t wordcloud_wordcloudseg -f wordcloud-seg/Dockerfile .
	docker build -t wordcloud_wordcloudtask -f wordcloud-task/Dockerfile .
	docker build -t wordcloud-ui -f wordcloud-ui/Dockerfile .
	docker build -t wordcloud_usersvc -f wordcloud-user/Dockerfile .

up:
	docker-compose -f docker-compose.yml up -d --remove-orphans

down:
	docker-compose -f docker-compose.yml down --remove-orphans

update:
	grep -rli 'github.com/unionj-cloud/go-doudou ' **/go.mod | xargs -I@ sed -E -i '' '/github.com\/unionj-cloud\/go-doudou /s/v[0-9]+(\.[0-9]+)*/$(version)/g' @
	./shellscripts/update.sh