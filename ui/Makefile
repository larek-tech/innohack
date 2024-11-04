.PHONY: docker-build-amd64
docker-build-amd64:
	docker build \
		--tag ghcr.io/larek-tech/innohack/ui:dev-79cc33a \
		--platform linux/amd64 .

.PHONY: docker-push-amd64
docker-push-amd64:
	docker push ghcr.io/larek-tech/innohack/ui:dev-79cc33a


.PHONY: docker-publish-amd64
docker-publish-amd64: docker-build-amd64 docker-push-amd64