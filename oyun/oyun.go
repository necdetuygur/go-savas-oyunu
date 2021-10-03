package oyun

type ISavasci interface {
	AtesEt()
	HasarAl()
}

type Savasci struct {
	Ad    string
	Can   int
	Mermi int
}

// Dunyali
type Dunyali struct {
	Savasci
	Birlik string
}

func (d *Dunyali) AtesEt() {
	if d.Mermi > 1 {
		d.Mermi--
	}
}

func (d *Dunyali) HasarAl() {
	if d.Can > 0 {
		d.Can -= 10
	}
}

// Marsli
type Marsli struct {
	Savasci
	UzayGemisi string
}

func (d *Marsli) AtesEt() {
	if d.Mermi > 1 {
		d.Mermi--
	}
}

func (d *Marsli) HasarAl() {
	if d.Can > 0 {
		d.Can -= 10
	}
}
