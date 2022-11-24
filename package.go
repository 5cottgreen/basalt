package basalt

type Packages []Package

type Package struct {
	Name      string
	Epoch     int
	Version   string
	Release   string
	Arch      string
	Disttag   string
	Buildtime int
	Source    string
}
