package config

import (
	"github.com/BurntSushi/toml"
	"github.com/ayn2op/tview/keybind"
)

type Keybind struct {
	keybind.Keybind
}

var _ toml.Unmarshaler = (*Keybind)(nil)

func (k *Keybind) UnmarshalTOML(value any) error {
	switch value := value.(type) {
	case string:
		k.SetKeys(value)
	case []any:
		keys := make([]string, 0, len(value))
		for _, key := range value {
			if key, ok := key.(string); ok {
				keys = append(keys, key)
			}
		}
		k.SetKeys(keys...)
	}
	// Keep displayed help key aligned with configured key(s).
	if keys := k.Keys(); len(keys) > 0 {
		k.SetHelp(keys[0], k.Help().Desc)
	}
	return nil
}

func newKeybind(key string, desc string, aliases ...string) Keybind {
	keys := append([]string{key}, aliases...)
	return Keybind{
		Keybind: keybind.NewKeybind(
			keybind.WithKeys(keys...),
			keybind.WithHelp(key, desc),
		),
	}
}

type NavigationKeybinds struct {
	Up     Keybind `toml:"up"`
	Down   Keybind `toml:"down"`
	Top    Keybind `toml:"top"`
	Bottom Keybind `toml:"bottom"`
}

type ScrollKeybinds struct {
	ScrollUp     Keybind `toml:"scroll_up"`
	ScrollDown   Keybind `toml:"scroll_down"`
	ScrollTop    Keybind `toml:"scroll_top"`
	ScrollBottom Keybind `toml:"scroll_bottom"`
}

type SelectionKeybinds struct {
	SelectUp     Keybind `toml:"select_up"`
	SelectDown   Keybind `toml:"select_down"`
	SelectTop    Keybind `toml:"select_top"`
	SelectBottom Keybind `toml:"select_bottom"`
}

type PickerKeybinds struct {
	NavigationKeybinds
	Select Keybind `toml:"select"`
	Cancel Keybind `toml:"cancel"`
}

type GuildsTreeKeybinds struct {
	NavigationKeybinds
	SelectCurrent Keybind `toml:"select_current"`
	YankID        Keybind `toml:"yank_id"`

	CollapseAll        Keybind `toml:"collapse_all"`
	CollapseParentNode Keybind `toml:"collapse_parent_node"`
	MoveToParentNode   Keybind `toml:"move_to_parent_node"`
}

type MessagesListKeybinds struct {
	SelectionKeybinds
	ScrollKeybinds

	SelectReply  Keybind `toml:"select_reply"`
	Reply        Keybind `toml:"reply"`
	ReplyMention Keybind `toml:"reply_mention"`

	Cancel        Keybind `toml:"cancel"`
	Edit          Keybind `toml:"edit"`
	Delete        Keybind `toml:"delete"`
	DeleteConfirm Keybind `toml:"delete_confirm"`
	Open          Keybind `toml:"open"`

	YankContent Keybind `toml:"yank_content"`
	YankURL     Keybind `toml:"yank_url"`
	YankID      Keybind `toml:"yank_id"`
}

type MessageInputKeybinds struct {
	Paste       Keybind `toml:"paste"`
	Send        Keybind `toml:"send"`
	Cancel      Keybind `toml:"cancel"`
	TabComplete Keybind `toml:"tab_complete"`
	Undo        Keybind `toml:"undo"`

	OpenEditor     Keybind `toml:"open_editor"`
	OpenFilePicker Keybind `toml:"open_file_picker"`
}

type MentionsListKeybinds struct {
	NavigationKeybinds
}

type Keybinds struct {
	ToggleGuildsTree     Keybind `toml:"toggle_guilds_tree"`
	ToggleChannelsPicker Keybind `toml:"toggle_channels_picker"`
	ToggleHelp           Keybind `toml:"toggle_help"`
	Suspend              Keybind `toml:"suspend"`

	FocusGuildsTree   Keybind `toml:"focus_guilds_tree"`
	FocusMessagesList Keybind `toml:"focus_messages_list"`
	FocusMessageInput Keybind `toml:"focus_message_input"`

	FocusPrevious Keybind `toml:"focus_previous"`
	FocusNext     Keybind `toml:"focus_next"`

	Picker       PickerKeybinds       `toml:"picker"`
	GuildsTree   GuildsTreeKeybinds   `toml:"guilds_tree"`
	MessagesList MessagesListKeybinds `toml:"messages_list"`
	MessageInput MessageInputKeybinds `toml:"message_input"`
	MentionsList MentionsListKeybinds `toml:"mentions_list"`

	Logout Keybind `toml:"logout"`
	Quit   Keybind `toml:"quit"`
}

func defaultPickerKeybinds() PickerKeybinds {
	return PickerKeybinds{
		NavigationKeybinds: NavigationKeybinds{
			Up:     newKeybind("up", "up", "ctrl+p"),
			Down:   newKeybind("down", "down", "ctrl+n"),
			Top:    newKeybind("home", "top"),
			Bottom: newKeybind("end", "bottom"),
		},
		Cancel: newKeybind("esc", "cancel"),
		Select: newKeybind("enter", "sel"),
	}
}

func defaultNavigationKeybinds() NavigationKeybinds {
	return NavigationKeybinds{
		Up:     newKeybind("up", "up", "k"),
		Down:   newKeybind("down", "down", "j"),
		Top:    newKeybind("home", "top", "g"),
		Bottom: newKeybind("end", "bottom", "G"),
	}
}

func defaultGuildsTreeKeybinds() GuildsTreeKeybinds {
	return GuildsTreeKeybinds{
		NavigationKeybinds: defaultNavigationKeybinds(),
		SelectCurrent:      newKeybind("enter", "select"),
		YankID:             newKeybind("i", "copy id"),

		CollapseAll:        newKeybind("_", "collapse all"),
		CollapseParentNode: newKeybind("-", "collapse parent"),
		MoveToParentNode:   newKeybind("p", "parent"),
	}
}

func defaultMessagesListKeybinds() MessagesListKeybinds {
	return MessagesListKeybinds{
		SelectionKeybinds: SelectionKeybinds{
			SelectUp:     newKeybind("k", "up"),
			SelectDown:   newKeybind("j", "down"),
			SelectTop:    newKeybind("home", "top", "g"),
			SelectBottom: newKeybind("end", "bottom", "G"),
		},
		ScrollKeybinds: ScrollKeybinds{
			ScrollUp:     newKeybind("up", "scroll up", "K"),
			ScrollDown:   newKeybind("down", "scroll down", "J"),
			ScrollTop:    newKeybind("home", "scroll top"),
			ScrollBottom: newKeybind("end", "scroll bottom"),
		},
		SelectReply:  newKeybind("s", "sel reply"),
		Reply:        newKeybind("R", "reply"),
		ReplyMention: newKeybind("r", "@reply"),
		Cancel:       newKeybind("esc", "cancel"),
		Edit:         newKeybind("e", "edit"),
		Delete:       newKeybind("D", "force delete"),
		DeleteConfirm: newKeybind(
			"d",
			"delete",
		),
		Open:        newKeybind("o", "open"),
		YankContent: newKeybind("y", "copy text"),
		YankURL:     newKeybind("u", "copy url"),
		YankID:      newKeybind("i", "copy id"),
	}
}

func defaultMessageInputKeybinds() MessageInputKeybinds {
	return MessageInputKeybinds{
		Paste:          newKeybind("ctrl+v", "paste"),
		Send:           newKeybind("enter", "send"),
		Cancel:         newKeybind("esc", "cancel"),
		TabComplete:    newKeybind("ctrl+space", "complete"),
		Undo:           newKeybind("ctrl+u", "undo"),
		OpenEditor:     newKeybind("ctrl+e", "editor"),
		OpenFilePicker: newKeybind("ctrl+\\", "attach"),
	}
}

func defaultMentionsListKeybinds() MentionsListKeybinds {
	return MentionsListKeybinds{
		NavigationKeybinds: NavigationKeybinds{
			Up:     newKeybind("up", "up", "ctrl+p"),
			Down:   newKeybind("down", "down", "ctrl+n"),
			Top:    newKeybind("home", "top"),
			Bottom: newKeybind("end", "bottom"),
		},
	}
}

func defaultKeybinds() Keybinds {
	return Keybinds{
		ToggleGuildsTree:     newKeybind("f2", "toggle guilds", "ctrl+b"),
		ToggleChannelsPicker: newKeybind("f3", "channels picker", "ctrl+k"),
		ToggleHelp:           newKeybind("ctrl+.", "help"),
		Suspend:              newKeybind("ctrl+z", "suspend"),

		FocusGuildsTree:   newKeybind("f4", "guilds", "ctrl+g"),
		FocusMessagesList: newKeybind("f5", "messages", "ctrl+t"),
		FocusMessageInput: newKeybind("f6", "input", "ctrl+i"),

		FocusPrevious: newKeybind("shift+tab", "focus prev", "ctrl+h"),
		FocusNext:     newKeybind("tab", "focus next", "ctrl+l"),

		Logout: newKeybind("ctrl+d", "logout"),
		Quit:   newKeybind("ctrl+c", "quit"),

		Picker:       defaultPickerKeybinds(),
		GuildsTree:   defaultGuildsTreeKeybinds(),
		MessagesList: defaultMessagesListKeybinds(),
		MessageInput: defaultMessageInputKeybinds(),
		MentionsList: defaultMentionsListKeybinds(),
	}
}
