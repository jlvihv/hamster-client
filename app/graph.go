package app

import (
	"context"
	"hamster-client/module/graph"
	"hamster-client/module/graph/cli"
)

type Graph struct {
	ctx          context.Context
	graphService graph.Service
	cliService   cli.Service
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
