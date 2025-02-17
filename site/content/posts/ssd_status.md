---
title: "SSD drive health status"
date: 2025-02-17T12:03:04+01:00
draft: false
tags: ["linux", "ssd"]
---

Should you replace your SSD drive? Here is a simple way to check the health status of your SSD drive.
```bash
sudo dnf install -y nvme-cli
sudo nvme smart-log /dev/nvme0
```

This will show you the following information:
```
Smart Log for NVME device:nvme0 namespace-id:ffffffff
critical_warning			: 0
temperature				: 43 °C (316 K)
available_spare				: 100%
available_spare_threshold		: 50%
percentage_used				: 1%
endurance group critical warning summary: 0
Data Units Read				: 10,529,778 (5.39 TB)
Data Units Written			: 36,729,164 (18.81 TB)
host_read_commands			: 179,044,182
host_write_commands			: 689,140,626
controller_busy_time			: 1,040
power_cycles				: 386
power_on_hours				: 6,494
unsafe_shutdowns			: 60
media_errors				: 0
num_err_log_entries			: 0
Warning Temperature Time		: 0
Critical Composite Temperature Time	: 0
Temperature Sensor 1			: 43 °C (316 K)
Temperature Sensor 2			: 44 °C (317 K)
Thermal Management T1 Trans Count	: 0
Thermal Management T2 Trans Count	: 0
Thermal Management T1 Total Time	: 0
Thermal Management T2 Total Time	: 0
```

`available_spare` - shows the percentage of spare memory blocks available on your SSD.

`available_spare_threshold` - is the critical threshold, expressed as a percentage.
When your `available_spare` drops below this threshold, it's a warning sign that
your drive is nearing the end of its life.

`percentage_used` - indicates the percentage of the drive's total estimated lifespan
that has been used up.
