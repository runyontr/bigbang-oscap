package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/defenseunicorns/bigbang-oscal/pkg/api"
	"github.com/dogenzaka/tsv"
)

type Control struct {
	ControlIdentifier           string `tsv:"Control Identifier"`
	ControlName                 string `tsv:"Control (or Control Enhancement) Name"`
	ControlText                 string `tsv:"Control Text"`
	SecurityControlBaselineHigh string `tsv:"Security Control Baseline"`
	TechnicalOrPolicy           string `tsv:"Technical or Policy"`
	Flux                        string `tsv:"Flux/BigBang"`
	Empty1                      string `tsv:""`
	Isito                       string `tsv:"Istio"`
	Authservice                 string `tsv:"Authservice"`
	Empty2                      string `tsv:""`
	Jaeger                      string `tsv:"Jaeger"`
	Kiali                       string `tsv:"Kiali"`
	OPA                         string `tsv:"OPA Gatekeeper"`
	ClusterAuditor              string `tsv:"Cluster Auditor"`
	Empty3                      string `tsv:""`
	Monitoring                  string `tsv:"Monitoring"`
	Logging                     string `tsv:"Logging"`
	Empty4                      string `tsv:""`
	Twistlock                   string `tsv:"Twistlock"`
	Velero                      string `tsv:"Velero"`
	Empty5                      string `tsv:""`
	Gitlab                      string `tsv:"Gitlab"`
	ArgoCD                      string `tsv:"ArgoCD"`
	Empty6                      string `tsv:""`
	Keycloak                    string `tsv:"Keycloak"`
	empty7                      string `tsv:""`
	CNAP                        string `tsv:"CNAP"`
	Comments                    string `tsv:"Comments"`
	Satisfied                   string `tsv:""`
}

func Parse(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %v: %v\n", filename, err)
		return err
	}
	defer file.Close()

	data := Control{}
	parser, err := tsv.NewParser(file, &data)

	InitComponents()

	for {
		eof, err := parser.Next()
		if eof {
			return nil
		}
		if err != nil {
			fmt.Printf("Error parsing row: %v\n", err)
		}
		// fmt.Println(data)
	}

	return nil
}

func InitComponents() map[string]api.Component {
	out := make(map[string]api.Component)

	out["Keycloak"] = api.NewComponent("Keycloak")
	out["Istio"] = api.NewComponent("Istio")

	b, _ := json.MarshalIndent(out, "", "\t")
	fmt.Println(string(b))
	return out
}
