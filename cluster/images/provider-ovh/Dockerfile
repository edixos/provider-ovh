FROM alpine:3.22
RUN apk --no-cache add ca-certificates

ARG TARGETOS
ARG TARGETARCH

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532

## Terraform provider metadata consumed by the controller at startup.
## All resources are reconciled in-process (Terraform Plugin SDK / Framework);
## no Terraform CLI or native provider binary is shipped in this image.
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION

ENV TERRAFORM_VERSION=${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE=${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
