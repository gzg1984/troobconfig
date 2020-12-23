package troobconfig

/*GetIndexPath will give the index path from the project name*/
func GetIndexPath(project string) string {

	InitGlobalDBManager()
	InitGlobalBasePath()

	indexPath := searchIndex(project)
	if len(indexPath) == 0 {
		return ""
	}
	return globalIndexBasePath + "/" + indexPath
}
