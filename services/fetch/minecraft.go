package fetch

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func writePacket(w io.Writer, packetId uint64, message[]byte) error {
	packet := binary.AppendUvarint(nil, uint64(packetId))
	body := append(packet, message...)

	length := len(body)
	lengthBytes := binary.AppendUvarint(nil, uint64(length))

	_, err := w.Write(append(lengthBytes, body...))
	return err
}

func writeHandshake(w io.Writer, host string, port uint16) error {
	payload := binary.AppendUvarint(nil, 767)
	payload = binary.AppendUvarint(payload, uint64(len(host)))
	payload = append(payload, []byte(host)...)
	payload = binary.BigEndian.AppendUint16(payload, port)
	payload = binary.AppendUvarint(payload, 1)

	return writePacket(w, 0x00, payload)
}

func writePing(w io.Writer, currentTime int64) error {
	payload := binary.BigEndian.AppendUint64(nil, uint64(currentTime))
	return writePacket(w, 0x01, payload)
}

func readPong(r *bufio.Reader) (int64, error) {
	_, err := binary.ReadUvarint(r)
	packet, err := binary.ReadUvarint(r)
	if err != nil {
		return 0, err
	}

	if packet != 0x01 {
		return 0, errors.New("Bad packet")
	}

	var payload int64
	err = binary.Read(r, binary.BigEndian, &payload)
	return payload, err
}

func parseUrl(rawUrl string) (string, int, string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", 0, "", err
	}
	hostname, rawPort := u.Hostname(), u.Port()
	if strings.TrimSpace(rawPort) == "" {
		rawPort = "25565"
	}

	port, err := strconv.Atoi(rawPort)
	if err != nil {
		return "", 0, "", err
	}

	host := fmt.Sprintf("%s:%d", hostname, port)

	return host, port, hostname, nil
}

func Minecraft(url string) (Data, error) {
	host, port, hostname, err := parseUrl(url)
	if err != nil {
		return Data{}, err
	}

	conn, err := net.DialTimeout("tcp", host, 1*time.Second)
	if err != nil {
		return Data{}, err
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err := writeHandshake(conn, hostname, uint16(port)); err != nil {
		return Data{}, err
	}
	if err := writePing(conn, time.Now().UnixMilli()); err != nil {
		return Data{}, err
	}

	res, err := readPong(bufio.NewReader(conn))
	if err != nil {
		return Data{}, err
	}

	t := time.Now().UnixMilli()
	latency := t - res

	return Data{
		Status: true,
		Latency: int(latency),
	}, nil
}
