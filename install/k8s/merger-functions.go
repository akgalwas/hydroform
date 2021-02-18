package k8s

import (
	"context"
	"github.com/kyma-incubator/hydroform/install/k8s/maputil"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1Client "k8s.io/client-go/kubernetes/typed/core/v1"
)

type SecretClient interface {
	Update(ctx context.Context, configMap *corev1Client.Secret, opts v1.UpdateOptions) (*Secret, error)
	Get(ctx context.Context, name string, opts v1.GetOptions) (*Secret, error)
}

func NewConfigMapMergeFunc(client corev1Client.ConfigMapInterfacet, newItem *ConfigMap) func() error {
	return func() error {
		old, err := client.Get(context.Background(), newItem.Name, v1.GetOptions{})
		if err == nil {
			maputil.MergeStringMaps(old.Data, newItem.Data)
		}

		return err
	}
}

func NewConfigMapUpdateFunc(client corev1Client.ConfigMapInterface, newItem *ConfigMap) func() error {
	return func() error {
		_, err := client.Update(context.Background(), newItem, v1.UpdateOptions{})
		return err
	}
}

func NewSecretMergeFunc(client corev1Client.SecretInterface, newItem *corev1Client.Secret) func() error {
	return func() error {
		old, err := client.Get(context.Background(), newItem.Name, v1.GetOptions{})
		if err == nil {
			maputil.MergeByteMaps(old.Data, newItem.Data)
		}

		return err
	}
}

func NewSecretUpdateFunc(client corev1Client.SecretInterface, newItem *corev1Client.Secret) func() error {
	return func() error {
		_, err := client.Update(context.Background(), newItem, v1.UpdateOptions{})
		return err
	}
}
