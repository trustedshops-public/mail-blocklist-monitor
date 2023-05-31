resource "null_resource" "this" {
  count = var.run_make_build ? 1 : 0

  provisioner "local-exec" {
    working_dir = "${path.module}/../../../monitor"
    command     = "make build"
    interpreter = ["bash", "-c"]
  }

  triggers = {
    uuid = uuid()
  }
}
