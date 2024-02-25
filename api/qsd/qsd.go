package qsd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/saracen/go-qmp/api"
)

type Null struct {
}

// RegisterEvents registers the events this package is aware of with the provided client
func RegisterEvents(client api.Client) {
	client.RegisterEvent("JOB_STATUS_CHANGE", func() api.EventType { return &JobStatusChangeEvent{} })
	client.RegisterEvent("BLOCK_IMAGE_CORRUPTED", func() api.EventType { return &BlockImageCorruptedEvent{} })
	client.RegisterEvent("BLOCK_IO_ERROR", func() api.EventType { return &BlockIoErrorEvent{} })
	client.RegisterEvent("BLOCK_JOB_COMPLETED", func() api.EventType { return &BlockJobCompletedEvent{} })
	client.RegisterEvent("BLOCK_JOB_CANCELLED", func() api.EventType { return &BlockJobCancelledEvent{} })
	client.RegisterEvent("BLOCK_JOB_ERROR", func() api.EventType { return &BlockJobErrorEvent{} })
	client.RegisterEvent("BLOCK_JOB_READY", func() api.EventType { return &BlockJobReadyEvent{} })
	client.RegisterEvent("BLOCK_JOB_PENDING", func() api.EventType { return &BlockJobPendingEvent{} })
	client.RegisterEvent("BLOCK_WRITE_THRESHOLD", func() api.EventType { return &BlockWriteThresholdEvent{} })
	client.RegisterEvent("QUORUM_FAILURE", func() api.EventType { return &QuorumFailureEvent{} })
	client.RegisterEvent("QUORUM_REPORT_BAD", func() api.EventType { return &QuorumReportBadEvent{} })
	client.RegisterEvent("BLOCK_EXPORT_DELETED", func() api.EventType { return &BlockExportDeletedEvent{} })
	client.RegisterEvent("VSERPORT_CHANGE", func() api.EventType { return &VserportChangeEvent{} })
}

// IoOperationType An enumeration of the I/O operation types
type IoOperationType string

const (
	// IoOperationTypeRead read operation
	IoOperationTypeRead IoOperationType = "read"
	// IoOperationTypeWrite write operation
	IoOperationTypeWrite IoOperationType = "write"
)

// OnOffAuto
type OnOffAuto string

const (
	// OnOffAutoAuto QEMU selects the value between on and off
	OnOffAutoAuto OnOffAuto = "auto"
	// OnOffAutoOn Enabled
	OnOffAutoOn OnOffAuto = "on"
	// OnOffAutoOff Disabled
	OnOffAutoOff OnOffAuto = "off"
)

// OnOffSplit
type OnOffSplit string

const (
	// OnOffSplitOn Enabled
	OnOffSplitOn OnOffSplit = "on"
	// OnOffSplitOff Disabled
	OnOffSplitOff OnOffSplit = "off"
	// OnOffSplitSplit Mixed
	OnOffSplitSplit OnOffSplit = "split"
)

// StrOrNull
//
// This is a string value or the explicit lack of a string (null pointer in C). Intended for cases when 'optional absent' already has a different meaning.
type StrOrNull struct {
	// S the string value
	S *string `json:"-"`
	// N no string value
	N *Null `json:"-"`
}

func (a StrOrNull) MarshalJSON() ([]byte, error) {
	switch {
	case a.S != nil:
		return json.Marshal(a.S)
	case a.N != nil:
		return json.Marshal(a.N)
	}

	return nil, fmt.Errorf("unknown type")
}

// OffAutoPCIBAR An enumeration of options for specifying a PCI BAR
type OffAutoPCIBAR string

const (
	// OffAutoPCIBAROff The specified feature is disabled
	OffAutoPCIBAROff OffAutoPCIBAR = "off"
	// OffAutoPCIBARAuto The PCI BAR for the feature is automatically selected
	OffAutoPCIBARAuto OffAutoPCIBAR = "auto"
	// OffAutoPCIBARBar0 PCI BAR0 is used for the feature
	OffAutoPCIBARBar0 OffAutoPCIBAR = "bar0"
	// OffAutoPCIBARBar1 PCI BAR1 is used for the feature
	OffAutoPCIBARBar1 OffAutoPCIBAR = "bar1"
	// OffAutoPCIBARBar2 PCI BAR2 is used for the feature
	OffAutoPCIBARBar2 OffAutoPCIBAR = "bar2"
	// OffAutoPCIBARBar3 PCI BAR3 is used for the feature
	OffAutoPCIBARBar3 OffAutoPCIBAR = "bar3"
	// OffAutoPCIBARBar4 PCI BAR4 is used for the feature
	OffAutoPCIBARBar4 OffAutoPCIBAR = "bar4"
	// OffAutoPCIBARBar5 PCI BAR5 is used for the feature
	OffAutoPCIBARBar5 OffAutoPCIBAR = "bar5"
)

// PCIELinkSpeed An enumeration of PCIe link speeds in units of GT/s
type PCIELinkSpeed string

const (
	// PCIELinkSpeed25 2.5GT/s
	PCIELinkSpeed25 PCIELinkSpeed = "2_5"
	// PCIELinkSpeed5 5.0GT/s
	PCIELinkSpeed5 PCIELinkSpeed = "5"
	// PCIELinkSpeed8 8.0GT/s
	PCIELinkSpeed8 PCIELinkSpeed = "8"
	// PCIELinkSpeed16 16.0GT/s
	PCIELinkSpeed16 PCIELinkSpeed = "16"
)

// PCIELinkWidth An enumeration of PCIe link width
type PCIELinkWidth string

const (
	// PCIELinkWidth1 x1
	PCIELinkWidth1 PCIELinkWidth = "1"
	// PCIELinkWidth2 x2
	PCIELinkWidth2 PCIELinkWidth = "2"
	// PCIELinkWidth4 x4
	PCIELinkWidth4 PCIELinkWidth = "4"
	// PCIELinkWidth8 x8
	PCIELinkWidth8 PCIELinkWidth = "8"
	// PCIELinkWidth12 x12
	PCIELinkWidth12 PCIELinkWidth = "12"
	// PCIELinkWidth16 x16
	PCIELinkWidth16 PCIELinkWidth = "16"
	// PCIELinkWidth32 x32
	PCIELinkWidth32 PCIELinkWidth = "32"
)

// HostMemPolicy Host memory policy types
type HostMemPolicy string

const (
	// HostMemPolicyDefault restore default policy, remove any nondefault policy
	HostMemPolicyDefault HostMemPolicy = "default"
	// HostMemPolicyPreferred set the preferred host nodes for allocation
	HostMemPolicyPreferred HostMemPolicy = "preferred"
	// HostMemPolicyBind a strict policy that restricts memory allocation to the host nodes specified
	HostMemPolicyBind HostMemPolicy = "bind"
	// HostMemPolicyInterleave memory allocations are interleaved across the set of host nodes specified
	HostMemPolicyInterleave HostMemPolicy = "interleave"
)

// NetFilterDirection Indicates whether a netfilter is attached to a netdev's transmit queue or receive queue or both.
type NetFilterDirection string

const (
	// NetFilterDirectionAll the filter is attached both to the receive and the transmit queue of the netdev (default).
	NetFilterDirectionAll NetFilterDirection = "all"
	// NetFilterDirectionRx the filter is attached to the receive queue of the netdev, where it will receive packets sent to the netdev.
	NetFilterDirectionRx NetFilterDirection = "rx"
	// NetFilterDirectionTx the filter is attached to the transmit queue of the netdev, where it will receive packets sent by the netdev.
	NetFilterDirectionTx NetFilterDirection = "tx"
)

// GrabToggleKeys Keys to toggle input-linux between host and guest.
type GrabToggleKeys string

const (
	GrabToggleKeysCtrlCtrl       GrabToggleKeys = "ctrl-ctrl"
	GrabToggleKeysAltAlt         GrabToggleKeys = "alt-alt"
	GrabToggleKeysShiftShift     GrabToggleKeys = "shift-shift"
	GrabToggleKeysMetaMeta       GrabToggleKeys = "meta-meta"
	GrabToggleKeysScrolllock     GrabToggleKeys = "scrolllock"
	GrabToggleKeysCtrlScrolllock GrabToggleKeys = "ctrl-scrolllock"
)

// HumanReadableText
type HumanReadableText struct {
	// HumanReadableText Formatted output intended for humans.
	HumanReadableText string `json:"human-readable-text"`
}

// NetworkAddressFamily The network address family
type NetworkAddressFamily string

const (
	// NetworkAddressFamilyIpv4 IPV4 family
	NetworkAddressFamilyIpv4 NetworkAddressFamily = "ipv4"
	// NetworkAddressFamilyIpv6 IPV6 family
	NetworkAddressFamilyIpv6 NetworkAddressFamily = "ipv6"
	// NetworkAddressFamilyUnix unix socket
	NetworkAddressFamilyUnix NetworkAddressFamily = "unix"
	// NetworkAddressFamilyVsock vsock family (since 2.8)
	NetworkAddressFamilyVsock NetworkAddressFamily = "vsock"
	// NetworkAddressFamilyUnknown otherwise
	NetworkAddressFamilyUnknown NetworkAddressFamily = "unknown"
)

// InetSocketAddressBase
type InetSocketAddressBase struct {
	// Host host part of the address
	Host string `json:"host"`
	// Port port part of the address
	Port string `json:"port"`
}

// InetSocketAddress
//
// Captures a socket address or address range in the Internet namespace.
type InetSocketAddress struct {
	InetSocketAddressBase

	// Numeric true if the host/port are guaranteed to be numeric, false if name resolution should be attempted. Defaults to false. (Since 2.9)
	Numeric *bool `json:"numeric,omitempty"`
	// To If present, this is range of possible addresses, with port between @port and @to.
	To *uint16 `json:"to,omitempty"`
	// Ipv4 whether to accept IPv4 addresses, default try both IPv4 and IPv6
	Ipv4 *bool `json:"ipv4,omitempty"`
	// Ipv6 whether to accept IPv6 addresses, default try both IPv4 and IPv6
	Ipv6 *bool `json:"ipv6,omitempty"`
	// KeepAlive enable keep-alive when connecting to this socket. Not supported for passive sockets. (Since 4.2)
	KeepAlive *bool `json:"keep-alive,omitempty"`
	// Mptcp enable multi-path TCP. (Since 6.1)
	Mptcp *bool `json:"mptcp,omitempty"`
}

// UnixSocketAddress
//
// Captures a socket address in the local ("Unix socket") namespace.
type UnixSocketAddress struct {
	// Path filesystem path to use
	Path string `json:"path"`
	// Abstract if true, this is a Linux abstract socket address. @path will be prefixed by a null byte, and optionally padded with null bytes. Defaults to false. (Since 5.1)
	Abstract *bool `json:"abstract,omitempty"`
	// Tight if false, pad an abstract socket address with enough null bytes to make it fill struct sockaddr_un member sun_path. Defaults to true. (Since 5.1)
	Tight *bool `json:"tight,omitempty"`
}

// VsockSocketAddress
//
// Captures a socket address in the vsock namespace.
type VsockSocketAddress struct {
	// Cid unique host identifier
	Cid string `json:"cid"`
	// Port port
	Port string `json:"port"`
}

// FdSocketAddress
//
// A file descriptor name or number.
type FdSocketAddress struct {
	// Str decimal is for file descriptor number, otherwise it's a file descriptor name. Named file descriptors are permitted in monitor commands, in combination with the 'getfd' command. Decimal file descriptors are permitted at startup or other contexts where no monitor context is active.
	Str string `json:"str"`
}

// InetSocketAddressWrapper
type InetSocketAddressWrapper struct {
	// Data internet domain socket address
	Data InetSocketAddress `json:"data"`
}

// UnixSocketAddressWrapper
type UnixSocketAddressWrapper struct {
	// Data UNIX domain socket address
	Data UnixSocketAddress `json:"data"`
}

// VsockSocketAddressWrapper
type VsockSocketAddressWrapper struct {
	// Data VSOCK domain socket address
	Data VsockSocketAddress `json:"data"`
}

// FdSocketAddressWrapper
type FdSocketAddressWrapper struct {
	// Data file descriptor name or number
	Data FdSocketAddress `json:"data"`
}

// SocketAddressLegacy
//
// Captures the address of a socket, which could also be a named file descriptor
type SocketAddressLegacy struct {
	// Discriminator: type

	// Type Transport type
	Type SocketAddressType `json:"type"`

	Inet  *InetSocketAddressWrapper  `json:"-"`
	Unix  *UnixSocketAddressWrapper  `json:"-"`
	Vsock *VsockSocketAddressWrapper `json:"-"`
	Fd    *FdSocketAddressWrapper    `json:"-"`
}

func (u SocketAddressLegacy) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "inet":
		if u.Inet == nil {
			return nil, fmt.Errorf("expected Inet to be set")
		}

		return json.Marshal(struct {
			Type SocketAddressType `json:"type"`
			*InetSocketAddressWrapper
		}{
			Type:                     u.Type,
			InetSocketAddressWrapper: u.Inet,
		})
	case "unix":
		if u.Unix == nil {
			return nil, fmt.Errorf("expected Unix to be set")
		}

		return json.Marshal(struct {
			Type SocketAddressType `json:"type"`
			*UnixSocketAddressWrapper
		}{
			Type:                     u.Type,
			UnixSocketAddressWrapper: u.Unix,
		})
	case "vsock":
		if u.Vsock == nil {
			return nil, fmt.Errorf("expected Vsock to be set")
		}

		return json.Marshal(struct {
			Type SocketAddressType `json:"type"`
			*VsockSocketAddressWrapper
		}{
			Type:                      u.Type,
			VsockSocketAddressWrapper: u.Vsock,
		})
	case "fd":
		if u.Fd == nil {
			return nil, fmt.Errorf("expected Fd to be set")
		}

		return json.Marshal(struct {
			Type SocketAddressType `json:"type"`
			*FdSocketAddressWrapper
		}{
			Type:                   u.Type,
			FdSocketAddressWrapper: u.Fd,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// SocketAddressType Available SocketAddress types
type SocketAddressType string

const (
	// SocketAddressTypeInet Internet address
	SocketAddressTypeInet SocketAddressType = "inet"
	// SocketAddressTypeUnix Unix domain socket
	SocketAddressTypeUnix SocketAddressType = "unix"
	// SocketAddressTypeVsock VMCI address
	SocketAddressTypeVsock SocketAddressType = "vsock"
	// SocketAddressTypeFd Socket file descriptor
	SocketAddressTypeFd SocketAddressType = "fd"
)

// SocketAddress
//
// Captures the address of a socket, which could also be a socket file descriptor
type SocketAddress struct {
	// Discriminator: type

	// Type Transport type
	Type SocketAddressType `json:"type"`

	Inet  *InetSocketAddress  `json:"-"`
	Unix  *UnixSocketAddress  `json:"-"`
	Vsock *VsockSocketAddress `json:"-"`
	Fd    *FdSocketAddress    `json:"-"`
}

func (u SocketAddress) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "inet":
		if u.Inet == nil {
			return nil, fmt.Errorf("expected Inet to be set")
		}

		return json.Marshal(struct {
			Type SocketAddressType `json:"type"`
			*InetSocketAddress
		}{
			Type:              u.Type,
			InetSocketAddress: u.Inet,
		})
	case "unix":
		if u.Unix == nil {
			return nil, fmt.Errorf("expected Unix to be set")
		}

		return json.Marshal(struct {
			Type SocketAddressType `json:"type"`
			*UnixSocketAddress
		}{
			Type:              u.Type,
			UnixSocketAddress: u.Unix,
		})
	case "vsock":
		if u.Vsock == nil {
			return nil, fmt.Errorf("expected Vsock to be set")
		}

		return json.Marshal(struct {
			Type SocketAddressType `json:"type"`
			*VsockSocketAddress
		}{
			Type:               u.Type,
			VsockSocketAddress: u.Vsock,
		})
	case "fd":
		if u.Fd == nil {
			return nil, fmt.Errorf("expected Fd to be set")
		}

		return json.Marshal(struct {
			Type SocketAddressType `json:"type"`
			*FdSocketAddress
		}{
			Type:            u.Type,
			FdSocketAddress: u.Fd,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// QCryptoTLSCredsEndpoint The type of network endpoint that will be using the credentials. Most types of credential require different setup / structures depending on whether they will be used in a server versus a client.
type QCryptoTLSCredsEndpoint string

const (
	// QcryptoTlsCredsEndpointClient the network endpoint is acting as the client
	QcryptoTlsCredsEndpointClient QCryptoTLSCredsEndpoint = "client"
	// QcryptoTlsCredsEndpointServer the network endpoint is acting as the server
	QcryptoTlsCredsEndpointServer QCryptoTLSCredsEndpoint = "server"
)

// QCryptoSecretFormat The data format that the secret is provided in
type QCryptoSecretFormat string

const (
	// QcryptoSecretFormatRaw raw bytes. When encoded in JSON only valid UTF-8 sequences can be used
	QcryptoSecretFormatRaw QCryptoSecretFormat = "raw"
	// QcryptoSecretFormatBase64 arbitrary base64 encoded binary data
	QcryptoSecretFormatBase64 QCryptoSecretFormat = "base64"
)

// QCryptoHashAlgorithm The supported algorithms for computing content digests
type QCryptoHashAlgorithm string

const (
	// QcryptoHashAlgMd5 MD5. Should not be used in any new code, legacy compat only
	QcryptoHashAlgMd5 QCryptoHashAlgorithm = "md5"
	// QcryptoHashAlgSha1 SHA-1. Should not be used in any new code, legacy compat only
	QcryptoHashAlgSha1 QCryptoHashAlgorithm = "sha1"
	// QcryptoHashAlgSha224 SHA-224. (since 2.7)
	QcryptoHashAlgSha224 QCryptoHashAlgorithm = "sha224"
	// QcryptoHashAlgSha256 SHA-256. Current recommended strong hash.
	QcryptoHashAlgSha256 QCryptoHashAlgorithm = "sha256"
	// QcryptoHashAlgSha384 SHA-384. (since 2.7)
	QcryptoHashAlgSha384 QCryptoHashAlgorithm = "sha384"
	// QcryptoHashAlgSha512 SHA-512. (since 2.7)
	QcryptoHashAlgSha512 QCryptoHashAlgorithm = "sha512"
	// QcryptoHashAlgRipemd160 RIPEMD-160. (since 2.7)
	QcryptoHashAlgRipemd160 QCryptoHashAlgorithm = "ripemd160"
)

// QCryptoCipherAlgorithm The supported algorithms for content encryption ciphers
type QCryptoCipherAlgorithm string

const (
	// QcryptoCipherAlgAes128 AES with 128 bit / 16 byte keys
	QcryptoCipherAlgAes128 QCryptoCipherAlgorithm = "aes-128"
	// QcryptoCipherAlgAes192 AES with 192 bit / 24 byte keys
	QcryptoCipherAlgAes192 QCryptoCipherAlgorithm = "aes-192"
	// QcryptoCipherAlgAes256 AES with 256 bit / 32 byte keys
	QcryptoCipherAlgAes256 QCryptoCipherAlgorithm = "aes-256"
	// QcryptoCipherAlgDes DES with 56 bit / 8 byte keys. Do not use except in VNC. (since 6.1)
	QcryptoCipherAlgDes QCryptoCipherAlgorithm = "des"
	// QcryptoCipherAlg3Des 3DES(EDE) with 192 bit / 24 byte keys (since 2.9)
	QcryptoCipherAlg3Des QCryptoCipherAlgorithm = "3des"
	// QcryptoCipherAlgCast5128 Cast5 with 128 bit / 16 byte keys
	QcryptoCipherAlgCast5128 QCryptoCipherAlgorithm = "cast5-128"
	// QcryptoCipherAlgSerpent128 Serpent with 128 bit / 16 byte keys
	QcryptoCipherAlgSerpent128 QCryptoCipherAlgorithm = "serpent-128"
	// QcryptoCipherAlgSerpent192 Serpent with 192 bit / 24 byte keys
	QcryptoCipherAlgSerpent192 QCryptoCipherAlgorithm = "serpent-192"
	// QcryptoCipherAlgSerpent256 Serpent with 256 bit / 32 byte keys
	QcryptoCipherAlgSerpent256 QCryptoCipherAlgorithm = "serpent-256"
	// QcryptoCipherAlgTwofish128 Twofish with 128 bit / 16 byte keys
	QcryptoCipherAlgTwofish128 QCryptoCipherAlgorithm = "twofish-128"
	// QcryptoCipherAlgTwofish192 Twofish with 192 bit / 24 byte keys
	QcryptoCipherAlgTwofish192 QCryptoCipherAlgorithm = "twofish-192"
	// QcryptoCipherAlgTwofish256 Twofish with 256 bit / 32 byte keys
	QcryptoCipherAlgTwofish256 QCryptoCipherAlgorithm = "twofish-256"
	// QcryptoCipherAlgSm4 SM4 with 128 bit / 16 byte keys (since 9.0)
	QcryptoCipherAlgSm4 QCryptoCipherAlgorithm = "sm4"
)

// QCryptoCipherMode The supported modes for content encryption ciphers
type QCryptoCipherMode string

const (
	// QcryptoCipherModeEcb Electronic Code Book
	QcryptoCipherModeEcb QCryptoCipherMode = "ecb"
	// QcryptoCipherModeCbc Cipher Block Chaining
	QcryptoCipherModeCbc QCryptoCipherMode = "cbc"
	// QcryptoCipherModeXts XEX with tweaked code book and ciphertext stealing
	QcryptoCipherModeXts QCryptoCipherMode = "xts"
	// QcryptoCipherModeCtr Counter (Since 2.8)
	QcryptoCipherModeCtr QCryptoCipherMode = "ctr"
)

// QCryptoIVGenAlgorithm The supported algorithms for generating initialization vectors for full disk encryption. The 'plain' generator should not be used for disks with sector numbers larger than 2^32, except where compatibility with pre-existing Linux dm-crypt volumes is required.
type QCryptoIVGenAlgorithm string

const (
	// QcryptoIvgenAlgPlain 64-bit sector number truncated to 32-bits
	QcryptoIvgenAlgPlain QCryptoIVGenAlgorithm = "plain"
	// QcryptoIvgenAlgPlain64 64-bit sector number
	QcryptoIvgenAlgPlain64 QCryptoIVGenAlgorithm = "plain64"
	// QcryptoIvgenAlgEssiv 64-bit sector number encrypted with a hash of the encryption key
	QcryptoIvgenAlgEssiv QCryptoIVGenAlgorithm = "essiv"
)

// QCryptoBlockFormat The supported full disk encryption formats
type QCryptoBlockFormat string

const (
	// QCryptoBlockFormatQcow QCow/QCow2 built-in AES-CBC encryption. Use only for liberating data from old images.
	QCryptoBlockFormatQcow QCryptoBlockFormat = "qcow"
	// QCryptoBlockFormatLuks LUKS encryption format. Recommended for new images
	QCryptoBlockFormatLuks QCryptoBlockFormat = "luks"
)

// QCryptoBlockOptionsBase
//
// The common options that apply to all full disk encryption formats
type QCryptoBlockOptionsBase struct {
	// Format the encryption format
	Format QCryptoBlockFormat `json:"format"`
}

// QCryptoBlockOptionsQCow
//
// The options that apply to QCow/QCow2 AES-CBC encryption format
type QCryptoBlockOptionsQCow struct {
	// KeySecret the ID of a QCryptoSecret object providing the decryption key. Mandatory except when probing image for metadata only.
	KeySecret *string `json:"key-secret,omitempty"`
}

// QCryptoBlockOptionsLUKS
//
// The options that apply to LUKS encryption format
type QCryptoBlockOptionsLUKS struct {
	// KeySecret the ID of a QCryptoSecret object providing the decryption key. Mandatory except when probing image for metadata only.
	KeySecret *string `json:"key-secret,omitempty"`
}

// QCryptoBlockCreateOptionsLUKS
//
// The options that apply to LUKS encryption format initialization
type QCryptoBlockCreateOptionsLUKS struct {
	QCryptoBlockOptionsLUKS

	// CipherAlg the cipher algorithm for data encryption Currently defaults to 'aes-256'.
	CipherAlg *QCryptoCipherAlgorithm `json:"cipher-alg,omitempty"`
	// CipherMode the cipher mode for data encryption Currently defaults to 'xts'
	CipherMode *QCryptoCipherMode `json:"cipher-mode,omitempty"`
	// IvgenAlg the initialization vector generator Currently defaults to 'plain64'
	IvgenAlg *QCryptoIVGenAlgorithm `json:"ivgen-alg,omitempty"`
	// IvgenHashAlg the initialization vector generator hash Currently defaults to 'sha256'
	IvgenHashAlg *QCryptoHashAlgorithm `json:"ivgen-hash-alg,omitempty"`
	// HashAlg the master key hash algorithm Currently defaults to 'sha256'
	HashAlg *QCryptoHashAlgorithm `json:"hash-alg,omitempty"`
	// IterTime number of milliseconds to spend in PBKDF passphrase processing. Currently defaults to 2000. (since 2.8)
	IterTime *int64 `json:"iter-time,omitempty"`
	// DetachedHeader create a detached LUKS header. (since 9.0)
	DetachedHeader *bool `json:"detached-header,omitempty"`
}

// QCryptoBlockOpenOptions
//
// The options that are available for all encryption formats when opening an existing volume
type QCryptoBlockOpenOptions struct {
	// Discriminator: format

	QCryptoBlockOptionsBase

	Qcow *QCryptoBlockOptionsQCow `json:"-"`
	Luks *QCryptoBlockOptionsLUKS `json:"-"`
}

func (u QCryptoBlockOpenOptions) MarshalJSON() ([]byte, error) {
	switch u.Format {
	case "qcow":
		if u.Qcow == nil {
			return nil, fmt.Errorf("expected Qcow to be set")
		}

		return json.Marshal(struct {
			QCryptoBlockOptionsBase
			*QCryptoBlockOptionsQCow
		}{
			QCryptoBlockOptionsBase: u.QCryptoBlockOptionsBase,
			QCryptoBlockOptionsQCow: u.Qcow,
		})
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			QCryptoBlockOptionsBase
			*QCryptoBlockOptionsLUKS
		}{
			QCryptoBlockOptionsBase: u.QCryptoBlockOptionsBase,
			QCryptoBlockOptionsLUKS: u.Luks,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// QCryptoBlockCreateOptions
//
// The options that are available for all encryption formats when initializing a new volume
type QCryptoBlockCreateOptions struct {
	// Discriminator: format

	QCryptoBlockOptionsBase

	Qcow *QCryptoBlockOptionsQCow       `json:"-"`
	Luks *QCryptoBlockCreateOptionsLUKS `json:"-"`
}

func (u QCryptoBlockCreateOptions) MarshalJSON() ([]byte, error) {
	switch u.Format {
	case "qcow":
		if u.Qcow == nil {
			return nil, fmt.Errorf("expected Qcow to be set")
		}

		return json.Marshal(struct {
			QCryptoBlockOptionsBase
			*QCryptoBlockOptionsQCow
		}{
			QCryptoBlockOptionsBase: u.QCryptoBlockOptionsBase,
			QCryptoBlockOptionsQCow: u.Qcow,
		})
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			QCryptoBlockOptionsBase
			*QCryptoBlockCreateOptionsLUKS
		}{
			QCryptoBlockOptionsBase:       u.QCryptoBlockOptionsBase,
			QCryptoBlockCreateOptionsLUKS: u.Luks,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// QCryptoBlockInfoBase
//
// The common information that applies to all full disk encryption formats
type QCryptoBlockInfoBase struct {
	// Format the encryption format
	Format QCryptoBlockFormat `json:"format"`
}

// QCryptoBlockInfoLUKSSlot
//
// Information about the LUKS block encryption key slot options
type QCryptoBlockInfoLUKSSlot struct {
	// Active whether the key slot is currently in use
	Active bool `json:"active"`
	// Iters number of PBKDF2 iterations for key material
	Iters *int64 `json:"iters,omitempty"`
	// Stripes number of stripes for splitting key material
	Stripes *int64 `json:"stripes,omitempty"`
	// KeyOffset offset to the key material in bytes
	KeyOffset int64 `json:"key-offset"`
}

// QCryptoBlockInfoLUKS
//
// Information about the LUKS block encryption options
type QCryptoBlockInfoLUKS struct {
	// CipherAlg the cipher algorithm for data encryption
	CipherAlg QCryptoCipherAlgorithm `json:"cipher-alg"`
	// CipherMode the cipher mode for data encryption
	CipherMode QCryptoCipherMode `json:"cipher-mode"`
	// IvgenAlg the initialization vector generator
	IvgenAlg QCryptoIVGenAlgorithm `json:"ivgen-alg"`
	// IvgenHashAlg the initialization vector generator hash
	IvgenHashAlg *QCryptoHashAlgorithm `json:"ivgen-hash-alg,omitempty"`
	// HashAlg the master key hash algorithm
	HashAlg QCryptoHashAlgorithm `json:"hash-alg"`
	// DetachedHeader whether the LUKS header is detached (Since 9.0)
	DetachedHeader bool `json:"detached-header"`
	// PayloadOffset offset to the payload data in bytes
	PayloadOffset int64 `json:"payload-offset"`
	// MasterKeyIters number of PBKDF2 iterations for key material
	MasterKeyIters int64 `json:"master-key-iters"`
	// Uuid unique identifier for the volume
	Uuid string `json:"uuid"`
	// Slots information about each key slot
	Slots []QCryptoBlockInfoLUKSSlot `json:"slots"`
}

// QCryptoBlockInfo
//
// Information about the block encryption options
type QCryptoBlockInfo struct {
	// Discriminator: format

	QCryptoBlockInfoBase

	Luks *QCryptoBlockInfoLUKS `json:"-"`
}

func (u QCryptoBlockInfo) MarshalJSON() ([]byte, error) {
	switch u.Format {
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			QCryptoBlockInfoBase
			*QCryptoBlockInfoLUKS
		}{
			QCryptoBlockInfoBase: u.QCryptoBlockInfoBase,
			QCryptoBlockInfoLUKS: u.Luks,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// QCryptoBlockLUKSKeyslotState Defines state of keyslots that are affected by the update
type QCryptoBlockLUKSKeyslotState string

const (
	// QCryptoBlockLUKSKeyslotStateActive The slots contain the given password and marked as active
	QCryptoBlockLUKSKeyslotStateActive QCryptoBlockLUKSKeyslotState = "active"
	// QCryptoBlockLUKSKeyslotStateInactive The slots are erased (contain garbage) and marked as inactive
	QCryptoBlockLUKSKeyslotStateInactive QCryptoBlockLUKSKeyslotState = "inactive"
)

// QCryptoBlockAmendOptionsLUKS
//
// This struct defines the update parameters that activate/de-activate set of keyslots
type QCryptoBlockAmendOptionsLUKS struct {
	// State the desired state of the keyslots
	State QCryptoBlockLUKSKeyslotState `json:"state"`
	// NewSecret The ID of a QCryptoSecret object providing the password to be written into added active keyslots
	NewSecret *string `json:"new-secret,omitempty"`
	// OldSecret Optional (for deactivation only) If given will deactivate all keyslots that match password located in QCryptoSecret with this ID
	OldSecret *string `json:"old-secret,omitempty"`
	// Keyslot Optional. ID of the keyslot to activate/deactivate. For keyslot activation, keyslot should not be active already (this is unsafe to update an active keyslot), but possible if 'force' parameter is given. If keyslot is not given, first free keyslot will be written. For keyslot deactivation, this parameter specifies the exact keyslot to deactivate
	Keyslot *int64 `json:"keyslot,omitempty"`
	// IterTime Optional (for activation only) Number of milliseconds to spend in PBKDF passphrase processing for the newly activated keyslot. Currently defaults to 2000.
	IterTime *int64 `json:"iter-time,omitempty"`
	// Secret Optional. The ID of a QCryptoSecret object providing the password to use to retrieve current master key. Defaults to the same secret that was used to open the image
	Secret *string `json:"secret,omitempty"`
}

// QCryptoBlockAmendOptions
//
// The options that are available for all encryption formats when amending encryption settings
type QCryptoBlockAmendOptions struct {
	// Discriminator: format

	QCryptoBlockOptionsBase

	Luks *QCryptoBlockAmendOptionsLUKS `json:"-"`
}

func (u QCryptoBlockAmendOptions) MarshalJSON() ([]byte, error) {
	switch u.Format {
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			QCryptoBlockOptionsBase
			*QCryptoBlockAmendOptionsLUKS
		}{
			QCryptoBlockOptionsBase:      u.QCryptoBlockOptionsBase,
			QCryptoBlockAmendOptionsLUKS: u.Luks,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// SecretCommonProperties
//
// Properties for objects of classes derived from secret-common.
type SecretCommonProperties struct {
	// Loaded if true, the secret is loaded immediately when applying this option and will probably fail when processing the next option. Don't use; only provided for compatibility.
	Loaded *bool `json:"loaded,omitempty"`
	// Format the data format that the secret is provided in
	Format *QCryptoSecretFormat `json:"format,omitempty"`
	// Keyid the name of another secret that should be used to decrypt the provided data. If not present, the data is assumed to be unencrypted.
	Keyid *string `json:"keyid,omitempty"`
	// Iv the random initialization vector used for encryption of this particular secret. Should be a base64 encrypted string of the 16-byte IV. Mandatory if @keyid is given. Ignored if @keyid is absent.
	Iv *string `json:"iv,omitempty"`
}

// SecretProperties
//
// Properties for secret objects. Either @data or @file must be provided, but not both.
type SecretProperties struct {
	SecretCommonProperties

	// Data the associated with the secret from
	Data *string `json:"data,omitempty"`
	// File the filename to load the data associated with the secret from
	File *string `json:"file,omitempty"`
}

// SecretKeyringProperties
//
// Properties for secret_keyring objects.
type SecretKeyringProperties struct {
	SecretCommonProperties

	// Serial serial number that identifies a key to get from the kernel
	Serial int32 `json:"serial"`
}

// TlsCredsProperties
//
// Properties for objects of classes derived from tls-creds.
type TlsCredsProperties struct {
	// VerifyPeer if true the peer credentials will be verified once the handshake is completed. This is a no-op for anonymous
	VerifyPeer *bool `json:"verify-peer,omitempty"`
	// Dir the path of the directory that contains the credential files
	Dir *string `json:"dir,omitempty"`
	// Endpoint whether the QEMU network backend that uses the credentials will be acting as a client or as a server
	Endpoint *QCryptoTLSCredsEndpoint `json:"endpoint,omitempty"`
	// Priority a gnutls priority string as described at
	Priority *string `json:"priority,omitempty"`
}

// TlsCredsAnonProperties
//
// Properties for tls-creds-anon objects.
type TlsCredsAnonProperties struct {
	TlsCredsProperties

	// Loaded if true, the credentials are loaded immediately when applying this option and will ignore options that are processed later. Don't use; only provided for compatibility.
	Loaded *bool `json:"loaded,omitempty"`
}

// TlsCredsPskProperties
//
// Properties for tls-creds-psk objects.
type TlsCredsPskProperties struct {
	TlsCredsProperties

	// Loaded if true, the credentials are loaded immediately when applying this option and will ignore options that are processed later. Don't use; only provided for compatibility.
	Loaded *bool `json:"loaded,omitempty"`
	// Username the username which will be sent to the server. For clients only. If absent, "qemu" is sent and the property will read back as an empty string.
	Username *string `json:"username,omitempty"`
}

// TlsCredsX509Properties
//
// Properties for tls-creds-x509 objects.
type TlsCredsX509Properties struct {
	TlsCredsProperties

	// Loaded if true, the credentials are loaded immediately when applying this option and will ignore options that are processed later. Don't use; only provided for compatibility.
	Loaded *bool `json:"loaded,omitempty"`
	// SanityCheck if true, perform some sanity checks before using the
	SanityCheck *bool `json:"sanity-check,omitempty"`
	// Passwordid For the server-key.pem and client-key.pem files which contain sensitive private keys, it is possible to use an encrypted version by providing the @passwordid parameter. This provides the ID of a previously created secret object containing the password for decryption.
	Passwordid *string `json:"passwordid,omitempty"`
}

// QCryptoAkCipherAlgorithm The supported algorithms for asymmetric encryption ciphers
type QCryptoAkCipherAlgorithm string

const (
	// QcryptoAkcipherAlgRsa RSA algorithm
	QcryptoAkcipherAlgRsa QCryptoAkCipherAlgorithm = "rsa"
)

// QCryptoAkCipherKeyType The type of asymmetric keys.
type QCryptoAkCipherKeyType string

const (
	QcryptoAkcipherKeyTypePublic  QCryptoAkCipherKeyType = "public"
	QcryptoAkcipherKeyTypePrivate QCryptoAkCipherKeyType = "private"
)

// QCryptoRSAPaddingAlgorithm The padding algorithm for RSA.
type QCryptoRSAPaddingAlgorithm string

const (
	// QcryptoRsaPaddingAlgRaw no padding used
	QcryptoRsaPaddingAlgRaw QCryptoRSAPaddingAlgorithm = "raw"
	// QcryptoRsaPaddingAlgPkcs1 pkcs1#v1.5
	QcryptoRsaPaddingAlgPkcs1 QCryptoRSAPaddingAlgorithm = "pkcs1"
)

// QCryptoAkCipherOptionsRSA
//
// Specific parameters for RSA algorithm.
type QCryptoAkCipherOptionsRSA struct {
	// HashAlg QCryptoHashAlgorithm
	HashAlg QCryptoHashAlgorithm `json:"hash-alg"`
	// PaddingAlg QCryptoRSAPaddingAlgorithm
	PaddingAlg QCryptoRSAPaddingAlgorithm `json:"padding-alg"`
}

// QCryptoAkCipherOptions
//
// The options that are available for all asymmetric key algorithms when creating a new QCryptoAkCipher.
type QCryptoAkCipherOptions struct {
	// Discriminator: alg

	// Alg encryption cipher algorithm
	Alg QCryptoAkCipherAlgorithm `json:"alg"`

	Rsa *QCryptoAkCipherOptionsRSA `json:"-"`
}

func (u QCryptoAkCipherOptions) MarshalJSON() ([]byte, error) {
	switch u.Alg {
	case "rsa":
		if u.Rsa == nil {
			return nil, fmt.Errorf("expected Rsa to be set")
		}

		return json.Marshal(struct {
			Alg QCryptoAkCipherAlgorithm `json:"alg"`
			*QCryptoAkCipherOptionsRSA
		}{
			Alg:                       u.Alg,
			QCryptoAkCipherOptionsRSA: u.Rsa,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// JobType Type of a background job.
type JobType string

const (
	// JobTypeCommit block commit job type, see "block-commit"
	JobTypeCommit JobType = "commit"
	// JobTypeStream block stream job type, see "block-stream"
	JobTypeStream JobType = "stream"
	// JobTypeMirror drive mirror job type, see "drive-mirror"
	JobTypeMirror JobType = "mirror"
	// JobTypeBackup drive backup job type, see "drive-backup"
	JobTypeBackup JobType = "backup"
	// JobTypeCreate image creation job type, see "blockdev-create" (since 3.0)
	JobTypeCreate JobType = "create"
	// JobTypeAmend image options amend job type, see "x-blockdev-amend" (since 5.1)
	JobTypeAmend JobType = "amend"
	// JobTypeSnapshotLoad snapshot load job type, see "snapshot-load" (since 6.0)
	JobTypeSnapshotLoad JobType = "snapshot-load"
	// JobTypeSnapshotSave snapshot save job type, see "snapshot-save" (since 6.0)
	JobTypeSnapshotSave JobType = "snapshot-save"
	// JobTypeSnapshotDelete snapshot delete job type, see "snapshot-delete" (since 6.0)
	JobTypeSnapshotDelete JobType = "snapshot-delete"
)

// JobStatus Indicates the present state of a given job in its lifetime.
type JobStatus string

const (
	// JobStatusUndefined Erroneous, default state. Should not ever be visible.
	JobStatusUndefined JobStatus = "undefined"
	// JobStatusCreated The job has been created, but not yet started.
	JobStatusCreated JobStatus = "created"
	// JobStatusRunning The job is currently running.
	JobStatusRunning JobStatus = "running"
	// JobStatusPaused The job is running, but paused. The pause may be requested by either the QMP user or by internal processes.
	JobStatusPaused JobStatus = "paused"
	// JobStatusReady The job is running, but is ready for the user to signal completion. This is used for long-running jobs like mirror that are designed to run indefinitely.
	JobStatusReady JobStatus = "ready"
	// JobStatusStandby The job is ready, but paused. This is nearly identical to @paused. The job may return to @ready or otherwise be canceled.
	JobStatusStandby JobStatus = "standby"
	// JobStatusWaiting The job is waiting for other jobs in the transaction to converge to the waiting state. This status will likely not be visible for the last job in a transaction.
	JobStatusWaiting JobStatus = "waiting"
	// JobStatusPending The job has finished its work, but has finalization steps that it needs to make prior to completing. These changes will require manual intervention via @job-finalize if auto-finalize was set to false. These pending changes may still fail.
	JobStatusPending JobStatus = "pending"
	// JobStatusAborting The job is in the process of being aborted, and will finish with an error. The job will afterwards report that it is @concluded. This status may not be visible to the management process.
	JobStatusAborting JobStatus = "aborting"
	// JobStatusConcluded The job has finished all work. If auto-dismiss was set to false, the job will remain in the query list until it is dismissed via @job-dismiss.
	JobStatusConcluded JobStatus = "concluded"
	// JobStatusNull The job is in the process of being dismantled. This state should not ever be visible externally.
	JobStatusNull JobStatus = "null"
)

// JobVerb Represents command verbs that can be applied to a job.
type JobVerb string

const (
	// JobVerbCancel see @job-cancel
	JobVerbCancel JobVerb = "cancel"
	// JobVerbPause see @job-pause
	JobVerbPause JobVerb = "pause"
	// JobVerbResume see @job-resume
	JobVerbResume JobVerb = "resume"
	// JobVerbSetSpeed see @block-job-set-speed
	JobVerbSetSpeed JobVerb = "set-speed"
	// JobVerbComplete see @job-complete
	JobVerbComplete JobVerb = "complete"
	// JobVerbDismiss see @job-dismiss
	JobVerbDismiss JobVerb = "dismiss"
	// JobVerbFinalize see @job-finalize
	JobVerbFinalize JobVerb = "finalize"
	// JobVerbChange see @block-job-change (since 8.2)
	JobVerbChange JobVerb = "change"
)

// JobStatusChangeEvent (JOB_STATUS_CHANGE)
//
// Emitted when a job transitions to a different status.
type JobStatusChangeEvent struct {
	// Id The job identifier
	Id string `json:"id"`
	// Status The new job status
	Status JobStatus `json:"status"`
}

func (JobStatusChangeEvent) Event() string {
	return "JOB_STATUS_CHANGE"
}

// JobPause
//
// Pause an active job. This command returns immediately after marking the active job for pausing. Pausing an already paused job is an error. The job will pause as soon as possible, which means transitioning into the PAUSED state if it was RUNNING, or into STANDBY if it was READY. The corresponding JOB_STATUS_CHANGE event will be emitted. Cancelling a paused job automatically resumes it.
type JobPause struct {
	// Id The job identifier.
	Id string `json:"id"`
}

func (JobPause) Command() string {
	return "job-pause"
}

func (cmd JobPause) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "job-pause", cmd, nil)
}

// JobResume
//
// Resume a paused job. This command returns immediately after resuming a paused job. Resuming an already running job is an error.
type JobResume struct {
	// Id The job identifier.
	Id string `json:"id"`
}

func (JobResume) Command() string {
	return "job-resume"
}

func (cmd JobResume) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "job-resume", cmd, nil)
}

// JobCancel
//
// Instruct an active background job to cancel at the next opportunity. This command returns immediately after marking the active job for cancellation. The job will cancel as soon as possible and then emit a JOB_STATUS_CHANGE event. Usually, the status will change to ABORTING, but it is possible that a job successfully completes (e.g. because it was almost done and there was no opportunity to cancel earlier than completing the job) and transitions to PENDING instead.
type JobCancel struct {
	// Id The job identifier.
	Id string `json:"id"`
}

func (JobCancel) Command() string {
	return "job-cancel"
}

func (cmd JobCancel) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "job-cancel", cmd, nil)
}

// JobComplete
//
// Manually trigger completion of an active job in the READY state.
type JobComplete struct {
	// Id The job identifier.
	Id string `json:"id"`
}

func (JobComplete) Command() string {
	return "job-complete"
}

func (cmd JobComplete) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "job-complete", cmd, nil)
}

// JobDismiss
//
// Deletes a job that is in the CONCLUDED state. This command only needs to be run explicitly for jobs that don't have automatic dismiss enabled. This command will refuse to operate on any job that has not yet reached its terminal state, JOB_STATUS_CONCLUDED. For jobs that make use of JOB_READY event, job-cancel or job-complete will still need to be used as appropriate.
type JobDismiss struct {
	// Id The job identifier.
	Id string `json:"id"`
}

func (JobDismiss) Command() string {
	return "job-dismiss"
}

func (cmd JobDismiss) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "job-dismiss", cmd, nil)
}

// JobFinalize
//
// Instructs all jobs in a transaction (or a single job if it is not part of any transaction) to finalize any graph changes and do any necessary cleanup. This command requires that all involved jobs are in the PENDING state. For jobs in a transaction, instructing one job to finalize will force ALL jobs in the transaction to finalize, so it is only necessary to instruct a single member job to finalize.
type JobFinalize struct {
	// Id The identifier of any job in the transaction, or of a job that is not part of any transaction.
	Id string `json:"id"`
}

func (JobFinalize) Command() string {
	return "job-finalize"
}

func (cmd JobFinalize) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "job-finalize", cmd, nil)
}

// JobInfo
//
// Information about a job.
type JobInfo struct {
	// Id The job identifier
	Id string `json:"id"`
	// Type The kind of job that is being performed
	Type JobType `json:"type"`
	// Status Current job state/status
	Status JobStatus `json:"status"`
	// CurrentProgress Progress made until now. The unit is arbitrary and the value can only meaningfully be used for the ratio of @current-progress to @total-progress. The value is monotonically increasing.
	CurrentProgress int64 `json:"current-progress"`
	// TotalProgress Estimated @current-progress value at the completion of the job. This value can arbitrarily change while the job is running, in both directions.
	TotalProgress int64 `json:"total-progress"`
	// Error If this field is present, the job failed; if it is still missing in the CONCLUDED state, this indicates successful completion. The value is a human-readable error message to describe the reason for the job failure. It should not be parsed by applications.
	Error *string `json:"error,omitempty"`
}

// QueryJobs
//
// Return information about jobs.
type QueryJobs struct {
}

func (QueryJobs) Command() string {
	return "query-jobs"
}

func (cmd QueryJobs) Execute(ctx context.Context, client api.Client) ([]JobInfo, error) {
	var ret []JobInfo

	return ret, client.Execute(ctx, "query-jobs", cmd, &ret)
}

// SnapshotInfo
type SnapshotInfo struct {
	// Id unique snapshot id
	Id string `json:"id"`
	// Name user chosen name
	Name string `json:"name"`
	// VmStateSize size of the VM state
	VmStateSize int64 `json:"vm-state-size"`
	// DateSec UTC date of the snapshot in seconds
	DateSec int64 `json:"date-sec"`
	// DateNsec fractional part in nano seconds to be used with date-sec
	DateNsec int64 `json:"date-nsec"`
	// VmClockSec VM clock relative to boot in seconds
	VmClockSec int64 `json:"vm-clock-sec"`
	// VmClockNsec fractional part in nano seconds to be used with vm-clock-sec
	VmClockNsec int64 `json:"vm-clock-nsec"`
	// Icount Current instruction count. Appears when execution record/replay is enabled. Used for "time-traveling" to match the moment in the recorded execution with the snapshots. This counter may be obtained through @query-replay command (since 5.2)
	Icount *int64 `json:"icount,omitempty"`
}

// ImageInfoSpecificQCow2EncryptionBase
type ImageInfoSpecificQCow2EncryptionBase struct {
	// Format The encryption format
	Format BlockdevQcow2EncryptionFormat `json:"format"`
}

// ImageInfoSpecificQCow2Encryption
type ImageInfoSpecificQCow2Encryption struct {
	// Discriminator: format

	ImageInfoSpecificQCow2EncryptionBase

	Luks *QCryptoBlockInfoLUKS `json:"-"`
}

func (u ImageInfoSpecificQCow2Encryption) MarshalJSON() ([]byte, error) {
	switch u.Format {
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			ImageInfoSpecificQCow2EncryptionBase
			*QCryptoBlockInfoLUKS
		}{
			ImageInfoSpecificQCow2EncryptionBase: u.ImageInfoSpecificQCow2EncryptionBase,
			QCryptoBlockInfoLUKS:                 u.Luks,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// ImageInfoSpecificQCow2
type ImageInfoSpecificQCow2 struct {
	// Compat compatibility level
	Compat string `json:"compat"`
	// DataFile the filename of the external data file that is stored in the image and used as a default for opening the image
	DataFile *string `json:"data-file,omitempty"`
	// DataFileRaw True if the external data file must stay valid as a standalone (read-only) raw image without looking at qcow2
	DataFileRaw *bool `json:"data-file-raw,omitempty"`
	// ExtendedL2 true if the image has extended L2 entries; only valid for compat >= 1.1 (since 5.2)
	ExtendedL2 *bool `json:"extended-l2,omitempty"`
	// LazyRefcounts on or off; only valid for compat >= 1.1
	LazyRefcounts *bool `json:"lazy-refcounts,omitempty"`
	// Corrupt true if the image has been marked corrupt; only valid for compat >= 1.1 (since 2.2)
	Corrupt *bool `json:"corrupt,omitempty"`
	// RefcountBits width of a refcount entry in bits (since 2.3)
	RefcountBits int64 `json:"refcount-bits"`
	// Encrypt details about encryption parameters; only set if image is encrypted (since 2.10)
	Encrypt *ImageInfoSpecificQCow2Encryption `json:"encrypt,omitempty"`
	// Bitmaps A list of qcow2 bitmap details (since 4.0)
	Bitmaps []Qcow2BitmapInfo `json:"bitmaps,omitempty"`
	// CompressionType the image cluster compression method (since 5.1)
	CompressionType Qcow2CompressionType `json:"compression-type"`
}

// ImageInfoSpecificVmdk
type ImageInfoSpecificVmdk struct {
	// CreateType The create type of VMDK image
	CreateType string `json:"create-type"`
	// Cid Content id of image
	Cid int64 `json:"cid"`
	// ParentCid Parent VMDK image's cid
	ParentCid int64 `json:"parent-cid"`
	// Extents List of extent files
	Extents []VmdkExtentInfo `json:"extents"`
}

// VmdkExtentInfo
//
// Information about a VMDK extent file
type VmdkExtentInfo struct {
	// Filename Name of the extent file
	Filename string `json:"filename"`
	// Format Extent type (e.g. FLAT or SPARSE)
	Format string `json:"format"`
	// VirtualSize Number of bytes covered by this extent
	VirtualSize int64 `json:"virtual-size"`
	// ClusterSize Cluster size in bytes (for non-flat extents)
	ClusterSize *int64 `json:"cluster-size,omitempty"`
	// Compressed Whether this extent contains compressed data
	Compressed *bool `json:"compressed,omitempty"`
}

// ImageInfoSpecificRbd
type ImageInfoSpecificRbd struct {
	// EncryptionFormat Image encryption format
	EncryptionFormat *RbdImageEncryptionFormat `json:"encryption-format,omitempty"`
}

// ImageInfoSpecificFile
type ImageInfoSpecificFile struct {
	// ExtentSizeHint Extent size hint (if available)
	ExtentSizeHint *uint64 `json:"extent-size-hint,omitempty"`
}

// ImageInfoSpecificKind
type ImageInfoSpecificKind string

const (
	ImageInfoSpecificKindQcow2 ImageInfoSpecificKind = "qcow2"
	ImageInfoSpecificKindVmdk  ImageInfoSpecificKind = "vmdk"
	// ImageInfoSpecificKindLuks Since 2.7
	ImageInfoSpecificKindLuks ImageInfoSpecificKind = "luks"
	// ImageInfoSpecificKindRbd Since 6.1
	ImageInfoSpecificKindRbd ImageInfoSpecificKind = "rbd"
	// ImageInfoSpecificKindFile Since 8.0
	ImageInfoSpecificKindFile ImageInfoSpecificKind = "file"
)

// ImageInfoSpecificQCow2Wrapper
type ImageInfoSpecificQCow2Wrapper struct {
	// Data image information specific to QCOW2
	Data ImageInfoSpecificQCow2 `json:"data"`
}

// ImageInfoSpecificVmdkWrapper
type ImageInfoSpecificVmdkWrapper struct {
	// Data image information specific to VMDK
	Data ImageInfoSpecificVmdk `json:"data"`
}

// ImageInfoSpecificLUKSWrapper
type ImageInfoSpecificLUKSWrapper struct {
	// Data image information specific to LUKS
	Data QCryptoBlockInfoLUKS `json:"data"`
}

// ImageInfoSpecificRbdWrapper
type ImageInfoSpecificRbdWrapper struct {
	// Data image information specific to RBD
	Data ImageInfoSpecificRbd `json:"data"`
}

// ImageInfoSpecificFileWrapper
type ImageInfoSpecificFileWrapper struct {
	// Data image information specific to files
	Data ImageInfoSpecificFile `json:"data"`
}

// ImageInfoSpecific
//
// A discriminated record of image format specific information structures.
type ImageInfoSpecific struct {
	// Discriminator: type

	// Type block driver name
	Type ImageInfoSpecificKind `json:"type"`

	Qcow2 *ImageInfoSpecificQCow2Wrapper `json:"-"`
	Vmdk  *ImageInfoSpecificVmdkWrapper  `json:"-"`
	Luks  *ImageInfoSpecificLUKSWrapper  `json:"-"`
	Rbd   *ImageInfoSpecificRbdWrapper   `json:"-"`
	File  *ImageInfoSpecificFileWrapper  `json:"-"`
}

func (u ImageInfoSpecific) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "qcow2":
		if u.Qcow2 == nil {
			return nil, fmt.Errorf("expected Qcow2 to be set")
		}

		return json.Marshal(struct {
			Type ImageInfoSpecificKind `json:"type"`
			*ImageInfoSpecificQCow2Wrapper
		}{
			Type:                          u.Type,
			ImageInfoSpecificQCow2Wrapper: u.Qcow2,
		})
	case "vmdk":
		if u.Vmdk == nil {
			return nil, fmt.Errorf("expected Vmdk to be set")
		}

		return json.Marshal(struct {
			Type ImageInfoSpecificKind `json:"type"`
			*ImageInfoSpecificVmdkWrapper
		}{
			Type:                         u.Type,
			ImageInfoSpecificVmdkWrapper: u.Vmdk,
		})
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			Type ImageInfoSpecificKind `json:"type"`
			*ImageInfoSpecificLUKSWrapper
		}{
			Type:                         u.Type,
			ImageInfoSpecificLUKSWrapper: u.Luks,
		})
	case "rbd":
		if u.Rbd == nil {
			return nil, fmt.Errorf("expected Rbd to be set")
		}

		return json.Marshal(struct {
			Type ImageInfoSpecificKind `json:"type"`
			*ImageInfoSpecificRbdWrapper
		}{
			Type:                        u.Type,
			ImageInfoSpecificRbdWrapper: u.Rbd,
		})
	case "file":
		if u.File == nil {
			return nil, fmt.Errorf("expected File to be set")
		}

		return json.Marshal(struct {
			Type ImageInfoSpecificKind `json:"type"`
			*ImageInfoSpecificFileWrapper
		}{
			Type:                         u.Type,
			ImageInfoSpecificFileWrapper: u.File,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockNodeInfo
//
// Information about a QEMU image file
type BlockNodeInfo struct {
	// Filename name of the image file
	Filename string `json:"filename"`
	// Format format of the image file
	Format string `json:"format"`
	// DirtyFlag true if image is not cleanly closed
	DirtyFlag *bool `json:"dirty-flag,omitempty"`
	// ActualSize actual size on disk in bytes of the image
	ActualSize *int64 `json:"actual-size,omitempty"`
	// VirtualSize maximum capacity in bytes of the image
	VirtualSize int64 `json:"virtual-size"`
	// ClusterSize size of a cluster in bytes
	ClusterSize *int64 `json:"cluster-size,omitempty"`
	// Encrypted true if the image is encrypted
	Encrypted *bool `json:"encrypted,omitempty"`
	// Compressed true if the image is compressed (Since 1.7)
	Compressed *bool `json:"compressed,omitempty"`
	// BackingFilename name of the backing file
	BackingFilename *string `json:"backing-filename,omitempty"`
	// FullBackingFilename full path of the backing file
	FullBackingFilename *string `json:"full-backing-filename,omitempty"`
	// BackingFilenameFormat the format of the backing file
	BackingFilenameFormat *string `json:"backing-filename-format,omitempty"`
	// Snapshots list of VM snapshots
	Snapshots []SnapshotInfo `json:"snapshots,omitempty"`
	// FormatSpecific structure supplying additional format-specific information (since 1.7)
	FormatSpecific *ImageInfoSpecific `json:"format-specific,omitempty"`
}

// ImageInfo
//
// Information about a QEMU image file, and potentially its backing image
type ImageInfo struct {
	BlockNodeInfo

	// BackingImage info of the backing image
	BackingImage *ImageInfo `json:"backing-image,omitempty"`
}

// BlockChildInfo
//
// Information about all nodes in the block graph starting at some node, annotated with information about that node in relation to its parent.
type BlockChildInfo struct {
	// Name Child name of the root node in the BlockGraphInfo struct, in its role as the child of some undescribed parent node
	Name string `json:"name"`
	// Info Block graph information starting at this node
	Info BlockGraphInfo `json:"info"`
}

// BlockGraphInfo
//
// Information about all nodes in a block (sub)graph in the form of BlockNodeInfo data. The base BlockNodeInfo struct contains the information for the (sub)graph's root node.
type BlockGraphInfo struct {
	BlockNodeInfo

	// Children Array of links to this node's child nodes' information
	Children []BlockChildInfo `json:"children"`
}

// ImageCheck
//
// Information about a QEMU image file check
type ImageCheck struct {
	// Filename name of the image file checked
	Filename string `json:"filename"`
	// Format format of the image file checked
	Format string `json:"format"`
	// CheckErrors number of unexpected errors occurred during check
	CheckErrors int64 `json:"check-errors"`
	// ImageEndOffset offset (in bytes) where the image ends, this field is present if the driver for the image format supports it
	ImageEndOffset *int64 `json:"image-end-offset,omitempty"`
	// Corruptions number of corruptions found during the check if any
	Corruptions *int64 `json:"corruptions,omitempty"`
	// Leaks number of leaks found during the check if any
	Leaks *int64 `json:"leaks,omitempty"`
	// CorruptionsFixed number of corruptions fixed during the check if any
	CorruptionsFixed *int64 `json:"corruptions-fixed,omitempty"`
	// LeaksFixed number of leaks fixed during the check if any
	LeaksFixed *int64 `json:"leaks-fixed,omitempty"`
	// TotalClusters total number of clusters, this field is present if the driver for the image format supports it
	TotalClusters *int64 `json:"total-clusters,omitempty"`
	// AllocatedClusters total number of allocated clusters, this field is present if the driver for the image format supports it
	AllocatedClusters *int64 `json:"allocated-clusters,omitempty"`
	// FragmentedClusters total number of fragmented clusters, this field is present if the driver for the image format supports it
	FragmentedClusters *int64 `json:"fragmented-clusters,omitempty"`
	// CompressedClusters total number of compressed clusters, this field is present if the driver for the image format supports it
	CompressedClusters *int64 `json:"compressed-clusters,omitempty"`
}

// MapEntry
//
// Mapping information from a virtual block range to a host file range
type MapEntry struct {
	// Start virtual (guest) offset of the first byte described by this entry
	Start int64 `json:"start"`
	// Length the number of bytes of the mapped virtual range
	Length int64 `json:"length"`
	// Data reading the image will actually read data from a file (in particular, if @offset is present this means that the sectors are not simply preallocated, but contain actual data in raw format)
	Data bool `json:"data"`
	// Zero whether the virtual blocks read as zeroes
	Zero bool `json:"zero"`
	// Compressed true if the data is stored compressed (since 8.2)
	Compressed bool `json:"compressed"`
	// Depth number of layers (0 = top image, 1 = top image's backing file, ..., n - 1 = bottom image (where n is the number of images in the chain)) before reaching one for which the range is allocated
	Depth int64 `json:"depth"`
	// Present true if this layer provides the data, false if adding a backing layer could impact this region (since 6.1)
	Present bool `json:"present"`
	// Offset if present, the image file stores the data for this range in raw format at the given (host) offset
	Offset *int64 `json:"offset,omitempty"`
	// Filename filename that is referred to by @offset
	Filename *string `json:"filename,omitempty"`
}

// BlockdevCacheInfo
//
// Cache mode information for a block device
type BlockdevCacheInfo struct {
	// Writeback true if writeback mode is enabled
	Writeback bool `json:"writeback"`
	// Direct true if the host page cache is bypassed (O_DIRECT)
	Direct bool `json:"direct"`
	// NoFlush true if flush requests are ignored for the device
	NoFlush bool `json:"no-flush"`
}

// BlockDeviceInfo
//
// Information about the backing device for a block device.
type BlockDeviceInfo struct {
	// File the filename of the backing device
	File string `json:"file"`
	// NodeName the name of the block driver node (Since 2.0)
	NodeName *string `json:"node-name,omitempty"`
	// Ro true if the backing device was open read-only
	Ro bool `json:"ro"`
	// Drv the name of the block format used to open the backing device.
	Drv string `json:"drv"`
	// BackingFile the name of the backing file (for copy-on-write)
	BackingFile *string `json:"backing_file,omitempty"`
	// BackingFileDepth number of files in the backing file chain
	BackingFileDepth int64 `json:"backing_file_depth"`
	// Encrypted true if the backing device is encrypted
	Encrypted bool `json:"encrypted"`
	// DetectZeroes detect and optimize zero writes (Since 2.1)
	DetectZeroes BlockdevDetectZeroesOptions `json:"detect_zeroes"`
	// Bps total throughput limit in bytes per second is specified
	Bps int64 `json:"bps"`
	// BpsRd read throughput limit in bytes per second is specified
	BpsRd int64 `json:"bps_rd"`
	// BpsWr write throughput limit in bytes per second is specified
	BpsWr int64 `json:"bps_wr"`
	// Iops total I/O operations per second is specified
	Iops int64 `json:"iops"`
	// IopsRd read I/O operations per second is specified
	IopsRd int64 `json:"iops_rd"`
	// IopsWr write I/O operations per second is specified
	IopsWr int64 `json:"iops_wr"`
	// Image the info of image used (since: 1.6)
	Image ImageInfo `json:"image"`
	// BpsMax total throughput limit during bursts, in bytes (Since 1.7)
	BpsMax *int64 `json:"bps_max,omitempty"`
	// BpsRdMax read throughput limit during bursts, in bytes (Since 1.7)
	BpsRdMax *int64 `json:"bps_rd_max,omitempty"`
	// BpsWrMax write throughput limit during bursts, in bytes (Since 1.7)
	BpsWrMax *int64 `json:"bps_wr_max,omitempty"`
	// IopsMax total I/O operations per second during bursts, in bytes (Since 1.7)
	IopsMax *int64 `json:"iops_max,omitempty"`
	// IopsRdMax read I/O operations per second during bursts, in bytes (Since 1.7)
	IopsRdMax *int64 `json:"iops_rd_max,omitempty"`
	// IopsWrMax write I/O operations per second during bursts, in bytes (Since 1.7)
	IopsWrMax *int64 `json:"iops_wr_max,omitempty"`
	// BpsMaxLength maximum length of the @bps_max burst period, in seconds. (Since 2.6)
	BpsMaxLength *int64 `json:"bps_max_length,omitempty"`
	// BpsRdMaxLength maximum length of the @bps_rd_max burst period, in seconds. (Since 2.6)
	BpsRdMaxLength *int64 `json:"bps_rd_max_length,omitempty"`
	// BpsWrMaxLength maximum length of the @bps_wr_max burst period, in seconds. (Since 2.6)
	BpsWrMaxLength *int64 `json:"bps_wr_max_length,omitempty"`
	// IopsMaxLength maximum length of the @iops burst period, in seconds. (Since 2.6)
	IopsMaxLength *int64 `json:"iops_max_length,omitempty"`
	// IopsRdMaxLength maximum length of the @iops_rd_max burst period, in seconds. (Since 2.6)
	IopsRdMaxLength *int64 `json:"iops_rd_max_length,omitempty"`
	// IopsWrMaxLength maximum length of the @iops_wr_max burst period, in seconds. (Since 2.6)
	IopsWrMaxLength *int64 `json:"iops_wr_max_length,omitempty"`
	// IopsSize an I/O size in bytes (Since 1.7)
	IopsSize *int64 `json:"iops_size,omitempty"`
	// Group throttle group name (Since 2.4)
	Group *string `json:"group,omitempty"`
	// Cache the cache mode used for the block device (since: 2.3)
	Cache BlockdevCacheInfo `json:"cache"`
	// WriteThreshold configured write threshold for the device. 0 if disabled. (Since 2.3)
	WriteThreshold int64 `json:"write_threshold"`
	// DirtyBitmaps dirty bitmaps information (only present if node has one or more dirty bitmaps) (Since 4.2)
	DirtyBitmaps []BlockDirtyInfo `json:"dirty-bitmaps,omitempty"`
}

// BlockDeviceIoStatus An enumeration of block device I/O status.
type BlockDeviceIoStatus string

const (
	// BlockDeviceIoStatusOk The last I/O operation has succeeded
	BlockDeviceIoStatusOk BlockDeviceIoStatus = "ok"
	// BlockDeviceIoStatusFailed The last I/O operation has failed
	BlockDeviceIoStatusFailed BlockDeviceIoStatus = "failed"
	// BlockDeviceIoStatusNospace The last I/O operation has failed due to a no-space condition
	BlockDeviceIoStatusNospace BlockDeviceIoStatus = "nospace"
)

// BlockDirtyInfo
//
// Block dirty bitmap information.
type BlockDirtyInfo struct {
	// Name the name of the dirty bitmap (Since 2.4)
	Name *string `json:"name,omitempty"`
	// Count number of dirty bytes according to the dirty bitmap
	Count int64 `json:"count"`
	// Granularity granularity of the dirty bitmap in bytes (since 1.4)
	Granularity uint32 `json:"granularity"`
	// Recording true if the bitmap is recording new writes from the guest. (since 4.0)
	Recording bool `json:"recording"`
	// Busy true if the bitmap is in-use by some operation (NBD or jobs) and cannot be modified via QMP or used by another operation. (since 4.0)
	Busy bool `json:"busy"`
	// Persistent true if the bitmap was stored on disk, is scheduled to be stored on disk, or both. (since 4.0)
	Persistent bool `json:"persistent"`
	// Inconsistent true if this is a persistent bitmap that was improperly stored. Implies @persistent to be true; @recording and @busy to be false. This bitmap cannot be used. To remove it, use @block-dirty-bitmap-remove. (Since 4.0)
	Inconsistent *bool `json:"inconsistent,omitempty"`
}

// Qcow2BitmapInfoFlags An enumeration of flags that a bitmap can report to the user.
type Qcow2BitmapInfoFlags string

const (
	// Qcow2BitmapInfoFlagsInUse This flag is set by any process actively modifying the qcow2 file, and cleared when the updated bitmap is flushed to the qcow2 image. The presence of this flag in an offline image means that the bitmap was not saved correctly after its last usage, and may contain inconsistent data.
	Qcow2BitmapInfoFlagsInUse Qcow2BitmapInfoFlags = "in-use"
	// Qcow2BitmapInfoFlagsAuto The bitmap must reflect all changes of the virtual disk by any application that would write to this qcow2 file.
	Qcow2BitmapInfoFlagsAuto Qcow2BitmapInfoFlags = "auto"
)

// Qcow2BitmapInfo
//
// Qcow2 bitmap information.
type Qcow2BitmapInfo struct {
	// Name the name of the bitmap
	Name string `json:"name"`
	// Granularity granularity of the bitmap in bytes
	Granularity uint32 `json:"granularity"`
	// Flags flags of the bitmap
	Flags []Qcow2BitmapInfoFlags `json:"flags"`
}

// BlockLatencyHistogramInfo
//
// Block latency histogram.
type BlockLatencyHistogramInfo struct {
	// Boundaries list of interval boundary values in nanoseconds, all greater than zero and in ascending order. For example, the list
	Boundaries []uint64 `json:"boundaries"`
	// Bins list of io request counts corresponding to histogram intervals, one more element than @boundaries has. For the example above, @bins may be something like [3, 1, 5, 2], and
	Bins []uint64 `json:"bins"`
}

// BlockInfo
//
// Block device information. This structure describes a virtual device and the backing device associated with it.
type BlockInfo struct {
	// Device The device name associated with the virtual device.
	Device string `json:"device"`
	// Qdev The qdev ID, or if no ID is assigned, the QOM path of the block device. (since 2.10)
	Qdev *string `json:"qdev,omitempty"`
	// Type This field is returned only for compatibility reasons, it should not be used (always returns 'unknown')
	Type string `json:"type"`
	// Removable True if the device supports removable media.
	Removable bool `json:"removable"`
	// Locked True if the guest has locked this device from having its media removed
	Locked bool `json:"locked"`
	// Inserted @BlockDeviceInfo describing the device if media is present
	Inserted *BlockDeviceInfo `json:"inserted,omitempty"`
	// TrayOpen True if the device's tray is open (only present if it has a tray)
	TrayOpen *bool `json:"tray_open,omitempty"`
	// IoStatus @BlockDeviceIoStatus. Only present if the device supports it and the VM is configured to stop on errors
	IoStatus *BlockDeviceIoStatus `json:"io-status,omitempty"`
}

// BlockMeasureInfo
//
// Image file size calculation information. This structure describes the size requirements for creating a new image file. The size requirements depend on the new image file format. File size always equals virtual disk size for the 'raw' format, even for sparse POSIX files. Compact formats such as 'qcow2' represent unallocated and zero regions efficiently so file size may be smaller than virtual disk size. The values are upper bounds that are guaranteed to fit the new image file. Subsequent modification, such as internal snapshot or further bitmap creation, may require additional space and is not covered here.
type BlockMeasureInfo struct {
	// Required Size required for a new image file, in bytes, when copying just allocated guest-visible contents.
	Required int64 `json:"required"`
	// FullyAllocated Image file size, in bytes, once data has been written to all sectors, when copying just guest-visible contents.
	FullyAllocated int64 `json:"fully-allocated"`
	// Bitmaps Additional size required if all the top-level bitmap metadata in the source image were to be copied to the destination, present only when source and destination both support persistent bitmaps. (since 5.1)
	Bitmaps *int64 `json:"bitmaps,omitempty"`
}

// QueryBlock
//
// Get a list of BlockInfo for all virtual block devices.
type QueryBlock struct {
}

func (QueryBlock) Command() string {
	return "query-block"
}

func (cmd QueryBlock) Execute(ctx context.Context, client api.Client) ([]BlockInfo, error) {
	var ret []BlockInfo

	return ret, client.Execute(ctx, "query-block", cmd, &ret)
}

// BlockDeviceTimedStats
//
// Statistics of a block device during a given interval of time.
type BlockDeviceTimedStats struct {
	// IntervalLength Interval used for calculating the statistics, in seconds.
	IntervalLength int64 `json:"interval_length"`
	// MinRdLatencyNs Minimum latency of read operations in the defined interval, in nanoseconds.
	MinRdLatencyNs int64 `json:"min_rd_latency_ns"`
	// MaxRdLatencyNs Maximum latency of read operations in the defined interval, in nanoseconds.
	MaxRdLatencyNs int64 `json:"max_rd_latency_ns"`
	// AvgRdLatencyNs Average latency of read operations in the defined interval, in nanoseconds.
	AvgRdLatencyNs int64 `json:"avg_rd_latency_ns"`
	// MinWrLatencyNs Minimum latency of write operations in the defined interval, in nanoseconds.
	MinWrLatencyNs int64 `json:"min_wr_latency_ns"`
	// MaxWrLatencyNs Maximum latency of write operations in the defined interval, in nanoseconds.
	MaxWrLatencyNs int64 `json:"max_wr_latency_ns"`
	// AvgWrLatencyNs Average latency of write operations in the defined interval, in nanoseconds.
	AvgWrLatencyNs int64 `json:"avg_wr_latency_ns"`
	// MinZoneAppendLatencyNs Minimum latency of zone append operations in the defined interval, in nanoseconds (since 8.1)
	MinZoneAppendLatencyNs int64 `json:"min_zone_append_latency_ns"`
	// MaxZoneAppendLatencyNs Maximum latency of zone append operations in the defined interval, in nanoseconds (since 8.1)
	MaxZoneAppendLatencyNs int64 `json:"max_zone_append_latency_ns"`
	// AvgZoneAppendLatencyNs Average latency of zone append operations in the defined interval, in nanoseconds (since 8.1)
	AvgZoneAppendLatencyNs int64 `json:"avg_zone_append_latency_ns"`
	// MinFlushLatencyNs Minimum latency of flush operations in the defined interval, in nanoseconds.
	MinFlushLatencyNs int64 `json:"min_flush_latency_ns"`
	// MaxFlushLatencyNs Maximum latency of flush operations in the defined interval, in nanoseconds.
	MaxFlushLatencyNs int64 `json:"max_flush_latency_ns"`
	// AvgFlushLatencyNs Average latency of flush operations in the defined interval, in nanoseconds.
	AvgFlushLatencyNs int64 `json:"avg_flush_latency_ns"`
	// AvgRdQueueDepth Average number of pending read operations in the defined interval.
	AvgRdQueueDepth float64 `json:"avg_rd_queue_depth"`
	// AvgWrQueueDepth Average number of pending write operations in the defined interval.
	AvgWrQueueDepth float64 `json:"avg_wr_queue_depth"`
	// AvgZoneAppendQueueDepth Average number of pending zone append operations in the defined interval (since 8.1).
	AvgZoneAppendQueueDepth float64 `json:"avg_zone_append_queue_depth"`
}

// BlockDeviceStats
//
// Statistics of a virtual block device or a block backing device.
type BlockDeviceStats struct {
	// RdBytes The number of bytes read by the device.
	RdBytes int64 `json:"rd_bytes"`
	// WrBytes The number of bytes written by the device.
	WrBytes int64 `json:"wr_bytes"`
	// ZoneAppendBytes The number of bytes appended by the zoned devices (since 8.1)
	ZoneAppendBytes int64 `json:"zone_append_bytes"`
	// UnmapBytes The number of bytes unmapped by the device (Since 4.2)
	UnmapBytes int64 `json:"unmap_bytes"`
	// RdOperations The number of read operations performed by the device.
	RdOperations int64 `json:"rd_operations"`
	// WrOperations The number of write operations performed by the device.
	WrOperations int64 `json:"wr_operations"`
	// ZoneAppendOperations The number of zone append operations performed by the zoned devices (since 8.1)
	ZoneAppendOperations int64 `json:"zone_append_operations"`
	// FlushOperations The number of cache flush operations performed by the device (since 0.15)
	FlushOperations int64 `json:"flush_operations"`
	// UnmapOperations The number of unmap operations performed by the device (Since 4.2)
	UnmapOperations int64 `json:"unmap_operations"`
	// RdTotalTimeNs Total time spent on reads in nanoseconds (since 0.15).
	RdTotalTimeNs int64 `json:"rd_total_time_ns"`
	// WrTotalTimeNs Total time spent on writes in nanoseconds (since 0.15).
	WrTotalTimeNs int64 `json:"wr_total_time_ns"`
	// ZoneAppendTotalTimeNs Total time spent on zone append writes in nanoseconds (since 8.1)
	ZoneAppendTotalTimeNs int64 `json:"zone_append_total_time_ns"`
	// FlushTotalTimeNs Total time spent on cache flushes in nanoseconds (since 0.15).
	FlushTotalTimeNs int64 `json:"flush_total_time_ns"`
	// UnmapTotalTimeNs Total time spent on unmap operations in nanoseconds (Since 4.2)
	UnmapTotalTimeNs int64 `json:"unmap_total_time_ns"`
	// WrHighestOffset The offset after the greatest byte written to the device. The intended use of this information is for growable sparse files (like qcow2) that are used on top of a physical device.
	WrHighestOffset int64 `json:"wr_highest_offset"`
	// RdMerged Number of read requests that have been merged into another request (Since 2.3).
	RdMerged int64 `json:"rd_merged"`
	// WrMerged Number of write requests that have been merged into another request (Since 2.3).
	WrMerged int64 `json:"wr_merged"`
	// ZoneAppendMerged Number of zone append requests that have been merged into another request (since 8.1)
	ZoneAppendMerged int64 `json:"zone_append_merged"`
	// UnmapMerged Number of unmap requests that have been merged into another request (Since 4.2)
	UnmapMerged int64 `json:"unmap_merged"`
	// IdleTimeNs Time since the last I/O operation, in nanoseconds. If the field is absent it means that there haven't been any operations yet (Since 2.5).
	IdleTimeNs *int64 `json:"idle_time_ns,omitempty"`
	// FailedRdOperations The number of failed read operations performed by the device (Since 2.5)
	FailedRdOperations int64 `json:"failed_rd_operations"`
	// FailedWrOperations The number of failed write operations performed by the device (Since 2.5)
	FailedWrOperations int64 `json:"failed_wr_operations"`
	// FailedZoneAppendOperations The number of failed zone append write operations performed by the zoned devices (since 8.1)
	FailedZoneAppendOperations int64 `json:"failed_zone_append_operations"`
	// FailedFlushOperations The number of failed flush operations performed by the device (Since 2.5)
	FailedFlushOperations int64 `json:"failed_flush_operations"`
	// FailedUnmapOperations The number of failed unmap operations performed by the device (Since 4.2)
	FailedUnmapOperations int64 `json:"failed_unmap_operations"`
	// InvalidRdOperations The number of invalid read operations performed by the device (Since 2.5)
	InvalidRdOperations int64 `json:"invalid_rd_operations"`
	// InvalidWrOperations The number of invalid write operations performed by the device (Since 2.5)
	InvalidWrOperations int64 `json:"invalid_wr_operations"`
	// InvalidZoneAppendOperations The number of invalid zone append operations performed by the zoned device (since 8.1)
	InvalidZoneAppendOperations int64 `json:"invalid_zone_append_operations"`
	// InvalidFlushOperations The number of invalid flush operations performed by the device (Since 2.5)
	InvalidFlushOperations int64 `json:"invalid_flush_operations"`
	// InvalidUnmapOperations The number of invalid unmap operations performed by the device (Since 4.2)
	InvalidUnmapOperations int64 `json:"invalid_unmap_operations"`
	// AccountInvalid Whether invalid operations are included in the last access statistics (Since 2.5)
	AccountInvalid bool `json:"account_invalid"`
	// AccountFailed Whether failed operations are included in the latency and last access statistics (Since 2.5)
	AccountFailed bool `json:"account_failed"`
	// TimedStats Statistics specific to the set of previously defined intervals of time (Since 2.5)
	TimedStats []BlockDeviceTimedStats `json:"timed_stats"`
	// RdLatencyHistogram @BlockLatencyHistogramInfo. (Since 4.0)
	RdLatencyHistogram *BlockLatencyHistogramInfo `json:"rd_latency_histogram,omitempty"`
	// WrLatencyHistogram @BlockLatencyHistogramInfo. (Since 4.0)
	WrLatencyHistogram *BlockLatencyHistogramInfo `json:"wr_latency_histogram,omitempty"`
	// ZoneAppendLatencyHistogram @BlockLatencyHistogramInfo. (since 8.1)
	ZoneAppendLatencyHistogram *BlockLatencyHistogramInfo `json:"zone_append_latency_histogram,omitempty"`
	// FlushLatencyHistogram @BlockLatencyHistogramInfo. (Since 4.0)
	FlushLatencyHistogram *BlockLatencyHistogramInfo `json:"flush_latency_histogram,omitempty"`
}

// BlockStatsSpecificFile
//
// File driver statistics
type BlockStatsSpecificFile struct {
	// DiscardNbOk The number of successful discard operations performed by the driver.
	DiscardNbOk uint64 `json:"discard-nb-ok"`
	// DiscardNbFailed The number of failed discard operations performed by the driver.
	DiscardNbFailed uint64 `json:"discard-nb-failed"`
	// DiscardBytesOk The number of bytes discarded by the driver.
	DiscardBytesOk uint64 `json:"discard-bytes-ok"`
}

// BlockStatsSpecificNvme
//
// NVMe driver statistics
type BlockStatsSpecificNvme struct {
	// CompletionErrors The number of completion errors.
	CompletionErrors uint64 `json:"completion-errors"`
	// AlignedAccesses The number of aligned accesses performed by the driver.
	AlignedAccesses uint64 `json:"aligned-accesses"`
	// UnalignedAccesses The number of unaligned accesses performed by the driver.
	UnalignedAccesses uint64 `json:"unaligned-accesses"`
}

// BlockStatsSpecific
//
// Block driver specific statistics
type BlockStatsSpecific struct {
	// Discriminator: driver

	// Driver block driver name
	Driver BlockdevDriver `json:"driver"`

	File       *BlockStatsSpecificFile `json:"-"`
	HostDevice *BlockStatsSpecificFile `json:"-"`
	Nvme       *BlockStatsSpecificNvme `json:"-"`
}

func (u BlockStatsSpecific) MarshalJSON() ([]byte, error) {
	switch u.Driver {
	case "file":
		if u.File == nil {
			return nil, fmt.Errorf("expected File to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockStatsSpecificFile
		}{
			Driver:                 u.Driver,
			BlockStatsSpecificFile: u.File,
		})
	case "host_device":
		if u.HostDevice == nil {
			return nil, fmt.Errorf("expected HostDevice to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockStatsSpecificFile
		}{
			Driver:                 u.Driver,
			BlockStatsSpecificFile: u.HostDevice,
		})
	case "nvme":
		if u.Nvme == nil {
			return nil, fmt.Errorf("expected Nvme to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockStatsSpecificNvme
		}{
			Driver:                 u.Driver,
			BlockStatsSpecificNvme: u.Nvme,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockStats
//
// Statistics of a virtual block device or a block backing device.
type BlockStats struct {
	// Device If the stats are for a virtual block device, the name corresponding to the virtual block device.
	Device *string `json:"device,omitempty"`
	// Qdev The qdev ID, or if no ID is assigned, the QOM path of the block device. (since 3.0)
	Qdev *string `json:"qdev,omitempty"`
	// NodeName The node name of the device. (Since 2.3)
	NodeName *string `json:"node-name,omitempty"`
	// Stats A @BlockDeviceStats for the device.
	Stats BlockDeviceStats `json:"stats"`
	// DriverSpecific Optional driver-specific stats. (Since 4.2)
	DriverSpecific *BlockStatsSpecific `json:"driver-specific,omitempty"`
	// Parent This describes the file block device if it has one. Contains recursively the statistics of the underlying protocol (e.g. the host file for a qcow2 image). If there is no underlying protocol, this field is omitted
	Parent *BlockStats `json:"parent,omitempty"`
	// Backing This describes the backing block device if it has one. (Since 2.0)
	Backing *BlockStats `json:"backing,omitempty"`
}

// QueryBlockstats
//
// Query the @BlockStats for all virtual block devices.
type QueryBlockstats struct {
	// QueryNodes If true, the command will query all the block nodes that have a node name, in a list which will include "parent" information, but not "backing". If false or omitted, the behavior is as before - query all the device backends, recursively including their "parent" and "backing". Filter nodes that were created implicitly are skipped over in this mode. (Since 2.3)
	QueryNodes *bool `json:"query-nodes,omitempty"`
}

func (QueryBlockstats) Command() string {
	return "query-blockstats"
}

func (cmd QueryBlockstats) Execute(ctx context.Context, client api.Client) ([]BlockStats, error) {
	var ret []BlockStats

	return ret, client.Execute(ctx, "query-blockstats", cmd, &ret)
}

// BlockdevOnError An enumeration of possible behaviors for errors on I/O operations. The exact meaning depends on whether the I/O was initiated by a guest or by a block job
type BlockdevOnError string

const (
	// BlockdevOnErrorReport for guest operations, report the error to the guest; for jobs, cancel the job
	BlockdevOnErrorReport BlockdevOnError = "report"
	// BlockdevOnErrorIgnore ignore the error, only report a QMP event (BLOCK_IO_ERROR or BLOCK_JOB_ERROR). The backup, mirror and commit block jobs retry the failing request later and may still complete successfully. The stream block job continues to stream and will complete with an error.
	BlockdevOnErrorIgnore BlockdevOnError = "ignore"
	// BlockdevOnErrorEnospc same as @stop on ENOSPC, same as @report otherwise.
	BlockdevOnErrorEnospc BlockdevOnError = "enospc"
	// BlockdevOnErrorStop for guest operations, stop the virtual machine; for jobs, pause the job
	BlockdevOnErrorStop BlockdevOnError = "stop"
	// BlockdevOnErrorAuto inherit the error handling policy of the backend (since: 2.7)
	BlockdevOnErrorAuto BlockdevOnError = "auto"
)

// MirrorSyncMode An enumeration of possible behaviors for the initial synchronization phase of storage mirroring.
type MirrorSyncMode string

const (
	// MirrorSyncModeTop copies data in the topmost image to the destination
	MirrorSyncModeTop MirrorSyncMode = "top"
	// MirrorSyncModeFull copies data from all images to the destination
	MirrorSyncModeFull MirrorSyncMode = "full"
	// MirrorSyncModeNone only copy data written from now on
	MirrorSyncModeNone MirrorSyncMode = "none"
	// MirrorSyncModeIncremental only copy data described by the dirty bitmap.
	MirrorSyncModeIncremental MirrorSyncMode = "incremental"
	// MirrorSyncModeBitmap only copy data described by the dirty bitmap. (since: 4.2) Behavior on completion is determined by the BitmapSyncMode.
	MirrorSyncModeBitmap MirrorSyncMode = "bitmap"
)

// BitmapSyncMode An enumeration of possible behaviors for the synchronization of a bitmap when used for data copy operations.
type BitmapSyncMode string

const (
	// BitmapSyncModeOnSuccess The bitmap is only synced when the operation is successful. This is the behavior always used for 'INCREMENTAL' backups.
	BitmapSyncModeOnSuccess BitmapSyncMode = "on-success"
	// BitmapSyncModeNever The bitmap is never synchronized with the operation, and is treated solely as a read-only manifest of blocks to copy.
	BitmapSyncModeNever BitmapSyncMode = "never"
	// BitmapSyncModeAlways The bitmap is always synchronized with the operation, regardless of whether or not the operation was successful.
	BitmapSyncModeAlways BitmapSyncMode = "always"
)

// MirrorCopyMode An enumeration whose values tell the mirror block job when to trigger writes to the target.
type MirrorCopyMode string

const (
	// MirrorCopyModeBackground copy data in background only.
	MirrorCopyModeBackground MirrorCopyMode = "background"
	// MirrorCopyModeWriteBlocking when data is written to the source, write it (synchronously) to the target as well. In addition, data is copied in background just like in @background mode.
	MirrorCopyModeWriteBlocking MirrorCopyMode = "write-blocking"
)

// BlockJobInfoMirror
//
// Information specific to mirror block jobs.
type BlockJobInfoMirror struct {
	// ActivelySynced Whether the source is actively synced to the target, i.e. same data and new writes are done synchronously to both.
	ActivelySynced bool `json:"actively-synced"`
}

// BlockJobInfo
//
// Information about a long-running block device operation.
type BlockJobInfo struct {
	// Discriminator: type

	// Type the job type ('stream' for image streaming)
	Type JobType `json:"type"`
	// Device The job identifier. Originally the device name but other values are allowed since QEMU 2.7
	Device string `json:"device"`
	// Len Estimated @offset value at the completion of the job. This value can arbitrarily change while the job is running, in both directions.
	Len int64 `json:"len"`
	// Offset Progress made until now. The unit is arbitrary and the value can only meaningfully be used for the ratio of @offset to @len. The value is monotonically increasing.
	Offset int64 `json:"offset"`
	// Busy false if the job is known to be in a quiescent state, with no pending I/O. (Since 1.3)
	Busy bool `json:"busy"`
	// Paused whether the job is paused or, if @busy is true, will pause itself as soon as possible. (Since 1.3)
	Paused bool `json:"paused"`
	// Speed the rate limit, bytes per second
	Speed int64 `json:"speed"`
	// IoStatus the status of the job (since 1.3)
	IoStatus BlockDeviceIoStatus `json:"io-status"`
	// Ready true if the job may be completed (since 2.2)
	Ready bool `json:"ready"`
	// Status Current job state/status (since 2.12)
	Status JobStatus `json:"status"`
	// AutoFinalize Job will finalize itself when PENDING, moving to the CONCLUDED state. (since 2.12)
	AutoFinalize bool `json:"auto-finalize"`
	// AutoDismiss Job will dismiss itself when CONCLUDED, moving to the NULL state and disappearing from the query list. (since 2.12)
	AutoDismiss bool `json:"auto-dismiss"`
	// Error Error information if the job did not complete successfully. Not set if the job completed successfully. (since 2.12.1)
	Error *string `json:"error,omitempty"`

	Mirror *BlockJobInfoMirror `json:"-"`
}

func (u BlockJobInfo) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "mirror":
		if u.Mirror == nil {
			return nil, fmt.Errorf("expected Mirror to be set")
		}

		return json.Marshal(struct {
			Type         JobType             `json:"type"`
			Device       string              `json:"device"`
			Len          int64               `json:"len"`
			Offset       int64               `json:"offset"`
			Busy         bool                `json:"busy"`
			Paused       bool                `json:"paused"`
			Speed        int64               `json:"speed"`
			IoStatus     BlockDeviceIoStatus `json:"io-status"`
			Ready        bool                `json:"ready"`
			Status       JobStatus           `json:"status"`
			AutoFinalize bool                `json:"auto-finalize"`
			AutoDismiss  bool                `json:"auto-dismiss"`
			Error        *string             `json:"error,omitempty"`
			*BlockJobInfoMirror
		}{
			Type:               u.Type,
			Device:             u.Device,
			Len:                u.Len,
			Offset:             u.Offset,
			Busy:               u.Busy,
			Paused:             u.Paused,
			Speed:              u.Speed,
			IoStatus:           u.IoStatus,
			Ready:              u.Ready,
			Status:             u.Status,
			AutoFinalize:       u.AutoFinalize,
			AutoDismiss:        u.AutoDismiss,
			Error:              u.Error,
			BlockJobInfoMirror: u.Mirror,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// QueryBlockJobs
//
// Return information about long-running block device operations.
type QueryBlockJobs struct {
}

func (QueryBlockJobs) Command() string {
	return "query-block-jobs"
}

func (cmd QueryBlockJobs) Execute(ctx context.Context, client api.Client) ([]BlockJobInfo, error) {
	var ret []BlockJobInfo

	return ret, client.Execute(ctx, "query-block-jobs", cmd, &ret)
}

// BlockResize
//
// Resize a block image while a guest is running. Either @device or @node-name must be set but not both.
type BlockResize struct {
	// Device the name of the device to get the image resized
	Device *string `json:"device,omitempty"`
	// NodeName graph node name to get the image resized (Since 2.0)
	NodeName *string `json:"node-name,omitempty"`
	// Size new image size in bytes
	Size int64 `json:"size"`
}

func (BlockResize) Command() string {
	return "block_resize"
}

func (cmd BlockResize) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block_resize", cmd, nil)
}

// NewImageMode An enumeration that tells QEMU how to set the backing file path in a new image file.
type NewImageMode string

const (
	// NewImageModeExisting QEMU should look for an existing image file.
	NewImageModeExisting NewImageMode = "existing"
	// NewImageModeAbsolutePaths QEMU should create a new image with absolute paths for the backing file. If there is no backing file available, the new image will not be backed either.
	NewImageModeAbsolutePaths NewImageMode = "absolute-paths"
)

// BlockdevSnapshotSync
//
// Either @device or @node-name must be set but not both.
type BlockdevSnapshotSync struct {
	// Device the name of the device to take a snapshot of.
	Device *string `json:"device,omitempty"`
	// NodeName graph node name to generate the snapshot from (Since 2.0)
	NodeName *string `json:"node-name,omitempty"`
	// SnapshotFile the target of the new overlay image. If the file exists, or if it is a device, the overlay will be created in the existing file/device. Otherwise, a new file will be created.
	SnapshotFile string `json:"snapshot-file"`
	// SnapshotNodeName the graph node name of the new image (Since 2.0)
	SnapshotNodeName *string `json:"snapshot-node-name,omitempty"`
	// Format the format of the overlay image, default is 'qcow2'.
	Format *string `json:"format,omitempty"`
	// Mode whether and how QEMU should create a new image, default is 'absolute-paths'.
	Mode *NewImageMode `json:"mode,omitempty"`
}

// BlockdevSnapshot
type BlockdevSnapshot struct {
	// Node device or node name that will have a snapshot taken.
	Node string `json:"node"`
	// Overlay reference to the existing block device that will become the overlay of @node, as part of taking the snapshot. It must not have a current backing file (this can be achieved by passing
	Overlay string `json:"overlay"`
}

// BackupPerf
//
// Optional parameters for backup. These parameters don't affect functionality, but may significantly affect performance.
type BackupPerf struct {
	// UseCopyRange Use copy offloading. Default false.
	UseCopyRange *bool `json:"use-copy-range,omitempty"`
	// MaxWorkers Maximum number of parallel requests for the sustained background copying process. Doesn't influence copy-before-write operations. Default 64.
	MaxWorkers *int64 `json:"max-workers,omitempty"`
	// MaxChunk Maximum request length for the sustained background copying process. Doesn't influence copy-before-write operations. 0 means unlimited. If max-chunk is non-zero then it should not be less than job cluster size which is calculated as maximum of target image cluster size and 64k. Default 0.
	MaxChunk *int64 `json:"max-chunk,omitempty"`
}

// BackupCommon
type BackupCommon struct {
	// JobId identifier for the newly-created block job. If omitted, the device name will be used. (Since 2.7)
	JobId *string `json:"job-id,omitempty"`
	// Device the device name or node-name of a root node which should be copied.
	Device string `json:"device"`
	// Sync what parts of the disk image should be copied to the destination (all the disk, only the sectors allocated in the topmost image, from a dirty bitmap, or only new I/O).
	Sync MirrorSyncMode `json:"sync"`
	// Speed the maximum speed, in bytes per second. The default is 0, for unlimited.
	Speed *int64 `json:"speed,omitempty"`
	// Bitmap The name of a dirty bitmap to use. Must be present if sync is "bitmap" or "incremental". Can be present if sync is "full" or "top". Must not be present otherwise. (Since 2.4 (drive-backup), 3.1 (blockdev-backup))
	Bitmap *string `json:"bitmap,omitempty"`
	// BitmapMode Specifies the type of data the bitmap should contain after the operation concludes. Must be present if a bitmap was provided, Must NOT be present otherwise. (Since 4.2)
	BitmapMode *BitmapSyncMode `json:"bitmap-mode,omitempty"`
	// Compress true to compress data, if the target format supports it.
	Compress *bool `json:"compress,omitempty"`
	// OnSourceError the action to take on an error on the source, default 'report'. 'stop' and 'enospc' can only be used if the block device supports io-status (see BlockInfo).
	OnSourceError *BlockdevOnError `json:"on-source-error,omitempty"`
	// OnTargetError the action to take on an error on the target, default 'report' (no limitations, since this applies to a different block device than @device).
	OnTargetError *BlockdevOnError `json:"on-target-error,omitempty"`
	// AutoFinalize When false, this job will wait in a PENDING state after it has finished its work, waiting for @block-job-finalize before making any block graph changes. When true, this job will automatically perform its abort or commit actions. Defaults to true. (Since 2.12)
	AutoFinalize *bool `json:"auto-finalize,omitempty"`
	// AutoDismiss When false, this job will wait in a CONCLUDED state after it has completely ceased all work, and awaits @block-job-dismiss. When true, this job will automatically disappear from the query list without user intervention. Defaults to true. (Since 2.12)
	AutoDismiss *bool `json:"auto-dismiss,omitempty"`
	// FilterNodeName the node name that should be assigned to the filter driver that the backup job inserts into the graph above node specified by @drive. If this option is not given, a node
	FilterNodeName *string `json:"filter-node-name,omitempty"`
	// Perf Performance options. (Since 6.0)
	Perf *BackupPerf `json:"x-perf,omitempty"`
}

// DriveBackup
type DriveBackup struct {
	BackupCommon

	// Target the target of the new image. If the file exists, or if it is a device, the existing file/device will be used as the new destination. If it does not exist, a new file will be created.
	Target string `json:"target"`
	// Format the format of the new destination, default is to probe if @mode is 'existing', else the format of the source
	Format *string `json:"format,omitempty"`
	// Mode whether and how QEMU should create a new image, default is 'absolute-paths'.
	Mode *NewImageMode `json:"mode,omitempty"`
}

// BlockdevBackup
type BlockdevBackup struct {
	BackupCommon

	// Target the device name or node-name of the backup target node.
	Target string `json:"target"`
}

func (BlockdevSnapshotSync) Command() string {
	return "blockdev-snapshot-sync"
}

func (cmd BlockdevSnapshotSync) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-snapshot-sync", cmd, nil)
}

func (BlockdevSnapshot) Command() string {
	return "blockdev-snapshot"
}

func (cmd BlockdevSnapshot) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-snapshot", cmd, nil)
}

// ChangeBackingFile
//
// Change the backing file in the image file metadata. This does not cause QEMU to reopen the image file to reparse the backing filename (it may, however, perform a reopen to change permissions from r/o -> r/w -> r/o, if needed). The new backing file string is written into the image file metadata, and the QEMU internal strings are updated.
type ChangeBackingFile struct {
	// Device The device name or node-name of the root node that owns image-node-name.
	Device string `json:"device"`
	// ImageNodeName The name of the block driver state node of the image to modify. The "device" argument is used to verify "image-node-name" is in the chain described by "device".
	ImageNodeName string `json:"image-node-name"`
	// BackingFile The string to write as the backing file. This string is not validated, so care should be taken when specifying the string or the image chain may not be able to be reopened again.
	BackingFile string `json:"backing-file"`
}

func (ChangeBackingFile) Command() string {
	return "change-backing-file"
}

func (cmd ChangeBackingFile) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "change-backing-file", cmd, nil)
}

// BlockCommit
//
// Live commit of data from overlay image nodes into backing nodes - i.e., writes data between 'top' and 'base' into 'base'. If top == base, that is an error. If top has no overlays on top of it, or if it is in use by a writer, the job will not be completed by itself. The user needs to complete the job with the block-job-complete command after getting the ready event. (Since 2.0) If the base image is smaller than top, then the base image will be resized to be the same size as top. If top is smaller than the base image, the base will not be truncated. If you want the base image size to match the size of the smaller top, you can safely truncate it yourself once the commit operation successfully completes.
type BlockCommit struct {
	// JobId identifier for the newly-created block job. If omitted, the device name will be used. (Since 2.7)
	JobId *string `json:"job-id,omitempty"`
	// Device the device name or node-name of a root node
	Device string `json:"device"`
	// BaseNode The node name of the backing image to write data into. If not specified, this is the deepest backing image.
	BaseNode *string `json:"base-node,omitempty"`
	// Base Same as @base-node, except that it is a file name rather than a node name. This must be the exact filename string that was used to open the node; other strings, even if addressing the same file, are not accepted
	Base *string `json:"base,omitempty"`
	// TopNode The node name of the backing image within the image chain which contains the topmost data to be committed down. If not
	TopNode *string `json:"top-node,omitempty"`
	// Top Same as @top-node, except that it is a file name rather than a node name. This must be the exact filename string that was used to open the node; other strings, even if addressing the same file, are not accepted
	Top *string `json:"top,omitempty"`
	// BackingFile The backing file string to write into the overlay image of 'top'. If 'top' does not have an overlay image, or if 'top' is in use by a writer, specifying a backing file string is an error. This filename is not validated. If a pathname string is such that it cannot be resolved by QEMU, that means that subsequent QMP or HMP commands must use node-names for the image in question, as filename lookup methods will fail. If not specified, QEMU will automatically determine the backing file string to use, or error out if there is no obvious choice. Care should be taken when specifying the string, to specify a valid filename or protocol. (Since 2.1)
	BackingFile *string `json:"backing-file,omitempty"`
	// BackingMaskProtocol If true, replace any protocol mentioned in the 'backing file format' with 'raw', rather than storing the protocol name as the backing format. Can be used even when no image header will be updated (default false; since 9.0).
	BackingMaskProtocol *bool `json:"backing-mask-protocol,omitempty"`
	// Speed the maximum speed, in bytes per second
	Speed *int64 `json:"speed,omitempty"`
	// OnError the action to take on an error. 'ignore' means that the
	OnError *BlockdevOnError `json:"on-error,omitempty"`
	// FilterNodeName the node name that should be assigned to the filter driver that the commit job inserts into the graph above @top. If this option is not given, a node name is
	FilterNodeName *string `json:"filter-node-name,omitempty"`
	// AutoFinalize When false, this job will wait in a PENDING state after it has finished its work, waiting for @block-job-finalize before making any block graph changes. When true, this job will automatically perform its abort or commit actions. Defaults to true. (Since 3.1)
	AutoFinalize *bool `json:"auto-finalize,omitempty"`
	// AutoDismiss When false, this job will wait in a CONCLUDED state after it has completely ceased all work, and awaits @block-job-dismiss. When true, this job will automatically disappear from the query list without user intervention. Defaults to true. (Since 3.1)
	AutoDismiss *bool `json:"auto-dismiss,omitempty"`
}

func (BlockCommit) Command() string {
	return "block-commit"
}

func (cmd BlockCommit) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-commit", cmd, nil)
}

func (DriveBackup) Command() string {
	return "drive-backup"
}

func (cmd DriveBackup) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "drive-backup", cmd, nil)
}

func (BlockdevBackup) Command() string {
	return "blockdev-backup"
}

func (cmd BlockdevBackup) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-backup", cmd, nil)
}

// QueryNamedBlockNodes
//
// Get the named block driver list
type QueryNamedBlockNodes struct {
	// Flat Omit the nested data about backing image ("backing-image" key) if true. Default is false (Since 5.0)
	Flat *bool `json:"flat,omitempty"`
}

func (QueryNamedBlockNodes) Command() string {
	return "query-named-block-nodes"
}

func (cmd QueryNamedBlockNodes) Execute(ctx context.Context, client api.Client) ([]BlockDeviceInfo, error) {
	var ret []BlockDeviceInfo

	return ret, client.Execute(ctx, "query-named-block-nodes", cmd, &ret)
}

// XDbgBlockGraphNodeType
type XDbgBlockGraphNodeType string

const (
	// XDbgBlockGraphNodeTypeBlockBackend corresponds to BlockBackend
	XDbgBlockGraphNodeTypeBlockBackend XDbgBlockGraphNodeType = "block-backend"
	// XDbgBlockGraphNodeTypeBlockJob corresponds to BlockJob
	XDbgBlockGraphNodeTypeBlockJob XDbgBlockGraphNodeType = "block-job"
	// XDbgBlockGraphNodeTypeBlockDriver corresponds to BlockDriverState
	XDbgBlockGraphNodeTypeBlockDriver XDbgBlockGraphNodeType = "block-driver"
)

// XDbgBlockGraphNode
type XDbgBlockGraphNode struct {
	// Id Block graph node identifier. This @id is generated only for x-debug-query-block-graph and does not relate to any other identifiers in Qemu.
	Id uint64 `json:"id"`
	// Type Type of graph node. Can be one of block-backend, block-job or block-driver-state.
	Type XDbgBlockGraphNodeType `json:"type"`
	// Name Human readable name of the node. Corresponds to node-name for block-driver-state nodes; is not guaranteed to be unique in the whole graph (with block-jobs and block-backends).
	Name string `json:"name"`
}

// BlockPermission Enum of base block permissions.
type BlockPermission string

const (
	// BlockPermissionConsistentRead A user that has the "permission" of consistent reads is guaranteed that their view of the contents of the block device is complete and self-consistent, representing the contents of a disk at a specific point. For most block devices (including their backing files) this is true, but the property cannot be maintained in a few situations like for intermediate nodes of a commit block job.
	BlockPermissionConsistentRead BlockPermission = "consistent-read"
	// BlockPermissionWrite This permission is required to change the visible disk contents.
	BlockPermissionWrite BlockPermission = "write"
	// BlockPermissionWriteUnchanged This permission (which is weaker than BLK_PERM_WRITE) is both enough and required for writes to the block node when the caller promises that the visible disk content doesn't change. As the BLK_PERM_WRITE permission is strictly stronger, either is sufficient to perform an unchanging write.
	BlockPermissionWriteUnchanged BlockPermission = "write-unchanged"
	// BlockPermissionResize This permission is required to change the size of a block node.
	BlockPermissionResize BlockPermission = "resize"
)

// XDbgBlockGraphEdge
//
// Block Graph edge description for x-debug-query-block-graph.
type XDbgBlockGraphEdge struct {
	// Parent parent id
	Parent uint64 `json:"parent"`
	// Child child id
	Child uint64 `json:"child"`
	// Name name of the relation (examples are 'file' and 'backing')
	Name string `json:"name"`
	// Perm granted permissions for the parent operating on the child
	Perm []BlockPermission `json:"perm"`
	// SharedPerm permissions that can still be granted to other users of the child while it is still attached to this parent
	SharedPerm []BlockPermission `json:"shared-perm"`
}

// XDbgBlockGraph
//
// Block Graph - list of nodes and list of edges.
type XDbgBlockGraph struct {
	Nodes []XDbgBlockGraphNode `json:"nodes"`
	Edges []XDbgBlockGraphEdge `json:"edges"`
}

// DebugQueryBlockGraph
//
// Get the block graph.
type DebugQueryBlockGraph struct {
}

func (DebugQueryBlockGraph) Command() string {
	return "x-debug-query-block-graph"
}

func (cmd DebugQueryBlockGraph) Execute(ctx context.Context, client api.Client) (XDbgBlockGraph, error) {
	var ret XDbgBlockGraph

	return ret, client.Execute(ctx, "x-debug-query-block-graph", cmd, &ret)
}

func (DriveMirror) Command() string {
	return "drive-mirror"
}

func (cmd DriveMirror) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "drive-mirror", cmd, nil)
}

// DriveMirror
//
// A set of parameters describing drive mirror setup.
type DriveMirror struct {
	// JobId identifier for the newly-created block job. If omitted, the device name will be used. (Since 2.7)
	JobId *string `json:"job-id,omitempty"`
	// Device the device name or node-name of a root node whose writes should be mirrored.
	Device string `json:"device"`
	// Target the target of the new image. If the file exists, or if it is a device, the existing file/device will be used as the new destination. If it does not exist, a new file will be created.
	Target string `json:"target"`
	// Format the format of the new destination, default is to probe if @mode is 'existing', else the format of the source
	Format *string `json:"format,omitempty"`
	// NodeName the new block driver state node name in the graph (Since 2.1)
	NodeName *string `json:"node-name,omitempty"`
	// Replaces with sync=full graph node name to be replaced by the new image when a whole image copy is done. This can be used to repair broken Quorum files. By default, @device is replaced, although implicitly created filters on it are kept. (Since 2.1)
	Replaces *string `json:"replaces,omitempty"`
	// Sync what parts of the disk image should be copied to the destination (all the disk, only the sectors allocated in the topmost image, or only new I/O).
	Sync MirrorSyncMode `json:"sync"`
	// Mode whether and how QEMU should create a new image, default is 'absolute-paths'.
	Mode *NewImageMode `json:"mode,omitempty"`
	// Speed the maximum speed, in bytes per second
	Speed *int64 `json:"speed,omitempty"`
	// Granularity granularity of the dirty bitmap, default is 64K if the image format doesn't have clusters, 4K if the clusters are smaller than that, else the cluster size. Must be a power of 2 between 512 and 64M (since 1.4).
	Granularity *uint32 `json:"granularity,omitempty"`
	// BufSize maximum amount of data in flight from source to target (since 1.4).
	BufSize *int64 `json:"buf-size,omitempty"`
	// OnSourceError the action to take on an error on the source, default 'report'. 'stop' and 'enospc' can only be used if the block device supports io-status (see BlockInfo).
	OnSourceError *BlockdevOnError `json:"on-source-error,omitempty"`
	// OnTargetError the action to take on an error on the target, default 'report' (no limitations, since this applies to a different block device than @device).
	OnTargetError *BlockdevOnError `json:"on-target-error,omitempty"`
	// Unmap Whether to try to unmap target sectors where source has only zero. If true, and target unallocated sectors will read as zero, target image sectors will be unmapped; otherwise, zeroes will be written. Both will result in identical contents. Default is true. (Since 2.4)
	Unmap *bool `json:"unmap,omitempty"`
	// CopyMode when to copy data to the destination; defaults to
	CopyMode *MirrorCopyMode `json:"copy-mode,omitempty"`
	// AutoFinalize When false, this job will wait in a PENDING state after it has finished its work, waiting for @block-job-finalize before making any block graph changes. When true, this job will automatically perform its abort or commit actions. Defaults to true. (Since 3.1)
	AutoFinalize *bool `json:"auto-finalize,omitempty"`
	// AutoDismiss When false, this job will wait in a CONCLUDED state after it has completely ceased all work, and awaits @block-job-dismiss. When true, this job will automatically disappear from the query list without user intervention. Defaults to true. (Since 3.1)
	AutoDismiss *bool `json:"auto-dismiss,omitempty"`
}

// BlockDirtyBitmap
type BlockDirtyBitmap struct {
	// Node name of device/node which the bitmap is tracking
	Node string `json:"node"`
	// Name name of the dirty bitmap
	Name string `json:"name"`
}

// BlockDirtyBitmapAdd
type BlockDirtyBitmapAdd struct {
	// Node name of device/node which the bitmap is tracking
	Node string `json:"node"`
	// Name name of the dirty bitmap (must be less than 1024 bytes)
	Name string `json:"name"`
	// Granularity the bitmap granularity, default is 64k for block-dirty-bitmap-add
	Granularity *uint32 `json:"granularity,omitempty"`
	// Persistent the bitmap is persistent, i.e. it will be saved to the corresponding block device image file on its close. For now only Qcow2 disks support persistent bitmaps. Default is false
	Persistent *bool `json:"persistent,omitempty"`
	// Disabled the bitmap is created in the disabled state, which means that it will not track drive changes. The bitmap may be enabled
	Disabled *bool `json:"disabled,omitempty"`
}

// BlockDirtyBitmapOrStr
type BlockDirtyBitmapOrStr struct {
	// Local name of the bitmap, attached to the same node as target bitmap.
	Local *string `json:"-"`
	// External bitmap with specified node
	External *BlockDirtyBitmap `json:"-"`
}

func (a BlockDirtyBitmapOrStr) MarshalJSON() ([]byte, error) {
	switch {
	case a.Local != nil:
		return json.Marshal(a.Local)
	case a.External != nil:
		return json.Marshal(a.External)
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockDirtyBitmapMerge
type BlockDirtyBitmapMerge struct {
	// Node name of device/node which the @target bitmap is tracking
	Node string `json:"node"`
	// Target name of the destination dirty bitmap
	Target string `json:"target"`
	// Bitmaps name(s) of the source dirty bitmap(s) at @node and/or fully specified BlockDirtyBitmap elements. The latter are supported since 4.1.
	Bitmaps []BlockDirtyBitmapOrStr `json:"bitmaps"`
}

func (BlockDirtyBitmapAdd) Command() string {
	return "block-dirty-bitmap-add"
}

func (cmd BlockDirtyBitmapAdd) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-dirty-bitmap-add", cmd, nil)
}

// BlockDirtyBitmapRemove
//
// Stop write tracking and remove the dirty bitmap that was created with block-dirty-bitmap-add. If the bitmap is persistent, remove it from its storage too.
type BlockDirtyBitmapRemove struct {
	BlockDirtyBitmap
}

func (BlockDirtyBitmapRemove) Command() string {
	return "block-dirty-bitmap-remove"
}

func (cmd BlockDirtyBitmapRemove) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-dirty-bitmap-remove", cmd, nil)
}

// BlockDirtyBitmapClear
//
// Clear (reset) a dirty bitmap on the device, so that an incremental backup from this point in time forward will only backup clusters modified after this clear operation.
type BlockDirtyBitmapClear struct {
	BlockDirtyBitmap
}

func (BlockDirtyBitmapClear) Command() string {
	return "block-dirty-bitmap-clear"
}

func (cmd BlockDirtyBitmapClear) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-dirty-bitmap-clear", cmd, nil)
}

// BlockDirtyBitmapEnable
//
// Enables a dirty bitmap so that it will begin tracking disk changes.
type BlockDirtyBitmapEnable struct {
	BlockDirtyBitmap
}

func (BlockDirtyBitmapEnable) Command() string {
	return "block-dirty-bitmap-enable"
}

func (cmd BlockDirtyBitmapEnable) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-dirty-bitmap-enable", cmd, nil)
}

// BlockDirtyBitmapDisable
//
// Disables a dirty bitmap so that it will stop tracking disk changes.
type BlockDirtyBitmapDisable struct {
	BlockDirtyBitmap
}

func (BlockDirtyBitmapDisable) Command() string {
	return "block-dirty-bitmap-disable"
}

func (cmd BlockDirtyBitmapDisable) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-dirty-bitmap-disable", cmd, nil)
}

func (BlockDirtyBitmapMerge) Command() string {
	return "block-dirty-bitmap-merge"
}

func (cmd BlockDirtyBitmapMerge) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-dirty-bitmap-merge", cmd, nil)
}

// BlockDirtyBitmapSha256
//
// SHA256 hash of dirty bitmap data
type BlockDirtyBitmapSha256 struct {
	// Sha256 ASCII representation of SHA256 bitmap hash
	Sha256 string `json:"sha256"`
}

// DebugBlockDirtyBitmapSha256
//
// Get bitmap SHA256.
type DebugBlockDirtyBitmapSha256 struct {
	BlockDirtyBitmap
}

func (DebugBlockDirtyBitmapSha256) Command() string {
	return "x-debug-block-dirty-bitmap-sha256"
}

func (cmd DebugBlockDirtyBitmapSha256) Execute(ctx context.Context, client api.Client) (BlockDirtyBitmapSha256, error) {
	var ret BlockDirtyBitmapSha256

	return ret, client.Execute(ctx, "x-debug-block-dirty-bitmap-sha256", cmd, &ret)
}

// BlockdevMirror
//
// Start mirroring a block device's writes to a new destination.
type BlockdevMirror struct {
	// JobId identifier for the newly-created block job. If omitted, the device name will be used. (Since 2.7)
	JobId *string `json:"job-id,omitempty"`
	// Device The device name or node-name of a root node whose writes should be mirrored.
	Device string `json:"device"`
	// Target the id or node-name of the block device to mirror to. This mustn't be attached to guest.
	Target string `json:"target"`
	// Replaces with sync=full graph node name to be replaced by the new image when a whole image copy is done. This can be used to repair broken Quorum files. By default, @device is replaced, although implicitly created filters on it are kept.
	Replaces *string `json:"replaces,omitempty"`
	// Sync what parts of the disk image should be copied to the destination (all the disk, only the sectors allocated in the topmost image, or only new I/O).
	Sync MirrorSyncMode `json:"sync"`
	// Speed the maximum speed, in bytes per second
	Speed *int64 `json:"speed,omitempty"`
	// Granularity granularity of the dirty bitmap, default is 64K if the image format doesn't have clusters, 4K if the clusters are smaller than that, else the cluster size. Must be a power of 2 between 512 and 64M
	Granularity *uint32 `json:"granularity,omitempty"`
	// BufSize maximum amount of data in flight from source to target
	BufSize *int64 `json:"buf-size,omitempty"`
	// OnSourceError the action to take on an error on the source, default 'report'. 'stop' and 'enospc' can only be used if the block device supports io-status (see BlockInfo).
	OnSourceError *BlockdevOnError `json:"on-source-error,omitempty"`
	// OnTargetError the action to take on an error on the target, default 'report' (no limitations, since this applies to a different block device than @device).
	OnTargetError *BlockdevOnError `json:"on-target-error,omitempty"`
	// FilterNodeName the node name that should be assigned to the filter driver that the mirror job inserts into the graph above @device. If this option is not given, a node name is
	FilterNodeName *string `json:"filter-node-name,omitempty"`
	// CopyMode when to copy data to the destination; defaults to
	CopyMode *MirrorCopyMode `json:"copy-mode,omitempty"`
	// AutoFinalize When false, this job will wait in a PENDING state after it has finished its work, waiting for @block-job-finalize before making any block graph changes. When true, this job will automatically perform its abort or commit actions. Defaults to true. (Since 3.1)
	AutoFinalize *bool `json:"auto-finalize,omitempty"`
	// AutoDismiss When false, this job will wait in a CONCLUDED state after it has completely ceased all work, and awaits @block-job-dismiss. When true, this job will automatically disappear from the query list without user intervention. Defaults to true. (Since 3.1)
	AutoDismiss *bool `json:"auto-dismiss,omitempty"`
}

func (BlockdevMirror) Command() string {
	return "blockdev-mirror"
}

func (cmd BlockdevMirror) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-mirror", cmd, nil)
}

// BlockIOThrottle
//
// A set of parameters describing block throttling.
type BlockIOThrottle struct {
	// Device Block device name
	Device *string `json:"device,omitempty"`
	// Id The name or QOM path of the guest device (since: 2.8)
	Id *string `json:"id,omitempty"`
	// Bps total throughput limit in bytes per second
	Bps int64 `json:"bps"`
	// BpsRd read throughput limit in bytes per second
	BpsRd int64 `json:"bps_rd"`
	// BpsWr write throughput limit in bytes per second
	BpsWr int64 `json:"bps_wr"`
	// Iops total I/O operations per second
	Iops int64 `json:"iops"`
	// IopsRd read I/O operations per second
	IopsRd int64 `json:"iops_rd"`
	// IopsWr write I/O operations per second
	IopsWr int64 `json:"iops_wr"`
	// BpsMax total throughput limit during bursts, in bytes (Since 1.7)
	BpsMax *int64 `json:"bps_max,omitempty"`
	// BpsRdMax read throughput limit during bursts, in bytes (Since 1.7)
	BpsRdMax *int64 `json:"bps_rd_max,omitempty"`
	// BpsWrMax write throughput limit during bursts, in bytes (Since 1.7)
	BpsWrMax *int64 `json:"bps_wr_max,omitempty"`
	// IopsMax total I/O operations per second during bursts, in bytes (Since 1.7)
	IopsMax *int64 `json:"iops_max,omitempty"`
	// IopsRdMax read I/O operations per second during bursts, in bytes (Since 1.7)
	IopsRdMax *int64 `json:"iops_rd_max,omitempty"`
	// IopsWrMax write I/O operations per second during bursts, in bytes (Since 1.7)
	IopsWrMax *int64 `json:"iops_wr_max,omitempty"`
	// BpsMaxLength maximum length of the @bps_max burst period, in seconds. It must only be set if @bps_max is set as well. Defaults to 1. (Since 2.6)
	BpsMaxLength *int64 `json:"bps_max_length,omitempty"`
	// BpsRdMaxLength maximum length of the @bps_rd_max burst period, in seconds. It must only be set if @bps_rd_max is set as well. Defaults to 1. (Since 2.6)
	BpsRdMaxLength *int64 `json:"bps_rd_max_length,omitempty"`
	// BpsWrMaxLength maximum length of the @bps_wr_max burst period, in seconds. It must only be set if @bps_wr_max is set as well. Defaults to 1. (Since 2.6)
	BpsWrMaxLength *int64 `json:"bps_wr_max_length,omitempty"`
	// IopsMaxLength maximum length of the @iops burst period, in seconds. It must only be set if @iops_max is set as well. Defaults to 1. (Since 2.6)
	IopsMaxLength *int64 `json:"iops_max_length,omitempty"`
	// IopsRdMaxLength maximum length of the @iops_rd_max burst period, in seconds. It must only be set if @iops_rd_max is set as well. Defaults to 1. (Since 2.6)
	IopsRdMaxLength *int64 `json:"iops_rd_max_length,omitempty"`
	// IopsWrMaxLength maximum length of the @iops_wr_max burst period, in seconds. It must only be set if @iops_wr_max is set as well. Defaults to 1. (Since 2.6)
	IopsWrMaxLength *int64 `json:"iops_wr_max_length,omitempty"`
	// IopsSize an I/O size in bytes (Since 1.7)
	IopsSize *int64 `json:"iops_size,omitempty"`
	// Group throttle group name (Since 2.4)
	Group *string `json:"group,omitempty"`
}

// ThrottleLimits
//
// Limit parameters for throttling. Since some limit combinations are illegal, limits should always be set in one transaction. All fields are optional. When setting limits, if a field is missing the current value is not changed.
type ThrottleLimits struct {
	// IopsTotal limit total I/O operations per second
	IopsTotal *int64 `json:"iops-total,omitempty"`
	// IopsTotalMax I/O operations burst
	IopsTotalMax *int64 `json:"iops-total-max,omitempty"`
	// IopsTotalMaxLength length of the iops-total-max burst period, in seconds It must only be set if @iops-total-max is set as well.
	IopsTotalMaxLength *int64 `json:"iops-total-max-length,omitempty"`
	// IopsRead limit read operations per second
	IopsRead *int64 `json:"iops-read,omitempty"`
	// IopsReadMax I/O operations read burst
	IopsReadMax *int64 `json:"iops-read-max,omitempty"`
	// IopsReadMaxLength length of the iops-read-max burst period, in seconds It must only be set if @iops-read-max is set as well.
	IopsReadMaxLength *int64 `json:"iops-read-max-length,omitempty"`
	// IopsWrite limit write operations per second
	IopsWrite *int64 `json:"iops-write,omitempty"`
	// IopsWriteMax I/O operations write burst
	IopsWriteMax *int64 `json:"iops-write-max,omitempty"`
	// IopsWriteMaxLength length of the iops-write-max burst period, in seconds It must only be set if @iops-write-max is set as well.
	IopsWriteMaxLength *int64 `json:"iops-write-max-length,omitempty"`
	// BpsTotal limit total bytes per second
	BpsTotal *int64 `json:"bps-total,omitempty"`
	// BpsTotalMax total bytes burst
	BpsTotalMax *int64 `json:"bps-total-max,omitempty"`
	// BpsTotalMaxLength length of the bps-total-max burst period, in seconds. It must only be set if @bps-total-max is set as well.
	BpsTotalMaxLength *int64 `json:"bps-total-max-length,omitempty"`
	// BpsRead limit read bytes per second
	BpsRead *int64 `json:"bps-read,omitempty"`
	// BpsReadMax total bytes read burst
	BpsReadMax *int64 `json:"bps-read-max,omitempty"`
	// BpsReadMaxLength length of the bps-read-max burst period, in seconds It must only be set if @bps-read-max is set as well.
	BpsReadMaxLength *int64 `json:"bps-read-max-length,omitempty"`
	// BpsWrite limit write bytes per second
	BpsWrite *int64 `json:"bps-write,omitempty"`
	// BpsWriteMax total bytes write burst
	BpsWriteMax *int64 `json:"bps-write-max,omitempty"`
	// BpsWriteMaxLength length of the bps-write-max burst period, in seconds It must only be set if @bps-write-max is set as well.
	BpsWriteMaxLength *int64 `json:"bps-write-max-length,omitempty"`
	// IopsSize when limiting by iops max size of an I/O in bytes
	IopsSize *int64 `json:"iops-size,omitempty"`
}

// ThrottleGroupProperties
//
// Properties for throttle-group objects.
type ThrottleGroupProperties struct {
	// Limits limits to apply for this throttle group
	Limits             *ThrottleLimits `json:"limits,omitempty"`
	IopsTotal          *int64          `json:"x-iops-total,omitempty"`
	IopsTotalMax       *int64          `json:"x-iops-total-max,omitempty"`
	IopsTotalMaxLength *int64          `json:"x-iops-total-max-length,omitempty"`
	IopsRead           *int64          `json:"x-iops-read,omitempty"`
	IopsReadMax        *int64          `json:"x-iops-read-max,omitempty"`
	IopsReadMaxLength  *int64          `json:"x-iops-read-max-length,omitempty"`
	IopsWrite          *int64          `json:"x-iops-write,omitempty"`
	IopsWriteMax       *int64          `json:"x-iops-write-max,omitempty"`
	IopsWriteMaxLength *int64          `json:"x-iops-write-max-length,omitempty"`
	BpsTotal           *int64          `json:"x-bps-total,omitempty"`
	BpsTotalMax        *int64          `json:"x-bps-total-max,omitempty"`
	BpsTotalMaxLength  *int64          `json:"x-bps-total-max-length,omitempty"`
	BpsRead            *int64          `json:"x-bps-read,omitempty"`
	BpsReadMax         *int64          `json:"x-bps-read-max,omitempty"`
	BpsReadMaxLength   *int64          `json:"x-bps-read-max-length,omitempty"`
	BpsWrite           *int64          `json:"x-bps-write,omitempty"`
	BpsWriteMax        *int64          `json:"x-bps-write-max,omitempty"`
	BpsWriteMaxLength  *int64          `json:"x-bps-write-max-length,omitempty"`
	IopsSize           *int64          `json:"x-iops-size,omitempty"`
}

// BlockStream
//
// Copy data from a backing file into a block device. The block streaming operation is performed in the background until the entire backing file has been copied. This command returns immediately once streaming has started. The status of ongoing block streaming operations can be checked with query-block-jobs. The operation can be stopped before it has completed using the block-job-cancel command. The node that receives the data is called the top image, can be located in any part of the chain (but always above the base image; see below) and can be specified using its device or node name. Earlier qemu versions only allowed 'device' to name the top level node; presence of the 'base-node' parameter during introspection can be used as a witness of the enhanced semantics of 'device'. If a base file is specified then sectors are not copied from that base file and its backing chain. This can be used to stream a subset of the backing file chain instead of flattening the entire image. When streaming completes the image file will have the base file as its backing file, unless that node was changed while the job was running. In that case, base's parent's backing (or filtered, whichever exists) child (i.e., base at the beginning of the job) will be the new backing file. On successful completion the image file is updated to drop the backing file and the BLOCK_JOB_COMPLETED event is emitted. In case @device is a filter node, block-stream modifies the first non-filter overlay node below it to point to the new backing node instead of modifying @device itself.
type BlockStream struct {
	// JobId identifier for the newly-created block job. If omitted, the device name will be used. (Since 2.7)
	JobId *string `json:"job-id,omitempty"`
	// Device the device or node name of the top image
	Device string `json:"device"`
	// Base the common backing file name. It cannot be set if @base-node or @bottom is also set.
	Base *string `json:"base,omitempty"`
	// BaseNode the node name of the backing file. It cannot be set if @base or @bottom is also set. (Since 2.8)
	BaseNode *string `json:"base-node,omitempty"`
	// BackingFile The backing file string to write into the top image. This filename is not validated. If a pathname string is such that it cannot be resolved by QEMU, that means that subsequent QMP or HMP commands must use node-names for the image in question, as filename lookup methods will fail. If not specified, QEMU will automatically determine the backing file string to use, or error out if there is no obvious choice. Care should be taken when specifying the string, to specify a valid filename or protocol. (Since 2.1)
	BackingFile *string `json:"backing-file,omitempty"`
	// BackingMaskProtocol If true, replace any protocol mentioned in the 'backing file format' with 'raw', rather than storing the protocol name as the backing format. Can be used even when no image header will be updated (default false; since 9.0).
	BackingMaskProtocol *bool `json:"backing-mask-protocol,omitempty"`
	// Bottom the last node in the chain that should be streamed into top. It cannot be set if @base or @base-node is also set. It cannot be filter node. (Since 6.0)
	Bottom *string `json:"bottom,omitempty"`
	// Speed the maximum speed, in bytes per second
	Speed *int64 `json:"speed,omitempty"`
	// OnError the action to take on an error (default report). 'stop' and 'enospc' can only be used if the block device supports io-status (see BlockInfo). (Since 1.3)
	OnError *BlockdevOnError `json:"on-error,omitempty"`
	// FilterNodeName the node name that should be assigned to the filter driver that the stream job inserts into the graph above @device. If this option is not given, a node name is
	FilterNodeName *string `json:"filter-node-name,omitempty"`
	// AutoFinalize When false, this job will wait in a PENDING state after it has finished its work, waiting for @block-job-finalize before making any block graph changes. When true, this job will automatically perform its abort or commit actions. Defaults to true. (Since 3.1)
	AutoFinalize *bool `json:"auto-finalize,omitempty"`
	// AutoDismiss When false, this job will wait in a CONCLUDED state after it has completely ceased all work, and awaits @block-job-dismiss. When true, this job will automatically disappear from the query list without user intervention. Defaults to true. (Since 3.1)
	AutoDismiss *bool `json:"auto-dismiss,omitempty"`
}

func (BlockStream) Command() string {
	return "block-stream"
}

func (cmd BlockStream) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-stream", cmd, nil)
}

// BlockJobSetSpeed
//
// Set maximum speed for a background block operation. This command can only be issued when there is an active block job. Throttling can be disabled by setting the speed to 0.
type BlockJobSetSpeed struct {
	// Device The job identifier. This used to be a device name (hence the name of the parameter), but since QEMU 2.7 it can have other values.
	Device string `json:"device"`
	// Speed the maximum speed, in bytes per second, or 0 for unlimited. Defaults to 0.
	Speed int64 `json:"speed"`
}

func (BlockJobSetSpeed) Command() string {
	return "block-job-set-speed"
}

func (cmd BlockJobSetSpeed) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-job-set-speed", cmd, nil)
}

// BlockJobCancel
//
// Stop an active background block operation. This command returns immediately after marking the active background block operation for cancellation. It is an error to call this command if no operation is in progress. The operation will cancel as soon as possible and then emit the BLOCK_JOB_CANCELLED event. Before that happens the job is still visible when enumerated using query-block-jobs. Note that if you issue 'block-job-cancel' after 'drive-mirror' has indicated (via the event BLOCK_JOB_READY) that the source and destination are synchronized, then the event triggered by this command changes to BLOCK_JOB_COMPLETED, to indicate that the mirroring has ended and the destination now has a point-in-time copy tied to the time of the cancellation. For streaming, the image file retains its backing file unless the streaming operation happens to complete just as it is being cancelled. A new streaming operation can be started at a later time to finish copying all data from the backing file.
type BlockJobCancel struct {
	// Device The job identifier. This used to be a device name (hence the name of the parameter), but since QEMU 2.7 it can have other values.
	Device string `json:"device"`
	// Force If true, and the job has already emitted the event BLOCK_JOB_READY, abandon the job immediately (even if it is paused) instead of waiting for the destination to complete its final synchronization (since 1.3)
	Force *bool `json:"force,omitempty"`
}

func (BlockJobCancel) Command() string {
	return "block-job-cancel"
}

func (cmd BlockJobCancel) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-job-cancel", cmd, nil)
}

// BlockJobPause
//
// Pause an active background block operation. This command returns immediately after marking the active background block operation for pausing. It is an error to call this command if no operation is in progress or if the job is already paused. The operation will pause as soon as possible. No event is emitted when the operation is actually paused. Cancelling a paused job automatically resumes it.
type BlockJobPause struct {
	// Device The job identifier. This used to be a device name (hence the name of the parameter), but since QEMU 2.7 it can have other values.
	Device string `json:"device"`
}

func (BlockJobPause) Command() string {
	return "block-job-pause"
}

func (cmd BlockJobPause) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-job-pause", cmd, nil)
}

// BlockJobResume
//
// Resume an active background block operation. This command returns immediately after resuming a paused background block operation. It is an error to call this command if no operation is in progress or if the job is not paused. This command also clears the error status of the job.
type BlockJobResume struct {
	// Device The job identifier. This used to be a device name (hence the name of the parameter), but since QEMU 2.7 it can have other values.
	Device string `json:"device"`
}

func (BlockJobResume) Command() string {
	return "block-job-resume"
}

func (cmd BlockJobResume) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-job-resume", cmd, nil)
}

// BlockJobComplete
//
// Manually trigger completion of an active background block operation. This is supported for drive mirroring, where it also switches the device to write to the target path only. The ability to complete is signaled with a BLOCK_JOB_READY event. This command completes an active background block operation synchronously. The ordering of this command's return with the BLOCK_JOB_COMPLETED event is not defined. Note that if an I/O error
type BlockJobComplete struct {
	// Device The job identifier. This used to be a device name (hence the name of the parameter), but since QEMU 2.7 it can have other values.
	Device string `json:"device"`
}

func (BlockJobComplete) Command() string {
	return "block-job-complete"
}

func (cmd BlockJobComplete) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-job-complete", cmd, nil)
}

// BlockJobDismiss
//
// For jobs that have already concluded, remove them from the block-job-query list. This command only needs to be run for jobs which were started with QEMU 2.12+ job lifetime management semantics. This command will refuse to operate on any job that has not yet reached its terminal state, JOB_STATUS_CONCLUDED. For jobs that make use of the BLOCK_JOB_READY event, block-job-cancel or block-job-complete will still need to be used as appropriate.
type BlockJobDismiss struct {
	// Id The job identifier.
	Id string `json:"id"`
}

func (BlockJobDismiss) Command() string {
	return "block-job-dismiss"
}

func (cmd BlockJobDismiss) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-job-dismiss", cmd, nil)
}

// BlockJobFinalize
//
// Once a job that has manual=true reaches the pending state, it can be instructed to finalize any graph changes and do any necessary cleanup via this command. For jobs in a transaction, instructing one job to finalize will force ALL jobs in the transaction to finalize, so it is only necessary to instruct a single member job to finalize.
type BlockJobFinalize struct {
	// Id The job identifier.
	Id string `json:"id"`
}

func (BlockJobFinalize) Command() string {
	return "block-job-finalize"
}

func (cmd BlockJobFinalize) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-job-finalize", cmd, nil)
}

// BlockJobChangeOptionsMirror
type BlockJobChangeOptionsMirror struct {
	// CopyMode Switch to this copy mode. Currently, only the switch from 'background' to 'write-blocking' is implemented.
	CopyMode MirrorCopyMode `json:"copy-mode"`
}

// BlockJobChangeOptions
//
// Block job options that can be changed after job creation.
type BlockJobChangeOptions struct {
	// Discriminator: type

	// Id The job identifier
	Id string `json:"id"`
	// Type The job type
	Type JobType `json:"type"`

	Mirror *BlockJobChangeOptionsMirror `json:"-"`
}

func (u BlockJobChangeOptions) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "mirror":
		if u.Mirror == nil {
			return nil, fmt.Errorf("expected Mirror to be set")
		}

		return json.Marshal(struct {
			Id   string  `json:"id"`
			Type JobType `json:"type"`
			*BlockJobChangeOptionsMirror
		}{
			Id:                          u.Id,
			Type:                        u.Type,
			BlockJobChangeOptionsMirror: u.Mirror,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockJobChange
//
// Change the block job's options.
type BlockJobChange struct {
	BlockJobChangeOptions
}

func (BlockJobChange) Command() string {
	return "block-job-change"
}

func (cmd BlockJobChange) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-job-change", cmd, nil)
}

// BlockdevDiscardOptions Determines how to handle discard requests.
type BlockdevDiscardOptions string

const (
	// BlockdevDiscardOptionsIgnore Ignore the request
	BlockdevDiscardOptionsIgnore BlockdevDiscardOptions = "ignore"
	// BlockdevDiscardOptionsUnmap Forward as an unmap request
	BlockdevDiscardOptionsUnmap BlockdevDiscardOptions = "unmap"
)

// BlockdevDetectZeroesOptions Describes the operation mode for the automatic conversion of plain zero writes by the OS to driver specific optimized zero write commands.
type BlockdevDetectZeroesOptions string

const (
	// BlockdevDetectZeroesOptionsOff Disabled (default)
	BlockdevDetectZeroesOptionsOff BlockdevDetectZeroesOptions = "off"
	// BlockdevDetectZeroesOptionsOn Enabled
	BlockdevDetectZeroesOptionsOn BlockdevDetectZeroesOptions = "on"
	// BlockdevDetectZeroesOptionsUnmap Enabled and even try to unmap blocks if possible. This requires also that @BlockdevDiscardOptions is set to unmap for this device.
	BlockdevDetectZeroesOptionsUnmap BlockdevDetectZeroesOptions = "unmap"
)

// BlockdevAioOptions Selects the AIO backend to handle I/O requests
type BlockdevAioOptions string

const (
	// BlockdevAioOptionsThreads Use qemu's thread pool
	BlockdevAioOptionsThreads BlockdevAioOptions = "threads"
	// BlockdevAioOptionsNative Use native AIO backend (only Linux and Windows)
	BlockdevAioOptionsNative BlockdevAioOptions = "native"
	// BlockdevAioOptionsIoUring Use linux io_uring (since 5.0)
	BlockdevAioOptionsIoUring BlockdevAioOptions = "io_uring"
)

// BlockdevCacheOptions
//
// Includes cache-related options for block devices
type BlockdevCacheOptions struct {
	// Direct enables use of O_DIRECT (bypass the host page cache;
	Direct *bool `json:"direct,omitempty"`
	// NoFlush ignore any flush requests for the device (default: false)
	NoFlush *bool `json:"no-flush,omitempty"`
}

// BlockdevDriver Drivers that are supported in block device operations.
type BlockdevDriver string

const (
	BlockdevDriverBlkdebug BlockdevDriver = "blkdebug"
	// BlockdevDriverBlklogwrites Since 3.0
	BlockdevDriverBlklogwrites BlockdevDriver = "blklogwrites"
	// BlockdevDriverBlkreplay Since 4.2
	BlockdevDriverBlkreplay BlockdevDriver = "blkreplay"
	BlockdevDriverBlkverify BlockdevDriver = "blkverify"
	BlockdevDriverBochs     BlockdevDriver = "bochs"
	BlockdevDriverCloop     BlockdevDriver = "cloop"
	// BlockdevDriverCompress Since 5.0
	BlockdevDriverCompress BlockdevDriver = "compress"
	// BlockdevDriverCopyBeforeWrite Since 6.2
	BlockdevDriverCopyBeforeWrite BlockdevDriver = "copy-before-write"
	// BlockdevDriverCopyOnRead Since 3.0
	BlockdevDriverCopyOnRead BlockdevDriver = "copy-on-read"
	BlockdevDriverDmg        BlockdevDriver = "dmg"
	BlockdevDriverFile       BlockdevDriver = "file"
	// BlockdevDriverSnapshotAccess Since 7.0
	BlockdevDriverSnapshotAccess BlockdevDriver = "snapshot-access"
	BlockdevDriverFtp            BlockdevDriver = "ftp"
	BlockdevDriverFtps           BlockdevDriver = "ftps"
	BlockdevDriverGluster        BlockdevDriver = "gluster"
	BlockdevDriverHostCdrom      BlockdevDriver = "host_cdrom"
	BlockdevDriverHostDevice     BlockdevDriver = "host_device"
	BlockdevDriverHttp           BlockdevDriver = "http"
	BlockdevDriverHttps          BlockdevDriver = "https"
	BlockdevDriverIoUring        BlockdevDriver = "io_uring"
	BlockdevDriverIscsi          BlockdevDriver = "iscsi"
	BlockdevDriverLuks           BlockdevDriver = "luks"
	BlockdevDriverNbd            BlockdevDriver = "nbd"
	BlockdevDriverNfs            BlockdevDriver = "nfs"
	BlockdevDriverNullAio        BlockdevDriver = "null-aio"
	BlockdevDriverNullCo         BlockdevDriver = "null-co"
	// BlockdevDriverNvme Since 2.12
	BlockdevDriverNvme        BlockdevDriver = "nvme"
	BlockdevDriverNvmeIoUring BlockdevDriver = "nvme-io_uring"
	BlockdevDriverParallels   BlockdevDriver = "parallels"
	BlockdevDriverPreallocate BlockdevDriver = "preallocate"
	BlockdevDriverQcow        BlockdevDriver = "qcow"
	BlockdevDriverQcow2       BlockdevDriver = "qcow2"
	BlockdevDriverQed         BlockdevDriver = "qed"
	BlockdevDriverQuorum      BlockdevDriver = "quorum"
	BlockdevDriverRaw         BlockdevDriver = "raw"
	BlockdevDriverRbd         BlockdevDriver = "rbd"
	BlockdevDriverReplication BlockdevDriver = "replication"
	BlockdevDriverSsh         BlockdevDriver = "ssh"
	// BlockdevDriverThrottle Since 2.11
	BlockdevDriverThrottle           BlockdevDriver = "throttle"
	BlockdevDriverVdi                BlockdevDriver = "vdi"
	BlockdevDriverVhdx               BlockdevDriver = "vhdx"
	BlockdevDriverVirtioBlkVfioPci   BlockdevDriver = "virtio-blk-vfio-pci"
	BlockdevDriverVirtioBlkVhostUser BlockdevDriver = "virtio-blk-vhost-user"
	BlockdevDriverVirtioBlkVhostVdpa BlockdevDriver = "virtio-blk-vhost-vdpa"
	BlockdevDriverVmdk               BlockdevDriver = "vmdk"
	BlockdevDriverVpc                BlockdevDriver = "vpc"
	BlockdevDriverVvfat              BlockdevDriver = "vvfat"
)

// BlockdevOptionsFile
//
// Driver specific block device options for the file backend.
type BlockdevOptionsFile struct {
	// Filename path to the image file
	Filename string `json:"filename"`
	// PrManager the id for the object that will handle persistent
	PrManager *string `json:"pr-manager,omitempty"`
	// Locking whether to enable file locking. If set to 'auto', only enable when Open File Descriptor (OFD) locking API is available
	Locking *OnOffAuto `json:"locking,omitempty"`
	// Aio AIO backend (default: threads) (since: 2.8)
	Aio *BlockdevAioOptions `json:"aio,omitempty"`
	// AioMaxBatch maximum number of requests to batch together into a single submission in the AIO backend. The smallest value between this and the aio-max-batch value of the IOThread object is chosen. 0 means that the AIO backend will handle it
	AioMaxBatch *int64 `json:"aio-max-batch,omitempty"`
	// DropCache invalidate page cache during live migration. This prevents stale data on the migration destination with cache.direct=off. Currently only supported on Linux hosts.
	DropCache *bool `json:"drop-cache,omitempty"`
	// CheckCacheDropped whether to check that page cache was dropped on live migration. May cause noticeable delays if the image
	CheckCacheDropped *bool `json:"x-check-cache-dropped,omitempty"`
}

// BlockdevOptionsNull
//
// Driver specific block device options for the null backend.
type BlockdevOptionsNull struct {
	// Size size of the device in bytes.
	Size *int64 `json:"size,omitempty"`
	// LatencyNs emulated latency (in nanoseconds) in processing requests. Default to zero which completes requests immediately. (Since 2.4)
	LatencyNs *uint64 `json:"latency-ns,omitempty"`
	// ReadZeroes if true, reads from the device produce zeroes; if false, the buffer is left unchanged.
	ReadZeroes *bool `json:"read-zeroes,omitempty"`
}

// BlockdevOptionsNVMe
//
// Driver specific block device options for the NVMe backend.
type BlockdevOptionsNVMe struct {
	// Device PCI controller address of the NVMe device in format
	Device string `json:"device"`
	// Namespace namespace number of the device, starting from 1. Note that the PCI @device must have been unbound from any host kernel driver before instructing QEMU to add the blockdev.
	Namespace int64 `json:"namespace"`
}

// BlockdevOptionsVVFAT
//
// Driver specific block device options for the vvfat protocol.
type BlockdevOptionsVVFAT struct {
	// Dir directory to be exported as FAT image
	Dir string `json:"dir"`
	// FatType FAT type: 12, 16 or 32
	FatType *int64 `json:"fat-type,omitempty"`
	// Floppy whether to export a floppy image (true) or partitioned hard disk (false; default)
	Floppy *bool `json:"floppy,omitempty"`
	// Label set the volume label, limited to 11 bytes. FAT16 and FAT32 traditionally have some restrictions on labels, which are ignored by most operating systems. Defaults to "QEMU VVFAT". (since 2.4)
	Label *string `json:"label,omitempty"`
	// Rw whether to allow write operations (default: false)
	Rw *bool `json:"rw,omitempty"`
}

// BlockdevOptionsGenericFormat
//
// Driver specific block device options for image format that have no option besides their data source.
type BlockdevOptionsGenericFormat struct {
	// File reference to or definition of the data source block device
	File BlockdevRef `json:"file"`
}

// BlockdevOptionsLUKS
//
// Driver specific block device options for LUKS.
type BlockdevOptionsLUKS struct {
	BlockdevOptionsGenericFormat

	// KeySecret the ID of a QCryptoSecret object providing the decryption key (since 2.6). Mandatory except when doing a metadata-only probe of the image.
	KeySecret *string `json:"key-secret,omitempty"`
	// Header block device holding a detached LUKS header. (since 9.0)
	Header *BlockdevRef `json:"header,omitempty"`
}

// BlockdevOptionsGenericCOWFormat
//
// Driver specific block device options for image format that have no option besides their data source and an optional backing file.
type BlockdevOptionsGenericCOWFormat struct {
	BlockdevOptionsGenericFormat

	// Backing reference to or definition of the backing file block device, null disables the backing file entirely. Defaults to the backing file stored the image file.
	Backing *BlockdevRefOrNull `json:"backing,omitempty"`
}

// Qcow2OverlapCheckMode General overlap check modes.
type Qcow2OverlapCheckMode string

const (
	// Qcow2OverlapCheckModeNone Do not perform any checks
	Qcow2OverlapCheckModeNone Qcow2OverlapCheckMode = "none"
	// Qcow2OverlapCheckModeConstant Perform only checks which can be done in constant time and without reading anything from disk
	Qcow2OverlapCheckModeConstant Qcow2OverlapCheckMode = "constant"
	// Qcow2OverlapCheckModeCached Perform only checks which can be done without reading anything from disk
	Qcow2OverlapCheckModeCached Qcow2OverlapCheckMode = "cached"
	// Qcow2OverlapCheckModeAll Perform all available overlap checks
	Qcow2OverlapCheckModeAll Qcow2OverlapCheckMode = "all"
)

// Qcow2OverlapCheckFlags
//
// Structure of flags for each metadata structure. Setting a field to 'true' makes qemu guard that structure against unintended overwriting. The default value is chosen according to the template given.
type Qcow2OverlapCheckFlags struct {
	// Template Specifies a template mode which can be adjusted using the other flags, defaults to 'cached'
	Template      *Qcow2OverlapCheckMode `json:"template,omitempty"`
	MainHeader    *bool                  `json:"main-header,omitempty"`
	ActiveL1      *bool                  `json:"active-l1,omitempty"`
	ActiveL2      *bool                  `json:"active-l2,omitempty"`
	RefcountTable *bool                  `json:"refcount-table,omitempty"`
	RefcountBlock *bool                  `json:"refcount-block,omitempty"`
	SnapshotTable *bool                  `json:"snapshot-table,omitempty"`
	InactiveL1    *bool                  `json:"inactive-l1,omitempty"`
	InactiveL2    *bool                  `json:"inactive-l2,omitempty"`
	// BitmapDirectory since 3.0
	BitmapDirectory *bool `json:"bitmap-directory,omitempty"`
}

// Qcow2OverlapChecks
//
// Specifies which metadata structures should be guarded against unintended overwriting.
type Qcow2OverlapChecks struct {
	// Flags set of flags for separate specification of each metadata structure type
	Flags *Qcow2OverlapCheckFlags `json:"-"`
	// Mode named mode which chooses a specific set of flags
	Mode *Qcow2OverlapCheckMode `json:"-"`
}

func (a Qcow2OverlapChecks) MarshalJSON() ([]byte, error) {
	switch {
	case a.Flags != nil:
		return json.Marshal(a.Flags)
	case a.Mode != nil:
		return json.Marshal(a.Mode)
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockdevQcowEncryptionFormat
type BlockdevQcowEncryptionFormat string

const (
	// BlockdevQcowEncryptionFormatAes AES-CBC with plain64 initialization vectors
	BlockdevQcowEncryptionFormatAes BlockdevQcowEncryptionFormat = "aes"
)

// BlockdevQcowEncryption
type BlockdevQcowEncryption struct {
	// Discriminator: format

	// Format encryption format
	Format BlockdevQcowEncryptionFormat `json:"format"`

	Aes *QCryptoBlockOptionsQCow `json:"-"`
}

func (u BlockdevQcowEncryption) MarshalJSON() ([]byte, error) {
	switch u.Format {
	case "aes":
		if u.Aes == nil {
			return nil, fmt.Errorf("expected Aes to be set")
		}

		return json.Marshal(struct {
			Format BlockdevQcowEncryptionFormat `json:"format"`
			*QCryptoBlockOptionsQCow
		}{
			Format:                  u.Format,
			QCryptoBlockOptionsQCow: u.Aes,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockdevOptionsQcow
//
// Driver specific block device options for qcow.
type BlockdevOptionsQcow struct {
	BlockdevOptionsGenericCOWFormat

	// Encrypt Image decryption options. Mandatory for encrypted images, except when doing a metadata-only probe of the image.
	Encrypt *BlockdevQcowEncryption `json:"encrypt,omitempty"`
}

// BlockdevQcow2EncryptionFormat
type BlockdevQcow2EncryptionFormat string

const (
	// BlockdevQcow2EncryptionFormatAes AES-CBC with plain64 initialization vectors
	BlockdevQcow2EncryptionFormatAes  BlockdevQcow2EncryptionFormat = "aes"
	BlockdevQcow2EncryptionFormatLuks BlockdevQcow2EncryptionFormat = "luks"
)

// BlockdevQcow2Encryption
type BlockdevQcow2Encryption struct {
	// Discriminator: format

	// Format encryption format
	Format BlockdevQcow2EncryptionFormat `json:"format"`

	Aes  *QCryptoBlockOptionsQCow `json:"-"`
	Luks *QCryptoBlockOptionsLUKS `json:"-"`
}

func (u BlockdevQcow2Encryption) MarshalJSON() ([]byte, error) {
	switch u.Format {
	case "aes":
		if u.Aes == nil {
			return nil, fmt.Errorf("expected Aes to be set")
		}

		return json.Marshal(struct {
			Format BlockdevQcow2EncryptionFormat `json:"format"`
			*QCryptoBlockOptionsQCow
		}{
			Format:                  u.Format,
			QCryptoBlockOptionsQCow: u.Aes,
		})
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			Format BlockdevQcow2EncryptionFormat `json:"format"`
			*QCryptoBlockOptionsLUKS
		}{
			Format:                  u.Format,
			QCryptoBlockOptionsLUKS: u.Luks,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockdevOptionsPreallocate
//
// Filter driver intended to be inserted between format and protocol node and do preallocation in protocol node on write.
type BlockdevOptionsPreallocate struct {
	BlockdevOptionsGenericFormat

	// PreallocAlign on preallocation, align file length to this number, default 1048576 (1M)
	PreallocAlign *int64 `json:"prealloc-align,omitempty"`
	// PreallocSize how much to preallocate, default 134217728 (128M)
	PreallocSize *int64 `json:"prealloc-size,omitempty"`
}

// BlockdevOptionsQcow2
//
// Driver specific block device options for qcow2.
type BlockdevOptionsQcow2 struct {
	BlockdevOptionsGenericCOWFormat

	// LazyRefcounts whether to enable the lazy refcounts feature (default is taken from the image file)
	LazyRefcounts *bool `json:"lazy-refcounts,omitempty"`
	// PassDiscardRequest whether discard requests to the qcow2 device should be forwarded to the data source
	PassDiscardRequest *bool `json:"pass-discard-request,omitempty"`
	// PassDiscardSnapshot whether discard requests for the data source should be issued when a snapshot operation (e.g. deleting a snapshot) frees clusters in the qcow2 file
	PassDiscardSnapshot *bool `json:"pass-discard-snapshot,omitempty"`
	// PassDiscardOther whether discard requests for the data source should be issued on other occasions where a cluster gets freed
	PassDiscardOther *bool `json:"pass-discard-other,omitempty"`
	// DiscardNoUnref when enabled, data clusters will remain preallocated when they are no longer used, e.g. because they are discarded or converted to zero clusters. As usual, whether the old data is discarded or kept on the protocol level (i.e. in the image file) depends on the setting of the pass-discard-request option. Keeping the clusters preallocated prevents qcow2 fragmentation that would otherwise be caused by freeing and re-allocating them later. Besides potential performance degradation, such fragmentation can lead to increased allocation of clusters past the end of the image file, resulting in image files whose file length can grow much larger than their guest disk size would suggest. If image file length is of concern (e.g. when storing qcow2 images directly on block devices), you should consider enabling this option. (since 8.1)
	DiscardNoUnref *bool `json:"discard-no-unref,omitempty"`
	// OverlapCheck which overlap checks to perform for writes to the image, defaults to 'cached' (since 2.2)
	OverlapCheck *Qcow2OverlapChecks `json:"overlap-check,omitempty"`
	// CacheSize the maximum total size of the L2 table and refcount block caches in bytes (since 2.2)
	CacheSize *int64 `json:"cache-size,omitempty"`
	// L2CacheSize the maximum size of the L2 table cache in bytes (since 2.2)
	L2CacheSize *int64 `json:"l2-cache-size,omitempty"`
	// L2CacheEntrySize the size of each entry in the L2 cache in bytes. It must be a power of two between 512 and the cluster size. The default value is the cluster size (since 2.12)
	L2CacheEntrySize *int64 `json:"l2-cache-entry-size,omitempty"`
	// RefcountCacheSize the maximum size of the refcount block cache in bytes (since 2.2)
	RefcountCacheSize *int64 `json:"refcount-cache-size,omitempty"`
	// CacheCleanInterval clean unused entries in the L2 and refcount caches. The interval is in seconds. The default value is 600 on supporting platforms, and 0 on other platforms. 0 disables this feature. (since 2.5)
	CacheCleanInterval *int64 `json:"cache-clean-interval,omitempty"`
	// Encrypt Image decryption options. Mandatory for encrypted images, except when doing a metadata-only probe of the image. (since 2.10)
	Encrypt *BlockdevQcow2Encryption `json:"encrypt,omitempty"`
	// DataFile reference to or definition of the external data file. This may only be specified for images that require an external data file. If it is not specified for such an image, the data file name is loaded from the image file. (since 4.0)
	DataFile *BlockdevRef `json:"data-file,omitempty"`
}

// SshHostKeyCheckMode
type SshHostKeyCheckMode string

const (
	// SshHostKeyCheckModeNone Don't check the host key at all
	SshHostKeyCheckModeNone SshHostKeyCheckMode = "none"
	// SshHostKeyCheckModeHash Compare the host key with a given hash
	SshHostKeyCheckModeHash SshHostKeyCheckMode = "hash"
	// SshHostKeyCheckModeKnownHosts Check the host key against the known_hosts file
	SshHostKeyCheckModeKnownHosts SshHostKeyCheckMode = "known_hosts"
)

// SshHostKeyCheckHashType
type SshHostKeyCheckHashType string

const (
	// SshHostKeyCheckHashTypeMd5 The given hash is an md5 hash
	SshHostKeyCheckHashTypeMd5 SshHostKeyCheckHashType = "md5"
	// SshHostKeyCheckHashTypeSha1 The given hash is an sha1 hash
	SshHostKeyCheckHashTypeSha1 SshHostKeyCheckHashType = "sha1"
	// SshHostKeyCheckHashTypeSha256 The given hash is an sha256 hash
	SshHostKeyCheckHashTypeSha256 SshHostKeyCheckHashType = "sha256"
)

// SshHostKeyHash
type SshHostKeyHash struct {
	// Type The hash algorithm used for the hash
	Type SshHostKeyCheckHashType `json:"type"`
	// Hash The expected hash value
	Hash string `json:"hash"`
}

// SshHostKeyCheck
type SshHostKeyCheck struct {
	// Discriminator: mode

	// Mode How to check the host key
	Mode SshHostKeyCheckMode `json:"mode"`

	Hash *SshHostKeyHash `json:"-"`
}

func (u SshHostKeyCheck) MarshalJSON() ([]byte, error) {
	switch u.Mode {
	case "hash":
		if u.Hash == nil {
			return nil, fmt.Errorf("expected Hash to be set")
		}

		return json.Marshal(struct {
			Mode SshHostKeyCheckMode `json:"mode"`
			*SshHostKeyHash
		}{
			Mode:           u.Mode,
			SshHostKeyHash: u.Hash,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockdevOptionsSsh
type BlockdevOptionsSsh struct {
	// Server host address
	Server InetSocketAddress `json:"server"`
	// Path path to the image on the host
	Path string `json:"path"`
	// User user as which to connect, defaults to current local user name
	User *string `json:"user,omitempty"`
	// HostKeyCheck Defines how and what to check the host key against
	HostKeyCheck *SshHostKeyCheck `json:"host-key-check,omitempty"`
}

// BlkdebugEvent Trigger events supported by blkdebug.
type BlkdebugEvent string

const (
	BlkdbgL1Update                 BlkdebugEvent = "l1_update"
	BlkdbgL1GrowAllocTable         BlkdebugEvent = "l1_grow_alloc_table"
	BlkdbgL1GrowWriteTable         BlkdebugEvent = "l1_grow_write_table"
	BlkdbgL1GrowActivateTable      BlkdebugEvent = "l1_grow_activate_table"
	BlkdbgL2Load                   BlkdebugEvent = "l2_load"
	BlkdbgL2Update                 BlkdebugEvent = "l2_update"
	BlkdbgL2UpdateCompressed       BlkdebugEvent = "l2_update_compressed"
	BlkdbgL2AllocCowRead           BlkdebugEvent = "l2_alloc_cow_read"
	BlkdbgL2AllocWrite             BlkdebugEvent = "l2_alloc_write"
	BlkdbgReadAio                  BlkdebugEvent = "read_aio"
	BlkdbgReadBackingAio           BlkdebugEvent = "read_backing_aio"
	BlkdbgReadCompressed           BlkdebugEvent = "read_compressed"
	BlkdbgWriteAio                 BlkdebugEvent = "write_aio"
	BlkdbgWriteCompressed          BlkdebugEvent = "write_compressed"
	BlkdbgVmstateLoad              BlkdebugEvent = "vmstate_load"
	BlkdbgVmstateSave              BlkdebugEvent = "vmstate_save"
	BlkdbgCowRead                  BlkdebugEvent = "cow_read"
	BlkdbgCowWrite                 BlkdebugEvent = "cow_write"
	BlkdbgReftableLoad             BlkdebugEvent = "reftable_load"
	BlkdbgReftableGrow             BlkdebugEvent = "reftable_grow"
	BlkdbgReftableUpdate           BlkdebugEvent = "reftable_update"
	BlkdbgRefblockLoad             BlkdebugEvent = "refblock_load"
	BlkdbgRefblockUpdate           BlkdebugEvent = "refblock_update"
	BlkdbgRefblockUpdatePart       BlkdebugEvent = "refblock_update_part"
	BlkdbgRefblockAlloc            BlkdebugEvent = "refblock_alloc"
	BlkdbgRefblockAllocHookup      BlkdebugEvent = "refblock_alloc_hookup"
	BlkdbgRefblockAllocWrite       BlkdebugEvent = "refblock_alloc_write"
	BlkdbgRefblockAllocWriteBlocks BlkdebugEvent = "refblock_alloc_write_blocks"
	BlkdbgRefblockAllocWriteTable  BlkdebugEvent = "refblock_alloc_write_table"
	BlkdbgRefblockAllocSwitchTable BlkdebugEvent = "refblock_alloc_switch_table"
	BlkdbgClusterAlloc             BlkdebugEvent = "cluster_alloc"
	BlkdbgClusterAllocBytes        BlkdebugEvent = "cluster_alloc_bytes"
	BlkdbgClusterFree              BlkdebugEvent = "cluster_free"
	BlkdbgFlushToOs                BlkdebugEvent = "flush_to_os"
	BlkdbgFlushToDisk              BlkdebugEvent = "flush_to_disk"
	BlkdbgPwritevRmwHead           BlkdebugEvent = "pwritev_rmw_head"
	BlkdbgPwritevRmwAfterHead      BlkdebugEvent = "pwritev_rmw_after_head"
	BlkdbgPwritevRmwTail           BlkdebugEvent = "pwritev_rmw_tail"
	BlkdbgPwritevRmwAfterTail      BlkdebugEvent = "pwritev_rmw_after_tail"
	BlkdbgPwritev                  BlkdebugEvent = "pwritev"
	BlkdbgPwritevZero              BlkdebugEvent = "pwritev_zero"
	BlkdbgPwritevDone              BlkdebugEvent = "pwritev_done"
	BlkdbgEmptyImagePrepare        BlkdebugEvent = "empty_image_prepare"
	// BlkdbgL1ShrinkWriteTable write zeros to the l1 table to shrink image. (since 2.11)
	BlkdbgL1ShrinkWriteTable BlkdebugEvent = "l1_shrink_write_table"
	// BlkdbgL1ShrinkFreeL2Clusters discard the l2 tables. (since 2.11)
	BlkdbgL1ShrinkFreeL2Clusters BlkdebugEvent = "l1_shrink_free_l2_clusters"
	// BlkdbgCorWrite a write due to copy-on-read (since 2.11)
	BlkdbgCorWrite BlkdebugEvent = "cor_write"
	// BlkdbgClusterAllocSpace an allocation of file space for a cluster (since 4.1)
	BlkdbgClusterAllocSpace BlkdebugEvent = "cluster_alloc_space"
	// BlkdbgNone triggers once at creation of the blkdebug node (since 4.1)
	BlkdbgNone BlkdebugEvent = "none"
)

// BlkdebugIOType Kinds of I/O that blkdebug can inject errors in.
type BlkdebugIOType string

const (
	// BlkdebugIoTypeRead .bdrv_co_preadv()
	BlkdebugIoTypeRead BlkdebugIOType = "read"
	// BlkdebugIoTypeWrite .bdrv_co_pwritev()
	BlkdebugIoTypeWrite BlkdebugIOType = "write"
	// BlkdebugIoTypeWriteZeroes .bdrv_co_pwrite_zeroes()
	BlkdebugIoTypeWriteZeroes BlkdebugIOType = "write-zeroes"
	// BlkdebugIoTypeDiscard .bdrv_co_pdiscard()
	BlkdebugIoTypeDiscard BlkdebugIOType = "discard"
	// BlkdebugIoTypeFlush .bdrv_co_flush_to_disk()
	BlkdebugIoTypeFlush BlkdebugIOType = "flush"
	// BlkdebugIoTypeBlockStatus .bdrv_co_block_status()
	BlkdebugIoTypeBlockStatus BlkdebugIOType = "block-status"
)

// BlkdebugInjectErrorOptions
//
// Describes a single error injection for blkdebug.
type BlkdebugInjectErrorOptions struct {
	// Event trigger event
	Event BlkdebugEvent `json:"event"`
	// State the state identifier blkdebug needs to be in to actually trigger the event; defaults to "any"
	State *int64 `json:"state,omitempty"`
	// Iotype the type of I/O operations on which this error should be injected; defaults to "all read, write, write-zeroes, discard,
	Iotype *BlkdebugIOType `json:"iotype,omitempty"`
	// Errno error identifier (errno) to be returned; defaults to EIO
	Errno *int64 `json:"errno,omitempty"`
	// Sector specifies the sector index which has to be affected in order to actually trigger the event; defaults to "any sector"
	Sector *int64 `json:"sector,omitempty"`
	// Once disables further events after this one has been triggered; defaults to false
	Once *bool `json:"once,omitempty"`
	// Immediately fail immediately; defaults to false
	Immediately *bool `json:"immediately,omitempty"`
}

// BlkdebugSetStateOptions
//
// Describes a single state-change event for blkdebug.
type BlkdebugSetStateOptions struct {
	// Event trigger event
	Event BlkdebugEvent `json:"event"`
	// State the current state identifier blkdebug needs to be in; defaults to "any"
	State *int64 `json:"state,omitempty"`
	// NewState the state identifier blkdebug is supposed to assume if this event is triggered
	NewState int64 `json:"new_state"`
}

// BlockdevOptionsBlkdebug
//
// Driver specific block device options for blkdebug.
type BlockdevOptionsBlkdebug struct {
	// Image underlying raw block device (or image file)
	Image BlockdevRef `json:"image"`
	// Config filename of the configuration file
	Config *string `json:"config,omitempty"`
	// Align required alignment for requests in bytes, must be positive power of 2, or 0 for default
	Align *int64 `json:"align,omitempty"`
	// MaxTransfer maximum size for I/O transfers in bytes, must be positive multiple of @align and of the underlying file's request alignment (but need not be a power of 2), or 0 for default (since 2.10)
	MaxTransfer *int32 `json:"max-transfer,omitempty"`
	// OptWriteZero preferred alignment for write zero requests in bytes, must be positive multiple of @align and of the underlying file's request alignment (but need not be a power of 2), or 0 for default (since 2.10)
	OptWriteZero *int32 `json:"opt-write-zero,omitempty"`
	// MaxWriteZero maximum size for write zero requests in bytes, must be positive multiple of @align, of @opt-write-zero, and of the underlying file's request alignment (but need not be a power of 2), or 0 for default (since 2.10)
	MaxWriteZero *int32 `json:"max-write-zero,omitempty"`
	// OptDiscard preferred alignment for discard requests in bytes, must be positive multiple of @align and of the underlying file's request alignment (but need not be a power of 2), or 0 for default (since 2.10)
	OptDiscard *int32 `json:"opt-discard,omitempty"`
	// MaxDiscard maximum size for discard requests in bytes, must be positive multiple of @align, of @opt-discard, and of the underlying file's request alignment (but need not be a power of 2), or 0 for default (since 2.10)
	MaxDiscard *int32 `json:"max-discard,omitempty"`
	// InjectError array of error injection descriptions
	InjectError []BlkdebugInjectErrorOptions `json:"inject-error,omitempty"`
	// SetState array of state-change descriptions
	SetState []BlkdebugSetStateOptions `json:"set-state,omitempty"`
	// TakeChildPerms Permissions to take on @image in addition to what is necessary anyway (which depends on how the blkdebug node is used). Defaults to none. (since 5.0)
	TakeChildPerms []BlockPermission `json:"take-child-perms,omitempty"`
	// UnshareChildPerms Permissions not to share on @image in addition to what cannot be shared anyway (which depends on how the blkdebug node is used). Defaults to none. (since 5.0)
	UnshareChildPerms []BlockPermission `json:"unshare-child-perms,omitempty"`
}

// BlockdevOptionsBlklogwrites
//
// Driver specific block device options for blklogwrites.
type BlockdevOptionsBlklogwrites struct {
	// File block device
	File BlockdevRef `json:"file"`
	// Log block device used to log writes to @file
	Log BlockdevRef `json:"log"`
	// LogSectorSize sector size used in logging writes to @file, determines granularity of offsets and sizes of writes
	LogSectorSize *uint32 `json:"log-sector-size,omitempty"`
	// LogAppend append to an existing log (default: false)
	LogAppend *bool `json:"log-append,omitempty"`
	// LogSuperUpdateInterval interval of write requests after which
	LogSuperUpdateInterval *uint64 `json:"log-super-update-interval,omitempty"`
}

// BlockdevOptionsBlkverify
//
// Driver specific block device options for blkverify.
type BlockdevOptionsBlkverify struct {
	// Test block device to be tested
	Test BlockdevRef `json:"test"`
	// Raw raw image used for verification
	Raw BlockdevRef `json:"raw"`
}

// BlockdevOptionsBlkreplay
//
// Driver specific block device options for blkreplay.
type BlockdevOptionsBlkreplay struct {
	// Image disk image which should be controlled with blkreplay
	Image BlockdevRef `json:"image"`
}

// QuorumReadPattern An enumeration of quorum read patterns.
type QuorumReadPattern string

const (
	// QuorumReadPatternQuorum read all the children and do a quorum vote on reads
	QuorumReadPatternQuorum QuorumReadPattern = "quorum"
	// QuorumReadPatternFifo read only from the first child that has not failed
	QuorumReadPatternFifo QuorumReadPattern = "fifo"
)

// BlockdevOptionsQuorum
//
// Driver specific block device options for Quorum
type BlockdevOptionsQuorum struct {
	// Blkverify true if the driver must print content mismatch set to false by default
	Blkverify *bool `json:"blkverify,omitempty"`
	// Children the children block devices to use
	Children []BlockdevRef `json:"children"`
	// VoteThreshold the vote limit under which a read will fail
	VoteThreshold int64 `json:"vote-threshold"`
	// RewriteCorrupted rewrite corrupted data when quorum is reached (Since 2.1)
	RewriteCorrupted *bool `json:"rewrite-corrupted,omitempty"`
	// ReadPattern choose read pattern and set to quorum by default (Since 2.2)
	ReadPattern *QuorumReadPattern `json:"read-pattern,omitempty"`
}

// BlockdevOptionsGluster
//
// Driver specific block device options for Gluster
type BlockdevOptionsGluster struct {
	// Volume name of gluster volume where VM image resides
	Volume string `json:"volume"`
	// Path absolute path to image file in gluster volume
	Path string `json:"path"`
	// Server gluster servers description
	Server []SocketAddress `json:"server"`
	// Debug libgfapi log level (default '4' which is Error) (Since 2.8)
	Debug *int64 `json:"debug,omitempty"`
	// Logfile libgfapi log file (default /dev/stderr) (Since 2.8)
	Logfile *string `json:"logfile,omitempty"`
}

// BlockdevOptionsIoUring
//
// Driver specific block device options for the io_uring backend.
type BlockdevOptionsIoUring struct {
	// Filename path to the image file
	Filename string `json:"filename"`
}

// BlockdevOptionsNvmeIoUring
//
// Driver specific block device options for the nvme-io_uring backend.
type BlockdevOptionsNvmeIoUring struct {
	// Path path to the NVMe namespace's character device (e.g. /dev/ng0n1).
	Path string `json:"path"`
}

// BlockdevOptionsVirtioBlkVfioPci
//
// Driver specific block device options for the virtio-blk-vfio-pci backend.
type BlockdevOptionsVirtioBlkVfioPci struct {
	// Path path to the PCI device's sysfs directory (e.g.
	Path string `json:"path"`
}

// BlockdevOptionsVirtioBlkVhostUser
//
// Driver specific block device options for the virtio-blk-vhost-user backend.
type BlockdevOptionsVirtioBlkVhostUser struct {
	// Path path to the vhost-user UNIX domain socket.
	Path string `json:"path"`
}

// BlockdevOptionsVirtioBlkVhostVdpa
//
// Driver specific block device options for the virtio-blk-vhost-vdpa backend.
type BlockdevOptionsVirtioBlkVhostVdpa struct {
	// Path path to the vhost-vdpa character device.
	Path string `json:"path"`
}

// IscsiTransport An enumeration of libiscsi transport types
type IscsiTransport string

const (
	IscsiTransportTcp  IscsiTransport = "tcp"
	IscsiTransportIser IscsiTransport = "iser"
)

// IscsiHeaderDigest An enumeration of header digests supported by libiscsi
type IscsiHeaderDigest string

const (
	QapiIscsiHeaderDigestCrc32c     IscsiHeaderDigest = "crc32c"
	QapiIscsiHeaderDigestNone       IscsiHeaderDigest = "none"
	QapiIscsiHeaderDigestCrc32cNone IscsiHeaderDigest = "crc32c-none"
	QapiIscsiHeaderDigestNoneCrc32c IscsiHeaderDigest = "none-crc32c"
)

// BlockdevOptionsIscsi
//
// Driver specific block device options for iscsi
type BlockdevOptionsIscsi struct {
	// Transport The iscsi transport type
	Transport IscsiTransport `json:"transport"`
	// Portal The address of the iscsi portal
	Portal string `json:"portal"`
	// Target The target iqn name
	Target string `json:"target"`
	// Lun LUN to connect to. Defaults to 0.
	Lun *int64 `json:"lun,omitempty"`
	// User User name to log in with. If omitted, no CHAP authentication is performed.
	User *string `json:"user,omitempty"`
	// PasswordSecret The ID of a QCryptoSecret object providing the password for the login. This option is required if @user is specified.
	PasswordSecret *string `json:"password-secret,omitempty"`
	// InitiatorName The iqn name we want to identify to the target as. If this option is not specified, an initiator name is generated automatically.
	InitiatorName *string `json:"initiator-name,omitempty"`
	// HeaderDigest The desired header digest. Defaults to none-crc32c.
	HeaderDigest *IscsiHeaderDigest `json:"header-digest,omitempty"`
	// Timeout Timeout in seconds after which a request will timeout. 0 means no timeout and is the default.
	Timeout *int64 `json:"timeout,omitempty"`
}

// RbdAuthMode
type RbdAuthMode string

const (
	RbdAuthModeCephx RbdAuthMode = "cephx"
	RbdAuthModeNone  RbdAuthMode = "none"
)

// RbdImageEncryptionFormat
type RbdImageEncryptionFormat string

const (
	RbdImageEncryptionFormatLuks  RbdImageEncryptionFormat = "luks"
	RbdImageEncryptionFormatLuks2 RbdImageEncryptionFormat = "luks2"
	// RbdImageEncryptionFormatLuksAny Used for opening either luks or luks2 (Since 8.0)
	RbdImageEncryptionFormatLuksAny RbdImageEncryptionFormat = "luks-any"
)

// RbdEncryptionOptionsLUKSBase
type RbdEncryptionOptionsLUKSBase struct {
	// KeySecret ID of a QCryptoSecret object providing a passphrase for unlocking the encryption
	KeySecret string `json:"key-secret"`
}

// RbdEncryptionCreateOptionsLUKSBase
type RbdEncryptionCreateOptionsLUKSBase struct {
	RbdEncryptionOptionsLUKSBase

	// CipherAlg The encryption algorithm
	CipherAlg *QCryptoCipherAlgorithm `json:"cipher-alg,omitempty"`
}

// RbdEncryptionOptionsLUKS
type RbdEncryptionOptionsLUKS struct {
	RbdEncryptionOptionsLUKSBase
}

// RbdEncryptionOptionsLUKS2
type RbdEncryptionOptionsLUKS2 struct {
	RbdEncryptionOptionsLUKSBase
}

// RbdEncryptionOptionsLUKSAny
type RbdEncryptionOptionsLUKSAny struct {
	RbdEncryptionOptionsLUKSBase
}

// RbdEncryptionCreateOptionsLUKS
type RbdEncryptionCreateOptionsLUKS struct {
	RbdEncryptionCreateOptionsLUKSBase
}

// RbdEncryptionCreateOptionsLUKS2
type RbdEncryptionCreateOptionsLUKS2 struct {
	RbdEncryptionCreateOptionsLUKSBase
}

// RbdEncryptionOptions
type RbdEncryptionOptions struct {
	// Discriminator: format

	// Format Encryption format.
	Format RbdImageEncryptionFormat `json:"format"`
	// Parent Parent image encryption options (for cloned images). Can be left unspecified if this cloned image is encrypted using the same format and secret as its parent image (i.e. not explicitly formatted) or if its parent image is not encrypted. (Since 8.0)
	Parent *RbdEncryptionOptions `json:"parent,omitempty"`

	Luks    *RbdEncryptionOptionsLUKS    `json:"-"`
	Luks2   *RbdEncryptionOptionsLUKS2   `json:"-"`
	LuksAny *RbdEncryptionOptionsLUKSAny `json:"-"`
}

func (u RbdEncryptionOptions) MarshalJSON() ([]byte, error) {
	switch u.Format {
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			Format RbdImageEncryptionFormat `json:"format"`
			Parent *RbdEncryptionOptions    `json:"parent,omitempty"`
			*RbdEncryptionOptionsLUKS
		}{
			Format:                   u.Format,
			Parent:                   u.Parent,
			RbdEncryptionOptionsLUKS: u.Luks,
		})
	case "luks2":
		if u.Luks2 == nil {
			return nil, fmt.Errorf("expected Luks2 to be set")
		}

		return json.Marshal(struct {
			Format RbdImageEncryptionFormat `json:"format"`
			Parent *RbdEncryptionOptions    `json:"parent,omitempty"`
			*RbdEncryptionOptionsLUKS2
		}{
			Format:                    u.Format,
			Parent:                    u.Parent,
			RbdEncryptionOptionsLUKS2: u.Luks2,
		})
	case "luks-any":
		if u.LuksAny == nil {
			return nil, fmt.Errorf("expected LuksAny to be set")
		}

		return json.Marshal(struct {
			Format RbdImageEncryptionFormat `json:"format"`
			Parent *RbdEncryptionOptions    `json:"parent,omitempty"`
			*RbdEncryptionOptionsLUKSAny
		}{
			Format:                      u.Format,
			Parent:                      u.Parent,
			RbdEncryptionOptionsLUKSAny: u.LuksAny,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// RbdEncryptionCreateOptions
type RbdEncryptionCreateOptions struct {
	// Discriminator: format

	// Format Encryption format.
	Format RbdImageEncryptionFormat `json:"format"`

	Luks  *RbdEncryptionCreateOptionsLUKS  `json:"-"`
	Luks2 *RbdEncryptionCreateOptionsLUKS2 `json:"-"`
}

func (u RbdEncryptionCreateOptions) MarshalJSON() ([]byte, error) {
	switch u.Format {
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			Format RbdImageEncryptionFormat `json:"format"`
			*RbdEncryptionCreateOptionsLUKS
		}{
			Format:                         u.Format,
			RbdEncryptionCreateOptionsLUKS: u.Luks,
		})
	case "luks2":
		if u.Luks2 == nil {
			return nil, fmt.Errorf("expected Luks2 to be set")
		}

		return json.Marshal(struct {
			Format RbdImageEncryptionFormat `json:"format"`
			*RbdEncryptionCreateOptionsLUKS2
		}{
			Format:                          u.Format,
			RbdEncryptionCreateOptionsLUKS2: u.Luks2,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockdevOptionsRbd
type BlockdevOptionsRbd struct {
	// Pool Ceph pool name.
	Pool string `json:"pool"`
	// Namespace Rados namespace name in the Ceph pool. (Since 5.0)
	Namespace *string `json:"namespace,omitempty"`
	// Image Image name in the Ceph pool.
	Image string `json:"image"`
	// Conf path to Ceph configuration file. Values in the configuration file will be overridden by options specified via QAPI.
	Conf *string `json:"conf,omitempty"`
	// Snapshot Ceph snapshot name.
	Snapshot *string `json:"snapshot,omitempty"`
	// Encrypt Image encryption options. (Since 6.1)
	Encrypt *RbdEncryptionOptions `json:"encrypt,omitempty"`
	// User Ceph id name.
	User *string `json:"user,omitempty"`
	// AuthClientRequired Acceptable authentication modes. This maps to Ceph configuration option "auth_client_required". (Since 3.0)
	AuthClientRequired []RbdAuthMode `json:"auth-client-required,omitempty"`
	// KeySecret ID of a QCryptoSecret object providing a key for cephx authentication. This maps to Ceph configuration option "key". (Since 3.0)
	KeySecret *string `json:"key-secret,omitempty"`
	// Server Monitor host address and port. This maps to the "mon_host" Ceph option.
	Server []InetSocketAddressBase `json:"server,omitempty"`
}

// ReplicationMode An enumeration of replication modes.
type ReplicationMode string

const (
	// ReplicationModePrimary Primary mode, the vm's state will be sent to secondary QEMU.
	ReplicationModePrimary ReplicationMode = "primary"
	// ReplicationModeSecondary Secondary mode, receive the vm's state from primary QEMU.
	ReplicationModeSecondary ReplicationMode = "secondary"
)

// BlockdevOptionsReplication
//
// Driver specific block device options for replication
type BlockdevOptionsReplication struct {
	BlockdevOptionsGenericFormat

	// Mode the replication mode
	Mode ReplicationMode `json:"mode"`
	// TopId In secondary mode, node name or device ID of the root node who owns the replication node chain. Must not be given in primary mode.
	TopId *string `json:"top-id,omitempty"`
}

// NFSTransport An enumeration of NFS transport types
type NFSTransport string

const (
	// NFSTransportInet TCP transport
	NFSTransportInet NFSTransport = "inet"
)

// NFSServer
//
// Captures the address of the socket
type NFSServer struct {
	// Type transport type used for NFS (only TCP supported)
	Type NFSTransport `json:"type"`
	// Host host address for NFS server
	Host string `json:"host"`
}

// BlockdevOptionsNfs
//
// Driver specific block device option for NFS
type BlockdevOptionsNfs struct {
	// Server host address
	Server NFSServer `json:"server"`
	// Path path of the image on the host
	Path string `json:"path"`
	// User UID value to use when talking to the server (defaults to 65534 on Windows and getuid() on unix)
	User *int64 `json:"user,omitempty"`
	// Group GID value to use when talking to the server (defaults to 65534 on Windows and getgid() in unix)
	Group *int64 `json:"group,omitempty"`
	// TcpSynCount number of SYNs during the session establishment (defaults to libnfs default)
	TcpSynCount *int64 `json:"tcp-syn-count,omitempty"`
	// ReadaheadSize set the readahead size in bytes (defaults to libnfs default)
	ReadaheadSize *int64 `json:"readahead-size,omitempty"`
	// PageCacheSize set the pagecache size in bytes (defaults to libnfs default)
	PageCacheSize *int64 `json:"page-cache-size,omitempty"`
	// Debug set the NFS debug level (max 2) (defaults to libnfs default)
	Debug *int64 `json:"debug,omitempty"`
}

// BlockdevOptionsCurlBase
//
// Driver specific block device options shared by all protocols supported by the curl backend.
type BlockdevOptionsCurlBase struct {
	// Url URL of the image file
	Url string `json:"url"`
	// Readahead Size of the read-ahead cache; must be a multiple of 512 (defaults to 256 kB)
	Readahead *int64 `json:"readahead,omitempty"`
	// Timeout Timeout for connections, in seconds (defaults to 5)
	Timeout *int64 `json:"timeout,omitempty"`
	// Username Username for authentication (defaults to none)
	Username *string `json:"username,omitempty"`
	// PasswordSecret ID of a QCryptoSecret object providing a password for authentication (defaults to no password)
	PasswordSecret *string `json:"password-secret,omitempty"`
	// ProxyUsername Username for proxy authentication (defaults to none)
	ProxyUsername *string `json:"proxy-username,omitempty"`
	// ProxyPasswordSecret ID of a QCryptoSecret object providing a password for proxy authentication (defaults to no password)
	ProxyPasswordSecret *string `json:"proxy-password-secret,omitempty"`
}

// BlockdevOptionsCurlHttp
//
// Driver specific block device options for HTTP connections over the
type BlockdevOptionsCurlHttp struct {
	BlockdevOptionsCurlBase

	// Cookie List of cookies to set; format is "name1=content1; name2=content2;" as explained by CURLOPT_COOKIE(3). Defaults to no cookies.
	Cookie *string `json:"cookie,omitempty"`
	// CookieSecret ID of a QCryptoSecret object providing the cookie data in a secure way. See @cookie for the format. (since 2.10)
	CookieSecret *string `json:"cookie-secret,omitempty"`
}

// BlockdevOptionsCurlHttps
//
// Driver specific block device options for HTTPS connections over the
type BlockdevOptionsCurlHttps struct {
	BlockdevOptionsCurlBase

	// Cookie List of cookies to set; format is "name1=content1; name2=content2;" as explained by CURLOPT_COOKIE(3). Defaults to no cookies.
	Cookie *string `json:"cookie,omitempty"`
	// Sslverify Whether to verify the SSL certificate's validity (defaults to true)
	Sslverify *bool `json:"sslverify,omitempty"`
	// CookieSecret ID of a QCryptoSecret object providing the cookie data in a secure way. See @cookie for the format. (since 2.10)
	CookieSecret *string `json:"cookie-secret,omitempty"`
}

// BlockdevOptionsCurlFtp
//
// Driver specific block device options for FTP connections over the
type BlockdevOptionsCurlFtp struct {
	BlockdevOptionsCurlBase
}

// BlockdevOptionsCurlFtps
//
// Driver specific block device options for FTPS connections over the
type BlockdevOptionsCurlFtps struct {
	BlockdevOptionsCurlBase

	// Sslverify Whether to verify the SSL certificate's validity (defaults to true)
	Sslverify *bool `json:"sslverify,omitempty"`
}

// BlockdevOptionsNbd
//
// Driver specific block device options for NBD.
type BlockdevOptionsNbd struct {
	// Server NBD server address
	Server SocketAddress `json:"server"`
	// Export export name
	Export *string `json:"export,omitempty"`
	// TlsCreds TLS credentials ID
	TlsCreds *string `json:"tls-creds,omitempty"`
	// TlsHostname TLS hostname override for certificate validation (Since 7.0)
	TlsHostname *string `json:"tls-hostname,omitempty"`
	// DirtyBitmap A metadata context name such as
	DirtyBitmap *string `json:"x-dirty-bitmap,omitempty"`
	// ReconnectDelay On an unexpected disconnect, the nbd client tries to connect again until succeeding or encountering a serious error. During the first @reconnect-delay seconds, all requests are paused and will be rerun on a successful reconnect. After that time, any delayed requests and all future requests before a successful reconnect will immediately fail. Default 0 (Since 4.2)
	ReconnectDelay *uint32 `json:"reconnect-delay,omitempty"`
	// OpenTimeout In seconds. If zero, the nbd driver tries the connection only once, and fails to open if the connection fails. If non-zero, the nbd driver will repeat connection attempts until successful or until @open-timeout seconds have elapsed. Default 0 (Since 7.0)
	OpenTimeout *uint32 `json:"open-timeout,omitempty"`
}

// BlockdevOptionsRaw
//
// Driver specific block device options for the raw driver.
type BlockdevOptionsRaw struct {
	BlockdevOptionsGenericFormat

	// Offset position where the block device starts
	Offset *int64 `json:"offset,omitempty"`
	// Size the assumed size of the device
	Size *int64 `json:"size,omitempty"`
}

// BlockdevOptionsThrottle
//
// Driver specific block device options for the throttle driver
type BlockdevOptionsThrottle struct {
	// ThrottleGroup the name of the throttle-group object to use. It must already exist.
	ThrottleGroup string `json:"throttle-group"`
	// File reference to or definition of the data source block device
	File BlockdevRef `json:"file"`
}

// BlockdevOptionsCor
//
// Driver specific block device options for the copy-on-read driver.
type BlockdevOptionsCor struct {
	BlockdevOptionsGenericFormat

	// Bottom The name of a non-filter node (allocation-bearing layer) that limits the COR operations in the backing chain (inclusive), so that no data below this node will be copied by this filter. If option is absent, the limit is not applied, so that data from all backing layers may be copied.
	Bottom *string `json:"bottom,omitempty"`
}

// OnCbwError An enumeration of possible behaviors for copy-before-write operation failures.
type OnCbwError string

const (
	// OnCbwErrorBreakGuestWrite report the error to the guest. This way, the guest will not be able to overwrite areas that cannot be backed up, so the backup process remains valid.
	OnCbwErrorBreakGuestWrite OnCbwError = "break-guest-write"
	// OnCbwErrorBreakSnapshot continue guest write. Doing so will make the provided snapshot state invalid and any backup or export process based on it will finally fail.
	OnCbwErrorBreakSnapshot OnCbwError = "break-snapshot"
)

// BlockdevOptionsCbw
//
// Driver specific block device options for the copy-before-write
type BlockdevOptionsCbw struct {
	BlockdevOptionsGenericFormat

	// Target The target for copy-before-write operations.
	Target BlockdevRef `json:"target"`
	// Bitmap If specified, copy-before-write filter will do copy-before-write operations only for dirty regions of the bitmap. Bitmap size must be equal to length of file and target child of the filter. Note also, that bitmap is used only to initialize internal bitmap of the process, so further modifications (or removing) of specified bitmap doesn't influence the filter. (Since 7.0)
	Bitmap *BlockDirtyBitmap `json:"bitmap,omitempty"`
	// OnCbwError Behavior on failure of copy-before-write operation. Default is @break-guest-write. (Since 7.1)
	OnCbwError *OnCbwError `json:"on-cbw-error,omitempty"`
	// CbwTimeout Zero means no limit. Non-zero sets the timeout in seconds for copy-before-write operation. When a timeout occurs, the respective copy-before-write operation will fail, and the @on-cbw-error parameter will decide how this failure is handled. Default 0. (Since 7.1)
	CbwTimeout *uint32 `json:"cbw-timeout,omitempty"`
}

// BlockdevOptions
//
// Options for creating a block device. Many options are available for
type BlockdevOptions struct {
	// Discriminator: driver

	// Driver block driver name
	Driver BlockdevDriver `json:"driver"`
	// NodeName the node name of the new node (Since 2.0). This option is required on the top level of blockdev-add. Valid node names start with an alphabetic character and may contain only alphanumeric characters, '-', '.' and '_'. Their maximum length is 31 characters.
	NodeName *string `json:"node-name,omitempty"`
	// Discard discard-related options (default: ignore)
	Discard *BlockdevDiscardOptions `json:"discard,omitempty"`
	// Cache cache-related options
	Cache *BlockdevCacheOptions `json:"cache,omitempty"`
	// ReadOnly whether the block device should be read-only (default: false). Note that some block drivers support only read-only access, either generally or in certain configurations. In this case, the default value does not work and the option must be specified explicitly.
	ReadOnly *bool `json:"read-only,omitempty"`
	// AutoReadOnly if true and @read-only is false, QEMU may automatically decide not to open the image read-write as requested, but fall back to read-only instead (and switch between the modes later), e.g. depending on whether the image file is writable or whether a writing user is attached to the
	AutoReadOnly *bool `json:"auto-read-only,omitempty"`
	// ForceShare force share all permission on added nodes. Requires read-only=true. (Since 2.10)
	ForceShare *bool `json:"force-share,omitempty"`
	// DetectZeroes detect and optimize zero writes (Since 2.1)
	DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`

	Blkdebug           *BlockdevOptionsBlkdebug           `json:"-"`
	Blklogwrites       *BlockdevOptionsBlklogwrites       `json:"-"`
	Blkverify          *BlockdevOptionsBlkverify          `json:"-"`
	Blkreplay          *BlockdevOptionsBlkreplay          `json:"-"`
	Bochs              *BlockdevOptionsGenericFormat      `json:"-"`
	Cloop              *BlockdevOptionsGenericFormat      `json:"-"`
	Compress           *BlockdevOptionsGenericFormat      `json:"-"`
	CopyBeforeWrite    *BlockdevOptionsCbw                `json:"-"`
	CopyOnRead         *BlockdevOptionsCor                `json:"-"`
	Dmg                *BlockdevOptionsGenericFormat      `json:"-"`
	File               *BlockdevOptionsFile               `json:"-"`
	Ftp                *BlockdevOptionsCurlFtp            `json:"-"`
	Ftps               *BlockdevOptionsCurlFtps           `json:"-"`
	Gluster            *BlockdevOptionsGluster            `json:"-"`
	HostCdrom          *BlockdevOptionsFile               `json:"-"`
	HostDevice         *BlockdevOptionsFile               `json:"-"`
	Http               *BlockdevOptionsCurlHttp           `json:"-"`
	Https              *BlockdevOptionsCurlHttps          `json:"-"`
	IoUring            *BlockdevOptionsIoUring            `json:"-"`
	Iscsi              *BlockdevOptionsIscsi              `json:"-"`
	Luks               *BlockdevOptionsLUKS               `json:"-"`
	Nbd                *BlockdevOptionsNbd                `json:"-"`
	Nfs                *BlockdevOptionsNfs                `json:"-"`
	NullAio            *BlockdevOptionsNull               `json:"-"`
	NullCo             *BlockdevOptionsNull               `json:"-"`
	Nvme               *BlockdevOptionsNVMe               `json:"-"`
	NvmeIoUring        *BlockdevOptionsNvmeIoUring        `json:"-"`
	Parallels          *BlockdevOptionsGenericFormat      `json:"-"`
	Preallocate        *BlockdevOptionsPreallocate        `json:"-"`
	Qcow2              *BlockdevOptionsQcow2              `json:"-"`
	Qcow               *BlockdevOptionsQcow               `json:"-"`
	Qed                *BlockdevOptionsGenericCOWFormat   `json:"-"`
	Quorum             *BlockdevOptionsQuorum             `json:"-"`
	Raw                *BlockdevOptionsRaw                `json:"-"`
	Rbd                *BlockdevOptionsRbd                `json:"-"`
	Replication        *BlockdevOptionsReplication        `json:"-"`
	SnapshotAccess     *BlockdevOptionsGenericFormat      `json:"-"`
	Ssh                *BlockdevOptionsSsh                `json:"-"`
	Throttle           *BlockdevOptionsThrottle           `json:"-"`
	Vdi                *BlockdevOptionsGenericFormat      `json:"-"`
	Vhdx               *BlockdevOptionsGenericFormat      `json:"-"`
	VirtioBlkVfioPci   *BlockdevOptionsVirtioBlkVfioPci   `json:"-"`
	VirtioBlkVhostUser *BlockdevOptionsVirtioBlkVhostUser `json:"-"`
	VirtioBlkVhostVdpa *BlockdevOptionsVirtioBlkVhostVdpa `json:"-"`
	Vmdk               *BlockdevOptionsGenericCOWFormat   `json:"-"`
	Vpc                *BlockdevOptionsGenericFormat      `json:"-"`
	Vvfat              *BlockdevOptionsVVFAT              `json:"-"`
}

func (u BlockdevOptions) MarshalJSON() ([]byte, error) {
	switch u.Driver {
	case "blkdebug":
		if u.Blkdebug == nil {
			return nil, fmt.Errorf("expected Blkdebug to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsBlkdebug
		}{
			Driver:                  u.Driver,
			NodeName:                u.NodeName,
			Discard:                 u.Discard,
			Cache:                   u.Cache,
			ReadOnly:                u.ReadOnly,
			AutoReadOnly:            u.AutoReadOnly,
			ForceShare:              u.ForceShare,
			DetectZeroes:            u.DetectZeroes,
			BlockdevOptionsBlkdebug: u.Blkdebug,
		})
	case "blklogwrites":
		if u.Blklogwrites == nil {
			return nil, fmt.Errorf("expected Blklogwrites to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsBlklogwrites
		}{
			Driver:                      u.Driver,
			NodeName:                    u.NodeName,
			Discard:                     u.Discard,
			Cache:                       u.Cache,
			ReadOnly:                    u.ReadOnly,
			AutoReadOnly:                u.AutoReadOnly,
			ForceShare:                  u.ForceShare,
			DetectZeroes:                u.DetectZeroes,
			BlockdevOptionsBlklogwrites: u.Blklogwrites,
		})
	case "blkverify":
		if u.Blkverify == nil {
			return nil, fmt.Errorf("expected Blkverify to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsBlkverify
		}{
			Driver:                   u.Driver,
			NodeName:                 u.NodeName,
			Discard:                  u.Discard,
			Cache:                    u.Cache,
			ReadOnly:                 u.ReadOnly,
			AutoReadOnly:             u.AutoReadOnly,
			ForceShare:               u.ForceShare,
			DetectZeroes:             u.DetectZeroes,
			BlockdevOptionsBlkverify: u.Blkverify,
		})
	case "blkreplay":
		if u.Blkreplay == nil {
			return nil, fmt.Errorf("expected Blkreplay to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsBlkreplay
		}{
			Driver:                   u.Driver,
			NodeName:                 u.NodeName,
			Discard:                  u.Discard,
			Cache:                    u.Cache,
			ReadOnly:                 u.ReadOnly,
			AutoReadOnly:             u.AutoReadOnly,
			ForceShare:               u.ForceShare,
			DetectZeroes:             u.DetectZeroes,
			BlockdevOptionsBlkreplay: u.Blkreplay,
		})
	case "bochs":
		if u.Bochs == nil {
			return nil, fmt.Errorf("expected Bochs to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGenericFormat
		}{
			Driver:                       u.Driver,
			NodeName:                     u.NodeName,
			Discard:                      u.Discard,
			Cache:                        u.Cache,
			ReadOnly:                     u.ReadOnly,
			AutoReadOnly:                 u.AutoReadOnly,
			ForceShare:                   u.ForceShare,
			DetectZeroes:                 u.DetectZeroes,
			BlockdevOptionsGenericFormat: u.Bochs,
		})
	case "cloop":
		if u.Cloop == nil {
			return nil, fmt.Errorf("expected Cloop to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGenericFormat
		}{
			Driver:                       u.Driver,
			NodeName:                     u.NodeName,
			Discard:                      u.Discard,
			Cache:                        u.Cache,
			ReadOnly:                     u.ReadOnly,
			AutoReadOnly:                 u.AutoReadOnly,
			ForceShare:                   u.ForceShare,
			DetectZeroes:                 u.DetectZeroes,
			BlockdevOptionsGenericFormat: u.Cloop,
		})
	case "compress":
		if u.Compress == nil {
			return nil, fmt.Errorf("expected Compress to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGenericFormat
		}{
			Driver:                       u.Driver,
			NodeName:                     u.NodeName,
			Discard:                      u.Discard,
			Cache:                        u.Cache,
			ReadOnly:                     u.ReadOnly,
			AutoReadOnly:                 u.AutoReadOnly,
			ForceShare:                   u.ForceShare,
			DetectZeroes:                 u.DetectZeroes,
			BlockdevOptionsGenericFormat: u.Compress,
		})
	case "copy-before-write":
		if u.CopyBeforeWrite == nil {
			return nil, fmt.Errorf("expected CopyBeforeWrite to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsCbw
		}{
			Driver:             u.Driver,
			NodeName:           u.NodeName,
			Discard:            u.Discard,
			Cache:              u.Cache,
			ReadOnly:           u.ReadOnly,
			AutoReadOnly:       u.AutoReadOnly,
			ForceShare:         u.ForceShare,
			DetectZeroes:       u.DetectZeroes,
			BlockdevOptionsCbw: u.CopyBeforeWrite,
		})
	case "copy-on-read":
		if u.CopyOnRead == nil {
			return nil, fmt.Errorf("expected CopyOnRead to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsCor
		}{
			Driver:             u.Driver,
			NodeName:           u.NodeName,
			Discard:            u.Discard,
			Cache:              u.Cache,
			ReadOnly:           u.ReadOnly,
			AutoReadOnly:       u.AutoReadOnly,
			ForceShare:         u.ForceShare,
			DetectZeroes:       u.DetectZeroes,
			BlockdevOptionsCor: u.CopyOnRead,
		})
	case "dmg":
		if u.Dmg == nil {
			return nil, fmt.Errorf("expected Dmg to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGenericFormat
		}{
			Driver:                       u.Driver,
			NodeName:                     u.NodeName,
			Discard:                      u.Discard,
			Cache:                        u.Cache,
			ReadOnly:                     u.ReadOnly,
			AutoReadOnly:                 u.AutoReadOnly,
			ForceShare:                   u.ForceShare,
			DetectZeroes:                 u.DetectZeroes,
			BlockdevOptionsGenericFormat: u.Dmg,
		})
	case "file":
		if u.File == nil {
			return nil, fmt.Errorf("expected File to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsFile
		}{
			Driver:              u.Driver,
			NodeName:            u.NodeName,
			Discard:             u.Discard,
			Cache:               u.Cache,
			ReadOnly:            u.ReadOnly,
			AutoReadOnly:        u.AutoReadOnly,
			ForceShare:          u.ForceShare,
			DetectZeroes:        u.DetectZeroes,
			BlockdevOptionsFile: u.File,
		})
	case "ftp":
		if u.Ftp == nil {
			return nil, fmt.Errorf("expected Ftp to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsCurlFtp
		}{
			Driver:                 u.Driver,
			NodeName:               u.NodeName,
			Discard:                u.Discard,
			Cache:                  u.Cache,
			ReadOnly:               u.ReadOnly,
			AutoReadOnly:           u.AutoReadOnly,
			ForceShare:             u.ForceShare,
			DetectZeroes:           u.DetectZeroes,
			BlockdevOptionsCurlFtp: u.Ftp,
		})
	case "ftps":
		if u.Ftps == nil {
			return nil, fmt.Errorf("expected Ftps to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsCurlFtps
		}{
			Driver:                  u.Driver,
			NodeName:                u.NodeName,
			Discard:                 u.Discard,
			Cache:                   u.Cache,
			ReadOnly:                u.ReadOnly,
			AutoReadOnly:            u.AutoReadOnly,
			ForceShare:              u.ForceShare,
			DetectZeroes:            u.DetectZeroes,
			BlockdevOptionsCurlFtps: u.Ftps,
		})
	case "gluster":
		if u.Gluster == nil {
			return nil, fmt.Errorf("expected Gluster to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGluster
		}{
			Driver:                 u.Driver,
			NodeName:               u.NodeName,
			Discard:                u.Discard,
			Cache:                  u.Cache,
			ReadOnly:               u.ReadOnly,
			AutoReadOnly:           u.AutoReadOnly,
			ForceShare:             u.ForceShare,
			DetectZeroes:           u.DetectZeroes,
			BlockdevOptionsGluster: u.Gluster,
		})
	case "host_cdrom":
		if u.HostCdrom == nil {
			return nil, fmt.Errorf("expected HostCdrom to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsFile
		}{
			Driver:              u.Driver,
			NodeName:            u.NodeName,
			Discard:             u.Discard,
			Cache:               u.Cache,
			ReadOnly:            u.ReadOnly,
			AutoReadOnly:        u.AutoReadOnly,
			ForceShare:          u.ForceShare,
			DetectZeroes:        u.DetectZeroes,
			BlockdevOptionsFile: u.HostCdrom,
		})
	case "host_device":
		if u.HostDevice == nil {
			return nil, fmt.Errorf("expected HostDevice to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsFile
		}{
			Driver:              u.Driver,
			NodeName:            u.NodeName,
			Discard:             u.Discard,
			Cache:               u.Cache,
			ReadOnly:            u.ReadOnly,
			AutoReadOnly:        u.AutoReadOnly,
			ForceShare:          u.ForceShare,
			DetectZeroes:        u.DetectZeroes,
			BlockdevOptionsFile: u.HostDevice,
		})
	case "http":
		if u.Http == nil {
			return nil, fmt.Errorf("expected Http to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsCurlHttp
		}{
			Driver:                  u.Driver,
			NodeName:                u.NodeName,
			Discard:                 u.Discard,
			Cache:                   u.Cache,
			ReadOnly:                u.ReadOnly,
			AutoReadOnly:            u.AutoReadOnly,
			ForceShare:              u.ForceShare,
			DetectZeroes:            u.DetectZeroes,
			BlockdevOptionsCurlHttp: u.Http,
		})
	case "https":
		if u.Https == nil {
			return nil, fmt.Errorf("expected Https to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsCurlHttps
		}{
			Driver:                   u.Driver,
			NodeName:                 u.NodeName,
			Discard:                  u.Discard,
			Cache:                    u.Cache,
			ReadOnly:                 u.ReadOnly,
			AutoReadOnly:             u.AutoReadOnly,
			ForceShare:               u.ForceShare,
			DetectZeroes:             u.DetectZeroes,
			BlockdevOptionsCurlHttps: u.Https,
		})
	case "io_uring":
		if u.IoUring == nil {
			return nil, fmt.Errorf("expected IoUring to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsIoUring
		}{
			Driver:                 u.Driver,
			NodeName:               u.NodeName,
			Discard:                u.Discard,
			Cache:                  u.Cache,
			ReadOnly:               u.ReadOnly,
			AutoReadOnly:           u.AutoReadOnly,
			ForceShare:             u.ForceShare,
			DetectZeroes:           u.DetectZeroes,
			BlockdevOptionsIoUring: u.IoUring,
		})
	case "iscsi":
		if u.Iscsi == nil {
			return nil, fmt.Errorf("expected Iscsi to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsIscsi
		}{
			Driver:               u.Driver,
			NodeName:             u.NodeName,
			Discard:              u.Discard,
			Cache:                u.Cache,
			ReadOnly:             u.ReadOnly,
			AutoReadOnly:         u.AutoReadOnly,
			ForceShare:           u.ForceShare,
			DetectZeroes:         u.DetectZeroes,
			BlockdevOptionsIscsi: u.Iscsi,
		})
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsLUKS
		}{
			Driver:              u.Driver,
			NodeName:            u.NodeName,
			Discard:             u.Discard,
			Cache:               u.Cache,
			ReadOnly:            u.ReadOnly,
			AutoReadOnly:        u.AutoReadOnly,
			ForceShare:          u.ForceShare,
			DetectZeroes:        u.DetectZeroes,
			BlockdevOptionsLUKS: u.Luks,
		})
	case "nbd":
		if u.Nbd == nil {
			return nil, fmt.Errorf("expected Nbd to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsNbd
		}{
			Driver:             u.Driver,
			NodeName:           u.NodeName,
			Discard:            u.Discard,
			Cache:              u.Cache,
			ReadOnly:           u.ReadOnly,
			AutoReadOnly:       u.AutoReadOnly,
			ForceShare:         u.ForceShare,
			DetectZeroes:       u.DetectZeroes,
			BlockdevOptionsNbd: u.Nbd,
		})
	case "nfs":
		if u.Nfs == nil {
			return nil, fmt.Errorf("expected Nfs to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsNfs
		}{
			Driver:             u.Driver,
			NodeName:           u.NodeName,
			Discard:            u.Discard,
			Cache:              u.Cache,
			ReadOnly:           u.ReadOnly,
			AutoReadOnly:       u.AutoReadOnly,
			ForceShare:         u.ForceShare,
			DetectZeroes:       u.DetectZeroes,
			BlockdevOptionsNfs: u.Nfs,
		})
	case "null-aio":
		if u.NullAio == nil {
			return nil, fmt.Errorf("expected NullAio to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsNull
		}{
			Driver:              u.Driver,
			NodeName:            u.NodeName,
			Discard:             u.Discard,
			Cache:               u.Cache,
			ReadOnly:            u.ReadOnly,
			AutoReadOnly:        u.AutoReadOnly,
			ForceShare:          u.ForceShare,
			DetectZeroes:        u.DetectZeroes,
			BlockdevOptionsNull: u.NullAio,
		})
	case "null-co":
		if u.NullCo == nil {
			return nil, fmt.Errorf("expected NullCo to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsNull
		}{
			Driver:              u.Driver,
			NodeName:            u.NodeName,
			Discard:             u.Discard,
			Cache:               u.Cache,
			ReadOnly:            u.ReadOnly,
			AutoReadOnly:        u.AutoReadOnly,
			ForceShare:          u.ForceShare,
			DetectZeroes:        u.DetectZeroes,
			BlockdevOptionsNull: u.NullCo,
		})
	case "nvme":
		if u.Nvme == nil {
			return nil, fmt.Errorf("expected Nvme to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsNVMe
		}{
			Driver:              u.Driver,
			NodeName:            u.NodeName,
			Discard:             u.Discard,
			Cache:               u.Cache,
			ReadOnly:            u.ReadOnly,
			AutoReadOnly:        u.AutoReadOnly,
			ForceShare:          u.ForceShare,
			DetectZeroes:        u.DetectZeroes,
			BlockdevOptionsNVMe: u.Nvme,
		})
	case "nvme-io_uring":
		if u.NvmeIoUring == nil {
			return nil, fmt.Errorf("expected NvmeIoUring to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsNvmeIoUring
		}{
			Driver:                     u.Driver,
			NodeName:                   u.NodeName,
			Discard:                    u.Discard,
			Cache:                      u.Cache,
			ReadOnly:                   u.ReadOnly,
			AutoReadOnly:               u.AutoReadOnly,
			ForceShare:                 u.ForceShare,
			DetectZeroes:               u.DetectZeroes,
			BlockdevOptionsNvmeIoUring: u.NvmeIoUring,
		})
	case "parallels":
		if u.Parallels == nil {
			return nil, fmt.Errorf("expected Parallels to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGenericFormat
		}{
			Driver:                       u.Driver,
			NodeName:                     u.NodeName,
			Discard:                      u.Discard,
			Cache:                        u.Cache,
			ReadOnly:                     u.ReadOnly,
			AutoReadOnly:                 u.AutoReadOnly,
			ForceShare:                   u.ForceShare,
			DetectZeroes:                 u.DetectZeroes,
			BlockdevOptionsGenericFormat: u.Parallels,
		})
	case "preallocate":
		if u.Preallocate == nil {
			return nil, fmt.Errorf("expected Preallocate to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsPreallocate
		}{
			Driver:                     u.Driver,
			NodeName:                   u.NodeName,
			Discard:                    u.Discard,
			Cache:                      u.Cache,
			ReadOnly:                   u.ReadOnly,
			AutoReadOnly:               u.AutoReadOnly,
			ForceShare:                 u.ForceShare,
			DetectZeroes:               u.DetectZeroes,
			BlockdevOptionsPreallocate: u.Preallocate,
		})
	case "qcow2":
		if u.Qcow2 == nil {
			return nil, fmt.Errorf("expected Qcow2 to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsQcow2
		}{
			Driver:               u.Driver,
			NodeName:             u.NodeName,
			Discard:              u.Discard,
			Cache:                u.Cache,
			ReadOnly:             u.ReadOnly,
			AutoReadOnly:         u.AutoReadOnly,
			ForceShare:           u.ForceShare,
			DetectZeroes:         u.DetectZeroes,
			BlockdevOptionsQcow2: u.Qcow2,
		})
	case "qcow":
		if u.Qcow == nil {
			return nil, fmt.Errorf("expected Qcow to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsQcow
		}{
			Driver:              u.Driver,
			NodeName:            u.NodeName,
			Discard:             u.Discard,
			Cache:               u.Cache,
			ReadOnly:            u.ReadOnly,
			AutoReadOnly:        u.AutoReadOnly,
			ForceShare:          u.ForceShare,
			DetectZeroes:        u.DetectZeroes,
			BlockdevOptionsQcow: u.Qcow,
		})
	case "qed":
		if u.Qed == nil {
			return nil, fmt.Errorf("expected Qed to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGenericCOWFormat
		}{
			Driver:                          u.Driver,
			NodeName:                        u.NodeName,
			Discard:                         u.Discard,
			Cache:                           u.Cache,
			ReadOnly:                        u.ReadOnly,
			AutoReadOnly:                    u.AutoReadOnly,
			ForceShare:                      u.ForceShare,
			DetectZeroes:                    u.DetectZeroes,
			BlockdevOptionsGenericCOWFormat: u.Qed,
		})
	case "quorum":
		if u.Quorum == nil {
			return nil, fmt.Errorf("expected Quorum to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsQuorum
		}{
			Driver:                u.Driver,
			NodeName:              u.NodeName,
			Discard:               u.Discard,
			Cache:                 u.Cache,
			ReadOnly:              u.ReadOnly,
			AutoReadOnly:          u.AutoReadOnly,
			ForceShare:            u.ForceShare,
			DetectZeroes:          u.DetectZeroes,
			BlockdevOptionsQuorum: u.Quorum,
		})
	case "raw":
		if u.Raw == nil {
			return nil, fmt.Errorf("expected Raw to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsRaw
		}{
			Driver:             u.Driver,
			NodeName:           u.NodeName,
			Discard:            u.Discard,
			Cache:              u.Cache,
			ReadOnly:           u.ReadOnly,
			AutoReadOnly:       u.AutoReadOnly,
			ForceShare:         u.ForceShare,
			DetectZeroes:       u.DetectZeroes,
			BlockdevOptionsRaw: u.Raw,
		})
	case "rbd":
		if u.Rbd == nil {
			return nil, fmt.Errorf("expected Rbd to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsRbd
		}{
			Driver:             u.Driver,
			NodeName:           u.NodeName,
			Discard:            u.Discard,
			Cache:              u.Cache,
			ReadOnly:           u.ReadOnly,
			AutoReadOnly:       u.AutoReadOnly,
			ForceShare:         u.ForceShare,
			DetectZeroes:       u.DetectZeroes,
			BlockdevOptionsRbd: u.Rbd,
		})
	case "replication":
		if u.Replication == nil {
			return nil, fmt.Errorf("expected Replication to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsReplication
		}{
			Driver:                     u.Driver,
			NodeName:                   u.NodeName,
			Discard:                    u.Discard,
			Cache:                      u.Cache,
			ReadOnly:                   u.ReadOnly,
			AutoReadOnly:               u.AutoReadOnly,
			ForceShare:                 u.ForceShare,
			DetectZeroes:               u.DetectZeroes,
			BlockdevOptionsReplication: u.Replication,
		})
	case "snapshot-access":
		if u.SnapshotAccess == nil {
			return nil, fmt.Errorf("expected SnapshotAccess to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGenericFormat
		}{
			Driver:                       u.Driver,
			NodeName:                     u.NodeName,
			Discard:                      u.Discard,
			Cache:                        u.Cache,
			ReadOnly:                     u.ReadOnly,
			AutoReadOnly:                 u.AutoReadOnly,
			ForceShare:                   u.ForceShare,
			DetectZeroes:                 u.DetectZeroes,
			BlockdevOptionsGenericFormat: u.SnapshotAccess,
		})
	case "ssh":
		if u.Ssh == nil {
			return nil, fmt.Errorf("expected Ssh to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsSsh
		}{
			Driver:             u.Driver,
			NodeName:           u.NodeName,
			Discard:            u.Discard,
			Cache:              u.Cache,
			ReadOnly:           u.ReadOnly,
			AutoReadOnly:       u.AutoReadOnly,
			ForceShare:         u.ForceShare,
			DetectZeroes:       u.DetectZeroes,
			BlockdevOptionsSsh: u.Ssh,
		})
	case "throttle":
		if u.Throttle == nil {
			return nil, fmt.Errorf("expected Throttle to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsThrottle
		}{
			Driver:                  u.Driver,
			NodeName:                u.NodeName,
			Discard:                 u.Discard,
			Cache:                   u.Cache,
			ReadOnly:                u.ReadOnly,
			AutoReadOnly:            u.AutoReadOnly,
			ForceShare:              u.ForceShare,
			DetectZeroes:            u.DetectZeroes,
			BlockdevOptionsThrottle: u.Throttle,
		})
	case "vdi":
		if u.Vdi == nil {
			return nil, fmt.Errorf("expected Vdi to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGenericFormat
		}{
			Driver:                       u.Driver,
			NodeName:                     u.NodeName,
			Discard:                      u.Discard,
			Cache:                        u.Cache,
			ReadOnly:                     u.ReadOnly,
			AutoReadOnly:                 u.AutoReadOnly,
			ForceShare:                   u.ForceShare,
			DetectZeroes:                 u.DetectZeroes,
			BlockdevOptionsGenericFormat: u.Vdi,
		})
	case "vhdx":
		if u.Vhdx == nil {
			return nil, fmt.Errorf("expected Vhdx to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGenericFormat
		}{
			Driver:                       u.Driver,
			NodeName:                     u.NodeName,
			Discard:                      u.Discard,
			Cache:                        u.Cache,
			ReadOnly:                     u.ReadOnly,
			AutoReadOnly:                 u.AutoReadOnly,
			ForceShare:                   u.ForceShare,
			DetectZeroes:                 u.DetectZeroes,
			BlockdevOptionsGenericFormat: u.Vhdx,
		})
	case "virtio-blk-vfio-pci":
		if u.VirtioBlkVfioPci == nil {
			return nil, fmt.Errorf("expected VirtioBlkVfioPci to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsVirtioBlkVfioPci
		}{
			Driver:                          u.Driver,
			NodeName:                        u.NodeName,
			Discard:                         u.Discard,
			Cache:                           u.Cache,
			ReadOnly:                        u.ReadOnly,
			AutoReadOnly:                    u.AutoReadOnly,
			ForceShare:                      u.ForceShare,
			DetectZeroes:                    u.DetectZeroes,
			BlockdevOptionsVirtioBlkVfioPci: u.VirtioBlkVfioPci,
		})
	case "virtio-blk-vhost-user":
		if u.VirtioBlkVhostUser == nil {
			return nil, fmt.Errorf("expected VirtioBlkVhostUser to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsVirtioBlkVhostUser
		}{
			Driver:                            u.Driver,
			NodeName:                          u.NodeName,
			Discard:                           u.Discard,
			Cache:                             u.Cache,
			ReadOnly:                          u.ReadOnly,
			AutoReadOnly:                      u.AutoReadOnly,
			ForceShare:                        u.ForceShare,
			DetectZeroes:                      u.DetectZeroes,
			BlockdevOptionsVirtioBlkVhostUser: u.VirtioBlkVhostUser,
		})
	case "virtio-blk-vhost-vdpa":
		if u.VirtioBlkVhostVdpa == nil {
			return nil, fmt.Errorf("expected VirtioBlkVhostVdpa to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsVirtioBlkVhostVdpa
		}{
			Driver:                            u.Driver,
			NodeName:                          u.NodeName,
			Discard:                           u.Discard,
			Cache:                             u.Cache,
			ReadOnly:                          u.ReadOnly,
			AutoReadOnly:                      u.AutoReadOnly,
			ForceShare:                        u.ForceShare,
			DetectZeroes:                      u.DetectZeroes,
			BlockdevOptionsVirtioBlkVhostVdpa: u.VirtioBlkVhostVdpa,
		})
	case "vmdk":
		if u.Vmdk == nil {
			return nil, fmt.Errorf("expected Vmdk to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGenericCOWFormat
		}{
			Driver:                          u.Driver,
			NodeName:                        u.NodeName,
			Discard:                         u.Discard,
			Cache:                           u.Cache,
			ReadOnly:                        u.ReadOnly,
			AutoReadOnly:                    u.AutoReadOnly,
			ForceShare:                      u.ForceShare,
			DetectZeroes:                    u.DetectZeroes,
			BlockdevOptionsGenericCOWFormat: u.Vmdk,
		})
	case "vpc":
		if u.Vpc == nil {
			return nil, fmt.Errorf("expected Vpc to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsGenericFormat
		}{
			Driver:                       u.Driver,
			NodeName:                     u.NodeName,
			Discard:                      u.Discard,
			Cache:                        u.Cache,
			ReadOnly:                     u.ReadOnly,
			AutoReadOnly:                 u.AutoReadOnly,
			ForceShare:                   u.ForceShare,
			DetectZeroes:                 u.DetectZeroes,
			BlockdevOptionsGenericFormat: u.Vpc,
		})
	case "vvfat":
		if u.Vvfat == nil {
			return nil, fmt.Errorf("expected Vvfat to be set")
		}

		return json.Marshal(struct {
			Driver       BlockdevDriver               `json:"driver"`
			NodeName     *string                      `json:"node-name,omitempty"`
			Discard      *BlockdevDiscardOptions      `json:"discard,omitempty"`
			Cache        *BlockdevCacheOptions        `json:"cache,omitempty"`
			ReadOnly     *bool                        `json:"read-only,omitempty"`
			AutoReadOnly *bool                        `json:"auto-read-only,omitempty"`
			ForceShare   *bool                        `json:"force-share,omitempty"`
			DetectZeroes *BlockdevDetectZeroesOptions `json:"detect-zeroes,omitempty"`
			*BlockdevOptionsVVFAT
		}{
			Driver:               u.Driver,
			NodeName:             u.NodeName,
			Discard:              u.Discard,
			Cache:                u.Cache,
			ReadOnly:             u.ReadOnly,
			AutoReadOnly:         u.AutoReadOnly,
			ForceShare:           u.ForceShare,
			DetectZeroes:         u.DetectZeroes,
			BlockdevOptionsVVFAT: u.Vvfat,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockdevRef
//
// Reference to a block device.
type BlockdevRef struct {
	// Definition defines a new block device inline
	Definition *BlockdevOptions `json:"-"`
	// Reference references the ID of an existing block device
	Reference *string `json:"-"`
}

func (a BlockdevRef) MarshalJSON() ([]byte, error) {
	switch {
	case a.Definition != nil:
		return json.Marshal(a.Definition)
	case a.Reference != nil:
		return json.Marshal(a.Reference)
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockdevRefOrNull
//
// Reference to a block device.
type BlockdevRefOrNull struct {
	// Definition defines a new block device inline
	Definition *BlockdevOptions `json:"-"`
	// Reference references the ID of an existing block device. An empty string means that no block device should be referenced. Deprecated; use null instead.
	Reference *string `json:"-"`
	// Null No block device should be referenced (since 2.10)
	Null *Null `json:"-"`
}

func (a BlockdevRefOrNull) MarshalJSON() ([]byte, error) {
	switch {
	case a.Definition != nil:
		return json.Marshal(a.Definition)
	case a.Reference != nil:
		return json.Marshal(a.Reference)
	case a.Null != nil:
		return json.Marshal(a.Null)
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockdevAdd
//
// Creates a new block device.
type BlockdevAdd struct {
	BlockdevOptions
}

func (BlockdevAdd) Command() string {
	return "blockdev-add"
}

func (cmd BlockdevAdd) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-add", cmd, nil)
}

// BlockdevReopen
//
// Reopens one or more block devices using the given set of options. Any option not specified will be reset to its default value regardless of its previous status. If an option cannot be changed or a particular driver does not support reopening then the command will return an error. All devices in the list are reopened in one transaction, so if one of them fails then the whole transaction is cancelled. The command receives a list of block devices to reopen. For each one of them, the top-level @node-name option (from BlockdevOptions) must be specified and is used to select the block device to be reopened. Other @node-name options must be either omitted or set to the current name of the appropriate node. This command won't change any node name and any attempt to do it will result in an error. In the case of options that refer to child nodes, the behavior of
type BlockdevReopen struct {
	Options []BlockdevOptions `json:"options"`
}

func (BlockdevReopen) Command() string {
	return "blockdev-reopen"
}

func (cmd BlockdevReopen) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-reopen", cmd, nil)
}

// BlockdevDel
//
// Deletes a block device that has been added using blockdev-add. The command will fail if the node is attached to a device or is otherwise being used.
type BlockdevDel struct {
	// NodeName Name of the graph node to delete.
	NodeName string `json:"node-name"`
}

func (BlockdevDel) Command() string {
	return "blockdev-del"
}

func (cmd BlockdevDel) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-del", cmd, nil)
}

// BlockdevCreateOptionsFile
//
// Driver specific image creation options for file.
type BlockdevCreateOptionsFile struct {
	// Filename Filename for the new image file
	Filename string `json:"filename"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// Preallocation Preallocation mode for the new image (default: off;
	Preallocation *PreallocMode `json:"preallocation,omitempty"`
	// Nocow Turn off copy-on-write (valid only on btrfs; default: off)
	Nocow *bool `json:"nocow,omitempty"`
	// ExtentSizeHint Extent size hint to add to the image file; 0 for
	ExtentSizeHint *uint64 `json:"extent-size-hint,omitempty"`
}

// BlockdevCreateOptionsGluster
//
// Driver specific image creation options for gluster.
type BlockdevCreateOptionsGluster struct {
	// Location Where to store the new image file
	Location BlockdevOptionsGluster `json:"location"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// Preallocation Preallocation mode for the new image (default: off;
	Preallocation *PreallocMode `json:"preallocation,omitempty"`
}

// BlockdevCreateOptionsLUKS
//
// Driver specific image creation options for LUKS.
type BlockdevCreateOptionsLUKS struct {
	QCryptoBlockCreateOptionsLUKS

	// File Node to create the image format on, mandatory except when 'preallocation' is not requested
	File *BlockdevRef `json:"file,omitempty"`
	// Header Block device holding a detached LUKS header. (since 9.0)
	Header *BlockdevRef `json:"header,omitempty"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// Preallocation Preallocation mode for the new image (since: 4.2)
	Preallocation *PreallocMode `json:"preallocation,omitempty"`
}

// BlockdevCreateOptionsNfs
//
// Driver specific image creation options for NFS.
type BlockdevCreateOptionsNfs struct {
	// Location Where to store the new image file
	Location BlockdevOptionsNfs `json:"location"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
}

// BlockdevCreateOptionsParallels
//
// Driver specific image creation options for parallels.
type BlockdevCreateOptionsParallels struct {
	// File Node to create the image format on
	File BlockdevRef `json:"file"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// ClusterSize Cluster size in bytes (default: 1 MB)
	ClusterSize *uint64 `json:"cluster-size,omitempty"`
}

// BlockdevCreateOptionsQcow
//
// Driver specific image creation options for qcow.
type BlockdevCreateOptionsQcow struct {
	// File Node to create the image format on
	File BlockdevRef `json:"file"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// BackingFile File name of the backing file if a backing file should be used
	BackingFile *string `json:"backing-file,omitempty"`
	// Encrypt Encryption options if the image should be encrypted
	Encrypt *QCryptoBlockCreateOptions `json:"encrypt,omitempty"`
}

// BlockdevQcow2Version
type BlockdevQcow2Version string

const (
	// BlockdevQcow2VersionV2 The original QCOW2 format as introduced in qemu 0.10 (version 2)
	BlockdevQcow2VersionV2 BlockdevQcow2Version = "v2"
	// BlockdevQcow2VersionV3 The extended QCOW2 format as introduced in qemu 1.1 (version 3)
	BlockdevQcow2VersionV3 BlockdevQcow2Version = "v3"
)

// Qcow2CompressionType Compression type used in qcow2 image file
type Qcow2CompressionType string

const (
	// Qcow2CompressionTypeZlib zlib compression, see <http://zlib.net/>
	Qcow2CompressionTypeZlib Qcow2CompressionType = "zlib"
	// Qcow2CompressionTypeZstd zstd compression, see <http://github.com/facebook/zstd>
	Qcow2CompressionTypeZstd Qcow2CompressionType = "zstd"
)

// BlockdevCreateOptionsQcow2
//
// Driver specific image creation options for qcow2.
type BlockdevCreateOptionsQcow2 struct {
	// File Node to create the image format on
	File BlockdevRef `json:"file"`
	// DataFile Node to use as an external data file in which all guest data is stored so that only metadata remains in the qcow2 file
	DataFile *BlockdevRef `json:"data-file,omitempty"`
	// DataFileRaw True if the external data file must stay valid as a standalone (read-only) raw image without looking at qcow2
	DataFileRaw *bool `json:"data-file-raw,omitempty"`
	// ExtendedL2 True to make the image have extended L2 entries
	ExtendedL2 *bool `json:"extended-l2,omitempty"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// Version Compatibility level (default: v3)
	Version *BlockdevQcow2Version `json:"version,omitempty"`
	// BackingFile File name of the backing file if a backing file should be used
	BackingFile *string `json:"backing-file,omitempty"`
	// BackingFmt Name of the block driver to use for the backing file
	BackingFmt *BlockdevDriver `json:"backing-fmt,omitempty"`
	// Encrypt Encryption options if the image should be encrypted
	Encrypt *QCryptoBlockCreateOptions `json:"encrypt,omitempty"`
	// ClusterSize qcow2 cluster size in bytes (default: 65536)
	ClusterSize *uint64 `json:"cluster-size,omitempty"`
	// Preallocation Preallocation mode for the new image (default: off;
	Preallocation *PreallocMode `json:"preallocation,omitempty"`
	// LazyRefcounts True if refcounts may be updated lazily
	LazyRefcounts *bool `json:"lazy-refcounts,omitempty"`
	// RefcountBits Width of reference counts in bits (default: 16)
	RefcountBits *int64 `json:"refcount-bits,omitempty"`
	// CompressionType The image cluster compression method
	CompressionType *Qcow2CompressionType `json:"compression-type,omitempty"`
}

// BlockdevCreateOptionsQed
//
// Driver specific image creation options for qed.
type BlockdevCreateOptionsQed struct {
	// File Node to create the image format on
	File BlockdevRef `json:"file"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// BackingFile File name of the backing file if a backing file should be used
	BackingFile *string `json:"backing-file,omitempty"`
	// BackingFmt Name of the block driver to use for the backing file
	BackingFmt *BlockdevDriver `json:"backing-fmt,omitempty"`
	// ClusterSize Cluster size in bytes (default: 65536)
	ClusterSize *uint64 `json:"cluster-size,omitempty"`
	// TableSize L1/L2 table size (in clusters)
	TableSize *int64 `json:"table-size,omitempty"`
}

// BlockdevCreateOptionsRbd
//
// Driver specific image creation options for rbd/Ceph.
type BlockdevCreateOptionsRbd struct {
	// Location Where to store the new image file. This location cannot point to a snapshot.
	Location BlockdevOptionsRbd `json:"location"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// ClusterSize RBD object size
	ClusterSize *uint64 `json:"cluster-size,omitempty"`
	// Encrypt Image encryption options. (Since 6.1)
	Encrypt *RbdEncryptionCreateOptions `json:"encrypt,omitempty"`
}

// BlockdevVmdkSubformat Subformat options for VMDK images
type BlockdevVmdkSubformat string

const (
	// BlockdevVmdkSubformatMonolithicsparse Single file image with sparse cluster allocation
	BlockdevVmdkSubformatMonolithicsparse BlockdevVmdkSubformat = "monolithicSparse"
	// BlockdevVmdkSubformatMonolithicflat Single flat data image and a descriptor file
	BlockdevVmdkSubformatMonolithicflat BlockdevVmdkSubformat = "monolithicFlat"
	// BlockdevVmdkSubformatTwogbmaxextentsparse Data is split into 2GB (per virtual LBA) sparse extent files, in addition to a descriptor file
	BlockdevVmdkSubformatTwogbmaxextentsparse BlockdevVmdkSubformat = "twoGbMaxExtentSparse"
	// BlockdevVmdkSubformatTwogbmaxextentflat Data is split into 2GB (per virtual LBA) flat extent files, in addition to a descriptor file
	BlockdevVmdkSubformatTwogbmaxextentflat BlockdevVmdkSubformat = "twoGbMaxExtentFlat"
	// BlockdevVmdkSubformatStreamoptimized Single file image sparse cluster allocation, optimized for streaming over network.
	BlockdevVmdkSubformatStreamoptimized BlockdevVmdkSubformat = "streamOptimized"
)

// BlockdevVmdkAdapterType Adapter type info for VMDK images
type BlockdevVmdkAdapterType string

const (
	BlockdevVmdkAdapterTypeIde       BlockdevVmdkAdapterType = "ide"
	BlockdevVmdkAdapterTypeBuslogic  BlockdevVmdkAdapterType = "buslogic"
	BlockdevVmdkAdapterTypeLsilogic  BlockdevVmdkAdapterType = "lsilogic"
	BlockdevVmdkAdapterTypeLegacyesx BlockdevVmdkAdapterType = "legacyESX"
)

// BlockdevCreateOptionsVmdk
//
// Driver specific image creation options for VMDK.
type BlockdevCreateOptionsVmdk struct {
	// File Where to store the new image file. This refers to the image file for monolithcSparse and streamOptimized format, or the descriptor file for other formats.
	File BlockdevRef `json:"file"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// Extents Where to store the data extents. Required for monolithcFlat, twoGbMaxExtentSparse and twoGbMaxExtentFlat formats. For monolithicFlat, only one entry is required; for twoGbMaxExtent* formats, the number of entries required is calculated as extent_number = virtual_size / 2GB. Providing more extents than will be used is an error.
	Extents []BlockdevRef `json:"extents,omitempty"`
	// Subformat The subformat of the VMDK image. Default: "monolithicSparse".
	Subformat *BlockdevVmdkSubformat `json:"subformat,omitempty"`
	// BackingFile The path of backing file. Default: no backing file is used.
	BackingFile *string `json:"backing-file,omitempty"`
	// AdapterType The adapter type used to fill in the descriptor.
	AdapterType *BlockdevVmdkAdapterType `json:"adapter-type,omitempty"`
	// Hwversion Hardware version. The meaningful options are "4" or
	Hwversion *string `json:"hwversion,omitempty"`
	// Toolsversion VMware guest tools version. Default: "2147483647" (Since 6.2)
	Toolsversion *string `json:"toolsversion,omitempty"`
	// ZeroedGrain Whether to enable zeroed-grain feature for sparse
	ZeroedGrain *bool `json:"zeroed-grain,omitempty"`
}

// BlockdevCreateOptionsSsh
//
// Driver specific image creation options for SSH.
type BlockdevCreateOptionsSsh struct {
	// Location Where to store the new image file
	Location BlockdevOptionsSsh `json:"location"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
}

// BlockdevCreateOptionsVdi
//
// Driver specific image creation options for VDI.
type BlockdevCreateOptionsVdi struct {
	// File Node to create the image format on
	File BlockdevRef `json:"file"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// Preallocation Preallocation mode for the new image (default: off;
	Preallocation *PreallocMode `json:"preallocation,omitempty"`
}

// BlockdevVhdxSubformat
type BlockdevVhdxSubformat string

const (
	// BlockdevVhdxSubformatDynamic Growing image file
	BlockdevVhdxSubformatDynamic BlockdevVhdxSubformat = "dynamic"
	// BlockdevVhdxSubformatFixed Preallocated fixed-size image file
	BlockdevVhdxSubformatFixed BlockdevVhdxSubformat = "fixed"
)

// BlockdevCreateOptionsVhdx
//
// Driver specific image creation options for vhdx.
type BlockdevCreateOptionsVhdx struct {
	// File Node to create the image format on
	File BlockdevRef `json:"file"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// LogSize Log size in bytes, must be a multiple of 1 MB (default: 1 MB)
	LogSize *uint64 `json:"log-size,omitempty"`
	// BlockSize Block size in bytes, must be a multiple of 1 MB and not
	BlockSize *uint64 `json:"block-size,omitempty"`
	// Subformat vhdx subformat (default: dynamic)
	Subformat *BlockdevVhdxSubformat `json:"subformat,omitempty"`
	// BlockStateZero Force use of payload blocks of type 'ZERO'. Non-standard, but default. Do not set to 'off' when using 'qemu-img convert' with subformat=dynamic.
	BlockStateZero *bool `json:"block-state-zero,omitempty"`
}

// BlockdevVpcSubformat
type BlockdevVpcSubformat string

const (
	// BlockdevVpcSubformatDynamic Growing image file
	BlockdevVpcSubformatDynamic BlockdevVpcSubformat = "dynamic"
	// BlockdevVpcSubformatFixed Preallocated fixed-size image file
	BlockdevVpcSubformatFixed BlockdevVpcSubformat = "fixed"
)

// BlockdevCreateOptionsVpc
//
// Driver specific image creation options for vpc (VHD).
type BlockdevCreateOptionsVpc struct {
	// File Node to create the image format on
	File BlockdevRef `json:"file"`
	// Size Size of the virtual disk in bytes
	Size uint64 `json:"size"`
	// Subformat vhdx subformat (default: dynamic)
	Subformat *BlockdevVpcSubformat `json:"subformat,omitempty"`
	// ForceSize Force use of the exact byte size instead of rounding to the next size that can be represented in CHS geometry
	ForceSize *bool `json:"force-size,omitempty"`
}

// BlockdevCreateOptions
//
// Options for creating an image format on a given node.
type BlockdevCreateOptions struct {
	// Discriminator: driver

	// Driver block driver to create the image format
	Driver BlockdevDriver `json:"driver"`

	File      *BlockdevCreateOptionsFile      `json:"-"`
	Gluster   *BlockdevCreateOptionsGluster   `json:"-"`
	Luks      *BlockdevCreateOptionsLUKS      `json:"-"`
	Nfs       *BlockdevCreateOptionsNfs       `json:"-"`
	Parallels *BlockdevCreateOptionsParallels `json:"-"`
	Qcow      *BlockdevCreateOptionsQcow      `json:"-"`
	Qcow2     *BlockdevCreateOptionsQcow2     `json:"-"`
	Qed       *BlockdevCreateOptionsQed       `json:"-"`
	Rbd       *BlockdevCreateOptionsRbd       `json:"-"`
	Ssh       *BlockdevCreateOptionsSsh       `json:"-"`
	Vdi       *BlockdevCreateOptionsVdi       `json:"-"`
	Vhdx      *BlockdevCreateOptionsVhdx      `json:"-"`
	Vmdk      *BlockdevCreateOptionsVmdk      `json:"-"`
	Vpc       *BlockdevCreateOptionsVpc       `json:"-"`
}

func (u BlockdevCreateOptions) MarshalJSON() ([]byte, error) {
	switch u.Driver {
	case "file":
		if u.File == nil {
			return nil, fmt.Errorf("expected File to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsFile
		}{
			Driver:                    u.Driver,
			BlockdevCreateOptionsFile: u.File,
		})
	case "gluster":
		if u.Gluster == nil {
			return nil, fmt.Errorf("expected Gluster to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsGluster
		}{
			Driver:                       u.Driver,
			BlockdevCreateOptionsGluster: u.Gluster,
		})
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsLUKS
		}{
			Driver:                    u.Driver,
			BlockdevCreateOptionsLUKS: u.Luks,
		})
	case "nfs":
		if u.Nfs == nil {
			return nil, fmt.Errorf("expected Nfs to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsNfs
		}{
			Driver:                   u.Driver,
			BlockdevCreateOptionsNfs: u.Nfs,
		})
	case "parallels":
		if u.Parallels == nil {
			return nil, fmt.Errorf("expected Parallels to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsParallels
		}{
			Driver:                         u.Driver,
			BlockdevCreateOptionsParallels: u.Parallels,
		})
	case "qcow":
		if u.Qcow == nil {
			return nil, fmt.Errorf("expected Qcow to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsQcow
		}{
			Driver:                    u.Driver,
			BlockdevCreateOptionsQcow: u.Qcow,
		})
	case "qcow2":
		if u.Qcow2 == nil {
			return nil, fmt.Errorf("expected Qcow2 to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsQcow2
		}{
			Driver:                     u.Driver,
			BlockdevCreateOptionsQcow2: u.Qcow2,
		})
	case "qed":
		if u.Qed == nil {
			return nil, fmt.Errorf("expected Qed to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsQed
		}{
			Driver:                   u.Driver,
			BlockdevCreateOptionsQed: u.Qed,
		})
	case "rbd":
		if u.Rbd == nil {
			return nil, fmt.Errorf("expected Rbd to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsRbd
		}{
			Driver:                   u.Driver,
			BlockdevCreateOptionsRbd: u.Rbd,
		})
	case "ssh":
		if u.Ssh == nil {
			return nil, fmt.Errorf("expected Ssh to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsSsh
		}{
			Driver:                   u.Driver,
			BlockdevCreateOptionsSsh: u.Ssh,
		})
	case "vdi":
		if u.Vdi == nil {
			return nil, fmt.Errorf("expected Vdi to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsVdi
		}{
			Driver:                   u.Driver,
			BlockdevCreateOptionsVdi: u.Vdi,
		})
	case "vhdx":
		if u.Vhdx == nil {
			return nil, fmt.Errorf("expected Vhdx to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsVhdx
		}{
			Driver:                    u.Driver,
			BlockdevCreateOptionsVhdx: u.Vhdx,
		})
	case "vmdk":
		if u.Vmdk == nil {
			return nil, fmt.Errorf("expected Vmdk to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsVmdk
		}{
			Driver:                    u.Driver,
			BlockdevCreateOptionsVmdk: u.Vmdk,
		})
	case "vpc":
		if u.Vpc == nil {
			return nil, fmt.Errorf("expected Vpc to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevCreateOptionsVpc
		}{
			Driver:                   u.Driver,
			BlockdevCreateOptionsVpc: u.Vpc,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockdevCreate
//
// Starts a job to create an image format on a given node. The job is automatically finalized, but a manual job-dismiss is required.
type BlockdevCreate struct {
	// JobId Identifier for the newly created job.
	JobId string `json:"job-id"`
	// Options Options for the image creation.
	Options BlockdevCreateOptions `json:"options"`
}

func (BlockdevCreate) Command() string {
	return "blockdev-create"
}

func (cmd BlockdevCreate) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-create", cmd, nil)
}

// BlockdevAmendOptionsLUKS
//
// Driver specific image amend options for LUKS.
type BlockdevAmendOptionsLUKS struct {
	QCryptoBlockAmendOptionsLUKS
}

// BlockdevAmendOptionsQcow2
//
// Driver specific image amend options for qcow2. For now, only encryption options can be amended
type BlockdevAmendOptionsQcow2 struct {
	// Encrypt Encryption options to be amended
	Encrypt *QCryptoBlockAmendOptions `json:"encrypt,omitempty"`
}

// BlockdevAmendOptions
//
// Options for amending an image format
type BlockdevAmendOptions struct {
	// Discriminator: driver

	// Driver Block driver of the node to amend.
	Driver BlockdevDriver `json:"driver"`

	Luks  *BlockdevAmendOptionsLUKS  `json:"-"`
	Qcow2 *BlockdevAmendOptionsQcow2 `json:"-"`
}

func (u BlockdevAmendOptions) MarshalJSON() ([]byte, error) {
	switch u.Driver {
	case "luks":
		if u.Luks == nil {
			return nil, fmt.Errorf("expected Luks to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevAmendOptionsLUKS
		}{
			Driver:                   u.Driver,
			BlockdevAmendOptionsLUKS: u.Luks,
		})
	case "qcow2":
		if u.Qcow2 == nil {
			return nil, fmt.Errorf("expected Qcow2 to be set")
		}

		return json.Marshal(struct {
			Driver BlockdevDriver `json:"driver"`
			*BlockdevAmendOptionsQcow2
		}{
			Driver:                    u.Driver,
			BlockdevAmendOptionsQcow2: u.Qcow2,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockdevAmend
//
// Starts a job to amend format specific options of an existing open block device The job is automatically finalized, but a manual job-dismiss is required.
type BlockdevAmend struct {
	// JobId Identifier for the newly created job.
	JobId string `json:"job-id"`
	// NodeName Name of the block node to work on
	NodeName string `json:"node-name"`
	// Options Options (driver specific)
	Options BlockdevAmendOptions `json:"options"`
	// Force Allow unsafe operations, format specific For luks that allows erase of the last active keyslot (permanent loss of data), and replacement of an active keyslot (possible loss of data if IO error happens)
	Force *bool `json:"force,omitempty"`
}

func (BlockdevAmend) Command() string {
	return "x-blockdev-amend"
}

func (cmd BlockdevAmend) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "x-blockdev-amend", cmd, nil)
}

// BlockErrorAction An enumeration of action that has been taken when a DISK I/O occurs
type BlockErrorAction string

const (
	// BlockErrorActionIgnore error has been ignored
	BlockErrorActionIgnore BlockErrorAction = "ignore"
	// BlockErrorActionReport error has been reported to the device
	BlockErrorActionReport BlockErrorAction = "report"
	// BlockErrorActionStop error caused VM to be stopped
	BlockErrorActionStop BlockErrorAction = "stop"
)

// BlockImageCorruptedEvent (BLOCK_IMAGE_CORRUPTED)
//
// Emitted when a disk image is being marked corrupt. The image can be identified by its device or node name. The 'device' field is always present for compatibility reasons, but it can be empty ("") if the image does not have a device name associated.
type BlockImageCorruptedEvent struct {
	// Device device name. This is always present for compatibility reasons, but it can be empty ("") if the image does not have a device name associated.
	Device string `json:"device"`
	// NodeName node name (Since: 2.4)
	NodeName *string `json:"node-name,omitempty"`
	// Msg informative message for human consumption, such as the kind of corruption being detected. It should not be parsed by machine as it is not guaranteed to be stable
	Msg string `json:"msg"`
	// Offset if the corruption resulted from an image access, this is the host's access offset into the image
	Offset *int64 `json:"offset,omitempty"`
	// Size if the corruption resulted from an image access, this is the access size
	Size *int64 `json:"size,omitempty"`
	// Fatal if set, the image is marked corrupt and therefore unusable after this event and must be repaired (Since 2.2; before, every BLOCK_IMAGE_CORRUPTED event was fatal)
	Fatal bool `json:"fatal"`
}

func (BlockImageCorruptedEvent) Event() string {
	return "BLOCK_IMAGE_CORRUPTED"
}

// BlockIoErrorEvent (BLOCK_IO_ERROR)
//
// Emitted when a disk I/O error occurs
type BlockIoErrorEvent struct {
	// Device device name. This is always present for compatibility reasons, but it can be empty ("") if the image does not have a device name associated.
	Device string `json:"device"`
	// NodeName node name. Note that errors may be reported for the root node that is directly attached to a guest device rather than for the node where the error occurred. The node name is
	NodeName *string `json:"node-name,omitempty"`
	// Operation I/O operation
	Operation IoOperationType `json:"operation"`
	// Action action that has been taken
	Action BlockErrorAction `json:"action"`
	// Nospace true if I/O error was caused due to a no-space condition. This key is only present if query-block's io-status is present, please see query-block documentation for more information
	Nospace *bool `json:"nospace,omitempty"`
	// Reason human readable string describing the error cause. (This field is a debugging aid for humans, it should not be parsed by
	Reason string `json:"reason"`
}

func (BlockIoErrorEvent) Event() string {
	return "BLOCK_IO_ERROR"
}

// BlockJobCompletedEvent (BLOCK_JOB_COMPLETED)
//
// Emitted when a block job has completed
type BlockJobCompletedEvent struct {
	// Type job type
	Type JobType `json:"type"`
	// Device The job identifier. Originally the device name but other values are allowed since QEMU 2.7
	Device string `json:"device"`
	// Len maximum progress value
	Len int64 `json:"len"`
	// Offset current progress value. On success this is equal to len. On failure this is less than len
	Offset int64 `json:"offset"`
	// Speed rate limit, bytes per second
	Speed int64 `json:"speed"`
	// Error error message. Only present on failure. This field contains a human-readable error message. There are no semantics other than that streaming has failed and clients should not try to interpret the error string
	Error *string `json:"error,omitempty"`
}

func (BlockJobCompletedEvent) Event() string {
	return "BLOCK_JOB_COMPLETED"
}

// BlockJobCancelledEvent (BLOCK_JOB_CANCELLED)
//
// Emitted when a block job has been cancelled
type BlockJobCancelledEvent struct {
	// Type job type
	Type JobType `json:"type"`
	// Device The job identifier. Originally the device name but other values are allowed since QEMU 2.7
	Device string `json:"device"`
	// Len maximum progress value
	Len int64 `json:"len"`
	// Offset current progress value. On success this is equal to len. On failure this is less than len
	Offset int64 `json:"offset"`
	// Speed rate limit, bytes per second
	Speed int64 `json:"speed"`
}

func (BlockJobCancelledEvent) Event() string {
	return "BLOCK_JOB_CANCELLED"
}

// BlockJobErrorEvent (BLOCK_JOB_ERROR)
//
// Emitted when a block job encounters an error
type BlockJobErrorEvent struct {
	// Device The job identifier. Originally the device name but other values are allowed since QEMU 2.7
	Device string `json:"device"`
	// Operation I/O operation
	Operation IoOperationType `json:"operation"`
	// Action action that has been taken
	Action BlockErrorAction `json:"action"`
}

func (BlockJobErrorEvent) Event() string {
	return "BLOCK_JOB_ERROR"
}

// BlockJobReadyEvent (BLOCK_JOB_READY)
//
// Emitted when a block job is ready to complete
type BlockJobReadyEvent struct {
	// Type job type
	Type JobType `json:"type"`
	// Device The job identifier. Originally the device name but other values are allowed since QEMU 2.7
	Device string `json:"device"`
	// Len maximum progress value
	Len int64 `json:"len"`
	// Offset current progress value. On success this is equal to len. On failure this is less than len
	Offset int64 `json:"offset"`
	// Speed rate limit, bytes per second
	Speed int64 `json:"speed"`
}

func (BlockJobReadyEvent) Event() string {
	return "BLOCK_JOB_READY"
}

// BlockJobPendingEvent (BLOCK_JOB_PENDING)
//
// Emitted when a block job is awaiting explicit authorization to finalize graph changes via @block-job-finalize. If this job is part of a transaction, it will not emit this event until the transaction has converged first.
type BlockJobPendingEvent struct {
	// Type job type
	Type JobType `json:"type"`
	// Id The job identifier.
	Id string `json:"id"`
}

func (BlockJobPendingEvent) Event() string {
	return "BLOCK_JOB_PENDING"
}

// PreallocMode Preallocation mode of QEMU image file
type PreallocMode string

const (
	// PreallocModeOff no preallocation
	PreallocModeOff PreallocMode = "off"
	// PreallocModeMetadata preallocate only for metadata
	PreallocModeMetadata PreallocMode = "metadata"
	// PreallocModeFalloc like @full preallocation but allocate disk space by posix_fallocate() rather than writing data.
	PreallocModeFalloc PreallocMode = "falloc"
	// PreallocModeFull preallocate all data by writing it to the device to ensure disk space is really available. This data may or may not be zero, depending on the image format and storage. @full preallocation also sets up metadata correctly.
	PreallocModeFull PreallocMode = "full"
)

// BlockWriteThresholdEvent (BLOCK_WRITE_THRESHOLD)
//
// Emitted when writes on block device reaches or exceeds the configured write threshold. For thin-provisioned devices, this means the device should be extended to avoid pausing for disk exhaustion. The event is one shot. Once triggered, it needs to be re-registered with another block-set-write-threshold command.
type BlockWriteThresholdEvent struct {
	// NodeName graph node name on which the threshold was exceeded.
	NodeName string `json:"node-name"`
	// AmountExceeded amount of data which exceeded the threshold, in bytes.
	AmountExceeded uint64 `json:"amount-exceeded"`
	// WriteThreshold last configured threshold, in bytes.
	WriteThreshold uint64 `json:"write-threshold"`
}

func (BlockWriteThresholdEvent) Event() string {
	return "BLOCK_WRITE_THRESHOLD"
}

// BlockSetWriteThreshold
//
// Change the write threshold for a block drive. An event will be delivered if a write to this block drive crosses the configured threshold. The threshold is an offset, thus must be non-negative. Default is no write threshold. Setting the threshold to zero disables it. This is useful to transparently resize thin-provisioned drives without the guest OS noticing.
type BlockSetWriteThreshold struct {
	// NodeName graph node name on which the threshold must be set.
	NodeName string `json:"node-name"`
	// WriteThreshold configured threshold for the block device, bytes. Use 0 to disable the threshold.
	WriteThreshold uint64 `json:"write-threshold"`
}

func (BlockSetWriteThreshold) Command() string {
	return "block-set-write-threshold"
}

func (cmd BlockSetWriteThreshold) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-set-write-threshold", cmd, nil)
}

// BlockdevChange
//
// Dynamically reconfigure the block driver state graph. It can be used to add, remove, insert or replace a graph node. Currently only the Quorum driver implements this feature to add or remove its child. This is useful to fix a broken quorum child. If @node is specified, it will be inserted under @parent. @child may not be specified in this case. If both @parent and @child are specified but @node is not, @child will be detached from @parent.
type BlockdevChange struct {
	// Parent the id or name of the parent node.
	Parent string `json:"parent"`
	// Child the name of a child under the given parent node.
	Child *string `json:"child,omitempty"`
	// Node the name of the node that will be added.
	Node *string `json:"node,omitempty"`
}

func (BlockdevChange) Command() string {
	return "x-blockdev-change"
}

func (cmd BlockdevChange) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "x-blockdev-change", cmd, nil)
}

// BlockdevSetIothread
//
// Move @node and its children into the @iothread. If @iothread is null then move @node and its children into the main loop. The node must not be attached to a BlockBackend.
type BlockdevSetIothread struct {
	// NodeName the name of the block driver node
	NodeName string `json:"node-name"`
	// Iothread the name of the IOThread object or null for the main loop
	Iothread StrOrNull `json:"iothread"`
	// Force true if the node and its children should be moved when a BlockBackend is already attached
	Force *bool `json:"force,omitempty"`
}

func (BlockdevSetIothread) Command() string {
	return "x-blockdev-set-iothread"
}

func (cmd BlockdevSetIothread) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "x-blockdev-set-iothread", cmd, nil)
}

// QuorumOpType An enumeration of the quorum operation types
type QuorumOpType string

const (
	// QuorumOpTypeRead read operation
	QuorumOpTypeRead QuorumOpType = "read"
	// QuorumOpTypeWrite write operation
	QuorumOpTypeWrite QuorumOpType = "write"
	// QuorumOpTypeFlush flush operation
	QuorumOpTypeFlush QuorumOpType = "flush"
)

// QuorumFailureEvent (QUORUM_FAILURE)
//
// Emitted by the Quorum block driver if it fails to establish a quorum
type QuorumFailureEvent struct {
	// Reference device name if defined else node name
	Reference string `json:"reference"`
	// SectorNum number of the first sector of the failed read operation
	SectorNum int64 `json:"sector-num"`
	// SectorsCount failed read operation sector count
	SectorsCount int64 `json:"sectors-count"`
}

func (QuorumFailureEvent) Event() string {
	return "QUORUM_FAILURE"
}

// QuorumReportBadEvent (QUORUM_REPORT_BAD)
//
// Emitted to report a corruption of a Quorum file
type QuorumReportBadEvent struct {
	// Type quorum operation type (Since 2.6)
	Type QuorumOpType `json:"type"`
	// Error error message. Only present on failure. This field contains a human-readable error message. There are no semantics other than that the block layer reported an error and clients should not try to interpret the error string.
	Error *string `json:"error,omitempty"`
	// NodeName the graph node name of the block driver state
	NodeName string `json:"node-name"`
	// SectorNum number of the first sector of the failed read operation
	SectorNum int64 `json:"sector-num"`
	// SectorsCount failed read operation sector count
	SectorsCount int64 `json:"sectors-count"`
}

func (QuorumReportBadEvent) Event() string {
	return "QUORUM_REPORT_BAD"
}

// BlockdevSnapshotInternal
type BlockdevSnapshotInternal struct {
	// Device the device name or node-name of a root node to generate the snapshot from
	Device string `json:"device"`
	// Name the name of the internal snapshot to be created
	Name string `json:"name"`
}

// BlockdevSnapshotInternalSync
//
// Synchronously take an internal snapshot of a block device, when the format of the image used supports it. If the name is an empty string, or a snapshot with name already exists, the operation will fail. For the arguments, see the documentation of BlockdevSnapshotInternal.
type BlockdevSnapshotInternalSync struct {
	BlockdevSnapshotInternal
}

func (BlockdevSnapshotInternalSync) Command() string {
	return "blockdev-snapshot-internal-sync"
}

func (cmd BlockdevSnapshotInternalSync) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-snapshot-internal-sync", cmd, nil)
}

// BlockdevSnapshotDeleteInternalSync
//
// Synchronously delete an internal snapshot of a block device, when the format of the image used support it. The snapshot is identified by name or id or both. One of the name or id is required. Return SnapshotInfo for the successfully deleted snapshot.
type BlockdevSnapshotDeleteInternalSync struct {
	// Device the device name or node-name of a root node to delete the snapshot from
	Device string `json:"device"`
	// Id optional the snapshot's ID to be deleted
	Id *string `json:"id,omitempty"`
	// Name optional the snapshot's name to be deleted
	Name *string `json:"name,omitempty"`
}

func (BlockdevSnapshotDeleteInternalSync) Command() string {
	return "blockdev-snapshot-delete-internal-sync"
}

func (cmd BlockdevSnapshotDeleteInternalSync) Execute(ctx context.Context, client api.Client) (SnapshotInfo, error) {
	var ret SnapshotInfo

	return ret, client.Execute(ctx, "blockdev-snapshot-delete-internal-sync", cmd, &ret)
}

// DummyBlockCoreForceArrays
//
// Not used by QMP; hack to let us use BlockGraphInfoList internally
type DummyBlockCoreForceArrays struct {
	UnusedBlockGraphInfo []BlockGraphInfo `json:"unused-block-graph-info"`
}

// NbdServerOptions
//
// Keep this type consistent with the nbd-server-start arguments. The only intended difference is using SocketAddress instead of SocketAddressLegacy.
type NbdServerOptions struct {
	// Addr Address on which to listen.
	Addr SocketAddress `json:"addr"`
	// TlsCreds ID of the TLS credentials object (since 2.6).
	TlsCreds *string `json:"tls-creds,omitempty"`
	// TlsAuthz ID of the QAuthZ authorization object used to validate the client's x509 distinguished name. This object is is only resolved at time of use, so can be deleted and recreated on the fly while the NBD server is active. If missing, it will default to denying access (since 4.0).
	TlsAuthz *string `json:"tls-authz,omitempty"`
	// MaxConnections The maximum number of connections to allow at the same time, 0 for unlimited. Setting this to 1 also stops the server from advertising multiple client support (since 5.2;
	MaxConnections *uint32 `json:"max-connections,omitempty"`
}

// NbdServerStart
//
// Start an NBD server listening on the given host and port. Block devices can then be exported using @nbd-server-add. The NBD server will present them as named exports; for example, another QEMU
type NbdServerStart struct {
	// Addr Address on which to listen.
	Addr SocketAddressLegacy `json:"addr"`
	// TlsCreds ID of the TLS credentials object (since 2.6).
	TlsCreds *string `json:"tls-creds,omitempty"`
	// TlsAuthz ID of the QAuthZ authorization object used to validate the client's x509 distinguished name. This object is is only resolved at time of use, so can be deleted and recreated on the fly while the NBD server is active. If missing, it will default to denying access (since 4.0).
	TlsAuthz *string `json:"tls-authz,omitempty"`
	// MaxConnections The maximum number of connections to allow at the same time, 0 for unlimited. Setting this to 1 also stops the server from advertising multiple client support (since 5.2;
	MaxConnections *uint32 `json:"max-connections,omitempty"`
}

func (NbdServerStart) Command() string {
	return "nbd-server-start"
}

func (cmd NbdServerStart) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "nbd-server-start", cmd, nil)
}

// BlockExportOptionsNbdBase
//
// An NBD block export (common options shared between nbd-server-add and the NBD branch of block-export-add).
type BlockExportOptionsNbdBase struct {
	// Name Export name. If unspecified, the @device parameter is used as the export name. (Since 2.12)
	Name *string `json:"name,omitempty"`
	// Description Free-form description of the export, up to 4096 bytes. (Since 5.0)
	Description *string `json:"description,omitempty"`
}

// BlockExportOptionsNbd
//
// An NBD block export (distinct options used in the NBD branch of block-export-add).
type BlockExportOptionsNbd struct {
	BlockExportOptionsNbdBase

	// Bitmaps Also export each of the named dirty bitmaps reachable from @device, so the NBD client can use NBD_OPT_SET_META_CONTEXT with
	Bitmaps []BlockDirtyBitmapOrStr `json:"bitmaps,omitempty"`
	// AllocationDepth Also export the allocation depth map for @device, so the NBD client can use NBD_OPT_SET_META_CONTEXT with the
	AllocationDepth *bool `json:"allocation-depth,omitempty"`
}

// BlockExportOptionsVhostUserBlk
//
// A vhost-user-blk block export.
type BlockExportOptionsVhostUserBlk struct {
	// Addr The vhost-user socket on which to listen. Both 'unix' and 'fd' SocketAddress types are supported. Passed fds must be UNIX domain sockets.
	Addr SocketAddress `json:"addr"`
	// LogicalBlockSize Logical block size in bytes. Defaults to 512 bytes.
	LogicalBlockSize *uint64 `json:"logical-block-size,omitempty"`
	// NumQueues Number of request virtqueues. Must be greater than 0. Defaults to 1.
	NumQueues *uint16 `json:"num-queues,omitempty"`
}

// FuseExportAllowOther Possible allow_other modes for FUSE exports.
type FuseExportAllowOther string

const (
	// FuseExportAllowOtherOff Do not pass allow_other as a mount option.
	FuseExportAllowOtherOff FuseExportAllowOther = "off"
	// FuseExportAllowOtherOn Pass allow_other as a mount option.
	FuseExportAllowOtherOn FuseExportAllowOther = "on"
	// FuseExportAllowOtherAuto Try mounting with allow_other first, and if that fails, retry without allow_other.
	FuseExportAllowOtherAuto FuseExportAllowOther = "auto"
)

// BlockExportOptionsFuse
//
// Options for exporting a block graph node on some (file) mountpoint as a raw image.
type BlockExportOptionsFuse struct {
	// Mountpoint Path on which to export the block device via FUSE. This must point to an existing regular file.
	Mountpoint string `json:"mountpoint"`
	// Growable Whether writes beyond the EOF should grow the block node
	Growable *bool `json:"growable,omitempty"`
	// AllowOther If this is off, only qemu's user is allowed access to this export. That cannot be changed even with chmod or chown. Enabling this option will allow other users access to the export with the FUSE mount option "allow_other". Note that using allow_other as a non-root user requires user_allow_other to be enabled in the global fuse.conf configuration file. In auto mode (the default), the FUSE export driver will first attempt to mount the export with allow_other, and if that fails, try again
	AllowOther *FuseExportAllowOther `json:"allow-other,omitempty"`
}

// BlockExportOptionsVduseBlk
//
// A vduse-blk block export.
type BlockExportOptionsVduseBlk struct {
	// Name the name of VDUSE device (must be unique across the host).
	Name string `json:"name"`
	// NumQueues the number of virtqueues. Defaults to 1.
	NumQueues *uint16 `json:"num-queues,omitempty"`
	// QueueSize the size of virtqueue. Defaults to 256.
	QueueSize *uint16 `json:"queue-size,omitempty"`
	// LogicalBlockSize Logical block size in bytes. Range [512, PAGE_SIZE] and must be power of 2. Defaults to 512 bytes.
	LogicalBlockSize *uint64 `json:"logical-block-size,omitempty"`
	// Serial the serial number of virtio block device. Defaults to empty string.
	Serial *string `json:"serial,omitempty"`
}

// NbdServerAddOptions
//
// An NBD block export, per legacy nbd-server-add command.
type NbdServerAddOptions struct {
	BlockExportOptionsNbdBase

	// Device The device name or node name of the node to be exported
	Device string `json:"device"`
	// Writable Whether clients should be able to write to the device via the NBD connection (default false).
	Writable *bool `json:"writable,omitempty"`
	// Bitmap Also export a single dirty bitmap reachable from @device, so the NBD client can use NBD_OPT_SET_META_CONTEXT with the
	Bitmap *string `json:"bitmap,omitempty"`
}

// NbdServerAdd
//
// Export a block node to QEMU's embedded NBD server. The export name will be used as the id for the resulting block export.
type NbdServerAdd struct {
	NbdServerAddOptions
}

func (NbdServerAdd) Command() string {
	return "nbd-server-add"
}

func (cmd NbdServerAdd) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "nbd-server-add", cmd, nil)
}

// BlockExportRemoveMode Mode for removing a block export.
type BlockExportRemoveMode string

const (
	// BlockExportRemoveModeSafe Remove export if there are no existing connections, fail otherwise.
	BlockExportRemoveModeSafe BlockExportRemoveMode = "safe"
	// BlockExportRemoveModeHard Drop all connections immediately and remove export.
	BlockExportRemoveModeHard BlockExportRemoveMode = "hard"
)

// NbdServerRemove
//
// Remove NBD export by name.
type NbdServerRemove struct {
	// Name Block export id.
	Name string `json:"name"`
	// Mode Mode of command operation. See @BlockExportRemoveMode description. Default is 'safe'.
	Mode *BlockExportRemoveMode `json:"mode,omitempty"`
}

func (NbdServerRemove) Command() string {
	return "nbd-server-remove"
}

func (cmd NbdServerRemove) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "nbd-server-remove", cmd, nil)
}

// NbdServerStop
//
// Stop QEMU's embedded NBD server, and unregister all devices previously added via @nbd-server-add.
type NbdServerStop struct {
}

func (NbdServerStop) Command() string {
	return "nbd-server-stop"
}

func (cmd NbdServerStop) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "nbd-server-stop", cmd, nil)
}

// BlockExportType An enumeration of block export types
type BlockExportType string

const (
	// BlockExportTypeNbd NBD export
	BlockExportTypeNbd BlockExportType = "nbd"
	// BlockExportTypeVhostUserBlk vhost-user-blk export (since 5.2)
	BlockExportTypeVhostUserBlk BlockExportType = "vhost-user-blk"
	// BlockExportTypeFuse FUSE export (since: 6.0)
	BlockExportTypeFuse BlockExportType = "fuse"
	// BlockExportTypeVduseBlk vduse-blk export (since 7.1)
	BlockExportTypeVduseBlk BlockExportType = "vduse-blk"
)

// BlockExportOptions
//
// Describes a block export, i.e. how single node should be exported on an external interface.
type BlockExportOptions struct {
	// Discriminator: type

	// Type Block export type
	Type BlockExportType `json:"type"`
	// Id A unique identifier for the block export (across all export types)
	Id string `json:"id"`
	// FixedIothread True prevents the block node from being moved to another thread while the export is active. If true and @iothread is given, export creation fails if the block node cannot be moved to the iothread. The default is false.
	FixedIothread *bool `json:"fixed-iothread,omitempty"`
	// Iothread The name of the iothread object where the export will run. The default is to use the thread currently associated with
	Iothread *string `json:"iothread,omitempty"`
	// NodeName The node name of the block node to be exported
	NodeName string `json:"node-name"`
	// Writable True if clients should be able to write to the export (default false)
	Writable *bool `json:"writable,omitempty"`
	// Writethrough If true, caches are flushed after every write request
	Writethrough *bool `json:"writethrough,omitempty"`

	Nbd          *BlockExportOptionsNbd          `json:"-"`
	VhostUserBlk *BlockExportOptionsVhostUserBlk `json:"-"`
	Fuse         *BlockExportOptionsFuse         `json:"-"`
	VduseBlk     *BlockExportOptionsVduseBlk     `json:"-"`
}

func (u BlockExportOptions) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "nbd":
		if u.Nbd == nil {
			return nil, fmt.Errorf("expected Nbd to be set")
		}

		return json.Marshal(struct {
			Type          BlockExportType `json:"type"`
			Id            string          `json:"id"`
			FixedIothread *bool           `json:"fixed-iothread,omitempty"`
			Iothread      *string         `json:"iothread,omitempty"`
			NodeName      string          `json:"node-name"`
			Writable      *bool           `json:"writable,omitempty"`
			Writethrough  *bool           `json:"writethrough,omitempty"`
			*BlockExportOptionsNbd
		}{
			Type:                  u.Type,
			Id:                    u.Id,
			FixedIothread:         u.FixedIothread,
			Iothread:              u.Iothread,
			NodeName:              u.NodeName,
			Writable:              u.Writable,
			Writethrough:          u.Writethrough,
			BlockExportOptionsNbd: u.Nbd,
		})
	case "vhost-user-blk":
		if u.VhostUserBlk == nil {
			return nil, fmt.Errorf("expected VhostUserBlk to be set")
		}

		return json.Marshal(struct {
			Type          BlockExportType `json:"type"`
			Id            string          `json:"id"`
			FixedIothread *bool           `json:"fixed-iothread,omitempty"`
			Iothread      *string         `json:"iothread,omitempty"`
			NodeName      string          `json:"node-name"`
			Writable      *bool           `json:"writable,omitempty"`
			Writethrough  *bool           `json:"writethrough,omitempty"`
			*BlockExportOptionsVhostUserBlk
		}{
			Type:                           u.Type,
			Id:                             u.Id,
			FixedIothread:                  u.FixedIothread,
			Iothread:                       u.Iothread,
			NodeName:                       u.NodeName,
			Writable:                       u.Writable,
			Writethrough:                   u.Writethrough,
			BlockExportOptionsVhostUserBlk: u.VhostUserBlk,
		})
	case "fuse":
		if u.Fuse == nil {
			return nil, fmt.Errorf("expected Fuse to be set")
		}

		return json.Marshal(struct {
			Type          BlockExportType `json:"type"`
			Id            string          `json:"id"`
			FixedIothread *bool           `json:"fixed-iothread,omitempty"`
			Iothread      *string         `json:"iothread,omitempty"`
			NodeName      string          `json:"node-name"`
			Writable      *bool           `json:"writable,omitempty"`
			Writethrough  *bool           `json:"writethrough,omitempty"`
			*BlockExportOptionsFuse
		}{
			Type:                   u.Type,
			Id:                     u.Id,
			FixedIothread:          u.FixedIothread,
			Iothread:               u.Iothread,
			NodeName:               u.NodeName,
			Writable:               u.Writable,
			Writethrough:           u.Writethrough,
			BlockExportOptionsFuse: u.Fuse,
		})
	case "vduse-blk":
		if u.VduseBlk == nil {
			return nil, fmt.Errorf("expected VduseBlk to be set")
		}

		return json.Marshal(struct {
			Type          BlockExportType `json:"type"`
			Id            string          `json:"id"`
			FixedIothread *bool           `json:"fixed-iothread,omitempty"`
			Iothread      *string         `json:"iothread,omitempty"`
			NodeName      string          `json:"node-name"`
			Writable      *bool           `json:"writable,omitempty"`
			Writethrough  *bool           `json:"writethrough,omitempty"`
			*BlockExportOptionsVduseBlk
		}{
			Type:                       u.Type,
			Id:                         u.Id,
			FixedIothread:              u.FixedIothread,
			Iothread:                   u.Iothread,
			NodeName:                   u.NodeName,
			Writable:                   u.Writable,
			Writethrough:               u.Writethrough,
			BlockExportOptionsVduseBlk: u.VduseBlk,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// BlockExportAdd
//
// Creates a new block export.
type BlockExportAdd struct {
	BlockExportOptions
}

func (BlockExportAdd) Command() string {
	return "block-export-add"
}

func (cmd BlockExportAdd) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-export-add", cmd, nil)
}

// BlockExportDel
//
// Request to remove a block export. This drops the user's reference to the export, but the export may still stay around after this command returns until the shutdown of the export has completed.
type BlockExportDel struct {
	// Id Block export id.
	Id string `json:"id"`
	// Mode Mode of command operation. See @BlockExportRemoveMode description. Default is 'safe'.
	Mode *BlockExportRemoveMode `json:"mode,omitempty"`
}

func (BlockExportDel) Command() string {
	return "block-export-del"
}

func (cmd BlockExportDel) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-export-del", cmd, nil)
}

// BlockExportDeletedEvent (BLOCK_EXPORT_DELETED)
//
// Emitted when a block export is removed and its id can be reused.
type BlockExportDeletedEvent struct {
	// Id Block export id.
	Id string `json:"id"`
}

func (BlockExportDeletedEvent) Event() string {
	return "BLOCK_EXPORT_DELETED"
}

// BlockExportInfo
//
// Information about a single block export.
type BlockExportInfo struct {
	// Id The unique identifier for the block export
	Id string `json:"id"`
	// Type The block export type
	Type BlockExportType `json:"type"`
	// NodeName The node name of the block node that is exported
	NodeName string `json:"node-name"`
	// ShuttingDown True if the export is shutting down (e.g. after a block-export-del command, but before the shutdown has completed)
	ShuttingDown bool `json:"shutting-down"`
}

// QueryBlockExports
type QueryBlockExports struct {
}

func (QueryBlockExports) Command() string {
	return "query-block-exports"
}

func (cmd QueryBlockExports) Execute(ctx context.Context, client api.Client) ([]BlockExportInfo, error) {
	var ret []BlockExportInfo

	return ret, client.Execute(ctx, "query-block-exports", cmd, &ret)
}

// ChardevInfo
//
// Information about a character device.
type ChardevInfo struct {
	// Label the label of the character device
	Label string `json:"label"`
	// Filename the filename of the character device
	Filename string `json:"filename"`
	// FrontendOpen shows whether the frontend device attached to this backend (e.g. with the chardev=... option) is in open or closed state (since 2.1)
	FrontendOpen bool `json:"frontend-open"`
}

// QueryChardev
//
// Returns information about current character devices.
type QueryChardev struct {
}

func (QueryChardev) Command() string {
	return "query-chardev"
}

func (cmd QueryChardev) Execute(ctx context.Context, client api.Client) ([]ChardevInfo, error) {
	var ret []ChardevInfo

	return ret, client.Execute(ctx, "query-chardev", cmd, &ret)
}

// ChardevBackendInfo
//
// Information about a character device backend
type ChardevBackendInfo struct {
	// Name The backend name
	Name string `json:"name"`
}

// QueryChardevBackends
//
// Returns information about character device backends.
type QueryChardevBackends struct {
}

func (QueryChardevBackends) Command() string {
	return "query-chardev-backends"
}

func (cmd QueryChardevBackends) Execute(ctx context.Context, client api.Client) ([]ChardevBackendInfo, error) {
	var ret []ChardevBackendInfo

	return ret, client.Execute(ctx, "query-chardev-backends", cmd, &ret)
}

// DataFormat An enumeration of data format.
type DataFormat string

const (
	// DataFormatUtf8 Data is a UTF-8 string (RFC 3629)
	DataFormatUtf8 DataFormat = "utf8"
	// DataFormatBase64 Data is Base64 encoded binary (RFC 3548)
	DataFormatBase64 DataFormat = "base64"
)

// RingbufWrite
//
// Write to a ring buffer character device.
type RingbufWrite struct {
	// Device the ring buffer character device name
	Device string `json:"device"`
	// Data data to write
	Data string `json:"data"`
	// Format data encoding (default 'utf8').
	Format *DataFormat `json:"format,omitempty"`
}

func (RingbufWrite) Command() string {
	return "ringbuf-write"
}

func (cmd RingbufWrite) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "ringbuf-write", cmd, nil)
}

// RingbufRead
//
// Read from a ring buffer character device.
type RingbufRead struct {
	// Device the ring buffer character device name
	Device string `json:"device"`
	// Size how many bytes to read at most
	Size int64 `json:"size"`
	// Format data encoding (default 'utf8').
	Format *DataFormat `json:"format,omitempty"`
}

func (RingbufRead) Command() string {
	return "ringbuf-read"
}

func (cmd RingbufRead) Execute(ctx context.Context, client api.Client) (string, error) {
	var ret string

	return ret, client.Execute(ctx, "ringbuf-read", cmd, &ret)
}

// ChardevCommon
//
// Configuration shared across all chardev backends
type ChardevCommon struct {
	// Logfile The name of a logfile to save output
	Logfile *string `json:"logfile,omitempty"`
	// Logappend true to append instead of truncate (default to false to truncate)
	Logappend *bool `json:"logappend,omitempty"`
}

// ChardevFile
//
// Configuration info for file chardevs.
type ChardevFile struct {
	ChardevCommon

	// In The name of the input file
	In *string `json:"in,omitempty"`
	// Out The name of the output file
	Out string `json:"out"`
	// Append Open the file in append mode (default false to truncate) (Since 2.6)
	Append *bool `json:"append,omitempty"`
}

// ChardevHostdev
//
// Configuration info for device and pipe chardevs.
type ChardevHostdev struct {
	ChardevCommon

	// Device The name of the special file for the device, i.e.
	Device string `json:"device"`
}

// ChardevSocket
//
// Configuration info for (stream) socket chardevs.
type ChardevSocket struct {
	ChardevCommon

	// Addr socket address to listen on (server=true) or connect to (server=false)
	Addr SocketAddressLegacy `json:"addr"`
	// TlsCreds the ID of the TLS credentials object (since 2.6)
	TlsCreds *string `json:"tls-creds,omitempty"`
	// TlsAuthz the ID of the QAuthZ authorization object against which the client's x509 distinguished name will be validated. This object is only resolved at time of use, so can be deleted and recreated on the fly while the chardev server is active. If missing, it will default to denying access (since 4.0)
	TlsAuthz *string `json:"tls-authz,omitempty"`
	// Server create server socket (default: true)
	Server *bool `json:"server,omitempty"`
	// Wait wait for incoming connection on server sockets (default:
	Wait *bool `json:"wait,omitempty"`
	// Nodelay set TCP_NODELAY socket option (default: false)
	Nodelay *bool `json:"nodelay,omitempty"`
	// Telnet enable telnet protocol on server sockets (default: false)
	Telnet *bool `json:"telnet,omitempty"`
	// Tn3270 enable tn3270 protocol on server sockets (default: false)
	Tn3270 *bool `json:"tn3270,omitempty"`
	// Websocket enable websocket protocol on server sockets
	Websocket *bool `json:"websocket,omitempty"`
	// Reconnect For a client socket, if a socket is disconnected, then attempt a reconnect after the given number of seconds. Setting
	Reconnect *int64 `json:"reconnect,omitempty"`
}

// ChardevUdp
//
// Configuration info for datagram socket chardevs.
type ChardevUdp struct {
	ChardevCommon

	// Remote remote address
	Remote SocketAddressLegacy `json:"remote"`
	// Local local address
	Local *SocketAddressLegacy `json:"local,omitempty"`
}

// ChardevMux
//
// Configuration info for mux chardevs.
type ChardevMux struct {
	ChardevCommon

	// Chardev name of the base chardev.
	Chardev string `json:"chardev"`
}

// ChardevStdio
//
// Configuration info for stdio chardevs.
type ChardevStdio struct {
	ChardevCommon

	// Signal Allow signals (such as SIGINT triggered by ^C) be delivered
	Signal *bool `json:"signal,omitempty"`
}

// ChardevSpiceChannel
//
// Configuration info for spice vm channel chardevs.
type ChardevSpiceChannel struct {
	ChardevCommon

	// Type kind of channel (for example vdagent).
	Type string `json:"type"`
}

// ChardevSpicePort
//
// Configuration info for spice port chardevs.
type ChardevSpicePort struct {
	ChardevCommon

	// Fqdn name of the channel (see docs/spice-port-fqdn.txt)
	Fqdn string `json:"fqdn"`
}

// ChardevDBus
//
// Configuration info for DBus chardevs.
type ChardevDBus struct {
	ChardevCommon

	// Name name of the channel (following docs/spice-port-fqdn.txt)
	Name string `json:"name"`
}

// ChardevVC
//
// Configuration info for virtual console chardevs.
type ChardevVC struct {
	ChardevCommon

	// Width console width, in pixels
	Width *int64 `json:"width,omitempty"`
	// Height console height, in pixels
	Height *int64 `json:"height,omitempty"`
	// Cols console width, in chars
	Cols *int64 `json:"cols,omitempty"`
	// Rows console height, in chars
	Rows *int64 `json:"rows,omitempty"`
}

// ChardevRingbuf
//
// Configuration info for ring buffer chardevs.
type ChardevRingbuf struct {
	ChardevCommon

	// Size ring buffer size, must be power of two, default is 65536
	Size *int64 `json:"size,omitempty"`
}

// ChardevQemuVDAgent
//
// Configuration info for qemu vdagent implementation.
type ChardevQemuVDAgent struct {
	ChardevCommon

	// Mouse enable/disable mouse, default is enabled.
	Mouse *bool `json:"mouse,omitempty"`
	// Clipboard enable/disable clipboard, default is disabled.
	Clipboard *bool `json:"clipboard,omitempty"`
}

// ChardevBackendKind
type ChardevBackendKind string

const (
	ChardevBackendKindFile     ChardevBackendKind = "file"
	ChardevBackendKindSerial   ChardevBackendKind = "serial"
	ChardevBackendKindParallel ChardevBackendKind = "parallel"
	// ChardevBackendKindPipe Since 1.5
	ChardevBackendKindPipe   ChardevBackendKind = "pipe"
	ChardevBackendKindSocket ChardevBackendKind = "socket"
	// ChardevBackendKindUdp Since 1.5
	ChardevBackendKindUdp  ChardevBackendKind = "udp"
	ChardevBackendKindPty  ChardevBackendKind = "pty"
	ChardevBackendKindNull ChardevBackendKind = "null"
	// ChardevBackendKindMux Since 1.5
	ChardevBackendKindMux ChardevBackendKind = "mux"
	// ChardevBackendKindMsmouse Since 1.5
	ChardevBackendKindMsmouse ChardevBackendKind = "msmouse"
	// ChardevBackendKindWctablet Since 2.9
	ChardevBackendKindWctablet ChardevBackendKind = "wctablet"
	// ChardevBackendKindBraille Since 1.5
	ChardevBackendKindBraille ChardevBackendKind = "braille"
	// ChardevBackendKindTestdev Since 2.2
	ChardevBackendKindTestdev ChardevBackendKind = "testdev"
	// ChardevBackendKindStdio Since 1.5
	ChardevBackendKindStdio ChardevBackendKind = "stdio"
	// ChardevBackendKindConsole Since 1.5
	ChardevBackendKindConsole ChardevBackendKind = "console"
	// ChardevBackendKindSpicevmc Since 1.5
	ChardevBackendKindSpicevmc ChardevBackendKind = "spicevmc"
	// ChardevBackendKindSpiceport Since 1.5
	ChardevBackendKindSpiceport ChardevBackendKind = "spiceport"
	// ChardevBackendKindQemuVdagent Since 6.1
	ChardevBackendKindQemuVdagent ChardevBackendKind = "qemu-vdagent"
	// ChardevBackendKindDbus Since 7.0
	ChardevBackendKindDbus ChardevBackendKind = "dbus"
	// ChardevBackendKindVc v1.5
	ChardevBackendKindVc ChardevBackendKind = "vc"
	// ChardevBackendKindRingbuf Since 1.6
	ChardevBackendKindRingbuf ChardevBackendKind = "ringbuf"
	// ChardevBackendKindMemory Since 1.5
	ChardevBackendKindMemory ChardevBackendKind = "memory"
)

// ChardevFileWrapper
type ChardevFileWrapper struct {
	// Data Configuration info for file chardevs
	Data ChardevFile `json:"data"`
}

// ChardevHostdevWrapper
type ChardevHostdevWrapper struct {
	// Data Configuration info for device and pipe chardevs
	Data ChardevHostdev `json:"data"`
}

// ChardevSocketWrapper
type ChardevSocketWrapper struct {
	// Data Configuration info for (stream) socket chardevs
	Data ChardevSocket `json:"data"`
}

// ChardevUdpWrapper
type ChardevUdpWrapper struct {
	// Data Configuration info for datagram socket chardevs
	Data ChardevUdp `json:"data"`
}

// ChardevCommonWrapper
type ChardevCommonWrapper struct {
	// Data Configuration shared across all chardev backends
	Data ChardevCommon `json:"data"`
}

// ChardevMuxWrapper
type ChardevMuxWrapper struct {
	// Data Configuration info for mux chardevs
	Data ChardevMux `json:"data"`
}

// ChardevStdioWrapper
type ChardevStdioWrapper struct {
	// Data Configuration info for stdio chardevs
	Data ChardevStdio `json:"data"`
}

// ChardevSpiceChannelWrapper
type ChardevSpiceChannelWrapper struct {
	// Data Configuration info for spice vm channel chardevs
	Data ChardevSpiceChannel `json:"data"`
}

// ChardevSpicePortWrapper
type ChardevSpicePortWrapper struct {
	// Data Configuration info for spice port chardevs
	Data ChardevSpicePort `json:"data"`
}

// ChardevQemuVDAgentWrapper
type ChardevQemuVDAgentWrapper struct {
	// Data Configuration info for qemu vdagent implementation
	Data ChardevQemuVDAgent `json:"data"`
}

// ChardevDBusWrapper
type ChardevDBusWrapper struct {
	// Data Configuration info for DBus chardevs
	Data ChardevDBus `json:"data"`
}

// ChardevVCWrapper
type ChardevVCWrapper struct {
	// Data Configuration info for virtual console chardevs
	Data ChardevVC `json:"data"`
}

// ChardevRingbufWrapper
type ChardevRingbufWrapper struct {
	// Data Configuration info for ring buffer chardevs
	Data ChardevRingbuf `json:"data"`
}

// ChardevBackend
//
// Configuration info for the new chardev backend.
type ChardevBackend struct {
	// Discriminator: type

	// Type backend type
	Type ChardevBackendKind `json:"type"`

	File        *ChardevFileWrapper         `json:"-"`
	Serial      *ChardevHostdevWrapper      `json:"-"`
	Parallel    *ChardevHostdevWrapper      `json:"-"`
	Pipe        *ChardevHostdevWrapper      `json:"-"`
	Socket      *ChardevSocketWrapper       `json:"-"`
	Udp         *ChardevUdpWrapper          `json:"-"`
	Pty         *ChardevCommonWrapper       `json:"-"`
	Null        *ChardevCommonWrapper       `json:"-"`
	Mux         *ChardevMuxWrapper          `json:"-"`
	Msmouse     *ChardevCommonWrapper       `json:"-"`
	Wctablet    *ChardevCommonWrapper       `json:"-"`
	Braille     *ChardevCommonWrapper       `json:"-"`
	Testdev     *ChardevCommonWrapper       `json:"-"`
	Stdio       *ChardevStdioWrapper        `json:"-"`
	Console     *ChardevCommonWrapper       `json:"-"`
	Spicevmc    *ChardevSpiceChannelWrapper `json:"-"`
	Spiceport   *ChardevSpicePortWrapper    `json:"-"`
	QemuVdagent *ChardevQemuVDAgentWrapper  `json:"-"`
	Dbus        *ChardevDBusWrapper         `json:"-"`
	Vc          *ChardevVCWrapper           `json:"-"`
	Ringbuf     *ChardevRingbufWrapper      `json:"-"`
	Memory      *ChardevRingbufWrapper      `json:"-"`
}

func (u ChardevBackend) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "file":
		if u.File == nil {
			return nil, fmt.Errorf("expected File to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevFileWrapper
		}{
			Type:               u.Type,
			ChardevFileWrapper: u.File,
		})
	case "serial":
		if u.Serial == nil {
			return nil, fmt.Errorf("expected Serial to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevHostdevWrapper
		}{
			Type:                  u.Type,
			ChardevHostdevWrapper: u.Serial,
		})
	case "parallel":
		if u.Parallel == nil {
			return nil, fmt.Errorf("expected Parallel to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevHostdevWrapper
		}{
			Type:                  u.Type,
			ChardevHostdevWrapper: u.Parallel,
		})
	case "pipe":
		if u.Pipe == nil {
			return nil, fmt.Errorf("expected Pipe to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevHostdevWrapper
		}{
			Type:                  u.Type,
			ChardevHostdevWrapper: u.Pipe,
		})
	case "socket":
		if u.Socket == nil {
			return nil, fmt.Errorf("expected Socket to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevSocketWrapper
		}{
			Type:                 u.Type,
			ChardevSocketWrapper: u.Socket,
		})
	case "udp":
		if u.Udp == nil {
			return nil, fmt.Errorf("expected Udp to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevUdpWrapper
		}{
			Type:              u.Type,
			ChardevUdpWrapper: u.Udp,
		})
	case "pty":
		if u.Pty == nil {
			return nil, fmt.Errorf("expected Pty to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevCommonWrapper
		}{
			Type:                 u.Type,
			ChardevCommonWrapper: u.Pty,
		})
	case "null":
		if u.Null == nil {
			return nil, fmt.Errorf("expected Null to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevCommonWrapper
		}{
			Type:                 u.Type,
			ChardevCommonWrapper: u.Null,
		})
	case "mux":
		if u.Mux == nil {
			return nil, fmt.Errorf("expected Mux to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevMuxWrapper
		}{
			Type:              u.Type,
			ChardevMuxWrapper: u.Mux,
		})
	case "msmouse":
		if u.Msmouse == nil {
			return nil, fmt.Errorf("expected Msmouse to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevCommonWrapper
		}{
			Type:                 u.Type,
			ChardevCommonWrapper: u.Msmouse,
		})
	case "wctablet":
		if u.Wctablet == nil {
			return nil, fmt.Errorf("expected Wctablet to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevCommonWrapper
		}{
			Type:                 u.Type,
			ChardevCommonWrapper: u.Wctablet,
		})
	case "braille":
		if u.Braille == nil {
			return nil, fmt.Errorf("expected Braille to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevCommonWrapper
		}{
			Type:                 u.Type,
			ChardevCommonWrapper: u.Braille,
		})
	case "testdev":
		if u.Testdev == nil {
			return nil, fmt.Errorf("expected Testdev to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevCommonWrapper
		}{
			Type:                 u.Type,
			ChardevCommonWrapper: u.Testdev,
		})
	case "stdio":
		if u.Stdio == nil {
			return nil, fmt.Errorf("expected Stdio to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevStdioWrapper
		}{
			Type:                u.Type,
			ChardevStdioWrapper: u.Stdio,
		})
	case "console":
		if u.Console == nil {
			return nil, fmt.Errorf("expected Console to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevCommonWrapper
		}{
			Type:                 u.Type,
			ChardevCommonWrapper: u.Console,
		})
	case "spicevmc":
		if u.Spicevmc == nil {
			return nil, fmt.Errorf("expected Spicevmc to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevSpiceChannelWrapper
		}{
			Type:                       u.Type,
			ChardevSpiceChannelWrapper: u.Spicevmc,
		})
	case "spiceport":
		if u.Spiceport == nil {
			return nil, fmt.Errorf("expected Spiceport to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevSpicePortWrapper
		}{
			Type:                    u.Type,
			ChardevSpicePortWrapper: u.Spiceport,
		})
	case "qemu-vdagent":
		if u.QemuVdagent == nil {
			return nil, fmt.Errorf("expected QemuVdagent to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevQemuVDAgentWrapper
		}{
			Type:                      u.Type,
			ChardevQemuVDAgentWrapper: u.QemuVdagent,
		})
	case "dbus":
		if u.Dbus == nil {
			return nil, fmt.Errorf("expected Dbus to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevDBusWrapper
		}{
			Type:               u.Type,
			ChardevDBusWrapper: u.Dbus,
		})
	case "vc":
		if u.Vc == nil {
			return nil, fmt.Errorf("expected Vc to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevVCWrapper
		}{
			Type:             u.Type,
			ChardevVCWrapper: u.Vc,
		})
	case "ringbuf":
		if u.Ringbuf == nil {
			return nil, fmt.Errorf("expected Ringbuf to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevRingbufWrapper
		}{
			Type:                  u.Type,
			ChardevRingbufWrapper: u.Ringbuf,
		})
	case "memory":
		if u.Memory == nil {
			return nil, fmt.Errorf("expected Memory to be set")
		}

		return json.Marshal(struct {
			Type ChardevBackendKind `json:"type"`
			*ChardevRingbufWrapper
		}{
			Type:                  u.Type,
			ChardevRingbufWrapper: u.Memory,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// ChardevReturn
//
// Return info about the chardev backend just created.
type ChardevReturn struct {
	// Pty name of the slave pseudoterminal device, present if and only if a chardev of type 'pty' was created
	Pty *string `json:"pty,omitempty"`
}

// ChardevAdd
//
// Add a character device backend
type ChardevAdd struct {
	// Id the chardev's ID, must be unique
	Id string `json:"id"`
	// Backend backend type and parameters
	Backend ChardevBackend `json:"backend"`
}

func (ChardevAdd) Command() string {
	return "chardev-add"
}

func (cmd ChardevAdd) Execute(ctx context.Context, client api.Client) (ChardevReturn, error) {
	var ret ChardevReturn

	return ret, client.Execute(ctx, "chardev-add", cmd, &ret)
}

// ChardevChange
//
// Change a character device backend
type ChardevChange struct {
	// Id the chardev's ID, must exist
	Id string `json:"id"`
	// Backend new backend type and parameters
	Backend ChardevBackend `json:"backend"`
}

func (ChardevChange) Command() string {
	return "chardev-change"
}

func (cmd ChardevChange) Execute(ctx context.Context, client api.Client) (ChardevReturn, error) {
	var ret ChardevReturn

	return ret, client.Execute(ctx, "chardev-change", cmd, &ret)
}

// ChardevRemove
//
// Remove a character device backend
type ChardevRemove struct {
	// Id the chardev's ID, must exist and not be in use
	Id string `json:"id"`
}

func (ChardevRemove) Command() string {
	return "chardev-remove"
}

func (cmd ChardevRemove) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "chardev-remove", cmd, nil)
}

// ChardevSendBreak
//
// Send a break to a character device
type ChardevSendBreak struct {
	// Id the chardev's ID, must exist
	Id string `json:"id"`
}

func (ChardevSendBreak) Command() string {
	return "chardev-send-break"
}

func (cmd ChardevSendBreak) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "chardev-send-break", cmd, nil)
}

// VserportChangeEvent (VSERPORT_CHANGE)
//
// Emitted when the guest opens or closes a virtio-serial port.
type VserportChangeEvent struct {
	// Id device identifier of the virtio-serial port
	Id string `json:"id"`
	// Open true if the guest has opened the virtio-serial port
	Open bool `json:"open"`
}

func (VserportChangeEvent) Event() string {
	return "VSERPORT_CHANGE"
}

// QAuthZListPolicy The authorization policy result
type QAuthZListPolicy string

const (
	// QauthzListPolicyDeny deny access
	QauthzListPolicyDeny QAuthZListPolicy = "deny"
	// QauthzListPolicyAllow allow access
	QauthzListPolicyAllow QAuthZListPolicy = "allow"
)

// QAuthZListFormat The authorization policy match format
type QAuthZListFormat string

const (
	// QauthzListFormatExact an exact string match
	QauthzListFormatExact QAuthZListFormat = "exact"
	// QauthzListFormatGlob string with ? and * shell wildcard support
	QauthzListFormatGlob QAuthZListFormat = "glob"
)

// QAuthZListRule
//
// A single authorization rule.
type QAuthZListRule struct {
	// Match a string or glob to match against a user identity
	Match string `json:"match"`
	// Policy the result to return if @match evaluates to true
	Policy QAuthZListPolicy `json:"policy"`
	// Format the format of the @match rule (default 'exact')
	Format *QAuthZListFormat `json:"format,omitempty"`
}

// AuthZListProperties
//
// Properties for authz-list objects.
type AuthZListProperties struct {
	// Policy Default policy to apply when no rule matches (default: deny)
	Policy *QAuthZListPolicy `json:"policy,omitempty"`
	// Rules Authorization rules based on matching user
	Rules []QAuthZListRule `json:"rules,omitempty"`
}

// AuthZListFileProperties
//
// Properties for authz-listfile objects.
type AuthZListFileProperties struct {
	// Filename File name to load the configuration from. The file must contain valid JSON for AuthZListProperties.
	Filename string `json:"filename"`
	// Refresh If true, inotify is used to monitor the file, automatically reloading changes. If an error occurs during reloading, all authorizations will fail until the file is next
	Refresh *bool `json:"refresh,omitempty"`
}

// AuthZPAMProperties
//
// Properties for authz-pam objects.
type AuthZPAMProperties struct {
	// Service PAM service name to use for authorization
	Service string `json:"service"`
}

// AuthZSimpleProperties
//
// Properties for authz-simple objects.
type AuthZSimpleProperties struct {
	// Identity Identifies the allowed user. Its format depends on the network service that authorization object is associated with. For authorizing based on TLS x509 certificates, the identity must be the x509 distinguished name.
	Identity string `json:"identity"`
}

// Abort
//
// This action can be used to test transaction failure.
type Abort struct {
}

// ActionCompletionMode An enumeration of Transactional completion modes.
type ActionCompletionMode string

const (
	// ActionCompletionModeIndividual Do not attempt to cancel any other Actions if any Actions fail after the Transaction request succeeds. All Actions that can complete successfully will do so without waiting on others. This is the default.
	ActionCompletionModeIndividual ActionCompletionMode = "individual"
	// ActionCompletionModeGrouped If any Action fails after the Transaction succeeds, cancel all Actions. Actions do not complete until all Actions are ready to complete. May be rejected by Actions that do not support this completion mode.
	ActionCompletionModeGrouped ActionCompletionMode = "grouped"
)

// TransactionActionKind
type TransactionActionKind string

const (
	// TransactionActionKindAbort Since 1.6
	TransactionActionKindAbort TransactionActionKind = "abort"
	// TransactionActionKindBlockDirtyBitmapAdd Since 2.5
	TransactionActionKindBlockDirtyBitmapAdd TransactionActionKind = "block-dirty-bitmap-add"
	// TransactionActionKindBlockDirtyBitmapRemove Since 4.2
	TransactionActionKindBlockDirtyBitmapRemove TransactionActionKind = "block-dirty-bitmap-remove"
	// TransactionActionKindBlockDirtyBitmapClear Since 2.5
	TransactionActionKindBlockDirtyBitmapClear TransactionActionKind = "block-dirty-bitmap-clear"
	// TransactionActionKindBlockDirtyBitmapEnable Since 4.0
	TransactionActionKindBlockDirtyBitmapEnable TransactionActionKind = "block-dirty-bitmap-enable"
	// TransactionActionKindBlockDirtyBitmapDisable Since 4.0
	TransactionActionKindBlockDirtyBitmapDisable TransactionActionKind = "block-dirty-bitmap-disable"
	// TransactionActionKindBlockDirtyBitmapMerge Since 4.0
	TransactionActionKindBlockDirtyBitmapMerge TransactionActionKind = "block-dirty-bitmap-merge"
	// TransactionActionKindBlockdevBackup Since 2.3
	TransactionActionKindBlockdevBackup TransactionActionKind = "blockdev-backup"
	// TransactionActionKindBlockdevSnapshot Since 2.5
	TransactionActionKindBlockdevSnapshot TransactionActionKind = "blockdev-snapshot"
	// TransactionActionKindBlockdevSnapshotInternalSync Since 1.7
	TransactionActionKindBlockdevSnapshotInternalSync TransactionActionKind = "blockdev-snapshot-internal-sync"
	// TransactionActionKindBlockdevSnapshotSync since 1.1
	TransactionActionKindBlockdevSnapshotSync TransactionActionKind = "blockdev-snapshot-sync"
	// TransactionActionKindDriveBackup Since 1.6
	TransactionActionKindDriveBackup TransactionActionKind = "drive-backup"
)

// AbortWrapper
type AbortWrapper struct {
	Data Abort `json:"data"`
}

// BlockDirtyBitmapAddWrapper
type BlockDirtyBitmapAddWrapper struct {
	Data BlockDirtyBitmapAdd `json:"data"`
}

// BlockDirtyBitmapWrapper
type BlockDirtyBitmapWrapper struct {
	Data BlockDirtyBitmap `json:"data"`
}

// BlockDirtyBitmapMergeWrapper
type BlockDirtyBitmapMergeWrapper struct {
	Data BlockDirtyBitmapMerge `json:"data"`
}

// BlockdevBackupWrapper
type BlockdevBackupWrapper struct {
	Data BlockdevBackup `json:"data"`
}

// BlockdevSnapshotWrapper
type BlockdevSnapshotWrapper struct {
	Data BlockdevSnapshot `json:"data"`
}

// BlockdevSnapshotInternalWrapper
type BlockdevSnapshotInternalWrapper struct {
	Data BlockdevSnapshotInternal `json:"data"`
}

// BlockdevSnapshotSyncWrapper
type BlockdevSnapshotSyncWrapper struct {
	Data BlockdevSnapshotSync `json:"data"`
}

// DriveBackupWrapper
type DriveBackupWrapper struct {
	Data DriveBackup `json:"data"`
}

// TransactionAction
//
// A discriminated record of operations that can be performed with @transaction.
type TransactionAction struct {
	// Discriminator: type

	// Type the operation to be performed
	Type TransactionActionKind `json:"type"`

	Abort                        *AbortWrapper                    `json:"-"`
	BlockDirtyBitmapAdd          *BlockDirtyBitmapAddWrapper      `json:"-"`
	BlockDirtyBitmapRemove       *BlockDirtyBitmapWrapper         `json:"-"`
	BlockDirtyBitmapClear        *BlockDirtyBitmapWrapper         `json:"-"`
	BlockDirtyBitmapEnable       *BlockDirtyBitmapWrapper         `json:"-"`
	BlockDirtyBitmapDisable      *BlockDirtyBitmapWrapper         `json:"-"`
	BlockDirtyBitmapMerge        *BlockDirtyBitmapMergeWrapper    `json:"-"`
	BlockdevBackup               *BlockdevBackupWrapper           `json:"-"`
	BlockdevSnapshot             *BlockdevSnapshotWrapper         `json:"-"`
	BlockdevSnapshotInternalSync *BlockdevSnapshotInternalWrapper `json:"-"`
	BlockdevSnapshotSync         *BlockdevSnapshotSyncWrapper     `json:"-"`
	DriveBackup                  *DriveBackupWrapper              `json:"-"`
}

func (u TransactionAction) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "abort":
		if u.Abort == nil {
			return nil, fmt.Errorf("expected Abort to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*AbortWrapper
		}{
			Type:         u.Type,
			AbortWrapper: u.Abort,
		})
	case "block-dirty-bitmap-add":
		if u.BlockDirtyBitmapAdd == nil {
			return nil, fmt.Errorf("expected BlockDirtyBitmapAdd to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*BlockDirtyBitmapAddWrapper
		}{
			Type:                       u.Type,
			BlockDirtyBitmapAddWrapper: u.BlockDirtyBitmapAdd,
		})
	case "block-dirty-bitmap-remove":
		if u.BlockDirtyBitmapRemove == nil {
			return nil, fmt.Errorf("expected BlockDirtyBitmapRemove to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*BlockDirtyBitmapWrapper
		}{
			Type:                    u.Type,
			BlockDirtyBitmapWrapper: u.BlockDirtyBitmapRemove,
		})
	case "block-dirty-bitmap-clear":
		if u.BlockDirtyBitmapClear == nil {
			return nil, fmt.Errorf("expected BlockDirtyBitmapClear to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*BlockDirtyBitmapWrapper
		}{
			Type:                    u.Type,
			BlockDirtyBitmapWrapper: u.BlockDirtyBitmapClear,
		})
	case "block-dirty-bitmap-enable":
		if u.BlockDirtyBitmapEnable == nil {
			return nil, fmt.Errorf("expected BlockDirtyBitmapEnable to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*BlockDirtyBitmapWrapper
		}{
			Type:                    u.Type,
			BlockDirtyBitmapWrapper: u.BlockDirtyBitmapEnable,
		})
	case "block-dirty-bitmap-disable":
		if u.BlockDirtyBitmapDisable == nil {
			return nil, fmt.Errorf("expected BlockDirtyBitmapDisable to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*BlockDirtyBitmapWrapper
		}{
			Type:                    u.Type,
			BlockDirtyBitmapWrapper: u.BlockDirtyBitmapDisable,
		})
	case "block-dirty-bitmap-merge":
		if u.BlockDirtyBitmapMerge == nil {
			return nil, fmt.Errorf("expected BlockDirtyBitmapMerge to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*BlockDirtyBitmapMergeWrapper
		}{
			Type:                         u.Type,
			BlockDirtyBitmapMergeWrapper: u.BlockDirtyBitmapMerge,
		})
	case "blockdev-backup":
		if u.BlockdevBackup == nil {
			return nil, fmt.Errorf("expected BlockdevBackup to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*BlockdevBackupWrapper
		}{
			Type:                  u.Type,
			BlockdevBackupWrapper: u.BlockdevBackup,
		})
	case "blockdev-snapshot":
		if u.BlockdevSnapshot == nil {
			return nil, fmt.Errorf("expected BlockdevSnapshot to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*BlockdevSnapshotWrapper
		}{
			Type:                    u.Type,
			BlockdevSnapshotWrapper: u.BlockdevSnapshot,
		})
	case "blockdev-snapshot-internal-sync":
		if u.BlockdevSnapshotInternalSync == nil {
			return nil, fmt.Errorf("expected BlockdevSnapshotInternalSync to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*BlockdevSnapshotInternalWrapper
		}{
			Type:                            u.Type,
			BlockdevSnapshotInternalWrapper: u.BlockdevSnapshotInternalSync,
		})
	case "blockdev-snapshot-sync":
		if u.BlockdevSnapshotSync == nil {
			return nil, fmt.Errorf("expected BlockdevSnapshotSync to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*BlockdevSnapshotSyncWrapper
		}{
			Type:                        u.Type,
			BlockdevSnapshotSyncWrapper: u.BlockdevSnapshotSync,
		})
	case "drive-backup":
		if u.DriveBackup == nil {
			return nil, fmt.Errorf("expected DriveBackup to be set")
		}

		return json.Marshal(struct {
			Type TransactionActionKind `json:"type"`
			*DriveBackupWrapper
		}{
			Type:               u.Type,
			DriveBackupWrapper: u.DriveBackup,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// TransactionProperties
//
// Optional arguments to modify the behavior of a Transaction.
type TransactionProperties struct {
	// CompletionMode Controls how jobs launched asynchronously by Actions will complete or fail as a group. See @ActionCompletionMode for details.
	CompletionMode *ActionCompletionMode `json:"completion-mode,omitempty"`
}

// Transaction
//
// Executes a number of transactionable QMP commands atomically. If any operation fails, then the entire set of actions will be abandoned and the appropriate error returned. For external snapshots, the dictionary contains the device, the file to use for the new snapshot, and the format. The default format, if not specified, is qcow2. Each new snapshot defaults to being created by QEMU (wiping any contents if the file already exists), but it is also possible to reuse an externally-created file. In the latter case, you should ensure that the new image file has the same contents as the current one; QEMU cannot perform any meaningful check. Typically this is achieved by using the current image file as the backing file for the new image. On failure, the original disks pre-snapshot attempt will be used. For internal snapshots, the dictionary contains the device and the snapshot's name. If an internal snapshot matching name already exists, the request will be rejected. Only some image formats support it, for example, qcow2, and rbd, On failure, qemu will try delete the newly created internal snapshot in the transaction. When an I/O error occurs during deletion, the user needs to fix it later with qemu-img or other command.
type Transaction struct {
	// Actions List of @TransactionAction; information needed for the respective operations.
	Actions []TransactionAction `json:"actions"`
	// Properties structure of additional options to control the execution of the transaction. See @TransactionProperties for additional detail.
	Properties *TransactionProperties `json:"properties,omitempty"`
}

func (Transaction) Command() string {
	return "transaction"
}

func (cmd Transaction) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "transaction", cmd, nil)
}

// QmpCapabilities
//
// Enable QMP capabilities.
type QmpCapabilities struct {
	// Enable An optional list of QMPCapability values to enable. The client must not enable any capability that is not mentioned in the QMP greeting message. If the field is not provided, it means no QMP capabilities will be enabled. (since 2.12)
	Enable []QMPCapability `json:"enable,omitempty"`
}

func (QmpCapabilities) Command() string {
	return "qmp_capabilities"
}

func (cmd QmpCapabilities) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "qmp_capabilities", cmd, nil)
}

// QMPCapability Enumeration of capabilities to be advertised during initial client connection, used for agreeing on particular QMP extension behaviors.
type QMPCapability string

const (
	// QMPCapabilityOob QMP ability to support out-of-band requests. (Please refer to qmp-spec.rst for more information on OOB)
	QMPCapabilityOob QMPCapability = "oob"
)

// VersionTriple
//
// A three-part version number.
type VersionTriple struct {
	// Major The major version number.
	Major int64 `json:"major"`
	// Minor The minor version number.
	Minor int64 `json:"minor"`
	// Micro The micro version number.
	Micro int64 `json:"micro"`
}

// VersionInfo
//
// A description of QEMU's version.
type VersionInfo struct {
	// Qemu The version of QEMU. By current convention, a micro version of 50 signifies a development branch. A micro version greater than or equal to 90 signifies a release candidate for the next minor version. A micro version of less than 50 signifies a stable release.
	Qemu VersionTriple `json:"qemu"`
	// Package QEMU will always set this field to an empty string. Downstream versions of QEMU should set this to a non-empty string. The exact format depends on the downstream however it highly recommended that a unique name is used.
	Package string `json:"package"`
}

// QueryVersion
//
// Returns the current version of QEMU.
type QueryVersion struct {
}

func (QueryVersion) Command() string {
	return "query-version"
}

func (cmd QueryVersion) Execute(ctx context.Context, client api.Client) (VersionInfo, error) {
	var ret VersionInfo

	return ret, client.Execute(ctx, "query-version", cmd, &ret)
}

// CommandInfo
//
// Information about a QMP command
type CommandInfo struct {
	// Name The command name
	Name string `json:"name"`
}

// QueryCommands
//
// Return a list of supported QMP commands by this server
type QueryCommands struct {
}

func (QueryCommands) Command() string {
	return "query-commands"
}

func (cmd QueryCommands) Execute(ctx context.Context, client api.Client) ([]CommandInfo, error) {
	var ret []CommandInfo

	return ret, client.Execute(ctx, "query-commands", cmd, &ret)
}

// Quit
//
// This command will cause the QEMU process to exit gracefully. While every attempt is made to send the QMP response before terminating, this is not guaranteed. When using this interface, a premature EOF would not be unexpected.
type Quit struct {
}

func (Quit) Command() string {
	return "quit"
}

func (cmd Quit) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "quit", cmd, nil)
}

// MonitorMode An enumeration of monitor modes.
type MonitorMode string

const (
	// MonitorModeReadline HMP monitor (human-oriented command line interface)
	MonitorModeReadline MonitorMode = "readline"
	// MonitorModeControl QMP monitor (JSON-based machine interface)
	MonitorModeControl MonitorMode = "control"
)

// MonitorOptions
//
// Options to be used for adding a new monitor.
type MonitorOptions struct {
	// Id Name of the monitor
	Id *string `json:"id,omitempty"`
	// Mode Selects the monitor mode (default: readline in the system emulator, control in qemu-storage-daemon)
	Mode *MonitorMode `json:"mode,omitempty"`
	// Pretty Enables pretty printing (QMP only)
	Pretty *bool `json:"pretty,omitempty"`
	// Chardev Name of a character device to expose the monitor on
	Chardev string `json:"chardev"`
}

// QueryQmpSchema
//
// Command query-qmp-schema exposes the QMP wire ABI as an array of SchemaInfo. This lets QMP clients figure out what commands and events are available in this QEMU, and their parameters and results. However, the SchemaInfo can't reflect all the rules and restrictions that apply to QMP. It's interface introspection (figuring out what's there), not interface specification. The specification is in the QAPI schema. Furthermore, while we strive to keep the QMP wire format backwards-compatible across qemu versions, the introspection output is not guaranteed to have the same stability. For example, one version of qemu may list an object member as an optional non-variant, while another lists the same member only through the object's variants; or the type of a member may change from a generic string into a specific enum or from one specific type into an alternate that includes the original type alongside something else.
type QueryQmpSchema struct {
}

func (QueryQmpSchema) Command() string {
	return "query-qmp-schema"
}

func (cmd QueryQmpSchema) Execute(ctx context.Context, client api.Client) ([]SchemaInfo, error) {
	var ret []SchemaInfo

	return ret, client.Execute(ctx, "query-qmp-schema", cmd, &ret)
}

// SchemaMetaType This is a @SchemaInfo's meta type, i.e. the kind of entity it describes.
type SchemaMetaType string

const (
	// SchemaMetaTypeBuiltin a predefined type such as 'int' or 'bool'.
	SchemaMetaTypeBuiltin SchemaMetaType = "builtin"
	// SchemaMetaTypeEnum an enumeration type
	SchemaMetaTypeEnum SchemaMetaType = "enum"
	// SchemaMetaTypeArray an array type
	SchemaMetaTypeArray SchemaMetaType = "array"
	// SchemaMetaTypeObject an object type (struct or union)
	SchemaMetaTypeObject SchemaMetaType = "object"
	// SchemaMetaTypeAlternate an alternate type
	SchemaMetaTypeAlternate SchemaMetaType = "alternate"
	// SchemaMetaTypeCommand a QMP command
	SchemaMetaTypeCommand SchemaMetaType = "command"
	// SchemaMetaTypeEvent a QMP event
	SchemaMetaTypeEvent SchemaMetaType = "event"
)

// SchemaInfo
type SchemaInfo struct {
	// Discriminator: meta-type

	// Name the entity's name, inherited from @base. The SchemaInfo is always referenced by this name. Commands and events have the name defined in the QAPI schema. Unlike command and event names, type names are not part of the wire ABI. Consequently, type names are meaningless strings here, although they are still guaranteed unique regardless of @meta-type.
	Name string `json:"name"`
	// MetaType the entity's meta type, inherited from @base.
	MetaType SchemaMetaType `json:"meta-type"`
	// Features names of features associated with the entity, in no particular order. (since 4.1 for object types, 4.2 for commands, 5.0 for the rest)
	Features []string `json:"features,omitempty"`

	Builtin   *SchemaInfoBuiltin   `json:"-"`
	Enum      *SchemaInfoEnum      `json:"-"`
	Array     *SchemaInfoArray     `json:"-"`
	Object    *SchemaInfoObject    `json:"-"`
	Alternate *SchemaInfoAlternate `json:"-"`
	Command   *SchemaInfoCommand   `json:"-"`
	Event     *SchemaInfoEvent     `json:"-"`
}

func (u SchemaInfo) MarshalJSON() ([]byte, error) {
	switch u.MetaType {
	case "builtin":
		if u.Builtin == nil {
			return nil, fmt.Errorf("expected Builtin to be set")
		}

		return json.Marshal(struct {
			Name     string         `json:"name"`
			MetaType SchemaMetaType `json:"meta-type"`
			Features []string       `json:"features,omitempty"`
			*SchemaInfoBuiltin
		}{
			Name:              u.Name,
			MetaType:          u.MetaType,
			Features:          u.Features,
			SchemaInfoBuiltin: u.Builtin,
		})
	case "enum":
		if u.Enum == nil {
			return nil, fmt.Errorf("expected Enum to be set")
		}

		return json.Marshal(struct {
			Name     string         `json:"name"`
			MetaType SchemaMetaType `json:"meta-type"`
			Features []string       `json:"features,omitempty"`
			*SchemaInfoEnum
		}{
			Name:           u.Name,
			MetaType:       u.MetaType,
			Features:       u.Features,
			SchemaInfoEnum: u.Enum,
		})
	case "array":
		if u.Array == nil {
			return nil, fmt.Errorf("expected Array to be set")
		}

		return json.Marshal(struct {
			Name     string         `json:"name"`
			MetaType SchemaMetaType `json:"meta-type"`
			Features []string       `json:"features,omitempty"`
			*SchemaInfoArray
		}{
			Name:            u.Name,
			MetaType:        u.MetaType,
			Features:        u.Features,
			SchemaInfoArray: u.Array,
		})
	case "object":
		if u.Object == nil {
			return nil, fmt.Errorf("expected Object to be set")
		}

		return json.Marshal(struct {
			Name     string         `json:"name"`
			MetaType SchemaMetaType `json:"meta-type"`
			Features []string       `json:"features,omitempty"`
			*SchemaInfoObject
		}{
			Name:             u.Name,
			MetaType:         u.MetaType,
			Features:         u.Features,
			SchemaInfoObject: u.Object,
		})
	case "alternate":
		if u.Alternate == nil {
			return nil, fmt.Errorf("expected Alternate to be set")
		}

		return json.Marshal(struct {
			Name     string         `json:"name"`
			MetaType SchemaMetaType `json:"meta-type"`
			Features []string       `json:"features,omitempty"`
			*SchemaInfoAlternate
		}{
			Name:                u.Name,
			MetaType:            u.MetaType,
			Features:            u.Features,
			SchemaInfoAlternate: u.Alternate,
		})
	case "command":
		if u.Command == nil {
			return nil, fmt.Errorf("expected Command to be set")
		}

		return json.Marshal(struct {
			Name     string         `json:"name"`
			MetaType SchemaMetaType `json:"meta-type"`
			Features []string       `json:"features,omitempty"`
			*SchemaInfoCommand
		}{
			Name:              u.Name,
			MetaType:          u.MetaType,
			Features:          u.Features,
			SchemaInfoCommand: u.Command,
		})
	case "event":
		if u.Event == nil {
			return nil, fmt.Errorf("expected Event to be set")
		}

		return json.Marshal(struct {
			Name     string         `json:"name"`
			MetaType SchemaMetaType `json:"meta-type"`
			Features []string       `json:"features,omitempty"`
			*SchemaInfoEvent
		}{
			Name:            u.Name,
			MetaType:        u.MetaType,
			Features:        u.Features,
			SchemaInfoEvent: u.Event,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// SchemaInfoBuiltin
//
// Additional SchemaInfo members for meta-type 'builtin'.
type SchemaInfoBuiltin struct {
	// JsonType the JSON type used for this type on the wire.
	JsonType JSONType `json:"json-type"`
}

// JSONType The four primitive and two structured types according to RFC 8259 section 1, plus 'int' (split off 'number'), plus the obvious top type 'value'.
type JSONType string

const (
	JSONTypeString  JSONType = "string"
	JSONTypeNumber  JSONType = "number"
	JSONTypeInt     JSONType = "int"
	JSONTypeBoolean JSONType = "boolean"
	JSONTypeNull    JSONType = "null"
	JSONTypeObject  JSONType = "object"
	JSONTypeArray   JSONType = "array"
	JSONTypeValue   JSONType = "value"
)

// SchemaInfoEnum
//
// Additional SchemaInfo members for meta-type 'enum'.
type SchemaInfoEnum struct {
	// Members the enum type's members, in no particular order (since 6.2).
	Members []SchemaInfoEnumMember `json:"members"`
	// Values the enumeration type's member names, in no particular order. Redundant with @members. Just for backward compatibility.
	Values []string `json:"values"`
}

// SchemaInfoEnumMember
//
// An object member.
type SchemaInfoEnumMember struct {
	// Name the member's name, as defined in the QAPI schema.
	Name string `json:"name"`
	// Features names of features associated with the member, in no particular order.
	Features []string `json:"features,omitempty"`
}

// SchemaInfoArray
//
// Additional SchemaInfo members for meta-type 'array'.
type SchemaInfoArray struct {
	// ElementType the array type's element type. Values of this type are JSON array on the wire.
	ElementType string `json:"element-type"`
}

// SchemaInfoObject
//
// Additional SchemaInfo members for meta-type 'object'.
type SchemaInfoObject struct {
	// Members the object type's (non-variant) members, in no particular order.
	Members []SchemaInfoObjectMember `json:"members"`
	// Tag the name of the member serving as type tag. An element of @members with this name must exist.
	Tag *string `json:"tag,omitempty"`
	// Variants variant members, i.e. additional members that depend on the type tag's value. Present exactly when @tag is present. The variants are in no particular order, and may even differ from the order of the values of the enum type of the @tag. Values of this type are JSON object on the wire.
	Variants []SchemaInfoObjectVariant `json:"variants,omitempty"`
}

// SchemaInfoObjectMember
//
// An object member.
type SchemaInfoObjectMember struct {
	// Name the member's name, as defined in the QAPI schema.
	Name string `json:"name"`
	// Type the name of the member's type.
	Type string `json:"type"`
	// Default default when used as command parameter. If absent, the parameter is mandatory. If present, the value must be null. The parameter is optional, and behavior when it's missing is not
	Default *any `json:"default,omitempty"`
	// Features names of features associated with the member, in no particular order. (since 5.0)
	Features []string `json:"features,omitempty"`
}

// SchemaInfoObjectVariant
//
// The variant members for a value of the type tag.
type SchemaInfoObjectVariant struct {
	// Case a value of the type tag.
	Case string `json:"case"`
	// Type the name of the object type that provides the variant members when the type tag has value @case.
	Type string `json:"type"`
}

// SchemaInfoAlternate
//
// Additional SchemaInfo members for meta-type 'alternate'.
type SchemaInfoAlternate struct {
	// Members the alternate type's members, in no particular order. The members' wire encoding is distinct, see
	Members []SchemaInfoAlternateMember `json:"members"`
}

// SchemaInfoAlternateMember
//
// An alternate member.
type SchemaInfoAlternateMember struct {
	// Type the name of the member's type.
	Type string `json:"type"`
}

// SchemaInfoCommand
//
// Additional SchemaInfo members for meta-type 'command'.
type SchemaInfoCommand struct {
	// ArgType the name of the object type that provides the command's parameters.
	ArgType string `json:"arg-type"`
	// RetType the name of the command's result type.
	RetType string `json:"ret-type"`
	// AllowOob whether the command allows out-of-band execution,
	AllowOob *bool `json:"allow-oob,omitempty"`
}

// SchemaInfoEvent
//
// Additional SchemaInfo members for meta-type 'event'.
type SchemaInfoEvent struct {
	// ArgType the name of the object type that provides the event's parameters.
	ArgType string `json:"arg-type"`
}

// ObjectPropertyInfo
type ObjectPropertyInfo struct {
	// Name the name of the property
	Name string `json:"name"`
	// Type the type of the property. This will typically come in one of
	Type string `json:"type"`
	// Description if specified, the description of the property.
	Description *string `json:"description,omitempty"`
	// DefaultValue the default value, if any (since 5.0)
	DefaultValue *any `json:"default-value,omitempty"`
}

// QomList
//
// This command will list any properties of a object given a path in the object model.
type QomList struct {
	// Path the path within the object model. See @qom-get for a description of this parameter.
	Path string `json:"path"`
}

func (QomList) Command() string {
	return "qom-list"
}

func (cmd QomList) Execute(ctx context.Context, client api.Client) ([]ObjectPropertyInfo, error) {
	var ret []ObjectPropertyInfo

	return ret, client.Execute(ctx, "qom-list", cmd, &ret)
}

// QomGet
//
// This command will get a property from a object model path and return the value.
type QomGet struct {
	// Path The path within the object model. There are two forms of supported paths--absolute and partial paths. Absolute paths are derived from the root object and can follow child<> or link<> properties. Since they can follow link<> properties, they can be arbitrarily long. Absolute paths look like absolute filenames and are prefixed with a leading slash. Partial paths look like relative filenames. They do not begin with a prefix. The matching rules for partial paths are subtle but designed to make specifying objects easy. At each level of the composition tree, the partial path is matched as an absolute path. The first match is not returned. At least two matches are searched for. A successful result is only returned if only one match is found. If more than one match is found, a flag is return to indicate that the match was ambiguous.
	Path string `json:"path"`
	// Property The property name to read
	Property string `json:"property"`
}

func (QomGet) Command() string {
	return "qom-get"
}

func (cmd QomGet) Execute(ctx context.Context, client api.Client) (any, error) {
	var ret any

	return ret, client.Execute(ctx, "qom-get", cmd, &ret)
}

// QomSet
//
// This command will set a property from a object model path.
type QomSet struct {
	// Path see @qom-get for a description of this parameter
	Path string `json:"path"`
	// Property the property name to set
	Property string `json:"property"`
	// Value a value who's type is appropriate for the property type. See @qom-get for a description of type mapping.
	Value any `json:"value"`
}

func (QomSet) Command() string {
	return "qom-set"
}

func (cmd QomSet) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "qom-set", cmd, nil)
}

// ObjectTypeInfo
//
// This structure describes a search result from @qom-list-types
type ObjectTypeInfo struct {
	// Name the type name found in the search
	Name string `json:"name"`
	// Abstract the type is abstract and can't be directly instantiated. Omitted if false. (since 2.10)
	Abstract *bool `json:"abstract,omitempty"`
	// Parent Name of parent type, if any (since 2.10)
	Parent *string `json:"parent,omitempty"`
}

// QomListTypes
//
// This command will return a list of types given search parameters
type QomListTypes struct {
	// Implements if specified, only return types that implement this type name
	Implements *string `json:"implements,omitempty"`
	// Abstract if true, include abstract types in the results
	Abstract *bool `json:"abstract,omitempty"`
}

func (QomListTypes) Command() string {
	return "qom-list-types"
}

func (cmd QomListTypes) Execute(ctx context.Context, client api.Client) ([]ObjectTypeInfo, error) {
	var ret []ObjectTypeInfo

	return ret, client.Execute(ctx, "qom-list-types", cmd, &ret)
}

// QomListProperties
//
// List properties associated with a QOM object.
type QomListProperties struct {
	// Typename the type name of an object
	Typename string `json:"typename"`
}

func (QomListProperties) Command() string {
	return "qom-list-properties"
}

func (cmd QomListProperties) Execute(ctx context.Context, client api.Client) ([]ObjectPropertyInfo, error) {
	var ret []ObjectPropertyInfo

	return ret, client.Execute(ctx, "qom-list-properties", cmd, &ret)
}

// CanHostSocketcanProperties
//
// Properties for can-host-socketcan objects.
type CanHostSocketcanProperties struct {
	// If interface name of the host system CAN bus to connect to
	If string `json:"if"`
	// Canbus object ID of the can-bus object to connect to the host interface
	Canbus string `json:"canbus"`
}

// ColoCompareProperties
//
// Properties for colo-compare objects.
type ColoCompareProperties struct {
	// PrimaryIn name of the character device backend to use for the primary input (incoming packets are redirected to @outdev)
	PrimaryIn string `json:"primary_in"`
	// SecondaryIn name of the character device backend to use for secondary input (incoming packets are only compared to the input on @primary_in and then dropped)
	SecondaryIn string `json:"secondary_in"`
	// Outdev name of the character device backend to use for output
	Outdev string `json:"outdev"`
	// Iothread name of the iothread to run in
	Iothread string `json:"iothread"`
	// NotifyDev name of the character device backend to be used to communicate with the remote colo-frame (only for Xen COLO)
	NotifyDev *string `json:"notify_dev,omitempty"`
	// CompareTimeout the maximum time to hold a packet from @primary_in for comparison with an incoming packet on @secondary_in in
	CompareTimeout *uint64 `json:"compare_timeout,omitempty"`
	// ExpiredScanCycle the interval at which colo-compare checks whether packets from @primary have timed out, in milliseconds
	ExpiredScanCycle *uint32 `json:"expired_scan_cycle,omitempty"`
	// MaxQueueSize the maximum number of packets to keep in the queue for comparing with incoming packets from @secondary_in. If the queue is full and additional packets are received, the
	MaxQueueSize *uint32 `json:"max_queue_size,omitempty"`
	// VnetHdrSupport if true, vnet header support is enabled
	VnetHdrSupport *bool `json:"vnet_hdr_support,omitempty"`
}

// CryptodevBackendProperties
//
// Properties for cryptodev-backend and cryptodev-backend-builtin objects.
type CryptodevBackendProperties struct {
	// Queues the number of queues for the cryptodev backend. Ignored for cryptodev-backend and must be 1 for
	Queues *uint32 `json:"queues,omitempty"`
	// ThrottleBps limit total bytes per second (Since 8.0)
	ThrottleBps *uint64 `json:"throttle-bps,omitempty"`
	// ThrottleOps limit total operations per second (Since 8.0)
	ThrottleOps *uint64 `json:"throttle-ops,omitempty"`
}

// CryptodevVhostUserProperties
//
// Properties for cryptodev-vhost-user objects.
type CryptodevVhostUserProperties struct {
	CryptodevBackendProperties

	// Chardev the name of a Unix domain socket character device that connects to the vhost-user server
	Chardev string `json:"chardev"`
}

// DBusVMStateProperties
//
// Properties for dbus-vmstate objects.
type DBusVMStateProperties struct {
	// Addr the name of the DBus bus to connect to
	Addr string `json:"addr"`
	// IdList a comma separated list of DBus IDs of helpers whose data should be included in the VM state on migration
	IdList *string `json:"id-list,omitempty"`
}

// NetfilterInsert Indicates where to insert a netfilter relative to a given other filter.
type NetfilterInsert string

const (
	// NetfilterInsertBefore insert before the specified filter
	NetfilterInsertBefore NetfilterInsert = "before"
	// NetfilterInsertBehind insert behind the specified filter
	NetfilterInsertBehind NetfilterInsert = "behind"
)

// NetfilterProperties
//
// Properties for objects of classes derived from netfilter.
type NetfilterProperties struct {
	// Netdev id of the network device backend to filter
	Netdev string `json:"netdev"`
	// Queue indicates which queue(s) to filter (default: all)
	Queue *NetFilterDirection `json:"queue,omitempty"`
	// Status indicates whether the filter is enabled ("on") or disabled
	Status *string `json:"status,omitempty"`
	// Position specifies where the filter should be inserted in the filter list. "head" means the filter is inserted at the head of the filter list, before any existing filters. "tail" means the filter is inserted at the tail of the filter list, behind any existing filters (default). "id=<id>" means the filter is inserted before or behind the filter specified by <id>,
	Position *string `json:"position,omitempty"`
	// Insert where to insert the filter relative to the filter given in @position. Ignored if @position is "head" or "tail".
	Insert *NetfilterInsert `json:"insert,omitempty"`
}

// FilterBufferProperties
//
// Properties for filter-buffer objects.
type FilterBufferProperties struct {
	NetfilterProperties

	// Interval a non-zero interval in microseconds. All packets arriving in the given interval are delayed until the end of the interval.
	Interval uint32 `json:"interval"`
}

// FilterDumpProperties
//
// Properties for filter-dump objects.
type FilterDumpProperties struct {
	NetfilterProperties

	// File the filename where the dumped packets should be stored
	File string `json:"file"`
	// Maxlen maximum number of bytes in a packet that are stored
	Maxlen *uint32 `json:"maxlen,omitempty"`
}

// FilterMirrorProperties
//
// Properties for filter-mirror objects.
type FilterMirrorProperties struct {
	NetfilterProperties

	// Outdev the name of a character device backend to which all incoming packets are mirrored
	Outdev string `json:"outdev"`
	// VnetHdrSupport if true, vnet header support is enabled
	VnetHdrSupport *bool `json:"vnet_hdr_support,omitempty"`
}

// FilterRedirectorProperties
//
// Properties for filter-redirector objects. At least one of @indev or @outdev must be present. If both are present, they must not refer to the same character device backend.
type FilterRedirectorProperties struct {
	NetfilterProperties

	// Indev the name of a character device backend from which packets are received and redirected to the filtered network device
	Indev *string `json:"indev,omitempty"`
	// Outdev the name of a character device backend to which all incoming packets are redirected
	Outdev *string `json:"outdev,omitempty"`
	// VnetHdrSupport if true, vnet header support is enabled
	VnetHdrSupport *bool `json:"vnet_hdr_support,omitempty"`
}

// FilterRewriterProperties
//
// Properties for filter-rewriter objects.
type FilterRewriterProperties struct {
	NetfilterProperties

	// VnetHdrSupport if true, vnet header support is enabled
	VnetHdrSupport *bool `json:"vnet_hdr_support,omitempty"`
}

// InputBarrierProperties
//
// Properties for input-barrier objects.
type InputBarrierProperties struct {
	// Name the screen name as declared in the screens section of barrier.conf
	Name string `json:"name"`
	// Server hostname of the Barrier server (default: "localhost")
	Server *string `json:"server,omitempty"`
	// Port TCP port of the Barrier server (default: "24800")
	Port *string `json:"port,omitempty"`
	// Origin x coordinate of the leftmost pixel on the guest screen
	Origin *string `json:"x-origin,omitempty"`
	// YOrigin y coordinate of the topmost pixel on the guest screen
	YOrigin *string `json:"y-origin,omitempty"`
	// Width the width of secondary screen in pixels (default: "1920")
	Width *string `json:"width,omitempty"`
	// Height the height of secondary screen in pixels (default: "1080")
	Height *string `json:"height,omitempty"`
}

// InputLinuxProperties
//
// Properties for input-linux objects.
type InputLinuxProperties struct {
	// Evdev the path of the host evdev device to use
	Evdev string `json:"evdev"`
	// GrabAll if true, grab is toggled for all devices (e.g. both
	GrabAll *bool `json:"grab_all,omitempty"`
	// Repeat enables auto-repeat events (default: false)
	Repeat *bool `json:"repeat,omitempty"`
	// GrabToggle the key or key combination that toggles device grab
	GrabToggle *GrabToggleKeys `json:"grab-toggle,omitempty"`
}

// EventLoopBaseProperties
//
// Common properties for event loops
type EventLoopBaseProperties struct {
	// AioMaxBatch maximum number of requests in a batch for the AIO engine, 0 means that the engine will use its default.
	AioMaxBatch *int64 `json:"aio-max-batch,omitempty"`
	// ThreadPoolMin minimum number of threads reserved in the thread
	ThreadPoolMin *int64 `json:"thread-pool-min,omitempty"`
	// ThreadPoolMax maximum number of threads the thread pool can
	ThreadPoolMax *int64 `json:"thread-pool-max,omitempty"`
}

// IothreadProperties
//
// Properties for iothread objects.
type IothreadProperties struct {
	EventLoopBaseProperties

	// PollMaxNs the maximum number of nanoseconds to busy wait for
	PollMaxNs *int64 `json:"poll-max-ns,omitempty"`
	// PollGrow the multiplier used to increase the polling time when the algorithm detects it is missing events due to not polling
	PollGrow *int64 `json:"poll-grow,omitempty"`
	// PollShrink the divisor used to decrease the polling time when the algorithm detects it is spending too long polling without
	PollShrink *int64 `json:"poll-shrink,omitempty"`
}

// MainLoopProperties
//
// Properties for the main-loop object.
type MainLoopProperties struct {
	EventLoopBaseProperties
}

// MemoryBackendProperties
//
// Properties for objects of classes derived from memory-backend.
type MemoryBackendProperties struct {
	// Dump if true, include the memory in core dumps (default depends on the machine type)
	Dump *bool `json:"dump,omitempty"`
	// HostNodes the list of NUMA host nodes to bind the memory to
	HostNodes []uint16 `json:"host-nodes,omitempty"`
	// Merge if true, mark the memory as mergeable (default depends on the machine type)
	Merge *bool `json:"merge,omitempty"`
	// Policy the NUMA policy (default: 'default')
	Policy *HostMemPolicy `json:"policy,omitempty"`
	// Prealloc if true, preallocate memory (default: false)
	Prealloc *bool `json:"prealloc,omitempty"`
	// PreallocThreads number of CPU threads to use for prealloc
	PreallocThreads *uint32 `json:"prealloc-threads,omitempty"`
	// PreallocContext thread context to use for creation of
	PreallocContext *string `json:"prealloc-context,omitempty"`
	// Share if false, the memory is private to QEMU; if true, it is
	Share *bool `json:"share,omitempty"`
	// Reserve if true, reserve swap space (or huge pages) if applicable
	Reserve *bool `json:"reserve,omitempty"`
	// Size size of the memory region in bytes
	Size uint64 `json:"size"`
	// UseCanonicalPathForRamblockId if true, the canonical path is used for ramblock-id. Disable this for 4.0 machine types or older to allow migration with newer QEMU versions.
	UseCanonicalPathForRamblockId *bool `json:"x-use-canonical-path-for-ramblock-id,omitempty"`
}

// MemoryBackendFileProperties
//
// Properties for memory-backend-file objects.
type MemoryBackendFileProperties struct {
	MemoryBackendProperties

	// Align the base address alignment when QEMU mmap(2)s @mem-path. Some backend stores specified by @mem-path require an alignment different than the default one used by QEMU, e.g. the device DAX /dev/dax0.0 requires 2M alignment rather than 4K. In such cases, users can specify the required alignment via this option. 0 selects a default alignment (currently the page size).
	Align *uint64 `json:"align,omitempty"`
	// Offset the offset into the target file that the region starts at. You can use this option to back multiple regions with a single file. Must be a multiple of the page size.
	Offset *uint64 `json:"offset,omitempty"`
	// DiscardData if true, the file contents can be destroyed when QEMU exits, to avoid unnecessarily flushing data to the backing file. Note that @discard-data is only an optimization, and QEMU might not discard file contents if it aborts unexpectedly or is
	DiscardData *bool `json:"discard-data,omitempty"`
	// MemPath the path to either a shared memory or huge page filesystem mount
	MemPath string `json:"mem-path"`
	// Pmem specifies whether the backing file specified by @mem-path is in host persistent memory that can be accessed using the SNIA NVM programming model (e.g. Intel NVDIMM).
	Pmem *bool `json:"pmem,omitempty"`
	// Readonly if true, the backing file is opened read-only; if false,
	Readonly *bool `json:"readonly,omitempty"`
	// Rom whether to create Read Only Memory (ROM) that cannot be modified by the VM. Any write attempts to such ROM will be denied. Most use cases want writable RAM instead of ROM. However, selected use cases, like R/O NVDIMMs, can benefit from ROM. If set to 'on', create ROM; if set to 'off', create writable RAM; if set to 'auto', the value of the @readonly property is used. This property is primarily helpful when we want to have proper RAM in configurations that would traditionally create ROM before this
	Rom *OnOffAuto `json:"rom,omitempty"`
}

// MemoryBackendMemfdProperties
//
// Properties for memory-backend-memfd objects. The @share boolean option is true by default with memfd.
type MemoryBackendMemfdProperties struct {
	MemoryBackendProperties

	// Hugetlb if true, the file to be created resides in the hugetlbfs
	Hugetlb *bool `json:"hugetlb,omitempty"`
	// Hugetlbsize the hugetlb page size on systems that support multiple hugetlb page sizes (it must be a power of 2 value supported by the system). 0 selects a default page size. This option is
	Hugetlbsize *uint64 `json:"hugetlbsize,omitempty"`
	// Seal if true, create a sealed-file, which will block further
	Seal *bool `json:"seal,omitempty"`
}

// MemoryBackendEpcProperties
//
// Properties for memory-backend-epc objects. The @share boolean option is true by default with epc The @merge boolean option is false by default with epc The @dump boolean option is false by default with epc
type MemoryBackendEpcProperties struct {
	MemoryBackendProperties
}

// PrManagerHelperProperties
//
// Properties for pr-manager-helper objects.
type PrManagerHelperProperties struct {
	// Path the path to a Unix domain socket for connecting to the external helper
	Path string `json:"path"`
}

// QtestProperties
//
// Properties for qtest objects.
type QtestProperties struct {
	// Chardev the chardev to be used to receive qtest commands on.
	Chardev string `json:"chardev"`
	// Log the path to a log file
	Log *string `json:"log,omitempty"`
}

// RemoteObjectProperties
//
// Properties for x-remote-object objects.
type RemoteObjectProperties struct {
	// Fd file descriptor name previously passed via 'getfd' command
	Fd string `json:"fd"`
	// Devid the id of the device to be associated with the file descriptor
	Devid string `json:"devid"`
}

// VfioUserServerProperties
//
// Properties for x-vfio-user-server objects.
type VfioUserServerProperties struct {
	// Socket socket to be used by the libvfio-user library
	Socket SocketAddress `json:"socket"`
	// Device the ID of the device to be emulated at the server
	Device string `json:"device"`
}

// IOMMUFDProperties
//
// Properties for iommufd objects.
type IOMMUFDProperties struct {
	// Fd file descriptor name previously passed via 'getfd' command, which represents a pre-opened /dev/iommu. This allows the iommufd object to be shared accross several subsystems (VFIO, VDPA, ...), and the file descriptor to be shared
	Fd *string `json:"fd,omitempty"`
}

// RngProperties
//
// Properties for objects of classes derived from rng.
type RngProperties struct {
	// Opened if true, the device is opened immediately when applying this option and will probably fail when processing the next option. Don't use; only provided for compatibility.
	Opened *bool `json:"opened,omitempty"`
}

// RngEgdProperties
//
// Properties for rng-egd objects.
type RngEgdProperties struct {
	RngProperties

	// Chardev the name of a character device backend that provides the connection to the RNG daemon
	Chardev string `json:"chardev"`
}

// RngRandomProperties
//
// Properties for rng-random objects.
type RngRandomProperties struct {
	RngProperties

	// Filename the filename of the device on the host to obtain entropy
	Filename *string `json:"filename,omitempty"`
}

// SevGuestProperties
//
// Properties for sev-guest objects.
type SevGuestProperties struct {
	// SevDevice SEV device to use (default: "/dev/sev")
	SevDevice *string `json:"sev-device,omitempty"`
	// DhCertFile guest owners DH certificate (encoded with base64)
	DhCertFile *string `json:"dh-cert-file,omitempty"`
	// SessionFile guest owners session parameters (encoded with base64)
	SessionFile *string `json:"session-file,omitempty"`
	// Policy SEV policy value (default: 0x1)
	Policy *uint32 `json:"policy,omitempty"`
	// Handle SEV firmware handle (default: 0)
	Handle *uint32 `json:"handle,omitempty"`
	// Cbitpos C-bit location in page table entry (default: 0)
	Cbitpos *uint32 `json:"cbitpos,omitempty"`
	// ReducedPhysBits number of bits in physical addresses that become unavailable when SEV is enabled
	ReducedPhysBits uint32 `json:"reduced-phys-bits"`
	// KernelHashes if true, add hashes of kernel/initrd/cmdline to a designated guest firmware page for measured boot with -kernel
	KernelHashes *bool `json:"kernel-hashes,omitempty"`
}

// ThreadContextProperties
//
// Properties for thread context objects.
type ThreadContextProperties struct {
	// CpuAffinity the list of host CPU numbers used as CPU affinity for
	CpuAffinity []uint16 `json:"cpu-affinity,omitempty"`
	// NodeAffinity the list of host node numbers that will be resolved to a list of host CPU numbers used as CPU affinity. This is a shortcut for specifying the list of host CPU numbers belonging to the host nodes manually by setting @cpu-affinity.
	NodeAffinity []uint16 `json:"node-affinity,omitempty"`
}

// ObjectType
type ObjectType string

const (
	ObjectTypeAuthzList               ObjectType = "authz-list"
	ObjectTypeAuthzListfile           ObjectType = "authz-listfile"
	ObjectTypeAuthzPam                ObjectType = "authz-pam"
	ObjectTypeAuthzSimple             ObjectType = "authz-simple"
	ObjectTypeCanBus                  ObjectType = "can-bus"
	ObjectTypeCanHostSocketcan        ObjectType = "can-host-socketcan"
	ObjectTypeColoCompare             ObjectType = "colo-compare"
	ObjectTypeCryptodevBackend        ObjectType = "cryptodev-backend"
	ObjectTypeCryptodevBackendBuiltin ObjectType = "cryptodev-backend-builtin"
	ObjectTypeCryptodevBackendLkcf    ObjectType = "cryptodev-backend-lkcf"
	ObjectTypeCryptodevVhostUser      ObjectType = "cryptodev-vhost-user"
	ObjectTypeDbusVmstate             ObjectType = "dbus-vmstate"
	ObjectTypeFilterBuffer            ObjectType = "filter-buffer"
	ObjectTypeFilterDump              ObjectType = "filter-dump"
	ObjectTypeFilterMirror            ObjectType = "filter-mirror"
	ObjectTypeFilterRedirector        ObjectType = "filter-redirector"
	ObjectTypeFilterReplay            ObjectType = "filter-replay"
	ObjectTypeFilterRewriter          ObjectType = "filter-rewriter"
	ObjectTypeInputBarrier            ObjectType = "input-barrier"
	ObjectTypeInputLinux              ObjectType = "input-linux"
	ObjectTypeIommufd                 ObjectType = "iommufd"
	ObjectTypeIothread                ObjectType = "iothread"
	ObjectTypeMainLoop                ObjectType = "main-loop"
	ObjectTypeMemoryBackendEpc        ObjectType = "memory-backend-epc"
	ObjectTypeMemoryBackendFile       ObjectType = "memory-backend-file"
	ObjectTypeMemoryBackendMemfd      ObjectType = "memory-backend-memfd"
	ObjectTypeMemoryBackendRam        ObjectType = "memory-backend-ram"
	ObjectTypePefGuest                ObjectType = "pef-guest"
	ObjectTypePrManagerHelper         ObjectType = "pr-manager-helper"
	ObjectTypeQtest                   ObjectType = "qtest"
	ObjectTypeRngBuiltin              ObjectType = "rng-builtin"
	ObjectTypeRngEgd                  ObjectType = "rng-egd"
	ObjectTypeRngRandom               ObjectType = "rng-random"
	ObjectTypeSecret                  ObjectType = "secret"
	ObjectTypeSecretKeyring           ObjectType = "secret_keyring"
	ObjectTypeSevGuest                ObjectType = "sev-guest"
	ObjectTypeThreadContext           ObjectType = "thread-context"
	ObjectTypeS390PvGuest             ObjectType = "s390-pv-guest"
	ObjectTypeThrottleGroup           ObjectType = "throttle-group"
	ObjectTypeTlsCredsAnon            ObjectType = "tls-creds-anon"
	ObjectTypeTlsCredsPsk             ObjectType = "tls-creds-psk"
	ObjectTypeTlsCredsX509            ObjectType = "tls-creds-x509"
	ObjectTypeTlsCipherSuites         ObjectType = "tls-cipher-suites"
	ObjectTypeRemoteObject            ObjectType = "x-remote-object"
	ObjectTypeVfioUserServer          ObjectType = "x-vfio-user-server"
)

// ObjectOptions
//
// Describes the options of a user creatable QOM object.
type ObjectOptions struct {
	// Discriminator: qom-type

	// QomType the class name for the object to be created
	QomType ObjectType `json:"qom-type"`
	// Id the name of the new object
	Id string `json:"id"`

	AuthzList               *AuthZListProperties          `json:"-"`
	AuthzListfile           *AuthZListFileProperties      `json:"-"`
	AuthzPam                *AuthZPAMProperties           `json:"-"`
	AuthzSimple             *AuthZSimpleProperties        `json:"-"`
	CanHostSocketcan        *CanHostSocketcanProperties   `json:"-"`
	ColoCompare             *ColoCompareProperties        `json:"-"`
	CryptodevBackend        *CryptodevBackendProperties   `json:"-"`
	CryptodevBackendBuiltin *CryptodevBackendProperties   `json:"-"`
	CryptodevBackendLkcf    *CryptodevBackendProperties   `json:"-"`
	CryptodevVhostUser      *CryptodevVhostUserProperties `json:"-"`
	DbusVmstate             *DBusVMStateProperties        `json:"-"`
	FilterBuffer            *FilterBufferProperties       `json:"-"`
	FilterDump              *FilterDumpProperties         `json:"-"`
	FilterMirror            *FilterMirrorProperties       `json:"-"`
	FilterRedirector        *FilterRedirectorProperties   `json:"-"`
	FilterReplay            *NetfilterProperties          `json:"-"`
	FilterRewriter          *FilterRewriterProperties     `json:"-"`
	InputBarrier            *InputBarrierProperties       `json:"-"`
	InputLinux              *InputLinuxProperties         `json:"-"`
	Iommufd                 *IOMMUFDProperties            `json:"-"`
	Iothread                *IothreadProperties           `json:"-"`
	MainLoop                *MainLoopProperties           `json:"-"`
	MemoryBackendEpc        *MemoryBackendEpcProperties   `json:"-"`
	MemoryBackendFile       *MemoryBackendFileProperties  `json:"-"`
	MemoryBackendMemfd      *MemoryBackendMemfdProperties `json:"-"`
	MemoryBackendRam        *MemoryBackendProperties      `json:"-"`
	PrManagerHelper         *PrManagerHelperProperties    `json:"-"`
	Qtest                   *QtestProperties              `json:"-"`
	RngBuiltin              *RngProperties                `json:"-"`
	RngEgd                  *RngEgdProperties             `json:"-"`
	RngRandom               *RngRandomProperties          `json:"-"`
	Secret                  *SecretProperties             `json:"-"`
	SecretKeyring           *SecretKeyringProperties      `json:"-"`
	SevGuest                *SevGuestProperties           `json:"-"`
	ThreadContext           *ThreadContextProperties      `json:"-"`
	ThrottleGroup           *ThrottleGroupProperties      `json:"-"`
	TlsCredsAnon            *TlsCredsAnonProperties       `json:"-"`
	TlsCredsPsk             *TlsCredsPskProperties        `json:"-"`
	TlsCredsX509            *TlsCredsX509Properties       `json:"-"`
	TlsCipherSuites         *TlsCredsProperties           `json:"-"`
	RemoteObject            *RemoteObjectProperties       `json:"-"`
	VfioUserServer          *VfioUserServerProperties     `json:"-"`
}

func (u ObjectOptions) MarshalJSON() ([]byte, error) {
	switch u.QomType {
	case "authz-list":
		if u.AuthzList == nil {
			return nil, fmt.Errorf("expected AuthzList to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*AuthZListProperties
		}{
			QomType:             u.QomType,
			Id:                  u.Id,
			AuthZListProperties: u.AuthzList,
		})
	case "authz-listfile":
		if u.AuthzListfile == nil {
			return nil, fmt.Errorf("expected AuthzListfile to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*AuthZListFileProperties
		}{
			QomType:                 u.QomType,
			Id:                      u.Id,
			AuthZListFileProperties: u.AuthzListfile,
		})
	case "authz-pam":
		if u.AuthzPam == nil {
			return nil, fmt.Errorf("expected AuthzPam to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*AuthZPAMProperties
		}{
			QomType:            u.QomType,
			Id:                 u.Id,
			AuthZPAMProperties: u.AuthzPam,
		})
	case "authz-simple":
		if u.AuthzSimple == nil {
			return nil, fmt.Errorf("expected AuthzSimple to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*AuthZSimpleProperties
		}{
			QomType:               u.QomType,
			Id:                    u.Id,
			AuthZSimpleProperties: u.AuthzSimple,
		})
	case "can-host-socketcan":
		if u.CanHostSocketcan == nil {
			return nil, fmt.Errorf("expected CanHostSocketcan to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*CanHostSocketcanProperties
		}{
			QomType:                    u.QomType,
			Id:                         u.Id,
			CanHostSocketcanProperties: u.CanHostSocketcan,
		})
	case "colo-compare":
		if u.ColoCompare == nil {
			return nil, fmt.Errorf("expected ColoCompare to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*ColoCompareProperties
		}{
			QomType:               u.QomType,
			Id:                    u.Id,
			ColoCompareProperties: u.ColoCompare,
		})
	case "cryptodev-backend":
		if u.CryptodevBackend == nil {
			return nil, fmt.Errorf("expected CryptodevBackend to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*CryptodevBackendProperties
		}{
			QomType:                    u.QomType,
			Id:                         u.Id,
			CryptodevBackendProperties: u.CryptodevBackend,
		})
	case "cryptodev-backend-builtin":
		if u.CryptodevBackendBuiltin == nil {
			return nil, fmt.Errorf("expected CryptodevBackendBuiltin to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*CryptodevBackendProperties
		}{
			QomType:                    u.QomType,
			Id:                         u.Id,
			CryptodevBackendProperties: u.CryptodevBackendBuiltin,
		})
	case "cryptodev-backend-lkcf":
		if u.CryptodevBackendLkcf == nil {
			return nil, fmt.Errorf("expected CryptodevBackendLkcf to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*CryptodevBackendProperties
		}{
			QomType:                    u.QomType,
			Id:                         u.Id,
			CryptodevBackendProperties: u.CryptodevBackendLkcf,
		})
	case "cryptodev-vhost-user":
		if u.CryptodevVhostUser == nil {
			return nil, fmt.Errorf("expected CryptodevVhostUser to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*CryptodevVhostUserProperties
		}{
			QomType:                      u.QomType,
			Id:                           u.Id,
			CryptodevVhostUserProperties: u.CryptodevVhostUser,
		})
	case "dbus-vmstate":
		if u.DbusVmstate == nil {
			return nil, fmt.Errorf("expected DbusVmstate to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*DBusVMStateProperties
		}{
			QomType:               u.QomType,
			Id:                    u.Id,
			DBusVMStateProperties: u.DbusVmstate,
		})
	case "filter-buffer":
		if u.FilterBuffer == nil {
			return nil, fmt.Errorf("expected FilterBuffer to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*FilterBufferProperties
		}{
			QomType:                u.QomType,
			Id:                     u.Id,
			FilterBufferProperties: u.FilterBuffer,
		})
	case "filter-dump":
		if u.FilterDump == nil {
			return nil, fmt.Errorf("expected FilterDump to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*FilterDumpProperties
		}{
			QomType:              u.QomType,
			Id:                   u.Id,
			FilterDumpProperties: u.FilterDump,
		})
	case "filter-mirror":
		if u.FilterMirror == nil {
			return nil, fmt.Errorf("expected FilterMirror to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*FilterMirrorProperties
		}{
			QomType:                u.QomType,
			Id:                     u.Id,
			FilterMirrorProperties: u.FilterMirror,
		})
	case "filter-redirector":
		if u.FilterRedirector == nil {
			return nil, fmt.Errorf("expected FilterRedirector to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*FilterRedirectorProperties
		}{
			QomType:                    u.QomType,
			Id:                         u.Id,
			FilterRedirectorProperties: u.FilterRedirector,
		})
	case "filter-replay":
		if u.FilterReplay == nil {
			return nil, fmt.Errorf("expected FilterReplay to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*NetfilterProperties
		}{
			QomType:             u.QomType,
			Id:                  u.Id,
			NetfilterProperties: u.FilterReplay,
		})
	case "filter-rewriter":
		if u.FilterRewriter == nil {
			return nil, fmt.Errorf("expected FilterRewriter to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*FilterRewriterProperties
		}{
			QomType:                  u.QomType,
			Id:                       u.Id,
			FilterRewriterProperties: u.FilterRewriter,
		})
	case "input-barrier":
		if u.InputBarrier == nil {
			return nil, fmt.Errorf("expected InputBarrier to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*InputBarrierProperties
		}{
			QomType:                u.QomType,
			Id:                     u.Id,
			InputBarrierProperties: u.InputBarrier,
		})
	case "input-linux":
		if u.InputLinux == nil {
			return nil, fmt.Errorf("expected InputLinux to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*InputLinuxProperties
		}{
			QomType:              u.QomType,
			Id:                   u.Id,
			InputLinuxProperties: u.InputLinux,
		})
	case "iommufd":
		if u.Iommufd == nil {
			return nil, fmt.Errorf("expected Iommufd to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*IOMMUFDProperties
		}{
			QomType:           u.QomType,
			Id:                u.Id,
			IOMMUFDProperties: u.Iommufd,
		})
	case "iothread":
		if u.Iothread == nil {
			return nil, fmt.Errorf("expected Iothread to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*IothreadProperties
		}{
			QomType:            u.QomType,
			Id:                 u.Id,
			IothreadProperties: u.Iothread,
		})
	case "main-loop":
		if u.MainLoop == nil {
			return nil, fmt.Errorf("expected MainLoop to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*MainLoopProperties
		}{
			QomType:            u.QomType,
			Id:                 u.Id,
			MainLoopProperties: u.MainLoop,
		})
	case "memory-backend-epc":
		if u.MemoryBackendEpc == nil {
			return nil, fmt.Errorf("expected MemoryBackendEpc to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*MemoryBackendEpcProperties
		}{
			QomType:                    u.QomType,
			Id:                         u.Id,
			MemoryBackendEpcProperties: u.MemoryBackendEpc,
		})
	case "memory-backend-file":
		if u.MemoryBackendFile == nil {
			return nil, fmt.Errorf("expected MemoryBackendFile to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*MemoryBackendFileProperties
		}{
			QomType:                     u.QomType,
			Id:                          u.Id,
			MemoryBackendFileProperties: u.MemoryBackendFile,
		})
	case "memory-backend-memfd":
		if u.MemoryBackendMemfd == nil {
			return nil, fmt.Errorf("expected MemoryBackendMemfd to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*MemoryBackendMemfdProperties
		}{
			QomType:                      u.QomType,
			Id:                           u.Id,
			MemoryBackendMemfdProperties: u.MemoryBackendMemfd,
		})
	case "memory-backend-ram":
		if u.MemoryBackendRam == nil {
			return nil, fmt.Errorf("expected MemoryBackendRam to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*MemoryBackendProperties
		}{
			QomType:                 u.QomType,
			Id:                      u.Id,
			MemoryBackendProperties: u.MemoryBackendRam,
		})
	case "pr-manager-helper":
		if u.PrManagerHelper == nil {
			return nil, fmt.Errorf("expected PrManagerHelper to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*PrManagerHelperProperties
		}{
			QomType:                   u.QomType,
			Id:                        u.Id,
			PrManagerHelperProperties: u.PrManagerHelper,
		})
	case "qtest":
		if u.Qtest == nil {
			return nil, fmt.Errorf("expected Qtest to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*QtestProperties
		}{
			QomType:         u.QomType,
			Id:              u.Id,
			QtestProperties: u.Qtest,
		})
	case "rng-builtin":
		if u.RngBuiltin == nil {
			return nil, fmt.Errorf("expected RngBuiltin to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*RngProperties
		}{
			QomType:       u.QomType,
			Id:            u.Id,
			RngProperties: u.RngBuiltin,
		})
	case "rng-egd":
		if u.RngEgd == nil {
			return nil, fmt.Errorf("expected RngEgd to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*RngEgdProperties
		}{
			QomType:          u.QomType,
			Id:               u.Id,
			RngEgdProperties: u.RngEgd,
		})
	case "rng-random":
		if u.RngRandom == nil {
			return nil, fmt.Errorf("expected RngRandom to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*RngRandomProperties
		}{
			QomType:             u.QomType,
			Id:                  u.Id,
			RngRandomProperties: u.RngRandom,
		})
	case "secret":
		if u.Secret == nil {
			return nil, fmt.Errorf("expected Secret to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*SecretProperties
		}{
			QomType:          u.QomType,
			Id:               u.Id,
			SecretProperties: u.Secret,
		})
	case "secret_keyring":
		if u.SecretKeyring == nil {
			return nil, fmt.Errorf("expected SecretKeyring to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*SecretKeyringProperties
		}{
			QomType:                 u.QomType,
			Id:                      u.Id,
			SecretKeyringProperties: u.SecretKeyring,
		})
	case "sev-guest":
		if u.SevGuest == nil {
			return nil, fmt.Errorf("expected SevGuest to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*SevGuestProperties
		}{
			QomType:            u.QomType,
			Id:                 u.Id,
			SevGuestProperties: u.SevGuest,
		})
	case "thread-context":
		if u.ThreadContext == nil {
			return nil, fmt.Errorf("expected ThreadContext to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*ThreadContextProperties
		}{
			QomType:                 u.QomType,
			Id:                      u.Id,
			ThreadContextProperties: u.ThreadContext,
		})
	case "throttle-group":
		if u.ThrottleGroup == nil {
			return nil, fmt.Errorf("expected ThrottleGroup to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*ThrottleGroupProperties
		}{
			QomType:                 u.QomType,
			Id:                      u.Id,
			ThrottleGroupProperties: u.ThrottleGroup,
		})
	case "tls-creds-anon":
		if u.TlsCredsAnon == nil {
			return nil, fmt.Errorf("expected TlsCredsAnon to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*TlsCredsAnonProperties
		}{
			QomType:                u.QomType,
			Id:                     u.Id,
			TlsCredsAnonProperties: u.TlsCredsAnon,
		})
	case "tls-creds-psk":
		if u.TlsCredsPsk == nil {
			return nil, fmt.Errorf("expected TlsCredsPsk to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*TlsCredsPskProperties
		}{
			QomType:               u.QomType,
			Id:                    u.Id,
			TlsCredsPskProperties: u.TlsCredsPsk,
		})
	case "tls-creds-x509":
		if u.TlsCredsX509 == nil {
			return nil, fmt.Errorf("expected TlsCredsX509 to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*TlsCredsX509Properties
		}{
			QomType:                u.QomType,
			Id:                     u.Id,
			TlsCredsX509Properties: u.TlsCredsX509,
		})
	case "tls-cipher-suites":
		if u.TlsCipherSuites == nil {
			return nil, fmt.Errorf("expected TlsCipherSuites to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*TlsCredsProperties
		}{
			QomType:            u.QomType,
			Id:                 u.Id,
			TlsCredsProperties: u.TlsCipherSuites,
		})
	case "x-remote-object":
		if u.RemoteObject == nil {
			return nil, fmt.Errorf("expected RemoteObject to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*RemoteObjectProperties
		}{
			QomType:                u.QomType,
			Id:                     u.Id,
			RemoteObjectProperties: u.RemoteObject,
		})
	case "x-vfio-user-server":
		if u.VfioUserServer == nil {
			return nil, fmt.Errorf("expected VfioUserServer to be set")
		}

		return json.Marshal(struct {
			QomType ObjectType `json:"qom-type"`
			Id      string     `json:"id"`
			*VfioUserServerProperties
		}{
			QomType:                  u.QomType,
			Id:                       u.Id,
			VfioUserServerProperties: u.VfioUserServer,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// ObjectAdd
//
// Create a QOM object.
type ObjectAdd struct {
	ObjectOptions
}

func (ObjectAdd) Command() string {
	return "object-add"
}

func (cmd ObjectAdd) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "object-add", cmd, nil)
}

// ObjectDel
//
// Remove a QOM object.
type ObjectDel struct {
	// Id the name of the QOM object to remove
	Id string `json:"id"`
}

func (ObjectDel) Command() string {
	return "object-del"
}

func (cmd ObjectDel) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "object-del", cmd, nil)
}
