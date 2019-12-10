package controller

import (
	"github.com/atgse/springboot-operator/pkg/controller/springbootapplication"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, springbootapplication.Add)
}
