package basalt

// Fetches packages from db and prints them
func SimpleFetch(names []string) (pkgs Packages) {
	for _, name := range names {
		resp := getPackages(name)
		pkgs = append(pkgs, resp.Packages...)
	}

	return pkgs
}

// Fetches packages from db, compares and prints them
func ComparativeFetch(names []string) (pkgs Packages) {
	resp1 := getPackages(names[0])
	resp2 := getPackages(names[1])

	for _, pkg1 := range resp1.Packages {
		exst := false
		for _, pkg2 := range resp2.Packages {
			if pkg1 == pkg2 {
				exst = true
				break
			}
		}
		if exst {
			pkgs = append(pkgs, pkg1)
		}
	}

	return pkgs
}

// Fetches packages from db, compares in reverse way and prints them
func ReverseComparativeFetch(names []string) (pkgs Packages) {
	resp1 := getPackages(names[0])
	resp2 := getPackages(names[1])

	for _, pkg2 := range resp2.Packages {
		exst := false
		for _, pkg1 := range resp1.Packages {
			if pkg2 == pkg1 {
				exst = true
				break
			}
		}
		if exst {
			pkgs = append(pkgs, pkg2)
		}
	}

	return pkgs
}

// Fetches packages from db, compares by version and prints them
func VersionComparativeFetch(names []string) (pkgs Packages) {
	resp1 := getPackages(names[0])
	resp2 := getPackages(names[1])

	for _, pkg1 := range resp1.Packages {
		for _, pkg2 := range resp2.Packages {
			if pkg1.Name == pkg2.Name {
				if pkg1.Version > pkg2.Version {
					pkgs = append(pkgs, pkg1)
				}
				break
			}
		}
	}

	return pkgs
}
