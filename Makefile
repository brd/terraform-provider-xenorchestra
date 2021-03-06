.PHONY: import testacc dist


ifdef TEST
    TEST := ./... -run $(TEST)
else
    TEST := ./...
endif

ifdef TF_LOG
    TF_LOG := TF_LOG=$(TF_LOG)
endif

build:
	go build -o terraform-provider-xenorchestra

clean:
	rm dist/*

dist:
	./scripts/dist.sh

plan: build
	terraform init
	terraform plan

apply:
	terraform apply

testacc:
	TF_ACC=1 $(TF_LOG) go test $(TEST) -v -count 1 -timeout 40m
