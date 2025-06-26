
Build locally
	go mod tidy
	go get hello_world/cmd
	go build -o hello_world ./cmd
	go run cmd/main.go


Docker And Local Testing
	docker build . -t ghcr.io/nuwas/hello-world-go:latest
	docker run -p 8080:8080 ghcr.io/nuwas/hello-world-go:latest

Local Testing
	http://localhost:8080/yyyy-yyyy/v1/message-y

Deployment

	kubectl apply -f docker-registory-secret.yaml
	kubectl apply -f deployment.yaml  
	kubectl apply -f service.yaml  
	kubectl apply -f virtual-service.yaml


