SCION Endhost Bootstrapping
=======================================

SCION endhost bootstrapping is a mechanism to automatically configure SCION on endhost machines within a SCION AS [1],[2]. For this, a central bootstrapping server instance needs to be setup within your SCION AS.

The bootstrapping server is a simple HTTP server that serves the endhost bootstrapping configuration files that are required to run SCION on an endhost.

The most straight forward way to setup the bootstrapping server is to use the boostrap-server repository [3] and run this setup on one of the machines that runs your SCION AS. This ensures that all configurations that are served to SCION endhosts match the current AS configuration.

The boostrap-server requires an IP address that is reachable by all potential SCION hosts within the AS. This is typically an IP address of the official IP range of the institution. Of course we can help setting finding this IP and setting up the server instance, please contact us in case any questions occur: marten.gartner@ovgu.de.

Once the bootstrap-server is set up, please follow the guide on SCION Endhost Installers to bootstrap SCION on your endhosts.

Related Links:
--------------
* `[1] Endhost Bootstrapping Code <https://github.com/netsec-ethz/bootstrapper>`_
* `[2] Endhost Bootstrapping Design <https://github.com/scionproto/scion/blob/master/doc/dev/design/endhost-bootstrap.rst>`_
* `[3] Endhost Bootstrapping Server <https://github.com/netsys-lab/bootstrap-server>`_