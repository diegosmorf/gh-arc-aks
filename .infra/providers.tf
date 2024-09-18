terraform {

  # Register common providers
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "<=3.116.0"
    }

    helm = {
      source  = "hashicorp/helm"
      version = "2.15.0"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.32.0"
    }
  }

  # Persist state in a storage account
  #   backend "azurerm" {
  #   }
}

# Configure the Azure Provider
provider "azurerm" {
  # skip_provider_registration = true
  subscription_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  features {}
}

# Data

# Provides client_id, tenant_id, subscription_id and object_id variables
data "azurerm_client_config" "current" {}