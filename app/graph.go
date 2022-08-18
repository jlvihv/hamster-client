package app

import (
	"context"
	"hamster-client/module/graph"
	"hamster-client/module/graph/cli"
	graphV2 "hamster-client/module/graph/v2"
)

type Graph struct {
	ctx            context.Context
	graphService   graph.Service
	cliService     cli.Service
	graphV2Service graphV2.Service
}

func NewGraphApp(graphService graph.Service, cliService cli.Service) Graph {
	return Graph{
		graphService: graphService,
		cliService:   cliService,
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
	info, err := g.graphV2Service.GraphRules(applicationId)
	return GraphRulesInfo{Info: info}, err
}

type GraphRulesInfo struct {
	Info []graphV2.GraphRule `json:"info"`
}
