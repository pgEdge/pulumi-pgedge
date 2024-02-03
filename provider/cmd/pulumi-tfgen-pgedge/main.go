package main

import (
	pgedge "github.com/pgEdge/pulumi-pgedge/provider"
	"github.com/pulumi/pulumi-terraform-bridge/pf/tfgen"
)

func main() {
	tfgen.Main("pgedge", pgedge.Provider())
}
