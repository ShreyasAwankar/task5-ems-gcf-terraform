
# Creating a storage bucket to store cloud function objects
resource "google_storage_bucket" "bucket" {
  project  = var.project_id
  name     = "${var.project_id}-bucket1"
  location = var.region
}

# data "archive_file" "function_src" {
#   for_each    = var.functions
#   type        = "zip"
#   output_path = "/tmp/${each.value.zip}"
#   source_dir  = "../function-zips"
# }

resource "google_storage_bucket_object" "function_zip" {
  for_each = var.functions
  name     = each.key
  bucket   = google_storage_bucket.bucket.name
  source   = each.value.zip
}


resource "google_cloudfunctions2_function" "function" {
  for_each = var.functions
  name     = each.value.name
  location = var.region

  build_config {
    runtime     = each.value.runtime
    entry_point = each.value.entrypoint

    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.function_zip[each.key].name
      }
    }
  }
  service_config {
    min_instance_count             = 1
    available_memory               = "128Mi"
    timeout_seconds                = 120
    all_traffic_on_latest_revision = false
    service_account_email          = "terraform-gcf@terraform-cloud-functions-ems.iam.gserviceaccount.com"
  }
}

resource "google_cloud_run_service_iam_member" "member" {
  for_each = var.functions

  location = google_cloudfunctions2_function.function[each.key].location
  service  = each.key
  role     = "roles/run.invoker"
  member   = "allUsers"
}
