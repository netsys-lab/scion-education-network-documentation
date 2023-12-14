Troubleshooting
=======================================

Open-Source Stack
------------
  
Check SCION service status
------------
  
.. code-block:: console
  
  sudo systemctl list-dependencies scionlab.target
  
This should show all entries as green. If there are any failed services in this list, start troubleshooting

Check the border router interface status
------------

Inspect the border router’s log (using sudo journalctl -u scion-border-router@br-1.service) to check that the bidirectional-forwarding detection (bfd) handshake completed and the interfaces are “active”:

Check that the log mentions Transitioned from state ... to state Up, not followed by a later ... to state Down.

Alternatively, you can check the same information in metrics of the border router. For the default SCIONLab User AS setup, this is exposed on localhost, port 30401.

.. code-block:: console

  $ curl -sfS localhost:30401/metrics | grep router_interface_up
  # HELP router_interface_up Either zero or one depending on whether the interface is up.
  # TYPE router_interface_up gauge
  router_interface_up{interface="1",isd_as="1-ff00:0:112",neighbor_isd_as="1-ff00:0:110"} 1


Check that beacons are registered
------------

An AS needs to receive path construction beacons from it’s upstream provider AS(es) in order to be able to communicate.

Inspect the control service’s log (using sudo journalctl -u scion-control-service@cs-1.service) to check that beacons are registered successfully.

Check that you find entries Registered beacons ..., with "count": 1 (any non-zero count).

Alternatively, you can check the same information in metrics of the control service, exposed by default on localhost, port 30454.

.. code-block:: console

  $ curl -sfS localhost:30454/metrics | grep control_beaconing_received_beacons_total
  # HELP control_beaconing_received_beacons_total Total number of beacons received.
  # TYPE control_beaconing_received_beacons_total counter
  control_beaconing_received_beacons_total{ingress_interface="41",neighbor_isd_as="1-ff00:0:110",result="ok_updated"} 38


Showpaths
------------
SCION provides youthe functionality to show the available paths to a particular AS. Check your connectivity by showing the available paths to one of the core ASes in the network:

.. code-block:: console

  scion showpaths 71-2:0:35
  scion showpaths 71-20965


Ping
------------
Ping somebody! Run scion ping to send an “SCMP echo request”; this is just like the ping command for IP.

The syntax is:

scion ping [destination SCION address]
where a SCION address has the form ISD-AS,IP. An example of pinging a host in the network would look like this:

.. code-block:: console

  scion ping 71-2:0:35,127.0.0.1
