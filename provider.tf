# specifying provider for our terraform configuration.
provider "google" {
  project = var.project_id
  region  = var.region
}

