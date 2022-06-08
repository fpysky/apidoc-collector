package generator

type project struct {
	name          string
	path          string
}

func (p project) isApp() bool {
	return isAppProjectDirName(p.name)
}

func (p project) isAdmin() bool {
	return isAdminProjectDirName(p.name)
}

func (p project) isMp() bool {
	return isMpProjectDirName(p.name)
}

func (p project) getActionDirName() string {
	return getProjectActionDirName(p.name)
}

func (p project) getActionDirPath(scanDir string) string {
	//todo: 这里要检查 scanDir 最后一个字符是否是斜杠
	return scanDir + p.name + "/app/" + getProjectActionDirName(p.name) + "/"
}

func isAppProjectDirName(dirName string) bool {
	return isProjectDirNameByApiNames(dirName, getAppApiNames())
}

func isAdminProjectDirName(dirName string) bool {
	return isProjectDirNameByApiNames(dirName, getAdminApiNames())
}

func isMpProjectDirName(dirName string) bool {
	return isProjectDirNameByApiNames(dirName, getMpApiNames())
}

func isProjectDirName(dirName string) bool {
	return isAppProjectDirName(dirName) || isAdminProjectDirName(dirName) || isMpProjectDirName(dirName)
}

func isProjectDirNameByApiNames(dirName string, apiNames []string) bool {
	for _, apiName := range apiNames {
		if dirName == apiName {
			return true
		}
	}
	return false
}

func getProjectActionDirName(projectName string) string {
	var actionDirName string
	switch projectName {
	case "pinhn-certify-api":
		actionDirName = "Action"
		break
	case "focus-on-mp-api":
		actionDirName = "Action"
		break
	case "focus-on-api":
		actionDirName = "Action"
		break
	case "life-message-api":
		actionDirName = "Action"
		break
	case "life-message-mp-api":
		actionDirName = "Action"
		break
	default:
		actionDirName = "Actions"
	}
	return actionDirName
}
