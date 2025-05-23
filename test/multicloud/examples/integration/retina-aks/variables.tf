variable "subscription_id" {
  description = "The subscription ID for the Azure account."
  type        = string
}

variable "tenant_id" {
  description = "The tenant ID for the Azure account."
  type        = string
}

variable "location" {
  description = "The Azure Cloud location where AKS will be deployed to."
  type        = string
  default     = "UK South"
}

variable "resource_group_name" {
  description = "The name of the resource group."
  type        = string
  default     = "example-rg"
}

variable "prefix" {
  description = "A prefix to add to all resources."
  type        = string
  default     = "example-mc"
}

variable "labels" {
  description = "A map of labels to add to all resources."
  type        = map(string)
  default     = {}
}


variable "retina_release_name" {
  description = "The name of the Helm release."
  type        = string
  default     = "retina"
}

variable "retina_repository_url" {
  description = "The URL of the Helm repository."
  type        = string
  default     = "oci://ghcr.io/microsoft/retina/charts"
}

variable "retina_chart_version" {
  description = "The version of the Helm chart to install."
  type        = string
  default     = "v0.0.24"
}

variable "retina_chart_name" {
  description = "The name of the Helm chart to install."
  type        = string
  default     = "retina"
}

variable "retina_values" {
  description = "This corresponds to Helm values.yaml"
  type        = any
}
