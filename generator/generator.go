package generator

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Generator struct {
	scanDir string
	appDocFile *os.File
	adminDocFile *os.File
	mpDocFile *os.File
}

func (g *Generator) init()  {
	g.appDocFile, _ = getAppDocFile()
	g.adminDocFile, _ = getAdminDocFile()
	g.mpDocFile, _ = getMpDocFile()
}

func (g *Generator) destroy() {
	g.appDocFile.Close()
	g.mpDocFile.Close()
	g.adminDocFile.Close()
}

func getRefreshFile(filePath string) (*os.File, error) {
	if dirOrFileExists(filePath) {
		_ = os.Remove(filePath)
	}
	return os.Create(filePath)
}

func getAppDocFile() (*os.File, error) {
	return getRefreshFile("./app-doc/doc.php")
}

func getAdminDocFile() (*os.File, error) {
	return getRefreshFile("./admin-doc/doc.php")
}

func getMpDocFile() (*os.File, error) {
	return getRefreshFile("./mp-doc/doc.php")
}

func New(scanDir string) *Generator {
	g := &Generator{scanDir: scanDir}
	return g
}

//开始扫描所以项目文件夹
func (g Generator) BeginScanningApiFolder() {
	g.init()
	defer g.destroy()
	for _,project := range initProjects(g.scanDir){
		g.processProjectAnnotates(project)
	}
}

func initProjects(scanDir string) []project {
	var projects []project
	dirEntries, _ := os.ReadDir(scanDir)
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() && isProjectDirName(dirEntry.Name()) {
			projects = append(projects, project{
				name: dirEntry.Name(),
				path: scanDir + dirEntry.Name(),
			})
		}
	}
	return projects
}

func (g *Generator) processProjectAnnotates(project project)  {
	filePaths := recursivelyReadDirFilesByExt(project.getActionDirPath(g.scanDir), "php")
	for _, filePath := range filePaths {
		fmt.Println("查找到文件：" + filePath)
		if !strings.Contains(filePath, "/Doc/Tips.php") {
			apiDocAnnotates := matchFileApiDocAnnotate(filePath, project.name)
			for _, annotate := range apiDocAnnotates {
				if project.isApp() {
					_, _ = g.appDocFile.WriteString(annotate + "\n\n")
				} else if project.isAdmin() {
					_, _ = g.adminDocFile.WriteString(annotate + "\n\n")
				} else if project.isMp() {
					_, _ = g.mpDocFile.WriteString(annotate + "\n\n")
				}
			}
		}
	}
}

//递归的读取文件夹下的所以文件
func recursivelyReadDirFilesByExt(dirName string, limitExt string) []string {
	var filePaths []string
	dirEntries, _ := os.ReadDir(dirName)
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			deepFileInfo := recursivelyReadDirFilesByExt(dirName+dirEntry.Name()+"/", limitExt)
			filePaths = append(filePaths, deepFileInfo...)
		} else {
			if strings.Contains(dirEntry.Name(), "."+limitExt) {
				filePaths = append(filePaths, dirName + dirEntry.Name())
			}
		}
	}
	return filePaths
}

//匹配文件中的所有apidoc注释
func matchFileApiDocAnnotate(filePath string, apiPrefix string) []string {
	dat, _ := os.ReadFile(filePath)
	datStr := string(dat)
	flowsRegexp := regexp.MustCompile(`\s*\* @api[\s\S]*?\*\/`)
	matches := flowsRegexp.FindStringSubmatch(datStr)
	for key, matchStr := range matches {
		flowsRegexp := regexp.MustCompile(`^\s*\*\s*@api\s*{.*`)
		matches0s := flowsRegexp.FindStringSubmatch(matchStr)
		for _, matches0 := range matches0s {
			if !strings.Contains(matches0, apiPrefix) {
				if strings.Contains(matches0, "} /") {
					matches0 = strings.Replace(matches0, " /", " /"+apiPrefix+"/", -1)
				} else if strings.Contains(matches0, "} ") {
					matches0 = strings.Replace(matches0, "} ", "} /"+apiPrefix+"/", -1)
				}
				matchStr = flowsRegexp.ReplaceAllString(matchStr, matches0)
			}
		}
		matches[key] = "    /**" + matchStr
	}
	return matches
}
