package utils

import (
	"fmt"

	"github.com/atgse/springboot-operator/pkg/apis/state"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_utils")

func WatchSecondaryResource(c controller.Controller, controllerName string, resourceKind string, objectTypeToWatch runtime.Object, cr runtime.Object) error {
	stateManager := state.GetStateManager()
	stateFieldName := GetStateFieldName(controllerName, resourceKind)

	// Check to see if the watch exists for this resource type already for this controller, if it does, we return so we don't set up another watch
	watchExists, keyExists := stateManager.GetState(stateFieldName).(bool)
	if keyExists || watchExists {
		return nil
	}

	// Set up the actual watch
	err := c.Watch(&source.Kind{Type: objectTypeToWatch}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    cr,
	})

	// Retry on error
	if err != nil {
		log.Error(err, "error creating watch")
		stateManager.SetState(stateFieldName, false)
		return err
	}

	stateManager.SetState(stateFieldName, true)
	log.Info(fmt.Sprintf("Watch created for '%s' resource in '%s'", resourceKind, controllerName))
	return nil
}

func GetStateFieldName(controllerName string, kind string) string {
	return controllerName + "-watch-" + kind
}
