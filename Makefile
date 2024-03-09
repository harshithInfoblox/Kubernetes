up:
	docker-compose build
	docker-compose down
	docker-compose up -d
remove-containers:
	kubectl delete deployment myapp-deployment
	kubectl delete deployment db-deployment
	kubectl delete service myapp-service
	kubectl delete service db-service

kuberemove:
	kubectl delete -f myapp-deployment.yaml
	kubectl delete -f db-deployment.yaml

build-image:
	docker build -t hp2411/myhnew-app:latest .

push-image:
	docker push hp2411/myhnew-app:latest

kubeapply:
	kubectl apply -f db-secret.yaml
	kubectl apply -f myapp-deployment.yaml
	kubectl apply -f db-deployment.yaml
get:
	kubectl get pods
	kubectl get deployments
	kubectl get services
	kubectl get secret db-secret
forward:
	kubectl port-forward service/myapp-service 8080:8080

# describe:
# 	kubectl describe pod myapp-deployment-59c944868b-8gmgh
# 	kubectl describe deployment myapp-deployment

# logs:
# 	kubectl logs myapp-deployment-59c944868b-8gmgh

# deletepod:
# 	kubectl delete pod myapp-deployment-59c944868b-8gmgh