package Manager

import (
	"fmt"
	"plugin"

	"github.com/shomali11/slacker"
)

type CommandPlugin interface {
    RegisterCommands(bot *slacker.Slacker)
}

type PluginManager struct {
    plugins []CommandPlugin
}

func (p *PluginManager) LoadPlugin(path string) error {
    // Load the plugin binary from the given path
    plugin, err := plugin.Open(path)
    if err != nil {
        return err
    }

    // Look for a CommandPlugin symbol in the plugin binary
    symbol, err := plugin.Lookup("CommandPlugin")
    if err != nil {
        return err
    }

    // Cast the symbol to a CommandPlugin interface and register the commands
    commandPlugin, ok := symbol.(CommandPlugin)
    if !ok {
        return fmt.Errorf("Plugin %s does not implement CommandPlugin", path)
    }
    commandPlugin.RegisterCommands(bot)

    // Add the plugin to the list of loaded plugins
    p.plugins = append(p.plugins, commandPlugin)

    return nil
}
