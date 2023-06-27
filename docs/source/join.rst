Join the Network
=====

Connect to GEANT
----------------

tbd

Connect to BRIDGES
----------------

tbd

Install SCION
----------------
So far we recommend to use Ubuntu to run SCION. You can install the packages from the official repository:

On Ubuntu, you can install SCION from our .deb-packages by running:

.. code-block:: console

  sudo apt-get install apt-transport-https ca-certificates
  echo "deb [trusted=yes] https://packages.netsec.inf.ethz.ch/debian all main" | sudo tee /etc/apt/sources.list.d/scionlab.list
  sudo apt-get update
  sudo apt-get install scionlab

Apply Configuration
----------------

After connecting to the network and setting up your host, the next step is to apply your dedicated SCION configuration. You will receive it as .tar.gz file and an install script from us. Please copy both files somewhere on your host and run 

.. code-block:: console

  sudo ./install.sh host1.tar.gz

This will install the proper SCION configuration and start all the services.

Check Connectivity
----------------

After applying the configuration, SCION needs a moment to retrieve beacons and to create paths. After a minute, try to run 

.. code-block:: console

  scion showpaths 71-20965 

to see if your AS has SCION connectivity to the network. If not, please have a look at troubleshooting.
