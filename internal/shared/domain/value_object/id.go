package value_object

import (
	"github.com/bwmarrin/snowflake"
	"strconv"
)

type ID int64

type IDGenerator interface {
	ID() ID
}

type idGenerator struct {
	node *snowflake.Node
}

func NewIDGenerator(ud int64) (IDGenerator, error) {
	node, err := snowflake.NewNode(ud)
	if err != nil {
		return nil, err
	}
	return &idGenerator{node: node}, nil
}

func (g *idGenerator) ID() ID {
	return ID(g.node.Generate().Int64())
}

func ToID(id string) (ID, error) {
	i, _ := strconv.ParseInt(id, 10, 64)
	return ID(i), nil
}
