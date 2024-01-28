package church

func Add2(m, n Numeral) Numeral { return Add(m)(n) }
func Mul2(m, n Numeral) Numeral { return Mul(m)(n) }
func Sub2(m, n Numeral) Numeral { return Sub(m)(n) }
func Div2(m, n Numeral) Numeral { return Div(m)(n) }
func Mod2(m, n Numeral) Numeral { return Mod(m)(n) }

func Xor2(p, q Boolean) Boolean { return Xor(p)(q) }
func And2(p, q Boolean) Boolean { return And(p)(q) }
func Or2(p, q Boolean) Boolean  { return Or(p)(q) }

func Pair2(x, y Term) Term { return Pair(x)(y) }

func LessOrEqual2(m, n Numeral) Boolean { return LessOrEqual(m)(n) }
func Less2(m, n Numeral) Boolean        { return Less(m)(n) }
func Equal2(m, n Numeral) Boolean       { return Equal(m)(n) }
