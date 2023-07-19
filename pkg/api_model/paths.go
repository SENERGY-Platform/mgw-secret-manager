package api_model

const SecretsPath = "/secrets"
const PathVariantPath = "/path-variant"
const SecretPath = SecretsPath + "/:id"
const LoadPathVariantPath = PathVariantPath + "/load"
const InitPathVariantPath = PathVariantPath + "/init"
const UnLoadPathVariantPath = PathVariantPath + "/unload"

const ConfidentialPath = "/confidential"
const ValueVariantPath = ConfidentialPath + "/value-variant"
