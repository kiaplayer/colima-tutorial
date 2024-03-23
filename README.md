Статья: https://habr.com/ru/companies/avito/articles/799689/

```
$ brew install colima
$ brew install kubernetes-cli
$ colima start --kubernetes --network-address
$ docker build -t kubeapp
$ kubectl apply -f deployment.yaml
$ kubectl expose --type=NodePort deployment/kubeapp-example
$ kubectl scale --replicas=5 deployment/kubeapp-example
```