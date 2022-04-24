tfapply:
	terraform -chdir=terraform apply -refresh=false
	
tfapplyauto:
	terraform -chdir=terraform apply -refresh=false -auto-approve

	