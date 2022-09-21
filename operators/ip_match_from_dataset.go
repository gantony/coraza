// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package operators

import (
	"fmt"
	"strings"

	"github.com/corazawaf/coraza/v3"
)

type ipMatchFromDataset struct {
	matcher *ipMatch
}

var _ coraza.RuleOperator = (*ipMatchFromDataset)(nil)

func (o *ipMatchFromDataset) Init(options coraza.RuleOperatorOptions) error {
	data := options.Arguments
	dataset, ok := options.Datasets[data]
	if !ok || len(dataset) == 0 {
		return fmt.Errorf("dataset %q not found", data)
	}

	datasetParsed := strings.Join(dataset, ",")

	o.matcher = &ipMatch{}
	opts := coraza.RuleOperatorOptions{
		Arguments: datasetParsed,
	}
	return o.matcher.Init(opts)
}

func (o *ipMatchFromDataset) Evaluate(tx *coraza.Transaction, value string) bool {
	return o.matcher.Evaluate(tx, value)
}

var _ coraza.RuleOperator = (*ipMatchFromDataset)(nil)