package routecm

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"

	configv1 "github.com/openshift/api/config/v1"
	hyperv1 "github.com/openshift/hypershift/api/v1beta1"
	"github.com/openshift/hypershift/support/config"
	"github.com/openshift/hypershift/support/globalconfig"
)

type OpenShiftRouteControllerManagerParams struct {
	OpenShiftControllerManagerImage string
	APIServer                       *configv1.APIServerSpec

	DeploymentConfig config.DeploymentConfig
	config.OwnerRef
}

func NewOpenShiftRouteControllerManagerParams(hcp *hyperv1.HostedControlPlane, observedConfig *globalconfig.ObservedConfig, images map[string]string, setDefaultSecurityContext bool) *OpenShiftRouteControllerManagerParams {
	params := &OpenShiftRouteControllerManagerParams{
		OpenShiftControllerManagerImage: images["route-controller-manager"],
	}
	if hcp.Spec.Configuration != nil {
		params.APIServer = hcp.Spec.Configuration.APIServer
	}

	params.DeploymentConfig = config.DeploymentConfig{
		Scheduling: config.Scheduling{
			PriorityClass: config.DefaultPriorityClass,
		},
		Resources: map[string]corev1.ResourceRequirements{
			routeOCMContainerMain().Name: {
				Requests: corev1.ResourceList{
					corev1.ResourceMemory: resource.MustParse("100Mi"),
					corev1.ResourceCPU:    resource.MustParse("100m"),
				},
			},
		},
	}
	params.DeploymentConfig.SetRestartAnnotation(hcp.ObjectMeta)
	params.DeploymentConfig.SetDefaults(hcp, openShiftRouteControllerManagerLabels(), nil)
	params.DeploymentConfig.SetDefaultSecurityContext = setDefaultSecurityContext

	params.OwnerRef = config.OwnerRefFrom(hcp)

	params.DeploymentConfig.LivenessProbes = config.LivenessProbes{
		routeOCMContainerMain().Name: {
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path:   "/healthz",
					Port:   intstr.FromInt(int(servingPort)),
					Scheme: corev1.URISchemeHTTPS,
				},
			},
			InitialDelaySeconds: 60,
			PeriodSeconds:       60,
			SuccessThreshold:    1,
			FailureThreshold:    5,
			TimeoutSeconds:      5,
		},
	}
	params.DeploymentConfig.ReadinessProbes = config.ReadinessProbes{
		routeOCMContainerMain().Name: {
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path:   "/healthz/ready",
					Port:   intstr.FromInt(int(servingPort)),
					Scheme: corev1.URISchemeHTTPS,
				},
			},
			InitialDelaySeconds: 15,
			PeriodSeconds:       60,
			SuccessThreshold:    1,
			FailureThreshold:    3,
			TimeoutSeconds:      5,
		},
	}

	return params
}
