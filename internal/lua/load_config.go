package lua

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

type Config struct {
	ExportVars map[string]string
	Aliases    map[string]string
}

func (c *Config) LoadConfig(filePath string) error {
	L := lua.NewState()
	defer L.Close()

	if err := L.DoFile(filePath); err != nil {
		return err
	}

	luaTable := L.Get(-1)
	if luaTable == nil {
		return fmt.Errorf("Lua file %s did not return a valid table", filePath)
	}

	// Ensure luaTable is a table
	luaTable, ok := luaTable.(*lua.LTable)
	if !ok {
		return fmt.Errorf("Lua file %s did not return a valid table", filePath)
	}

	// Extract 'export' subtable
	exportTable := luaTable.(*lua.LTable).RawGetString("export")
	c.ExportVars = make(map[string]string)

	exportTable.(*lua.LTable).ForEach(func(key lua.LValue, value lua.LValue) {
		if keyStr, ok := key.(lua.LString); ok {
			c.ExportVars[string(keyStr)] = string(value.(lua.LString))
		}
	})

	// Extract 'aliases' subtable
	aliasesTable := luaTable.(*lua.LTable).RawGetString("aliases")
	c.Aliases = make(map[string]string)

	aliasesTable.(*lua.LTable).ForEach(func(key lua.LValue, value lua.LValue) {
		if keyStr, ok := key.(lua.LString); ok {
			c.Aliases[string(keyStr)] = string(value.(lua.LString))
		}
	})

	return nil
}
