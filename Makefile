beecow: beecow.go
	go build beecow.go
	if command -v setcap; then sudo setcap cap_net_bind_service+ep beecow; fi
