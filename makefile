deploy-check-all: go-mod-tidy-all
	cd generator; make deploy-check
	cd partners; make deploy-check
	cd pkg; make deploy-check

go-mod-tidy-all:
	cd generator; go mod tidy
	cd partners; go mod tidy
	cd pkg; go mod tidy

update-pkg:
	cd pkg; git pull origin main
