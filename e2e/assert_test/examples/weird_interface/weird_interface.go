package weirdinterface

import "github.com/wimspaargaren/final-unit/e2e/assert_test/examples/weird_interface/interfaceimport"

type y struct{}

type x struct{}

type weird interface {
	X(x) y
	x(x) y
}

func WeirdInterface(x weird) weird {
	return nil
}

func WeirdInterfaceImported(x interfaceimport.Weird) interfaceimport.Weird {
	return nil
}

func WeirdInterfaceImportedMap(x map[interfaceimport.Weird]interfaceimport.Weird) interfaceimport.Weird {
	return nil
}

func WeirdInterfaceImportedArray(x []interfaceimport.Weird) interfaceimport.Weird {
	return nil
}

func WeirdInterfaceOnlyMethodNotExportedImported(x interfaceimport.OnlyMethodNotExported) interfaceimport.OnlyMethodNotExported {
	return nil
}

func WeirdInterfaceOnlyImportNotExportedImported(x interfaceimport.OnlyImportNotExported) interfaceimport.OnlyImportNotExported {
	return nil
}

func WeirdInterfaceOnlyOutputNotExportedImported(x interfaceimport.OnlyOutputNotExported) interfaceimport.OnlyOutputNotExported {
	return nil
}
