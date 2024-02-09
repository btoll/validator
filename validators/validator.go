package validators

import "fmt"

type Manifest interface {
	PrintTopLevelManifest()
	PrintSpec()
}

type Validator struct {
	Documents []Document
}

func New(d1, d2 *Document) *Validator {
	return &Validator{
		Documents: []Document{*d1, *d2},
	}
}

func (v *Validator) Validate() {
	d1 := v.Documents[0]
	d2 := v.Documents[1]

	d1.Print()
	fmt.Println("---------------------------------------------------\n")
	d2.Print()

	//	InterleaveDocuments([]Document{d1, d2})

}
