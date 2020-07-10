package main

import (
    "fmt"
    "log"

    corev1 "k8s.io/api/core/v1"
    "k8s.io/apimachinery/pkg/util/runtime"

    "k8s.io/client-go/informers"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/cache"
    "k8s.io/client-go/tools/clientcmd"
)

func main() {
    log.Print("Starting the shared informed app")
    config, err := clientcmd.BuildConfigFromFlags("", "/Users/.../.kube/config")
    if err != nil {
        log.Panic(err.Error())
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Panic(err.Error())
    }

    factory := informers.NewSharedInformerFactory(clientset, 0)
    informer := factory.Core().V1().Pods().Informer()
    stopper := make(chan struct{})
    defer close(stopper)
    defer runtime.HandleCrash()
    informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: onAdd,
		DeleteFunc: onDelete,
    })
    go informer.Run(stopper)
    if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
        runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
        return
    }
    <-stopper
}

// when a new pod is deployed the onAdd function would be invoked
// for now just print the event.
func onAdd(obj interface{}) {
    // Cast the obj as Pod
    pod := obj.(*corev1.Pod)
    podName := pod.GetName()
    fmt.Println("Pod started -> ", podName)
}

// when a pod is deleted the onDelete function would be invoked
// for now just print the event 
func onDelete(obj interface{}) {
	pod := obj.(*corev1.Pod)
    podName := pod.GetName()
    fmt.Println("Pod deleted -> ", podName)	
}