package chat

import (
	"testing"

	"github.com/ayn2op/tview"
	"github.com/ayn2op/tview/list"
)

func TestFocusedPaneRecognizesEmbeddedFocus(t *testing.T) {
	tests := []struct {
		name  string
		focus func(*Model)
		want  focusedPane
	}{
		{
			name: "guilds tree",
			focus: func(m *Model) {
				m.guildsTree.Focus(nil)
			},
			want: focusedPaneGuildsTree,
		},
		{
			name: "messages list wrapper",
			focus: func(m *Model) {
				m.messagesList.Focus(nil)
			},
			want: focusedPaneMessagesList,
		},
		{
			name: "message input text area",
			focus: func(m *Model) {
				m.messageInput.Focus(nil)
			},
			want: focusedPaneMessageInput,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := &Model{
				guildsTree:   &guildsTree{TreeView: tview.NewTreeView()},
				messagesList: &messagesList{Model: list.NewModel()},
				messageInput: &messageInput{TextArea: tview.NewTextArea()},
			}
			test.focus(m)

			if got := m.focusedPane(); got != test.want {
				t.Fatalf("focusedPane() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestGuildsTreeWrapMove(t *testing.T) {
	root := tview.NewTreeNode("")
	first := tview.NewTreeNode("first")
	folder := tview.NewTreeNode("folder").SetExpanded(true)
	child := tview.NewTreeNode("child")
	last := tview.NewTreeNode("last")
	collapsed := tview.NewTreeNode("collapsed").SetExpanded(false)
	hidden := tview.NewTreeNode("hidden")

	folder.AddChild(child)
	collapsed.AddChild(hidden)
	root.AddChild(first).AddChild(folder).AddChild(collapsed).AddChild(last)

	gt := &guildsTree{TreeView: tview.NewTreeView()}
	gt.SetRoot(root).SetTopLevel(1)

	gt.SetCurrentNode(first)
	if !gt.wrapMove(true) {
		t.Fatal("wrapMove(up) did not wrap from first node")
	}
	if got := gt.GetCurrentNode(); got != last {
		t.Fatal("current node after up wrap was not the last node")
	}

	if !gt.wrapMove(false) {
		t.Fatal("wrapMove(down) did not wrap from last node")
	}
	if got := gt.GetCurrentNode(); got != first {
		t.Fatal("current node after down wrap was not the first node")
	}

	gt.SetCurrentNode(folder)
	if gt.wrapMove(true) {
		t.Fatal("wrapMove(up) wrapped from a middle node")
	}
	if gt.wrapMove(false) {
		t.Fatal("wrapMove(down) wrapped from a middle node")
	}
}
