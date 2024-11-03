
Docker And Local Testing
	docker build . -t ghcr.io/nuwas/hello-world-go:latest
	docker run -p 8080:8080 ghcr.io/nuwas/hello-world-go:latest

Local Testing
	http://localhost:8080/yyyy-yyyy/v1/message-y


Istio

    istioctl install --set profile=default
    kubectl label namespace default istio-injection=enabled

Certificate
    export DOMAIN_NAME=nuwas.nn
    openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=$DOMAIN_NAME Inc./CN=$DOMAIN_NAME' -keyout $DOMAIN_NAME.key -out $DOMAIN_NAME.crt
    openssl req -out tt.$DOMAIN_NAME.csr -newkey rsa:2048 -nodes -keyout tt.$DOMAIN_NAME.key -subj "/CN=tt.$DOMAIN_NAME/O=tt from $DOMAIN_NAME"
    openssl x509 -req -days 365 -CA $DOMAIN_NAME.crt -CAkey $DOMAIN_NAME.key -set_serial 0 -in tt.$DOMAIN_NAME.csr -out tt.$DOMAIN_NAME.crt

    kubectl create secret tls global-istio-tls --key tt.$DOMAIN_NAME.key --cert tt.$DOMAIN_NAME.crt -n istio-system

Gateway-Crd
    kubectl apply -f gateway_crd.yaml

Deployment

	kubectl apply -f docker-registory-secret.yaml
	kubectl apply -f deployment.yaml  
	kubectl apply -f service.yaml  
	kubectl apply -f virtual-service.yaml


