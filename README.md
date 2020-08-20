# prog_kubernetes

# to generate code

./k8s.io/code-generator/generate-groups.sh "deepcopy,client,informer,lister" ./k8s.io/cnat-client-go/cnat/pkg/generated k8s.io/cnat-client-go/cnat/pkg/apis cnat:v1alpha1 --output-base "${GOPATH}/src" --go-header-file "k8s.io/cnat-client-go/hack/boilerplate.go.txt"

# after updating with crd controller

./k8s.io/code-generator/generate-groups.sh "deepcopy,client,informer,lister" ./k8s.io/cnat-client-go/cnat/pkg/generated k8s.io/cnat-client-go/cnat/pkg/apis cnat:v1alpha1 --output-base "${GOPATH}/src" --go-header-file "k8s.io/cnat-client-go/hack/boilerplate.go.txt"

go build -o cnat-controller .

./cnat-controller -kubeconfig=$HOME/.kube/config

# create CRD and register and then create CR in a seperate terminal

kubectl create -f cnat-client-go/cnat-crd.yaml
kubectl create -f cnat-client-go/cnat-example.yaml

