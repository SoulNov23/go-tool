package netpoll

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func Socket(network string) (int, error) {
	switch network {
	case "tcp", "tcp4":
		fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, unix.IPPROTO_TCP)
		if err != nil {
			return 0, err
		}
		return fd, nil
	case "tcp6":
		fd, err := unix.Socket(unix.AF_INET6, unix.SOCK_STREAM|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, unix.IPPROTO_TCP)
		if err != nil {
			return 0, err
		}
		return fd, nil
	case "udp", "udp4":
		fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, unix.IPPROTO_UDP)
		if err != nil {
			return 0, err
		}
		return fd, nil
	case "udp6":
		fd, err := unix.Socket(unix.AF_INET6, unix.SOCK_DGRAM|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, unix.IPPROTO_UDP)
		if err != nil {
			return 0, err
		}
		return fd, nil
	case "unix":
		fd, err := unix.Socket(unix.AF_UNIX, unix.SOCK_STREAM|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, 0)
		if err != nil {
			return 0, err
		}
		return fd, nil
	case "unixgram":
		fd, err := unix.Socket(unix.AF_UNIX, unix.SOCK_DGRAM|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, 0)
		if err != nil {
			return 0, err
		}
		return fd, nil
	case "unixpacket":
		fd, err := unix.Socket(unix.AF_UNIX, unix.SOCK_SEQPACKET|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, 0)
		if err != nil {
			return 0, err
		}
		return fd, nil
	case "ip", "ip4":
		fd, err := unix.Socket(unix.AF_INET, unix.SOCK_RAW|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, unix.IPPROTO_IP)
		if err != nil {
			return 0, err
		}
		return fd, nil
	case "ip6":
		fd, err := unix.Socket(unix.AF_INET6, unix.SOCK_RAW|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, unix.IPPROTO_IPV6)
		if err != nil {
			return 0, err
		}
		return fd, nil
	default:
		return 0, fmt.Errorf("network[%s] not support", network)
	}
}
