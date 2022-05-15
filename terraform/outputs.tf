# output "deployment_sha" {
#   value = sha1(join("", [
#     for file in fileset(path.module, "apigw_*.tf") : filesha1(file)
#   ]))
# }

