package main

import (
	"flag"
	"github.com/appscode/go/signals"
	clientset "go.bytebuilders.dev/crd-learner-template/client/clientset/versioned"
	clusterInformer "go.bytebuilders.dev/crd-learner-template/client/informers/externalversions"
	k8sInformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"time"
)

var (
	k8sconfig string
	masterUrl string

)

func main()  {
	flag.Parse()
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	//get the kubeconfig
	config , err := clientcmd.BuildConfigFromFlags(masterUrl,k8sconfig)
	if err != nil {
		klog.Fatalf("Build config error. the error details : ",err.Error())
	}
	k8sClientset , err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatalf("error in getting k8sClient , error details : ", err.Error())
	}
	clusterClientSet , err := clientset.NewForConfig(config)
	if err != nil {
		klog.Fatalf("error in getting clientset , error details : ", err.Error())
	}
	kubeInformerFactory := k8sInformers.NewSharedInformerFactory(k8sClientset,time.Minute*1)

	clusterInformerFactory := clusterInformer.NewSharedInformerFactory(clusterClientSet,time.Minute*1)

	kubeInformerFactory.Start(stopCh)
	clusterInformerFactory.Start(stopCh)

}

func init() {


	flag.StringVar(&k8sconfig,"kubeconfig","","set the kubeconfig file path via flag command line");
	flag.StringVar(&masterUrl, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}