package websupport

import (
	"html/template"
	"text/template/parse"
)

func AddFlush(t *template.Template) {
	addFlushToCommands(t.Tree.Root)
}

func flushCommand(position parse.Pos) *parse.CommandNode {
	return &parse.CommandNode{
		NodeType: parse.NodeCommand,
		Args:     []parse.Node{parse.NewIdentifier("flush").SetTree(nil).SetPos(position)},
	}
}

func addFlushToCommands(node parse.Node) {
	if pipeNode, ok := node.(*parse.PipeNode); ok {
		pipeNode.Cmds = append(pipeNode.Cmds, flushCommand(pipeNode.Position()))
	}

	if templateNode, ok := node.(*parse.TemplateNode); ok {
		if templateNode.Pipe != nil {
			addFlushToCommands(templateNode.Pipe)
		}
	}

	if ln, ok := node.(*parse.ListNode); ok {
		for _, n := range ln.Nodes {
			addFlushToCommands(n)
		}
	}
}
