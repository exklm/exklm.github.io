---
layout: post
title: IoT platform simulation
tags: [IoT, fault-injection]
---

This is an on-going personal project to build a system to test IoT applications.

Typically, IoT platorm will transfer data (from sensors) to somewhere, as well as receiving commands and perform some actions. When testing such systems, concerns are:
- reliability: would the system malfunction if network is loss or if individual devices malfunction.
- security: what are the attack surfaces & their impact.

These are big and intractable problems, however. My smaller plans is as followed:

- Get IoT nodes that run Linux
- Get a wireless router that run OpenWrt
- Get IoT nodes to connect to the network.
- Have an external process that can control the network (via `tc` on OpenWrt)
  and can randomly cause nodes to fault (think Jepsen)
- Script tests that setup something, then cause random failures
- Observe and analyze results (hopefully with scripts)

Hopefully, with this, we can simulates IoT running on both 3G and wireless/wired setup.

## Setup

### IoT nodes

- 2x Raspberry Pi 3 (model B)
- 2x Raspberry Pi Zero

Both models are selected because they offer Wifi chipset.
This result in less wiring and doesn't require a switch, too.

Since it's a bit hard to plug things into the Pi Zero (keyboard and HDMI), I needed Pi 3.
However, one Pi 3 should have been enough. Pi Zero seems to be the best candidate for this test
setup, since it's cheap and have a small physical profile.

I tried to install a custom Ubuntu server image for ARM on them, but haven't managed to boot up.
Hence, Raspbian was installed.

After booting, the Wifi can be configured by adding this into `/etc/wpa_supplicant/wpa_supplicant.conf`:

```
network={
  ssid="openwrt.network"
  psk="password"
}
```

After reboot, the Pi got their IP from the router.

### Power supply

- [Anker AH231 USB 3.0 10-Port Hub]
- 2.0 A micro USB cable (pack of 6)

Granted, this might not guarantee enough power supply for the Pi 3 (max at 2.1 A, instead of 2.5),
but things has been working okay so far

### Router

- Router: Netgear N600 with WNDR3700v4

I picked this specifically because the chipset is known to support OpenWRT well.
Flashing a OpenWRT image on this seems easy, as the stock router provided GUI to flash itself.
Hence, all I needed to do was downloading an image & use the web UI.

### Control node

- a VM

