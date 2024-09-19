Services
=======================================

We provide the following services to facilitate an easy setup of endhosts within SCIERA:

SCION Endhost Bootstrapping
----------------

SCION endhost bootstrapping is a mechanism to automatically configure SCION on endhost machines within a SCION AS (cf. `Code <https://github.com/netsec-ethz/bootstrapper>`_, `Design <https://github.com/scionproto/scion/blob/master/doc/dev/design/endhost-bootstrap.rst>`_). For this, a central bootstrapping server instance needs to be setup within your SCION AS.

The bootstrapping server is a simple HTTP server that serves the endhost bootstrapping configuration files that are required to run SCION on an endhost.

The most straight forward way to setup the bootstrapping server is to use the `boostrap-server repository <https://github.com/netsys-lab/bootstrap-server>`_ and run this setup on one of the machines that runs your SCION AS. This ensures that all configurations that are served to SCION endhosts match the current AS configuration.

The boostrap-server requires an IP address that is reachable by all potential SCION hosts within the AS. This is typically an IP address of the official IP range of the institution. Of course we can help setting finding this IP and setting up the server instance, please contact us in case any questions occur: marten.gartner@ovgu.de.

Once the bootstrap-server is set up, please follow the guide on SCION Endhost Installers to bootstrap SCION on your endhosts.

SCION Endhost Installer
----------------

The SCION Endhost installer facilitates the setup of SCION on pure endhosts, to achieve native SCION connectivity. We provide a tool called `scion-host <https://github.com/netsys-lab/scion-host>`_ that automates the installation of SCION on endhosts. It requires a configured `bootstrap-server <https://github.com/netsys-lab/bootstrap-server>`_ in your SCION AS. It consists of the following core steps:

* Installation of the SCION Binaries onto the machine.
* Connect to the given bootstrap-server to fetch a SCION configuration.
* Run all required SCION endhost services.

So far we provide preliminary support to run SCION on the following operating systems:

* Linux (amd64)
* Windows (amd64) 
* MacOS (arm64/amd64)

You can follow the guidelines in the `Readme <https://github.com/netsys-lab/scion-host>`_ to install SCION on your endhost. Please note that the installation process requires root privileges at the moment. Furthermore, we observed some issues with the prebuilt binaries for Windows/MacOS (antivirus systems may think it is a virus, and on MacOS a signature for the binary is misssing). Consequently, we recommend to build the binaries from source using the makefile on the host on which you want to run SCION.

We currently work on providing signed installers for all operating systems. Furthermore, we are working on a more user-friendly installation process, which does not require root privileges. Please stay tuned for updates.
