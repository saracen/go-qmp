package qemu

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
	client.RegisterEvent("SHUTDOWN", func() api.EventType { return &ShutdownEvent{} })
	client.RegisterEvent("POWERDOWN", func() api.EventType { return &PowerdownEvent{} })
	client.RegisterEvent("RESET", func() api.EventType { return &ResetEvent{} })
	client.RegisterEvent("STOP", func() api.EventType { return &StopEvent{} })
	client.RegisterEvent("RESUME", func() api.EventType { return &ResumeEvent{} })
	client.RegisterEvent("SUSPEND", func() api.EventType { return &SuspendEvent{} })
	client.RegisterEvent("SUSPEND_DISK", func() api.EventType { return &SuspendDiskEvent{} })
	client.RegisterEvent("WAKEUP", func() api.EventType { return &WakeupEvent{} })
	client.RegisterEvent("WATCHDOG", func() api.EventType { return &WatchdogEvent{} })
	client.RegisterEvent("GUEST_PANICKED", func() api.EventType { return &GuestPanickedEvent{} })
	client.RegisterEvent("GUEST_CRASHLOADED", func() api.EventType { return &GuestCrashloadedEvent{} })
	client.RegisterEvent("MEMORY_FAILURE", func() api.EventType { return &MemoryFailureEvent{} })
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
	client.RegisterEvent("DEVICE_TRAY_MOVED", func() api.EventType { return &DeviceTrayMovedEvent{} })
	client.RegisterEvent("PR_MANAGER_STATUS_CHANGED", func() api.EventType { return &PrManagerStatusChangedEvent{} })
	client.RegisterEvent("BLOCK_EXPORT_DELETED", func() api.EventType { return &BlockExportDeletedEvent{} })
	client.RegisterEvent("VSERPORT_CHANGE", func() api.EventType { return &VserportChangeEvent{} })
	client.RegisterEvent("DUMP_COMPLETED", func() api.EventType { return &DumpCompletedEvent{} })
	client.RegisterEvent("NIC_RX_FILTER_CHANGED", func() api.EventType { return &NicRxFilterChangedEvent{} })
	client.RegisterEvent("FAILOVER_NEGOTIATED", func() api.EventType { return &FailoverNegotiatedEvent{} })
	client.RegisterEvent("NETDEV_STREAM_CONNECTED", func() api.EventType { return &NetdevStreamConnectedEvent{} })
	client.RegisterEvent("NETDEV_STREAM_DISCONNECTED", func() api.EventType { return &NetdevStreamDisconnectedEvent{} })
	client.RegisterEvent("RDMA_GID_STATUS_CHANGED", func() api.EventType { return &RdmaGidStatusChangedEvent{} })
	client.RegisterEvent("SPICE_CONNECTED", func() api.EventType { return &SpiceConnectedEvent{} })
	client.RegisterEvent("SPICE_INITIALIZED", func() api.EventType { return &SpiceInitializedEvent{} })
	client.RegisterEvent("SPICE_DISCONNECTED", func() api.EventType { return &SpiceDisconnectedEvent{} })
	client.RegisterEvent("SPICE_MIGRATE_COMPLETED", func() api.EventType { return &SpiceMigrateCompletedEvent{} })
	client.RegisterEvent("VNC_CONNECTED", func() api.EventType { return &VncConnectedEvent{} })
	client.RegisterEvent("VNC_INITIALIZED", func() api.EventType { return &VncInitializedEvent{} })
	client.RegisterEvent("VNC_DISCONNECTED", func() api.EventType { return &VncDisconnectedEvent{} })
	client.RegisterEvent("MIGRATION", func() api.EventType { return &MigrationEvent{} })
	client.RegisterEvent("MIGRATION_PASS", func() api.EventType { return &MigrationPassEvent{} })
	client.RegisterEvent("COLO_EXIT", func() api.EventType { return &ColoExitEvent{} })
	client.RegisterEvent("UNPLUG_PRIMARY", func() api.EventType { return &UnplugPrimaryEvent{} })
	client.RegisterEvent("DEVICE_DELETED", func() api.EventType { return &DeviceDeletedEvent{} })
	client.RegisterEvent("DEVICE_UNPLUG_GUEST_ERROR", func() api.EventType { return &DeviceUnplugGuestErrorEvent{} })
	client.RegisterEvent("BALLOON_CHANGE", func() api.EventType { return &BalloonChangeEvent{} })
	client.RegisterEvent("HV_BALLOON_STATUS_REPORT", func() api.EventType { return &HvBalloonStatusReportEvent{} })
	client.RegisterEvent("MEMORY_DEVICE_SIZE_CHANGE", func() api.EventType { return &MemoryDeviceSizeChangeEvent{} })
	client.RegisterEvent("MEM_UNPLUG_ERROR", func() api.EventType { return &MemUnplugErrorEvent{} })
	client.RegisterEvent("CPU_POLARIZATION_CHANGE", func() api.EventType { return &CpuPolarizationChangeEvent{} })
	client.RegisterEvent("RTC_CHANGE", func() api.EventType { return &RtcChangeEvent{} })
	client.RegisterEvent("VFU_CLIENT_HANGUP", func() api.EventType { return &VfuClientHangupEvent{} })
	client.RegisterEvent("ACPI_DEVICE_OST", func() api.EventType { return &AcpiDeviceOstEvent{} })
}

// QapiErrorClass QEMU error classes
type QapiErrorClass string

const (
	// QapiErrorClassGenericerror this is used for errors that don't require a specific error class. This should be the default case for most errors
	QapiErrorClassGenericerror QapiErrorClass = "GenericError"
	// QapiErrorClassCommandnotfound the requested command has not been found
	QapiErrorClassCommandnotfound QapiErrorClass = "CommandNotFound"
	// QapiErrorClassDevicenotactive a device has failed to be become active
	QapiErrorClassDevicenotactive QapiErrorClass = "DeviceNotActive"
	// QapiErrorClassDevicenotfound the requested device has not been found
	QapiErrorClassDevicenotfound QapiErrorClass = "DeviceNotFound"
	// QapiErrorClassKvmmissingcap the requested operation can't be fulfilled because a required KVM capability is missing
	QapiErrorClassKvmmissingcap QapiErrorClass = "KVMMissingCap"
)

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

// RunState An enumeration of VM run states.
type RunState string

const (
	// RunStateDebug QEMU is running on a debugger
	RunStateDebug RunState = "debug"
	// RunStateInmigrate guest is paused waiting for an incoming migration. Note that this state does not tell whether the machine will start at the end of the migration. This depends on the command-line -S option and any invocation of 'stop' or 'cont' that has happened since QEMU was started.
	RunStateInmigrate RunState = "inmigrate"
	// RunStateInternalError An internal error that prevents further guest execution has occurred
	RunStateInternalError RunState = "internal-error"
	// RunStateIoError the last IOP has failed and the device is configured to pause on I/O errors
	RunStateIoError RunState = "io-error"
	// RunStatePaused guest has been paused via the 'stop' command
	RunStatePaused RunState = "paused"
	// RunStatePostmigrate guest is paused following a successful 'migrate'
	RunStatePostmigrate RunState = "postmigrate"
	// RunStatePrelaunch QEMU was started with -S and guest has not started
	RunStatePrelaunch RunState = "prelaunch"
	// RunStateFinishMigrate guest is paused to finish the migration process
	RunStateFinishMigrate RunState = "finish-migrate"
	// RunStateRestoreVm guest is paused to restore VM state
	RunStateRestoreVm RunState = "restore-vm"
	// RunStateRunning guest is actively running
	RunStateRunning RunState = "running"
	// RunStateSaveVm guest is paused to save the VM state
	RunStateSaveVm RunState = "save-vm"
	// RunStateShutdown guest is shut down (and -no-shutdown is in use)
	RunStateShutdown RunState = "shutdown"
	// RunStateSuspended guest is suspended (ACPI S3)
	RunStateSuspended RunState = "suspended"
	// RunStateWatchdog the watchdog action is configured to pause and has been triggered
	RunStateWatchdog RunState = "watchdog"
	// RunStateGuestPanicked guest has been panicked as a result of guest OS panic
	RunStateGuestPanicked RunState = "guest-panicked"
	// RunStateColo guest is paused to save/restore VM state under colo checkpoint, VM can not get into this state unless colo capability is enabled for migration. (since 2.8)
	RunStateColo RunState = "colo"
)

// ShutdownCause An enumeration of reasons for a Shutdown.
type ShutdownCause string

const (
	// ShutdownCauseNone No shutdown request pending
	ShutdownCauseNone ShutdownCause = "none"
	// ShutdownCauseHostError An error prevents further use of guest
	ShutdownCauseHostError ShutdownCause = "host-error"
	// ShutdownCauseHostQmpQuit Reaction to the QMP command 'quit'
	ShutdownCauseHostQmpQuit ShutdownCause = "host-qmp-quit"
	// ShutdownCauseHostQmpSystemReset Reaction to the QMP command 'system_reset'
	ShutdownCauseHostQmpSystemReset ShutdownCause = "host-qmp-system-reset"
	// ShutdownCauseHostSignal Reaction to a signal, such as SIGINT
	ShutdownCauseHostSignal ShutdownCause = "host-signal"
	// ShutdownCauseHostUi Reaction to a UI event, like window close
	ShutdownCauseHostUi ShutdownCause = "host-ui"
	// ShutdownCauseGuestShutdown Guest shutdown/suspend request, via ACPI or other hardware-specific means
	ShutdownCauseGuestShutdown ShutdownCause = "guest-shutdown"
	// ShutdownCauseGuestReset Guest reset request, and command line turns that into a shutdown
	ShutdownCauseGuestReset ShutdownCause = "guest-reset"
	// ShutdownCauseGuestPanic Guest panicked, and command line turns that into a shutdown
	ShutdownCauseGuestPanic ShutdownCause = "guest-panic"
	// ShutdownCauseSubsystemReset Partial guest reset that does not trigger QMP events and ignores --no-reboot. This is useful for sanitizing hypercalls on s390 that are used during kexec/kdump/boot
	ShutdownCauseSubsystemReset ShutdownCause = "subsystem-reset"
	// ShutdownCauseSnapshotLoad A snapshot is being loaded by the record & replay subsystem. This value is used only within QEMU. It doesn't occur in QMP. (since 7.2)
	ShutdownCauseSnapshotLoad ShutdownCause = "snapshot-load"
)

// StatusInfo
//
// Information about VM run state
type StatusInfo struct {
	// Running true if all VCPUs are runnable, false if not runnable
	Running bool `json:"running"`
	// Status the virtual machine @RunState
	Status RunState `json:"status"`
}

// QueryStatus
//
// Query the run status of the VM
type QueryStatus struct {
}

func (QueryStatus) Command() string {
	return "query-status"
}

func (cmd QueryStatus) Execute(ctx context.Context, client api.Client) (StatusInfo, error) {
	var ret StatusInfo

	return ret, client.Execute(ctx, "query-status", cmd, &ret)
}

// ShutdownEvent (SHUTDOWN)
//
// Emitted when the virtual machine has shut down, indicating that qemu is about to exit.
type ShutdownEvent struct {
	// Guest If true, the shutdown was triggered by a guest request (such as a guest-initiated ACPI shutdown request or other hardware-specific action) rather than a host request (such as sending qemu a SIGINT). (since 2.10)
	Guest bool `json:"guest"`
	// Reason The @ShutdownCause which resulted in the SHUTDOWN. (since 4.0)
	Reason ShutdownCause `json:"reason"`
}

func (ShutdownEvent) Event() string {
	return "SHUTDOWN"
}

// PowerdownEvent (POWERDOWN)
//
// Emitted when the virtual machine is powered down through the power control system, such as via ACPI.
type PowerdownEvent struct {
}

func (PowerdownEvent) Event() string {
	return "POWERDOWN"
}

// ResetEvent (RESET)
//
// Emitted when the virtual machine is reset
type ResetEvent struct {
	// Guest If true, the reset was triggered by a guest request (such as a guest-initiated ACPI reboot request or other hardware-specific action) rather than a host request (such as the QMP command system_reset). (since 2.10)
	Guest bool `json:"guest"`
	// Reason The @ShutdownCause of the RESET. (since 4.0)
	Reason ShutdownCause `json:"reason"`
}

func (ResetEvent) Event() string {
	return "RESET"
}

// StopEvent (STOP)
//
// Emitted when the virtual machine is stopped
type StopEvent struct {
}

func (StopEvent) Event() string {
	return "STOP"
}

// ResumeEvent (RESUME)
//
// Emitted when the virtual machine resumes execution
type ResumeEvent struct {
}

func (ResumeEvent) Event() string {
	return "RESUME"
}

// SuspendEvent (SUSPEND)
//
// Emitted when guest enters a hardware suspension state, for example, S3 state, which is sometimes called standby state
type SuspendEvent struct {
}

func (SuspendEvent) Event() string {
	return "SUSPEND"
}

// SuspendDiskEvent (SUSPEND_DISK)
//
// Emitted when guest enters a hardware suspension state with data saved on disk, for example, S4 state, which is sometimes called hibernate state
type SuspendDiskEvent struct {
}

func (SuspendDiskEvent) Event() string {
	return "SUSPEND_DISK"
}

// WakeupEvent (WAKEUP)
//
// Emitted when the guest has woken up from suspend state and is running
type WakeupEvent struct {
}

func (WakeupEvent) Event() string {
	return "WAKEUP"
}

// WatchdogEvent (WATCHDOG)
//
// Emitted when the watchdog device's timer is expired
type WatchdogEvent struct {
	// Action action that has been taken
	Action WatchdogAction `json:"action"`
}

func (WatchdogEvent) Event() string {
	return "WATCHDOG"
}

// WatchdogAction An enumeration of the actions taken when the watchdog device's timer is expired
type WatchdogAction string

const (
	// WatchdogActionReset system resets
	WatchdogActionReset WatchdogAction = "reset"
	// WatchdogActionShutdown system shutdown, note that it is similar to @powerdown, which tries to set to system status and notify guest
	WatchdogActionShutdown WatchdogAction = "shutdown"
	// WatchdogActionPoweroff system poweroff, the emulator program exits
	WatchdogActionPoweroff WatchdogAction = "poweroff"
	// WatchdogActionPause system pauses, similar to @stop
	WatchdogActionPause WatchdogAction = "pause"
	// WatchdogActionDebug system enters debug state
	WatchdogActionDebug WatchdogAction = "debug"
	// WatchdogActionNone nothing is done
	WatchdogActionNone WatchdogAction = "none"
	// WatchdogActionInjectNmi a non-maskable interrupt is injected into the first VCPU (all VCPUS on x86) (since 2.4)
	WatchdogActionInjectNmi WatchdogAction = "inject-nmi"
)

// RebootAction Possible QEMU actions upon guest reboot
type RebootAction string

const (
	// RebootActionReset Reset the VM
	RebootActionReset RebootAction = "reset"
	// RebootActionShutdown Shutdown the VM and exit, according to the shutdown action
	RebootActionShutdown RebootAction = "shutdown"
)

// ShutdownAction Possible QEMU actions upon guest shutdown
type ShutdownAction string

const (
	// ShutdownActionPoweroff Shutdown the VM and exit
	ShutdownActionPoweroff ShutdownAction = "poweroff"
	// ShutdownActionPause pause the VM
	ShutdownActionPause ShutdownAction = "pause"
)

// PanicAction
type PanicAction string

const (
	// PanicActionPause Pause the VM
	PanicActionPause PanicAction = "pause"
	// PanicActionShutdown Shutdown the VM and exit, according to the shutdown action
	PanicActionShutdown PanicAction = "shutdown"
	// PanicActionExitFailure Shutdown the VM and exit with nonzero status (since 7.1)
	PanicActionExitFailure PanicAction = "exit-failure"
	// PanicActionNone Continue VM execution
	PanicActionNone PanicAction = "none"
)

// WatchdogSetAction
//
// Set watchdog action
type WatchdogSetAction struct {
	Action WatchdogAction `json:"action"`
}

func (WatchdogSetAction) Command() string {
	return "watchdog-set-action"
}

func (cmd WatchdogSetAction) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "watchdog-set-action", cmd, nil)
}

// SetAction
//
// Set the actions that will be taken by the emulator in response to guest events.
type SetAction struct {
	// Reboot @RebootAction action taken on guest reboot.
	Reboot *RebootAction `json:"reboot,omitempty"`
	// Shutdown @ShutdownAction action taken on guest shutdown.
	Shutdown *ShutdownAction `json:"shutdown,omitempty"`
	// Panic @PanicAction action taken on guest panic.
	Panic *PanicAction `json:"panic,omitempty"`
	// Watchdog @WatchdogAction action taken when watchdog timer expires .
	Watchdog *WatchdogAction `json:"watchdog,omitempty"`
}

func (SetAction) Command() string {
	return "set-action"
}

func (cmd SetAction) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "set-action", cmd, nil)
}

// GuestPanickedEvent (GUEST_PANICKED)
//
// Emitted when guest OS panic is detected
type GuestPanickedEvent struct {
	// Action action that has been taken, currently always "pause"
	Action GuestPanicAction `json:"action"`
	// Info information about a panic (since 2.9)
	Info *GuestPanicInformation `json:"info,omitempty"`
}

func (GuestPanickedEvent) Event() string {
	return "GUEST_PANICKED"
}

// GuestCrashloadedEvent (GUEST_CRASHLOADED)
//
// Emitted when guest OS crash loaded is detected
type GuestCrashloadedEvent struct {
	// Action action that has been taken, currently always "run"
	Action GuestPanicAction `json:"action"`
	// Info information about a panic
	Info *GuestPanicInformation `json:"info,omitempty"`
}

func (GuestCrashloadedEvent) Event() string {
	return "GUEST_CRASHLOADED"
}

// GuestPanicAction An enumeration of the actions taken when guest OS panic is detected
type GuestPanicAction string

const (
	// GuestPanicActionPause system pauses
	GuestPanicActionPause GuestPanicAction = "pause"
	// GuestPanicActionPoweroff system powers off (since 2.8)
	GuestPanicActionPoweroff GuestPanicAction = "poweroff"
	// GuestPanicActionRun system continues to run (since 5.0)
	GuestPanicActionRun GuestPanicAction = "run"
)

// GuestPanicInformationType An enumeration of the guest panic information types
type GuestPanicInformationType string

const (
	// GuestPanicInformationTypeHyperV hyper-v guest panic information type
	GuestPanicInformationTypeHyperV GuestPanicInformationType = "hyper-v"
	// GuestPanicInformationTypeS390 s390 guest panic information type (Since: 2.12)
	GuestPanicInformationTypeS390 GuestPanicInformationType = "s390"
)

// GuestPanicInformation
//
// Information about a guest panic
type GuestPanicInformation struct {
	// Discriminator: type

	// Type Crash type that defines the hypervisor specific information
	Type GuestPanicInformationType `json:"type"`

	HyperV *GuestPanicInformationHyperV `json:"-"`
	S390   *GuestPanicInformationS390   `json:"-"`
}

func (u GuestPanicInformation) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "hyper-v":
		if u.HyperV == nil {
			return nil, fmt.Errorf("expected HyperV to be set")
		}

		return json.Marshal(struct {
			Type GuestPanicInformationType `json:"type"`
			*GuestPanicInformationHyperV
		}{
			Type:                        u.Type,
			GuestPanicInformationHyperV: u.HyperV,
		})
	case "s390":
		if u.S390 == nil {
			return nil, fmt.Errorf("expected S390 to be set")
		}

		return json.Marshal(struct {
			Type GuestPanicInformationType `json:"type"`
			*GuestPanicInformationS390
		}{
			Type:                      u.Type,
			GuestPanicInformationS390: u.S390,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// GuestPanicInformationHyperV
//
// Hyper-V specific guest panic information (HV crash MSRs)
type GuestPanicInformationHyperV struct {
	Arg1 uint64 `json:"arg1"`
	Arg2 uint64 `json:"arg2"`
	Arg3 uint64 `json:"arg3"`
	Arg4 uint64 `json:"arg4"`
	Arg5 uint64 `json:"arg5"`
}

// S390CrashReason Reason why the CPU is in a crashed state.
type S390CrashReason string

const (
	// S390CrashReasonUnknown no crash reason was set
	S390CrashReasonUnknown S390CrashReason = "unknown"
	// S390CrashReasonDisabledWait the CPU has entered a disabled wait state
	S390CrashReasonDisabledWait S390CrashReason = "disabled-wait"
	// S390CrashReasonExtintLoop clock comparator or cpu timer interrupt with new PSW enabled for external interrupts
	S390CrashReasonExtintLoop S390CrashReason = "extint-loop"
	// S390CrashReasonPgmintLoop program interrupt with BAD new PSW
	S390CrashReasonPgmintLoop S390CrashReason = "pgmint-loop"
	// S390CrashReasonOpintLoop operation exception interrupt with invalid code at the program interrupt new PSW
	S390CrashReasonOpintLoop S390CrashReason = "opint-loop"
)

// GuestPanicInformationS390
//
// S390 specific guest panic information (PSW)
type GuestPanicInformationS390 struct {
	// Core core id of the CPU that crashed
	Core uint32 `json:"core"`
	// PswMask control fields of guest PSW
	PswMask uint64 `json:"psw-mask"`
	// PswAddr guest instruction address
	PswAddr uint64 `json:"psw-addr"`
	// Reason guest crash reason
	Reason S390CrashReason `json:"reason"`
}

// MemoryFailureEvent (MEMORY_FAILURE)
//
// Emitted when a memory failure occurs on host side.
type MemoryFailureEvent struct {
	// Recipient recipient is defined as @MemoryFailureRecipient.
	Recipient MemoryFailureRecipient `json:"recipient"`
	// Action action that has been taken. action is defined as @MemoryFailureAction.
	Action MemoryFailureAction `json:"action"`
	// Flags flags for MemoryFailureAction. action is defined as @MemoryFailureFlags.
	Flags MemoryFailureFlags `json:"flags"`
}

func (MemoryFailureEvent) Event() string {
	return "MEMORY_FAILURE"
}

// MemoryFailureRecipient Hardware memory failure occurs, handled by recipient.
type MemoryFailureRecipient string

const (
	// MemoryFailureRecipientHypervisor memory failure at QEMU process address space. (none guest memory, but used by QEMU itself).
	MemoryFailureRecipientHypervisor MemoryFailureRecipient = "hypervisor"
	// MemoryFailureRecipientGuest memory failure at guest memory,
	MemoryFailureRecipientGuest MemoryFailureRecipient = "guest"
)

// MemoryFailureAction Actions taken by QEMU in response to a hardware memory failure.
type MemoryFailureAction string

const (
	// MemoryFailureActionIgnore the memory failure could be ignored. This will only be the case for action-optional failures.
	MemoryFailureActionIgnore MemoryFailureAction = "ignore"
	// MemoryFailureActionInject memory failure occurred in guest memory, the guest enabled MCE handling mechanism, and QEMU could inject the MCE into the guest successfully.
	MemoryFailureActionInject MemoryFailureAction = "inject"
	// MemoryFailureActionFatal the failure is unrecoverable. This occurs for action-required failures if the recipient is the hypervisor; QEMU will exit.
	MemoryFailureActionFatal MemoryFailureAction = "fatal"
	// MemoryFailureActionReset the failure is unrecoverable but confined to the guest. This occurs if the recipient is a guest guest which is not ready to handle memory failures.
	MemoryFailureActionReset MemoryFailureAction = "reset"
)

// MemoryFailureFlags
//
// Additional information on memory failures.
type MemoryFailureFlags struct {
	// ActionRequired whether a memory failure event is action-required or action-optional (e.g. a failure during memory scrub).
	ActionRequired bool `json:"action-required"`
	// Recursive whether the failure occurred while the previous failure was still in progress.
	Recursive bool `json:"recursive"`
}

// NotifyVmexitOption An enumeration of the options specified when enabling notify VM exit
type NotifyVmexitOption string

const (
	// NotifyVmexitOptionRun enable the feature, do nothing and continue if the notify VM exit happens.
	NotifyVmexitOptionRun NotifyVmexitOption = "run"
	// NotifyVmexitOptionInternalError enable the feature, raise a internal error if the notify VM exit happens.
	NotifyVmexitOptionInternalError NotifyVmexitOption = "internal-error"
	// NotifyVmexitOptionDisable disable the feature.
	NotifyVmexitOptionDisable NotifyVmexitOption = "disable"
)

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

// BiosAtaTranslation Policy that BIOS should use to interpret cylinder/head/sector addresses. Note that Bochs BIOS and SeaBIOS will not actually translate logical CHS to physical; instead, they will use logical block addressing.
type BiosAtaTranslation string

const (
	// BiosAtaTranslationAuto If cylinder/heads/sizes are passed, choose between none and LBA depending on the size of the disk. If they are not passed, choose none if QEMU can guess that the disk had 16 or fewer heads, large if QEMU can guess that the disk had 131072 or fewer tracks across all heads (i.e. cylinders*heads<131072), otherwise LBA.
	BiosAtaTranslationAuto BiosAtaTranslation = "auto"
	// BiosAtaTranslationNone The physical disk geometry is equal to the logical geometry.
	BiosAtaTranslationNone BiosAtaTranslation = "none"
	// BiosAtaTranslationLba Assume 63 sectors per track and one of 16, 32, 64, 128 or 255 heads (if fewer than 255 are enough to cover the whole disk with 1024 cylinders/head). The number of cylinders/head is then computed based on the number of sectors and heads.
	BiosAtaTranslationLba BiosAtaTranslation = "lba"
	// BiosAtaTranslationLarge The number of cylinders per head is scaled down to 1024 by correspondingly scaling up the number of heads.
	BiosAtaTranslationLarge BiosAtaTranslation = "large"
	// BiosAtaTranslationRechs Same as @large, but first convert a 16-head geometry to 15-head, by proportionally scaling up the number of cylinders/head.
	BiosAtaTranslationRechs BiosAtaTranslation = "rechs"
)

// FloppyDriveType Type of Floppy drive to be emulated by the Floppy Disk Controller.
type FloppyDriveType string

const (
	// FloppyDriveType144 1.44MB 3.5" drive
	FloppyDriveType144 FloppyDriveType = "144"
	// FloppyDriveType288 2.88MB 3.5" drive
	FloppyDriveType288 FloppyDriveType = "288"
	// FloppyDriveType120 1.2MB 5.25" drive
	FloppyDriveType120 FloppyDriveType = "120"
	// FloppyDriveTypeNone No drive connected
	FloppyDriveTypeNone FloppyDriveType = "none"
	// FloppyDriveTypeAuto Automatically determined by inserted media at boot
	FloppyDriveTypeAuto FloppyDriveType = "auto"
)

// PRManagerInfo
//
// Information about a persistent reservation manager
type PRManagerInfo struct {
	// Id the identifier of the persistent reservation manager
	Id string `json:"id"`
	// Connected true if the persistent reservation manager is connected to the underlying storage or helper
	Connected bool `json:"connected"`
}

// QueryPrManagers
//
// Returns a list of information about each persistent reservation manager.
type QueryPrManagers struct {
}

func (QueryPrManagers) Command() string {
	return "query-pr-managers"
}

func (cmd QueryPrManagers) Execute(ctx context.Context, client api.Client) ([]PRManagerInfo, error) {
	var ret []PRManagerInfo

	return ret, client.Execute(ctx, "query-pr-managers", cmd, &ret)
}

// Eject
//
// Ejects the medium from a removable drive.
type Eject struct {
	// Device Block device name
	Device *string `json:"device,omitempty"`
	// Id The name or QOM path of the guest device (since: 2.8)
	Id *string `json:"id,omitempty"`
	// Force If true, eject regardless of whether the drive is locked. If not specified, the default value is false.
	Force *bool `json:"force,omitempty"`
}

func (Eject) Command() string {
	return "eject"
}

func (cmd Eject) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "eject", cmd, nil)
}

// BlockdevOpenTray
//
// Opens a block device's tray. If there is a block driver state tree inserted as a medium, it will become inaccessible to the guest (but it will remain associated to the block device, so closing the tray will make it accessible again). If the tray was already open before, this will be a no-op. Once the tray opens, a DEVICE_TRAY_MOVED event is emitted. There
type BlockdevOpenTray struct {
	// Device Block device name
	Device *string `json:"device,omitempty"`
	// Id The name or QOM path of the guest device (since: 2.8)
	Id *string `json:"id,omitempty"`
	// Force if false (the default), an eject request will be sent to the guest if it has locked the tray (and the tray will not be opened immediately); if true, the tray will be opened regardless of whether it is locked
	Force *bool `json:"force,omitempty"`
}

func (BlockdevOpenTray) Command() string {
	return "blockdev-open-tray"
}

func (cmd BlockdevOpenTray) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-open-tray", cmd, nil)
}

// BlockdevCloseTray
//
// Closes a block device's tray. If there is a block driver state tree associated with the block device (which is currently ejected), that tree will be loaded as the medium. If the tray was already closed before, this will be a no-op.
type BlockdevCloseTray struct {
	// Device Block device name
	Device *string `json:"device,omitempty"`
	// Id The name or QOM path of the guest device (since: 2.8)
	Id *string `json:"id,omitempty"`
}

func (BlockdevCloseTray) Command() string {
	return "blockdev-close-tray"
}

func (cmd BlockdevCloseTray) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-close-tray", cmd, nil)
}

// BlockdevRemoveMedium
//
// Removes a medium (a block driver state tree) from a block device. That block device's tray must currently be open (unless there is no attached guest device). If the tray is open and there is no medium inserted, this will be a no-op.
type BlockdevRemoveMedium struct {
	// Id The name or QOM path of the guest device
	Id string `json:"id"`
}

func (BlockdevRemoveMedium) Command() string {
	return "blockdev-remove-medium"
}

func (cmd BlockdevRemoveMedium) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-remove-medium", cmd, nil)
}

// BlockdevInsertMedium
//
// Inserts a medium (a block driver state tree) into a block device. That block device's tray must currently be open (unless there is no attached guest device) and there must be no medium inserted already.
type BlockdevInsertMedium struct {
	// Id The name or QOM path of the guest device
	Id string `json:"id"`
	// NodeName name of a node in the block driver state graph
	NodeName string `json:"node-name"`
}

func (BlockdevInsertMedium) Command() string {
	return "blockdev-insert-medium"
}

func (cmd BlockdevInsertMedium) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-insert-medium", cmd, nil)
}

// BlockdevChangeReadOnlyMode Specifies the new read-only mode of a block device subject to the @blockdev-change-medium command.
type BlockdevChangeReadOnlyMode string

const (
	// BlockdevChangeReadOnlyModeRetain Retains the current read-only mode
	BlockdevChangeReadOnlyModeRetain BlockdevChangeReadOnlyMode = "retain"
	// BlockdevChangeReadOnlyModeReadOnly Makes the device read-only
	BlockdevChangeReadOnlyModeReadOnly BlockdevChangeReadOnlyMode = "read-only"
	// BlockdevChangeReadOnlyModeReadWrite Makes the device writable
	BlockdevChangeReadOnlyModeReadWrite BlockdevChangeReadOnlyMode = "read-write"
)

// BlockdevChangeMedium
//
// Changes the medium inserted into a block device by ejecting the current medium and loading a new image file which is inserted as the new medium (this command combines blockdev-open-tray, blockdev-remove-medium, blockdev-insert-medium and blockdev-close-tray).
type BlockdevChangeMedium struct {
	// Device Block device name
	Device *string `json:"device,omitempty"`
	// Id The name or QOM path of the guest device (since: 2.8)
	Id *string `json:"id,omitempty"`
	// Filename filename of the new image to be loaded
	Filename string `json:"filename"`
	// Format format to open the new image with (defaults to the probed format)
	Format *string `json:"format,omitempty"`
	// Force if false (the default), an eject request through blockdev-open-tray will be sent to the guest if it has locked the tray (and the tray will not be opened immediately); if true, the tray will be opened regardless of whether it is locked. (since 7.1)
	Force *bool `json:"force,omitempty"`
	// ReadOnlyMode change the read-only mode of the device; defaults to 'retain'
	ReadOnlyMode *BlockdevChangeReadOnlyMode `json:"read-only-mode,omitempty"`
}

func (BlockdevChangeMedium) Command() string {
	return "blockdev-change-medium"
}

func (cmd BlockdevChangeMedium) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "blockdev-change-medium", cmd, nil)
}

// DeviceTrayMovedEvent (DEVICE_TRAY_MOVED)
//
// Emitted whenever the tray of a removable device is moved by the guest or by HMP/QMP commands
type DeviceTrayMovedEvent struct {
	// Device Block device name. This is always present for compatibility reasons, but it can be empty ("") if the image does not have a device name associated.
	Device string `json:"device"`
	// Id The name or QOM path of the guest device (since 2.8)
	Id string `json:"id"`
	// TrayOpen true if the tray has been opened or false if it has been closed
	TrayOpen bool `json:"tray-open"`
}

func (DeviceTrayMovedEvent) Event() string {
	return "DEVICE_TRAY_MOVED"
}

// PrManagerStatusChangedEvent (PR_MANAGER_STATUS_CHANGED)
//
// Emitted whenever the connected status of a persistent reservation manager changes.
type PrManagerStatusChangedEvent struct {
	// Id The id of the PR manager object
	Id string `json:"id"`
	// Connected true if the PR manager is connected to a backend
	Connected bool `json:"connected"`
}

func (PrManagerStatusChangedEvent) Event() string {
	return "PR_MANAGER_STATUS_CHANGED"
}

// BlockSetIoThrottle
//
// Change I/O throttle limits for a block drive. Since QEMU 2.4, each device with I/O limits is member of a throttle group. If two or more devices are members of the same group, the limits will apply to the combined I/O of the whole group in a round-robin fashion. Therefore, setting new I/O limits to a device will affect the whole group. The name of the group can be specified using the 'group' parameter. If the parameter is unset, it is assumed to be the current group of that device. If it's not in any group yet, the name of the device will be used as the name for its group. The 'group' parameter can also be used to move a device to a different group. In this case the limits specified in the parameters will be applied to the new group only. I/O limits can be disabled by setting all of them to 0. In this case the device will be removed from its group and the rest of its members will not be affected. The 'group' parameter is ignored.
type BlockSetIoThrottle struct {
	BlockIOThrottle
}

func (BlockSetIoThrottle) Command() string {
	return "block_set_io_throttle"
}

func (cmd BlockSetIoThrottle) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block_set_io_throttle", cmd, nil)
}

// BlockLatencyHistogramSet
//
// Manage read, write and flush latency histograms for the device. If only @id parameter is specified, remove all present latency histograms for the device. Otherwise, add/reset some of (or all) latency histograms.
type BlockLatencyHistogramSet struct {
	// Id The name or QOM path of the guest device.
	Id string `json:"id"`
	// Boundaries list of interval boundary values (see description in BlockLatencyHistogramInfo definition). If specified, all latency histograms are removed, and empty ones created for all io types with intervals corresponding to @boundaries (except for io types, for which specific boundaries are set through the following parameters).
	Boundaries []uint64 `json:"boundaries,omitempty"`
	// BoundariesRead list of interval boundary values for read latency histogram. If specified, old read latency histogram is removed, and empty one created with intervals corresponding to @boundaries-read. The parameter has higher priority then @boundaries.
	BoundariesRead []uint64 `json:"boundaries-read,omitempty"`
	// BoundariesWrite list of interval boundary values for write latency histogram.
	BoundariesWrite []uint64 `json:"boundaries-write,omitempty"`
	// BoundariesZap list of interval boundary values for zone append write latency histogram.
	BoundariesZap []uint64 `json:"boundaries-zap,omitempty"`
	// BoundariesFlush list of interval boundary values for flush latency histogram.
	BoundariesFlush []uint64 `json:"boundaries-flush,omitempty"`
}

func (BlockLatencyHistogramSet) Command() string {
	return "block-latency-histogram-set"
}

func (cmd BlockLatencyHistogramSet) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "block-latency-histogram-set", cmd, nil)
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

// DumpGuestMemoryFormat An enumeration of guest-memory-dump's format.
type DumpGuestMemoryFormat string

const (
	// DumpGuestMemoryFormatElf elf format
	DumpGuestMemoryFormatElf DumpGuestMemoryFormat = "elf"
	// DumpGuestMemoryFormatKdumpZlib makedumpfile flattened, kdump-compressed format with zlib compression
	DumpGuestMemoryFormatKdumpZlib DumpGuestMemoryFormat = "kdump-zlib"
	// DumpGuestMemoryFormatKdumpLzo makedumpfile flattened, kdump-compressed format with lzo compression
	DumpGuestMemoryFormatKdumpLzo DumpGuestMemoryFormat = "kdump-lzo"
	// DumpGuestMemoryFormatKdumpSnappy makedumpfile flattened, kdump-compressed format with snappy compression
	DumpGuestMemoryFormatKdumpSnappy DumpGuestMemoryFormat = "kdump-snappy"
	// DumpGuestMemoryFormatKdumpRawZlib raw assembled kdump-compressed format with zlib compression (since 8.2)
	DumpGuestMemoryFormatKdumpRawZlib DumpGuestMemoryFormat = "kdump-raw-zlib"
	// DumpGuestMemoryFormatKdumpRawLzo raw assembled kdump-compressed format with lzo compression (since 8.2)
	DumpGuestMemoryFormatKdumpRawLzo DumpGuestMemoryFormat = "kdump-raw-lzo"
	// DumpGuestMemoryFormatKdumpRawSnappy raw assembled kdump-compressed format with snappy compression (since 8.2)
	DumpGuestMemoryFormatKdumpRawSnappy DumpGuestMemoryFormat = "kdump-raw-snappy"
	// DumpGuestMemoryFormatWinDmp Windows full crashdump format, can be used instead of ELF converting (since 2.13)
	DumpGuestMemoryFormatWinDmp DumpGuestMemoryFormat = "win-dmp"
)

// DumpGuestMemory
//
// Dump guest's memory to vmcore. It is a synchronous operation that can take very long depending on the amount of guest memory.
type DumpGuestMemory struct {
	// Paging if true, do paging to get guest's memory mapping. This allows using gdb to process the core file.
	Paging bool `json:"paging"`
	// Protocol the filename or file descriptor of the vmcore. The
	Protocol string `json:"protocol"`
	// Detach if true, QMP will return immediately rather than waiting for the dump to finish. The user can track progress using "query-dump". (since 2.6).
	Detach *bool `json:"detach,omitempty"`
	// Begin if specified, the starting physical address.
	Begin *int64 `json:"begin,omitempty"`
	// Length if specified, the memory size, in bytes. If you don't want to dump all guest's memory, please specify the start @begin and @length
	Length *int64 `json:"length,omitempty"`
	// Format if specified, the format of guest memory dump. But non-elf format is conflict with paging and filter, ie. @paging, @begin and @length is not allowed to be specified with non-elf @format at the same time (since 2.0)
	Format *DumpGuestMemoryFormat `json:"format,omitempty"`
}

func (DumpGuestMemory) Command() string {
	return "dump-guest-memory"
}

func (cmd DumpGuestMemory) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "dump-guest-memory", cmd, nil)
}

// DumpStatus Describe the status of a long-running background guest memory dump.
type DumpStatus string

const (
	// DumpStatusNone no dump-guest-memory has started yet.
	DumpStatusNone DumpStatus = "none"
	// DumpStatusActive there is one dump running in background.
	DumpStatusActive DumpStatus = "active"
	// DumpStatusCompleted the last dump has finished successfully.
	DumpStatusCompleted DumpStatus = "completed"
	// DumpStatusFailed the last dump has failed.
	DumpStatusFailed DumpStatus = "failed"
)

// DumpQueryResult
//
// The result format for 'query-dump'.
type DumpQueryResult struct {
	// Status enum of @DumpStatus, which shows current dump status
	Status DumpStatus `json:"status"`
	// Completed bytes written in latest dump (uncompressed)
	Completed int64 `json:"completed"`
	// Total total bytes to be written in latest dump (uncompressed)
	Total int64 `json:"total"`
}

// QueryDump
//
// Query latest dump status.
type QueryDump struct {
}

func (QueryDump) Command() string {
	return "query-dump"
}

func (cmd QueryDump) Execute(ctx context.Context, client api.Client) (DumpQueryResult, error) {
	var ret DumpQueryResult

	return ret, client.Execute(ctx, "query-dump", cmd, &ret)
}

// DumpCompletedEvent (DUMP_COMPLETED)
//
// Emitted when background dump has completed
type DumpCompletedEvent struct {
	// Result final dump status
	Result DumpQueryResult `json:"result"`
	// Error human-readable error string that provides hint on why dump failed. Only presents on failure. The user should not try to interpret the error string.
	Error *string `json:"error,omitempty"`
}

func (DumpCompletedEvent) Event() string {
	return "DUMP_COMPLETED"
}

// DumpGuestMemoryCapability
type DumpGuestMemoryCapability struct {
	// Formats the available formats for dump-guest-memory
	Formats []DumpGuestMemoryFormat `json:"formats"`
}

// QueryDumpGuestMemoryCapability
//
// Returns the available formats for dump-guest-memory
type QueryDumpGuestMemoryCapability struct {
}

func (QueryDumpGuestMemoryCapability) Command() string {
	return "query-dump-guest-memory-capability"
}

func (cmd QueryDumpGuestMemoryCapability) Execute(ctx context.Context, client api.Client) (DumpGuestMemoryCapability, error) {
	var ret DumpGuestMemoryCapability

	return ret, client.Execute(ctx, "query-dump-guest-memory-capability", cmd, &ret)
}

// SetLink
//
// Sets the link status of a virtual network adapter.
type SetLink struct {
	// Name the device name of the virtual network adapter
	Name string `json:"name"`
	// Up true to set the link status to be up
	Up bool `json:"up"`
}

func (SetLink) Command() string {
	return "set_link"
}

func (cmd SetLink) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "set_link", cmd, nil)
}

// NetdevAdd
//
// Add a network backend. Additional arguments depend on the type.
type NetdevAdd struct {
	Netdev
}

func (NetdevAdd) Command() string {
	return "netdev_add"
}

func (cmd NetdevAdd) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "netdev_add", cmd, nil)
}

// NetdevDel
//
// Remove a network backend.
type NetdevDel struct {
	// Id the name of the network backend to remove
	Id string `json:"id"`
}

func (NetdevDel) Command() string {
	return "netdev_del"
}

func (cmd NetdevDel) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "netdev_del", cmd, nil)
}

// NetLegacyNicOptions
//
// Create a new Network Interface Card.
type NetLegacyNicOptions struct {
	// Netdev id of -netdev to connect to
	Netdev *string `json:"netdev,omitempty"`
	// Macaddr MAC address
	Macaddr *string `json:"macaddr,omitempty"`
	// Model device model (e1000, rtl8139, virtio etc.)
	Model *string `json:"model,omitempty"`
	// Addr PCI device address
	Addr *string `json:"addr,omitempty"`
	// Vectors number of MSI-x vectors, 0 to disable MSI-X
	Vectors *uint32 `json:"vectors,omitempty"`
}

// String
//
// A fat type wrapping 'str', to be embedded in lists.
type String struct {
	Str string `json:"str"`
}

// NetdevUserOptions
//
// Use the user mode network stack which requires no administrator privilege to run.
type NetdevUserOptions struct {
	// Hostname client hostname reported by the builtin DHCP server
	Hostname *string `json:"hostname,omitempty"`
	// Restrict isolate the guest from the host
	Restrict *bool `json:"restrict,omitempty"`
	// Ipv4 whether to support IPv4, default true for enabled (since 2.6)
	Ipv4 *bool `json:"ipv4,omitempty"`
	// Ipv6 whether to support IPv6, default true for enabled (since 2.6)
	Ipv6 *bool `json:"ipv6,omitempty"`
	// Ip legacy parameter, use net= instead
	Ip *string `json:"ip,omitempty"`
	// Net IP network address that the guest will see, in the form addr[/netmask] The netmask is optional, and can be either in the form a.b.c.d or as a number of valid top-most bits. Default is 10.0.2.0/24.
	Net *string `json:"net,omitempty"`
	// Host guest-visible address of the host
	Host *string `json:"host,omitempty"`
	// Tftp root directory of the built-in TFTP server
	Tftp *string `json:"tftp,omitempty"`
	// Bootfile BOOTP filename, for use with tftp=
	Bootfile *string `json:"bootfile,omitempty"`
	// Dhcpstart the first of the 16 IPs the built-in DHCP server can assign
	Dhcpstart *string `json:"dhcpstart,omitempty"`
	// Dns guest-visible address of the virtual nameserver
	Dns *string `json:"dns,omitempty"`
	// Dnssearch list of DNS suffixes to search, passed as DHCP option to the guest
	Dnssearch []String `json:"dnssearch,omitempty"`
	// Domainname guest-visible domain name of the virtual nameserver (since 3.0)
	Domainname *string `json:"domainname,omitempty"`
	// Ipv6Prefix IPv6 network prefix (default is fec0::) (since 2.6). The network prefix is given in the usual hexadecimal IPv6 address notation.
	Ipv6Prefix *string `json:"ipv6-prefix,omitempty"`
	// Ipv6Prefixlen IPv6 network prefix length (default is 64) (since 2.6)
	Ipv6Prefixlen *int64 `json:"ipv6-prefixlen,omitempty"`
	// Ipv6Host guest-visible IPv6 address of the host (since 2.6)
	Ipv6Host *string `json:"ipv6-host,omitempty"`
	// Ipv6Dns guest-visible IPv6 address of the virtual nameserver (since 2.6)
	Ipv6Dns *string `json:"ipv6-dns,omitempty"`
	// Smb root directory of the built-in SMB server
	Smb *string `json:"smb,omitempty"`
	// Smbserver IP address of the built-in SMB server
	Smbserver *string `json:"smbserver,omitempty"`
	// Hostfwd redirect incoming TCP or UDP host connections to guest endpoints
	Hostfwd []String `json:"hostfwd,omitempty"`
	// Guestfwd forward guest TCP connections
	Guestfwd []String `json:"guestfwd,omitempty"`
	// TftpServerName RFC2132 "TFTP server name" string (Since 3.1)
	TftpServerName *string `json:"tftp-server-name,omitempty"`
}

// NetdevTapOptions
//
// Used to configure a host TAP network interface backend.
type NetdevTapOptions struct {
	// Ifname interface name
	Ifname *string `json:"ifname,omitempty"`
	// Fd file descriptor of an already opened tap
	Fd *string `json:"fd,omitempty"`
	// Fds multiple file descriptors of already opened multiqueue capable tap
	Fds *string `json:"fds,omitempty"`
	// Script script to initialize the interface
	Script *string `json:"script,omitempty"`
	// Downscript script to shut down the interface
	Downscript *string `json:"downscript,omitempty"`
	// Br bridge name (since 2.8)
	Br *string `json:"br,omitempty"`
	// Helper command to execute to configure bridge
	Helper *string `json:"helper,omitempty"`
	// Sndbuf send buffer limit. Understands [TGMKkb] suffixes.
	Sndbuf *uint64 `json:"sndbuf,omitempty"`
	// VnetHdr enable the IFF_VNET_HDR flag on the tap interface
	VnetHdr *bool `json:"vnet_hdr,omitempty"`
	// Vhost enable vhost-net network accelerator
	Vhost *bool `json:"vhost,omitempty"`
	// Vhostfd file descriptor of an already opened vhost net device
	Vhostfd *string `json:"vhostfd,omitempty"`
	// Vhostfds file descriptors of multiple already opened vhost net devices
	Vhostfds *string `json:"vhostfds,omitempty"`
	// Vhostforce vhost on for non-MSIX virtio guests
	Vhostforce *bool `json:"vhostforce,omitempty"`
	// Queues number of queues to be created for multiqueue capable tap
	Queues *uint32 `json:"queues,omitempty"`
	// PollUs maximum number of microseconds that could be spent on busy polling for tap (since 2.7)
	PollUs *uint32 `json:"poll-us,omitempty"`
}

// NetdevSocketOptions
//
// Socket netdevs are used to establish a network connection to another QEMU virtual machine via a TCP socket.
type NetdevSocketOptions struct {
	// Fd file descriptor of an already opened socket
	Fd *string `json:"fd,omitempty"`
	// Listen port number, and optional hostname, to listen on
	Listen *string `json:"listen,omitempty"`
	// Connect port number, and optional hostname, to connect to
	Connect *string `json:"connect,omitempty"`
	// Mcast UDP multicast address and port number
	Mcast *string `json:"mcast,omitempty"`
	// Localaddr source address and port for multicast and udp packets
	Localaddr *string `json:"localaddr,omitempty"`
	// Udp UDP unicast address and port number
	Udp *string `json:"udp,omitempty"`
}

// NetdevL2TPv3Options
//
// Configure an Ethernet over L2TPv3 tunnel.
type NetdevL2TPv3Options struct {
	// Src source address
	Src string `json:"src"`
	// Dst destination address
	Dst string `json:"dst"`
	// Srcport source port - mandatory for udp, optional for ip
	Srcport *string `json:"srcport,omitempty"`
	// Dstport destination port - mandatory for udp, optional for ip
	Dstport *string `json:"dstport,omitempty"`
	// Ipv6 force the use of ipv6
	Ipv6 *bool `json:"ipv6,omitempty"`
	// Udp use the udp version of l2tpv3 encapsulation
	Udp *bool `json:"udp,omitempty"`
	// Cookie64 use 64 bit cookies
	Cookie64 *bool `json:"cookie64,omitempty"`
	// Counter have sequence counter
	Counter *bool `json:"counter,omitempty"`
	// Pincounter pin sequence counter to zero - workaround for buggy implementations or networks with packet reorder
	Pincounter *bool `json:"pincounter,omitempty"`
	// Txcookie 32 or 64 bit transmit cookie
	Txcookie *uint64 `json:"txcookie,omitempty"`
	// Rxcookie 32 or 64 bit receive cookie
	Rxcookie *uint64 `json:"rxcookie,omitempty"`
	// Txsession 32 bit transmit session
	Txsession uint32 `json:"txsession"`
	// Rxsession 32 bit receive session - if not specified set to the same value as transmit
	Rxsession *uint32 `json:"rxsession,omitempty"`
	// Offset additional offset - allows the insertion of additional application-specific data before the packet payload
	Offset *uint32 `json:"offset,omitempty"`
}

// NetdevVdeOptions
//
// Connect to a vde switch running on the host.
type NetdevVdeOptions struct {
	// Sock socket path
	Sock *string `json:"sock,omitempty"`
	// Port port number
	Port *uint16 `json:"port,omitempty"`
	// Group group owner of socket
	Group *string `json:"group,omitempty"`
	// Mode permissions for socket
	Mode *uint16 `json:"mode,omitempty"`
}

// NetdevBridgeOptions
//
// Connect a host TAP network interface to a host bridge device.
type NetdevBridgeOptions struct {
	// Br bridge name
	Br *string `json:"br,omitempty"`
	// Helper command to execute to configure bridge
	Helper *string `json:"helper,omitempty"`
}

// NetdevHubPortOptions
//
// Connect two or more net clients through a software hub.
type NetdevHubPortOptions struct {
	// Hubid hub identifier number
	Hubid int32 `json:"hubid"`
	// Netdev used to connect hub to a netdev instead of a device (since 2.12)
	Netdev *string `json:"netdev,omitempty"`
}

// NetdevNetmapOptions
//
// Connect a client to a netmap-enabled NIC or to a VALE switch port
type NetdevNetmapOptions struct {
	// Ifname Either the name of an existing network interface supported by netmap, or the name of a VALE port (created on the fly). A
	Ifname string `json:"ifname"`
	// Devname path of the netmap device (default: '/dev/netmap').
	Devname *string `json:"devname,omitempty"`
}

// AFXDPMode Attach mode for a default XDP program
type AFXDPMode string

const (
	// AFXDPModeNative DRV mode, program is attached to a driver, packets are passed to the socket without allocation of skb.
	AFXDPModeNative AFXDPMode = "native"
	// AFXDPModeSkb generic mode, no driver support necessary
	AFXDPModeSkb AFXDPMode = "skb"
)

// NetdevAFXDPOptions
//
// AF_XDP network backend
type NetdevAFXDPOptions struct {
	// Ifname The name of an existing network interface.
	Ifname string `json:"ifname"`
	// Mode Attach mode for a default XDP program. If not specified, then 'native' will be tried first, then 'skb'.
	Mode *AFXDPMode `json:"mode,omitempty"`
	// ForceCopy Force XDP copy mode even if device supports zero-copy.
	ForceCopy *bool `json:"force-copy,omitempty"`
	// Queues number of queues to be used for multiqueue interfaces (default: 1).
	Queues *int64 `json:"queues,omitempty"`
	// StartQueue Use @queues starting from this queue number (default: 0).
	StartQueue *int64 `json:"start-queue,omitempty"`
	// Inhibit Don't load a default XDP program, use one already loaded to
	Inhibit *bool `json:"inhibit,omitempty"`
	// SockFds A colon (:) separated list of file descriptors for already open but not bound AF_XDP sockets in the queue order. One fd per queue. These descriptors should already be added into XDP socket map for corresponding queues. Requires @inhibit.
	SockFds *string `json:"sock-fds,omitempty"`
}

// NetdevVhostUserOptions
//
// Vhost-user network backend
type NetdevVhostUserOptions struct {
	// Chardev name of a unix socket chardev
	Chardev string `json:"chardev"`
	// Vhostforce vhost on for non-MSIX virtio guests (default: false).
	Vhostforce *bool `json:"vhostforce,omitempty"`
	// Queues number of queues to be created for multiqueue vhost-user
	Queues *int64 `json:"queues,omitempty"`
}

// NetdevVhostVDPAOptions
//
// Vhost-vdpa network backend vDPA device is a device that uses a datapath which complies with the virtio specifications with a vendor specific control path.
type NetdevVhostVDPAOptions struct {
	// Vhostdev path of vhost-vdpa device (default:'/dev/vhost-vdpa-0')
	Vhostdev *string `json:"vhostdev,omitempty"`
	// Vhostfd file descriptor of an already opened vhost vdpa device
	Vhostfd *string `json:"vhostfd,omitempty"`
	// Queues number of queues to be created for multiqueue vhost-vdpa
	Queues *int64 `json:"queues,omitempty"`
	// Svq Start device with (experimental) shadow virtqueue. (Since
	Svq *bool `json:"x-svq,omitempty"`
}

// NetdevVmnetHostOptions
//
// vmnet (host mode) network backend. Allows the vmnet interface to communicate with other vmnet interfaces that are in host mode and also with the host.
type NetdevVmnetHostOptions struct {
	// StartAddress The starting IPv4 address to use for the interface. Must be in the private IP range (RFC 1918). Must be specified along with @end-address and @subnet-mask. This address is used as the gateway address. The subsequent address up to and including end-address are placed in the DHCP pool.
	StartAddress *string `json:"start-address,omitempty"`
	// EndAddress The DHCP IPv4 range end address to use for the interface. Must be in the private IP range (RFC 1918). Must be specified along with @start-address and @subnet-mask.
	EndAddress *string `json:"end-address,omitempty"`
	// SubnetMask The IPv4 subnet mask to use on the interface. Must be specified along with @start-address and @subnet-mask.
	SubnetMask *string `json:"subnet-mask,omitempty"`
	// Isolated Enable isolation for this interface. Interface isolation ensures that vmnet interface is not able to communicate with any other vmnet interfaces. Only communication with host is allowed. Requires at least macOS Big Sur 11.0.
	Isolated *bool `json:"isolated,omitempty"`
	// NetUuid The identifier (UUID) to uniquely identify the isolated network vmnet interface should be added to. If set, no DHCP service is provided for this interface and network communication is allowed only with other interfaces added to this network identified by the UUID. Requires at least macOS Big Sur 11.0.
	NetUuid *string `json:"net-uuid,omitempty"`
}

// NetdevVmnetSharedOptions
//
// vmnet (shared mode) network backend. Allows traffic originating from the vmnet interface to reach the Internet through a network address translator (NAT). The vmnet interface can communicate with the host and with other shared mode interfaces on the same subnet. If no DHCP settings, subnet mask and IPv6 prefix specified, the interface can communicate with any of other interfaces in shared mode.
type NetdevVmnetSharedOptions struct {
	// StartAddress The starting IPv4 address to use for the interface. Must be in the private IP range (RFC 1918). Must be specified along with @end-address and @subnet-mask. This address is used as the gateway address. The subsequent address up to and including end-address are placed in the DHCP pool.
	StartAddress *string `json:"start-address,omitempty"`
	// EndAddress The DHCP IPv4 range end address to use for the interface. Must be in the private IP range (RFC 1918). Must be specified along with @start-address and @subnet-mask.
	EndAddress *string `json:"end-address,omitempty"`
	// SubnetMask The IPv4 subnet mask to use on the interface. Must be specified along with @start-address and @subnet-mask.
	SubnetMask *string `json:"subnet-mask,omitempty"`
	// Isolated Enable isolation for this interface. Interface isolation ensures that vmnet interface is not able to communicate with any other vmnet interfaces. Only communication with host is allowed. Requires at least macOS Big Sur 11.0.
	Isolated *bool `json:"isolated,omitempty"`
	// Nat66Prefix The IPv6 prefix to use into guest network. Must be a
	Nat66Prefix *string `json:"nat66-prefix,omitempty"`
}

// NetdevVmnetBridgedOptions
//
// vmnet (bridged mode) network backend. Bridges the vmnet interface with a physical network interface.
type NetdevVmnetBridgedOptions struct {
	// Ifname The name of the physical interface to be bridged.
	Ifname string `json:"ifname"`
	// Isolated Enable isolation for this interface. Interface isolation ensures that vmnet interface is not able to communicate with any other vmnet interfaces. Only communication with host is allowed. Requires at least macOS Big Sur 11.0.
	Isolated *bool `json:"isolated,omitempty"`
}

// NetdevStreamOptions
//
// Configuration info for stream socket netdev
type NetdevStreamOptions struct {
	// Addr socket address to listen on (server=true) or connect to (server=false)
	Addr SocketAddress `json:"addr"`
	// Server create server socket (default: false)
	Server *bool `json:"server,omitempty"`
	// Reconnect For a client socket, if a socket is disconnected, then attempt a reconnect after the given number of seconds. Setting
	Reconnect *uint32 `json:"reconnect,omitempty"`
}

// NetdevDgramOptions
//
// Configuration info for datagram socket netdev.
type NetdevDgramOptions struct {
	// Local local address Only SocketAddress types 'unix', 'inet' and 'fd' are supported. If remote address is present and it's a multicast address, local address is optional. Otherwise local address is required and remote address is optional.
	Local *SocketAddress `json:"local,omitempty"`
	// Remote remote address
	Remote *SocketAddress `json:"remote,omitempty"`
}

// NetClientDriver Available netdev drivers.
type NetClientDriver string

const (
	NetClientDriverNone NetClientDriver = "none"
	NetClientDriverNic  NetClientDriver = "nic"
	NetClientDriverUser NetClientDriver = "user"
	NetClientDriverTap  NetClientDriver = "tap"
	// NetClientDriverL2tpv3 since 2.1
	NetClientDriverL2tpv3 NetClientDriver = "l2tpv3"
	NetClientDriverSocket NetClientDriver = "socket"
	// NetClientDriverStream since 7.2
	NetClientDriverStream NetClientDriver = "stream"
	// NetClientDriverDgram since 7.2
	NetClientDriverDgram     NetClientDriver = "dgram"
	NetClientDriverVde       NetClientDriver = "vde"
	NetClientDriverBridge    NetClientDriver = "bridge"
	NetClientDriverHubport   NetClientDriver = "hubport"
	NetClientDriverNetmap    NetClientDriver = "netmap"
	NetClientDriverVhostUser NetClientDriver = "vhost-user"
	// NetClientDriverVhostVdpa since 5.1
	NetClientDriverVhostVdpa NetClientDriver = "vhost-vdpa"
	// NetClientDriverAfXdp since 8.2
	NetClientDriverAfXdp NetClientDriver = "af-xdp"
	// NetClientDriverVmnetHost since 7.1
	NetClientDriverVmnetHost NetClientDriver = "vmnet-host"
	// NetClientDriverVmnetShared since 7.1
	NetClientDriverVmnetShared NetClientDriver = "vmnet-shared"
	// NetClientDriverVmnetBridged since 7.1
	NetClientDriverVmnetBridged NetClientDriver = "vmnet-bridged"
)

// Netdev
//
// Captures the configuration of a network device.
type Netdev struct {
	// Discriminator: type

	// Id identifier for monitor commands.
	Id string `json:"id"`
	// Type Specify the driver used for interpreting remaining arguments.
	Type NetClientDriver `json:"type"`

	Nic          *NetLegacyNicOptions       `json:"-"`
	User         *NetdevUserOptions         `json:"-"`
	Tap          *NetdevTapOptions          `json:"-"`
	L2tpv3       *NetdevL2TPv3Options       `json:"-"`
	Socket       *NetdevSocketOptions       `json:"-"`
	Stream       *NetdevStreamOptions       `json:"-"`
	Dgram        *NetdevDgramOptions        `json:"-"`
	Vde          *NetdevVdeOptions          `json:"-"`
	Bridge       *NetdevBridgeOptions       `json:"-"`
	Hubport      *NetdevHubPortOptions      `json:"-"`
	Netmap       *NetdevNetmapOptions       `json:"-"`
	AfXdp        *NetdevAFXDPOptions        `json:"-"`
	VhostUser    *NetdevVhostUserOptions    `json:"-"`
	VhostVdpa    *NetdevVhostVDPAOptions    `json:"-"`
	VmnetHost    *NetdevVmnetHostOptions    `json:"-"`
	VmnetShared  *NetdevVmnetSharedOptions  `json:"-"`
	VmnetBridged *NetdevVmnetBridgedOptions `json:"-"`
}

func (u Netdev) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "nic":
		if u.Nic == nil {
			return nil, fmt.Errorf("expected Nic to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetLegacyNicOptions
		}{
			Id:                  u.Id,
			Type:                u.Type,
			NetLegacyNicOptions: u.Nic,
		})
	case "user":
		if u.User == nil {
			return nil, fmt.Errorf("expected User to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevUserOptions
		}{
			Id:                u.Id,
			Type:              u.Type,
			NetdevUserOptions: u.User,
		})
	case "tap":
		if u.Tap == nil {
			return nil, fmt.Errorf("expected Tap to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevTapOptions
		}{
			Id:               u.Id,
			Type:             u.Type,
			NetdevTapOptions: u.Tap,
		})
	case "l2tpv3":
		if u.L2tpv3 == nil {
			return nil, fmt.Errorf("expected L2tpv3 to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevL2TPv3Options
		}{
			Id:                  u.Id,
			Type:                u.Type,
			NetdevL2TPv3Options: u.L2tpv3,
		})
	case "socket":
		if u.Socket == nil {
			return nil, fmt.Errorf("expected Socket to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevSocketOptions
		}{
			Id:                  u.Id,
			Type:                u.Type,
			NetdevSocketOptions: u.Socket,
		})
	case "stream":
		if u.Stream == nil {
			return nil, fmt.Errorf("expected Stream to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevStreamOptions
		}{
			Id:                  u.Id,
			Type:                u.Type,
			NetdevStreamOptions: u.Stream,
		})
	case "dgram":
		if u.Dgram == nil {
			return nil, fmt.Errorf("expected Dgram to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevDgramOptions
		}{
			Id:                 u.Id,
			Type:               u.Type,
			NetdevDgramOptions: u.Dgram,
		})
	case "vde":
		if u.Vde == nil {
			return nil, fmt.Errorf("expected Vde to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevVdeOptions
		}{
			Id:               u.Id,
			Type:             u.Type,
			NetdevVdeOptions: u.Vde,
		})
	case "bridge":
		if u.Bridge == nil {
			return nil, fmt.Errorf("expected Bridge to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevBridgeOptions
		}{
			Id:                  u.Id,
			Type:                u.Type,
			NetdevBridgeOptions: u.Bridge,
		})
	case "hubport":
		if u.Hubport == nil {
			return nil, fmt.Errorf("expected Hubport to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevHubPortOptions
		}{
			Id:                   u.Id,
			Type:                 u.Type,
			NetdevHubPortOptions: u.Hubport,
		})
	case "netmap":
		if u.Netmap == nil {
			return nil, fmt.Errorf("expected Netmap to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevNetmapOptions
		}{
			Id:                  u.Id,
			Type:                u.Type,
			NetdevNetmapOptions: u.Netmap,
		})
	case "af-xdp":
		if u.AfXdp == nil {
			return nil, fmt.Errorf("expected AfXdp to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevAFXDPOptions
		}{
			Id:                 u.Id,
			Type:               u.Type,
			NetdevAFXDPOptions: u.AfXdp,
		})
	case "vhost-user":
		if u.VhostUser == nil {
			return nil, fmt.Errorf("expected VhostUser to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevVhostUserOptions
		}{
			Id:                     u.Id,
			Type:                   u.Type,
			NetdevVhostUserOptions: u.VhostUser,
		})
	case "vhost-vdpa":
		if u.VhostVdpa == nil {
			return nil, fmt.Errorf("expected VhostVdpa to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevVhostVDPAOptions
		}{
			Id:                     u.Id,
			Type:                   u.Type,
			NetdevVhostVDPAOptions: u.VhostVdpa,
		})
	case "vmnet-host":
		if u.VmnetHost == nil {
			return nil, fmt.Errorf("expected VmnetHost to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevVmnetHostOptions
		}{
			Id:                     u.Id,
			Type:                   u.Type,
			NetdevVmnetHostOptions: u.VmnetHost,
		})
	case "vmnet-shared":
		if u.VmnetShared == nil {
			return nil, fmt.Errorf("expected VmnetShared to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevVmnetSharedOptions
		}{
			Id:                       u.Id,
			Type:                     u.Type,
			NetdevVmnetSharedOptions: u.VmnetShared,
		})
	case "vmnet-bridged":
		if u.VmnetBridged == nil {
			return nil, fmt.Errorf("expected VmnetBridged to be set")
		}

		return json.Marshal(struct {
			Id   string          `json:"id"`
			Type NetClientDriver `json:"type"`
			*NetdevVmnetBridgedOptions
		}{
			Id:                        u.Id,
			Type:                      u.Type,
			NetdevVmnetBridgedOptions: u.VmnetBridged,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// RxState Packets receiving state
type RxState string

const (
	// RxStateNormal filter assigned packets according to the mac-table
	RxStateNormal RxState = "normal"
	// RxStateNone don't receive any assigned packet
	RxStateNone RxState = "none"
	// RxStateAll receive all assigned packets
	RxStateAll RxState = "all"
)

// RxFilterInfo
//
// Rx-filter information for a NIC.
type RxFilterInfo struct {
	// Name net client name
	Name string `json:"name"`
	// Promiscuous whether promiscuous mode is enabled
	Promiscuous bool `json:"promiscuous"`
	// Multicast multicast receive state
	Multicast RxState `json:"multicast"`
	// Unicast unicast receive state
	Unicast RxState `json:"unicast"`
	// Vlan vlan receive state (Since 2.0)
	Vlan RxState `json:"vlan"`
	// BroadcastAllowed whether to receive broadcast
	BroadcastAllowed bool `json:"broadcast-allowed"`
	// MulticastOverflow multicast table is overflowed or not
	MulticastOverflow bool `json:"multicast-overflow"`
	// UnicastOverflow unicast table is overflowed or not
	UnicastOverflow bool `json:"unicast-overflow"`
	// MainMac the main macaddr string
	MainMac string `json:"main-mac"`
	// VlanTable a list of active vlan id
	VlanTable []int64 `json:"vlan-table"`
	// UnicastTable a list of unicast macaddr string
	UnicastTable []string `json:"unicast-table"`
	// MulticastTable a list of multicast macaddr string
	MulticastTable []string `json:"multicast-table"`
}

// QueryRxFilter
//
// Return rx-filter information for all NICs (or for the given NIC).
type QueryRxFilter struct {
	// Name net client name
	Name *string `json:"name,omitempty"`
}

func (QueryRxFilter) Command() string {
	return "query-rx-filter"
}

func (cmd QueryRxFilter) Execute(ctx context.Context, client api.Client) ([]RxFilterInfo, error) {
	var ret []RxFilterInfo

	return ret, client.Execute(ctx, "query-rx-filter", cmd, &ret)
}

// NicRxFilterChangedEvent (NIC_RX_FILTER_CHANGED)
//
// Emitted once until the 'query-rx-filter' command is executed, the first event will always be emitted
type NicRxFilterChangedEvent struct {
	// Name net client name
	Name *string `json:"name,omitempty"`
	// Path device path
	Path string `json:"path"`
}

func (NicRxFilterChangedEvent) Event() string {
	return "NIC_RX_FILTER_CHANGED"
}

// AnnounceParameters
//
// Parameters for self-announce timers
type AnnounceParameters struct {
	// Initial Initial delay (in ms) before sending the first GARP/RARP announcement
	Initial int64 `json:"initial"`
	// Max Maximum delay (in ms) between GARP/RARP announcement packets
	Max int64 `json:"max"`
	// Rounds Number of self-announcement attempts
	Rounds int64 `json:"rounds"`
	// Step Delay increase (in ms) after each self-announcement attempt
	Step int64 `json:"step"`
	// Interfaces An optional list of interface names, which restricts the announcement to the listed interfaces. (Since 4.1)
	Interfaces []string `json:"interfaces,omitempty"`
	// Id A name to be used to identify an instance of announce-timers and to allow it to modified later. Not for use as part of the migration parameters. (Since 4.1)
	Id *string `json:"id,omitempty"`
}

// AnnounceSelf
//
// Trigger generation of broadcast RARP frames to update network switches. This can be useful when network bonds fail-over the active slave.
type AnnounceSelf struct {
	AnnounceParameters
}

func (AnnounceSelf) Command() string {
	return "announce-self"
}

func (cmd AnnounceSelf) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "announce-self", cmd, nil)
}

// FailoverNegotiatedEvent (FAILOVER_NEGOTIATED)
//
// Emitted when VIRTIO_NET_F_STANDBY was enabled during feature negotiation. Failover primary devices which were hidden (not hotplugged when requested) before will now be hotplugged by the virtio-net standby device.
type FailoverNegotiatedEvent struct {
	// DeviceId QEMU device id of the unplugged device
	DeviceId string `json:"device-id"`
}

func (FailoverNegotiatedEvent) Event() string {
	return "FAILOVER_NEGOTIATED"
}

// NetdevStreamConnectedEvent (NETDEV_STREAM_CONNECTED)
//
// Emitted when the netdev stream backend is connected
type NetdevStreamConnectedEvent struct {
	// NetdevId QEMU netdev id that is connected
	NetdevId string `json:"netdev-id"`
	// Addr The destination address
	Addr SocketAddress `json:"addr"`
}

func (NetdevStreamConnectedEvent) Event() string {
	return "NETDEV_STREAM_CONNECTED"
}

// NetdevStreamDisconnectedEvent (NETDEV_STREAM_DISCONNECTED)
//
// Emitted when the netdev stream backend is disconnected
type NetdevStreamDisconnectedEvent struct {
	// NetdevId QEMU netdev id that is disconnected
	NetdevId string `json:"netdev-id"`
}

func (NetdevStreamDisconnectedEvent) Event() string {
	return "NETDEV_STREAM_DISCONNECTED"
}

// RdmaGidStatusChangedEvent (RDMA_GID_STATUS_CHANGED)
//
// Emitted when guest driver adds/deletes GID to/from device
type RdmaGidStatusChangedEvent struct {
	// Netdev RoCE Network Device name
	Netdev string `json:"netdev"`
	// GidStatus Add or delete indication
	GidStatus bool `json:"gid-status"`
	// SubnetPrefix Subnet Prefix
	SubnetPrefix uint64 `json:"subnet-prefix"`
	// InterfaceId Interface ID
	InterfaceId uint64 `json:"interface-id"`
}

func (RdmaGidStatusChangedEvent) Event() string {
	return "RDMA_GID_STATUS_CHANGED"
}

// RockerSwitch
//
// Rocker switch information.
type RockerSwitch struct {
	// Name switch name
	Name string `json:"name"`
	// Id switch ID
	Id uint64 `json:"id"`
	// Ports number of front-panel ports
	Ports uint32 `json:"ports"`
}

// QueryRocker
//
// Return rocker switch information.
type QueryRocker struct {
	Name string `json:"name"`
}

func (QueryRocker) Command() string {
	return "query-rocker"
}

func (cmd QueryRocker) Execute(ctx context.Context, client api.Client) (RockerSwitch, error) {
	var ret RockerSwitch

	return ret, client.Execute(ctx, "query-rocker", cmd, &ret)
}

// RockerPortDuplex An eumeration of port duplex states.
type RockerPortDuplex string

const (
	// RockerPortDuplexHalf half duplex
	RockerPortDuplexHalf RockerPortDuplex = "half"
	// RockerPortDuplexFull full duplex
	RockerPortDuplexFull RockerPortDuplex = "full"
)

// RockerPortAutoneg An eumeration of port autoneg states.
type RockerPortAutoneg string

const (
	// RockerPortAutonegOff autoneg is off
	RockerPortAutonegOff RockerPortAutoneg = "off"
	// RockerPortAutonegOn autoneg is on
	RockerPortAutonegOn RockerPortAutoneg = "on"
)

// RockerPort
//
// Rocker switch port information.
type RockerPort struct {
	// Name port name
	Name string `json:"name"`
	// Enabled port is enabled for I/O
	Enabled bool `json:"enabled"`
	// LinkUp physical link is UP on port
	LinkUp bool `json:"link-up"`
	// Speed port link speed in Mbps
	Speed uint32 `json:"speed"`
	// Duplex port link duplex
	Duplex RockerPortDuplex `json:"duplex"`
	// Autoneg port link autoneg
	Autoneg RockerPortAutoneg `json:"autoneg"`
}

// QueryRockerPorts
//
// Return rocker switch port information.
type QueryRockerPorts struct {
	Name string `json:"name"`
}

func (QueryRockerPorts) Command() string {
	return "query-rocker-ports"
}

func (cmd QueryRockerPorts) Execute(ctx context.Context, client api.Client) ([]RockerPort, error) {
	var ret []RockerPort

	return ret, client.Execute(ctx, "query-rocker-ports", cmd, &ret)
}

// RockerOfDpaFlowKey
//
// Rocker switch OF-DPA flow key
type RockerOfDpaFlowKey struct {
	// Priority key priority, 0 being lowest priority
	Priority uint32 `json:"priority"`
	// TblId flow table ID
	TblId uint32 `json:"tbl-id"`
	// InPport physical input port
	InPport *uint32 `json:"in-pport,omitempty"`
	// TunnelId tunnel ID
	TunnelId *uint32 `json:"tunnel-id,omitempty"`
	// VlanId VLAN ID
	VlanId *uint16 `json:"vlan-id,omitempty"`
	// EthType Ethernet header type
	EthType *uint16 `json:"eth-type,omitempty"`
	// EthSrc Ethernet header source MAC address
	EthSrc *string `json:"eth-src,omitempty"`
	// EthDst Ethernet header destination MAC address
	EthDst *string `json:"eth-dst,omitempty"`
	// IpProto IP Header protocol field
	IpProto *uint8 `json:"ip-proto,omitempty"`
	// IpTos IP header TOS field
	IpTos *uint8 `json:"ip-tos,omitempty"`
	// IpDst IP header destination address
	IpDst *string `json:"ip-dst,omitempty"`
}

// RockerOfDpaFlowMask
//
// Rocker switch OF-DPA flow mask
type RockerOfDpaFlowMask struct {
	// InPport physical input port
	InPport *uint32 `json:"in-pport,omitempty"`
	// TunnelId tunnel ID
	TunnelId *uint32 `json:"tunnel-id,omitempty"`
	// VlanId VLAN ID
	VlanId *uint16 `json:"vlan-id,omitempty"`
	// EthSrc Ethernet header source MAC address
	EthSrc *string `json:"eth-src,omitempty"`
	// EthDst Ethernet header destination MAC address
	EthDst *string `json:"eth-dst,omitempty"`
	// IpProto IP Header protocol field
	IpProto *uint8 `json:"ip-proto,omitempty"`
	// IpTos IP header TOS field
	IpTos *uint8 `json:"ip-tos,omitempty"`
}

// RockerOfDpaFlowAction
//
// Rocker switch OF-DPA flow action
type RockerOfDpaFlowAction struct {
	// GotoTbl next table ID
	GotoTbl *uint32 `json:"goto-tbl,omitempty"`
	// GroupId group ID
	GroupId *uint32 `json:"group-id,omitempty"`
	// TunnelLport tunnel logical port ID
	TunnelLport *uint32 `json:"tunnel-lport,omitempty"`
	// VlanId VLAN ID
	VlanId *uint16 `json:"vlan-id,omitempty"`
	// NewVlanId new VLAN ID
	NewVlanId *uint16 `json:"new-vlan-id,omitempty"`
	// OutPport physical output port
	OutPport *uint32 `json:"out-pport,omitempty"`
}

// RockerOfDpaFlow
//
// Rocker switch OF-DPA flow
type RockerOfDpaFlow struct {
	// Cookie flow unique cookie ID
	Cookie uint64 `json:"cookie"`
	// Hits count of matches (hits) on flow
	Hits uint64 `json:"hits"`
	// Key flow key
	Key RockerOfDpaFlowKey `json:"key"`
	// Mask flow mask
	Mask RockerOfDpaFlowMask `json:"mask"`
	// Action flow action
	Action RockerOfDpaFlowAction `json:"action"`
}

// QueryRockerOfDpaFlows
//
// Return rocker OF-DPA flow information.
type QueryRockerOfDpaFlows struct {
	// Name switch name
	Name string `json:"name"`
	// TblId flow table ID. If tbl-id is not specified, returns flow information for all tables.
	TblId *uint32 `json:"tbl-id,omitempty"`
}

func (QueryRockerOfDpaFlows) Command() string {
	return "query-rocker-of-dpa-flows"
}

func (cmd QueryRockerOfDpaFlows) Execute(ctx context.Context, client api.Client) ([]RockerOfDpaFlow, error) {
	var ret []RockerOfDpaFlow

	return ret, client.Execute(ctx, "query-rocker-of-dpa-flows", cmd, &ret)
}

// RockerOfDpaGroup
//
// Rocker switch OF-DPA group
type RockerOfDpaGroup struct {
	// Id group unique ID
	Id uint32 `json:"id"`
	// Type group type
	Type uint8 `json:"type"`
	// VlanId VLAN ID
	VlanId *uint16 `json:"vlan-id,omitempty"`
	// Pport physical port number
	Pport *uint32 `json:"pport,omitempty"`
	// Index group index, unique with group type
	Index *uint32 `json:"index,omitempty"`
	// OutPport output physical port number
	OutPport *uint32 `json:"out-pport,omitempty"`
	// GroupId next group ID
	GroupId *uint32 `json:"group-id,omitempty"`
	// SetVlanId VLAN ID to set
	SetVlanId *uint16 `json:"set-vlan-id,omitempty"`
	// PopVlan pop VLAN headr from packet
	PopVlan *uint8 `json:"pop-vlan,omitempty"`
	// GroupIds list of next group IDs
	GroupIds []uint32 `json:"group-ids,omitempty"`
	// SetEthSrc set source MAC address in Ethernet header
	SetEthSrc *string `json:"set-eth-src,omitempty"`
	// SetEthDst set destination MAC address in Ethernet header
	SetEthDst *string `json:"set-eth-dst,omitempty"`
	// TtlCheck perform TTL check
	TtlCheck *uint8 `json:"ttl-check,omitempty"`
}

// QueryRockerOfDpaGroups
//
// Return rocker OF-DPA group information.
type QueryRockerOfDpaGroups struct {
	// Name switch name
	Name string `json:"name"`
	// Type group type. If type is not specified, returns group information for all group types.
	Type *uint8 `json:"type,omitempty"`
}

func (QueryRockerOfDpaGroups) Command() string {
	return "query-rocker-of-dpa-groups"
}

func (cmd QueryRockerOfDpaGroups) Execute(ctx context.Context, client api.Client) ([]RockerOfDpaGroup, error) {
	var ret []RockerOfDpaGroup

	return ret, client.Execute(ctx, "query-rocker-of-dpa-groups", cmd, &ret)
}

// TpmModel An enumeration of TPM models
type TpmModel string

const (
	// TpmModelTpmTis TPM TIS model
	TpmModelTpmTis TpmModel = "tpm-tis"
	// TpmModelTpmCrb TPM CRB model (since 2.12)
	TpmModelTpmCrb TpmModel = "tpm-crb"
	// TpmModelTpmSpapr TPM SPAPR model (since 5.0)
	TpmModelTpmSpapr TpmModel = "tpm-spapr"
)

// QueryTpmModels
//
// Return a list of supported TPM models
type QueryTpmModels struct {
}

func (QueryTpmModels) Command() string {
	return "query-tpm-models"
}

func (cmd QueryTpmModels) Execute(ctx context.Context, client api.Client) ([]TpmModel, error) {
	var ret []TpmModel

	return ret, client.Execute(ctx, "query-tpm-models", cmd, &ret)
}

// TpmType An enumeration of TPM types
type TpmType string

const (
	// TpmTypePassthrough TPM passthrough type
	TpmTypePassthrough TpmType = "passthrough"
	// TpmTypeEmulator Software Emulator TPM type (since 2.11)
	TpmTypeEmulator TpmType = "emulator"
)

// QueryTpmTypes
//
// Return a list of supported TPM types
type QueryTpmTypes struct {
}

func (QueryTpmTypes) Command() string {
	return "query-tpm-types"
}

func (cmd QueryTpmTypes) Execute(ctx context.Context, client api.Client) ([]TpmType, error) {
	var ret []TpmType

	return ret, client.Execute(ctx, "query-tpm-types", cmd, &ret)
}

// TPMPassthroughOptions
//
// Information about the TPM passthrough type
type TPMPassthroughOptions struct {
	// Path string describing the path used for accessing the TPM device
	Path *string `json:"path,omitempty"`
	// CancelPath string showing the TPM's sysfs cancel file for cancellation of TPM commands while they are executing
	CancelPath *string `json:"cancel-path,omitempty"`
}

// TPMEmulatorOptions
//
// Information about the TPM emulator type
type TPMEmulatorOptions struct {
	// Chardev Name of a unix socket chardev
	Chardev string `json:"chardev"`
}

// TPMPassthroughOptionsWrapper
type TPMPassthroughOptionsWrapper struct {
	// Data Information about the TPM passthrough type
	Data TPMPassthroughOptions `json:"data"`
}

// TPMEmulatorOptionsWrapper
type TPMEmulatorOptionsWrapper struct {
	// Data Information about the TPM emulator type
	Data TPMEmulatorOptions `json:"data"`
}

// TpmTypeOptions
//
// A union referencing different TPM backend types' configuration options
type TpmTypeOptions struct {
	// Discriminator: type

	// Type - 'passthrough' The configuration options for the TPM passthrough type - 'emulator' The configuration options for TPM emulator backend type
	Type TpmType `json:"type"`

	Passthrough *TPMPassthroughOptionsWrapper `json:"-"`
	Emulator    *TPMEmulatorOptionsWrapper    `json:"-"`
}

func (u TpmTypeOptions) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "passthrough":
		if u.Passthrough == nil {
			return nil, fmt.Errorf("expected Passthrough to be set")
		}

		return json.Marshal(struct {
			Type TpmType `json:"type"`
			*TPMPassthroughOptionsWrapper
		}{
			Type:                         u.Type,
			TPMPassthroughOptionsWrapper: u.Passthrough,
		})
	case "emulator":
		if u.Emulator == nil {
			return nil, fmt.Errorf("expected Emulator to be set")
		}

		return json.Marshal(struct {
			Type TpmType `json:"type"`
			*TPMEmulatorOptionsWrapper
		}{
			Type:                      u.Type,
			TPMEmulatorOptionsWrapper: u.Emulator,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// TPMInfo
//
// Information about the TPM
type TPMInfo struct {
	// Id The Id of the TPM
	Id string `json:"id"`
	// Model The TPM frontend model
	Model TpmModel `json:"model"`
	// Options The TPM (backend) type configuration options
	Options TpmTypeOptions `json:"options"`
}

// QueryTpm
//
// Return information about the TPM device
type QueryTpm struct {
}

func (QueryTpm) Command() string {
	return "query-tpm"
}

func (cmd QueryTpm) Execute(ctx context.Context, client api.Client) ([]TPMInfo, error) {
	var ret []TPMInfo

	return ret, client.Execute(ctx, "query-tpm", cmd, &ret)
}

// DisplayProtocol Display protocols which support changing password options.
type DisplayProtocol string

const (
	DisplayProtocolVnc   DisplayProtocol = "vnc"
	DisplayProtocolSpice DisplayProtocol = "spice"
)

// SetPasswordAction An action to take on changing a password on a connection with active clients.
type SetPasswordAction string

const (
	// SetPasswordActionKeep maintain existing clients
	SetPasswordActionKeep SetPasswordAction = "keep"
	// SetPasswordActionFail fail the command if clients are connected
	SetPasswordActionFail SetPasswordAction = "fail"
	// SetPasswordActionDisconnect disconnect existing clients
	SetPasswordActionDisconnect SetPasswordAction = "disconnect"
)

// SetPasswordOptions
//
// Options for set_password.
type SetPasswordOptions struct {
	// Discriminator: protocol

	// Protocol - 'vnc' to modify the VNC server password - 'spice' to modify the Spice server password
	Protocol DisplayProtocol `json:"protocol"`
	// Password the new password
	Password string `json:"password"`
	// Connected How to handle existing clients when changing the password. If nothing is specified, defaults to 'keep'. For VNC, only 'keep' is currently implemented.
	Connected *SetPasswordAction `json:"connected,omitempty"`

	Vnc *SetPasswordOptionsVnc `json:"-"`
}

func (u SetPasswordOptions) MarshalJSON() ([]byte, error) {
	switch u.Protocol {
	case "vnc":
		if u.Vnc == nil {
			return nil, fmt.Errorf("expected Vnc to be set")
		}

		return json.Marshal(struct {
			Protocol  DisplayProtocol    `json:"protocol"`
			Password  string             `json:"password"`
			Connected *SetPasswordAction `json:"connected,omitempty"`
			*SetPasswordOptionsVnc
		}{
			Protocol:              u.Protocol,
			Password:              u.Password,
			Connected:             u.Connected,
			SetPasswordOptionsVnc: u.Vnc,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// SetPasswordOptionsVnc
//
// Options for set_password specific to the VNC protocol.
type SetPasswordOptionsVnc struct {
	// Display The id of the display where the password should be changed. Defaults to the first.
	Display *string `json:"display,omitempty"`
}

// SetPassword
//
// Set the password of a remote display server.
type SetPassword struct {
	SetPasswordOptions
}

func (SetPassword) Command() string {
	return "set_password"
}

func (cmd SetPassword) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "set_password", cmd, nil)
}

// ExpirePasswordOptions
//
// General options for expire_password.
type ExpirePasswordOptions struct {
	// Discriminator: protocol

	// Protocol - 'vnc' to modify the VNC server expiration - 'spice' to modify the Spice server expiration
	Protocol DisplayProtocol `json:"protocol"`
	// Time when to expire the password. - 'now' to expire the password immediately - 'never' to cancel password expiration - '+INT' where INT is the number of seconds from now (integer) - 'INT' where INT is the absolute time in seconds
	Time string `json:"time"`

	Vnc *ExpirePasswordOptionsVnc `json:"-"`
}

func (u ExpirePasswordOptions) MarshalJSON() ([]byte, error) {
	switch u.Protocol {
	case "vnc":
		if u.Vnc == nil {
			return nil, fmt.Errorf("expected Vnc to be set")
		}

		return json.Marshal(struct {
			Protocol DisplayProtocol `json:"protocol"`
			Time     string          `json:"time"`
			*ExpirePasswordOptionsVnc
		}{
			Protocol:                 u.Protocol,
			Time:                     u.Time,
			ExpirePasswordOptionsVnc: u.Vnc,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// ExpirePasswordOptionsVnc
//
// Options for expire_password specific to the VNC protocol.
type ExpirePasswordOptionsVnc struct {
	// Display The id of the display where the expiration should be changed. Defaults to the first.
	Display *string `json:"display,omitempty"`
}

// ExpirePassword
//
// Expire the password of a remote display server.
type ExpirePassword struct {
	ExpirePasswordOptions
}

func (ExpirePassword) Command() string {
	return "expire_password"
}

func (cmd ExpirePassword) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "expire_password", cmd, nil)
}

// ImageFormat Supported image format types.
type ImageFormat string

const (
	// ImageFormatPpm PPM format
	ImageFormatPpm ImageFormat = "ppm"
	// ImageFormatPng PNG format
	ImageFormatPng ImageFormat = "png"
)

// Screendump
//
// Capture the contents of a screen and write it to a file.
type Screendump struct {
	// Filename the path of a new file to store the image
	Filename string `json:"filename"`
	// Device ID of the display device that should be dumped. If this parameter is missing, the primary display will be used. (Since 2.12)
	Device *string `json:"device,omitempty"`
	// Head head to use in case the device supports multiple heads. If this parameter is missing, head #0 will be used. Also note that the head can only be specified in conjunction with the device ID. (Since 2.12)
	Head *int64 `json:"head,omitempty"`
	// Format image format for screendump. (default: ppm) (Since 7.1)
	Format *ImageFormat `json:"format,omitempty"`
}

func (Screendump) Command() string {
	return "screendump"
}

func (cmd Screendump) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "screendump", cmd, nil)
}

// SpiceBasicInfo
//
// The basic information for SPICE network connection
type SpiceBasicInfo struct {
	// Host IP address
	Host string `json:"host"`
	// Port port number
	Port string `json:"port"`
	// Family address family
	Family NetworkAddressFamily `json:"family"`
}

// SpiceServerInfo
//
// Information about a SPICE server
type SpiceServerInfo struct {
	SpiceBasicInfo

	// Auth authentication method
	Auth *string `json:"auth,omitempty"`
}

// SpiceChannel
//
// Information about a SPICE client channel.
type SpiceChannel struct {
	SpiceBasicInfo

	// ConnectionId SPICE connection id number. All channels with the same id belong to the same SPICE session.
	ConnectionId int64 `json:"connection-id"`
	// ChannelType SPICE channel type number. "1" is the main control channel, filter for this one if you want to track spice sessions only
	ChannelType int64 `json:"channel-type"`
	// ChannelId SPICE channel ID number. Usually "0", might be different when multiple channels of the same type exist, such as multiple display channels in a multihead setup
	ChannelId int64 `json:"channel-id"`
	// Tls true if the channel is encrypted, false otherwise.
	Tls bool `json:"tls"`
}

// SpiceQueryMouseMode An enumeration of Spice mouse states.
type SpiceQueryMouseMode string

const (
	// SpiceQueryMouseModeClient Mouse cursor position is determined by the client.
	SpiceQueryMouseModeClient SpiceQueryMouseMode = "client"
	// SpiceQueryMouseModeServer Mouse cursor position is determined by the server.
	SpiceQueryMouseModeServer SpiceQueryMouseMode = "server"
	// SpiceQueryMouseModeUnknown No information is available about mouse mode used by the spice server.
	SpiceQueryMouseModeUnknown SpiceQueryMouseMode = "unknown"
)

// SpiceInfo
//
// Information about the SPICE session.
type SpiceInfo struct {
	// Enabled true if the SPICE server is enabled, false otherwise
	Enabled bool `json:"enabled"`
	// Migrated true if the last guest migration completed and spice migration had completed as well. false otherwise. (since 1.4)
	Migrated bool `json:"migrated"`
	// Host The hostname the SPICE server is bound to. This depends on the name resolution on the host and may be an IP address.
	Host *string `json:"host,omitempty"`
	// Port The SPICE server's port number.
	Port *int64 `json:"port,omitempty"`
	// TlsPort The SPICE server's TLS port number.
	TlsPort *int64 `json:"tls-port,omitempty"`
	// Auth the current authentication type used by the server - 'none' if no authentication is being used - 'spice' uses SASL or direct TLS authentication, depending on command line options
	Auth *string `json:"auth,omitempty"`
	// CompiledVersion SPICE server version.
	CompiledVersion *string `json:"compiled-version,omitempty"`
	// MouseMode The mode in which the mouse cursor is displayed currently. Can be determined by the client or the server, or unknown if spice server doesn't provide this information.
	MouseMode SpiceQueryMouseMode `json:"mouse-mode"`
	// Channels a list of @SpiceChannel for each active spice channel
	Channels []SpiceChannel `json:"channels,omitempty"`
}

// QuerySpice
//
// Returns information about the current SPICE server
type QuerySpice struct {
}

func (QuerySpice) Command() string {
	return "query-spice"
}

func (cmd QuerySpice) Execute(ctx context.Context, client api.Client) (SpiceInfo, error) {
	var ret SpiceInfo

	return ret, client.Execute(ctx, "query-spice", cmd, &ret)
}

// SpiceConnectedEvent (SPICE_CONNECTED)
//
// Emitted when a SPICE client establishes a connection
type SpiceConnectedEvent struct {
	// Server server information
	Server SpiceBasicInfo `json:"server"`
	// Client client information
	Client SpiceBasicInfo `json:"client"`
}

func (SpiceConnectedEvent) Event() string {
	return "SPICE_CONNECTED"
}

// SpiceInitializedEvent (SPICE_INITIALIZED)
//
// Emitted after initial handshake and authentication takes place (if any) and the SPICE channel is up and running
type SpiceInitializedEvent struct {
	// Server server information
	Server SpiceServerInfo `json:"server"`
	// Client client information
	Client SpiceChannel `json:"client"`
}

func (SpiceInitializedEvent) Event() string {
	return "SPICE_INITIALIZED"
}

// SpiceDisconnectedEvent (SPICE_DISCONNECTED)
//
// Emitted when the SPICE connection is closed
type SpiceDisconnectedEvent struct {
	// Server server information
	Server SpiceBasicInfo `json:"server"`
	// Client client information
	Client SpiceBasicInfo `json:"client"`
}

func (SpiceDisconnectedEvent) Event() string {
	return "SPICE_DISCONNECTED"
}

// SpiceMigrateCompletedEvent (SPICE_MIGRATE_COMPLETED)
//
// Emitted when SPICE migration has completed
type SpiceMigrateCompletedEvent struct {
}

func (SpiceMigrateCompletedEvent) Event() string {
	return "SPICE_MIGRATE_COMPLETED"
}

// VncBasicInfo
//
// The basic information for vnc network connection
type VncBasicInfo struct {
	// Host IP address
	Host string `json:"host"`
	// Service The service name of the vnc port. This may depend on the host system's service database so symbolic names should not be relied on.
	Service string `json:"service"`
	// Family address family
	Family NetworkAddressFamily `json:"family"`
	// Websocket true in case the socket is a websocket (since 2.3).
	Websocket bool `json:"websocket"`
}

// VncServerInfo
//
// The network connection information for server
type VncServerInfo struct {
	VncBasicInfo

	// Auth authentication method used for the plain (non-websocket) VNC server
	Auth *string `json:"auth,omitempty"`
}

// VncClientInfo
//
// Information about a connected VNC client.
type VncClientInfo struct {
	VncBasicInfo

	// X509Dname If x509 authentication is in use, the Distinguished Name of the client.
	X509Dname *string `json:"x509_dname,omitempty"`
	// SaslUsername If SASL authentication is in use, the SASL username used for authentication.
	SaslUsername *string `json:"sasl_username,omitempty"`
}

// VncInfo
//
// Information about the VNC session.
type VncInfo struct {
	// Enabled true if the VNC server is enabled, false otherwise
	Enabled bool `json:"enabled"`
	// Host The hostname the VNC server is bound to. This depends on the name resolution on the host and may be an IP address.
	Host *string `json:"host,omitempty"`
	// Family - 'ipv6' if the host is listening for IPv6 connections - 'ipv4' if the host is listening for IPv4 connections - 'unix' if the host is listening on a unix domain socket - 'unknown' otherwise
	Family *NetworkAddressFamily `json:"family,omitempty"`
	// Service The service name of the server's port. This may depends on the host system's service database so symbolic names should not be relied on.
	Service *string `json:"service,omitempty"`
	// Auth the current authentication type used by the server - 'none' if no authentication is being used - 'vnc' if VNC authentication is being used - 'vencrypt+plain' if VEncrypt is used with plain text authentication - 'vencrypt+tls+none' if VEncrypt is used with TLS and no authentication - 'vencrypt+tls+vnc' if VEncrypt is used with TLS and VNC authentication - 'vencrypt+tls+plain' if VEncrypt is used with TLS and plain text auth - 'vencrypt+x509+none' if VEncrypt is used with x509 and no auth - 'vencrypt+x509+vnc' if VEncrypt is used with x509 and VNC auth - 'vencrypt+x509+plain' if VEncrypt is used with x509 and plain text auth - 'vencrypt+tls+sasl' if VEncrypt is used with TLS and SASL auth - 'vencrypt+x509+sasl' if VEncrypt is used with x509 and SASL auth
	Auth *string `json:"auth,omitempty"`
	// Clients a list of @VncClientInfo of all currently connected clients
	Clients []VncClientInfo `json:"clients,omitempty"`
}

// VncPrimaryAuth vnc primary authentication method.
type VncPrimaryAuth string

const (
	VncPrimaryAuthNone     VncPrimaryAuth = "none"
	VncPrimaryAuthVnc      VncPrimaryAuth = "vnc"
	VncPrimaryAuthRa2      VncPrimaryAuth = "ra2"
	VncPrimaryAuthRa2ne    VncPrimaryAuth = "ra2ne"
	VncPrimaryAuthTight    VncPrimaryAuth = "tight"
	VncPrimaryAuthUltra    VncPrimaryAuth = "ultra"
	VncPrimaryAuthTls      VncPrimaryAuth = "tls"
	VncPrimaryAuthVencrypt VncPrimaryAuth = "vencrypt"
	VncPrimaryAuthSasl     VncPrimaryAuth = "sasl"
)

// VncVencryptSubAuth vnc sub authentication method with vencrypt.
type VncVencryptSubAuth string

const (
	VncVencryptSubAuthPlain     VncVencryptSubAuth = "plain"
	VncVencryptSubAuthTlsNone   VncVencryptSubAuth = "tls-none"
	VncVencryptSubAuthX509None  VncVencryptSubAuth = "x509-none"
	VncVencryptSubAuthTlsVnc    VncVencryptSubAuth = "tls-vnc"
	VncVencryptSubAuthX509Vnc   VncVencryptSubAuth = "x509-vnc"
	VncVencryptSubAuthTlsPlain  VncVencryptSubAuth = "tls-plain"
	VncVencryptSubAuthX509Plain VncVencryptSubAuth = "x509-plain"
	VncVencryptSubAuthTlsSasl   VncVencryptSubAuth = "tls-sasl"
	VncVencryptSubAuthX509Sasl  VncVencryptSubAuth = "x509-sasl"
)

// VncServerInfo2
//
// The network connection information for server
type VncServerInfo2 struct {
	VncBasicInfo

	// Auth The current authentication type used by the servers
	Auth VncPrimaryAuth `json:"auth"`
	// Vencrypt The vencrypt sub authentication type used by the servers, only specified in case auth == vencrypt.
	Vencrypt *VncVencryptSubAuth `json:"vencrypt,omitempty"`
}

// VncInfo2
//
// Information about a vnc server
type VncInfo2 struct {
	// Id vnc server name.
	Id string `json:"id"`
	// Server A list of @VncBasincInfo describing all listening sockets. The list can be empty (in case the vnc server is disabled). It
	Server []VncServerInfo2 `json:"server"`
	// Clients A list of @VncClientInfo of all currently connected clients. The list can be empty, for obvious reasons.
	Clients []VncClientInfo `json:"clients"`
	// Auth The current authentication type used by the non-websockets servers
	Auth VncPrimaryAuth `json:"auth"`
	// Vencrypt The vencrypt authentication type used by the servers, only specified in case auth == vencrypt.
	Vencrypt *VncVencryptSubAuth `json:"vencrypt,omitempty"`
	// Display The display device the vnc server is linked to.
	Display *string `json:"display,omitempty"`
}

// QueryVnc
//
// Returns information about the current VNC server
type QueryVnc struct {
}

func (QueryVnc) Command() string {
	return "query-vnc"
}

func (cmd QueryVnc) Execute(ctx context.Context, client api.Client) (VncInfo, error) {
	var ret VncInfo

	return ret, client.Execute(ctx, "query-vnc", cmd, &ret)
}

// QueryVncServers
//
// Returns a list of vnc servers. The list can be empty.
type QueryVncServers struct {
}

func (QueryVncServers) Command() string {
	return "query-vnc-servers"
}

func (cmd QueryVncServers) Execute(ctx context.Context, client api.Client) ([]VncInfo2, error) {
	var ret []VncInfo2

	return ret, client.Execute(ctx, "query-vnc-servers", cmd, &ret)
}

// ChangeVncPassword
//
// Change the VNC server password.
type ChangeVncPassword struct {
	// Password the new password to use with VNC authentication
	Password string `json:"password"`
}

func (ChangeVncPassword) Command() string {
	return "change-vnc-password"
}

func (cmd ChangeVncPassword) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "change-vnc-password", cmd, nil)
}

// VncConnectedEvent (VNC_CONNECTED)
//
// Emitted when a VNC client establishes a connection
type VncConnectedEvent struct {
	// Server server information
	Server VncServerInfo `json:"server"`
	// Client client information
	Client VncBasicInfo `json:"client"`
}

func (VncConnectedEvent) Event() string {
	return "VNC_CONNECTED"
}

// VncInitializedEvent (VNC_INITIALIZED)
//
// Emitted after authentication takes place (if any) and the VNC session is made active
type VncInitializedEvent struct {
	// Server server information
	Server VncServerInfo `json:"server"`
	// Client client information
	Client VncClientInfo `json:"client"`
}

func (VncInitializedEvent) Event() string {
	return "VNC_INITIALIZED"
}

// VncDisconnectedEvent (VNC_DISCONNECTED)
//
// Emitted when the connection is closed
type VncDisconnectedEvent struct {
	// Server server information
	Server VncServerInfo `json:"server"`
	// Client client information
	Client VncClientInfo `json:"client"`
}

func (VncDisconnectedEvent) Event() string {
	return "VNC_DISCONNECTED"
}

// MouseInfo
//
// Information about a mouse device.
type MouseInfo struct {
	// Name the name of the mouse device
	Name string `json:"name"`
	// Index the index of the mouse device
	Index int64 `json:"index"`
	// Current true if this device is currently receiving mouse events
	Current bool `json:"current"`
	// Absolute true if this device supports absolute coordinates as input
	Absolute bool `json:"absolute"`
}

// QueryMice
//
// Returns information about each active mouse device
type QueryMice struct {
}

func (QueryMice) Command() string {
	return "query-mice"
}

func (cmd QueryMice) Execute(ctx context.Context, client api.Client) ([]MouseInfo, error) {
	var ret []MouseInfo

	return ret, client.Execute(ctx, "query-mice", cmd, &ret)
}

// QKeyCode An enumeration of key name. This is used by the @send-key command.
type QKeyCode string

const (
	// QKeyCodeUnmapped since 2.0
	QKeyCodeUnmapped     QKeyCode = "unmapped"
	QKeyCodeShift        QKeyCode = "shift"
	QKeyCodeShiftR       QKeyCode = "shift_r"
	QKeyCodeAlt          QKeyCode = "alt"
	QKeyCodeAltR         QKeyCode = "alt_r"
	QKeyCodeCtrl         QKeyCode = "ctrl"
	QKeyCodeCtrlR        QKeyCode = "ctrl_r"
	QKeyCodeMenu         QKeyCode = "menu"
	QKeyCodeEsc          QKeyCode = "esc"
	QKeyCode1            QKeyCode = "1"
	QKeyCode2            QKeyCode = "2"
	QKeyCode3            QKeyCode = "3"
	QKeyCode4            QKeyCode = "4"
	QKeyCode5            QKeyCode = "5"
	QKeyCode6            QKeyCode = "6"
	QKeyCode7            QKeyCode = "7"
	QKeyCode8            QKeyCode = "8"
	QKeyCode9            QKeyCode = "9"
	QKeyCode0            QKeyCode = "0"
	QKeyCodeMinus        QKeyCode = "minus"
	QKeyCodeEqual        QKeyCode = "equal"
	QKeyCodeBackspace    QKeyCode = "backspace"
	QKeyCodeTab          QKeyCode = "tab"
	QKeyCodeQ            QKeyCode = "q"
	QKeyCodeW            QKeyCode = "w"
	QKeyCodeE            QKeyCode = "e"
	QKeyCodeR            QKeyCode = "r"
	QKeyCodeT            QKeyCode = "t"
	QKeyCodeY            QKeyCode = "y"
	QKeyCodeU            QKeyCode = "u"
	QKeyCodeI            QKeyCode = "i"
	QKeyCodeO            QKeyCode = "o"
	QKeyCodeP            QKeyCode = "p"
	QKeyCodeBracketLeft  QKeyCode = "bracket_left"
	QKeyCodeBracketRight QKeyCode = "bracket_right"
	QKeyCodeRet          QKeyCode = "ret"
	QKeyCodeA            QKeyCode = "a"
	QKeyCodeS            QKeyCode = "s"
	QKeyCodeD            QKeyCode = "d"
	QKeyCodeF            QKeyCode = "f"
	QKeyCodeG            QKeyCode = "g"
	QKeyCodeH            QKeyCode = "h"
	QKeyCodeJ            QKeyCode = "j"
	QKeyCodeK            QKeyCode = "k"
	QKeyCodeL            QKeyCode = "l"
	QKeyCodeSemicolon    QKeyCode = "semicolon"
	QKeyCodeApostrophe   QKeyCode = "apostrophe"
	QKeyCodeGraveAccent  QKeyCode = "grave_accent"
	QKeyCodeBackslash    QKeyCode = "backslash"
	QKeyCodeZ            QKeyCode = "z"
	QKeyCodeX            QKeyCode = "x"
	QKeyCodeC            QKeyCode = "c"
	QKeyCodeV            QKeyCode = "v"
	QKeyCodeB            QKeyCode = "b"
	QKeyCodeN            QKeyCode = "n"
	QKeyCodeM            QKeyCode = "m"
	QKeyCodeComma        QKeyCode = "comma"
	QKeyCodeDot          QKeyCode = "dot"
	QKeyCodeSlash        QKeyCode = "slash"
	QKeyCodeAsterisk     QKeyCode = "asterisk"
	QKeyCodeSpc          QKeyCode = "spc"
	QKeyCodeCapsLock     QKeyCode = "caps_lock"
	QKeyCodeF1           QKeyCode = "f1"
	QKeyCodeF2           QKeyCode = "f2"
	QKeyCodeF3           QKeyCode = "f3"
	QKeyCodeF4           QKeyCode = "f4"
	QKeyCodeF5           QKeyCode = "f5"
	QKeyCodeF6           QKeyCode = "f6"
	QKeyCodeF7           QKeyCode = "f7"
	QKeyCodeF8           QKeyCode = "f8"
	QKeyCodeF9           QKeyCode = "f9"
	QKeyCodeF10          QKeyCode = "f10"
	QKeyCodeNumLock      QKeyCode = "num_lock"
	QKeyCodeScrollLock   QKeyCode = "scroll_lock"
	QKeyCodeKpDivide     QKeyCode = "kp_divide"
	QKeyCodeKpMultiply   QKeyCode = "kp_multiply"
	QKeyCodeKpSubtract   QKeyCode = "kp_subtract"
	QKeyCodeKpAdd        QKeyCode = "kp_add"
	QKeyCodeKpEnter      QKeyCode = "kp_enter"
	QKeyCodeKpDecimal    QKeyCode = "kp_decimal"
	QKeyCodeSysrq        QKeyCode = "sysrq"
	QKeyCodeKp0          QKeyCode = "kp_0"
	QKeyCodeKp1          QKeyCode = "kp_1"
	QKeyCodeKp2          QKeyCode = "kp_2"
	QKeyCodeKp3          QKeyCode = "kp_3"
	QKeyCodeKp4          QKeyCode = "kp_4"
	QKeyCodeKp5          QKeyCode = "kp_5"
	QKeyCodeKp6          QKeyCode = "kp_6"
	QKeyCodeKp7          QKeyCode = "kp_7"
	QKeyCodeKp8          QKeyCode = "kp_8"
	QKeyCodeKp9          QKeyCode = "kp_9"
	QKeyCodeLess         QKeyCode = "less"
	QKeyCodeF11          QKeyCode = "f11"
	QKeyCodeF12          QKeyCode = "f12"
	QKeyCodePrint        QKeyCode = "print"
	QKeyCodeHome         QKeyCode = "home"
	QKeyCodePgup         QKeyCode = "pgup"
	QKeyCodePgdn         QKeyCode = "pgdn"
	QKeyCodeEnd          QKeyCode = "end"
	QKeyCodeLeft         QKeyCode = "left"
	QKeyCodeUp           QKeyCode = "up"
	QKeyCodeDown         QKeyCode = "down"
	QKeyCodeRight        QKeyCode = "right"
	QKeyCodeInsert       QKeyCode = "insert"
	QKeyCodeDelete       QKeyCode = "delete"
	QKeyCodeStop         QKeyCode = "stop"
	QKeyCodeAgain        QKeyCode = "again"
	QKeyCodeProps        QKeyCode = "props"
	QKeyCodeUndo         QKeyCode = "undo"
	QKeyCodeFront        QKeyCode = "front"
	QKeyCodeCopy         QKeyCode = "copy"
	QKeyCodeOpen         QKeyCode = "open"
	QKeyCodePaste        QKeyCode = "paste"
	QKeyCodeFind         QKeyCode = "find"
	QKeyCodeCut          QKeyCode = "cut"
	QKeyCodeLf           QKeyCode = "lf"
	QKeyCodeHelp         QKeyCode = "help"
	QKeyCodeMetaL        QKeyCode = "meta_l"
	QKeyCodeMetaR        QKeyCode = "meta_r"
	QKeyCodeCompose      QKeyCode = "compose"
	// QKeyCodePause since 2.0
	QKeyCodePause QKeyCode = "pause"
	// QKeyCodeRo since 2.4
	QKeyCodeRo QKeyCode = "ro"
	// QKeyCodeHiragana since 2.9
	QKeyCodeHiragana QKeyCode = "hiragana"
	// QKeyCodeHenkan since 2.9
	QKeyCodeHenkan QKeyCode = "henkan"
	// QKeyCodeYen since 2.9
	QKeyCodeYen QKeyCode = "yen"
	// QKeyCodeMuhenkan since 2.12
	QKeyCodeMuhenkan QKeyCode = "muhenkan"
	// QKeyCodeKatakanahiragana since 2.12
	QKeyCodeKatakanahiragana QKeyCode = "katakanahiragana"
	// QKeyCodeKpComma since 2.4
	QKeyCodeKpComma QKeyCode = "kp_comma"
	// QKeyCodeKpEquals since 2.6
	QKeyCodeKpEquals QKeyCode = "kp_equals"
	// QKeyCodePower since 2.6
	QKeyCodePower QKeyCode = "power"
	// QKeyCodeSleep since 2.10
	QKeyCodeSleep QKeyCode = "sleep"
	// QKeyCodeWake since 2.10
	QKeyCodeWake QKeyCode = "wake"
	// QKeyCodeAudionext since 2.10
	QKeyCodeAudionext QKeyCode = "audionext"
	// QKeyCodeAudioprev since 2.10
	QKeyCodeAudioprev QKeyCode = "audioprev"
	// QKeyCodeAudiostop since 2.10
	QKeyCodeAudiostop QKeyCode = "audiostop"
	// QKeyCodeAudioplay since 2.10
	QKeyCodeAudioplay QKeyCode = "audioplay"
	// QKeyCodeAudiomute since 2.10
	QKeyCodeAudiomute QKeyCode = "audiomute"
	// QKeyCodeVolumeup since 2.10
	QKeyCodeVolumeup QKeyCode = "volumeup"
	// QKeyCodeVolumedown since 2.10
	QKeyCodeVolumedown QKeyCode = "volumedown"
	// QKeyCodeMediaselect since 2.10
	QKeyCodeMediaselect QKeyCode = "mediaselect"
	// QKeyCodeMail since 2.10
	QKeyCodeMail QKeyCode = "mail"
	// QKeyCodeCalculator since 2.10
	QKeyCodeCalculator QKeyCode = "calculator"
	// QKeyCodeComputer since 2.10
	QKeyCodeComputer QKeyCode = "computer"
	// QKeyCodeAcHome since 2.10
	QKeyCodeAcHome QKeyCode = "ac_home"
	// QKeyCodeAcBack since 2.10
	QKeyCodeAcBack QKeyCode = "ac_back"
	// QKeyCodeAcForward since 2.10
	QKeyCodeAcForward QKeyCode = "ac_forward"
	// QKeyCodeAcRefresh since 2.10
	QKeyCodeAcRefresh QKeyCode = "ac_refresh"
	// QKeyCodeAcBookmarks since 2.10
	QKeyCodeAcBookmarks QKeyCode = "ac_bookmarks"
	// QKeyCodeLang1 since 6.1
	QKeyCodeLang1 QKeyCode = "lang1"
	// QKeyCodeLang2 since 6.1
	QKeyCodeLang2 QKeyCode = "lang2"
	// QKeyCodeF13 since 8.0
	QKeyCodeF13 QKeyCode = "f13"
	// QKeyCodeF14 since 8.0
	QKeyCodeF14 QKeyCode = "f14"
	// QKeyCodeF15 since 8.0
	QKeyCodeF15 QKeyCode = "f15"
	// QKeyCodeF16 since 8.0
	QKeyCodeF16 QKeyCode = "f16"
	// QKeyCodeF17 since 8.0
	QKeyCodeF17 QKeyCode = "f17"
	// QKeyCodeF18 since 8.0
	QKeyCodeF18 QKeyCode = "f18"
	// QKeyCodeF19 since 8.0
	QKeyCodeF19 QKeyCode = "f19"
	// QKeyCodeF20 since 8.0
	QKeyCodeF20 QKeyCode = "f20"
	// QKeyCodeF21 since 8.0
	QKeyCodeF21 QKeyCode = "f21"
	// QKeyCodeF22 since 8.0
	QKeyCodeF22 QKeyCode = "f22"
	// QKeyCodeF23 since 8.0
	QKeyCodeF23 QKeyCode = "f23"
	// QKeyCodeF24 since 8.0 'sysrq' was mistakenly added to hack around the fact that the ps2 driver was not generating correct scancodes sequences when 'alt+print' was pressed. This flaw is now fixed and the 'sysrq' key serves no further purpose. Any further use of 'sysrq' will be transparently changed to 'print', so they are effectively synonyms.
	QKeyCodeF24 QKeyCode = "f24"
)

// KeyValueKind
type KeyValueKind string

const (
	KeyValueKindNumber KeyValueKind = "number"
	KeyValueKindQcode  KeyValueKind = "qcode"
)

// IntWrapper
type IntWrapper struct {
	// Data a numeric key code
	Data int64 `json:"data"`
}

// QKeyCodeWrapper
type QKeyCodeWrapper struct {
	// Data An enumeration of key name
	Data QKeyCode `json:"data"`
}

// KeyValue
//
// Represents a keyboard key.
type KeyValue struct {
	// Discriminator: type

	// Type key encoding
	Type KeyValueKind `json:"type"`

	Number *IntWrapper      `json:"-"`
	Qcode  *QKeyCodeWrapper `json:"-"`
}

func (u KeyValue) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "number":
		if u.Number == nil {
			return nil, fmt.Errorf("expected Number to be set")
		}

		return json.Marshal(struct {
			Type KeyValueKind `json:"type"`
			*IntWrapper
		}{
			Type:       u.Type,
			IntWrapper: u.Number,
		})
	case "qcode":
		if u.Qcode == nil {
			return nil, fmt.Errorf("expected Qcode to be set")
		}

		return json.Marshal(struct {
			Type KeyValueKind `json:"type"`
			*QKeyCodeWrapper
		}{
			Type:            u.Type,
			QKeyCodeWrapper: u.Qcode,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// SendKey
//
// Send keys to guest.
type SendKey struct {
	// Keys An array of @KeyValue elements. All @KeyValues in this array are simultaneously sent to the guest. A @KeyValue.number value is sent directly to the guest, while @KeyValue.qcode must be a valid @QKeyCode value
	Keys []KeyValue `json:"keys"`
	// HoldTime time to delay key up events, milliseconds. Defaults to 100
	HoldTime *int64 `json:"hold-time,omitempty"`
}

func (SendKey) Command() string {
	return "send-key"
}

func (cmd SendKey) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "send-key", cmd, nil)
}

// InputButton Button of a pointer input device (mouse, tablet).
type InputButton string

const (
	InputButtonLeft      InputButton = "left"
	InputButtonMiddle    InputButton = "middle"
	InputButtonRight     InputButton = "right"
	InputButtonWheelUp   InputButton = "wheel-up"
	InputButtonWheelDown InputButton = "wheel-down"
	// InputButtonSide front side button of a 5-button mouse (since 2.9)
	InputButtonSide InputButton = "side"
	// InputButtonExtra rear side button of a 5-button mouse (since 2.9)
	InputButtonExtra      InputButton = "extra"
	InputButtonWheelLeft  InputButton = "wheel-left"
	InputButtonWheelRight InputButton = "wheel-right"
	// InputButtonTouch screen contact on a multi-touch device (since 8.1)
	InputButtonTouch InputButton = "touch"
)

// InputAxis Position axis of a pointer input device (mouse, tablet).
type InputAxis string

const (
	InputAxisX InputAxis = "x"
	InputAxisY InputAxis = "y"
)

// InputMultiTouchType Type of a multi-touch event.
type InputMultiTouchType string

const (
	InputMultiTouchTypeBegin  InputMultiTouchType = "begin"
	InputMultiTouchTypeUpdate InputMultiTouchType = "update"
	InputMultiTouchTypeEnd    InputMultiTouchType = "end"
	InputMultiTouchTypeCancel InputMultiTouchType = "cancel"
	InputMultiTouchTypeData   InputMultiTouchType = "data"
)

// InputKeyEvent
//
// Keyboard input event.
type InputKeyEvent struct {
	// Key Which key this event is for.
	Key KeyValue `json:"key"`
	// Down True for key-down and false for key-up events.
	Down bool `json:"down"`
}

// InputBtnEvent
//
// Pointer button input event.
type InputBtnEvent struct {
	// Button Which button this event is for.
	Button InputButton `json:"button"`
	// Down True for key-down and false for key-up events.
	Down bool `json:"down"`
}

// InputMoveEvent
//
// Pointer motion input event.
type InputMoveEvent struct {
	// Axis Which axis is referenced by @value.
	Axis InputAxis `json:"axis"`
	// Value Pointer position. For absolute coordinates the valid range is 0 -> 0x7ffff
	Value int64 `json:"value"`
}

// InputMultiTouchEvent
//
// MultiTouch input event.
type InputMultiTouchEvent struct {
	Type InputMultiTouchType `json:"type"`
	// Slot Which slot has generated the event.
	Slot int64 `json:"slot"`
	// TrackingId ID to correlate this event with previously generated events.
	TrackingId int64 `json:"tracking-id"`
	// Axis Which axis is referenced by @value.
	Axis InputAxis `json:"axis"`
	// Value Contact position.
	Value int64 `json:"value"`
}

// InputEventKind
type InputEventKind string

const (
	// InputEventKindKey a keyboard input event
	InputEventKindKey InputEventKind = "key"
	// InputEventKindBtn a pointer button input event
	InputEventKindBtn InputEventKind = "btn"
	// InputEventKindRel a relative pointer motion input event
	InputEventKindRel InputEventKind = "rel"
	// InputEventKindAbs an absolute pointer motion input event
	InputEventKindAbs InputEventKind = "abs"
	// InputEventKindMtt a multi-touch input event
	InputEventKindMtt InputEventKind = "mtt"
)

// InputKeyEventWrapper
type InputKeyEventWrapper struct {
	// Data Keyboard input event
	Data InputKeyEvent `json:"data"`
}

// InputBtnEventWrapper
type InputBtnEventWrapper struct {
	// Data Pointer button input event
	Data InputBtnEvent `json:"data"`
}

// InputMoveEventWrapper
type InputMoveEventWrapper struct {
	// Data Pointer motion input event
	Data InputMoveEvent `json:"data"`
}

// InputMultiTouchEventWrapper
type InputMultiTouchEventWrapper struct {
	// Data MultiTouch input event
	Data InputMultiTouchEvent `json:"data"`
}

// InputEvent
//
// Input event union.
type InputEvent struct {
	// Discriminator: type

	// Type the type of input event
	Type InputEventKind `json:"type"`

	Key *InputKeyEventWrapper        `json:"-"`
	Btn *InputBtnEventWrapper        `json:"-"`
	Rel *InputMoveEventWrapper       `json:"-"`
	Abs *InputMoveEventWrapper       `json:"-"`
	Mtt *InputMultiTouchEventWrapper `json:"-"`
}

func (u InputEvent) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "key":
		if u.Key == nil {
			return nil, fmt.Errorf("expected Key to be set")
		}

		return json.Marshal(struct {
			Type InputEventKind `json:"type"`
			*InputKeyEventWrapper
		}{
			Type:                 u.Type,
			InputKeyEventWrapper: u.Key,
		})
	case "btn":
		if u.Btn == nil {
			return nil, fmt.Errorf("expected Btn to be set")
		}

		return json.Marshal(struct {
			Type InputEventKind `json:"type"`
			*InputBtnEventWrapper
		}{
			Type:                 u.Type,
			InputBtnEventWrapper: u.Btn,
		})
	case "rel":
		if u.Rel == nil {
			return nil, fmt.Errorf("expected Rel to be set")
		}

		return json.Marshal(struct {
			Type InputEventKind `json:"type"`
			*InputMoveEventWrapper
		}{
			Type:                  u.Type,
			InputMoveEventWrapper: u.Rel,
		})
	case "abs":
		if u.Abs == nil {
			return nil, fmt.Errorf("expected Abs to be set")
		}

		return json.Marshal(struct {
			Type InputEventKind `json:"type"`
			*InputMoveEventWrapper
		}{
			Type:                  u.Type,
			InputMoveEventWrapper: u.Abs,
		})
	case "mtt":
		if u.Mtt == nil {
			return nil, fmt.Errorf("expected Mtt to be set")
		}

		return json.Marshal(struct {
			Type InputEventKind `json:"type"`
			*InputMultiTouchEventWrapper
		}{
			Type:                        u.Type,
			InputMultiTouchEventWrapper: u.Mtt,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// InputSendEvent
//
// Send input event(s) to guest. The @device and @head parameters can be used to send the input event to specific input devices in case (a) multiple input devices of the same kind are added to the virtual machine and (b) you have configured input routing (see docs/multiseat.txt) for those input devices. The parameters work exactly like the device and head properties of input devices. If @device is missing, only devices that have no input routing config are admissible. If @device is specified, both input devices with and without input routing config are admissible, but devices with input routing config take precedence.
type InputSendEvent struct {
	// Device display device to send event(s) to.
	Device *string `json:"device,omitempty"`
	// Head head to send event(s) to, in case the display device supports multiple scanouts.
	Head *int64 `json:"head,omitempty"`
	// Events List of InputEvent union.
	Events []InputEvent `json:"events"`
}

func (InputSendEvent) Command() string {
	return "input-send-event"
}

func (cmd InputSendEvent) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "input-send-event", cmd, nil)
}

// DisplayGTK
//
// GTK display options.
type DisplayGTK struct {
	// GrabOnHover Grab keyboard input on mouse hover.
	GrabOnHover *bool `json:"grab-on-hover,omitempty"`
	// ZoomToFit Zoom guest display to fit into the host window. When turned off the host window will be resized instead. In case the display device can notify the guest on window resizes (virtio-gpu) this will default to "on", assuming the guest will resize the display to match the window size then. Otherwise it defaults to "off". (Since 3.1)
	ZoomToFit *bool `json:"zoom-to-fit,omitempty"`
	// ShowTabs Display the tab bar for switching between the various graphical interfaces (e.g. VGA and virtual console character devices) by default. (Since 7.1)
	ShowTabs *bool `json:"show-tabs,omitempty"`
	// ShowMenubar Display the main window menubar. Defaults to "on". (Since 8.0)
	ShowMenubar *bool `json:"show-menubar,omitempty"`
}

// DisplayEGLHeadless
//
// EGL headless display options.
type DisplayEGLHeadless struct {
	// Rendernode Which DRM render node should be used. Default is the first available node on the host.
	Rendernode *string `json:"rendernode,omitempty"`
}

// DisplayDBus
//
// DBus display options.
type DisplayDBus struct {
	// Rendernode Which DRM render node should be used. Default is the first available node on the host.
	Rendernode *string `json:"rendernode,omitempty"`
	// Addr The D-Bus bus address (default to the session bus).
	Addr *string `json:"addr,omitempty"`
	// P2p Whether to use peer-to-peer connections (accepted through @add_client).
	P2p *bool `json:"p2p,omitempty"`
	// Audiodev Use the specified DBus audiodev to export audio.
	Audiodev *string `json:"audiodev,omitempty"`
}

// DisplayGLMode Display OpenGL mode.
type DisplayGLMode string

const (
	// DisplayGLModeOff Disable OpenGL (default).
	DisplayGLModeOff DisplayGLMode = "off"
	// DisplayGLModeOn Use OpenGL, pick context type automatically. Would better be named 'auto' but is called 'on' for backward compatibility with bool type.
	DisplayGLModeOn DisplayGLMode = "on"
	// DisplayGLModeCore Use OpenGL with Core (desktop) Context.
	DisplayGLModeCore DisplayGLMode = "core"
	// DisplayGLModeEs Use OpenGL with ES (embedded systems) Context.
	DisplayGLModeEs DisplayGLMode = "es"
)

// DisplayCurses
//
// Curses display options.
type DisplayCurses struct {
	// Charset Font charset used by guest (default: CP437).
	Charset *string `json:"charset,omitempty"`
}

// DisplayCocoa
//
// Cocoa display options.
type DisplayCocoa struct {
	// LeftCommandKey Enable/disable forwarding of left command key to guest. Allows command-tab window switching on the host without sending this key to the guest when "off". Defaults to "on"
	LeftCommandKey *bool `json:"left-command-key,omitempty"`
	// FullGrab Capture all key presses, including system combos. This requires accessibility permissions, since it performs a global
	FullGrab *bool `json:"full-grab,omitempty"`
	// SwapOptCmd Swap the Option and Command keys so that their key codes match their position on non-Mac keyboards and you can use
	SwapOptCmd *bool `json:"swap-opt-cmd,omitempty"`
	// ZoomToFit Zoom guest display to fit into the host window. When turned off the host window will be resized instead. Defaults to "off". (Since 8.2)
	ZoomToFit *bool `json:"zoom-to-fit,omitempty"`
}

// HotKeyMod Set of modifier keys that need to be held for shortcut key actions.
type HotKeyMod string

const (
	HotKeyModLctrlLalt       HotKeyMod = "lctrl-lalt"
	HotKeyModLshiftLctrlLalt HotKeyMod = "lshift-lctrl-lalt"
	HotKeyModRctrl           HotKeyMod = "rctrl"
)

// DisplaySDL
//
// SDL2 display options.
type DisplaySDL struct {
	// GrabMod Modifier keys that should be pressed together with the "G" key to release the mouse grab.
	GrabMod *HotKeyMod `json:"grab-mod,omitempty"`
}

// DisplayType Display (user interface) type.
type DisplayType string

const (
	// DisplayTypeDefault The default user interface, selecting from the first available of gtk, sdl, cocoa, and vnc.
	DisplayTypeDefault DisplayType = "default"
	// DisplayTypeNone No user interface or video output display. The guest will still see an emulated graphics card, but its output will not be displayed to the QEMU user.
	DisplayTypeNone DisplayType = "none"
	// DisplayTypeGtk The GTK user interface.
	DisplayTypeGtk DisplayType = "gtk"
	// DisplayTypeSdl The SDL user interface.
	DisplayTypeSdl DisplayType = "sdl"
	// DisplayTypeEglHeadless No user interface, offload GL operations to a local DRI device. Graphical display need to be paired with VNC or Spice. (Since 3.1)
	DisplayTypeEglHeadless DisplayType = "egl-headless"
	// DisplayTypeCurses Display video output via curses. For graphics device models which support a text mode, QEMU can display this output using a curses/ncurses interface. Nothing is displayed when the graphics device is in graphical mode or if the graphics device does not support a text mode. Generally only the VGA device models support text mode.
	DisplayTypeCurses DisplayType = "curses"
	// DisplayTypeCocoa The Cocoa user interface.
	DisplayTypeCocoa DisplayType = "cocoa"
	// DisplayTypeSpiceApp Set up a Spice server and run the default associated application to connect to it. The server will redirect the serial console and QEMU monitors. (Since 4.0)
	DisplayTypeSpiceApp DisplayType = "spice-app"
	// DisplayTypeDbus Start a D-Bus service for the display. (Since 7.0)
	DisplayTypeDbus DisplayType = "dbus"
)

// DisplayOptions
//
// Display (user interface) options.
type DisplayOptions struct {
	// Discriminator: type

	// Type Which DisplayType qemu should use.
	Type DisplayType `json:"type"`
	// FullScreen Start user interface in fullscreen mode
	FullScreen *bool `json:"full-screen,omitempty"`
	// WindowClose Allow to quit qemu with window close button
	WindowClose *bool `json:"window-close,omitempty"`
	// ShowCursor Force showing the mouse cursor (default: off).
	ShowCursor *bool `json:"show-cursor,omitempty"`
	// Gl Enable OpenGL support (default: off).
	Gl *DisplayGLMode `json:"gl,omitempty"`

	Gtk         *DisplayGTK         `json:"-"`
	Cocoa       *DisplayCocoa       `json:"-"`
	Curses      *DisplayCurses      `json:"-"`
	EglHeadless *DisplayEGLHeadless `json:"-"`
	Dbus        *DisplayDBus        `json:"-"`
	Sdl         *DisplaySDL         `json:"-"`
}

func (u DisplayOptions) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "gtk":
		if u.Gtk == nil {
			return nil, fmt.Errorf("expected Gtk to be set")
		}

		return json.Marshal(struct {
			Type        DisplayType    `json:"type"`
			FullScreen  *bool          `json:"full-screen,omitempty"`
			WindowClose *bool          `json:"window-close,omitempty"`
			ShowCursor  *bool          `json:"show-cursor,omitempty"`
			Gl          *DisplayGLMode `json:"gl,omitempty"`
			*DisplayGTK
		}{
			Type:        u.Type,
			FullScreen:  u.FullScreen,
			WindowClose: u.WindowClose,
			ShowCursor:  u.ShowCursor,
			Gl:          u.Gl,
			DisplayGTK:  u.Gtk,
		})
	case "cocoa":
		if u.Cocoa == nil {
			return nil, fmt.Errorf("expected Cocoa to be set")
		}

		return json.Marshal(struct {
			Type        DisplayType    `json:"type"`
			FullScreen  *bool          `json:"full-screen,omitempty"`
			WindowClose *bool          `json:"window-close,omitempty"`
			ShowCursor  *bool          `json:"show-cursor,omitempty"`
			Gl          *DisplayGLMode `json:"gl,omitempty"`
			*DisplayCocoa
		}{
			Type:         u.Type,
			FullScreen:   u.FullScreen,
			WindowClose:  u.WindowClose,
			ShowCursor:   u.ShowCursor,
			Gl:           u.Gl,
			DisplayCocoa: u.Cocoa,
		})
	case "curses":
		if u.Curses == nil {
			return nil, fmt.Errorf("expected Curses to be set")
		}

		return json.Marshal(struct {
			Type        DisplayType    `json:"type"`
			FullScreen  *bool          `json:"full-screen,omitempty"`
			WindowClose *bool          `json:"window-close,omitempty"`
			ShowCursor  *bool          `json:"show-cursor,omitempty"`
			Gl          *DisplayGLMode `json:"gl,omitempty"`
			*DisplayCurses
		}{
			Type:          u.Type,
			FullScreen:    u.FullScreen,
			WindowClose:   u.WindowClose,
			ShowCursor:    u.ShowCursor,
			Gl:            u.Gl,
			DisplayCurses: u.Curses,
		})
	case "egl-headless":
		if u.EglHeadless == nil {
			return nil, fmt.Errorf("expected EglHeadless to be set")
		}

		return json.Marshal(struct {
			Type        DisplayType    `json:"type"`
			FullScreen  *bool          `json:"full-screen,omitempty"`
			WindowClose *bool          `json:"window-close,omitempty"`
			ShowCursor  *bool          `json:"show-cursor,omitempty"`
			Gl          *DisplayGLMode `json:"gl,omitempty"`
			*DisplayEGLHeadless
		}{
			Type:               u.Type,
			FullScreen:         u.FullScreen,
			WindowClose:        u.WindowClose,
			ShowCursor:         u.ShowCursor,
			Gl:                 u.Gl,
			DisplayEGLHeadless: u.EglHeadless,
		})
	case "dbus":
		if u.Dbus == nil {
			return nil, fmt.Errorf("expected Dbus to be set")
		}

		return json.Marshal(struct {
			Type        DisplayType    `json:"type"`
			FullScreen  *bool          `json:"full-screen,omitempty"`
			WindowClose *bool          `json:"window-close,omitempty"`
			ShowCursor  *bool          `json:"show-cursor,omitempty"`
			Gl          *DisplayGLMode `json:"gl,omitempty"`
			*DisplayDBus
		}{
			Type:        u.Type,
			FullScreen:  u.FullScreen,
			WindowClose: u.WindowClose,
			ShowCursor:  u.ShowCursor,
			Gl:          u.Gl,
			DisplayDBus: u.Dbus,
		})
	case "sdl":
		if u.Sdl == nil {
			return nil, fmt.Errorf("expected Sdl to be set")
		}

		return json.Marshal(struct {
			Type        DisplayType    `json:"type"`
			FullScreen  *bool          `json:"full-screen,omitempty"`
			WindowClose *bool          `json:"window-close,omitempty"`
			ShowCursor  *bool          `json:"show-cursor,omitempty"`
			Gl          *DisplayGLMode `json:"gl,omitempty"`
			*DisplaySDL
		}{
			Type:        u.Type,
			FullScreen:  u.FullScreen,
			WindowClose: u.WindowClose,
			ShowCursor:  u.ShowCursor,
			Gl:          u.Gl,
			DisplaySDL:  u.Sdl,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// QueryDisplayOptions
//
// Returns information about display configuration
type QueryDisplayOptions struct {
}

func (QueryDisplayOptions) Command() string {
	return "query-display-options"
}

func (cmd QueryDisplayOptions) Execute(ctx context.Context, client api.Client) (DisplayOptions, error) {
	var ret DisplayOptions

	return ret, client.Execute(ctx, "query-display-options", cmd, &ret)
}

// DisplayReloadType Available DisplayReload types.
type DisplayReloadType string

const (
	// DisplayReloadTypeVnc VNC display
	DisplayReloadTypeVnc DisplayReloadType = "vnc"
)

// DisplayReloadOptionsVNC
//
// Specify the VNC reload options.
type DisplayReloadOptionsVNC struct {
	// TlsCerts reload tls certs or not.
	TlsCerts *bool `json:"tls-certs,omitempty"`
}

// DisplayReloadOptions
//
// Options of the display configuration reload.
type DisplayReloadOptions struct {
	// Discriminator: type

	// Type Specify the display type.
	Type DisplayReloadType `json:"type"`

	Vnc *DisplayReloadOptionsVNC `json:"-"`
}

func (u DisplayReloadOptions) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "vnc":
		if u.Vnc == nil {
			return nil, fmt.Errorf("expected Vnc to be set")
		}

		return json.Marshal(struct {
			Type DisplayReloadType `json:"type"`
			*DisplayReloadOptionsVNC
		}{
			Type:                    u.Type,
			DisplayReloadOptionsVNC: u.Vnc,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// DisplayReload
//
// Reload display configuration.
type DisplayReload struct {
	DisplayReloadOptions
}

func (DisplayReload) Command() string {
	return "display-reload"
}

func (cmd DisplayReload) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "display-reload", cmd, nil)
}

// DisplayUpdateType Available DisplayUpdate types.
type DisplayUpdateType string

const (
	// DisplayUpdateTypeVnc VNC display
	DisplayUpdateTypeVnc DisplayUpdateType = "vnc"
)

// DisplayUpdateOptionsVNC
//
// Specify the VNC reload options.
type DisplayUpdateOptionsVNC struct {
	// Addresses If specified, change set of addresses to listen for connections. Addresses configured for websockets are not touched.
	Addresses []SocketAddress `json:"addresses,omitempty"`
}

// DisplayUpdateOptions
//
// Options of the display configuration reload.
type DisplayUpdateOptions struct {
	// Discriminator: type

	// Type Specify the display type.
	Type DisplayUpdateType `json:"type"`

	Vnc *DisplayUpdateOptionsVNC `json:"-"`
}

func (u DisplayUpdateOptions) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "vnc":
		if u.Vnc == nil {
			return nil, fmt.Errorf("expected Vnc to be set")
		}

		return json.Marshal(struct {
			Type DisplayUpdateType `json:"type"`
			*DisplayUpdateOptionsVNC
		}{
			Type:                    u.Type,
			DisplayUpdateOptionsVNC: u.Vnc,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// DisplayUpdate
//
// Update display configuration.
type DisplayUpdate struct {
	DisplayUpdateOptions
}

func (DisplayUpdate) Command() string {
	return "display-update"
}

func (cmd DisplayUpdate) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "display-update", cmd, nil)
}

// ClientMigrateInfo
//
// Set migration information for remote display. This makes the server ask the client to automatically reconnect using the new parameters once migration finished successfully. Only implemented for SPICE.
type ClientMigrateInfo struct {
	// Protocol must be "spice"
	Protocol string `json:"protocol"`
	// Hostname migration target hostname
	Hostname string `json:"hostname"`
	// Port spice tcp port for plaintext channels
	Port *int64 `json:"port,omitempty"`
	// TlsPort spice tcp port for tls-secured channels
	TlsPort *int64 `json:"tls-port,omitempty"`
	// CertSubject server certificate subject
	CertSubject *string `json:"cert-subject,omitempty"`
}

func (ClientMigrateInfo) Command() string {
	return "client_migrate_info"
}

func (cmd ClientMigrateInfo) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "client_migrate_info", cmd, nil)
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

// MigrationStats
//
// Detailed migration status.
type MigrationStats struct {
	// Transferred amount of bytes already transferred to the target VM
	Transferred int64 `json:"transferred"`
	// Remaining amount of bytes remaining to be transferred to the target VM
	Remaining int64 `json:"remaining"`
	// Total total amount of bytes involved in the migration process
	Total int64 `json:"total"`
	// Duplicate number of duplicate (zero) pages (since 1.2)
	Duplicate int64 `json:"duplicate"`
	// Skipped number of skipped zero pages. Always zero, only provided for compatibility (since 1.5)
	Skipped int64 `json:"skipped"`
	// Normal number of normal pages (since 1.2)
	Normal int64 `json:"normal"`
	// NormalBytes number of normal bytes sent (since 1.2)
	NormalBytes int64 `json:"normal-bytes"`
	// DirtyPagesRate number of pages dirtied by second by the guest (since 1.3)
	DirtyPagesRate int64 `json:"dirty-pages-rate"`
	// Mbps throughput in megabits/sec. (since 1.6)
	Mbps float64 `json:"mbps"`
	// DirtySyncCount number of times that dirty ram was synchronized (since 2.1)
	DirtySyncCount int64 `json:"dirty-sync-count"`
	// PostcopyRequests The number of page requests received from the destination (since 2.7)
	PostcopyRequests int64 `json:"postcopy-requests"`
	// PageSize The number of bytes per page for the various page-based statistics (since 2.10)
	PageSize int64 `json:"page-size"`
	// MultifdBytes The number of bytes sent through multifd (since 3.0)
	MultifdBytes uint64 `json:"multifd-bytes"`
	// PagesPerSecond the number of memory pages transferred per second (Since 4.0)
	PagesPerSecond uint64 `json:"pages-per-second"`
	// PrecopyBytes The number of bytes sent in the pre-copy phase (since 7.0).
	PrecopyBytes uint64 `json:"precopy-bytes"`
	// DowntimeBytes The number of bytes sent while the guest is paused (since 7.0).
	DowntimeBytes uint64 `json:"downtime-bytes"`
	// PostcopyBytes The number of bytes sent during the post-copy phase (since 7.0).
	PostcopyBytes uint64 `json:"postcopy-bytes"`
	// DirtySyncMissedZeroCopy Number of times dirty RAM synchronization could not avoid copying dirty pages. This is between 0 and @dirty-sync-count * @multifd-channels. (since 7.1)
	DirtySyncMissedZeroCopy uint64 `json:"dirty-sync-missed-zero-copy"`
}

// XBZRLECacheStats
//
// Detailed XBZRLE migration cache statistics
type XBZRLECacheStats struct {
	// CacheSize XBZRLE cache size
	CacheSize uint64 `json:"cache-size"`
	// Bytes amount of bytes already transferred to the target VM
	Bytes int64 `json:"bytes"`
	// Pages amount of pages transferred to the target VM
	Pages int64 `json:"pages"`
	// CacheMiss number of cache miss
	CacheMiss int64 `json:"cache-miss"`
	// CacheMissRate rate of cache miss (since 2.1)
	CacheMissRate float64 `json:"cache-miss-rate"`
	// EncodingRate rate of encoded bytes (since 5.1)
	EncodingRate float64 `json:"encoding-rate"`
	// Overflow number of overflows
	Overflow int64 `json:"overflow"`
}

// CompressionStats
//
// Detailed migration compression statistics
type CompressionStats struct {
	// Pages amount of pages compressed and transferred to the target VM
	Pages int64 `json:"pages"`
	// Busy count of times that no free thread was available to compress data
	Busy int64 `json:"busy"`
	// BusyRate rate of thread busy
	BusyRate float64 `json:"busy-rate"`
	// CompressedSize amount of bytes after compression
	CompressedSize int64 `json:"compressed-size"`
	// CompressionRate rate of compressed size
	CompressionRate float64 `json:"compression-rate"`
}

// MigrationStatus An enumeration of migration status.
type MigrationStatus string

const (
	// MigrationStatusNone no migration has ever happened.
	MigrationStatusNone MigrationStatus = "none"
	// MigrationStatusSetup migration process has been initiated.
	MigrationStatusSetup MigrationStatus = "setup"
	// MigrationStatusCancelling in the process of cancelling migration.
	MigrationStatusCancelling MigrationStatus = "cancelling"
	// MigrationStatusCancelled cancelling migration is finished.
	MigrationStatusCancelled MigrationStatus = "cancelled"
	// MigrationStatusActive in the process of doing migration.
	MigrationStatusActive MigrationStatus = "active"
	// MigrationStatusPostcopyActive like active, but now in postcopy mode. (since 2.5)
	MigrationStatusPostcopyActive MigrationStatus = "postcopy-active"
	// MigrationStatusPostcopyPaused during postcopy but paused. (since 3.0)
	MigrationStatusPostcopyPaused MigrationStatus = "postcopy-paused"
	// MigrationStatusPostcopyRecover trying to recover from a paused postcopy. (since 3.0)
	MigrationStatusPostcopyRecover MigrationStatus = "postcopy-recover"
	// MigrationStatusCompleted migration is finished.
	MigrationStatusCompleted MigrationStatus = "completed"
	// MigrationStatusFailed some error occurred during migration process.
	MigrationStatusFailed MigrationStatus = "failed"
	// MigrationStatusColo VM is in the process of fault tolerance, VM can not get into this state unless colo capability is enabled for migration. (since 2.8)
	MigrationStatusColo MigrationStatus = "colo"
	// MigrationStatusPreSwitchover Paused before device serialisation. (since 2.11)
	MigrationStatusPreSwitchover MigrationStatus = "pre-switchover"
	// MigrationStatusDevice During device serialisation when pause-before-switchover is enabled (since 2.11)
	MigrationStatusDevice MigrationStatus = "device"
	// MigrationStatusWaitUnplug wait for device unplug request by guest OS to be completed. (since 4.2)
	MigrationStatusWaitUnplug MigrationStatus = "wait-unplug"
)

// VfioStats
//
// Detailed VFIO devices migration statistics
type VfioStats struct {
	// Transferred amount of bytes transferred to the target VM by VFIO devices
	Transferred int64 `json:"transferred"`
}

// MigrationInfo
//
// Information about current migration process.
type MigrationInfo struct {
	// Status @MigrationStatus describing the current migration status. If this field is not returned, no migration process has been initiated
	Status *MigrationStatus `json:"status,omitempty"`
	// Ram @MigrationStats containing detailed migration status, only returned if status is 'active' or 'completed'(since 1.2)
	Ram *MigrationStats `json:"ram,omitempty"`
	// Disk @MigrationStats containing detailed disk migration status, only returned if status is 'active' and it is a block migration
	Disk *MigrationStats `json:"disk,omitempty"`
	// Vfio @VfioStats containing detailed VFIO devices migration statistics, only returned if VFIO device is present, migration is supported by all VFIO devices and status is 'active' or 'completed' (since 5.2)
	Vfio *VfioStats `json:"vfio,omitempty"`
	// XbzrleCache @XBZRLECacheStats containing detailed XBZRLE migration statistics, only returned if XBZRLE feature is on and status is 'active' or 'completed' (since 1.2)
	XbzrleCache *XBZRLECacheStats `json:"xbzrle-cache,omitempty"`
	// TotalTime total amount of milliseconds since migration started. If migration has ended, it returns the total migration time. (since 1.2)
	TotalTime *int64 `json:"total-time,omitempty"`
	// ExpectedDowntime only present while migration is active expected downtime in milliseconds for the guest in last walk of the dirty bitmap. (since 1.3)
	ExpectedDowntime *int64 `json:"expected-downtime,omitempty"`
	// Downtime only present when migration finishes correctly total downtime in milliseconds for the guest. (since 1.3)
	Downtime *int64 `json:"downtime,omitempty"`
	// SetupTime amount of setup time in milliseconds *before* the iterations begin but *after* the QMP command is issued. This is designed to provide an accounting of any activities (such as RDMA pinning) which may be expensive, but do not actually occur during the iterative migration rounds themselves. (since 1.6)
	SetupTime *int64 `json:"setup-time,omitempty"`
	// CpuThrottlePercentage percentage of time guest cpus are being throttled during auto-converge. This is only present when auto-converge has started throttling guest cpus. (Since 2.7)
	CpuThrottlePercentage *int64 `json:"cpu-throttle-percentage,omitempty"`
	// ErrorDesc the human readable error description string. Clients should not attempt to parse the error strings. (Since 2.7)
	ErrorDesc *string `json:"error-desc,omitempty"`
	// BlockedReasons A list of reasons an outgoing migration is blocked. Present and non-empty when migration is blocked. (since 6.0)
	BlockedReasons []string `json:"blocked-reasons,omitempty"`
	// PostcopyBlocktime total time when all vCPU were blocked during postcopy live migration. This is only present when the postcopy-blocktime migration capability is enabled. (Since 3.0)
	PostcopyBlocktime *uint32 `json:"postcopy-blocktime,omitempty"`
	// PostcopyVcpuBlocktime list of the postcopy blocktime per vCPU. This is only present when the postcopy-blocktime migration capability is enabled. (Since 3.0)
	PostcopyVcpuBlocktime []uint32 `json:"postcopy-vcpu-blocktime,omitempty"`
	// Compression migration compression statistics, only returned if compression feature is on and status is 'active' or 'completed' (Since 3.1)
	Compression *CompressionStats `json:"compression,omitempty"`
	// SocketAddress Only used for tcp, to know what the real port is (Since 4.0)
	SocketAddress []SocketAddress `json:"socket-address,omitempty"`
	// DirtyLimitThrottleTimePerRound Maximum throttle time (in microseconds) of virtual CPUs each dirty ring full round, which shows how MigrationCapability dirty-limit affects the guest during live migration. (Since 8.1)
	DirtyLimitThrottleTimePerRound *uint64 `json:"dirty-limit-throttle-time-per-round,omitempty"`
	// DirtyLimitRingFullTime Estimated average dirty ring full time (in microseconds) for each dirty ring full round. The value equals the dirty ring memory size divided by the average dirty page rate of the virtual CPU, which can be used to observe the average memory load of the virtual CPU indirectly. Note that zero means guest doesn't dirty memory. (Since 8.1)
	DirtyLimitRingFullTime *uint64 `json:"dirty-limit-ring-full-time,omitempty"`
}

// QueryMigrate
//
// Returns information about current migration process. If migration is active there will be another json-object with RAM migration status and if block migration is active another one with block migration status.
type QueryMigrate struct {
}

func (QueryMigrate) Command() string {
	return "query-migrate"
}

func (cmd QueryMigrate) Execute(ctx context.Context, client api.Client) (MigrationInfo, error) {
	var ret MigrationInfo

	return ret, client.Execute(ctx, "query-migrate", cmd, &ret)
}

// MigrationCapability Migration capabilities enumeration
type MigrationCapability string

const (
	// MigrationCapabilityXbzrle Migration supports xbzrle (Xor Based Zero Run Length Encoding). This feature allows us to minimize migration traffic for certain work loads, by sending compressed difference of the pages
	MigrationCapabilityXbzrle MigrationCapability = "xbzrle"
	// MigrationCapabilityRdmaPinAll Controls whether or not the entire VM memory footprint is mlock()'d on demand or all at once. Refer to docs/rdma.txt for usage. Disabled by default. (since 2.0)
	MigrationCapabilityRdmaPinAll MigrationCapability = "rdma-pin-all"
	// MigrationCapabilityAutoConverge If enabled, QEMU will automatically throttle down the guest to speed up convergence of RAM migration. (since 1.6)
	MigrationCapabilityAutoConverge MigrationCapability = "auto-converge"
	// MigrationCapabilityZeroBlocks During storage migration encode blocks of zeroes efficiently. This essentially saves 1MB of zeroes per block on the wire. Enabling requires source and target VM to support this feature. To enable it is sufficient to enable the capability on the source VM. The feature is disabled by default. (since 1.6)
	MigrationCapabilityZeroBlocks MigrationCapability = "zero-blocks"
	// MigrationCapabilityCompress Use multiple compression threads to accelerate live migration. This feature can help to reduce the migration traffic, by sending compressed pages. Please note that if compress and xbzrle are both on, compress only takes effect in the ram bulk stage, after that, it will be disabled and only xbzrle takes effect, this can help to minimize migration traffic. The feature is disabled by default. (since 2.4)
	MigrationCapabilityCompress MigrationCapability = "compress"
	// MigrationCapabilityEvents generate events for each migration state change (since 2.4)
	MigrationCapabilityEvents MigrationCapability = "events"
	// MigrationCapabilityPostcopyRam Start executing on the migration target before all of RAM has been migrated, pulling the remaining pages along as needed. The capacity must have the same setting on both source
	MigrationCapabilityPostcopyRam MigrationCapability = "postcopy-ram"
	// MigrationCapabilityColo If enabled, migration will never end, and the state of the VM on the primary side will be migrated continuously to the VM on secondary side, this process is called COarse-Grain LOck Stepping (COLO) for Non-stop Service. (since 2.8)
	MigrationCapabilityColo MigrationCapability = "x-colo"
	// MigrationCapabilityReleaseRam if enabled, qemu will free the migrated ram pages on the source during postcopy-ram migration. (since 2.9)
	MigrationCapabilityReleaseRam MigrationCapability = "release-ram"
	// MigrationCapabilityBlock If enabled, QEMU will also migrate the contents of all block devices. Default is disabled. A possible alternative uses mirror jobs to a builtin NBD server on the destination, which offers more flexibility. (Since 2.10)
	MigrationCapabilityBlock MigrationCapability = "block"
	// MigrationCapabilityReturnPath If enabled, migration will use the return path even for precopy. (since 2.10)
	MigrationCapabilityReturnPath MigrationCapability = "return-path"
	// MigrationCapabilityPauseBeforeSwitchover Pause outgoing migration before serialising device state and before disabling block IO (since 2.11)
	MigrationCapabilityPauseBeforeSwitchover MigrationCapability = "pause-before-switchover"
	// MigrationCapabilityMultifd Use more than one fd for migration (since 4.0)
	MigrationCapabilityMultifd MigrationCapability = "multifd"
	// MigrationCapabilityDirtyBitmaps If enabled, QEMU will migrate named dirty bitmaps. (since 2.12)
	MigrationCapabilityDirtyBitmaps MigrationCapability = "dirty-bitmaps"
	// MigrationCapabilityPostcopyBlocktime Calculate downtime for postcopy live migration (since 3.0)
	MigrationCapabilityPostcopyBlocktime MigrationCapability = "postcopy-blocktime"
	// MigrationCapabilityLateBlockActivate If enabled, the destination will not activate block devices (and thus take locks) immediately at the end of migration. (since 3.0)
	MigrationCapabilityLateBlockActivate MigrationCapability = "late-block-activate"
	// MigrationCapabilityIgnoreShared If enabled, QEMU will not migrate shared memory that is accessible on the destination machine. (since 4.0)
	MigrationCapabilityIgnoreShared MigrationCapability = "x-ignore-shared"
	// MigrationCapabilityValidateUuid Send the UUID of the source to allow the destination to ensure it is the same. (since 4.2)
	MigrationCapabilityValidateUuid MigrationCapability = "validate-uuid"
	// MigrationCapabilityBackgroundSnapshot If enabled, the migration stream will be a snapshot of the VM exactly at the point when the migration procedure starts. The VM RAM is saved with running VM. (since 6.0)
	MigrationCapabilityBackgroundSnapshot MigrationCapability = "background-snapshot"
	// MigrationCapabilityZeroCopySend Controls behavior on sending memory pages on migration. When true, enables a zero-copy mechanism for sending memory pages, if host supports it. Requires that QEMU be permitted to use locked memory for guest RAM pages. (since 7.1)
	MigrationCapabilityZeroCopySend MigrationCapability = "zero-copy-send"
	// MigrationCapabilityPostcopyPreempt If enabled, the migration process will allow postcopy requests to preempt precopy stream, so postcopy requests will be handled faster. This is a performance feature and should not affect the correctness of postcopy migration. (since 7.1)
	MigrationCapabilityPostcopyPreempt MigrationCapability = "postcopy-preempt"
	// MigrationCapabilitySwitchoverAck If enabled, migration will not stop the source VM and complete the migration until an ACK is received from the destination that it's OK to do so. Exactly when this ACK is sent depends on the migrated devices that use this feature. For example, a device can use it to make sure some of its data is sent and loaded in the destination before doing switchover. This can reduce downtime if devices that support this capability are present. 'return-path' capability must be enabled to use it. (since 8.1)
	MigrationCapabilitySwitchoverAck MigrationCapability = "switchover-ack"
	// MigrationCapabilityDirtyLimit If enabled, migration will throttle vCPUs as needed to keep their dirty page rate within @vcpu-dirty-limit. This can improve responsiveness of large guests during live migration, and can result in more stable read performance. Requires KVM with accelerator property "dirty-ring-size" set. (Since 8.1)
	MigrationCapabilityDirtyLimit MigrationCapability = "dirty-limit"
)

// MigrationCapabilityStatus
//
// Migration capability information
type MigrationCapabilityStatus struct {
	// Capability capability enum
	Capability MigrationCapability `json:"capability"`
	// State capability state bool
	State bool `json:"state"`
}

// MigrateSetCapabilities
//
// Enable/Disable the following migration capabilities (like xbzrle)
type MigrateSetCapabilities struct {
	// Capabilities json array of capability modifications to make
	Capabilities []MigrationCapabilityStatus `json:"capabilities"`
}

func (MigrateSetCapabilities) Command() string {
	return "migrate-set-capabilities"
}

func (cmd MigrateSetCapabilities) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "migrate-set-capabilities", cmd, nil)
}

// QueryMigrateCapabilities
//
// Returns information about the current migration capabilities status
type QueryMigrateCapabilities struct {
}

func (QueryMigrateCapabilities) Command() string {
	return "query-migrate-capabilities"
}

func (cmd QueryMigrateCapabilities) Execute(ctx context.Context, client api.Client) ([]MigrationCapabilityStatus, error) {
	var ret []MigrationCapabilityStatus

	return ret, client.Execute(ctx, "query-migrate-capabilities", cmd, &ret)
}

// MultiFDCompression An enumeration of multifd compression methods.
type MultiFDCompression string

const (
	// MultiFDCompressionNone no compression.
	MultiFDCompressionNone MultiFDCompression = "none"
	// MultiFDCompressionZlib use zlib compression method.
	MultiFDCompressionZlib MultiFDCompression = "zlib"
	// MultiFDCompressionZstd use zstd compression method.
	MultiFDCompressionZstd MultiFDCompression = "zstd"
)

// MigMode
type MigMode string

const (
	// MigModeNormal the original form of migration. (since 8.2)
	MigModeNormal MigMode = "normal"
	// MigModeCprReboot The migrate command saves state to a file, allowing one to quit qemu, reboot to an updated kernel, and restart an updated version of qemu. The caller must specify a migration URI that writes to and reads from a file. Unlike normal mode, the use of certain local storage options does not block the migration, but the caller must not modify guest block devices between the quit and restart. To avoid saving guest RAM to the file, the memory backend must be shared, and the @x-ignore-shared migration capability must be set. Guest RAM must be non-volatile across reboot, such as by backing it with a dax device, but this is not enforced. The restarted qemu arguments must match those used to initially start qemu, plus the -incoming option. (since 8.2)
	MigModeCprReboot MigMode = "cpr-reboot"
)

// BitmapMigrationBitmapAliasTransform
type BitmapMigrationBitmapAliasTransform struct {
	// Persistent If present, the bitmap will be made persistent or transient depending on this parameter.
	Persistent *bool `json:"persistent,omitempty"`
}

// BitmapMigrationBitmapAlias
type BitmapMigrationBitmapAlias struct {
	// Name The name of the bitmap.
	Name string `json:"name"`
	// Alias An alias name for migration (for example the bitmap name on the opposite site).
	Alias string `json:"alias"`
	// Transform Allows the modification of the migrated bitmap. (since 6.0)
	Transform *BitmapMigrationBitmapAliasTransform `json:"transform,omitempty"`
}

// BitmapMigrationNodeAlias
//
// Maps a block node name and the bitmaps it has to aliases for dirty bitmap migration.
type BitmapMigrationNodeAlias struct {
	// NodeName A block node name.
	NodeName string `json:"node-name"`
	// Alias An alias block node name for migration (for example the node name on the opposite site).
	Alias string `json:"alias"`
	// Bitmaps Mappings for the bitmaps on this node.
	Bitmaps []BitmapMigrationBitmapAlias `json:"bitmaps"`
}

// MigrationParameter Migration parameters enumeration
type MigrationParameter string

const (
	// MigrationParameterAnnounceInitial Initial delay (in milliseconds) before sending the first announce (Since 4.0)
	MigrationParameterAnnounceInitial MigrationParameter = "announce-initial"
	// MigrationParameterAnnounceMax Maximum delay (in milliseconds) between packets in the announcement (Since 4.0)
	MigrationParameterAnnounceMax MigrationParameter = "announce-max"
	// MigrationParameterAnnounceRounds Number of self-announce packets sent after migration (Since 4.0)
	MigrationParameterAnnounceRounds MigrationParameter = "announce-rounds"
	// MigrationParameterAnnounceStep Increase in delay (in milliseconds) between subsequent packets in the announcement (Since 4.0)
	MigrationParameterAnnounceStep MigrationParameter = "announce-step"
	// MigrationParameterCompressLevel Set the compression level to be used in live migration, the compression level is an integer between 0 and 9, where 0 means no compression, 1 means the best compression speed, and 9 means best compression ratio which will consume more CPU.
	MigrationParameterCompressLevel MigrationParameter = "compress-level"
	// MigrationParameterCompressThreads Set compression thread count to be used in live migration, the compression thread count is an integer between 1 and 255.
	MigrationParameterCompressThreads MigrationParameter = "compress-threads"
	// MigrationParameterDecompressThreads Set decompression thread count to be used in live migration, the decompression thread count is an integer between 1 and 255. Usually, decompression is at least 4 times as fast as compression, so set the decompress-threads to the number about 1/4 of compress-threads is adequate.
	MigrationParameterDecompressThreads MigrationParameter = "decompress-threads"
	// MigrationParameterCompressWaitThread Controls behavior when all compression threads are currently busy. If true (default), wait for a free compression thread to become available; otherwise, send the page uncompressed. (Since 3.1)
	MigrationParameterCompressWaitThread MigrationParameter = "compress-wait-thread"
	// MigrationParameterThrottleTriggerThreshold The ratio of bytes_dirty_period and bytes_xfer_period to trigger throttling. It is expressed as percentage. The default value is 50. (Since 5.0)
	MigrationParameterThrottleTriggerThreshold MigrationParameter = "throttle-trigger-threshold"
	// MigrationParameterCpuThrottleInitial Initial percentage of time guest cpus are throttled when migration auto-converge is activated. The default value is 20. (Since 2.7)
	MigrationParameterCpuThrottleInitial MigrationParameter = "cpu-throttle-initial"
	// MigrationParameterCpuThrottleIncrement throttle percentage increase each time auto-converge detects that migration is not making progress. The default value is 10. (Since 2.7)
	MigrationParameterCpuThrottleIncrement MigrationParameter = "cpu-throttle-increment"
	// MigrationParameterCpuThrottleTailslow Make CPU throttling slower at tail stage At the tail stage of throttling, the Guest is very sensitive to CPU percentage while the @cpu-throttle -increment is excessive usually at tail stage. If this parameter is true, we will compute the ideal CPU percentage used by the Guest, which may exactly make the dirty rate match the dirty rate threshold. Then we will choose a smaller throttle increment between the one specified by @cpu-throttle-increment and the one generated by ideal CPU percentage. Therefore, it is compatible to traditional throttling, meanwhile the throttle increment won't be excessive at tail stage. The default value is false. (Since 5.1)
	MigrationParameterCpuThrottleTailslow MigrationParameter = "cpu-throttle-tailslow"
	// MigrationParameterTlsCreds ID of the 'tls-creds' object that provides credentials for establishing a TLS connection over the migration data channel. On the outgoing side of the migration, the credentials must be for a 'client' endpoint, while for the incoming side the credentials must be for a 'server' endpoint. Setting this will enable TLS for all migrations. The default is unset, resulting in unsecured migration at the QEMU level. (Since 2.7)
	MigrationParameterTlsCreds MigrationParameter = "tls-creds"
	// MigrationParameterTlsHostname hostname of the target host for the migration. This is required when using x509 based TLS credentials and the migration URI does not already include a hostname. For example
	MigrationParameterTlsHostname MigrationParameter = "tls-hostname"
	// MigrationParameterTlsAuthz ID of the 'authz' object subclass that provides access control checking of the TLS x509 certificate distinguished name. This object is only resolved at time of use, so can be deleted and recreated on the fly while the migration server is active. If missing, it will default to denying access (Since 4.0)
	MigrationParameterTlsAuthz MigrationParameter = "tls-authz"
	// MigrationParameterMaxBandwidth to set maximum speed for migration. maximum speed in bytes per second. (Since 2.8)
	MigrationParameterMaxBandwidth MigrationParameter = "max-bandwidth"
	// MigrationParameterAvailSwitchoverBandwidth to set the available bandwidth that migration can use during switchover phase. NOTE! This does not limit the bandwidth during switchover, but only for calculations when making decisions to switchover. By default, this value is zero, which means QEMU will estimate the bandwidth automatically. This can be set when the estimated value is not accurate, while the user is able to guarantee such bandwidth is available when switching over. When specified correctly, this can make the switchover decision much more accurate. (Since 8.2)
	MigrationParameterAvailSwitchoverBandwidth MigrationParameter = "avail-switchover-bandwidth"
	// MigrationParameterDowntimeLimit set maximum tolerated downtime for migration. maximum downtime in milliseconds (Since 2.8)
	MigrationParameterDowntimeLimit MigrationParameter = "downtime-limit"
	// MigrationParameterCheckpointDelay The delay time (in ms) between two COLO checkpoints in periodic mode. (Since 2.8)
	MigrationParameterCheckpointDelay MigrationParameter = "x-checkpoint-delay"
	// MigrationParameterBlockIncremental Affects how much storage is migrated when the block migration capability is enabled. When false, the entire storage backing chain is migrated into a flattened image at the destination; when true, only the active qcow2 layer is migrated and the destination must already have access to the same backing chain as was used on the source. (since 2.10)
	MigrationParameterBlockIncremental MigrationParameter = "block-incremental"
	// MigrationParameterMultifdChannels Number of channels used to migrate data in parallel. This is the same number that the number of sockets used for migration. The default value is 2 (since 4.0)
	MigrationParameterMultifdChannels MigrationParameter = "multifd-channels"
	// MigrationParameterXbzrleCacheSize cache size to be used by XBZRLE migration. It needs to be a multiple of the target page size and a power of 2 (Since 2.11)
	MigrationParameterXbzrleCacheSize MigrationParameter = "xbzrle-cache-size"
	// MigrationParameterMaxPostcopyBandwidth Background transfer bandwidth during postcopy. Defaults to 0 (unlimited). In bytes per second. (Since 3.0)
	MigrationParameterMaxPostcopyBandwidth MigrationParameter = "max-postcopy-bandwidth"
	// MigrationParameterMaxCpuThrottle maximum cpu throttle percentage. Defaults to 99. (Since 3.1)
	MigrationParameterMaxCpuThrottle MigrationParameter = "max-cpu-throttle"
	// MigrationParameterMultifdCompression Which compression method to use. Defaults to none. (Since 5.0)
	MigrationParameterMultifdCompression MigrationParameter = "multifd-compression"
	// MigrationParameterMultifdZlibLevel Set the compression level to be used in live migration, the compression level is an integer between 0 and 9, where 0 means no compression, 1 means the best compression speed, and 9 means best compression ratio which will consume more CPU. Defaults to 1. (Since 5.0)
	MigrationParameterMultifdZlibLevel MigrationParameter = "multifd-zlib-level"
	// MigrationParameterMultifdZstdLevel Set the compression level to be used in live migration, the compression level is an integer between 0 and 20, where 0 means no compression, 1 means the best compression speed, and 20 means best compression ratio which will consume more CPU. Defaults to 1. (Since 5.0)
	MigrationParameterMultifdZstdLevel MigrationParameter = "multifd-zstd-level"
	// MigrationParameterBlockBitmapMapping Maps block nodes and bitmaps on them to aliases for the purpose of dirty bitmap migration. Such aliases may for example be the corresponding names on the opposite site.
	MigrationParameterBlockBitmapMapping MigrationParameter = "block-bitmap-mapping"
	// MigrationParameterVcpuDirtyLimitPeriod Periodic time (in milliseconds) of dirty limit during live migration. Should be in the range 1 to 1000ms. Defaults to 1000ms. (Since 8.1)
	MigrationParameterVcpuDirtyLimitPeriod MigrationParameter = "x-vcpu-dirty-limit-period"
	// MigrationParameterVcpuDirtyLimit Dirtyrate limit (MB/s) during live migration. Defaults to 1. (Since 8.1)
	MigrationParameterVcpuDirtyLimit MigrationParameter = "vcpu-dirty-limit"
	// MigrationParameterMode Migration mode. See description in @MigMode. Default is 'normal'. (Since 8.2)
	MigrationParameterMode MigrationParameter = "mode"
)

// MigrateSetParameters
type MigrateSetParameters struct {
	// AnnounceInitial Initial delay (in milliseconds) before sending the first announce (Since 4.0)
	AnnounceInitial *uint64 `json:"announce-initial,omitempty"`
	// AnnounceMax Maximum delay (in milliseconds) between packets in the announcement (Since 4.0)
	AnnounceMax *uint64 `json:"announce-max,omitempty"`
	// AnnounceRounds Number of self-announce packets sent after migration (Since 4.0)
	AnnounceRounds *uint64 `json:"announce-rounds,omitempty"`
	// AnnounceStep Increase in delay (in milliseconds) between subsequent packets in the announcement (Since 4.0)
	AnnounceStep *uint64 `json:"announce-step,omitempty"`
	// CompressLevel compression level
	CompressLevel *uint8 `json:"compress-level,omitempty"`
	// CompressThreads compression thread count
	CompressThreads *uint8 `json:"compress-threads,omitempty"`
	// CompressWaitThread Controls behavior when all compression threads are currently busy. If true (default), wait for a free compression thread to become available; otherwise, send the page uncompressed. (Since 3.1)
	CompressWaitThread *bool `json:"compress-wait-thread,omitempty"`
	// DecompressThreads decompression thread count
	DecompressThreads *uint8 `json:"decompress-threads,omitempty"`
	// ThrottleTriggerThreshold The ratio of bytes_dirty_period and bytes_xfer_period to trigger throttling. It is expressed as percentage. The default value is 50. (Since 5.0)
	ThrottleTriggerThreshold *uint8 `json:"throttle-trigger-threshold,omitempty"`
	// CpuThrottleInitial Initial percentage of time guest cpus are throttled when migration auto-converge is activated. The default value is 20. (Since 2.7)
	CpuThrottleInitial *uint8 `json:"cpu-throttle-initial,omitempty"`
	// CpuThrottleIncrement throttle percentage increase each time auto-converge detects that migration is not making progress. The default value is 10. (Since 2.7)
	CpuThrottleIncrement *uint8 `json:"cpu-throttle-increment,omitempty"`
	// CpuThrottleTailslow Make CPU throttling slower at tail stage At the tail stage of throttling, the Guest is very sensitive to CPU percentage while the @cpu-throttle -increment is excessive usually at tail stage. If this parameter is true, we will compute the ideal CPU percentage used by the Guest, which may exactly make the dirty rate match the dirty rate threshold. Then we will choose a smaller throttle increment between the one specified by @cpu-throttle-increment and the one generated by ideal CPU percentage. Therefore, it is compatible to traditional throttling, meanwhile the throttle increment won't be excessive at tail stage. The default value is false. (Since 5.1)
	CpuThrottleTailslow *bool `json:"cpu-throttle-tailslow,omitempty"`
	// TlsCreds ID of the 'tls-creds' object that provides credentials for establishing a TLS connection over the migration data channel. On the outgoing side of the migration, the credentials must be for a 'client' endpoint, while for the incoming side the credentials must be for a 'server' endpoint. Setting this to a non-empty string enables TLS for all migrations. An empty string means that QEMU will use plain text mode for migration, rather than TLS (Since 2.9) Previously (since 2.7), this was reported by omitting tls-creds instead.
	TlsCreds *StrOrNull `json:"tls-creds,omitempty"`
	// TlsHostname hostname of the target host for the migration. This is required when using x509 based TLS credentials and the migration URI does not already include a hostname. For example
	TlsHostname *StrOrNull `json:"tls-hostname,omitempty"`
	// TlsAuthz ID of the 'authz' object subclass that provides access control checking of the TLS x509 certificate distinguished name. (Since 4.0)
	TlsAuthz *StrOrNull `json:"tls-authz,omitempty"`
	// MaxBandwidth to set maximum speed for migration. maximum speed in bytes per second. (Since 2.8)
	MaxBandwidth *uint64 `json:"max-bandwidth,omitempty"`
	// AvailSwitchoverBandwidth to set the available bandwidth that migration can use during switchover phase. NOTE! This does not limit the bandwidth during switchover, but only for calculations when making decisions to switchover. By default, this value is zero, which means QEMU will estimate the bandwidth automatically. This can be set when the estimated value is not accurate, while the user is able to guarantee such bandwidth is available when switching over. When specified correctly, this can make the switchover decision much more accurate. (Since 8.2)
	AvailSwitchoverBandwidth *uint64 `json:"avail-switchover-bandwidth,omitempty"`
	// DowntimeLimit set maximum tolerated downtime for migration. maximum downtime in milliseconds (Since 2.8)
	DowntimeLimit *uint64 `json:"downtime-limit,omitempty"`
	// CheckpointDelay the delay time between two COLO checkpoints. (Since 2.8)
	CheckpointDelay *uint32 `json:"x-checkpoint-delay,omitempty"`
	// BlockIncremental Affects how much storage is migrated when the block migration capability is enabled. When false, the entire storage backing chain is migrated into a flattened image at the destination; when true, only the active qcow2 layer is migrated and the destination must already have access to the same backing chain as was used on the source. (since 2.10)
	BlockIncremental *bool `json:"block-incremental,omitempty"`
	// MultifdChannels Number of channels used to migrate data in parallel. This is the same number that the number of sockets used for migration. The default value is 2 (since 4.0)
	MultifdChannels *uint8 `json:"multifd-channels,omitempty"`
	// XbzrleCacheSize cache size to be used by XBZRLE migration. It needs to be a multiple of the target page size and a power of 2 (Since 2.11)
	XbzrleCacheSize *uint64 `json:"xbzrle-cache-size,omitempty"`
	// MaxPostcopyBandwidth Background transfer bandwidth during postcopy. Defaults to 0 (unlimited). In bytes per second. (Since 3.0)
	MaxPostcopyBandwidth *uint64 `json:"max-postcopy-bandwidth,omitempty"`
	// MaxCpuThrottle maximum cpu throttle percentage. The default value is 99. (Since 3.1)
	MaxCpuThrottle *uint8 `json:"max-cpu-throttle,omitempty"`
	// MultifdCompression Which compression method to use. Defaults to none. (Since 5.0)
	MultifdCompression *MultiFDCompression `json:"multifd-compression,omitempty"`
	// MultifdZlibLevel Set the compression level to be used in live migration, the compression level is an integer between 0 and 9, where 0 means no compression, 1 means the best compression speed, and 9 means best compression ratio which will consume more CPU. Defaults to 1. (Since 5.0)
	MultifdZlibLevel *uint8 `json:"multifd-zlib-level,omitempty"`
	// MultifdZstdLevel Set the compression level to be used in live migration, the compression level is an integer between 0 and 20, where 0 means no compression, 1 means the best compression speed, and 20 means best compression ratio which will consume more CPU. Defaults to 1. (Since 5.0)
	MultifdZstdLevel *uint8 `json:"multifd-zstd-level,omitempty"`
	// BlockBitmapMapping Maps block nodes and bitmaps on them to aliases for the purpose of dirty bitmap migration. Such aliases may for example be the corresponding names on the opposite site.
	BlockBitmapMapping []BitmapMigrationNodeAlias `json:"block-bitmap-mapping,omitempty"`
	// VcpuDirtyLimitPeriod Periodic time (in milliseconds) of dirty limit during live migration. Should be in the range 1 to 1000ms. Defaults to 1000ms. (Since 8.1)
	VcpuDirtyLimitPeriod *uint64 `json:"x-vcpu-dirty-limit-period,omitempty"`
	// VcpuDirtyLimit Dirtyrate limit (MB/s) during live migration. Defaults to 1. (Since 8.1)
	VcpuDirtyLimit *uint64 `json:"vcpu-dirty-limit,omitempty"`
	// Mode Migration mode. See description in @MigMode. Default is 'normal'. (Since 8.2)
	Mode *MigMode `json:"mode,omitempty"`
}

func (MigrateSetParameters) Command() string {
	return "migrate-set-parameters"
}

func (cmd MigrateSetParameters) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "migrate-set-parameters", cmd, nil)
}

// MigrationParameters
//
// The optional members aren't actually optional.
type MigrationParameters struct {
	// AnnounceInitial Initial delay (in milliseconds) before sending the first announce (Since 4.0)
	AnnounceInitial *uint64 `json:"announce-initial,omitempty"`
	// AnnounceMax Maximum delay (in milliseconds) between packets in the announcement (Since 4.0)
	AnnounceMax *uint64 `json:"announce-max,omitempty"`
	// AnnounceRounds Number of self-announce packets sent after migration (Since 4.0)
	AnnounceRounds *uint64 `json:"announce-rounds,omitempty"`
	// AnnounceStep Increase in delay (in milliseconds) between subsequent packets in the announcement (Since 4.0)
	AnnounceStep *uint64 `json:"announce-step,omitempty"`
	// CompressLevel compression level
	CompressLevel *uint8 `json:"compress-level,omitempty"`
	// CompressThreads compression thread count
	CompressThreads *uint8 `json:"compress-threads,omitempty"`
	// CompressWaitThread Controls behavior when all compression threads are currently busy. If true (default), wait for a free compression thread to become available; otherwise, send the page uncompressed. (Since 3.1)
	CompressWaitThread *bool `json:"compress-wait-thread,omitempty"`
	// DecompressThreads decompression thread count
	DecompressThreads *uint8 `json:"decompress-threads,omitempty"`
	// ThrottleTriggerThreshold The ratio of bytes_dirty_period and bytes_xfer_period to trigger throttling. It is expressed as percentage. The default value is 50. (Since 5.0)
	ThrottleTriggerThreshold *uint8 `json:"throttle-trigger-threshold,omitempty"`
	// CpuThrottleInitial Initial percentage of time guest cpus are throttled when migration auto-converge is activated. (Since 2.7)
	CpuThrottleInitial *uint8 `json:"cpu-throttle-initial,omitempty"`
	// CpuThrottleIncrement throttle percentage increase each time auto-converge detects that migration is not making progress. (Since 2.7)
	CpuThrottleIncrement *uint8 `json:"cpu-throttle-increment,omitempty"`
	// CpuThrottleTailslow Make CPU throttling slower at tail stage At the tail stage of throttling, the Guest is very sensitive to CPU percentage while the @cpu-throttle -increment is excessive usually at tail stage. If this parameter is true, we will compute the ideal CPU percentage used by the Guest, which may exactly make the dirty rate match the dirty rate threshold. Then we will choose a smaller throttle increment between the one specified by @cpu-throttle-increment and the one generated by ideal CPU percentage. Therefore, it is compatible to traditional throttling, meanwhile the throttle increment won't be excessive at tail stage. The default value is false. (Since 5.1)
	CpuThrottleTailslow *bool `json:"cpu-throttle-tailslow,omitempty"`
	// TlsCreds ID of the 'tls-creds' object that provides credentials for establishing a TLS connection over the migration data channel. On the outgoing side of the migration, the credentials must be for a 'client' endpoint, while for the incoming side the credentials must be for a 'server' endpoint. An empty string means that QEMU will use plain text mode for migration, rather
	TlsCreds *string `json:"tls-creds,omitempty"`
	// TlsHostname hostname of the target host for the migration. This is required when using x509 based TLS credentials and the migration URI does not already include a hostname. For example
	TlsHostname *string `json:"tls-hostname,omitempty"`
	// TlsAuthz ID of the 'authz' object subclass that provides access control checking of the TLS x509 certificate distinguished name. (Since 4.0)
	TlsAuthz *string `json:"tls-authz,omitempty"`
	// MaxBandwidth to set maximum speed for migration. maximum speed in bytes per second. (Since 2.8)
	MaxBandwidth *uint64 `json:"max-bandwidth,omitempty"`
	// AvailSwitchoverBandwidth to set the available bandwidth that migration can use during switchover phase. NOTE! This does not limit the bandwidth during switchover, but only for calculations when making decisions to switchover. By default, this value is zero, which means QEMU will estimate the bandwidth automatically. This can be set when the estimated value is not accurate, while the user is able to guarantee such bandwidth is available when switching over. When specified correctly, this can make the switchover decision much more accurate. (Since 8.2)
	AvailSwitchoverBandwidth *uint64 `json:"avail-switchover-bandwidth,omitempty"`
	// DowntimeLimit set maximum tolerated downtime for migration. maximum downtime in milliseconds (Since 2.8)
	DowntimeLimit *uint64 `json:"downtime-limit,omitempty"`
	// CheckpointDelay the delay time between two COLO checkpoints. (Since 2.8)
	CheckpointDelay *uint32 `json:"x-checkpoint-delay,omitempty"`
	// BlockIncremental Affects how much storage is migrated when the block migration capability is enabled. When false, the entire storage backing chain is migrated into a flattened image at the destination; when true, only the active qcow2 layer is migrated and the destination must already have access to the same backing chain as was used on the source. (since 2.10)
	BlockIncremental *bool `json:"block-incremental,omitempty"`
	// MultifdChannels Number of channels used to migrate data in parallel. This is the same number that the number of sockets used for migration. The default value is 2 (since 4.0)
	MultifdChannels *uint8 `json:"multifd-channels,omitempty"`
	// XbzrleCacheSize cache size to be used by XBZRLE migration. It needs to be a multiple of the target page size and a power of 2 (Since 2.11)
	XbzrleCacheSize *uint64 `json:"xbzrle-cache-size,omitempty"`
	// MaxPostcopyBandwidth Background transfer bandwidth during postcopy. Defaults to 0 (unlimited). In bytes per second. (Since 3.0)
	MaxPostcopyBandwidth *uint64 `json:"max-postcopy-bandwidth,omitempty"`
	// MaxCpuThrottle maximum cpu throttle percentage. Defaults to 99. (Since 3.1)
	MaxCpuThrottle *uint8 `json:"max-cpu-throttle,omitempty"`
	// MultifdCompression Which compression method to use. Defaults to none. (Since 5.0)
	MultifdCompression *MultiFDCompression `json:"multifd-compression,omitempty"`
	// MultifdZlibLevel Set the compression level to be used in live migration, the compression level is an integer between 0 and 9, where 0 means no compression, 1 means the best compression speed, and 9 means best compression ratio which will consume more CPU. Defaults to 1. (Since 5.0)
	MultifdZlibLevel *uint8 `json:"multifd-zlib-level,omitempty"`
	// MultifdZstdLevel Set the compression level to be used in live migration, the compression level is an integer between 0 and 20, where 0 means no compression, 1 means the best compression speed, and 20 means best compression ratio which will consume more CPU. Defaults to 1. (Since 5.0)
	MultifdZstdLevel *uint8 `json:"multifd-zstd-level,omitempty"`
	// BlockBitmapMapping Maps block nodes and bitmaps on them to aliases for the purpose of dirty bitmap migration. Such aliases may for example be the corresponding names on the opposite site.
	BlockBitmapMapping []BitmapMigrationNodeAlias `json:"block-bitmap-mapping,omitempty"`
	// VcpuDirtyLimitPeriod Periodic time (in milliseconds) of dirty limit during live migration. Should be in the range 1 to 1000ms. Defaults to 1000ms. (Since 8.1)
	VcpuDirtyLimitPeriod *uint64 `json:"x-vcpu-dirty-limit-period,omitempty"`
	// VcpuDirtyLimit Dirtyrate limit (MB/s) during live migration. Defaults to 1. (Since 8.1)
	VcpuDirtyLimit *uint64 `json:"vcpu-dirty-limit,omitempty"`
	// Mode Migration mode. See description in @MigMode. Default is 'normal'. (Since 8.2)
	Mode *MigMode `json:"mode,omitempty"`
}

// QueryMigrateParameters
//
// Returns information about the current migration parameters
type QueryMigrateParameters struct {
}

func (QueryMigrateParameters) Command() string {
	return "query-migrate-parameters"
}

func (cmd QueryMigrateParameters) Execute(ctx context.Context, client api.Client) (MigrationParameters, error) {
	var ret MigrationParameters

	return ret, client.Execute(ctx, "query-migrate-parameters", cmd, &ret)
}

// MigrateStartPostcopy
//
// Followup to a migration command to switch the migration to postcopy mode. The postcopy-ram capability must be set on both source and destination before the original migration command.
type MigrateStartPostcopy struct {
}

func (MigrateStartPostcopy) Command() string {
	return "migrate-start-postcopy"
}

func (cmd MigrateStartPostcopy) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "migrate-start-postcopy", cmd, nil)
}

// MigrationEvent (MIGRATION)
//
// Emitted when a migration event happens
type MigrationEvent struct {
	// Status @MigrationStatus describing the current migration status.
	Status MigrationStatus `json:"status"`
}

func (MigrationEvent) Event() string {
	return "MIGRATION"
}

// MigrationPassEvent (MIGRATION_PASS)
//
// Emitted from the source side of a migration at the start of each pass (when it syncs the dirty bitmap)
type MigrationPassEvent struct {
	// Pass An incrementing count (starting at 1 on the first pass)
	Pass int64 `json:"pass"`
}

func (MigrationPassEvent) Event() string {
	return "MIGRATION_PASS"
}

// COLOMessage The message transmission between Primary side and Secondary side.
type COLOMessage string

const (
	// COLOMessageCheckpointReady Secondary VM (SVM) is ready for checkpointing
	COLOMessageCheckpointReady COLOMessage = "checkpoint-ready"
	// COLOMessageCheckpointRequest Primary VM (PVM) tells SVM to prepare for checkpointing
	COLOMessageCheckpointRequest COLOMessage = "checkpoint-request"
	// COLOMessageCheckpointReply SVM gets PVM's checkpoint request
	COLOMessageCheckpointReply COLOMessage = "checkpoint-reply"
	// COLOMessageVmstateSend VM's state will be sent by PVM.
	COLOMessageVmstateSend COLOMessage = "vmstate-send"
	// COLOMessageVmstateSize The total size of VMstate.
	COLOMessageVmstateSize COLOMessage = "vmstate-size"
	// COLOMessageVmstateReceived VM's state has been received by SVM.
	COLOMessageVmstateReceived COLOMessage = "vmstate-received"
	// COLOMessageVmstateLoaded VM's state has been loaded by SVM.
	COLOMessageVmstateLoaded COLOMessage = "vmstate-loaded"
)

// COLOMode The COLO current mode.
type COLOMode string

const (
	// COLOModeNone COLO is disabled.
	COLOModeNone COLOMode = "none"
	// COLOModePrimary COLO node in primary side.
	COLOModePrimary COLOMode = "primary"
	// COLOModeSecondary COLO node in slave side.
	COLOModeSecondary COLOMode = "secondary"
)

// FailoverStatus An enumeration of COLO failover status
type FailoverStatus string

const (
	// FailoverStatusNone no failover has ever happened
	FailoverStatusNone FailoverStatus = "none"
	// FailoverStatusRequire got failover requirement but not handled
	FailoverStatusRequire FailoverStatus = "require"
	// FailoverStatusActive in the process of doing failover
	FailoverStatusActive FailoverStatus = "active"
	// FailoverStatusCompleted finish the process of failover
	FailoverStatusCompleted FailoverStatus = "completed"
	// FailoverStatusRelaunch restart the failover process, from 'none' -> 'completed' (Since 2.9)
	FailoverStatusRelaunch FailoverStatus = "relaunch"
)

// ColoExitEvent (COLO_EXIT)
//
// Emitted when VM finishes COLO mode due to some errors happening or at the request of users.
type ColoExitEvent struct {
	// Mode report COLO mode when COLO exited.
	Mode COLOMode `json:"mode"`
	// Reason describes the reason for the COLO exit.
	Reason COLOExitReason `json:"reason"`
}

func (ColoExitEvent) Event() string {
	return "COLO_EXIT"
}

// COLOExitReason The reason for a COLO exit.
type COLOExitReason string

const (
	// COLOExitReasonNone failover has never happened. This state does not occur in the COLO_EXIT event, and is only visible in the result of query-colo-status.
	COLOExitReasonNone COLOExitReason = "none"
	// COLOExitReasonRequest COLO exit is due to an external request.
	COLOExitReasonRequest COLOExitReason = "request"
	// COLOExitReasonError COLO exit is due to an internal error.
	COLOExitReasonError COLOExitReason = "error"
	// COLOExitReasonProcessing COLO is currently handling a failover (since 4.0).
	COLOExitReasonProcessing COLOExitReason = "processing"
)

// ColoLostHeartbeat
//
// Tell qemu that heartbeat is lost, request it to do takeover procedures. If this command is sent to the PVM, the Primary side will exit COLO mode. If sent to the Secondary, the Secondary side will run failover work, then takes over server operation to become the service VM.
type ColoLostHeartbeat struct {
}

func (ColoLostHeartbeat) Command() string {
	return "x-colo-lost-heartbeat"
}

func (cmd ColoLostHeartbeat) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "x-colo-lost-heartbeat", cmd, nil)
}

// MigrateCancel
//
// Cancel the current executing migration process.
type MigrateCancel struct {
}

func (MigrateCancel) Command() string {
	return "migrate_cancel"
}

func (cmd MigrateCancel) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "migrate_cancel", cmd, nil)
}

// MigrateContinue
//
// Continue migration when it's in a paused state.
type MigrateContinue struct {
	// State The state the migration is currently expected to be in
	State MigrationStatus `json:"state"`
}

func (MigrateContinue) Command() string {
	return "migrate-continue"
}

func (cmd MigrateContinue) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "migrate-continue", cmd, nil)
}

// MigrationAddressType The migration stream transport mechanisms.
type MigrationAddressType string

const (
	// MigrationAddressTypeSocket Migrate via socket.
	MigrationAddressTypeSocket MigrationAddressType = "socket"
	// MigrationAddressTypeExec Direct the migration stream to another process.
	MigrationAddressTypeExec MigrationAddressType = "exec"
	// MigrationAddressTypeRdma Migrate via RDMA.
	MigrationAddressTypeRdma MigrationAddressType = "rdma"
	// MigrationAddressTypeFile Direct the migration stream to a file.
	MigrationAddressTypeFile MigrationAddressType = "file"
)

// FileMigrationArgs
type FileMigrationArgs struct {
	// Filename The file to receive the migration stream
	Filename string `json:"filename"`
	// Offset The file offset where the migration stream will start
	Offset uint64 `json:"offset"`
}

// MigrationExecCommand
type MigrationExecCommand struct {
	// Args command (list head) and arguments to execute.
	Args []string `json:"args"`
}

// MigrationAddress
//
// Migration endpoint configuration.
type MigrationAddress struct {
	// Discriminator: transport

	// Transport The migration stream transport mechanism
	Transport MigrationAddressType `json:"transport"`

	Socket *SocketAddress        `json:"-"`
	Exec   *MigrationExecCommand `json:"-"`
	Rdma   *InetSocketAddress    `json:"-"`
	File   *FileMigrationArgs    `json:"-"`
}

func (u MigrationAddress) MarshalJSON() ([]byte, error) {
	switch u.Transport {
	case "socket":
		if u.Socket == nil {
			return nil, fmt.Errorf("expected Socket to be set")
		}

		return json.Marshal(struct {
			Transport MigrationAddressType `json:"transport"`
			*SocketAddress
		}{
			Transport:     u.Transport,
			SocketAddress: u.Socket,
		})
	case "exec":
		if u.Exec == nil {
			return nil, fmt.Errorf("expected Exec to be set")
		}

		return json.Marshal(struct {
			Transport MigrationAddressType `json:"transport"`
			*MigrationExecCommand
		}{
			Transport:            u.Transport,
			MigrationExecCommand: u.Exec,
		})
	case "rdma":
		if u.Rdma == nil {
			return nil, fmt.Errorf("expected Rdma to be set")
		}

		return json.Marshal(struct {
			Transport MigrationAddressType `json:"transport"`
			*InetSocketAddress
		}{
			Transport:         u.Transport,
			InetSocketAddress: u.Rdma,
		})
	case "file":
		if u.File == nil {
			return nil, fmt.Errorf("expected File to be set")
		}

		return json.Marshal(struct {
			Transport MigrationAddressType `json:"transport"`
			*FileMigrationArgs
		}{
			Transport:         u.Transport,
			FileMigrationArgs: u.File,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// MigrationChannelType The migration channel-type request options.
type MigrationChannelType string

const (
	// MigrationChannelTypeMain Main outbound migration channel.
	MigrationChannelTypeMain MigrationChannelType = "main"
)

// MigrationChannel
//
// Migration stream channel parameters.
type MigrationChannel struct {
	// ChannelType Channel type for transferring packet information.
	ChannelType MigrationChannelType `json:"channel-type"`
	// Addr Migration endpoint configuration on destination interface.
	Addr MigrationAddress `json:"addr"`
}

// Migrate
//
// Migrates the current running guest to another Virtual Machine.
type Migrate struct {
	// Uri the Uniform Resource Identifier of the destination VM
	Uri *string `json:"uri,omitempty"`
	// Channels list of migration stream channels with each stream in the list connected to a destination interface endpoint.
	Channels []MigrationChannel `json:"channels,omitempty"`
	// Blk do block migration (full disk copy)
	Blk *bool `json:"blk,omitempty"`
	// Inc incremental disk copy migration
	Inc *bool `json:"inc,omitempty"`
	// Detach this argument exists only for compatibility reasons and is ignored by QEMU
	Detach *bool `json:"detach,omitempty"`
	// Resume resume one paused migration, default "off". (since 3.0)
	Resume *bool `json:"resume,omitempty"`
}

func (Migrate) Command() string {
	return "migrate"
}

func (cmd Migrate) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "migrate", cmd, nil)
}

// MigrateIncoming
//
// Start an incoming migration, the qemu must have been started with -incoming defer
type MigrateIncoming struct {
	// Uri The Uniform Resource Identifier identifying the source or address to listen on
	Uri *string `json:"uri,omitempty"`
	// Channels list of migration stream channels with each stream in the list connected to a destination interface endpoint.
	Channels []MigrationChannel `json:"channels,omitempty"`
}

func (MigrateIncoming) Command() string {
	return "migrate-incoming"
}

func (cmd MigrateIncoming) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "migrate-incoming", cmd, nil)
}

// XenSaveDevicesState
//
// Save the state of all devices to file. The RAM and the block devices of the VM are not saved by this command.
type XenSaveDevicesState struct {
	// Filename the file to save the state of the devices to as binary data. See xen-save-devices-state.txt for a description of the binary format.
	Filename string `json:"filename"`
	// Live Optional argument to ask QEMU to treat this command as part of a live migration. Default to true. (since 2.11)
	Live *bool `json:"live,omitempty"`
}

func (XenSaveDevicesState) Command() string {
	return "xen-save-devices-state"
}

func (cmd XenSaveDevicesState) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "xen-save-devices-state", cmd, nil)
}

// XenSetGlobalDirtyLog
//
// Enable or disable the global dirty log mode.
type XenSetGlobalDirtyLog struct {
	// Enable true to enable, false to disable.
	Enable bool `json:"enable"`
}

func (XenSetGlobalDirtyLog) Command() string {
	return "xen-set-global-dirty-log"
}

func (cmd XenSetGlobalDirtyLog) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "xen-set-global-dirty-log", cmd, nil)
}

// XenLoadDevicesState
//
// Load the state of all devices from file. The RAM and the block devices of the VM are not loaded by this command.
type XenLoadDevicesState struct {
	// Filename the file to load the state of the devices from as binary data. See xen-save-devices-state.txt for a description of the binary format.
	Filename string `json:"filename"`
}

func (XenLoadDevicesState) Command() string {
	return "xen-load-devices-state"
}

func (cmd XenLoadDevicesState) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "xen-load-devices-state", cmd, nil)
}

// XenSetReplication
//
// Enable or disable replication.
type XenSetReplication struct {
	// Enable true to enable, false to disable.
	Enable bool `json:"enable"`
	// Primary true for primary or false for secondary.
	Primary bool `json:"primary"`
	// Failover true to do failover, false to stop. but cannot be specified if 'enable' is true. default value is false.
	Failover *bool `json:"failover,omitempty"`
}

func (XenSetReplication) Command() string {
	return "xen-set-replication"
}

func (cmd XenSetReplication) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "xen-set-replication", cmd, nil)
}

// ReplicationStatus
//
// The result format for 'query-xen-replication-status'.
type ReplicationStatus struct {
	// Error true if an error happened, false if replication is normal.
	Error bool `json:"error"`
	// Desc the human readable error description string, when @error is 'true'.
	Desc *string `json:"desc,omitempty"`
}

// QueryXenReplicationStatus
//
// Query replication status while the vm is running.
type QueryXenReplicationStatus struct {
}

func (QueryXenReplicationStatus) Command() string {
	return "query-xen-replication-status"
}

func (cmd QueryXenReplicationStatus) Execute(ctx context.Context, client api.Client) (ReplicationStatus, error) {
	var ret ReplicationStatus

	return ret, client.Execute(ctx, "query-xen-replication-status", cmd, &ret)
}

// XenColoDoCheckpoint
//
// Xen uses this command to notify replication to trigger a checkpoint.
type XenColoDoCheckpoint struct {
}

func (XenColoDoCheckpoint) Command() string {
	return "xen-colo-do-checkpoint"
}

func (cmd XenColoDoCheckpoint) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "xen-colo-do-checkpoint", cmd, nil)
}

// COLOStatus
//
// The result format for 'query-colo-status'.
type COLOStatus struct {
	// Mode COLO running mode. If COLO is running, this field will return 'primary' or 'secondary'.
	Mode COLOMode `json:"mode"`
	// LastMode COLO last running mode. If COLO is running, this field will return same like mode field, after failover we can use this field to get last colo mode. (since 4.0)
	LastMode COLOMode `json:"last-mode"`
	// Reason describes the reason for the COLO exit.
	Reason COLOExitReason `json:"reason"`
}

// QueryColoStatus
//
// Query COLO status while the vm is running.
type QueryColoStatus struct {
}

func (QueryColoStatus) Command() string {
	return "query-colo-status"
}

func (cmd QueryColoStatus) Execute(ctx context.Context, client api.Client) (COLOStatus, error) {
	var ret COLOStatus

	return ret, client.Execute(ctx, "query-colo-status", cmd, &ret)
}

// MigrateRecover
//
// Provide a recovery migration stream URI.
type MigrateRecover struct {
	// Uri the URI to be used for the recovery of migration stream.
	Uri string `json:"uri"`
}

func (MigrateRecover) Command() string {
	return "migrate-recover"
}

func (cmd MigrateRecover) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "migrate-recover", cmd, nil)
}

// MigratePause
//
// Pause a migration. Currently it only supports postcopy.
type MigratePause struct {
}

func (MigratePause) Command() string {
	return "migrate-pause"
}

func (cmd MigratePause) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "migrate-pause", cmd, nil)
}

// UnplugPrimaryEvent (UNPLUG_PRIMARY)
//
// Emitted from source side of a migration when migration state is WAIT_UNPLUG. Device was unplugged by guest operating system. Device resources in QEMU are kept on standby to be able to re-plug it in case of migration failure.
type UnplugPrimaryEvent struct {
	// DeviceId QEMU device id of the unplugged device
	DeviceId string `json:"device-id"`
}

func (UnplugPrimaryEvent) Event() string {
	return "UNPLUG_PRIMARY"
}

// DirtyRateVcpu
//
// Dirty rate of vcpu.
type DirtyRateVcpu struct {
	// Id vcpu index.
	Id int64 `json:"id"`
	// DirtyRate dirty rate.
	DirtyRate int64 `json:"dirty-rate"`
}

// DirtyRateStatus Dirty page rate measurement status.
type DirtyRateStatus string

const (
	// DirtyRateStatusUnstarted measuring thread has not been started yet
	DirtyRateStatusUnstarted DirtyRateStatus = "unstarted"
	// DirtyRateStatusMeasuring measuring thread is running
	DirtyRateStatusMeasuring DirtyRateStatus = "measuring"
	// DirtyRateStatusMeasured dirty page rate is measured and the results are available
	DirtyRateStatusMeasured DirtyRateStatus = "measured"
)

// DirtyRateMeasureMode Method used to measure dirty page rate. Differences between available methods are explained in @calc-dirty-rate.
type DirtyRateMeasureMode string

const (
	// DirtyRateMeasureModePageSampling use page sampling
	DirtyRateMeasureModePageSampling DirtyRateMeasureMode = "page-sampling"
	// DirtyRateMeasureModeDirtyRing use dirty ring
	DirtyRateMeasureModeDirtyRing DirtyRateMeasureMode = "dirty-ring"
	// DirtyRateMeasureModeDirtyBitmap use dirty bitmap
	DirtyRateMeasureModeDirtyBitmap DirtyRateMeasureMode = "dirty-bitmap"
)

// TimeUnit Specifies unit in which time-related value is specified.
type TimeUnit string

const (
	// TimeUnitSecond value is in seconds
	TimeUnitSecond TimeUnit = "second"
	// TimeUnitMillisecond value is in milliseconds
	TimeUnitMillisecond TimeUnit = "millisecond"
)

// DirtyRateInfo
//
// Information about measured dirty page rate.
type DirtyRateInfo struct {
	// DirtyRate an estimate of the dirty page rate of the VM in units of MiB/s. Value is present only when @status is 'measured'.
	DirtyRate *int64 `json:"dirty-rate,omitempty"`
	// Status current status of dirty page rate measurements
	Status DirtyRateStatus `json:"status"`
	// StartTime start time in units of second for calculation
	StartTime int64 `json:"start-time"`
	// CalcTime time period for which dirty page rate was measured, expressed and rounded down to @calc-time-unit.
	CalcTime int64 `json:"calc-time"`
	// CalcTimeUnit time unit of @calc-time (Since 8.2)
	CalcTimeUnit TimeUnit `json:"calc-time-unit"`
	// SamplePages number of sampled pages per GiB of guest memory. Valid only in page-sampling mode (Since 6.1)
	SamplePages uint64 `json:"sample-pages"`
	// Mode mode that was used to measure dirty page rate (Since 6.2)
	Mode DirtyRateMeasureMode `json:"mode"`
	// VcpuDirtyRate dirty rate for each vCPU if dirty-ring mode was specified (Since 6.2)
	VcpuDirtyRate []DirtyRateVcpu `json:"vcpu-dirty-rate,omitempty"`
}

// CalcDirtyRate
//
// Start measuring dirty page rate of the VM. Results can be retrieved with @query-dirty-rate after measurements are completed. Dirty page rate is the number of pages changed in a given time period expressed in MiB/s. The following methods of calculation are
type CalcDirtyRate struct {
	// CalcTime time period for which dirty page rate is calculated. By default it is specified in seconds, but the unit can be set explicitly with @calc-time-unit. Note that larger @calc-time values will typically result in smaller dirty page rates because page dirtying is a one-time event. Once some page is counted as dirty during @calc-time period, further writes to this page will not increase dirty page rate anymore.
	CalcTime int64 `json:"calc-time"`
	// CalcTimeUnit time unit in which @calc-time is specified. By default it is seconds. (Since 8.2)
	CalcTimeUnit *TimeUnit `json:"calc-time-unit,omitempty"`
	// SamplePages number of sampled pages per each GiB of guest memory. Default value is 512. For 4KiB guest pages this corresponds to sampling ratio of 0.2%. This argument is used only in page sampling mode. (Since 6.1)
	SamplePages *int64 `json:"sample-pages,omitempty"`
	// Mode mechanism for tracking dirty pages. Default value is 'page-sampling'. Others are 'dirty-bitmap' and 'dirty-ring'. (Since 6.1)
	Mode *DirtyRateMeasureMode `json:"mode,omitempty"`
}

func (CalcDirtyRate) Command() string {
	return "calc-dirty-rate"
}

func (cmd CalcDirtyRate) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "calc-dirty-rate", cmd, nil)
}

// QueryDirtyRate
//
// Query results of the most recent invocation of @calc-dirty-rate.
type QueryDirtyRate struct {
	// CalcTimeUnit time unit in which to report calculation time. By default it is reported in seconds. (Since 8.2)
	CalcTimeUnit *TimeUnit `json:"calc-time-unit,omitempty"`
}

func (QueryDirtyRate) Command() string {
	return "query-dirty-rate"
}

func (cmd QueryDirtyRate) Execute(ctx context.Context, client api.Client) (DirtyRateInfo, error) {
	var ret DirtyRateInfo

	return ret, client.Execute(ctx, "query-dirty-rate", cmd, &ret)
}

// DirtyLimitInfo
//
// Dirty page rate limit information of a virtual CPU.
type DirtyLimitInfo struct {
	// CpuIndex index of a virtual CPU.
	CpuIndex int64 `json:"cpu-index"`
	// LimitRate upper limit of dirty page rate (MB/s) for a virtual CPU, 0 means unlimited.
	LimitRate uint64 `json:"limit-rate"`
	// CurrentRate current dirty page rate (MB/s) for a virtual CPU.
	CurrentRate uint64 `json:"current-rate"`
}

// SetVcpuDirtyLimit
//
// Set the upper limit of dirty page rate for virtual CPUs. Requires KVM with accelerator property "dirty-ring-size" set. A virtual CPU's dirty page rate is a measure of its memory load. To observe dirty page rates, use @calc-dirty-rate.
type SetVcpuDirtyLimit struct {
	// CpuIndex index of a virtual CPU, default is all.
	CpuIndex *int64 `json:"cpu-index,omitempty"`
	// DirtyRate upper limit of dirty page rate (MB/s) for virtual CPUs.
	DirtyRate uint64 `json:"dirty-rate"`
}

func (SetVcpuDirtyLimit) Command() string {
	return "set-vcpu-dirty-limit"
}

func (cmd SetVcpuDirtyLimit) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "set-vcpu-dirty-limit", cmd, nil)
}

// CancelVcpuDirtyLimit
//
// Cancel the upper limit of dirty page rate for virtual CPUs. Cancel the dirty page limit for the vCPU which has been set with set-vcpu-dirty-limit command. Note that this command requires support from dirty ring, same as the "set-vcpu-dirty-limit".
type CancelVcpuDirtyLimit struct {
	// CpuIndex index of a virtual CPU, default is all.
	CpuIndex *int64 `json:"cpu-index,omitempty"`
}

func (CancelVcpuDirtyLimit) Command() string {
	return "cancel-vcpu-dirty-limit"
}

func (cmd CancelVcpuDirtyLimit) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "cancel-vcpu-dirty-limit", cmd, nil)
}

// QueryVcpuDirtyLimit
//
// Returns information about virtual CPU dirty page rate limits, if any.
type QueryVcpuDirtyLimit struct {
}

func (QueryVcpuDirtyLimit) Command() string {
	return "query-vcpu-dirty-limit"
}

func (cmd QueryVcpuDirtyLimit) Execute(ctx context.Context, client api.Client) ([]DirtyLimitInfo, error) {
	var ret []DirtyLimitInfo

	return ret, client.Execute(ctx, "query-vcpu-dirty-limit", cmd, &ret)
}

// MigrationThreadInfo
//
// Information about migrationthreads
type MigrationThreadInfo struct {
	// Name the name of migration thread
	Name string `json:"name"`
	// ThreadId ID of the underlying host thread
	ThreadId int64 `json:"thread-id"`
}

// QueryMigrationthreads
//
// Returns information of migration threads
type QueryMigrationthreads struct {
}

func (QueryMigrationthreads) Command() string {
	return "query-migrationthreads"
}

func (cmd QueryMigrationthreads) Execute(ctx context.Context, client api.Client) ([]MigrationThreadInfo, error) {
	var ret []MigrationThreadInfo

	return ret, client.Execute(ctx, "query-migrationthreads", cmd, &ret)
}

// SnapshotSave
//
// Save a VM snapshot
type SnapshotSave struct {
	// JobId identifier for the newly created job
	JobId string `json:"job-id"`
	// Tag name of the snapshot to create
	Tag string `json:"tag"`
	// Vmstate block device node name to save vmstate to
	Vmstate string `json:"vmstate"`
	// Devices list of block device node names to save a snapshot to Applications should not assume that the snapshot save is complete when this command returns. The job commands / events must be used to determine completion and to fetch details of any errors that arise. Note that execution of the guest CPUs may be stopped during the time it takes to save the snapshot. A future version of QEMU may ensure CPUs are executing continuously. It is strongly recommended that @devices contain all writable block device nodes if a consistent snapshot is required. If @tag already exists, an error will be reported
	Devices []string `json:"devices"`
}

func (SnapshotSave) Command() string {
	return "snapshot-save"
}

func (cmd SnapshotSave) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "snapshot-save", cmd, nil)
}

// SnapshotLoad
//
// Load a VM snapshot
type SnapshotLoad struct {
	// JobId identifier for the newly created job
	JobId string `json:"job-id"`
	// Tag name of the snapshot to load.
	Tag string `json:"tag"`
	// Vmstate block device node name to load vmstate from
	Vmstate string `json:"vmstate"`
	// Devices list of block device node names to load a snapshot from Applications should not assume that the snapshot load is complete when this command returns. The job commands / events must be used to determine completion and to fetch details of any errors that arise. Note that execution of the guest CPUs will be stopped during the time it takes to load the snapshot. It is strongly recommended that @devices contain all writable block device nodes that can have changed since the original @snapshot-save command execution.
	Devices []string `json:"devices"`
}

func (SnapshotLoad) Command() string {
	return "snapshot-load"
}

func (cmd SnapshotLoad) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "snapshot-load", cmd, nil)
}

// SnapshotDelete
//
// Delete a VM snapshot
type SnapshotDelete struct {
	// JobId identifier for the newly created job
	JobId string `json:"job-id"`
	// Tag name of the snapshot to delete.
	Tag string `json:"tag"`
	// Devices list of block device node names to delete a snapshot from Applications should not assume that the snapshot delete is complete when this command returns. The job commands / events must be used to determine completion and to fetch details of any errors that arise.
	Devices []string `json:"devices"`
}

func (SnapshotDelete) Command() string {
	return "snapshot-delete"
}

func (cmd SnapshotDelete) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "snapshot-delete", cmd, nil)
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

// TraceEventState State of a tracing event.
type TraceEventState string

const (
	// TraceEventStateUnavailable The event is statically disabled.
	TraceEventStateUnavailable TraceEventState = "unavailable"
	// TraceEventStateDisabled The event is dynamically disabled.
	TraceEventStateDisabled TraceEventState = "disabled"
	// TraceEventStateEnabled The event is dynamically enabled.
	TraceEventStateEnabled TraceEventState = "enabled"
)

// TraceEventInfo
//
// Information of a tracing event.
type TraceEventInfo struct {
	// Name Event name.
	Name string `json:"name"`
	// State Tracing state.
	State TraceEventState `json:"state"`
	// Vcpu Whether this is a per-vCPU event (since 2.7).
	Vcpu bool `json:"vcpu"`
}

// TraceEventGetState
//
// Query the state of events.
type TraceEventGetState struct {
	// Name Event name pattern (case-sensitive glob).
	Name string `json:"name"`
	// Vcpu The vCPU to query (since 2.7).
	Vcpu *int64 `json:"vcpu,omitempty"`
}

func (TraceEventGetState) Command() string {
	return "trace-event-get-state"
}

func (cmd TraceEventGetState) Execute(ctx context.Context, client api.Client) ([]TraceEventInfo, error) {
	var ret []TraceEventInfo

	return ret, client.Execute(ctx, "trace-event-get-state", cmd, &ret)
}

// TraceEventSetState
//
// Set the dynamic tracing state of events.
type TraceEventSetState struct {
	// Name Event name pattern (case-sensitive glob).
	Name string `json:"name"`
	// Enable Whether to enable tracing.
	Enable bool `json:"enable"`
	// IgnoreUnavailable Do not match unavailable events with @name.
	IgnoreUnavailable *bool `json:"ignore-unavailable,omitempty"`
	// Vcpu The vCPU to act upon (all by default; since 2.7).
	Vcpu *int64 `json:"vcpu,omitempty"`
}

func (TraceEventSetState) Command() string {
	return "trace-event-set-state"
}

func (cmd TraceEventSetState) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "trace-event-set-state", cmd, nil)
}

// CompatPolicyInput Policy for handling "funny" input.
type CompatPolicyInput string

const (
	// CompatPolicyInputAccept Accept silently
	CompatPolicyInputAccept CompatPolicyInput = "accept"
	// CompatPolicyInputReject Reject with an error
	CompatPolicyInputReject CompatPolicyInput = "reject"
	// CompatPolicyInputCrash abort() the process
	CompatPolicyInputCrash CompatPolicyInput = "crash"
)

// CompatPolicyOutput Policy for handling "funny" output.
type CompatPolicyOutput string

const (
	// CompatPolicyOutputAccept Pass on unchanged
	CompatPolicyOutputAccept CompatPolicyOutput = "accept"
	// CompatPolicyOutputHide Filter out
	CompatPolicyOutputHide CompatPolicyOutput = "hide"
)

// CompatPolicy
//
// Policy for handling deprecated management interfaces. This is intended for testing users of the management interfaces.
type CompatPolicy struct {
	// DeprecatedInput how to handle deprecated input (default 'accept')
	DeprecatedInput *CompatPolicyInput `json:"deprecated-input,omitempty"`
	// DeprecatedOutput how to handle deprecated output (default 'accept')
	DeprecatedOutput *CompatPolicyOutput `json:"deprecated-output,omitempty"`
	// UnstableInput how to handle unstable input (default 'accept') (since 6.2)
	UnstableInput *CompatPolicyInput `json:"unstable-input,omitempty"`
	// UnstableOutput how to handle unstable output (default 'accept') (since 6.2)
	UnstableOutput *CompatPolicyOutput `json:"unstable-output,omitempty"`
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

// DeviceListProperties
//
// List properties associated with a device.
type DeviceListProperties struct {
	// Typename the type name of a device
	Typename string `json:"typename"`
}

func (DeviceListProperties) Command() string {
	return "device-list-properties"
}

func (cmd DeviceListProperties) Execute(ctx context.Context, client api.Client) ([]ObjectPropertyInfo, error) {
	var ret []ObjectPropertyInfo

	return ret, client.Execute(ctx, "device-list-properties", cmd, &ret)
}

// DeviceAdd
//
// Add a device.
type DeviceAdd struct {
	// Driver the name of the new device's driver
	Driver string `json:"driver"`
	// Bus the device's parent bus (device tree path)
	Bus *string `json:"bus,omitempty"`
	// Id the device's ID, must be unique
	Id *string `json:"id,omitempty"`
}

func (DeviceAdd) Command() string {
	return "device_add"
}

func (cmd DeviceAdd) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "device_add", cmd, nil)
}

// DeviceDel
//
// Remove a device from a guest
type DeviceDel struct {
	// Id the device's ID or QOM path
	Id string `json:"id"`
}

func (DeviceDel) Command() string {
	return "device_del"
}

func (cmd DeviceDel) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "device_del", cmd, nil)
}

// DeviceDeletedEvent (DEVICE_DELETED)
//
// Emitted whenever the device removal completion is acknowledged by the guest. At this point, it's safe to reuse the specified device ID. Device removal can be initiated by the guest or by HMP/QMP commands.
type DeviceDeletedEvent struct {
	// Device the device's ID if it has one
	Device *string `json:"device,omitempty"`
	// Path the device's QOM path
	Path string `json:"path"`
}

func (DeviceDeletedEvent) Event() string {
	return "DEVICE_DELETED"
}

// DeviceUnplugGuestErrorEvent (DEVICE_UNPLUG_GUEST_ERROR)
//
// Emitted when a device hot unplug fails due to a guest reported error.
type DeviceUnplugGuestErrorEvent struct {
	// Device the device's ID if it has one
	Device *string `json:"device,omitempty"`
	// Path the device's QOM path
	Path string `json:"path"`
}

func (DeviceUnplugGuestErrorEvent) Event() string {
	return "DEVICE_UNPLUG_GUEST_ERROR"
}

// CpuS390Entitlement An enumeration of CPU entitlements that can be assumed by a virtual S390 CPU
type CpuS390Entitlement string

const (
	S390CpuEntitlementAuto   CpuS390Entitlement = "auto"
	S390CpuEntitlementLow    CpuS390Entitlement = "low"
	S390CpuEntitlementMedium CpuS390Entitlement = "medium"
	S390CpuEntitlementHigh   CpuS390Entitlement = "high"
)

// SysEmuTarget The comprehensive enumeration of QEMU system emulation ("softmmu") targets. Run "./configure --help" in the project root directory, and look for the \*-softmmu targets near the "--target-list" option. The individual target constants are not documented here, for the time being.
type SysEmuTarget string

const (
	SysEmuTargetAarch64 SysEmuTarget = "aarch64"
	SysEmuTargetAlpha   SysEmuTarget = "alpha"
	SysEmuTargetArm     SysEmuTarget = "arm"
	// SysEmuTargetAvr since 5.1
	SysEmuTargetAvr          SysEmuTarget = "avr"
	SysEmuTargetCris         SysEmuTarget = "cris"
	SysEmuTargetHppa         SysEmuTarget = "hppa"
	SysEmuTargetI386         SysEmuTarget = "i386"
	SysEmuTargetLoongarch64  SysEmuTarget = "loongarch64"
	SysEmuTargetM68k         SysEmuTarget = "m68k"
	SysEmuTargetMicroblaze   SysEmuTarget = "microblaze"
	SysEmuTargetMicroblazeel SysEmuTarget = "microblazeel"
	SysEmuTargetMips         SysEmuTarget = "mips"
	SysEmuTargetMips64       SysEmuTarget = "mips64"
	SysEmuTargetMips64el     SysEmuTarget = "mips64el"
	SysEmuTargetMipsel       SysEmuTarget = "mipsel"
	SysEmuTargetNios2        SysEmuTarget = "nios2"
	SysEmuTargetOr1k         SysEmuTarget = "or1k"
	SysEmuTargetPpc          SysEmuTarget = "ppc"
	SysEmuTargetPpc64        SysEmuTarget = "ppc64"
	SysEmuTargetRiscv32      SysEmuTarget = "riscv32"
	SysEmuTargetRiscv64      SysEmuTarget = "riscv64"
	// SysEmuTargetRx since 5.0
	SysEmuTargetRx       SysEmuTarget = "rx"
	SysEmuTargetS390x    SysEmuTarget = "s390x"
	SysEmuTargetSh4      SysEmuTarget = "sh4"
	SysEmuTargetSh4eb    SysEmuTarget = "sh4eb"
	SysEmuTargetSparc    SysEmuTarget = "sparc"
	SysEmuTargetSparc64  SysEmuTarget = "sparc64"
	SysEmuTargetTricore  SysEmuTarget = "tricore"
	SysEmuTargetX8664    SysEmuTarget = "x86_64"
	SysEmuTargetXtensa   SysEmuTarget = "xtensa"
	SysEmuTargetXtensaeb SysEmuTarget = "xtensaeb"
)

// CpuS390State An enumeration of cpu states that can be assumed by a virtual S390 CPU
type CpuS390State string

const (
	S390CpuStateUninitialized CpuS390State = "uninitialized"
	S390CpuStateStopped       CpuS390State = "stopped"
	S390CpuStateCheckStop     CpuS390State = "check-stop"
	S390CpuStateOperating     CpuS390State = "operating"
	S390CpuStateLoad          CpuS390State = "load"
)

// CpuInfoS390
//
// Additional information about a virtual S390 CPU
type CpuInfoS390 struct {
	// CpuState the virtual CPU's state
	CpuState CpuS390State `json:"cpu-state"`
	// Dedicated the virtual CPU's dedication (since 8.2)
	Dedicated *bool `json:"dedicated,omitempty"`
	// Entitlement the virtual CPU's entitlement (since 8.2)
	Entitlement *CpuS390Entitlement `json:"entitlement,omitempty"`
}

// CpuInfoFast
//
// Information about a virtual CPU
type CpuInfoFast struct {
	// Discriminator: target

	// CpuIndex index of the virtual CPU
	CpuIndex int64 `json:"cpu-index"`
	// QomPath path to the CPU object in the QOM tree
	QomPath string `json:"qom-path"`
	// ThreadId ID of the underlying host thread
	ThreadId int64 `json:"thread-id"`
	// Props properties associated with a virtual CPU, e.g. the socket id
	Props *CpuInstanceProperties `json:"props,omitempty"`
	// Target the QEMU system emulation target, which determines which additional fields will be listed (since 3.0)
	Target SysEmuTarget `json:"target"`

	S390x *CpuInfoS390 `json:"-"`
}

func (u CpuInfoFast) MarshalJSON() ([]byte, error) {
	switch u.Target {
	case "s390x":
		if u.S390x == nil {
			return nil, fmt.Errorf("expected S390x to be set")
		}

		return json.Marshal(struct {
			CpuIndex int64                  `json:"cpu-index"`
			QomPath  string                 `json:"qom-path"`
			ThreadId int64                  `json:"thread-id"`
			Props    *CpuInstanceProperties `json:"props,omitempty"`
			Target   SysEmuTarget           `json:"target"`
			*CpuInfoS390
		}{
			CpuIndex:    u.CpuIndex,
			QomPath:     u.QomPath,
			ThreadId:    u.ThreadId,
			Props:       u.Props,
			Target:      u.Target,
			CpuInfoS390: u.S390x,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// QueryCpusFast
//
// Returns information about all virtual CPUs.
type QueryCpusFast struct {
}

func (QueryCpusFast) Command() string {
	return "query-cpus-fast"
}

func (cmd QueryCpusFast) Execute(ctx context.Context, client api.Client) ([]CpuInfoFast, error) {
	var ret []CpuInfoFast

	return ret, client.Execute(ctx, "query-cpus-fast", cmd, &ret)
}

// MachineInfo
//
// Information describing a machine.
type MachineInfo struct {
	// Name the name of the machine
	Name string `json:"name"`
	// Alias an alias for the machine name
	Alias *string `json:"alias,omitempty"`
	// IsDefault whether the machine is default
	IsDefault *bool `json:"is-default,omitempty"`
	// CpuMax maximum number of CPUs supported by the machine type (since 1.5)
	CpuMax int64 `json:"cpu-max"`
	// HotpluggableCpus cpu hotplug via -device is supported (since 2.7)
	HotpluggableCpus bool `json:"hotpluggable-cpus"`
	// NumaMemSupported true if '-numa node,mem' option is supported by the machine type and false otherwise (since 4.1)
	NumaMemSupported bool `json:"numa-mem-supported"`
	// Deprecated if true, the machine type is deprecated and may be removed in future versions of QEMU according to the QEMU deprecation policy (since 4.1)
	Deprecated bool `json:"deprecated"`
	// DefaultCpuType default CPU model typename if none is requested via the -cpu argument. (since 4.2)
	DefaultCpuType *string `json:"default-cpu-type,omitempty"`
	// DefaultRamId the default ID of initial RAM memory backend (since 5.2)
	DefaultRamId *string `json:"default-ram-id,omitempty"`
	// Acpi machine type supports ACPI (since 8.0)
	Acpi bool `json:"acpi"`
}

// QueryMachines
//
// Return a list of supported machines
type QueryMachines struct {
}

func (QueryMachines) Command() string {
	return "query-machines"
}

func (cmd QueryMachines) Execute(ctx context.Context, client api.Client) ([]MachineInfo, error) {
	var ret []MachineInfo

	return ret, client.Execute(ctx, "query-machines", cmd, &ret)
}

// CurrentMachineParams
//
// Information describing the running machine parameters.
type CurrentMachineParams struct {
	// WakeupSuspendSupport true if the machine supports wake up from suspend
	WakeupSuspendSupport bool `json:"wakeup-suspend-support"`
}

// QueryCurrentMachine
//
// Return information on the current virtual machine.
type QueryCurrentMachine struct {
}

func (QueryCurrentMachine) Command() string {
	return "query-current-machine"
}

func (cmd QueryCurrentMachine) Execute(ctx context.Context, client api.Client) (CurrentMachineParams, error) {
	var ret CurrentMachineParams

	return ret, client.Execute(ctx, "query-current-machine", cmd, &ret)
}

// TargetInfo
//
// Information describing the QEMU target.
type TargetInfo struct {
	// Arch the target architecture
	Arch SysEmuTarget `json:"arch"`
}

// QueryTarget
//
// Return information about the target for this QEMU
type QueryTarget struct {
}

func (QueryTarget) Command() string {
	return "query-target"
}

func (cmd QueryTarget) Execute(ctx context.Context, client api.Client) (TargetInfo, error) {
	var ret TargetInfo

	return ret, client.Execute(ctx, "query-target", cmd, &ret)
}

// UuidInfo
//
// Guest UUID information (Universally Unique Identifier).
type UuidInfo struct {
	// Uuid the UUID of the guest
	Uuid string `json:"UUID"`
}

// QueryUuid
//
// Query the guest UUID information.
type QueryUuid struct {
}

func (QueryUuid) Command() string {
	return "query-uuid"
}

func (cmd QueryUuid) Execute(ctx context.Context, client api.Client) (UuidInfo, error) {
	var ret UuidInfo

	return ret, client.Execute(ctx, "query-uuid", cmd, &ret)
}

// GuidInfo
//
// GUID information.
type GuidInfo struct {
	// Guid the globally unique identifier
	Guid string `json:"guid"`
}

// QueryVmGenerationId
//
// Show Virtual Machine Generation ID
type QueryVmGenerationId struct {
}

func (QueryVmGenerationId) Command() string {
	return "query-vm-generation-id"
}

func (cmd QueryVmGenerationId) Execute(ctx context.Context, client api.Client) (GuidInfo, error) {
	var ret GuidInfo

	return ret, client.Execute(ctx, "query-vm-generation-id", cmd, &ret)
}

// SystemReset
//
// Performs a hard reset of a guest.
type SystemReset struct {
}

func (SystemReset) Command() string {
	return "system_reset"
}

func (cmd SystemReset) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "system_reset", cmd, nil)
}

// SystemPowerdown
//
// Requests that a guest perform a powerdown operation.
type SystemPowerdown struct {
}

func (SystemPowerdown) Command() string {
	return "system_powerdown"
}

func (cmd SystemPowerdown) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "system_powerdown", cmd, nil)
}

// SystemWakeup
//
// Wake up guest from suspend. If the guest has wake-up from suspend support enabled (wakeup-suspend-support flag from query-current-machine), wake-up guest from suspend if the guest is in SUSPENDED state. Return an error otherwise.
type SystemWakeup struct {
}

func (SystemWakeup) Command() string {
	return "system_wakeup"
}

func (cmd SystemWakeup) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "system_wakeup", cmd, nil)
}

// LostTickPolicy Policy for handling lost ticks in timer devices. Ticks end up getting lost when, for example, the guest is paused.
type LostTickPolicy string

const (
	// LostTickPolicyDiscard throw away the missed ticks and continue with future injection normally. The guest OS will see the timer jump ahead by a potentially quite significant amount all at once, as if the intervening chunk of time had simply not existed; needless to say, such a sudden jump can easily confuse a guest OS which is not specifically prepared to deal with it. Assuming the guest OS can deal correctly with the time jump, the time in the guest and in the host should now match.
	LostTickPolicyDiscard LostTickPolicy = "discard"
	// LostTickPolicyDelay continue to deliver ticks at the normal rate. The guest OS will not notice anything is amiss, as from its point of view time will have continued to flow normally. The time in the guest should now be behind the time in the host by exactly the amount of time during which ticks have been missed.
	LostTickPolicyDelay LostTickPolicy = "delay"
	// LostTickPolicySlew deliver ticks at a higher rate to catch up with the missed ticks. The guest OS will not notice anything is amiss, as from its point of view time will have continued to flow normally. Once the timer has managed to catch up with all the missing ticks, the time in the guest and in the host should match.
	LostTickPolicySlew LostTickPolicy = "slew"
)

// InjectNmi
//
// Injects a Non-Maskable Interrupt into the default CPU (x86/s390) or all CPUs (ppc64). The command fails when the guest doesn't support injecting.
type InjectNmi struct {
}

func (InjectNmi) Command() string {
	return "inject-nmi"
}

func (cmd InjectNmi) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "inject-nmi", cmd, nil)
}

// KvmInfo
//
// Information about support for KVM acceleration
type KvmInfo struct {
	// Enabled true if KVM acceleration is active
	Enabled bool `json:"enabled"`
	// Present true if KVM acceleration is built into this executable
	Present bool `json:"present"`
}

// QueryKvm
//
// Returns information about KVM acceleration
type QueryKvm struct {
}

func (QueryKvm) Command() string {
	return "query-kvm"
}

func (cmd QueryKvm) Execute(ctx context.Context, client api.Client) (KvmInfo, error) {
	var ret KvmInfo

	return ret, client.Execute(ctx, "query-kvm", cmd, &ret)
}

// NumaOptionsType
type NumaOptionsType string

const (
	// NumaOptionsTypeNode NUMA nodes configuration
	NumaOptionsTypeNode NumaOptionsType = "node"
	// NumaOptionsTypeDist NUMA distance configuration (since 2.10)
	NumaOptionsTypeDist NumaOptionsType = "dist"
	// NumaOptionsTypeCpu property based CPU(s) to node mapping (Since: 2.10)
	NumaOptionsTypeCpu NumaOptionsType = "cpu"
	// NumaOptionsTypeHmatLb memory latency and bandwidth information (Since: 5.0)
	NumaOptionsTypeHmatLb NumaOptionsType = "hmat-lb"
	// NumaOptionsTypeHmatCache memory side cache information (Since: 5.0)
	NumaOptionsTypeHmatCache NumaOptionsType = "hmat-cache"
)

// NumaOptions
//
// A discriminated record of NUMA options. (for OptsVisitor)
type NumaOptions struct {
	// Discriminator: type

	// Type NUMA option type
	Type NumaOptionsType `json:"type"`

	Node      *NumaNodeOptions      `json:"-"`
	Dist      *NumaDistOptions      `json:"-"`
	Cpu       *NumaCpuOptions       `json:"-"`
	HmatLb    *NumaHmatLBOptions    `json:"-"`
	HmatCache *NumaHmatCacheOptions `json:"-"`
}

func (u NumaOptions) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "node":
		if u.Node == nil {
			return nil, fmt.Errorf("expected Node to be set")
		}

		return json.Marshal(struct {
			Type NumaOptionsType `json:"type"`
			*NumaNodeOptions
		}{
			Type:            u.Type,
			NumaNodeOptions: u.Node,
		})
	case "dist":
		if u.Dist == nil {
			return nil, fmt.Errorf("expected Dist to be set")
		}

		return json.Marshal(struct {
			Type NumaOptionsType `json:"type"`
			*NumaDistOptions
		}{
			Type:            u.Type,
			NumaDistOptions: u.Dist,
		})
	case "cpu":
		if u.Cpu == nil {
			return nil, fmt.Errorf("expected Cpu to be set")
		}

		return json.Marshal(struct {
			Type NumaOptionsType `json:"type"`
			*NumaCpuOptions
		}{
			Type:           u.Type,
			NumaCpuOptions: u.Cpu,
		})
	case "hmat-lb":
		if u.HmatLb == nil {
			return nil, fmt.Errorf("expected HmatLb to be set")
		}

		return json.Marshal(struct {
			Type NumaOptionsType `json:"type"`
			*NumaHmatLBOptions
		}{
			Type:              u.Type,
			NumaHmatLBOptions: u.HmatLb,
		})
	case "hmat-cache":
		if u.HmatCache == nil {
			return nil, fmt.Errorf("expected HmatCache to be set")
		}

		return json.Marshal(struct {
			Type NumaOptionsType `json:"type"`
			*NumaHmatCacheOptions
		}{
			Type:                 u.Type,
			NumaHmatCacheOptions: u.HmatCache,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// NumaNodeOptions
//
// Create a guest NUMA node. (for OptsVisitor)
type NumaNodeOptions struct {
	// Nodeid NUMA node ID (increase by 1 from 0 if omitted)
	Nodeid *uint16 `json:"nodeid,omitempty"`
	// Cpus VCPUs belonging to this node (assign VCPUS round-robin if omitted)
	Cpus []uint16 `json:"cpus,omitempty"`
	// Mem memory size of this node; mutually exclusive with @memdev. Equally divide total memory among nodes if both @mem and @memdev are omitted.
	Mem *uint64 `json:"mem,omitempty"`
	// Memdev memory backend object. If specified for one node, it must be specified for all nodes.
	Memdev *string `json:"memdev,omitempty"`
	// Initiator defined in ACPI 6.3 Chapter 5.2.27.3 Table 5-145, points to the nodeid which has the memory controller responsible for this NUMA node. This field provides additional information as to the initiator node that is closest (as in directly attached) to this node, and therefore has the best performance (since 5.0)
	Initiator *uint16 `json:"initiator,omitempty"`
}

// NumaDistOptions
//
// Set the distance between 2 NUMA nodes.
type NumaDistOptions struct {
	// Src source NUMA node.
	Src uint16 `json:"src"`
	// Dst destination NUMA node.
	Dst uint16 `json:"dst"`
	// Val NUMA distance from source node to destination node. When a node is unreachable from another node, set the distance between them to 255.
	Val uint8 `json:"val"`
}

// CXLFixedMemoryWindowOptions
//
// Create a CXL Fixed Memory Window
type CXLFixedMemoryWindowOptions struct {
	// Size Size of the Fixed Memory Window in bytes. Must be a multiple of 256MiB.
	Size uint64 `json:"size"`
	// InterleaveGranularity Number of contiguous bytes for which accesses will go to a given interleave target. Accepted values [256, 512, 1k, 2k, 4k, 8k, 16k]
	InterleaveGranularity *uint64 `json:"interleave-granularity,omitempty"`
	// Targets Target root bridge IDs from -device ...,id=<ID> for each root bridge.
	Targets []string `json:"targets"`
}

// CXLFMWProperties
//
// List of CXL Fixed Memory Windows.
type CXLFMWProperties struct {
	// CxlFmw List of CXLFixedMemoryWindowOptions
	CxlFmw []CXLFixedMemoryWindowOptions `json:"cxl-fmw"`
}

// X86CPURegister32 A X86 32-bit register
type X86CPURegister32 string

const (
	X86CPURegister32Eax X86CPURegister32 = "EAX"
	X86CPURegister32Ebx X86CPURegister32 = "EBX"
	X86CPURegister32Ecx X86CPURegister32 = "ECX"
	X86CPURegister32Edx X86CPURegister32 = "EDX"
	X86CPURegister32Esp X86CPURegister32 = "ESP"
	X86CPURegister32Ebp X86CPURegister32 = "EBP"
	X86CPURegister32Esi X86CPURegister32 = "ESI"
	X86CPURegister32Edi X86CPURegister32 = "EDI"
)

// X86CPUFeatureWordInfo
//
// Information about a X86 CPU feature word
type X86CPUFeatureWordInfo struct {
	// CpuidInputEax Input EAX value for CPUID instruction for that feature word
	CpuidInputEax int64 `json:"cpuid-input-eax"`
	// CpuidInputEcx Input ECX value for CPUID instruction for that feature word
	CpuidInputEcx *int64 `json:"cpuid-input-ecx,omitempty"`
	// CpuidRegister Output register containing the feature bits
	CpuidRegister X86CPURegister32 `json:"cpuid-register"`
	// Features value of output register, containing the feature bits
	Features int64 `json:"features"`
}

// DummyForceArrays
//
// Not used by QMP; hack to let us use X86CPUFeatureWordInfoList internally
type DummyForceArrays struct {
	Unused []X86CPUFeatureWordInfo `json:"unused"`
}

// NumaCpuOptions
//
// Option "-numa cpu" overrides default cpu to node mapping. It accepts the same set of cpu properties as returned by query-hotpluggable-cpus[].props, where node-id could be used to override default node mapping.
type NumaCpuOptions struct {
	CpuInstanceProperties
}

// HmatLBMemoryHierarchy The memory hierarchy in the System Locality Latency and Bandwidth Information Structure of HMAT (Heterogeneous Memory Attribute Table) For more information about @HmatLBMemoryHierarchy, see chapter
type HmatLBMemoryHierarchy string

const (
	// HmatLBMemoryHierarchyMemory the structure represents the memory performance
	HmatLBMemoryHierarchyMemory HmatLBMemoryHierarchy = "memory"
	// HmatLBMemoryHierarchyFirstLevel first level of memory side cache
	HmatLBMemoryHierarchyFirstLevel HmatLBMemoryHierarchy = "first-level"
	// HmatLBMemoryHierarchySecondLevel second level of memory side cache
	HmatLBMemoryHierarchySecondLevel HmatLBMemoryHierarchy = "second-level"
	// HmatLBMemoryHierarchyThirdLevel third level of memory side cache
	HmatLBMemoryHierarchyThirdLevel HmatLBMemoryHierarchy = "third-level"
)

// HmatLBDataType Data type in the System Locality Latency and Bandwidth Information Structure of HMAT (Heterogeneous Memory Attribute Table)
type HmatLBDataType string

const (
	// HmatLBDataTypeAccessLatency access latency (nanoseconds)
	HmatLBDataTypeAccessLatency HmatLBDataType = "access-latency"
	// HmatLBDataTypeReadLatency read latency (nanoseconds)
	HmatLBDataTypeReadLatency HmatLBDataType = "read-latency"
	// HmatLBDataTypeWriteLatency write latency (nanoseconds)
	HmatLBDataTypeWriteLatency HmatLBDataType = "write-latency"
	// HmatLBDataTypeAccessBandwidth access bandwidth (Bytes per second)
	HmatLBDataTypeAccessBandwidth HmatLBDataType = "access-bandwidth"
	// HmatLBDataTypeReadBandwidth read bandwidth (Bytes per second)
	HmatLBDataTypeReadBandwidth HmatLBDataType = "read-bandwidth"
	// HmatLBDataTypeWriteBandwidth write bandwidth (Bytes per second)
	HmatLBDataTypeWriteBandwidth HmatLBDataType = "write-bandwidth"
)

// NumaHmatLBOptions
//
// Set the system locality latency and bandwidth information between Initiator and Target proximity Domains.
type NumaHmatLBOptions struct {
	// Initiator the Initiator Proximity Domain.
	Initiator uint16 `json:"initiator"`
	// Target the Target Proximity Domain.
	Target uint16 `json:"target"`
	// Hierarchy the Memory Hierarchy. Indicates the performance of memory or side cache.
	Hierarchy HmatLBMemoryHierarchy `json:"hierarchy"`
	// DataType presents the type of data, access/read/write latency or hit latency.
	DataType HmatLBDataType `json:"data-type"`
	// Latency the value of latency from @initiator to @target proximity domain, the latency unit is "ns(nanosecond)".
	Latency *uint64 `json:"latency,omitempty"`
	// Bandwidth the value of bandwidth between @initiator and @target proximity domain, the bandwidth unit is "Bytes per second".
	Bandwidth *uint64 `json:"bandwidth,omitempty"`
}

// HmatCacheAssociativity Cache associativity in the Memory Side Cache Information Structure of HMAT For more information of @HmatCacheAssociativity, see chapter
type HmatCacheAssociativity string

const (
	// HmatCacheAssociativityNone None (no memory side cache in this proximity domain, or cache associativity unknown)
	HmatCacheAssociativityNone HmatCacheAssociativity = "none"
	// HmatCacheAssociativityDirect Direct Mapped
	HmatCacheAssociativityDirect HmatCacheAssociativity = "direct"
	// HmatCacheAssociativityComplex Complex Cache Indexing (implementation specific)
	HmatCacheAssociativityComplex HmatCacheAssociativity = "complex"
)

// HmatCacheWritePolicy Cache write policy in the Memory Side Cache Information Structure of HMAT
type HmatCacheWritePolicy string

const (
	// HmatCacheWritePolicyNone None (no memory side cache in this proximity domain, or cache write policy unknown)
	HmatCacheWritePolicyNone HmatCacheWritePolicy = "none"
	// HmatCacheWritePolicyWriteBack Write Back (WB)
	HmatCacheWritePolicyWriteBack HmatCacheWritePolicy = "write-back"
	// HmatCacheWritePolicyWriteThrough Write Through (WT)
	HmatCacheWritePolicyWriteThrough HmatCacheWritePolicy = "write-through"
)

// NumaHmatCacheOptions
//
// Set the memory side cache information for a given memory domain.
type NumaHmatCacheOptions struct {
	// NodeId the memory proximity domain to which the memory belongs.
	NodeId uint32 `json:"node-id"`
	// Size the size of memory side cache in bytes.
	Size uint64 `json:"size"`
	// Level the cache level described in this structure.
	Level uint8 `json:"level"`
	// Associativity the cache associativity, none/direct-mapped/complex(complex cache indexing).
	Associativity HmatCacheAssociativity `json:"associativity"`
	// Policy the write policy, none/write-back/write-through.
	Policy HmatCacheWritePolicy `json:"policy"`
	// Line the cache Line size in bytes.
	Line uint16 `json:"line"`
}

// Memsave
//
// Save a portion of guest memory to a file.
type Memsave struct {
	// Val the virtual address of the guest to start from
	Val int64 `json:"val"`
	// Size the size of memory region to save
	Size int64 `json:"size"`
	// Filename the file to save the memory to as binary data
	Filename string `json:"filename"`
	// CpuIndex the index of the virtual CPU to use for translating the virtual address (defaults to CPU 0)
	CpuIndex *int64 `json:"cpu-index,omitempty"`
}

func (Memsave) Command() string {
	return "memsave"
}

func (cmd Memsave) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "memsave", cmd, nil)
}

// Pmemsave
//
// Save a portion of guest physical memory to a file.
type Pmemsave struct {
	// Val the physical address of the guest to start from
	Val int64 `json:"val"`
	// Size the size of memory region to save
	Size int64 `json:"size"`
	// Filename the file to save the memory to as binary data
	Filename string `json:"filename"`
}

func (Pmemsave) Command() string {
	return "pmemsave"
}

func (cmd Pmemsave) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "pmemsave", cmd, nil)
}

// Memdev
//
// Information about memory backend
type Memdev struct {
	// Id backend's ID if backend has 'id' property (since 2.9)
	Id *string `json:"id,omitempty"`
	// Size memory backend size
	Size uint64 `json:"size"`
	// Merge whether memory merge support is enabled
	Merge bool `json:"merge"`
	// Dump whether memory backend's memory is included in a core dump
	Dump bool `json:"dump"`
	// Prealloc whether memory was preallocated
	Prealloc bool `json:"prealloc"`
	// Share whether memory is private to QEMU or shared (since 6.1)
	Share bool `json:"share"`
	// Reserve whether swap space (or huge pages) was reserved if applicable. This corresponds to the user configuration and not the actual behavior implemented in the OS to perform the reservation. For example, Linux will never reserve swap space for shared file mappings. (since 6.1)
	Reserve *bool `json:"reserve,omitempty"`
	// HostNodes host nodes for its memory policy
	HostNodes []uint16 `json:"host-nodes"`
	// Policy memory policy of memory backend
	Policy HostMemPolicy `json:"policy"`
}

// QueryMemdev
//
// Returns information for all memory backends.
type QueryMemdev struct {
}

func (QueryMemdev) Command() string {
	return "query-memdev"
}

func (cmd QueryMemdev) Execute(ctx context.Context, client api.Client) ([]Memdev, error) {
	var ret []Memdev

	return ret, client.Execute(ctx, "query-memdev", cmd, &ret)
}

// CpuInstanceProperties
//
// List of properties to be used for hotplugging a CPU instance, it should be passed by management with device_add command when a CPU is being hotplugged. Which members are optional and which mandatory depends on the architecture and board.
type CpuInstanceProperties struct {
	// NodeId NUMA node ID the CPU belongs to
	NodeId *int64 `json:"node-id,omitempty"`
	// DrawerId drawer number within CPU topology the CPU belongs to (since 8.2)
	DrawerId *int64 `json:"drawer-id,omitempty"`
	// BookId book number within parent container the CPU belongs to (since 8.2)
	BookId *int64 `json:"book-id,omitempty"`
	// SocketId socket number within parent container the CPU belongs to
	SocketId *int64 `json:"socket-id,omitempty"`
	// DieId die number within the parent container the CPU belongs to (since 4.1)
	DieId *int64 `json:"die-id,omitempty"`
	// ClusterId cluster number within the parent container the CPU belongs to (since 7.1)
	ClusterId *int64 `json:"cluster-id,omitempty"`
	// CoreId core number within the parent container the CPU belongs to
	CoreId *int64 `json:"core-id,omitempty"`
	// ThreadId thread number within the core the CPU belongs to
	ThreadId *int64 `json:"thread-id,omitempty"`
}

// HotpluggableCPU
type HotpluggableCPU struct {
	// Type CPU object type for usage with device_add command
	Type string `json:"type"`
	// VcpusCount number of logical VCPU threads @HotpluggableCPU provides
	VcpusCount int64 `json:"vcpus-count"`
	// Props list of properties to be used for hotplugging CPU
	Props CpuInstanceProperties `json:"props"`
	// QomPath link to existing CPU object if CPU is present or omitted if CPU is not present.
	QomPath *string `json:"qom-path,omitempty"`
}

// QueryHotpluggableCpus
type QueryHotpluggableCpus struct {
}

func (QueryHotpluggableCpus) Command() string {
	return "query-hotpluggable-cpus"
}

func (cmd QueryHotpluggableCpus) Execute(ctx context.Context, client api.Client) ([]HotpluggableCPU, error) {
	var ret []HotpluggableCPU

	return ret, client.Execute(ctx, "query-hotpluggable-cpus", cmd, &ret)
}

// SetNumaNode
//
// Runtime equivalent of '-numa' CLI option, available at preconfigure stage to configure numa mapping before initializing machine.
type SetNumaNode struct {
	NumaOptions
}

func (SetNumaNode) Command() string {
	return "set-numa-node"
}

func (cmd SetNumaNode) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "set-numa-node", cmd, nil)
}

// Balloon
//
// Request the balloon driver to change its balloon size.
type Balloon struct {
	// Value the target logical size of the VM in bytes. We can deduce
	Value int64 `json:"value"`
}

func (Balloon) Command() string {
	return "balloon"
}

func (cmd Balloon) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "balloon", cmd, nil)
}

// BalloonInfo
//
// Information about the guest balloon device.
type BalloonInfo struct {
	// Actual the logical size of the VM in bytes Formula used: logical_vm_size = vm_ram_size - balloon_size
	Actual int64 `json:"actual"`
}

// QueryBalloon
//
// Return information about the balloon device.
type QueryBalloon struct {
}

func (QueryBalloon) Command() string {
	return "query-balloon"
}

func (cmd QueryBalloon) Execute(ctx context.Context, client api.Client) (BalloonInfo, error) {
	var ret BalloonInfo

	return ret, client.Execute(ctx, "query-balloon", cmd, &ret)
}

// BalloonChangeEvent (BALLOON_CHANGE)
//
// Emitted when the guest changes the actual BALLOON level. This value is equivalent to the @actual field return by the 'query-balloon' command
type BalloonChangeEvent struct {
	// Actual the logical size of the VM in bytes Formula used: logical_vm_size = vm_ram_size - balloon_size
	Actual int64 `json:"actual"`
}

func (BalloonChangeEvent) Event() string {
	return "BALLOON_CHANGE"
}

// HvBalloonInfo
//
// hv-balloon guest-provided memory status information.
type HvBalloonInfo struct {
	// Committed the amount of memory in use inside the guest plus the amount of the memory unusable inside the guest (ballooned out, offline, etc.)
	Committed uint64 `json:"committed"`
	// Available the amount of the memory inside the guest available for new allocations ("free")
	Available uint64 `json:"available"`
}

// QueryHvBalloonStatusReport
//
// Returns the hv-balloon driver data contained in the last received "STATUS" message from the guest.
type QueryHvBalloonStatusReport struct {
}

func (QueryHvBalloonStatusReport) Command() string {
	return "query-hv-balloon-status-report"
}

func (cmd QueryHvBalloonStatusReport) Execute(ctx context.Context, client api.Client) (HvBalloonInfo, error) {
	var ret HvBalloonInfo

	return ret, client.Execute(ctx, "query-hv-balloon-status-report", cmd, &ret)
}

// HvBalloonStatusReportEvent (HV_BALLOON_STATUS_REPORT)
//
// Emitted when the hv-balloon driver receives a "STATUS" message from the guest.
type HvBalloonStatusReportEvent struct {
}

func (HvBalloonStatusReportEvent) Event() string {
	return "HV_BALLOON_STATUS_REPORT"
}

// MemoryInfo
//
// Actual memory information in bytes.
type MemoryInfo struct {
	// BaseMemory size of "base" memory specified with command line option -m.
	BaseMemory uint64 `json:"base-memory"`
	// PluggedMemory size of memory that can be hot-unplugged. This field is omitted if target doesn't support memory hotplug (i.e. CONFIG_MEM_DEVICE not defined at build time).
	PluggedMemory *uint64 `json:"plugged-memory,omitempty"`
}

// QueryMemorySizeSummary
//
// Return the amount of initially allocated and present hotpluggable (if enabled) memory in bytes.
type QueryMemorySizeSummary struct {
}

func (QueryMemorySizeSummary) Command() string {
	return "query-memory-size-summary"
}

func (cmd QueryMemorySizeSummary) Execute(ctx context.Context, client api.Client) (MemoryInfo, error) {
	var ret MemoryInfo

	return ret, client.Execute(ctx, "query-memory-size-summary", cmd, &ret)
}

// PCDIMMDeviceInfo
//
// PCDIMMDevice state information
type PCDIMMDeviceInfo struct {
	// Id device's ID
	Id *string `json:"id,omitempty"`
	// Addr physical address, where device is mapped
	Addr int64 `json:"addr"`
	// Size size of memory that the device provides
	Size int64 `json:"size"`
	// Slot slot number at which device is plugged in
	Slot int64 `json:"slot"`
	// Node NUMA node number where device is plugged in
	Node int64 `json:"node"`
	// Memdev memory backend linked with device
	Memdev string `json:"memdev"`
	// Hotplugged true if device was hotplugged
	Hotplugged bool `json:"hotplugged"`
	// Hotpluggable true if device if could be added/removed while machine is running
	Hotpluggable bool `json:"hotpluggable"`
}

// VirtioPMEMDeviceInfo
//
// VirtioPMEM state information
type VirtioPMEMDeviceInfo struct {
	// Id device's ID
	Id *string `json:"id,omitempty"`
	// Memaddr physical address in memory, where device is mapped
	Memaddr uint64 `json:"memaddr"`
	// Size size of memory that the device provides
	Size uint64 `json:"size"`
	// Memdev memory backend linked with device
	Memdev string `json:"memdev"`
}

// VirtioMEMDeviceInfo
//
// VirtioMEMDevice state information
type VirtioMEMDeviceInfo struct {
	// Id device's ID
	Id *string `json:"id,omitempty"`
	// Memaddr physical address in memory, where device is mapped
	Memaddr uint64 `json:"memaddr"`
	// RequestedSize the user requested size of the device
	RequestedSize uint64 `json:"requested-size"`
	// Size the (current) size of memory that the device provides
	Size uint64 `json:"size"`
	// MaxSize the maximum size of memory that the device can provide
	MaxSize uint64 `json:"max-size"`
	// BlockSize the block size of memory that the device provides
	BlockSize uint64 `json:"block-size"`
	// Node NUMA node number where device is assigned to
	Node int64 `json:"node"`
	// Memdev memory backend linked with the region
	Memdev string `json:"memdev"`
}

// SgxEPCDeviceInfo
//
// Sgx EPC state information
type SgxEPCDeviceInfo struct {
	// Id device's ID
	Id *string `json:"id,omitempty"`
	// Memaddr physical address in memory, where device is mapped
	Memaddr uint64 `json:"memaddr"`
	// Size size of memory that the device provides
	Size uint64 `json:"size"`
	// Node the numa node (Since: 7.0)
	Node int64 `json:"node"`
	// Memdev memory backend linked with device
	Memdev string `json:"memdev"`
}

// HvBalloonDeviceInfo
//
// hv-balloon provided memory state information
type HvBalloonDeviceInfo struct {
	// Id device's ID
	Id *string `json:"id,omitempty"`
	// Memaddr physical address in memory, where device is mapped
	Memaddr *uint64 `json:"memaddr,omitempty"`
	// MaxSize the maximum size of memory that the device can provide
	MaxSize uint64 `json:"max-size"`
	// Memdev memory backend linked with device
	Memdev *string `json:"memdev,omitempty"`
}

// MemoryDeviceInfoKind
type MemoryDeviceInfoKind string

const (
	MemoryDeviceInfoKindDimm MemoryDeviceInfoKind = "dimm"
	// MemoryDeviceInfoKindNvdimm since 2.12
	MemoryDeviceInfoKindNvdimm MemoryDeviceInfoKind = "nvdimm"
	// MemoryDeviceInfoKindVirtioPmem since 4.1
	MemoryDeviceInfoKindVirtioPmem MemoryDeviceInfoKind = "virtio-pmem"
	// MemoryDeviceInfoKindVirtioMem since 5.1
	MemoryDeviceInfoKindVirtioMem MemoryDeviceInfoKind = "virtio-mem"
	// MemoryDeviceInfoKindSgxEpc since 6.2.
	MemoryDeviceInfoKindSgxEpc MemoryDeviceInfoKind = "sgx-epc"
	// MemoryDeviceInfoKindHvBalloon since 8.2.
	MemoryDeviceInfoKindHvBalloon MemoryDeviceInfoKind = "hv-balloon"
)

// PCDIMMDeviceInfoWrapper
type PCDIMMDeviceInfoWrapper struct {
	// Data PCDIMMDevice state information
	Data PCDIMMDeviceInfo `json:"data"`
}

// VirtioPMEMDeviceInfoWrapper
type VirtioPMEMDeviceInfoWrapper struct {
	// Data VirtioPMEM state information
	Data VirtioPMEMDeviceInfo `json:"data"`
}

// VirtioMEMDeviceInfoWrapper
type VirtioMEMDeviceInfoWrapper struct {
	// Data VirtioMEMDevice state information
	Data VirtioMEMDeviceInfo `json:"data"`
}

// SgxEPCDeviceInfoWrapper
type SgxEPCDeviceInfoWrapper struct {
	// Data Sgx EPC state information
	Data SgxEPCDeviceInfo `json:"data"`
}

// HvBalloonDeviceInfoWrapper
type HvBalloonDeviceInfoWrapper struct {
	// Data hv-balloon provided memory state information
	Data HvBalloonDeviceInfo `json:"data"`
}

// MemoryDeviceInfo
//
// Union containing information about a memory device
type MemoryDeviceInfo struct {
	// Discriminator: type

	// Type memory device type
	Type MemoryDeviceInfoKind `json:"type"`

	Dimm       *PCDIMMDeviceInfoWrapper     `json:"-"`
	Nvdimm     *PCDIMMDeviceInfoWrapper     `json:"-"`
	VirtioPmem *VirtioPMEMDeviceInfoWrapper `json:"-"`
	VirtioMem  *VirtioMEMDeviceInfoWrapper  `json:"-"`
	SgxEpc     *SgxEPCDeviceInfoWrapper     `json:"-"`
	HvBalloon  *HvBalloonDeviceInfoWrapper  `json:"-"`
}

func (u MemoryDeviceInfo) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "dimm":
		if u.Dimm == nil {
			return nil, fmt.Errorf("expected Dimm to be set")
		}

		return json.Marshal(struct {
			Type MemoryDeviceInfoKind `json:"type"`
			*PCDIMMDeviceInfoWrapper
		}{
			Type:                    u.Type,
			PCDIMMDeviceInfoWrapper: u.Dimm,
		})
	case "nvdimm":
		if u.Nvdimm == nil {
			return nil, fmt.Errorf("expected Nvdimm to be set")
		}

		return json.Marshal(struct {
			Type MemoryDeviceInfoKind `json:"type"`
			*PCDIMMDeviceInfoWrapper
		}{
			Type:                    u.Type,
			PCDIMMDeviceInfoWrapper: u.Nvdimm,
		})
	case "virtio-pmem":
		if u.VirtioPmem == nil {
			return nil, fmt.Errorf("expected VirtioPmem to be set")
		}

		return json.Marshal(struct {
			Type MemoryDeviceInfoKind `json:"type"`
			*VirtioPMEMDeviceInfoWrapper
		}{
			Type:                        u.Type,
			VirtioPMEMDeviceInfoWrapper: u.VirtioPmem,
		})
	case "virtio-mem":
		if u.VirtioMem == nil {
			return nil, fmt.Errorf("expected VirtioMem to be set")
		}

		return json.Marshal(struct {
			Type MemoryDeviceInfoKind `json:"type"`
			*VirtioMEMDeviceInfoWrapper
		}{
			Type:                       u.Type,
			VirtioMEMDeviceInfoWrapper: u.VirtioMem,
		})
	case "sgx-epc":
		if u.SgxEpc == nil {
			return nil, fmt.Errorf("expected SgxEpc to be set")
		}

		return json.Marshal(struct {
			Type MemoryDeviceInfoKind `json:"type"`
			*SgxEPCDeviceInfoWrapper
		}{
			Type:                    u.Type,
			SgxEPCDeviceInfoWrapper: u.SgxEpc,
		})
	case "hv-balloon":
		if u.HvBalloon == nil {
			return nil, fmt.Errorf("expected HvBalloon to be set")
		}

		return json.Marshal(struct {
			Type MemoryDeviceInfoKind `json:"type"`
			*HvBalloonDeviceInfoWrapper
		}{
			Type:                       u.Type,
			HvBalloonDeviceInfoWrapper: u.HvBalloon,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// SgxEPC
//
// Sgx EPC cmdline information
type SgxEPC struct {
	// Memdev memory backend linked with device
	Memdev string `json:"memdev"`
	// Node the numa node (Since: 7.0)
	Node int64 `json:"node"`
}

// SgxEPCProperties
//
// SGX properties of machine types.
type SgxEPCProperties struct {
	// SgxEpc list of ids of memory-backend-epc objects.
	SgxEpc []SgxEPC `json:"sgx-epc"`
}

// QueryMemoryDevices
//
// Lists available memory devices and their state
type QueryMemoryDevices struct {
}

func (QueryMemoryDevices) Command() string {
	return "query-memory-devices"
}

func (cmd QueryMemoryDevices) Execute(ctx context.Context, client api.Client) ([]MemoryDeviceInfo, error) {
	var ret []MemoryDeviceInfo

	return ret, client.Execute(ctx, "query-memory-devices", cmd, &ret)
}

// MemoryDeviceSizeChangeEvent (MEMORY_DEVICE_SIZE_CHANGE)
//
// Emitted when the size of a memory device changes. Only emitted for memory devices that can actually change the size (e.g., virtio-mem due to guest action).
type MemoryDeviceSizeChangeEvent struct {
	// Id device's ID
	Id *string `json:"id,omitempty"`
	// Size the new size of memory that the device provides
	Size uint64 `json:"size"`
	// QomPath path to the device object in the QOM tree (since 6.2)
	QomPath string `json:"qom-path"`
}

func (MemoryDeviceSizeChangeEvent) Event() string {
	return "MEMORY_DEVICE_SIZE_CHANGE"
}

// MemUnplugErrorEvent (MEM_UNPLUG_ERROR)
//
// Emitted when memory hot unplug error occurs.
type MemUnplugErrorEvent struct {
	// Device device name
	Device string `json:"device"`
	// Msg Informative message
	Msg string `json:"msg"`
}

func (MemUnplugErrorEvent) Event() string {
	return "MEM_UNPLUG_ERROR"
}

// BootConfiguration
//
// Schema for virtual machine boot configuration.
type BootConfiguration struct {
	// Order Boot order (a=floppy, c=hard disk, d=CD-ROM, n=network)
	Order *string `json:"order,omitempty"`
	// Once Boot order to apply on first boot
	Once *string `json:"once,omitempty"`
	// Menu Whether to show a boot menu
	Menu *bool `json:"menu,omitempty"`
	// Splash The name of the file to be passed to the firmware as logo picture, if @menu is true.
	Splash *string `json:"splash,omitempty"`
	// SplashTime How long to show the logo picture, in milliseconds
	SplashTime *int64 `json:"splash-time,omitempty"`
	// RebootTimeout Timeout before guest reboots after boot fails
	RebootTimeout *int64 `json:"reboot-timeout,omitempty"`
	// Strict Whether to attempt booting from devices not included in the boot order
	Strict *bool `json:"strict,omitempty"`
}

// SMPConfiguration
//
// Schema for CPU topology configuration. A missing value lets QEMU figure out a suitable value based on the ones that are provided. The members other than @cpus and @maxcpus define a topology of containers.
type SMPConfiguration struct {
	// Cpus number of virtual CPUs in the virtual machine
	Cpus *int64 `json:"cpus,omitempty"`
	// Drawers number of drawers in the CPU topology (since 8.2)
	Drawers *int64 `json:"drawers,omitempty"`
	// Books number of books in the CPU topology (since 8.2)
	Books *int64 `json:"books,omitempty"`
	// Sockets number of sockets per parent container
	Sockets *int64 `json:"sockets,omitempty"`
	// Dies number of dies per parent container
	Dies *int64 `json:"dies,omitempty"`
	// Clusters number of clusters per parent container (since 7.0)
	Clusters *int64 `json:"clusters,omitempty"`
	// Cores number of cores per parent container
	Cores *int64 `json:"cores,omitempty"`
	// Threads number of threads per core
	Threads *int64 `json:"threads,omitempty"`
	// Maxcpus maximum number of hotpluggable virtual CPUs in the virtual machine
	Maxcpus *int64 `json:"maxcpus,omitempty"`
}

// QueryIrq
//
// Query interrupt statistics
type QueryIrq struct {
}

func (QueryIrq) Command() string {
	return "x-query-irq"
}

func (cmd QueryIrq) Execute(ctx context.Context, client api.Client) (HumanReadableText, error) {
	var ret HumanReadableText

	return ret, client.Execute(ctx, "x-query-irq", cmd, &ret)
}

// QueryJit
//
// Query TCG compiler statistics
type QueryJit struct {
}

func (QueryJit) Command() string {
	return "x-query-jit"
}

func (cmd QueryJit) Execute(ctx context.Context, client api.Client) (HumanReadableText, error) {
	var ret HumanReadableText

	return ret, client.Execute(ctx, "x-query-jit", cmd, &ret)
}

// QueryNuma
//
// Query NUMA topology information
type QueryNuma struct {
}

func (QueryNuma) Command() string {
	return "x-query-numa"
}

func (cmd QueryNuma) Execute(ctx context.Context, client api.Client) (HumanReadableText, error) {
	var ret HumanReadableText

	return ret, client.Execute(ctx, "x-query-numa", cmd, &ret)
}

// QueryOpcount
//
// Query TCG opcode counters
type QueryOpcount struct {
}

func (QueryOpcount) Command() string {
	return "x-query-opcount"
}

func (cmd QueryOpcount) Execute(ctx context.Context, client api.Client) (HumanReadableText, error) {
	var ret HumanReadableText

	return ret, client.Execute(ctx, "x-query-opcount", cmd, &ret)
}

// QueryRamblock
//
// Query system ramblock information
type QueryRamblock struct {
}

func (QueryRamblock) Command() string {
	return "x-query-ramblock"
}

func (cmd QueryRamblock) Execute(ctx context.Context, client api.Client) (HumanReadableText, error) {
	var ret HumanReadableText

	return ret, client.Execute(ctx, "x-query-ramblock", cmd, &ret)
}

// QueryRdma
//
// Query RDMA state
type QueryRdma struct {
}

func (QueryRdma) Command() string {
	return "x-query-rdma"
}

func (cmd QueryRdma) Execute(ctx context.Context, client api.Client) (HumanReadableText, error) {
	var ret HumanReadableText

	return ret, client.Execute(ctx, "x-query-rdma", cmd, &ret)
}

// QueryRoms
//
// Query information on the registered ROMS
type QueryRoms struct {
}

func (QueryRoms) Command() string {
	return "x-query-roms"
}

func (cmd QueryRoms) Execute(ctx context.Context, client api.Client) (HumanReadableText, error) {
	var ret HumanReadableText

	return ret, client.Execute(ctx, "x-query-roms", cmd, &ret)
}

// QueryUsb
//
// Query information on the USB devices
type QueryUsb struct {
}

func (QueryUsb) Command() string {
	return "x-query-usb"
}

func (cmd QueryUsb) Execute(ctx context.Context, client api.Client) (HumanReadableText, error) {
	var ret HumanReadableText

	return ret, client.Execute(ctx, "x-query-usb", cmd, &ret)
}

// SmbiosEntryPointType
type SmbiosEntryPointType string

const (
	// SmbiosEntryPointType32 SMBIOS version 2.1 (32-bit) Entry Point
	SmbiosEntryPointType32 SmbiosEntryPointType = "32"
	// SmbiosEntryPointType64 SMBIOS version 3.0 (64-bit) Entry Point
	SmbiosEntryPointType64 SmbiosEntryPointType = "64"
)

// MemorySizeConfiguration
//
// Schema for memory size configuration.
type MemorySizeConfiguration struct {
	// Size memory size in bytes
	Size *uint64 `json:"size,omitempty"`
	// MaxSize maximum hotpluggable memory size in bytes
	MaxSize *uint64 `json:"max-size,omitempty"`
	// Slots number of available memory slots for hotplug
	Slots *uint64 `json:"slots,omitempty"`
}

// Dumpdtb
//
// Save the FDT in dtb format.
type Dumpdtb struct {
	// Filename name of the dtb file to be created
	Filename string `json:"filename"`
}

func (Dumpdtb) Command() string {
	return "dumpdtb"
}

func (cmd Dumpdtb) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "dumpdtb", cmd, nil)
}

// CpuModelInfo
//
// Virtual CPU model. A CPU model consists of the name of a CPU definition, to which delta changes are applied (e.g. features added/removed). Most magic values that an architecture might require should be hidden behind the name. However, if required, architectures can expose relevant properties.
type CpuModelInfo struct {
	// Name the name of the CPU definition the model is based on
	Name string `json:"name"`
	// Props a dictionary of QOM properties to be applied
	Props *any `json:"props,omitempty"`
}

// CpuModelExpansionType An enumeration of CPU model expansion types.
type CpuModelExpansionType string

const (
	// CpuModelExpansionTypeStatic Expand to a static CPU model, a combination of a static base model name and property delta changes. As the static base model will never change, the expanded CPU model will be the same, independent of QEMU version, machine type, machine options, and accelerator options. Therefore, the resulting model can be used by tooling without having to specify a compatibility machine - e.g. when displaying the "host" model. The @static CPU models are migration-safe.
	CpuModelExpansionTypeStatic CpuModelExpansionType = "static"
	// CpuModelExpansionTypeFull Expand all properties. The produced model is not guaranteed to be migration-safe, but allows tooling to get an insight and work with model details.
	CpuModelExpansionTypeFull CpuModelExpansionType = "full"
)

// CpuModelCompareResult An enumeration of CPU model comparison results. The result is usually calculated using e.g. CPU features or CPU generations.
type CpuModelCompareResult string

const (
	// CpuModelCompareResultIncompatible If model A is incompatible to model B, model A is not guaranteed to run where model B runs and the other way around.
	CpuModelCompareResultIncompatible CpuModelCompareResult = "incompatible"
	// CpuModelCompareResultIdentical If model A is identical to model B, model A is guaranteed to run where model B runs and the other way around.
	CpuModelCompareResultIdentical CpuModelCompareResult = "identical"
	// CpuModelCompareResultSuperset If model A is a superset of model B, model B is guaranteed to run where model A runs. There are no guarantees about the other way.
	CpuModelCompareResultSuperset CpuModelCompareResult = "superset"
	// CpuModelCompareResultSubset If model A is a subset of model B, model A is guaranteed to run where model B runs. There are no guarantees about the other way.
	CpuModelCompareResultSubset CpuModelCompareResult = "subset"
)

// CpuModelBaselineInfo
//
// The result of a CPU model baseline.
type CpuModelBaselineInfo struct {
	// Model the baselined CpuModelInfo.
	Model CpuModelInfo `json:"model"`
}

// CpuModelCompareInfo
//
// The result of a CPU model comparison.
type CpuModelCompareInfo struct {
	// Result The result of the compare operation.
	Result CpuModelCompareResult `json:"result"`
	// ResponsibleProperties List of properties that led to the comparison result not being identical. @responsible-properties is a list of QOM property names that led to both CPUs not being detected as identical. For identical models, this list is empty. If a QOM property is read-only, that means there's no known way to make the CPU models identical. If the special property name "type" is included, the models are by definition not identical and cannot be made identical.
	ResponsibleProperties []string `json:"responsible-properties"`
}

// QueryCpuModelComparison
//
// Compares two CPU models, returning how they compare in a specific configuration. The results indicates how both models compare regarding runnability. This result can be used by tooling to make decisions if a certain CPU model will run in a certain configuration or if a compatible CPU model has to be created by baselining. Usually, a CPU model is compared against the maximum possible CPU model of a certain configuration (e.g. the "host" model for KVM). If that CPU model is identical or a subset, it will run in that configuration.
type QueryCpuModelComparison struct {
	Modela CpuModelInfo `json:"modela"`
	Modelb CpuModelInfo `json:"modelb"`
}

func (QueryCpuModelComparison) Command() string {
	return "query-cpu-model-comparison"
}

func (cmd QueryCpuModelComparison) Execute(ctx context.Context, client api.Client) (CpuModelCompareInfo, error) {
	var ret CpuModelCompareInfo

	return ret, client.Execute(ctx, "query-cpu-model-comparison", cmd, &ret)
}

// QueryCpuModelBaseline
//
// Baseline two CPU models, creating a compatible third model. The created model will always be a static, migration-safe CPU model (see "static" CPU model expansion for details). This interface can be used by tooling to create a compatible CPU model out two CPU models. The created CPU model will be identical to or a subset of both CPU models when comparing them. Therefore, the created CPU model is guaranteed to run where the given CPU models run.
type QueryCpuModelBaseline struct {
	Modela CpuModelInfo `json:"modela"`
	Modelb CpuModelInfo `json:"modelb"`
}

func (QueryCpuModelBaseline) Command() string {
	return "query-cpu-model-baseline"
}

func (cmd QueryCpuModelBaseline) Execute(ctx context.Context, client api.Client) (CpuModelBaselineInfo, error) {
	var ret CpuModelBaselineInfo

	return ret, client.Execute(ctx, "query-cpu-model-baseline", cmd, &ret)
}

// CpuModelExpansionInfo
//
// The result of a cpu model expansion.
type CpuModelExpansionInfo struct {
	// Model the expanded CpuModelInfo.
	Model CpuModelInfo `json:"model"`
}

// QueryCpuModelExpansion
//
// Expands a given CPU model (or a combination of CPU model + additional options) to different granularities, allowing tooling to get an understanding what a specific CPU model looks like in QEMU under a certain configuration. This interface can be used to query the "host" CPU model.
type QueryCpuModelExpansion struct {
	Type  CpuModelExpansionType `json:"type"`
	Model CpuModelInfo          `json:"model"`
}

func (QueryCpuModelExpansion) Command() string {
	return "query-cpu-model-expansion"
}

func (cmd QueryCpuModelExpansion) Execute(ctx context.Context, client api.Client) (CpuModelExpansionInfo, error) {
	var ret CpuModelExpansionInfo

	return ret, client.Execute(ctx, "query-cpu-model-expansion", cmd, &ret)
}

// CpuDefinitionInfo
//
// Virtual CPU definition.
type CpuDefinitionInfo struct {
	// Name the name of the CPU definition
	Name string `json:"name"`
	// MigrationSafe whether a CPU definition can be safely used for migration in combination with a QEMU compatibility machine when migrating between different QEMU versions and between hosts with different sets of (hardware or software) capabilities. If not provided, information is not available and callers should not assume the CPU definition to be migration-safe. (since 2.8)
	MigrationSafe *bool `json:"migration-safe,omitempty"`
	// Static whether a CPU definition is static and will not change depending on QEMU version, machine type, machine options and accelerator options. A static model is always migration-safe. (since 2.8)
	Static bool `json:"static"`
	// UnavailableFeatures List of properties that prevent the CPU model from running in the current host. (since 2.8)
	UnavailableFeatures []string `json:"unavailable-features,omitempty"`
	// Typename Type name that can be used as argument to @device-list-properties, to introspect properties configurable using -cpu or -global. (since 2.9)
	Typename string `json:"typename"`
	// AliasOf Name of CPU model this model is an alias for. The target of the CPU model alias may change depending on the machine type. Management software is supposed to translate CPU model aliases in the VM configuration, because aliases may stop being migration-safe in the future (since 4.1)
	AliasOf *string `json:"alias-of,omitempty"`
	// Deprecated If true, this CPU model is deprecated and may be removed in in some future version of QEMU according to the QEMU deprecation policy. (since 5.2) @unavailable-features is a list of QOM property names that represent CPU model attributes that prevent the CPU from running. If the QOM property is read-only, that means there's no known way to make the CPU model run in the current host. Implementations that choose not to provide specific information return the property name "type". If the property is read-write, it means that it MAY be possible to run the CPU model in the current host if that property is changed. Management software can use it as hints to suggest or choose an alternative for the user, or just to generate meaningful error messages explaining why the CPU model can't be used. If @unavailable-features is an empty list, the CPU model is runnable using the current host and machine-type. If @unavailable-features is not present, runnability information for the CPU is not available.
	Deprecated bool `json:"deprecated"`
}

// QueryCpuDefinitions
//
// Return a list of supported virtual CPU definitions
type QueryCpuDefinitions struct {
}

func (QueryCpuDefinitions) Command() string {
	return "query-cpu-definitions"
}

func (cmd QueryCpuDefinitions) Execute(ctx context.Context, client api.Client) ([]CpuDefinitionInfo, error) {
	var ret []CpuDefinitionInfo

	return ret, client.Execute(ctx, "query-cpu-definitions", cmd, &ret)
}

// CpuS390Polarization An enumeration of CPU polarization that can be assumed by a virtual S390 CPU
type CpuS390Polarization string

const (
	S390CpuPolarizationHorizontal CpuS390Polarization = "horizontal"
	S390CpuPolarizationVertical   CpuS390Polarization = "vertical"
)

// SetCpuTopology
//
// Modify the topology by moving the CPU inside the topology tree, or by changing a modifier attribute of a CPU. Absent values will not be modified.
type SetCpuTopology struct {
	// CoreId the vCPU ID to be moved
	CoreId uint16 `json:"core-id"`
	// SocketId destination socket to move the vCPU to
	SocketId *uint16 `json:"socket-id,omitempty"`
	// BookId destination book to move the vCPU to
	BookId *uint16 `json:"book-id,omitempty"`
	// DrawerId destination drawer to move the vCPU to
	DrawerId *uint16 `json:"drawer-id,omitempty"`
	// Entitlement entitlement to set
	Entitlement *CpuS390Entitlement `json:"entitlement,omitempty"`
	// Dedicated whether the provisioning of real to virtual CPU is dedicated
	Dedicated *bool `json:"dedicated,omitempty"`
}

func (SetCpuTopology) Command() string {
	return "set-cpu-topology"
}

func (cmd SetCpuTopology) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "set-cpu-topology", cmd, nil)
}

// CpuPolarizationChangeEvent (CPU_POLARIZATION_CHANGE)
//
// Emitted when the guest asks to change the polarization. The guest can tell the host (via the PTF instruction) whether the CPUs should be provisioned using horizontal or vertical polarization. On horizontal polarization the host is expected to provision all vCPUs equally. On vertical polarization the host can provision each vCPU differently. The guest will get information on the details of the provisioning the next time it uses the STSI(15) instruction.
type CpuPolarizationChangeEvent struct {
	// Polarization polarization specified by the guest
	Polarization CpuS390Polarization `json:"polarization"`
}

func (CpuPolarizationChangeEvent) Event() string {
	return "CPU_POLARIZATION_CHANGE"
}

// CpuPolarizationInfo
//
// The result of a CPU polarization query.
type CpuPolarizationInfo struct {
	// Polarization the CPU polarization
	Polarization CpuS390Polarization `json:"polarization"`
}

// QueryS390xCpuPolarization
type QueryS390xCpuPolarization struct {
}

func (QueryS390xCpuPolarization) Command() string {
	return "query-s390x-cpu-polarization"
}

func (cmd QueryS390xCpuPolarization) Execute(ctx context.Context, client api.Client) (CpuPolarizationInfo, error) {
	var ret CpuPolarizationInfo

	return ret, client.Execute(ctx, "query-s390x-cpu-polarization", cmd, &ret)
}

// ReplayMode Mode of the replay subsystem.
type ReplayMode string

const (
	// ReplayModeNone normal execution mode. Replay or record are not enabled.
	ReplayModeNone ReplayMode = "none"
	// ReplayModeRecord record mode. All non-deterministic data is written into the replay log.
	ReplayModeRecord ReplayMode = "record"
	// ReplayModePlay replay mode. Non-deterministic data required for system execution is read from the log.
	ReplayModePlay ReplayMode = "play"
)

// ReplayInfo
//
// Record/replay information.
type ReplayInfo struct {
	// Mode current mode.
	Mode ReplayMode `json:"mode"`
	// Filename name of the record/replay log file. It is present only in record or replay modes, when the log is recorded or replayed.
	Filename *string `json:"filename,omitempty"`
	// Icount current number of executed instructions.
	Icount int64 `json:"icount"`
}

// QueryReplay
//
// Retrieve the record/replay information. It includes current instruction count which may be used for @replay-break and @replay-seek commands.
type QueryReplay struct {
}

func (QueryReplay) Command() string {
	return "query-replay"
}

func (cmd QueryReplay) Execute(ctx context.Context, client api.Client) (ReplayInfo, error) {
	var ret ReplayInfo

	return ret, client.Execute(ctx, "query-replay", cmd, &ret)
}

// ReplayBreak
//
// Set replay breakpoint at instruction count @icount. Execution stops when the specified instruction is reached. There can be at most one breakpoint. When breakpoint is set, any prior one is removed. The breakpoint may be set only in replay mode and only "in the future", i.e. at instruction counts greater than the current one. The current instruction count can be observed with @query-replay.
type ReplayBreak struct {
	// Icount instruction count to stop at
	Icount int64 `json:"icount"`
}

func (ReplayBreak) Command() string {
	return "replay-break"
}

func (cmd ReplayBreak) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "replay-break", cmd, nil)
}

// ReplayDeleteBreak
//
// Remove replay breakpoint which was set with @replay-break. The command is ignored when there are no replay breakpoints.
type ReplayDeleteBreak struct {
}

func (ReplayDeleteBreak) Command() string {
	return "replay-delete-break"
}

func (cmd ReplayDeleteBreak) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "replay-delete-break", cmd, nil)
}

// ReplaySeek
//
// Automatically proceed to the instruction count @icount, when replaying the execution. The command automatically loads nearest snapshot and replays the execution to find the desired instruction. When there is no preceding snapshot or the execution is not replayed, then the command fails. icount for the reference may be obtained with @query-replay command.
type ReplaySeek struct {
	// Icount target instruction count
	Icount int64 `json:"icount"`
}

func (ReplaySeek) Command() string {
	return "replay-seek"
}

func (cmd ReplaySeek) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "replay-seek", cmd, nil)
}

// YankInstanceType An enumeration of yank instance types. See @YankInstance for more information.
type YankInstanceType string

const (
	YankInstanceTypeBlockNode YankInstanceType = "block-node"
	YankInstanceTypeChardev   YankInstanceType = "chardev"
	YankInstanceTypeMigration YankInstanceType = "migration"
)

// YankInstanceBlockNode
//
// Specifies which block graph node to yank. See @YankInstance for more information.
type YankInstanceBlockNode struct {
	// NodeName the name of the block graph node
	NodeName string `json:"node-name"`
}

// YankInstanceChardev
//
// Specifies which character device to yank. See @YankInstance for more information.
type YankInstanceChardev struct {
	// Id the chardev's ID
	Id string `json:"id"`
}

// YankInstance
//
// A yank instance can be yanked with the @yank qmp command to recover from a hanging QEMU.
type YankInstance struct {
	// Discriminator: type

	// Type yank instance type
	Type YankInstanceType `json:"type"`

	BlockNode *YankInstanceBlockNode `json:"-"`
	Chardev   *YankInstanceChardev   `json:"-"`
}

func (u YankInstance) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "block-node":
		if u.BlockNode == nil {
			return nil, fmt.Errorf("expected BlockNode to be set")
		}

		return json.Marshal(struct {
			Type YankInstanceType `json:"type"`
			*YankInstanceBlockNode
		}{
			Type:                  u.Type,
			YankInstanceBlockNode: u.BlockNode,
		})
	case "chardev":
		if u.Chardev == nil {
			return nil, fmt.Errorf("expected Chardev to be set")
		}

		return json.Marshal(struct {
			Type YankInstanceType `json:"type"`
			*YankInstanceChardev
		}{
			Type:                u.Type,
			YankInstanceChardev: u.Chardev,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// Yank
//
// Try to recover from hanging QEMU by yanking the specified instances. See @YankInstance for more information.
type Yank struct {
	// Instances the instances to be yanked
	Instances []YankInstance `json:"instances"`
}

func (Yank) Command() string {
	return "yank"
}

func (cmd Yank) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "yank", cmd, nil)
}

// QueryYank
//
// Query yank instances. See @YankInstance for more information.
type QueryYank struct {
}

func (QueryYank) Command() string {
	return "query-yank"
}

func (cmd QueryYank) Execute(ctx context.Context, client api.Client) ([]YankInstance, error) {
	var ret []YankInstance

	return ret, client.Execute(ctx, "query-yank", cmd, &ret)
}

// AddClient
//
// Allow client connections for VNC, Spice and socket based character devices to be passed in to QEMU via SCM_RIGHTS. If the FD associated with @fdname is not a socket, the command will fail and the FD will be closed.
type AddClient struct {
	// Protocol protocol name. Valid names are "vnc", "spice", "@dbus-display" or the name of a character device (e.g. from -chardev id=XXXX)
	Protocol string `json:"protocol"`
	// Fdname file descriptor name previously passed via 'getfd' command
	Fdname string `json:"fdname"`
	// Skipauth whether to skip authentication. Only applies to "vnc" and "spice" protocols
	Skipauth *bool `json:"skipauth,omitempty"`
	// Tls whether to perform TLS. Only applies to the "spice" protocol
	Tls *bool `json:"tls,omitempty"`
}

func (AddClient) Command() string {
	return "add_client"
}

func (cmd AddClient) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "add_client", cmd, nil)
}

// NameInfo
//
// Guest name information.
type NameInfo struct {
	// Name The name of the guest
	Name *string `json:"name,omitempty"`
}

// QueryName
//
// Return the name information of a guest.
type QueryName struct {
}

func (QueryName) Command() string {
	return "query-name"
}

func (cmd QueryName) Execute(ctx context.Context, client api.Client) (NameInfo, error) {
	var ret NameInfo

	return ret, client.Execute(ctx, "query-name", cmd, &ret)
}

// IOThreadInfo
//
// Information about an iothread
type IOThreadInfo struct {
	// Id the identifier of the iothread
	Id string `json:"id"`
	// ThreadId ID of the underlying host thread
	ThreadId int64 `json:"thread-id"`
	// PollMaxNs maximum polling time in ns, 0 means polling is disabled (since 2.9)
	PollMaxNs int64 `json:"poll-max-ns"`
	// PollGrow how many ns will be added to polling time, 0 means that it's not configured (since 2.9)
	PollGrow int64 `json:"poll-grow"`
	// PollShrink how many ns will be removed from polling time, 0 means that it's not configured (since 2.9)
	PollShrink int64 `json:"poll-shrink"`
	// AioMaxBatch maximum number of requests in a batch for the AIO engine, 0 means that the engine will use its default (since 6.1)
	AioMaxBatch int64 `json:"aio-max-batch"`
}

// QueryIothreads
//
// Returns a list of information about each iothread.
type QueryIothreads struct {
}

func (QueryIothreads) Command() string {
	return "query-iothreads"
}

func (cmd QueryIothreads) Execute(ctx context.Context, client api.Client) ([]IOThreadInfo, error) {
	var ret []IOThreadInfo

	return ret, client.Execute(ctx, "query-iothreads", cmd, &ret)
}

// Stop
//
// Stop guest VM execution.
type Stop struct {
}

func (Stop) Command() string {
	return "stop"
}

func (cmd Stop) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "stop", cmd, nil)
}

// Cont
//
// Resume guest VM execution.
type Cont struct {
}

func (Cont) Command() string {
	return "cont"
}

func (cmd Cont) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "cont", cmd, nil)
}

// ExitPreconfig
//
// Exit from "preconfig" state This command makes QEMU exit the preconfig state and proceed with VM initialization using configuration data provided on the command line and via the QMP monitor during the preconfig state. The command is only available during the preconfig state (i.e. when the --preconfig command line option was in use).
type ExitPreconfig struct {
}

func (ExitPreconfig) Command() string {
	return "x-exit-preconfig"
}

func (cmd ExitPreconfig) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "x-exit-preconfig", cmd, nil)
}

// HumanMonitorCommand
//
// Execute a command on the human monitor and return the output.
type HumanMonitorCommand struct {
	// CommandLine the command to execute in the human monitor
	CommandLine string `json:"command-line"`
	// CpuIndex The CPU to use for commands that require an implicit CPU
	CpuIndex *int64 `json:"cpu-index,omitempty"`
}

func (HumanMonitorCommand) Command() string {
	return "human-monitor-command"
}

func (cmd HumanMonitorCommand) Execute(ctx context.Context, client api.Client) (string, error) {
	var ret string

	return ret, client.Execute(ctx, "human-monitor-command", cmd, &ret)
}

// Getfd
//
// Receive a file descriptor via SCM rights and assign it a name
type Getfd struct {
	// Fdname file descriptor name
	Fdname string `json:"fdname"`
}

func (Getfd) Command() string {
	return "getfd"
}

func (cmd Getfd) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "getfd", cmd, nil)
}

// GetWin32Socket
//
// Add a socket that was duplicated to QEMU process with WSADuplicateSocketW() via WSASocket() & WSAPROTOCOL_INFOW structure and assign it a name (the SOCKET is associated with a CRT file descriptor)
type GetWin32Socket struct {
	// Info the WSAPROTOCOL_INFOW structure (encoded in base64)
	Info string `json:"info"`
	// Fdname file descriptor name
	Fdname string `json:"fdname"`
}

func (GetWin32Socket) Command() string {
	return "get-win32-socket"
}

func (cmd GetWin32Socket) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "get-win32-socket", cmd, nil)
}

// Closefd
//
// Close a file descriptor previously passed via SCM rights
type Closefd struct {
	// Fdname file descriptor name
	Fdname string `json:"fdname"`
}

func (Closefd) Command() string {
	return "closefd"
}

func (cmd Closefd) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "closefd", cmd, nil)
}

// AddfdInfo
//
// Information about a file descriptor that was added to an fd set.
type AddfdInfo struct {
	// FdsetId The ID of the fd set that @fd was added to.
	FdsetId int64 `json:"fdset-id"`
	// Fd The file descriptor that was received via SCM rights and added to the fd set.
	Fd int64 `json:"fd"`
}

// AddFd
//
// Add a file descriptor, that was passed via SCM rights, to an fd set.
type AddFd struct {
	// FdsetId The ID of the fd set to add the file descriptor to.
	FdsetId *int64 `json:"fdset-id,omitempty"`
	// Opaque A free-form string that can be used to describe the fd.
	Opaque *string `json:"opaque,omitempty"`
}

func (AddFd) Command() string {
	return "add-fd"
}

func (cmd AddFd) Execute(ctx context.Context, client api.Client) (AddfdInfo, error) {
	var ret AddfdInfo

	return ret, client.Execute(ctx, "add-fd", cmd, &ret)
}

// RemoveFd
//
// Remove a file descriptor from an fd set.
type RemoveFd struct {
	// FdsetId The ID of the fd set that the file descriptor belongs to.
	FdsetId int64 `json:"fdset-id"`
	// Fd The file descriptor that is to be removed.
	Fd *int64 `json:"fd,omitempty"`
}

func (RemoveFd) Command() string {
	return "remove-fd"
}

func (cmd RemoveFd) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "remove-fd", cmd, nil)
}

// FdsetFdInfo
//
// Information about a file descriptor that belongs to an fd set.
type FdsetFdInfo struct {
	// Fd The file descriptor value.
	Fd int64 `json:"fd"`
	// Opaque A free-form string that can be used to describe the fd.
	Opaque *string `json:"opaque,omitempty"`
}

// FdsetInfo
//
// Information about an fd set.
type FdsetInfo struct {
	// FdsetId The ID of the fd set.
	FdsetId int64 `json:"fdset-id"`
	// Fds A list of file descriptors that belong to this fd set.
	Fds []FdsetFdInfo `json:"fds"`
}

// QueryFdsets
//
// Return information describing all fd sets.
type QueryFdsets struct {
}

func (QueryFdsets) Command() string {
	return "query-fdsets"
}

func (cmd QueryFdsets) Execute(ctx context.Context, client api.Client) ([]FdsetInfo, error) {
	var ret []FdsetInfo

	return ret, client.Execute(ctx, "query-fdsets", cmd, &ret)
}

// CommandLineParameterType Possible types for an option parameter.
type CommandLineParameterType string

const (
	// CommandLineParameterTypeString accepts a character string
	CommandLineParameterTypeString CommandLineParameterType = "string"
	// CommandLineParameterTypeBoolean accepts "on" or "off"
	CommandLineParameterTypeBoolean CommandLineParameterType = "boolean"
	// CommandLineParameterTypeNumber accepts a number
	CommandLineParameterTypeNumber CommandLineParameterType = "number"
	// CommandLineParameterTypeSize accepts a number followed by an optional suffix (K)ilo, (M)ega, (G)iga, (T)era
	CommandLineParameterTypeSize CommandLineParameterType = "size"
)

// CommandLineParameterInfo
//
// Details about a single parameter of a command line option.
type CommandLineParameterInfo struct {
	// Name parameter name
	Name string `json:"name"`
	// Type parameter @CommandLineParameterType
	Type CommandLineParameterType `json:"type"`
	// Help human readable text string, not suitable for parsing.
	Help *string `json:"help,omitempty"`
	// Default default value string (since 2.1)
	Default *string `json:"default,omitempty"`
}

// CommandLineOptionInfo
//
// Details about a command line option, including its list of parameter details
type CommandLineOptionInfo struct {
	// Option option name
	Option string `json:"option"`
	// Parameters an array of @CommandLineParameterInfo
	Parameters []CommandLineParameterInfo `json:"parameters"`
}

// QueryCommandLineOptions
//
// Query command line option schema.
type QueryCommandLineOptions struct {
	// Option option name
	Option *string `json:"option,omitempty"`
}

func (QueryCommandLineOptions) Command() string {
	return "query-command-line-options"
}

func (cmd QueryCommandLineOptions) Execute(ctx context.Context, client api.Client) ([]CommandLineOptionInfo, error) {
	var ret []CommandLineOptionInfo

	return ret, client.Execute(ctx, "query-command-line-options", cmd, &ret)
}

// RtcChangeEvent (RTC_CHANGE)
//
// Emitted when the guest changes the RTC time.
type RtcChangeEvent struct {
	// Offset offset in seconds between base RTC clock (as specified by -rtc base), and new RTC clock value
	Offset int64 `json:"offset"`
	// QomPath path to the RTC object in the QOM tree
	QomPath string `json:"qom-path"`
}

func (RtcChangeEvent) Event() string {
	return "RTC_CHANGE"
}

// VfuClientHangupEvent (VFU_CLIENT_HANGUP)
//
// Emitted when the client of a TYPE_VFIO_USER_SERVER closes the communication channel
type VfuClientHangupEvent struct {
	// VfuId ID of the TYPE_VFIO_USER_SERVER object. It is the last component of @vfu-qom-path referenced below
	VfuId string `json:"vfu-id"`
	// VfuQomPath path to the TYPE_VFIO_USER_SERVER object in the QOM tree
	VfuQomPath string `json:"vfu-qom-path"`
	// DevId ID of attached PCI device
	DevId string `json:"dev-id"`
	// DevQomPath path to attached PCI device in the QOM tree
	DevQomPath string `json:"dev-qom-path"`
}

func (VfuClientHangupEvent) Event() string {
	return "VFU_CLIENT_HANGUP"
}

// RtcResetReinjection
//
// This command will reset the RTC interrupt reinjection backlog. Can be used if another mechanism to synchronize guest time is in effect, for example QEMU guest agent's guest-set-time command.
type RtcResetReinjection struct {
}

func (RtcResetReinjection) Command() string {
	return "rtc-reset-reinjection"
}

func (cmd RtcResetReinjection) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "rtc-reset-reinjection", cmd, nil)
}

// SevState An enumeration of SEV state information used during @query-sev.
type SevState string

const (
	// SevStateUninit The guest is uninitialized.
	SevStateUninit SevState = "uninit"
	// SevStateLaunchUpdate The guest is currently being launched; plaintext data and register state is being imported.
	SevStateLaunchUpdate SevState = "launch-update"
	// SevStateLaunchSecret The guest is currently being launched; ciphertext data is being imported.
	SevStateLaunchSecret SevState = "launch-secret"
	// SevStateRunning The guest is fully launched or migrated in.
	SevStateRunning SevState = "running"
	// SevStateSendUpdate The guest is currently being migrated out to another machine.
	SevStateSendUpdate SevState = "send-update"
	// SevStateReceiveUpdate The guest is currently being migrated from another machine.
	SevStateReceiveUpdate SevState = "receive-update"
)

// SevInfo
//
// Information about Secure Encrypted Virtualization (SEV) support
type SevInfo struct {
	// Enabled true if SEV is active
	Enabled bool `json:"enabled"`
	// ApiMajor SEV API major version
	ApiMajor uint8 `json:"api-major"`
	// ApiMinor SEV API minor version
	ApiMinor uint8 `json:"api-minor"`
	// BuildId SEV FW build id
	BuildId uint8 `json:"build-id"`
	// Policy SEV policy value
	Policy uint32 `json:"policy"`
	// State SEV guest state
	State SevState `json:"state"`
	// Handle SEV firmware handle
	Handle uint32 `json:"handle"`
}

// QuerySev
//
// Returns information about SEV
type QuerySev struct {
}

func (QuerySev) Command() string {
	return "query-sev"
}

func (cmd QuerySev) Execute(ctx context.Context, client api.Client) (SevInfo, error) {
	var ret SevInfo

	return ret, client.Execute(ctx, "query-sev", cmd, &ret)
}

// SevLaunchMeasureInfo
//
// SEV Guest Launch measurement information
type SevLaunchMeasureInfo struct {
	// Data the measurement value encoded in base64
	Data string `json:"data"`
}

// QuerySevLaunchMeasure
//
// Query the SEV guest launch information.
type QuerySevLaunchMeasure struct {
}

func (QuerySevLaunchMeasure) Command() string {
	return "query-sev-launch-measure"
}

func (cmd QuerySevLaunchMeasure) Execute(ctx context.Context, client api.Client) (SevLaunchMeasureInfo, error) {
	var ret SevLaunchMeasureInfo

	return ret, client.Execute(ctx, "query-sev-launch-measure", cmd, &ret)
}

// SevCapability
//
// The struct describes capability for a Secure Encrypted Virtualization feature.
type SevCapability struct {
	// Pdh Platform Diffie-Hellman key (base64 encoded)
	Pdh string `json:"pdh"`
	// CertChain PDH certificate chain (base64 encoded)
	CertChain string `json:"cert-chain"`
	// Cpu0Id Unique ID of CPU0 (base64 encoded) (since 7.1)
	Cpu0Id string `json:"cpu0-id"`
	// Cbitpos C-bit location in page table entry
	Cbitpos int64 `json:"cbitpos"`
	// ReducedPhysBits Number of physical Address bit reduction when SEV is enabled
	ReducedPhysBits int64 `json:"reduced-phys-bits"`
}

// QuerySevCapabilities
//
// This command is used to get the SEV capabilities, and is supported on AMD X86 platforms only.
type QuerySevCapabilities struct {
}

func (QuerySevCapabilities) Command() string {
	return "query-sev-capabilities"
}

func (cmd QuerySevCapabilities) Execute(ctx context.Context, client api.Client) (SevCapability, error) {
	var ret SevCapability

	return ret, client.Execute(ctx, "query-sev-capabilities", cmd, &ret)
}

// SevInjectLaunchSecret
//
// This command injects a secret blob into memory of SEV guest.
type SevInjectLaunchSecret struct {
	// PacketHeader the launch secret packet header encoded in base64
	PacketHeader string `json:"packet-header"`
	// Secret the launch secret data to be injected encoded in base64
	Secret string `json:"secret"`
	// Gpa the guest physical address where secret will be injected.
	Gpa *uint64 `json:"gpa,omitempty"`
}

func (SevInjectLaunchSecret) Command() string {
	return "sev-inject-launch-secret"
}

func (cmd SevInjectLaunchSecret) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "sev-inject-launch-secret", cmd, nil)
}

// SevAttestationReport
//
// The struct describes attestation report for a Secure Encrypted Virtualization feature.
type SevAttestationReport struct {
	// Data guest attestation report (base64 encoded)
	Data string `json:"data"`
}

// QuerySevAttestationReport
//
// This command is used to get the SEV attestation report, and is supported on AMD X86 platforms only.
type QuerySevAttestationReport struct {
	// Mnonce a random 16 bytes value encoded in base64 (it will be included in report)
	Mnonce string `json:"mnonce"`
}

func (QuerySevAttestationReport) Command() string {
	return "query-sev-attestation-report"
}

func (cmd QuerySevAttestationReport) Execute(ctx context.Context, client api.Client) (SevAttestationReport, error) {
	var ret SevAttestationReport

	return ret, client.Execute(ctx, "query-sev-attestation-report", cmd, &ret)
}

// DumpSkeys
//
// Dump guest's storage keys
type DumpSkeys struct {
	// Filename the path to the file to dump to
	Filename string `json:"filename"`
}

func (DumpSkeys) Command() string {
	return "dump-skeys"
}

func (cmd DumpSkeys) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "dump-skeys", cmd, nil)
}

// GICCapability
//
// The struct describes capability for a specific GIC (Generic Interrupt Controller) version. These bits are not only decided by QEMU/KVM software version, but also decided by the hardware that the program is running upon.
type GICCapability struct {
	// Version version of GIC to be described. Currently, only 2 and 3 are supported.
	Version int64 `json:"version"`
	// Emulated whether current QEMU/hardware supports emulated GIC device in user space.
	Emulated bool `json:"emulated"`
	// Kernel whether current QEMU/hardware supports hardware accelerated GIC device in kernel.
	Kernel bool `json:"kernel"`
}

// QueryGicCapabilities
//
// This command is ARM-only. It will return a list of GICCapability objects that describe its capability bits.
type QueryGicCapabilities struct {
}

func (QueryGicCapabilities) Command() string {
	return "query-gic-capabilities"
}

func (cmd QueryGicCapabilities) Execute(ctx context.Context, client api.Client) ([]GICCapability, error) {
	var ret []GICCapability

	return ret, client.Execute(ctx, "query-gic-capabilities", cmd, &ret)
}

// SGXEPCSection
//
// Information about intel SGX EPC section info
type SGXEPCSection struct {
	// Node the numa node
	Node int64 `json:"node"`
	// Size the size of EPC section
	Size uint64 `json:"size"`
}

// SGXInfo
//
// Information about intel Safe Guard eXtension (SGX) support
type SGXInfo struct {
	// Sgx true if SGX is supported
	Sgx bool `json:"sgx"`
	// Sgx1 true if SGX1 is supported
	Sgx1 bool `json:"sgx1"`
	// Sgx2 true if SGX2 is supported
	Sgx2 bool `json:"sgx2"`
	// Flc true if FLC is supported
	Flc bool `json:"flc"`
	// Sections The EPC sections info for guest (Since: 7.0)
	Sections []SGXEPCSection `json:"sections"`
}

// QuerySgx
//
// Returns information about SGX
type QuerySgx struct {
}

func (QuerySgx) Command() string {
	return "query-sgx"
}

func (cmd QuerySgx) Execute(ctx context.Context, client api.Client) (SGXInfo, error) {
	var ret SGXInfo

	return ret, client.Execute(ctx, "query-sgx", cmd, &ret)
}

// QuerySgxCapabilities
//
// Returns information from host SGX capabilities
type QuerySgxCapabilities struct {
}

func (QuerySgxCapabilities) Command() string {
	return "query-sgx-capabilities"
}

func (cmd QuerySgxCapabilities) Execute(ctx context.Context, client api.Client) (SGXInfo, error) {
	var ret SGXInfo

	return ret, client.Execute(ctx, "query-sgx-capabilities", cmd, &ret)
}

// EvtchnPortType An enumeration of Xen event channel port types.
type EvtchnPortType string

const (
	// EvtchnPortTypeClosed The port is unused.
	EvtchnPortTypeClosed EvtchnPortType = "closed"
	// EvtchnPortTypeUnbound The port is allocated and ready to be bound.
	EvtchnPortTypeUnbound EvtchnPortType = "unbound"
	// EvtchnPortTypeInterdomain The port is connected as an interdomain interrupt.
	EvtchnPortTypeInterdomain EvtchnPortType = "interdomain"
	// EvtchnPortTypePirq The port is bound to a physical IRQ (PIRQ).
	EvtchnPortTypePirq EvtchnPortType = "pirq"
	// EvtchnPortTypeVirq The port is bound to a virtual IRQ (VIRQ).
	EvtchnPortTypeVirq EvtchnPortType = "virq"
	// EvtchnPortTypeIpi The post is an inter-processor interrupt (IPI).
	EvtchnPortTypeIpi EvtchnPortType = "ipi"
)

// EvtchnInfo
//
// Information about a Xen event channel port
type EvtchnInfo struct {
	// Port the port number
	Port uint16 `json:"port"`
	// Vcpu target vCPU for this port
	Vcpu uint32 `json:"vcpu"`
	// Type the port type
	Type EvtchnPortType `json:"type"`
	// RemoteDomain remote domain for interdomain ports
	RemoteDomain string `json:"remote-domain"`
	// Target remote port ID, or virq/pirq number
	Target uint16 `json:"target"`
	// Pending port is currently active pending delivery
	Pending bool `json:"pending"`
	// Masked port is masked
	Masked bool `json:"masked"`
}

// XenEventList
//
// Query the Xen event channels opened by the guest.
type XenEventList struct {
}

func (XenEventList) Command() string {
	return "xen-event-list"
}

func (cmd XenEventList) Execute(ctx context.Context, client api.Client) ([]EvtchnInfo, error) {
	var ret []EvtchnInfo

	return ret, client.Execute(ctx, "xen-event-list", cmd, &ret)
}

// XenEventInject
//
// Inject a Xen event channel port (interrupt) to the guest.
type XenEventInject struct {
	// Port The port number
	Port uint32 `json:"port"`
}

func (XenEventInject) Command() string {
	return "xen-event-inject"
}

func (cmd XenEventInject) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "xen-event-inject", cmd, nil)
}

// AudiodevPerDirectionOptions
//
// General audio backend options that are used for both playback and recording.
type AudiodevPerDirectionOptions struct {
	// MixingEngine use QEMU's mixing engine to mix all streams inside QEMU and convert audio formats when not supported by the backend. When set to off, fixed-settings must be also off (default on, since 4.2)
	MixingEngine *bool `json:"mixing-engine,omitempty"`
	// FixedSettings use fixed settings for host input/output. When off, frequency, channels and format must not be specified (default true)
	FixedSettings *bool `json:"fixed-settings,omitempty"`
	// Frequency frequency to use when using fixed settings (default 44100)
	Frequency *uint32 `json:"frequency,omitempty"`
	// Channels number of channels when using fixed settings (default 2)
	Channels *uint32 `json:"channels,omitempty"`
	// Voices number of voices to use (default 1)
	Voices *uint32 `json:"voices,omitempty"`
	// Format sample format to use when using fixed settings (default s16)
	Format *AudioFormat `json:"format,omitempty"`
	// BufferLength the buffer length in microseconds
	BufferLength *uint32 `json:"buffer-length,omitempty"`
}

// AudiodevGenericOptions
//
// Generic driver-specific options.
type AudiodevGenericOptions struct {
	// In options of the capture stream
	In *AudiodevPerDirectionOptions `json:"in,omitempty"`
	// Out options of the playback stream
	Out *AudiodevPerDirectionOptions `json:"out,omitempty"`
}

// AudiodevAlsaPerDirectionOptions
//
// Options of the ALSA backend that are used for both playback and recording.
type AudiodevAlsaPerDirectionOptions struct {
	AudiodevPerDirectionOptions

	// Dev the name of the ALSA device to use (default 'default')
	Dev *string `json:"dev,omitempty"`
	// PeriodLength the period length in microseconds
	PeriodLength *uint32 `json:"period-length,omitempty"`
	// TryPoll attempt to use poll mode, falling back to non-polling access on failure (default true)
	TryPoll *bool `json:"try-poll,omitempty"`
}

// AudiodevAlsaOptions
//
// Options of the ALSA audio backend.
type AudiodevAlsaOptions struct {
	// In options of the capture stream
	In *AudiodevAlsaPerDirectionOptions `json:"in,omitempty"`
	// Out options of the playback stream
	Out *AudiodevAlsaPerDirectionOptions `json:"out,omitempty"`
	// Threshold set the threshold (in microseconds) when playback starts
	Threshold *uint32 `json:"threshold,omitempty"`
}

// AudiodevSndioOptions
//
// Options of the sndio audio backend.
type AudiodevSndioOptions struct {
	// In options of the capture stream
	In *AudiodevPerDirectionOptions `json:"in,omitempty"`
	// Out options of the playback stream
	Out *AudiodevPerDirectionOptions `json:"out,omitempty"`
	// Dev the name of the sndio device to use (default 'default')
	Dev *string `json:"dev,omitempty"`
	// Latency play buffer size (in microseconds)
	Latency *uint32 `json:"latency,omitempty"`
}

// AudiodevCoreaudioPerDirectionOptions
//
// Options of the Core Audio backend that are used for both playback and recording.
type AudiodevCoreaudioPerDirectionOptions struct {
	AudiodevPerDirectionOptions

	// BufferCount number of buffers
	BufferCount *uint32 `json:"buffer-count,omitempty"`
}

// AudiodevCoreaudioOptions
//
// Options of the coreaudio audio backend.
type AudiodevCoreaudioOptions struct {
	// In options of the capture stream
	In *AudiodevCoreaudioPerDirectionOptions `json:"in,omitempty"`
	// Out options of the playback stream
	Out *AudiodevCoreaudioPerDirectionOptions `json:"out,omitempty"`
}

// AudiodevDsoundOptions
//
// Options of the DirectSound audio backend.
type AudiodevDsoundOptions struct {
	// In options of the capture stream
	In *AudiodevPerDirectionOptions `json:"in,omitempty"`
	// Out options of the playback stream
	Out *AudiodevPerDirectionOptions `json:"out,omitempty"`
	// Latency add extra latency to playback in microseconds (default 10000)
	Latency *uint32 `json:"latency,omitempty"`
}

// AudiodevJackPerDirectionOptions
//
// Options of the JACK backend that are used for both playback and recording.
type AudiodevJackPerDirectionOptions struct {
	AudiodevPerDirectionOptions

	// ServerName select from among several possible concurrent server
	ServerName *string `json:"server-name,omitempty"`
	// ClientName the client name to use. The server will modify this name to create a unique variant, if needed unless @exact-name is
	ClientName *string `json:"client-name,omitempty"`
	// ConnectPorts if set, a regular expression of JACK client port name(s) to monitor for and automatically connect to
	ConnectPorts *string `json:"connect-ports,omitempty"`
	// StartServer start a jack server process if one is not already
	StartServer *bool `json:"start-server,omitempty"`
	// ExactName use the exact name requested otherwise JACK
	ExactName *bool `json:"exact-name,omitempty"`
}

// AudiodevJackOptions
//
// Options of the JACK audio backend.
type AudiodevJackOptions struct {
	// In options of the capture stream
	In *AudiodevJackPerDirectionOptions `json:"in,omitempty"`
	// Out options of the playback stream
	Out *AudiodevJackPerDirectionOptions `json:"out,omitempty"`
}

// AudiodevOssPerDirectionOptions
//
// Options of the OSS backend that are used for both playback and recording.
type AudiodevOssPerDirectionOptions struct {
	AudiodevPerDirectionOptions

	// Dev file name of the OSS device (default '/dev/dsp')
	Dev *string `json:"dev,omitempty"`
	// BufferCount number of buffers
	BufferCount *uint32 `json:"buffer-count,omitempty"`
	// TryPoll attempt to use poll mode, falling back to non-polling access on failure (default true)
	TryPoll *bool `json:"try-poll,omitempty"`
}

// AudiodevOssOptions
//
// Options of the OSS audio backend.
type AudiodevOssOptions struct {
	// In options of the capture stream
	In *AudiodevOssPerDirectionOptions `json:"in,omitempty"`
	// Out options of the playback stream
	Out *AudiodevOssPerDirectionOptions `json:"out,omitempty"`
	// TryMmap try using memory-mapped access, falling back to non-memory-mapped access on failure (default true)
	TryMmap *bool `json:"try-mmap,omitempty"`
	// Exclusive open device in exclusive mode (vmix won't work) (default false)
	Exclusive *bool `json:"exclusive,omitempty"`
	// DspPolicy set the timing policy of the device (between 0 and 10, where smaller number means smaller latency but higher CPU usage) or -1 to use fragment mode (option ignored on some platforms) (default 5)
	DspPolicy *uint32 `json:"dsp-policy,omitempty"`
}

// AudiodevPaPerDirectionOptions
//
// Options of the Pulseaudio backend that are used for both playback and recording.
type AudiodevPaPerDirectionOptions struct {
	AudiodevPerDirectionOptions

	// Name name of the sink/source to use
	Name *string `json:"name,omitempty"`
	// StreamName name of the PulseAudio stream created by qemu. Can be used to identify the stream in PulseAudio when you create multiple PulseAudio devices or run multiple qemu instances
	StreamName *string `json:"stream-name,omitempty"`
	// Latency latency you want PulseAudio to achieve in microseconds (default 15000)
	Latency *uint32 `json:"latency,omitempty"`
}

// AudiodevPaOptions
//
// Options of the PulseAudio audio backend.
type AudiodevPaOptions struct {
	// In options of the capture stream
	In *AudiodevPaPerDirectionOptions `json:"in,omitempty"`
	// Out options of the playback stream
	Out *AudiodevPaPerDirectionOptions `json:"out,omitempty"`
	// Server PulseAudio server address (default: let PulseAudio choose)
	Server *string `json:"server,omitempty"`
}

// AudiodevPipewirePerDirectionOptions
//
// Options of the PipeWire backend that are used for both playback and recording.
type AudiodevPipewirePerDirectionOptions struct {
	AudiodevPerDirectionOptions

	// Name name of the sink/source to use
	Name *string `json:"name,omitempty"`
	// StreamName name of the PipeWire stream created by qemu. Can be used to identify the stream in PipeWire when you create multiple
	StreamName *string `json:"stream-name,omitempty"`
	// Latency latency you want PipeWire to achieve in microseconds (default 46000)
	Latency *uint32 `json:"latency,omitempty"`
}

// AudiodevPipewireOptions
//
// Options of the PipeWire audio backend.
type AudiodevPipewireOptions struct {
	// In options of the capture stream
	In *AudiodevPipewirePerDirectionOptions `json:"in,omitempty"`
	// Out options of the playback stream
	Out *AudiodevPipewirePerDirectionOptions `json:"out,omitempty"`
}

// AudiodevSdlPerDirectionOptions
//
// Options of the SDL audio backend that are used for both playback and recording.
type AudiodevSdlPerDirectionOptions struct {
	AudiodevPerDirectionOptions

	// BufferCount number of buffers (default 4)
	BufferCount *uint32 `json:"buffer-count,omitempty"`
}

// AudiodevSdlOptions
//
// Options of the SDL audio backend.
type AudiodevSdlOptions struct {
	// In options of the recording stream
	In *AudiodevSdlPerDirectionOptions `json:"in,omitempty"`
	// Out options of the playback stream
	Out *AudiodevSdlPerDirectionOptions `json:"out,omitempty"`
}

// AudiodevWavOptions
//
// Options of the wav audio backend.
type AudiodevWavOptions struct {
	// In options of the capture stream
	In *AudiodevPerDirectionOptions `json:"in,omitempty"`
	// Out options of the playback stream
	Out *AudiodevPerDirectionOptions `json:"out,omitempty"`
	// Path name of the wav file to record (default 'qemu.wav')
	Path *string `json:"path,omitempty"`
}

// AudioFormat An enumeration of possible audio formats.
type AudioFormat string

const (
	// AudioFormatU8 unsigned 8 bit integer
	AudioFormatU8 AudioFormat = "u8"
	// AudioFormatS8 signed 8 bit integer
	AudioFormatS8 AudioFormat = "s8"
	// AudioFormatU16 unsigned 16 bit integer
	AudioFormatU16 AudioFormat = "u16"
	// AudioFormatS16 signed 16 bit integer
	AudioFormatS16 AudioFormat = "s16"
	// AudioFormatU32 unsigned 32 bit integer
	AudioFormatU32 AudioFormat = "u32"
	// AudioFormatS32 signed 32 bit integer
	AudioFormatS32 AudioFormat = "s32"
	// AudioFormatF32 single precision floating-point (since 5.0)
	AudioFormatF32 AudioFormat = "f32"
)

// AudiodevDriver An enumeration of possible audio backend drivers.
type AudiodevDriver string

const (
	AudiodevDriverNone      AudiodevDriver = "none"
	AudiodevDriverAlsa      AudiodevDriver = "alsa"
	AudiodevDriverCoreaudio AudiodevDriver = "coreaudio"
	AudiodevDriverDbus      AudiodevDriver = "dbus"
	AudiodevDriverDsound    AudiodevDriver = "dsound"
	// AudiodevDriverJack JACK audio backend (since 5.1)
	AudiodevDriverJack     AudiodevDriver = "jack"
	AudiodevDriverOss      AudiodevDriver = "oss"
	AudiodevDriverPa       AudiodevDriver = "pa"
	AudiodevDriverPipewire AudiodevDriver = "pipewire"
	AudiodevDriverSdl      AudiodevDriver = "sdl"
	AudiodevDriverSndio    AudiodevDriver = "sndio"
	AudiodevDriverSpice    AudiodevDriver = "spice"
	AudiodevDriverWav      AudiodevDriver = "wav"
)

// Audiodev
//
// Options of an audio backend.
type Audiodev struct {
	// Discriminator: driver

	// Id identifier of the backend
	Id string `json:"id"`
	// Driver the backend driver to use
	Driver AudiodevDriver `json:"driver"`
	// TimerPeriod timer period (in microseconds, 0: use lowest possible)
	TimerPeriod *uint32 `json:"timer-period,omitempty"`

	None      *AudiodevGenericOptions   `json:"-"`
	Alsa      *AudiodevAlsaOptions      `json:"-"`
	Coreaudio *AudiodevCoreaudioOptions `json:"-"`
	Dbus      *AudiodevGenericOptions   `json:"-"`
	Dsound    *AudiodevDsoundOptions    `json:"-"`
	Jack      *AudiodevJackOptions      `json:"-"`
	Oss       *AudiodevOssOptions       `json:"-"`
	Pa        *AudiodevPaOptions        `json:"-"`
	Pipewire  *AudiodevPipewireOptions  `json:"-"`
	Sdl       *AudiodevSdlOptions       `json:"-"`
	Sndio     *AudiodevSndioOptions     `json:"-"`
	Spice     *AudiodevGenericOptions   `json:"-"`
	Wav       *AudiodevWavOptions       `json:"-"`
}

func (u Audiodev) MarshalJSON() ([]byte, error) {
	switch u.Driver {
	case "none":
		if u.None == nil {
			return nil, fmt.Errorf("expected None to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevGenericOptions
		}{
			Id:                     u.Id,
			Driver:                 u.Driver,
			TimerPeriod:            u.TimerPeriod,
			AudiodevGenericOptions: u.None,
		})
	case "alsa":
		if u.Alsa == nil {
			return nil, fmt.Errorf("expected Alsa to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevAlsaOptions
		}{
			Id:                  u.Id,
			Driver:              u.Driver,
			TimerPeriod:         u.TimerPeriod,
			AudiodevAlsaOptions: u.Alsa,
		})
	case "coreaudio":
		if u.Coreaudio == nil {
			return nil, fmt.Errorf("expected Coreaudio to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevCoreaudioOptions
		}{
			Id:                       u.Id,
			Driver:                   u.Driver,
			TimerPeriod:              u.TimerPeriod,
			AudiodevCoreaudioOptions: u.Coreaudio,
		})
	case "dbus":
		if u.Dbus == nil {
			return nil, fmt.Errorf("expected Dbus to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevGenericOptions
		}{
			Id:                     u.Id,
			Driver:                 u.Driver,
			TimerPeriod:            u.TimerPeriod,
			AudiodevGenericOptions: u.Dbus,
		})
	case "dsound":
		if u.Dsound == nil {
			return nil, fmt.Errorf("expected Dsound to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevDsoundOptions
		}{
			Id:                    u.Id,
			Driver:                u.Driver,
			TimerPeriod:           u.TimerPeriod,
			AudiodevDsoundOptions: u.Dsound,
		})
	case "jack":
		if u.Jack == nil {
			return nil, fmt.Errorf("expected Jack to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevJackOptions
		}{
			Id:                  u.Id,
			Driver:              u.Driver,
			TimerPeriod:         u.TimerPeriod,
			AudiodevJackOptions: u.Jack,
		})
	case "oss":
		if u.Oss == nil {
			return nil, fmt.Errorf("expected Oss to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevOssOptions
		}{
			Id:                 u.Id,
			Driver:             u.Driver,
			TimerPeriod:        u.TimerPeriod,
			AudiodevOssOptions: u.Oss,
		})
	case "pa":
		if u.Pa == nil {
			return nil, fmt.Errorf("expected Pa to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevPaOptions
		}{
			Id:                u.Id,
			Driver:            u.Driver,
			TimerPeriod:       u.TimerPeriod,
			AudiodevPaOptions: u.Pa,
		})
	case "pipewire":
		if u.Pipewire == nil {
			return nil, fmt.Errorf("expected Pipewire to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevPipewireOptions
		}{
			Id:                      u.Id,
			Driver:                  u.Driver,
			TimerPeriod:             u.TimerPeriod,
			AudiodevPipewireOptions: u.Pipewire,
		})
	case "sdl":
		if u.Sdl == nil {
			return nil, fmt.Errorf("expected Sdl to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevSdlOptions
		}{
			Id:                 u.Id,
			Driver:             u.Driver,
			TimerPeriod:        u.TimerPeriod,
			AudiodevSdlOptions: u.Sdl,
		})
	case "sndio":
		if u.Sndio == nil {
			return nil, fmt.Errorf("expected Sndio to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevSndioOptions
		}{
			Id:                   u.Id,
			Driver:               u.Driver,
			TimerPeriod:          u.TimerPeriod,
			AudiodevSndioOptions: u.Sndio,
		})
	case "spice":
		if u.Spice == nil {
			return nil, fmt.Errorf("expected Spice to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevGenericOptions
		}{
			Id:                     u.Id,
			Driver:                 u.Driver,
			TimerPeriod:            u.TimerPeriod,
			AudiodevGenericOptions: u.Spice,
		})
	case "wav":
		if u.Wav == nil {
			return nil, fmt.Errorf("expected Wav to be set")
		}

		return json.Marshal(struct {
			Id          string         `json:"id"`
			Driver      AudiodevDriver `json:"driver"`
			TimerPeriod *uint32        `json:"timer-period,omitempty"`
			*AudiodevWavOptions
		}{
			Id:                 u.Id,
			Driver:             u.Driver,
			TimerPeriod:        u.TimerPeriod,
			AudiodevWavOptions: u.Wav,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// QueryAudiodevs
//
// Returns information about audiodev configuration
type QueryAudiodevs struct {
}

func (QueryAudiodevs) Command() string {
	return "query-audiodevs"
}

func (cmd QueryAudiodevs) Execute(ctx context.Context, client api.Client) ([]Audiodev, error) {
	var ret []Audiodev

	return ret, client.Execute(ctx, "query-audiodevs", cmd, &ret)
}

// AcpiTableOptions
//
// Specify an ACPI table on the command line to load. At most one of @file and @data can be specified. The list of files specified by any one of them is loaded and concatenated in order. If both are omitted, @data is implied. Other fields / optargs can be used to override fields of the generic ACPI table header; refer to the ACPI specification 5.0, section 5.2.6 System Description Table Header. If a header field is not overridden, then the corresponding value from the concatenated blob is used (in case of @file), or it is filled in with a hard-coded value (in case of @data). String fields are copied into the matching ACPI member from lowest address upwards, and silently truncated / NUL-padded to length.
type AcpiTableOptions struct {
	// Sig table signature / identifier (4 bytes)
	Sig *string `json:"sig,omitempty"`
	// Rev table revision number (dependent on signature, 1 byte)
	Rev *uint8 `json:"rev,omitempty"`
	// OemId OEM identifier (6 bytes)
	OemId *string `json:"oem_id,omitempty"`
	// OemTableId OEM table identifier (8 bytes)
	OemTableId *string `json:"oem_table_id,omitempty"`
	// OemRev OEM-supplied revision number (4 bytes)
	OemRev *uint32 `json:"oem_rev,omitempty"`
	// AslCompilerId identifier of the utility that created the table (4 bytes)
	AslCompilerId *string `json:"asl_compiler_id,omitempty"`
	// AslCompilerRev revision number of the utility that created the table (4 bytes)
	AslCompilerRev *uint32 `json:"asl_compiler_rev,omitempty"`
	// File colon (:) separated list of pathnames to load and concatenate as table data. The resultant binary blob is expected to have an ACPI table header. At least one file is required. This field excludes @data.
	File *string `json:"file,omitempty"`
	// Data colon (:) separated list of pathnames to load and concatenate as table data. The resultant binary blob must not have an ACPI table header. At least one file is required. This field excludes @file.
	Data *string `json:"data,omitempty"`
}

// ACPISlotType
type ACPISlotType string

const (
	// ACPISlotTypeDimm memory slot
	ACPISlotTypeDimm ACPISlotType = "DIMM"
	// ACPISlotTypeCpu logical CPU slot (since 2.7)
	ACPISlotTypeCpu ACPISlotType = "CPU"
)

// ACPIOSTInfo
//
// OSPM Status Indication for a device For description of possible values of @source and @status fields see "_OST (OSPM Status Indication)" chapter of ACPI5.0 spec.
type ACPIOSTInfo struct {
	// Device device ID associated with slot
	Device *string `json:"device,omitempty"`
	// Slot slot ID, unique per slot of a given @slot-type
	Slot string `json:"slot"`
	// SlotType type of the slot
	SlotType ACPISlotType `json:"slot-type"`
	// Source an integer containing the source event
	Source int64 `json:"source"`
	// Status an integer containing the status code
	Status int64 `json:"status"`
}

// QueryAcpiOspmStatus
//
// Return a list of ACPIOSTInfo for devices that support status reporting via ACPI _OST method.
type QueryAcpiOspmStatus struct {
}

func (QueryAcpiOspmStatus) Command() string {
	return "query-acpi-ospm-status"
}

func (cmd QueryAcpiOspmStatus) Execute(ctx context.Context, client api.Client) ([]ACPIOSTInfo, error) {
	var ret []ACPIOSTInfo

	return ret, client.Execute(ctx, "query-acpi-ospm-status", cmd, &ret)
}

// AcpiDeviceOstEvent (ACPI_DEVICE_OST)
//
// Emitted when guest executes ACPI _OST method.
type AcpiDeviceOstEvent struct {
	// Info OSPM Status Indication
	Info ACPIOSTInfo `json:"info"`
}

func (AcpiDeviceOstEvent) Event() string {
	return "ACPI_DEVICE_OST"
}

// PciMemoryRange
//
// A PCI device memory region
type PciMemoryRange struct {
	// Base the starting address (guest physical)
	Base int64 `json:"base"`
	// Limit the ending address (guest physical)
	Limit int64 `json:"limit"`
}

// PciMemoryRegion
//
// Information about a PCI device I/O region.
type PciMemoryRegion struct {
	// Bar the index of the Base Address Register for this region
	Bar int64 `json:"bar"`
	// Type - 'io' if the region is a PIO region - 'memory' if the region is a MMIO region
	Type    string `json:"type"`
	Address int64  `json:"address"`
	// Size memory size
	Size int64 `json:"size"`
	// Prefetch if @type is 'memory', true if the memory is prefetchable
	Prefetch *bool `json:"prefetch,omitempty"`
	// MemType64 if @type is 'memory', true if the BAR is 64-bit
	MemType64 *bool `json:"mem_type_64,omitempty"`
}

// PciBusInfo
//
// Information about a bus of a PCI Bridge device
type PciBusInfo struct {
	// Number primary bus interface number. This should be the number of the bus the device resides on.
	Number int64 `json:"number"`
	// Secondary secondary bus interface number. This is the number of the main bus for the bridge
	Secondary int64 `json:"secondary"`
	// Subordinate This is the highest number bus that resides below the bridge.
	Subordinate int64 `json:"subordinate"`
	// IoRange The PIO range for all devices on this bridge
	IoRange PciMemoryRange `json:"io_range"`
	// MemoryRange The MMIO range for all devices on this bridge
	MemoryRange PciMemoryRange `json:"memory_range"`
	// PrefetchableRange The range of prefetchable MMIO for all devices on this bridge
	PrefetchableRange PciMemoryRange `json:"prefetchable_range"`
}

// PciBridgeInfo
//
// Information about a PCI Bridge device
type PciBridgeInfo struct {
	// Bus information about the bus the device resides on
	Bus PciBusInfo `json:"bus"`
	// Devices a list of @PciDeviceInfo for each device on this bridge
	Devices []PciDeviceInfo `json:"devices,omitempty"`
}

// PciDeviceClass
//
// Information about the Class of a PCI device
type PciDeviceClass struct {
	// Desc a string description of the device's class
	Desc *string `json:"desc,omitempty"`
	// Class the class code of the device
	Class int64 `json:"class"`
}

// PciDeviceId
//
// Information about the Id of a PCI device
type PciDeviceId struct {
	// Device the PCI device id
	Device int64 `json:"device"`
	// Vendor the PCI vendor id
	Vendor int64 `json:"vendor"`
	// Subsystem the PCI subsystem id (since 3.1)
	Subsystem *int64 `json:"subsystem,omitempty"`
	// SubsystemVendor the PCI subsystem vendor id (since 3.1)
	SubsystemVendor *int64 `json:"subsystem-vendor,omitempty"`
}

// PciDeviceInfo
//
// Information about a PCI device
type PciDeviceInfo struct {
	// Bus the bus number of the device
	Bus int64 `json:"bus"`
	// Slot the slot the device is located in
	Slot int64 `json:"slot"`
	// Function the function of the slot used by the device
	Function int64 `json:"function"`
	// ClassInfo the class of the device
	ClassInfo PciDeviceClass `json:"class_info"`
	// Id the PCI device id
	Id PciDeviceId `json:"id"`
	// Irq if an IRQ is assigned to the device, the IRQ number
	Irq *int64 `json:"irq,omitempty"`
	// IrqPin the IRQ pin, zero means no IRQ (since 5.1)
	IrqPin int64 `json:"irq_pin"`
	// QdevId the device name of the PCI device
	QdevId string `json:"qdev_id"`
	// PciBridge if the device is a PCI bridge, the bridge information
	PciBridge *PciBridgeInfo `json:"pci_bridge,omitempty"`
	// Regions a list of the PCI I/O regions associated with the device
	Regions []PciMemoryRegion `json:"regions"`
}

// PciInfo
//
// Information about a PCI bus
type PciInfo struct {
	// Bus the bus index
	Bus int64 `json:"bus"`
	// Devices a list of devices on this bus
	Devices []PciDeviceInfo `json:"devices"`
}

// QueryPci
//
// Return information about the PCI bus topology of the guest.
type QueryPci struct {
}

func (QueryPci) Command() string {
	return "query-pci"
}

func (cmd QueryPci) Execute(ctx context.Context, client api.Client) ([]PciInfo, error) {
	var ret []PciInfo

	return ret, client.Execute(ctx, "query-pci", cmd, &ret)
}

// StatsType Enumeration of statistics types
type StatsType string

const (
	// StatsTypeCumulative stat is cumulative; value can only increase.
	StatsTypeCumulative StatsType = "cumulative"
	// StatsTypeInstant stat is instantaneous; value can increase or decrease.
	StatsTypeInstant StatsType = "instant"
	// StatsTypePeak stat is the peak value; value can only increase.
	StatsTypePeak StatsType = "peak"
	// StatsTypeLinearHistogram stat is a linear histogram.
	StatsTypeLinearHistogram StatsType = "linear-histogram"
	// StatsTypeLog2Histogram stat is a logarithmic histogram, with one bucket for each power of two.
	StatsTypeLog2Histogram StatsType = "log2-histogram"
)

// StatsUnit Enumeration of unit of measurement for statistics
type StatsUnit string

const (
	// StatsUnitBytes stat reported in bytes.
	StatsUnitBytes StatsUnit = "bytes"
	// StatsUnitSeconds stat reported in seconds.
	StatsUnitSeconds StatsUnit = "seconds"
	// StatsUnitCycles stat reported in clock cycles.
	StatsUnitCycles StatsUnit = "cycles"
	// StatsUnitBoolean stat is a boolean value.
	StatsUnitBoolean StatsUnit = "boolean"
)

// StatsProvider Enumeration of statistics providers.
type StatsProvider string

const (
	// StatsProviderKvm since 7.1
	StatsProviderKvm StatsProvider = "kvm"
	// StatsProviderCryptodev since 8.0
	StatsProviderCryptodev StatsProvider = "cryptodev"
)

// StatsTarget The kinds of objects on which one can request statistics.
type StatsTarget string

const (
	// StatsTargetVm statistics that apply to the entire virtual machine or the entire QEMU process.
	StatsTargetVm StatsTarget = "vm"
	// StatsTargetVcpu statistics that apply to a single virtual CPU.
	StatsTargetVcpu StatsTarget = "vcpu"
	// StatsTargetCryptodev statistics that apply to a crypto device (since 8.0)
	StatsTargetCryptodev StatsTarget = "cryptodev"
)

// StatsRequest
//
// Indicates a set of statistics that should be returned by query-stats.
type StatsRequest struct {
	// Provider provider for which to return statistics.
	Provider StatsProvider `json:"provider"`
	// Names statistics to be returned (all if omitted).
	Names []string `json:"names,omitempty"`
}

// StatsVCPUFilter
type StatsVCPUFilter struct {
	// Vcpus list of QOM paths for the desired vCPU objects.
	Vcpus []string `json:"vcpus,omitempty"`
}

// StatsFilter
//
// The arguments to the query-stats command; specifies a target for which to request statistics and optionally the required subset of
type StatsFilter struct {
	// Discriminator: target

	// Target the kind of objects to query
	Target    StatsTarget    `json:"target"`
	Providers []StatsRequest `json:"providers,omitempty"`

	Vcpu *StatsVCPUFilter `json:"-"`
}

func (u StatsFilter) MarshalJSON() ([]byte, error) {
	switch u.Target {
	case "vcpu":
		if u.Vcpu == nil {
			return nil, fmt.Errorf("expected Vcpu to be set")
		}

		return json.Marshal(struct {
			Target    StatsTarget    `json:"target"`
			Providers []StatsRequest `json:"providers,omitempty"`
			*StatsVCPUFilter
		}{
			Target:          u.Target,
			Providers:       u.Providers,
			StatsVCPUFilter: u.Vcpu,
		})
	}

	return nil, fmt.Errorf("unknown type")
}

// StatsValue
type StatsValue struct {
	// Scalar single unsigned 64-bit integers.
	Scalar  *uint64 `json:"-"`
	Boolean *bool   `json:"-"`
	// List list of unsigned 64-bit integers (used for histograms).
	List []uint64 `json:"-"`
}

func (a StatsValue) MarshalJSON() ([]byte, error) {
	switch {
	case a.Scalar != nil:
		return json.Marshal(a.Scalar)
	case a.Boolean != nil:
		return json.Marshal(a.Boolean)
	case a.List != nil:
		return json.Marshal(a.List)
	}

	return nil, fmt.Errorf("unknown type")
}

// Stats
type Stats struct {
	// Name name of stat.
	Name string `json:"name"`
	// Value stat value.
	Value StatsValue `json:"value"`
}

// StatsResult
type StatsResult struct {
	// Provider provider for this set of statistics.
	Provider StatsProvider `json:"provider"`
	// QomPath Path to the object for which the statistics are returned, if the object is exposed in the QOM tree
	QomPath *string `json:"qom-path,omitempty"`
	// Stats list of statistics.
	Stats []Stats `json:"stats"`
}

// QueryStats
//
// Return runtime-collected statistics for objects such as the VM or its vCPUs. The arguments are a StatsFilter and specify the provider and objects to return statistics about.
type QueryStats struct {
	StatsFilter
}

func (QueryStats) Command() string {
	return "query-stats"
}

func (cmd QueryStats) Execute(ctx context.Context, client api.Client) ([]StatsResult, error) {
	var ret []StatsResult

	return ret, client.Execute(ctx, "query-stats", cmd, &ret)
}

// StatsSchemaValue
//
// Schema for a single statistic.
type StatsSchemaValue struct {
	// Name name of the statistic; each element of the schema is uniquely identified by a target, a provider (both available in @StatsSchema) and the name.
	Name string `json:"name"`
	// Type kind of statistic.
	Type StatsType `json:"type"`
	// Unit basic unit of measure for the statistic; if missing, the statistic is a simple number or counter.
	Unit *StatsUnit `json:"unit,omitempty"`
	// Base base for the multiple of @unit in which the statistic is measured. Only present if @exponent is non-zero; @base and @exponent together form a SI prefix (e.g., _nano-_ for ``base=10`` and ``exponent=-9``) or IEC binary prefix (e.g. _kibi-_ for ``base=2`` and ``exponent=10``)
	Base *int8 `json:"base,omitempty"`
	// Exponent exponent for the multiple of @unit in which the statistic is expressed, or 0 for the basic unit
	Exponent int16 `json:"exponent"`
	// BucketSize Present when @type is "linear-histogram", contains the width of each bucket of the histogram.
	BucketSize *uint32 `json:"bucket-size,omitempty"`
}

// StatsSchema
//
// Schema for all available statistics for a provider and target.
type StatsSchema struct {
	// Provider provider for this set of statistics.
	Provider StatsProvider `json:"provider"`
	// Target the kind of object that can be queried through the provider.
	Target StatsTarget `json:"target"`
	// Stats list of statistics.
	Stats []StatsSchemaValue `json:"stats"`
}

// QueryStatsSchemas
//
// Return the schema for all available runtime-collected statistics.
type QueryStatsSchemas struct {
	Provider *StatsProvider `json:"provider,omitempty"`
}

func (QueryStatsSchemas) Command() string {
	return "query-stats-schemas"
}

func (cmd QueryStatsSchemas) Execute(ctx context.Context, client api.Client) ([]StatsSchema, error) {
	var ret []StatsSchema

	return ret, client.Execute(ctx, "query-stats-schemas", cmd, &ret)
}

// VirtioInfo
//
// Basic information about a given VirtIODevice
type VirtioInfo struct {
	// Path The VirtIODevice's canonical QOM path
	Path string `json:"path"`
	// Name Name of the VirtIODevice
	Name string `json:"name"`
}

// QueryVirtio
//
// Returns a list of all realized VirtIODevices
type QueryVirtio struct {
}

func (QueryVirtio) Command() string {
	return "x-query-virtio"
}

func (cmd QueryVirtio) Execute(ctx context.Context, client api.Client) ([]VirtioInfo, error) {
	var ret []VirtioInfo

	return ret, client.Execute(ctx, "x-query-virtio", cmd, &ret)
}

// VhostStatus
//
// Information about a vhost device. This information will only be displayed if the vhost device is active.
type VhostStatus struct {
	// NMemSections vhost_dev n_mem_sections
	NMemSections int64 `json:"n-mem-sections"`
	// NTmpSections vhost_dev n_tmp_sections
	NTmpSections int64 `json:"n-tmp-sections"`
	// Nvqs vhost_dev nvqs (number of virtqueues being used)
	Nvqs uint32 `json:"nvqs"`
	// VqIndex vhost_dev vq_index
	VqIndex int64 `json:"vq-index"`
	// Features vhost_dev features
	Features VirtioDeviceFeatures `json:"features"`
	// AckedFeatures vhost_dev acked_features
	AckedFeatures VirtioDeviceFeatures `json:"acked-features"`
	// BackendFeatures vhost_dev backend_features
	BackendFeatures VirtioDeviceFeatures `json:"backend-features"`
	// ProtocolFeatures vhost_dev protocol_features
	ProtocolFeatures VhostDeviceProtocols `json:"protocol-features"`
	// MaxQueues vhost_dev max_queues
	MaxQueues uint64 `json:"max-queues"`
	// BackendCap vhost_dev backend_cap
	BackendCap uint64 `json:"backend-cap"`
	// LogEnabled vhost_dev log_enabled flag
	LogEnabled bool `json:"log-enabled"`
	// LogSize vhost_dev log_size
	LogSize uint64 `json:"log-size"`
}

// VirtioStatus
//
// Full status of the virtio device with most VirtIODevice members. Also includes the full status of the corresponding vhost device if the vhost device is active.
type VirtioStatus struct {
	// Name VirtIODevice name
	Name string `json:"name"`
	// DeviceId VirtIODevice ID
	DeviceId uint16 `json:"device-id"`
	// VhostStarted VirtIODevice vhost_started flag
	VhostStarted bool `json:"vhost-started"`
	// DeviceEndian VirtIODevice device_endian
	DeviceEndian string `json:"device-endian"`
	// GuestFeatures VirtIODevice guest_features
	GuestFeatures VirtioDeviceFeatures `json:"guest-features"`
	// HostFeatures VirtIODevice host_features
	HostFeatures VirtioDeviceFeatures `json:"host-features"`
	// BackendFeatures VirtIODevice backend_features
	BackendFeatures VirtioDeviceFeatures `json:"backend-features"`
	// NumVqs VirtIODevice virtqueue count. This is the number of active virtqueues being used by the VirtIODevice.
	NumVqs int64 `json:"num-vqs"`
	// Status VirtIODevice configuration status (VirtioDeviceStatus)
	Status VirtioDeviceStatus `json:"status"`
	// Isr VirtIODevice ISR
	Isr uint8 `json:"isr"`
	// QueueSel VirtIODevice queue_sel
	QueueSel uint16 `json:"queue-sel"`
	// VmRunning VirtIODevice vm_running flag
	VmRunning bool `json:"vm-running"`
	// Broken VirtIODevice broken flag
	Broken bool `json:"broken"`
	// Disabled VirtIODevice disabled flag
	Disabled bool `json:"disabled"`
	// UseStarted VirtIODevice use_started flag
	UseStarted bool `json:"use-started"`
	// Started VirtIODevice started flag
	Started bool `json:"started"`
	// StartOnKick VirtIODevice start_on_kick flag
	StartOnKick bool `json:"start-on-kick"`
	// DisableLegacyCheck VirtIODevice disabled_legacy_check flag
	DisableLegacyCheck bool `json:"disable-legacy-check"`
	// BusName VirtIODevice bus_name
	BusName string `json:"bus-name"`
	// UseGuestNotifierMask VirtIODevice use_guest_notifier_mask flag
	UseGuestNotifierMask bool `json:"use-guest-notifier-mask"`
	// VhostDev Corresponding vhost device info for a given VirtIODevice. Present if the given VirtIODevice has an active vhost device.
	VhostDev *VhostStatus `json:"vhost-dev,omitempty"`
}

// QueryVirtioStatus
//
// Poll for a comprehensive status of a given virtio device
type QueryVirtioStatus struct {
	// Path Canonical QOM path of the VirtIODevice
	Path string `json:"path"`
}

func (QueryVirtioStatus) Command() string {
	return "x-query-virtio-status"
}

func (cmd QueryVirtioStatus) Execute(ctx context.Context, client api.Client) (VirtioStatus, error) {
	var ret VirtioStatus

	return ret, client.Execute(ctx, "x-query-virtio-status", cmd, &ret)
}

// VirtioDeviceStatus
//
// A structure defined to list the configuration statuses of a virtio device
type VirtioDeviceStatus struct {
	// Statuses List of decoded configuration statuses of the virtio device
	Statuses []string `json:"statuses"`
	// UnknownStatuses Virtio device statuses bitmap that have not been decoded
	UnknownStatuses *uint8 `json:"unknown-statuses,omitempty"`
}

// VhostDeviceProtocols
//
// A structure defined to list the vhost user protocol features of a Vhost User device
type VhostDeviceProtocols struct {
	// Protocols List of decoded vhost user protocol features of a vhost user device
	Protocols []string `json:"protocols"`
	// UnknownProtocols Vhost user device protocol features bitmap that have not been decoded
	UnknownProtocols *uint64 `json:"unknown-protocols,omitempty"`
}

// VirtioDeviceFeatures
//
// The common fields that apply to most Virtio devices. Some devices may not have their own device-specific features (e.g. virtio-rng).
type VirtioDeviceFeatures struct {
	// Transports List of transport features of the virtio device
	Transports []string `json:"transports"`
	// DevFeatures List of device-specific features (if the device has unique features)
	DevFeatures []string `json:"dev-features,omitempty"`
	// UnknownDevFeatures Virtio device features bitmap that have not been decoded
	UnknownDevFeatures *uint64 `json:"unknown-dev-features,omitempty"`
}

// VirtQueueStatus
//
// Information of a VirtIODevice VirtQueue, including most members of the VirtQueue data structure.
type VirtQueueStatus struct {
	// Name Name of the VirtIODevice that uses this VirtQueue
	Name string `json:"name"`
	// QueueIndex VirtQueue queue_index
	QueueIndex uint16 `json:"queue-index"`
	// Inuse VirtQueue inuse
	Inuse uint32 `json:"inuse"`
	// VringNum VirtQueue vring.num
	VringNum uint32 `json:"vring-num"`
	// VringNumDefault VirtQueue vring.num_default
	VringNumDefault uint32 `json:"vring-num-default"`
	// VringAlign VirtQueue vring.align
	VringAlign uint32 `json:"vring-align"`
	// VringDesc VirtQueue vring.desc (descriptor area)
	VringDesc uint64 `json:"vring-desc"`
	// VringAvail VirtQueue vring.avail (driver area)
	VringAvail uint64 `json:"vring-avail"`
	// VringUsed VirtQueue vring.used (device area)
	VringUsed uint64 `json:"vring-used"`
	// LastAvailIdx VirtQueue last_avail_idx or return of vhost_dev vhost_get_vring_base (if vhost active)
	LastAvailIdx *uint16 `json:"last-avail-idx,omitempty"`
	// ShadowAvailIdx VirtQueue shadow_avail_idx
	ShadowAvailIdx *uint16 `json:"shadow-avail-idx,omitempty"`
	// UsedIdx VirtQueue used_idx
	UsedIdx uint16 `json:"used-idx"`
	// SignalledUsed VirtQueue signalled_used
	SignalledUsed uint16 `json:"signalled-used"`
	// SignalledUsedValid VirtQueue signalled_used_valid flag
	SignalledUsedValid bool `json:"signalled-used-valid"`
}

// QueryVirtioQueueStatus
//
// Return the status of a given VirtIODevice's VirtQueue
type QueryVirtioQueueStatus struct {
	// Path VirtIODevice canonical QOM path
	Path string `json:"path"`
	// Queue VirtQueue index to examine
	Queue uint16 `json:"queue"`
}

func (QueryVirtioQueueStatus) Command() string {
	return "x-query-virtio-queue-status"
}

func (cmd QueryVirtioQueueStatus) Execute(ctx context.Context, client api.Client) (VirtQueueStatus, error) {
	var ret VirtQueueStatus

	return ret, client.Execute(ctx, "x-query-virtio-queue-status", cmd, &ret)
}

// VirtVhostQueueStatus
//
// Information of a vhost device's vhost_virtqueue, including most members of the vhost_dev vhost_virtqueue data structure.
type VirtVhostQueueStatus struct {
	// Name Name of the VirtIODevice that uses this vhost_virtqueue
	Name string `json:"name"`
	// Kick vhost_virtqueue kick
	Kick int64 `json:"kick"`
	// Call vhost_virtqueue call
	Call int64 `json:"call"`
	// Desc vhost_virtqueue desc
	Desc uint64 `json:"desc"`
	// Avail vhost_virtqueue avail
	Avail uint64 `json:"avail"`
	// Used vhost_virtqueue used
	Used uint64 `json:"used"`
	// Num vhost_virtqueue num
	Num int64 `json:"num"`
	// DescPhys vhost_virtqueue desc_phys (descriptor area phys. addr.)
	DescPhys uint64 `json:"desc-phys"`
	// DescSize vhost_virtqueue desc_size
	DescSize uint32 `json:"desc-size"`
	// AvailPhys vhost_virtqueue avail_phys (driver area phys. addr.)
	AvailPhys uint64 `json:"avail-phys"`
	// AvailSize vhost_virtqueue avail_size
	AvailSize uint32 `json:"avail-size"`
	// UsedPhys vhost_virtqueue used_phys (device area phys. addr.)
	UsedPhys uint64 `json:"used-phys"`
	// UsedSize vhost_virtqueue used_size
	UsedSize uint32 `json:"used-size"`
}

// QueryVirtioVhostQueueStatus
//
// Return information of a given vhost device's vhost_virtqueue
type QueryVirtioVhostQueueStatus struct {
	// Path VirtIODevice canonical QOM path
	Path string `json:"path"`
	// Queue vhost_virtqueue index to examine
	Queue uint16 `json:"queue"`
}

func (QueryVirtioVhostQueueStatus) Command() string {
	return "x-query-virtio-vhost-queue-status"
}

func (cmd QueryVirtioVhostQueueStatus) Execute(ctx context.Context, client api.Client) (VirtVhostQueueStatus, error) {
	var ret VirtVhostQueueStatus

	return ret, client.Execute(ctx, "x-query-virtio-vhost-queue-status", cmd, &ret)
}

// VirtioRingDesc
//
// Information regarding the vring descriptor area
type VirtioRingDesc struct {
	// Addr Guest physical address of the descriptor area
	Addr uint64 `json:"addr"`
	// Len Length of the descriptor area
	Len uint32 `json:"len"`
	// Flags List of descriptor flags
	Flags []string `json:"flags"`
}

// VirtioRingAvail
//
// Information regarding the avail vring (a.k.a. driver area)
type VirtioRingAvail struct {
	// Flags VRingAvail flags
	Flags uint16 `json:"flags"`
	// Idx VRingAvail index
	Idx uint16 `json:"idx"`
	// Ring VRingAvail ring[] entry at provided index
	Ring uint16 `json:"ring"`
}

// VirtioRingUsed
//
// Information regarding the used vring (a.k.a. device area)
type VirtioRingUsed struct {
	// Flags VRingUsed flags
	Flags uint16 `json:"flags"`
	// Idx VRingUsed index
	Idx uint16 `json:"idx"`
}

// VirtioQueueElement
//
// Information regarding a VirtQueue's VirtQueueElement including descriptor, driver, and device areas
type VirtioQueueElement struct {
	// Name Name of the VirtIODevice that uses this VirtQueue
	Name string `json:"name"`
	// Index Index of the element in the queue
	Index uint32 `json:"index"`
	// Descs List of descriptors (VirtioRingDesc)
	Descs []VirtioRingDesc `json:"descs"`
	// Avail VRingAvail info
	Avail VirtioRingAvail `json:"avail"`
	// Used VRingUsed info
	Used VirtioRingUsed `json:"used"`
}

// QueryVirtioQueueElement
//
// Return the information about a VirtQueue's VirtQueueElement
type QueryVirtioQueueElement struct {
	// Path VirtIODevice canonical QOM path
	Path string `json:"path"`
	// Queue VirtQueue index to examine
	Queue uint16 `json:"queue"`
	// Index Index of the element in the queue (default: head of the queue)
	Index *uint16 `json:"index,omitempty"`
}

func (QueryVirtioQueueElement) Command() string {
	return "x-query-virtio-queue-element"
}

func (cmd QueryVirtioQueueElement) Execute(ctx context.Context, client api.Client) (VirtioQueueElement, error) {
	var ret VirtioQueueElement

	return ret, client.Execute(ctx, "x-query-virtio-queue-element", cmd, &ret)
}

// IOThreadVirtQueueMapping
//
// Describes the subset of virtqueues assigned to an IOThread.
type IOThreadVirtQueueMapping struct {
	// Iothread the id of IOThread object
	Iothread string `json:"iothread"`
	// Vqs an optional array of virtqueue indices that will be handled by this IOThread. When absent, virtqueues are assigned round-robin across all IOThreadVirtQueueMappings provided. Either all IOThreadVirtQueueMappings must have @vqs or none of them must have it.
	Vqs []uint16 `json:"vqs,omitempty"`
}

// DummyVirtioForceArrays
//
// Not used by QMP; hack to let us use IOThreadVirtQueueMappingList internally
type DummyVirtioForceArrays struct {
	UnusedIothreadVqMapping []IOThreadVirtQueueMapping `json:"unused-iothread-vq-mapping"`
}

// QCryptodevBackendAlgType The supported algorithm types of a crypto device.
type QCryptodevBackendAlgType string

const (
	// QcryptodevBackendAlgSym symmetric encryption
	QcryptodevBackendAlgSym QCryptodevBackendAlgType = "sym"
	// QcryptodevBackendAlgAsym asymmetric Encryption
	QcryptodevBackendAlgAsym QCryptodevBackendAlgType = "asym"
)

// QCryptodevBackendServiceType The supported service types of a crypto device.
type QCryptodevBackendServiceType string

const (
	QcryptodevBackendServiceCipher   QCryptodevBackendServiceType = "cipher"
	QcryptodevBackendServiceHash     QCryptodevBackendServiceType = "hash"
	QcryptodevBackendServiceMac      QCryptodevBackendServiceType = "mac"
	QcryptodevBackendServiceAead     QCryptodevBackendServiceType = "aead"
	QcryptodevBackendServiceAkcipher QCryptodevBackendServiceType = "akcipher"
)

// QCryptodevBackendType The crypto device backend type
type QCryptodevBackendType string

const (
	// QcryptodevBackendTypeBuiltin the QEMU builtin support
	QcryptodevBackendTypeBuiltin QCryptodevBackendType = "builtin"
	// QcryptodevBackendTypeVhostUser vhost-user
	QcryptodevBackendTypeVhostUser QCryptodevBackendType = "vhost-user"
	// QcryptodevBackendTypeLkcf Linux kernel cryptographic framework
	QcryptodevBackendTypeLkcf QCryptodevBackendType = "lkcf"
)

// QCryptodevBackendClient
//
// Information about a queue of crypto device.
type QCryptodevBackendClient struct {
	// Queue the queue index of the crypto device
	Queue uint32 `json:"queue"`
	// Type the type of the crypto device
	Type QCryptodevBackendType `json:"type"`
}

// QCryptodevInfo
//
// Information about a crypto device.
type QCryptodevInfo struct {
	// Id the id of the crypto device
	Id string `json:"id"`
	// Service supported service types of a crypto device
	Service []QCryptodevBackendServiceType `json:"service"`
	// Client the additional information of the crypto device
	Client []QCryptodevBackendClient `json:"client"`
}

// QueryCryptodev
//
// Returns information about current crypto devices.
type QueryCryptodev struct {
}

func (QueryCryptodev) Command() string {
	return "query-cryptodev"
}

func (cmd QueryCryptodev) Execute(ctx context.Context, client api.Client) ([]QCryptodevInfo, error) {
	var ret []QCryptodevInfo

	return ret, client.Execute(ctx, "query-cryptodev", cmd, &ret)
}

// CxlEventLog CXL has a number of separate event logs for different types of events. Each such event log is handled and signaled independently.
type CxlEventLog string

const (
	// CxlEventLogInformational Information Event Log
	CxlEventLogInformational CxlEventLog = "informational"
	// CxlEventLogWarning Warning Event Log
	CxlEventLogWarning CxlEventLog = "warning"
	// CxlEventLogFailure Failure Event Log
	CxlEventLogFailure CxlEventLog = "failure"
	// CxlEventLogFatal Fatal Event Log
	CxlEventLogFatal CxlEventLog = "fatal"
)

// CxlInjectGeneralMediaEvent
//
// Inject an event record for a General Media Event (CXL r3.0 8.2.9.2.1.1). This event type is reported via one of the event logs specified via the log parameter.
type CxlInjectGeneralMediaEvent struct {
	// Path CXL type 3 device canonical QOM path
	Path string `json:"path"`
	// Log event log to add the event to
	Log CxlEventLog `json:"log"`
	// Flags Event Record Flags. See CXL r3.0 Table 8-42 Common Event Record Format, Event Record Flags for subfield definitions.
	Flags uint8 `json:"flags"`
	// Dpa Device Physical Address (relative to @path device). Note lower bits include some flags. See CXL r3.0 Table 8-43 General Media Event Record, Physical Address.
	Dpa uint64 `json:"dpa"`
	// Descriptor Memory Event Descriptor with additional memory event information. See CXL r3.0 Table 8-43 General Media Event Record, Memory Event Descriptor for bit definitions.
	Descriptor uint8 `json:"descriptor"`
	// Type Type of memory event that occurred. See CXL r3.0 Table 8-43 General Media Event Record, Memory Event Type for possible values.
	Type uint8 `json:"type"`
	// TransactionType Type of first transaction that caused the event to occur. See CXL r3.0 Table 8-43 General Media Event Record, Transaction Type for possible values.
	TransactionType uint8 `json:"transaction-type"`
	// Channel The channel of the memory event location. A channel is an interface that can be independently accessed for a transaction.
	Channel *uint8 `json:"channel,omitempty"`
	// Rank The rank of the memory event location. A rank is a set of memory devices on a channel that together execute a transaction.
	Rank *uint8 `json:"rank,omitempty"`
	// Device Bitmask that represents all devices in the rank associated with the memory event location.
	Device *uint32 `json:"device,omitempty"`
	// ComponentId Device specific component identifier for the event. May describe a field replaceable sub-component of the device.
	ComponentId *string `json:"component-id,omitempty"`
}

func (CxlInjectGeneralMediaEvent) Command() string {
	return "cxl-inject-general-media-event"
}

func (cmd CxlInjectGeneralMediaEvent) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "cxl-inject-general-media-event", cmd, nil)
}

// CxlInjectDramEvent
//
// Inject an event record for a DRAM Event (CXL r3.0 8.2.9.2.1.2). This event type is reported via one of the event logs specified via the log parameter.
type CxlInjectDramEvent struct {
	// Path CXL type 3 device canonical QOM path
	Path string `json:"path"`
	// Log Event log to add the event to
	Log CxlEventLog `json:"log"`
	// Flags Event Record Flags. See CXL r3.0 Table 8-42 Common Event Record Format, Event Record Flags for subfield definitions.
	Flags uint8 `json:"flags"`
	// Dpa Device Physical Address (relative to @path device). Note lower bits include some flags. See CXL r3.0 Table 8-44 DRAM Event Record, Physical Address.
	Dpa uint64 `json:"dpa"`
	// Descriptor Memory Event Descriptor with additional memory event information. See CXL r3.0 Table 8-44 DRAM Event Record, Memory Event Descriptor for bit definitions.
	Descriptor uint8 `json:"descriptor"`
	// Type Type of memory event that occurred. See CXL r3.0 Table 8-44 DRAM Event Record, Memory Event Type for possible values.
	Type uint8 `json:"type"`
	// TransactionType Type of first transaction that caused the event to occur. See CXL r3.0 Table 8-44 DRAM Event Record, Transaction Type for possible values.
	TransactionType uint8 `json:"transaction-type"`
	// Channel The channel of the memory event location. A channel is an interface that can be independently accessed for a transaction.
	Channel *uint8 `json:"channel,omitempty"`
	// Rank The rank of the memory event location. A rank is a set of memory devices on a channel that together execute a transaction.
	Rank *uint8 `json:"rank,omitempty"`
	// NibbleMask Identifies one or more nibbles that the error affects
	NibbleMask *uint32 `json:"nibble-mask,omitempty"`
	// BankGroup Bank group of the memory event location, incorporating a number of Banks.
	BankGroup *uint8 `json:"bank-group,omitempty"`
	// Bank Bank of the memory event location. A single bank is accessed per read or write of the memory.
	Bank *uint8 `json:"bank,omitempty"`
	// Row Row address within the DRAM.
	Row *uint32 `json:"row,omitempty"`
	// Column Column address within the DRAM.
	Column *uint16 `json:"column,omitempty"`
	// CorrectionMask Bits within each nibble. Used in order of bits set in the nibble-mask. Up to 4 nibbles may be covered.
	CorrectionMask []uint64 `json:"correction-mask,omitempty"`
}

func (CxlInjectDramEvent) Command() string {
	return "cxl-inject-dram-event"
}

func (cmd CxlInjectDramEvent) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "cxl-inject-dram-event", cmd, nil)
}

// CxlInjectMemoryModuleEvent
//
// Inject an event record for a Memory Module Event (CXL r3.0 8.2.9.2.1.3). This event includes a copy of the Device Health info at the time of the event.
type CxlInjectMemoryModuleEvent struct {
	// Path CXL type 3 device canonical QOM path
	Path string `json:"path"`
	// Log Event Log to add the event to
	Log CxlEventLog `json:"log"`
	// Flags Event Record Flags. See CXL r3.0 Table 8-42 Common Event Record Format, Event Record Flags for subfield definitions.
	Flags uint8 `json:"flags"`
	// Type Device Event Type. See CXL r3.0 Table 8-45 Memory Module Event Record for bit definitions for bit definiions.
	Type uint8 `json:"type"`
	// HealthStatus Overall health summary bitmap. See CXL r3.0 Table 8-100 Get Health Info Output Payload, Health Status for bit definitions.
	HealthStatus uint8 `json:"health-status"`
	// MediaStatus Overall media health summary. See CXL r3.0 Table 8-100 Get Health Info Output Payload, Media Status for bit definitions.
	MediaStatus uint8 `json:"media-status"`
	// AdditionalStatus See CXL r3.0 Table 8-100 Get Health Info Output Payload, Additional Status for subfield definitions.
	AdditionalStatus uint8 `json:"additional-status"`
	// LifeUsed Percentage (0-100) of factory expected life span.
	LifeUsed uint8 `json:"life-used"`
	// Temperature Device temperature in degrees Celsius.
	Temperature int16 `json:"temperature"`
	// DirtyShutdownCount Number of times the device has been unable to determine whether data loss may have occurred.
	DirtyShutdownCount uint32 `json:"dirty-shutdown-count"`
	// CorrectedVolatileErrorCount Total number of correctable errors in volatile memory.
	CorrectedVolatileErrorCount uint32 `json:"corrected-volatile-error-count"`
	// CorrectedPersistentErrorCount Total number of correctable errors in persistent memory
	CorrectedPersistentErrorCount uint32 `json:"corrected-persistent-error-count"`
}

func (CxlInjectMemoryModuleEvent) Command() string {
	return "cxl-inject-memory-module-event"
}

func (cmd CxlInjectMemoryModuleEvent) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "cxl-inject-memory-module-event", cmd, nil)
}

// CxlInjectPoison
//
// Poison records indicate that a CXL memory device knows that a particular memory region may be corrupted. This may be because of locally detected errors (e.g. ECC failure) or poisoned writes received from other components in the system. This injection mechanism enables testing of the OS handling of poison records which may be queried via the CXL mailbox.
type CxlInjectPoison struct {
	// Path CXL type 3 device canonical QOM path
	Path string `json:"path"`
	// Start Start address; must be 64 byte aligned.
	Start uint64 `json:"start"`
	// Length Length of poison to inject; must be a multiple of 64 bytes.
	Length uint64 `json:"length"`
}

func (CxlInjectPoison) Command() string {
	return "cxl-inject-poison"
}

func (cmd CxlInjectPoison) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "cxl-inject-poison", cmd, nil)
}

// CxlUncorErrorType Type of uncorrectable CXL error to inject. These errors are reported via an AER uncorrectable internal error with additional information logged at the CXL device.
type CxlUncorErrorType string

const (
	// CxlUncorErrorTypeCacheDataParity Data error such as data parity or data ECC error CXL.cache
	CxlUncorErrorTypeCacheDataParity CxlUncorErrorType = "cache-data-parity"
	// CxlUncorErrorTypeCacheAddressParity Address parity or other errors associated with the address field on CXL.cache
	CxlUncorErrorTypeCacheAddressParity CxlUncorErrorType = "cache-address-parity"
	// CxlUncorErrorTypeCacheBeParity Byte enable parity or other byte enable errors on CXL.cache
	CxlUncorErrorTypeCacheBeParity CxlUncorErrorType = "cache-be-parity"
	// CxlUncorErrorTypeCacheDataEcc ECC error on CXL.cache
	CxlUncorErrorTypeCacheDataEcc CxlUncorErrorType = "cache-data-ecc"
	// CxlUncorErrorTypeMemDataParity Data error such as data parity or data ECC error on CXL.mem
	CxlUncorErrorTypeMemDataParity CxlUncorErrorType = "mem-data-parity"
	// CxlUncorErrorTypeMemAddressParity Address parity or other errors associated with the address field on CXL.mem
	CxlUncorErrorTypeMemAddressParity CxlUncorErrorType = "mem-address-parity"
	// CxlUncorErrorTypeMemBeParity Byte enable parity or other byte enable errors on CXL.mem.
	CxlUncorErrorTypeMemBeParity CxlUncorErrorType = "mem-be-parity"
	// CxlUncorErrorTypeMemDataEcc Data ECC error on CXL.mem.
	CxlUncorErrorTypeMemDataEcc CxlUncorErrorType = "mem-data-ecc"
	// CxlUncorErrorTypeReinitThreshold REINIT threshold hit.
	CxlUncorErrorTypeReinitThreshold CxlUncorErrorType = "reinit-threshold"
	// CxlUncorErrorTypeRsvdEncoding Received unrecognized encoding.
	CxlUncorErrorTypeRsvdEncoding CxlUncorErrorType = "rsvd-encoding"
	// CxlUncorErrorTypePoisonReceived Received poison from the peer.
	CxlUncorErrorTypePoisonReceived CxlUncorErrorType = "poison-received"
	// CxlUncorErrorTypeReceiverOverflow Buffer overflows (first 3 bits of header log indicate which)
	CxlUncorErrorTypeReceiverOverflow CxlUncorErrorType = "receiver-overflow"
	// CxlUncorErrorTypeInternal Component specific error
	CxlUncorErrorTypeInternal CxlUncorErrorType = "internal"
	// CxlUncorErrorTypeCxlIdeTx Integrity and data encryption tx error.
	CxlUncorErrorTypeCxlIdeTx CxlUncorErrorType = "cxl-ide-tx"
	// CxlUncorErrorTypeCxlIdeRx Integrity and data encryption rx error.
	CxlUncorErrorTypeCxlIdeRx CxlUncorErrorType = "cxl-ide-rx"
)

// CXLUncorErrorRecord
//
// Record of a single error including header log.
type CXLUncorErrorRecord struct {
	// Type Type of error
	Type CxlUncorErrorType `json:"type"`
	// Header 16 DWORD of header.
	Header []uint32 `json:"header"`
}

// CxlInjectUncorrectableErrors
//
// Command to allow injection of multiple errors in one go. This allows testing of multiple header log handling in the OS.
type CxlInjectUncorrectableErrors struct {
	// Path CXL Type 3 device canonical QOM path
	Path string `json:"path"`
	// Errors Errors to inject
	Errors []CXLUncorErrorRecord `json:"errors"`
}

func (CxlInjectUncorrectableErrors) Command() string {
	return "cxl-inject-uncorrectable-errors"
}

func (cmd CxlInjectUncorrectableErrors) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "cxl-inject-uncorrectable-errors", cmd, nil)
}

// CxlCorErrorType Type of CXL correctable error to inject
type CxlCorErrorType string

const (
	// CxlCorErrorTypeCacheDataEcc Data ECC error on CXL.cache
	CxlCorErrorTypeCacheDataEcc CxlCorErrorType = "cache-data-ecc"
	// CxlCorErrorTypeMemDataEcc Data ECC error on CXL.mem
	CxlCorErrorTypeMemDataEcc CxlCorErrorType = "mem-data-ecc"
	// CxlCorErrorTypeCrcThreshold Component specific and applicable to 68 byte Flit mode only.
	CxlCorErrorTypeCrcThreshold   CxlCorErrorType = "crc-threshold"
	CxlCorErrorTypeRetryThreshold CxlCorErrorType = "retry-threshold"
	// CxlCorErrorTypeCachePoisonReceived Received poison from a peer on CXL.cache.
	CxlCorErrorTypeCachePoisonReceived CxlCorErrorType = "cache-poison-received"
	// CxlCorErrorTypeMemPoisonReceived Received poison from a peer on CXL.mem
	CxlCorErrorTypeMemPoisonReceived CxlCorErrorType = "mem-poison-received"
	// CxlCorErrorTypePhysical Received error indication from the physical layer.
	CxlCorErrorTypePhysical CxlCorErrorType = "physical"
)

// CxlInjectCorrectableError
//
// Command to inject a single correctable error. Multiple error injection of this error type is not interesting as there is no associated header log. These errors are reported via AER as a correctable internal error, with additional detail available from the CXL device.
type CxlInjectCorrectableError struct {
	// Path CXL Type 3 device canonical QOM path
	Path string `json:"path"`
	// Type Type of error.
	Type CxlCorErrorType `json:"type"`
}

func (CxlInjectCorrectableError) Command() string {
	return "cxl-inject-correctable-error"
}

func (cmd CxlInjectCorrectableError) Execute(ctx context.Context, client api.Client) error {
	return client.Execute(ctx, "cxl-inject-correctable-error", cmd, nil)
}
