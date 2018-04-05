---
title:  "Install Qualcomm E2400 Ethernet on Ubuntu"
date:   2016-04-20 00:00:00
description: Modify Qualcomm Atheros alx driver to recognize E2400 chip
---

A recent updated screwed over some driver and make my Ethernet device unrecognizable on
Ubuntu 14.04. After some googling and helped from a colleagues,
I got it fixed by patching the driver :)

Below are the steps, plus some explanation of what it does so new Ubuntu/Linux
users can get started.

## The issue

The Ethernet adapter (Qualcomm E2400) is attached to the mainboard, and it was not recognized by
Ubuntu only.

A quick call to `lspci -k` shows this

{% highlight bash %}
03:00.0 Ethernet controller: Qualcomm Atheros Device e0a1 (rev 10)
        Subsystem: Micro-Star International Co., Ltd. [MSI] Device 7976
{% endhighlight %}

There is no `Kernel driver in use: `. This means that Ubuntu itself wasn't able
to find and load a driver for it.

Checking only, it seems that E2400 use the `alx` driver, but the driver doesn't seem to recognize
the device code. So people recommend to patch the driver

## Driver patching

> [alx driver](http://www.linuxfoundation.org/collaborate/workgroups/networking/alx)
>
> The manufacter seems have provided the instruction on
[killer ethernet support page](http://www.killernetworking.com/support/knowledge-base/17-linux/21-killer-e2400-in-linux-ubuntu-debian)

The Killer E2400 has been confirmed to work fully by modifying and loading the alx driver
in Ubuntu 14.04 and 15.04.
These changes have been upstreamed and should work natively in future Linux kernels.

> So the first step is getting the driver source, which has already been included in the Linux kernel I am using

### 1. Make sure your complier enviroment is ready.

{% highlight bash %}
$ apt-get source linux-image-$(uname -r)
{% endhighlight %}

Alternative: If you do not have an alternative network connection, you can download the linux kernel source manually from Ubuntu's site:

<http://packages.ubuntu.com/vivid/kernel/>

Download and unpack it.

{% highlight bash %}
$ cd ~/linux-image-$(uname-r)
$ make oldconfig
$ make prepare
$ make scripts
$ apt-get install linux-headers-$(uname -r)
{% endhighlight %}

### 2. Prevent the message "no symbol version for module_layout" when loading the module with insmod or modprobe.

{% highlight bash %}
$ cd ~/linux-source
$ cp -v /usr/src/linux-headers-$(uname -r)/Module.symvers .
{% endhighlight %}

### 3. Make changes to main.c and reg.h files in ./drivers/net/ethernet/atheros/alx :

{% highlight c %}
// diff --git a/main.c.orig b/main.c
// --- a/main.c.orig
// +++ b/main.c
// @@ -1537,6 +1537,7 @@ static const struct pci_device_id alx_pci_tbl[] = {
        { PCI_VDEVICE(ATTANSIC, ALX_DEV_ID_AR8162),
          .driver_data = ALX_DEV_QUIRK_MSI_INTX_DISABLE_BUG },
        { PCI_VDEVICE(ATTANSIC, ALX_DEV_ID_AR8171) },
        { PCI_VDEVICE(ATTANSIC, ALX_DEV_ID_E2400) },
        { PCI_VDEVICE(ATTANSIC, ALX_DEV_ID_AR8172) },
        {}
};
{% endhighlight %}

{% highlight bash %}
diff --git a/reg.h.orig b/reg.h
index af006b4..1396483 100644
--- a/reg.h.orig
+++ b/reg.h
@@ -39,6 +39,7 @@
{% endhighlight %}
{% highlight c %}
#define ALX_DEV_ID_E2200                                0xe091
#define ALX_DEV_ID_AR8162                               0x1090
#define ALX_DEV_ID_AR8171                               0x10A1
#define ALX_DEV_ID_E2400                                0xe0A1
#define ALX_DEV_ID_AR8172                               0x10A0

/\* rev definition,
{% endhighlight %}

### 4. Build and install module

{% highlight bash %}
$ cd drivers/net/ethernet/alx
$ make -C /lib/modules/$(uname -r)/build M=$(pwd) modules
$ make -C /lib/modules/$(uname -r)/build M=$(pwd) modules_install
$ modprobe -r alx
$ depmod
$ modprobe -v alx
{% endhighlight %}

> One interesting thing here is that my desktop seems to put the driver into `/lib/modules/3.19.0-58-generic/extra`, instead of going to `/lib/modules/3.19.0-58-generic/kernel/drivers/net/ethernet/atheros/alx/alx.lo`
>
> So I forced copy it instead

After that, all is good :)
