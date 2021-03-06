tfapply:
	terraform -chdir=terraform apply -refresh=false
	
tfapplyauto:
	terraform -chdir=terraform apply -refresh=false -auto-approve

deploy:
	go mod tidy
	@echo "Deploying ${LAMBDA} Lambda"
	GOOS=linux GOARCH=amd64 go build -o ./dist/${LAMBDA}/${LAMBDA} ./functions/${LAMBDA}/*.go
	zip -j ./dist/${LAMBDA}_handler.zip ./dist/${LAMBDA}/${LAMBDA}
	aws --profile=ots s3 cp ./dist/${LAMBDA}_handler.zip s3://ppc-lambda-functions/${LAMBDA}_handler.zip
	aws --profile=ots lambda update-function-code --function-name ${LAMBDA}_handler --s3-bucket ppc-lambda-functions --s3-key ${LAMBDA}_handler.zip
	rm -r ./dist/${LAMBDA} ./dist/${LAMBDA}_handler.zip