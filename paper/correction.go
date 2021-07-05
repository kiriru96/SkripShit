package paper

type Correction struct {
	PaperUniqueID string
	Teacher       string
	File          string
	Line          int
	Message       string
}

func (c Correction) ReadCorrection() {

}
