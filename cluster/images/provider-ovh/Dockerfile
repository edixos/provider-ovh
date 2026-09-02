FROM --platform=$BUILDPLATFORM golang:1.25.13-alpine3.24 AS dependency-builder

ARG TARGETOS
ARG TARGETARCH
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_REPO
ARG TERRAFORM_PROVIDER_VERSION

RUN apk --no-cache add ca-certificates

WORKDIR /build/terraform
ADD "https://github.com/hashicorp/terraform/archive/refs/tags/v${TERRAFORM_VERSION}.tar.gz" /tmp/terraform.tar.gz
RUN tar -xzf /tmp/terraform.tar.gz --strip-components=1 \
  && CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -mod=mod \
    -o /out/terraform .

WORKDIR /build/terraform-provider-ovh
ADD "${TERRAFORM_PROVIDER_REPO}/archive/refs/tags/v${TERRAFORM_PROVIDER_VERSION}.tar.gz" /tmp/terraform-provider-ovh.tar.gz
RUN tar -xzf /tmp/terraform-provider-ovh.tar.gz --strip-components=1 \
  && go mod edit -require=golang.org/x/mod@v0.40.0 \
  && go mod edit -require=golang.org/x/net@v0.58.0 \
  && go mod edit -require=golang.org/x/text@v0.41.0 \
  && CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -mod=mod \
    -o /out/terraform-provider-ovh .

FROM alpine:3.24.1
RUN apk --no-cache add ca-certificates bash

ARG TARGETOS
ARG TARGETARCH

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532

# Setup Terraform environment

## Provider-dependent configuration
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME
ARG TERRAFORM_NATIVE_PROVIDER_BINARY
ARG TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX

## End of - Provider-dependent configuration

ENV PLUGIN_DIR=/terraform/provider-mirror/registry.terraform.io/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}
ENV TF_CLI_CONFIG_FILE=/terraform/.terraformrc
ENV TF_FORK=0

RUN mkdir -p ${PLUGIN_DIR}

ADD terraformrc.hcl ${TF_CLI_CONFIG_FILE}

COPY --from=dependency-builder /out/terraform /usr/local/bin/terraform
COPY --from=dependency-builder /out/terraform-provider-ovh ${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

RUN chmod +x /usr/local/bin/terraform ${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY} \
  && chown -R ${USER_ID}:${USER_ID} /terraform
# End of - Setup Terraform environment

# Provider controller needs these environment variable at runtime
ENV TERRAFORM_VERSION=${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE=${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}
ENV TERRAFORM_NATIVE_PROVIDER_PATH=${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
