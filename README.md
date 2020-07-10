# k8s-sharedinformer
Simple example of the `client-go` SharedInformer that can be used to do all kinds of fun stuff with K8S deployment events

## To run the example

1)  Clone this repo - `git clone https://github.com/venkyvb/k8s-sharedinformer.git`
2)  `cd k8s-sharedinformer`
3)  [Update](https://github.com/venkyvb/k8s-sharedinformer/blob/master/main.go#L18) the path to the KUBECONFIG file (assumes that this app is running outside the cluster)
4)  Do a local K8S cluster set-up either using the Docker Desktop or the much more lightweight and awesome [Kind](https://kind.sigs.k8s.io/)
5)  To update the go-mod dependencies run `go get`
6)  Run the app - `go run main.go`.
7)  In another terminal session run do a simple K8S deployment (the file `mongodb.yaml`) - `kubectl apply -f mongodb.yaml`
8)  You would be able to see the message - `Pod started ->  mongodb-XXXXXXXXX-XXXXX` in the waiting console session created in step (6)
9)  Now run `kubectl delete deployment mongodb` in the 2nd terminal session 
10)  8)  You would be able to see the message - `Pod deleted ->  mongodb-XXXXXXXXX-XXXXX` in the waiting console session created in step (6)

Have fun !!
