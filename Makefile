tfapply:
	terraform -chdir=terraform apply -refresh=false
	
tfapplyauto:
	terraform -chdir=terraform apply -refresh=false -auto-approve

deploy:
	@echo ${LAMBDA}
	GOOS=linux GOARCH=amd64 go build -o ./dist/${LAMBDA}/${LAMBDA} ./functions/${LAMBDA}/*.go
	zip -j ./dist/${LAMBDA}_handler.zip ./dist/${LAMBDA}/${LAMBDA}
	aws --profile=onetwentyseven s3 cp ./dist/${LAMBDA}_handler.zip s3://ppc-lambda-functions/${LAMBDA}_handler.zip
	aws --profile=onetwentyseven lambda update-function-code --function-name ${LAMBDA}_handler --s3-bucket ppc-lambda-functions --s3-key ${LAMBDA}_handler.zip
	rm -r ./dist