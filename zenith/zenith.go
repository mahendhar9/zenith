package zenith

const version = "1.0.0"

type Zenith struct {
	AppName string
	Debug   bool
	Version string
}

func (z *Zenith) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := z.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}

func (z *Zenith) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := z.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
