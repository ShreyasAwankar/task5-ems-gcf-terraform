variable "project_id" {
  default = "terraform-cloud-functions-ems"
}

variable "region" {
  default = "us-central1"
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
    "create-employee" : {
      zip        = "function-zips/create-employee.zip"
      name       = "createEmployee"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "CreateEmployee"
    },
    "delete-employee" : {
      zip        = "function-zips/delete-employee.zip"
      name       = "deleteEmployee"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "DeleteEmployee"
    },
    "search-employees" : {
      zip        = "function-zips/search-employee.zip"
      name       = "searchEmployee"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "SearchEmployees"
    },
    "get-employee-by-id" : {
      zip        = "function-zips/get-employee-by-id.zip"
      name       = "getEmployeeById"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "GetEmployee"
    },
    "get-all-employees" : {
      zip        = "function-zips/get-all-employees.zip"
      name       = "getAllEmployees"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "GetAllEmployees"
    },
    "update-employee" : {
      zip        = "function-zips/update-employee.zip"
      name       = "updateEmployee"
      trigger    = "http-trigger"
      runtime    = "go121"
      entrypoint = "UpdateEmployee"
    }
  }
}
