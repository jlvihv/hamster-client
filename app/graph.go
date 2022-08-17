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

func (g *Graph) GraphConnect() error {
	return g.graphV2Service.GraphConnect(34003)
}

func (g *Graph) GraphStart(deploymentID string) error {
	return g.graphV2Service.GraphStart(34003, deploymentID)
}

func (g *Graph) GraphRules() (GraphRulesInfo, error) {
	info, err := g.graphV2Service.GraphRules(34003)
	return GraphRulesInfo{Info: info}, err
}

type GraphRulesInfo struct {
	Info []map[string]interface{} `json:"info"`
}
