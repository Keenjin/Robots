// +build windows

package main

import (
	"flag"
	"fmt"
	registry "golang.org/x/sys/windows/registry"
	"strings"
)

func QueryUninstallDir(key registry.Key, flag uint32, displayReg, displayName, installReg string) string {
	k, err := registry.OpenKey(key, `Software\Microsoft\Windows\CurrentVersion\Uninstall\`, flag)
	if err != nil {
		//log.Printf("open key failed. err: %v\n", err)
		return ""
	}
	defer k.Close()

	skns, err := k.ReadSubKeyNames(0)
	if err != nil {
		//log.Printf("read subkeys failed. err: %v\n", err)
		return ""
	}
	for _, skn := range skns {
		sk, err := registry.OpenKey(k, skn, flag)
		if err != nil {
			continue
		}

		defer sk.Close()

		name, _, err := sk.GetStringValue(displayReg)
		if err != nil {
			continue
		}

		if strings.EqualFold(name, displayName) {
			// 查询
			instDir, _, err := sk.GetStringValue(installReg)
			if err == nil {
				return instDir
			}
			break
		}
	}

	return ""
}

func main() {
	displayName := flag.String("DisplayName", "", "软件在控制面板-程序-卸载中的名称")
	displayNameRegKey := flag.String("DisplayNameRegKey", "DisplayName", "DisplayName在注册表中实际的key，正常情况下，均为DisplayName")
	installDirRegKey := flag.String("InstallDirRegKey", "InstallLocation", "存储安装目录路径的注册表Key")
	flag.Parse()

	if displayName == nil || displayNameRegKey == nil || installDirRegKey == nil {
		//log.Fatalf("invalid param.\n")
		return
	}
	instDir := QueryUninstallDir(registry.LOCAL_MACHINE, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS, *displayNameRegKey, *displayName, *installDirRegKey)
	if instDir != "" {
		fmt.Print(instDir)
		return
	}
	instDir = QueryUninstallDir(registry.LOCAL_MACHINE, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS|registry.WOW64_32KEY, *displayNameRegKey, *displayName, *installDirRegKey)
	if instDir != "" {
		fmt.Print(instDir)
		return
	}
	instDir = QueryUninstallDir(registry.CURRENT_USER, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS, *displayNameRegKey, *displayName, *installDirRegKey)
	if instDir != "" {
		fmt.Print(instDir)
		return
	}
	instDir = QueryUninstallDir(registry.CURRENT_USER, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS|registry.WOW64_32KEY, *displayNameRegKey, *displayName, *installDirRegKey)
	if instDir != "" {
		fmt.Print(instDir)
		return
	}
	instDir = QueryUninstallDir(registry.CURRENT_USER, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS, *displayNameRegKey, *displayName, *installDirRegKey)
	if instDir != "" {
		fmt.Print(instDir)
		return
	}
	instDir = QueryUninstallDir(registry.CURRENT_USER, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS|registry.WOW64_32KEY, *displayNameRegKey, *displayName, *installDirRegKey)
	if instDir != "" {
		fmt.Print(instDir)
		return
	}
}
