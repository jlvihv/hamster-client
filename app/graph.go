package app

import (
	"context"
	"hamster-client/module/graph"
)

type Graph struct {
	ctx          context.Context
	graphService graph.Service
}

func NewGraphApp(graphService graph.Service) Graph {
	return Graph{
		graphService: graphService,
	}
}

func (g *Graph) WailsInit(ctx context.Context) error {
	g.ctx = ctx
	return nil
}

func (g *Graph) QueryApplyAndParams(applicationId int) (graph.GraphParameter, error) {
	return g.graphService.QueryParamByApplyId(applicationId)
}
