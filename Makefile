VERSION ?= stable

generate:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
			-i https://repo.jellyfin.org/files/openapi/jellyfin-openapi-${VERSION}.json \
			-g go \
			--git-repo-id go-jellyfin-api-client --git-user-id shamelin \
			-c /local/config.yaml \
			-o /local
