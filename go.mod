module icm.ferth/practica1

go 1.13

require (
	icm.ferth/practica1/v2/racional v0.0.0

	icm.ferth/practica1/v3/racional v0.0.0

	icm.ferth/practica1/v4/racional v0.0.0
)

replace (
	icm.ferth/practica1/v2/racional => ./v2/racional

	icm.ferth/practica1/v3/racional => ./v3/racional

	icm.ferth/practica1/v4/racional => ./v4/racional
)
