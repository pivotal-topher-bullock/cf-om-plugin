package cmd

type Login struct {
	OpsManager struct {
		Username string `short:"u"  long:"username" description:"Ops Manager admin username"`
		Password string `short:"p"  long:"password" description:"Ops Manager admin password"`
		Target   string `short:"t"  long:"target"   description:"location of the Ops Manager VM"`
	} `group:"Ops Manager Authentication" namespace:"om"`

	Username string `short:"u"  long:"username" description:"CF User to read from ERT Tile Config"`
}

func (cmd *Login) Run() error {
	return nil
}
