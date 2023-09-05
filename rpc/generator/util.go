package generator

import (
	"strconv"
	"strings"

	"github.com/zhangbao138208/goctls/rpc/parser"
)

func GetGroup(service parser.Service) (data []string) {
	groupNames := map[string]struct{}{}
	for _, rpc := range service.RPC {
		if name := GetGroupName(rpc); name != "" {
			groupNames[name] = struct{}{}
		}
	}

	for k, _ := range groupNames {
		data = append(data, k)
	}

	return data
}

func GetGroupName(rpc *parser.RPC) string {
	if rpc.Comment != nil && len(rpc.Comment.Lines) > 0 {
		for _, v := range rpc.Comment.Lines {
			if strings.Contains(v, "group") {
				groupData := strings.Split(v, ":")
				if len(groupData) == 2 {
					return strings.TrimSpace(groupData[1])
				}
			}
		}
	}
	return ""
}

func HasLock(service parser.Service) bool {
	for _, rpc := range service.RPC {
		if key, _, _ := GetLockKey(rpc); key != "" {
			return true
		}
	}

	return false
}

func GetLockKey(rpc *parser.RPC) (string, string, int) {
	if rpc.Comment != nil && len(rpc.Comment.Lines) > 0 {
		for _, v := range rpc.Comment.Lines {
			if strings.Contains(v, "lock") {
				groupData := strings.Split(v, ":")
				if len(groupData) == 2 {
					return strings.TrimSpace(groupData[1]), "", 0
				}
				if len(groupData) == 3 {
					return strings.TrimSpace(groupData[1]), strings.TrimSpace(groupData[2]), 0
				}
				if len(groupData) == 4 {
					exceed, _ := strconv.Atoi(groupData[3])

					return strings.TrimSpace(groupData[1]), strings.TrimSpace(groupData[2]), exceed
				}
			}
		}
	}
	return "", "", 0
}
