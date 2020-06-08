package crd_util

import (
	"context"

	"k8s.io/klog"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func AddFinalizer(m metav1.ObjectMeta, finalizerName string) metav1.ObjectMeta {
	klog.Info("added finilizer")
	for _, name := range m.Finalizers {
		if name == finalizerName {
			return m
		}
	}
	m.Finalizers = append(m.Finalizers, finalizerName)
	return m
}

func HasFinalizer(m metav1.ObjectMeta, finalizerName string) bool {
	for _, name := range m.Finalizers {
		if name == finalizerName {
			return true
		}
	}
	return false
}

func RemoveFinalizer(m metav1.ObjectMeta, finalizerName string) metav1.ObjectMeta {
	var farray []string
	for _, name := range m.Finalizers {
		if name != finalizerName {
			farray = append(farray, name)
		}
	}
	m.Finalizers = farray
	return m
}

func RemoveDeployment(clientset kubernetes.Interface, Namespace string, deploymentName string) error {
	klog.Info("in the remove deployment")
	_, err := clientset.AppsV1().Deployments(Namespace).Get(context.TODO(),
		deploymentName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	err = clientset.AppsV1().Deployments(Namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	return err
}
