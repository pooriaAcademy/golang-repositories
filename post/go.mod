module github.com/pooriaAcademy/golang-repositories

go 1.15

replace github.com/pooriaAcademy/golang-repositories/user => ../user

replace github.com/pooriaAcademy/golang-repositories/post => ./

require (
	github.com/pooriaAcademy/golang-repositories/post v0.0.0-00010101000000-000000000000
	github.com/pooriaAcademy/golang-repositories/user v0.0.0-00010101000000-000000000000
)
