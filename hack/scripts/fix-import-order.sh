gci write -s standard -s default -s "prefix(github.com/mykytaserdiuk9)" --skip-generated --skip-vendor $(go list ./...)
