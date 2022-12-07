run:
	docker build -t go-app .
# run:
	docker run --name test1 -p 8080:8080 -dt go-app
	@echo ""
	@echo "GO to: http://127.0.0.1:8080"
stop:
	docker stop test1
	
remove:
	@docker rm $(docker ps -a -q)