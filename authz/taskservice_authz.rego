package authz

# Only owner can create a task
# Ownership information is provided as part of OPA's input
default allow = false
allow {
    input.method == "POST"
    input.path = ["tasks"]
    input.user == input.owner
}

# Only owner can read a task
# Ownership information is provided as part of OPA's input
allow {
    input.method == "GET"
    some taskid
    input.path = ["tasks", taskid]
    input.user == input.owner
}
