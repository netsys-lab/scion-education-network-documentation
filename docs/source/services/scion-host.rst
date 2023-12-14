SCION Endhost Installer
==============================================

This site covers the setup of SCION on pure endhosts, to achieve native SCION connectivity. We provide a tool called scion-host [1] that automates the installation of SCION on endhosts. It requires a configured bootstrap-server in your SCION AS [2]. It consists of the following core steps:

* Installation of the SCION Binaries onto the machine.
* Connect to the given bootstrap-server to fetch a SCION configuration.
* Run all required SCION endhost services.

So far we provide preliminary support to run SCION on the following operating systems:

* Linux (amd64)
* Windows (amd64) 
* MacOS (arm64/amd64)

You can follow the guidelines in the Readme [1] to install SCION on your endhost. Please note that the installation process requires root privileges at the moment. Furthermore, we observed some issues with the prebuilt binaries for Windows/MacOS (antivirus systems may think it is a virus, and on MacOS a signature for the binary is misssing). Consequently, we recommend to build the binaries from source using the makefile on the host on which you want to run SCION.

We currently work on providing signed installers for all operating systems. Furthermore, we are working on a more user-friendly installation process, which does not require root privileges. Please stay tuned for updates.

Related Links:
----------------
* `[1] SCION Installer/Binaries <https://github.com/netsys-lab/scion-host>`_
* `[2] Endhost Bootstrapping Server <https://github.com/netsys-lab/bootstrap-server>`_
