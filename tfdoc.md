## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_google"></a> [google](#provider\_google) | 5.9.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [google_cloud_run_service_iam_member.member](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_service_iam_member) | resource |
| [google_cloudfunctions2_function.function](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloudfunctions2_function) | resource |
| [google_storage_bucket.bucket](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/storage_bucket) | resource |
| [google_storage_bucket_object.function_zip](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/storage_bucket_object) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_functions"></a> [functions](#input\_functions) | n/a | <pre>map(object({<br>    zip        = string<br>    name       = string<br>    trigger    = string<br>    runtime    = string<br>    entrypoint = string<br>  }))</pre> | <pre>{<br>  "createemployee": {<br>    "entrypoint": "CreateEmployee",<br>    "name": "createEmployee",<br>    "runtime": "go121",<br>    "trigger": "http-trigger",<br>    "zip": "function-zips/create-employee.zip"<br>  },<br>  "deleteemployee": {<br>    "entrypoint": "DeleteEmployee",<br>    "name": "deleteEmployee",<br>    "runtime": "go121",<br>    "trigger": "http-trigger",<br>    "zip": "function-zips/delete-employee.zip"<br>  },<br>  "getallemployees": {<br>    "entrypoint": "GetAllEmployees",<br>    "name": "getAllEmployees",<br>    "runtime": "go121",<br>    "trigger": "http-trigger",<br>    "zip": "function-zips/get-all-employees.zip"<br>  },<br>  "getemployeebyid": {<br>    "entrypoint": "GetEmployee",<br>    "name": "getEmployeeById",<br>    "runtime": "go121",<br>    "trigger": "http-trigger",<br>    "zip": "function-zips/get-employee-by-id.zip"<br>  },<br>  "searchemployees": {<br>    "entrypoint": "SearchEmployees",<br>    "name": "searchEmployees",<br>    "runtime": "go121",<br>    "trigger": "http-trigger",<br>    "zip": "function-zips/search-employee.zip"<br>  },<br>  "updateemployee": {<br>    "entrypoint": "UpdateEmployee",<br>    "name": "updateEmployee",<br>    "runtime": "go121",<br>    "trigger": "http-trigger",<br>    "zip": "function-zips/update-employee.zip"<br>  }<br>}</pre> | no |
| <a name="input_project_id"></a> [project\_id](#input\_project\_id) | n/a | `string` | `"terraform-cloud-functions-ems"` | no |
| <a name="input_region"></a> [region](#input\_region) | n/a | `string` | `"us-central1"` | no |

## Outputs

No outputs.
