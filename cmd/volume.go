package main

import corev1 "k8s.io/api/core/v1"

// -v /var/lib/lxc/lxcfs/proc/cpuinfo:/proc/cpuinfo:ro
// -v /var/lib/lxc/lxcfs/proc/diskstats:/proc/diskstats:ro
// -v /var/lib/lxc/lxcfs/proc/loadavg:/proc/loadavg:ro
// -v /var/lib/lxc/lxcfs/proc/meminfo:/proc/meminfo:ro
// -v /var/lib/lxc/lxcfs/proc/stat:/proc/stat:ro
// -v /var/lib/lxc/lxcfs/proc/swaps:/proc/swaps:ro
// -v /var/lib/lxc/lxcfs/proc/uptime:/proc/uptime:ro
// -v /var/lib/lxc/lxcfs/sys/devices/system/cpu/online:/sys/devices/system/cpu/online:ro
// -v /var/lib/lxc/:/var/lib/lxc/:ro

const lxcfsVol = "lxcfs"

var volumeMountsTemplate = []corev1.VolumeMount{

	{
		Name:      "lxcfs-proc-cpuinfo",
		MountPath: "/proc/cpuinfo",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-meminfo",
		MountPath: "/proc/meminfo",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-diskstats",
		MountPath: "/proc/diskstats",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-stat",
		MountPath: "/proc/stat",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-swaps",
		MountPath: "/proc/swaps",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-uptime",
		MountPath: "/proc/uptime",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-proc-loadavg",
		MountPath: "/proc/loadavg",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs-sys-devices-system-cpu-online",
		MountPath: "/sys/devices/system/cpu/online",
		ReadOnly:  true,
	},
	{
		Name:      "lxcfs",
		MountPath: "/var/lib/lxc/",
		ReadOnly:  true,
		MountPropagation: func() *corev1.MountPropagationMode {
			pt := corev1.MountPropagationHostToContainer
			return &pt
		}(),
	},
}

var volumesTemplate = []corev1.Volume{
	{
		Name: "lxcfs-proc-cpuinfo",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxc/lxcfs/proc/cpuinfo",
				Type: func() *corev1.HostPathType {
					pt := corev1.HostPathFile
					return &pt
				}(),
			},
		},
	},
	{
		Name: "lxcfs-proc-diskstats",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxc/lxcfs/proc/diskstats",
				Type: func() *corev1.HostPathType {
					pt := corev1.HostPathFile
					return &pt
				}(),
			},
		},
	},
	{
		Name: "lxcfs-proc-meminfo",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxc/lxcfs/proc/meminfo",
				Type: func() *corev1.HostPathType {
					pt := corev1.HostPathFile
					return &pt
				}(),
			},
		},
	},
	{
		Name: "lxcfs-proc-stat",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxc/lxcfs/proc/stat",
				Type: func() *corev1.HostPathType {
					pt := corev1.HostPathFile
					return &pt
				}(),
			},
		},
	},
	{
		Name: "lxcfs-proc-swaps",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxc/lxcfs/proc/swaps",
				Type: func() *corev1.HostPathType {
					pt := corev1.HostPathFile
					return &pt
				}(),
			},
		},
	},
	{
		Name: "lxcfs-proc-uptime",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxc/lxcfs/proc/uptime",
				Type: func() *corev1.HostPathType {
					pt := corev1.HostPathFile
					return &pt
				}(),
			},
		},
	},
	{
		Name: "lxcfs-proc-loadavg",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxc/lxcfs/proc/loadavg",
				Type: func() *corev1.HostPathType {
					pt := corev1.HostPathFile
					return &pt
				}(),
			},
		},
	},
	{
		Name: "lxcfs-sys-devices-system-cpu-online",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxc/lxcfs/sys/devices/system/cpu/online",
				Type: func() *corev1.HostPathType {
					pt := corev1.HostPathFile
					return &pt
				}(),
			},
		},
	},
	{
		Name: lxcfsVol,
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxc/",
				Type: func() *corev1.HostPathType {
					pt := corev1.HostPathDirectoryOrCreate
					return &pt
				}(),
			},
		},
	},
}
