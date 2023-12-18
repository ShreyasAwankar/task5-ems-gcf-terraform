variable "project_id" {
  default = "terraform-cloud-functions-ems"
}

variable "region" {
  default = "us-central1"
}

variable "create_bucket" {
  description = "Set to true if you want to create a new storage bucket, false to import an existing one"
  default     = true
}

variable "functions" {
  type = map(object({
    zip        = string
    name       = string
    trigger    = string
    runtime    = string
    entrypoint = string
  }))

  default = {
    "createemployee" : {
      zip        = "function-zips/create-employee.zip"
      name       = "createEmployee"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "CreateEmployee"
    },
    "deleteemployee" : {
      zip        = "function-zips/delete-employee.zip"
      name       = "deleteEmployee"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "DeleteEmployee"
    },
    "searchemployees" : {
      zip        = "function-zips/search-employee.zip"
      name       = "searchEmployees"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "SearchEmployees"
    },
    "getemployeebyid" : {
      zip        = "function-zips/get-employee-by-id.zip"
      name       = "getEmployeeById"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "GetEmployee"
    },
    "getallemployees" : {
      zip        = "function-zips/get-all-employees.zip"
      name       = "getAllEmployees"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "GetAllEmployees"
    },
    "updateemployee" : {
      zip        = "function-zips/update-employee.zip"
      name       = "updateEmployee"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "UpdateEmployee"
    }
  }
}
