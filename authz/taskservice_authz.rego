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
# taskid may or may not be part of the path
allow {
    input.method == "GET"
    some taskid
    input.path = ["tasks", taskid]
    input.taskid = taskid
    input.user == input.owner
}
allow {
    input.method == "GET"
    input.path = ["tasks"]
    some taskid
    input.taskid = taskid
    input.user == input.owner
}

# Only owner can read a task
# Ownership information is provided as part of OPA's input
allow {
    input.method == "DELETE"
    some taskid
    input.path = ["tasks", taskid]
    input.taskid = taskid
    input.user == input.owner
}