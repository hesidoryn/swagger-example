swagger:
	@swagger -apiPackage="github.com/hesidoryn/swagger-example/handler" -mainApiFile="router/router.go" -output="swagger"
	@# remove second line (`package main`) and save file
	@sed '2d' ./swagger/docs.go > newdocs.go && mv newdocs.go ./swagger/docs.go
	@# insert `package swagger` in first line of file
	@sed -i '1 i\package swagger' ./swagger/docs.go