package main

// const (

// E = 2.71828182845904523536028747135266249775724709369995957496696763
// Pi = 3.14159265358979323846264338327950288419716939937510582097494459
// Phi = 1.61803398874989484820458683436563811772030917980576286213544862

// // square roots

// squrt2 = 1.41421356237309504880168872420969807856967187537694807317667974
// squrt3 = 1.73205080756887729352744634150587236694280525381038062805580698
// squrtPi = 1.77245385090551602729816748334114518279754945612238712821380779
// squrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038

// // natural logarithms
// ln2 = 0.693147180559945309417232121458176568075500134360255254120680009
// ln10 = 2.30258509299404568401799145468436420760110148862877297603332790
// log2E = 1.44269504088896340735992468100189213742664595415298593413544934
// log10E = 0.43429448190325182765112891891660508229439700580366656611445381

// // Mathemtical constants
// MaxFloat32 = 3.40282346638528859811704183484516925440e+38
// SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45

// MaxFloat64 = 1.797693134862315708145274237317043567981e+308
// SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324

// )

func main() {
	// // RoundToEven rounds x to the nearest even integer, choosing the even integer farther away from zero if x is equidistant between two integers.
	// a := math.RoundToEven(2.5)

	// // Cbrt returns the cube root of x.
	// b := math.Cbrt(27)

	// // Ceil returns the least integer value greater than or equal to x.
	// c := math.Ceil(2.5)

	// // Copysign returns a value with the magnitude of x and the sign of y.
	// d := math.Copysign(2.5, -1)

	// // Cos returns the cosine of the radian argument x.
	// e := math.Cos(0)

	// // Erf returns the error function of x.
	// f := math.Erf(0)

	// // Erfc returns the complementary error function of x.
	// g := math.Erfc(0)

	// // Exp returns e**x, the base-e exponential of x.
	// h := math.Exp(0)

	// // Exp2 returns 2**x, the base-2 exponential of x.
	// i := math.Exp2(0)

	// // Expm1 returns e**x - 1, the base-e exponential of x minus 1.
	// j := math.Expm1(0)

	// // Floor returns the greatest integer value less than or equal to x.
	// l := math.Floor(2.5)

	// // Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow.
	// q := math.Hypot(3, 4)

	// // Ilogb returns the binary exponent of x, which must be finite and non-zero.
	// r := math.Ilogb(8)

	// // Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
	// s := math.Inf(1)

	// // IsInf reports whether f is an infinity, according to sign.
	// t := math.IsInf(2, 1)

	// // IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
	// u := math.IsNaN(2)

	// // J0 returns the order-zero Bessel function of the first kind.
	// v := math.J0(0)

	// // J1 returns the order-one Bessel function of the first kind.
	// w := math.J1(0)

	// // Jn returns the order-n Bessel function of the first kind.
	// x := math.Jn(2, 0)

	// // Ldexp is the inverse of Frexp. It sets x to the (possibly) normalized fraction of f and returns the binary exponent.
	// y := math.Ldexp(2, 3)

	// // Ln2 returns the natural logarithm of 2.
	// aa := math.Ln2

	// // Ln10 returns the natural logarithm of 10.
	// ab := math.Ln10

	// // Log2E returns the base 2 logarithm of E (approximately 0.693147180559945309417).
	// ac := math.Log(2)

	// // Log10E returns the base 10 logarithm of E (approximately 0.434294481903251827651).
	// ad := math.Log10(2)

	// // Log1p returns the natural logarithm of 1 plus its argument.
	// ae := math.Log1p(2)

	// // Logb returns the binary exponent of x.
	// af := math.Log2(2)

	// // Max returns the larger of x or y.
	// ag := math.Max(2, 3)

	// // Min returns the smaller of x or y.
	// ah := math.Min(2, 3)

	// // Mod returns the floating-point remainder of x/y.
	// ai := math.Mod(5, 3)

	// // NaN returns an IEEE 754 ``not-a-number'' value.
	// ak := math.NaN()

	// // Nextafter returns the next representable float32 value after x towards y.
	// al := math.Nextafter(2, 3)

	// // Nextafter32 returns the next representable float32 value after x towards y.
	// am := math.Nextafter32(2, 3)

	// // Pow returns x**y, the base-x exponential of y.
	// ao := math.Pow(2, 3)

	// // Pow10 returns 10**e, the base-10 exponential of e.
	// ap := math.Pow10(2)

	// // Remainder is the IEEE 754 floating-point remainder of x/y.
	// aq := math.Remainder(5, 3)

	// // Round returns the nearest integer, rounding half away from zero.
	// ar := math.Round(2.5)

	// // RoundToEven rounds x to the nearest even integer, choosing the even integer farther away from zero if x is equidistant between two integers.
	// as := math.RoundToEven(2.5)

	// // Signbit reports whether x is negative or negative zero.
	// at := math.Signbit(2)

	// // Sin returns the sine of the radian argument x.
	// au := math.Sin(0)

	// // Sinh returns the hyperbolic sine of x.
	// aw := math.Sinh(0)

	// // Sqrt returns the square root of x.
	// ax := math.Sqrt(4)

	// // Tan returns the tangent of the radian argument x.
	// ay := math.Tan(0)

	// // Tanh returns the hyperbolic tangent of x.
	// az := math.Tanh(0)

	// // Trunc returns the integer value of x.
	// ba := math.Trunc(2.5)

	// // Y0 returns the order-zero Bessel function of the second kind.
	// bb := math.Y0(0)

	// // Y1 returns the order-one Bessel function of the second kind.
	// bc := math.Y1(0)

	// // Yn returns the order-n Bessel function of the second kind.
	// bd := math.Yn(2, 0)

	// // Mathematical constants
	// be := math.Pi

	// // E = 2.71828182845904523536028747135266249775724709369995957496696763
	// bf := math.Phi

	// // square roots
	// bg := math.Sqrt2

	// // squrt3 = 1.73205080756887729352744634150587236694280525381038062805580698
	// bh := math.SqrtE

	// // squrtPi = 1.77245385090551602729816748334114518279754945612238712821380779
	// bi := math.SqrtPi

	// // squrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038
	// bj := math.SqrtPhi

	// // MaxFloat32 = 3.40282346638528859811704183484516925440e+38
	// bk := math.MaxFloat32

	// // SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45
	// bl := math.SmallestNonzeroFloat32

	// // MaxFloat64 = 1.797693134862315708145274237317043567981e+308
	// bm := math.MaxFloat64

	// // SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324
	// bn := math.SmallestNonzeroFloat64

	// fmt.Println(a, b, c, d, e, f, g, h, i, j, l, q, r, s, t, u, v, w, x, y, aa, ab, ac, ad, ae, af, ag, ah, ai, ak, al, am, ao, ap, aq, ar, as, at, au, aw, ax, ay, az, ba, bb, bc, bd, be, bf, bg, bh, bi, bj, bk, bl, bm, bn)

}
