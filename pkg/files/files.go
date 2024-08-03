package files

import "os"

type FileInfo struct {
	Name  string
	IsDir bool
}

func GetFilesStruct(files []os.FileInfo) []FileInfo {
	var fileInfos []FileInfo

  fileInfos = append(fileInfos, FileInfo{
    Name:  "..",
    IsDir: true,
  })

	for _, entry := range files {
		fileInfos = append(fileInfos, FileInfo{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
		})
	}


	return fileInfos
}

func DeleteFile(dir string, file string) {
  err := os.Remove(dir + "/" + file)
  if err != nil {
    panic(err)
  }
}

func CreateFile(dir string, file string) {
  if file[len(file)-1] == '/' {
    err := os.Mkdir(dir + "/" + file, 0755)
    if err != nil {
      panic(err)
    }
  } else {
    _, err := os.Create(dir + "/" + file)
    if err != nil {
      panic(err)
    }
  }
}
