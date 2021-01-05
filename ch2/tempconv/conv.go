package tempconv

// CtoF converts a Celsius temperature to Fahrenheit
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CtoK converts a Celsius temperature to Fahrenheit
func CtoK(c Celsius) Kelvin { return Kelvin(c) + FreezingK }

// FtoC converts a Fahrenheit temperature to Celsius
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FtoK converts a Fahrenheit temperature to Celsius
func FtoK(f Fahrenheit) Kelvin { return Kelvin(CtoK(FtoC(f))) }

// KtoC converts a Fahrenheit temperature to Celsius
func KtoC(k Kelvin) Celsius { return Celsius(k - FreezingK) }

// KtoF converts a Celsius temperature to Fahrenheit
func KtoF(k Kelvin) Fahrenheit { return Fahrenheit(KtoC(k)) }
