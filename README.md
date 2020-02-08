![Go](https://github.com/ReverentEngineer/workspace/workflows/Go/badge.svg)

# Summary

Manages workspace in DigitalOcean

# Configuration

The workspace created is configured by $HOME/.config/workspace/config.yaml
```yaml
access_token: {access_token}
slug: centos-8-x86
size: s-1vcpu-1gb
region: nyc3
ssh_fingerprint: {ssh fingerprint}
```

# Usage

The usage is fairly simple for the moment.

To start the workspace:
```
workspace start
```

To stop the workspace:
```
workspace stop
```
