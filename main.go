package main

import (
	"flag"
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)



func main() {
	res:=flag.String("res","pods","Resource name")
	flag.Parse()
	
	configFlag := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	matchedVersionFlag := cmdutil.NewMatchVersionFlags(configFlag)
	mapper, err := cmdutil.NewFactory(matchedVersionFlag).ToRESTMapper()
	if err != nil {
		log.Fatal(err)
	}

	gvr, err := mapper.ResourceFor(schema.GroupVersionResource{
		Resource: *res,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Complete gvr: %v", gvr)
}
