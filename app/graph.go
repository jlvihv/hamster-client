package app

import (
	"context"
	"fmt"
	"hamster-client/module/graph"
	"hamster-client/module/graph/cli"
	param "hamster-client/module/graph/v2"
)

type Graph struct {
	ctx            context.Context
	graphService   graph.Service
	cliService     cli.Service
	graphV2Service param.Service
}

func NewGraphApp(graphService graph.Service, cliService cli.Service, graphV2Service param.Service) Graph {
	return Graph{
		graphService:   graphService,
		cliService:     cliService,
		graphV2Service: graphV2Service,
	}
}

func (g *Graph) WailsInit(ctx context.Context) error {
	g.ctx = ctx
	return nil
}

func (g *Graph) QueryApplyAndParams(applicationId int) (graph.GraphParameterVo, error) {
	return g.graphService.QueryParamByApplyId(applicationId)
}

func (g *Graph) CliLink(applicationId int) (int, error) {
	return g.cliService.CliLink(applicationId)
}

func (g *Graph) GraphStart(applicationId int, deploymentID string) error {
	return g.graphV2Service.GraphStart(applicationId, deploymentID)
}

func (g *Graph) GraphRules(applicationId int) (GraphRulesInfo, error) {
	fmt.Println("##### applicationId: ", applicationId)
	fmt.Println("##### g.graphV2Service: ", g.graphV2Service)
	info, err := g.graphV2Service.GraphRules(applicationId)
	fmt.Println(info)
	return GraphRulesInfo{Info: info}, err
}

type GraphRulesInfo struct {
	Info []param.GraphRule `json:"info"`
}
