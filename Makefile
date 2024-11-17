VERSION ?= stable

generate:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
			-i https://repo.jellyfin.org/files/openapi/jellyfin-openapi-${VERSION}.json \
			-g go \
			--package-name jellyfin-api-client \
			--git-repo-id jellyfin-api-client-go --git-user-id shamelin \
			-o /local
