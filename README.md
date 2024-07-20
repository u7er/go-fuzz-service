# go-fuzz-service
> Runtime of this code stores yours ip address. Do not run it as provided here. All responsibilities are yours.

# Build
```bash
go build -o fuzz ./cmd/main.go
```

# Run
```shell
./fuzz --address "*:43000"
```

# Daemon
```shell
cat << EOT > /etc/systemd/system/fuzz.service
[Unit]
Description=Go Web Fuzz
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/fuzz
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOT

systemctl enable fuzz.service \
&& systemctl start fuzz.service \
&& systemctl status fuzz.service
```

# Envoy
```shell
docker run -d --name envoy -e ENVOY_UID=0 --network host -v /root/envoy.yaml:/etc/envoy/envoy.yaml:ro envoyproxy/envoy:v1.20.0
```