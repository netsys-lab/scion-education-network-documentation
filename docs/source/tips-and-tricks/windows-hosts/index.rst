Windows as a SCION Host
=======================

SCION is usually deployed on Linux hosts. Since the open-source version of SCION is implemented in Go, other platforms such as Windows work as well. There are some limitations, but support is good enough to run SCION applications. This guide will walk you through running a Windows 11 PC as an end host in a SCION network.

Connecting to SCIERA
--------------------
A simple and quick way to connect to SCIERA points is the `scion-orchestrator <https://github.com/netsys-lab/scion-orchestrator>`_. If all you need is a one-off connection to a SCIERA network and you institution is in the :ref:`supported list<scion-endhost-installer>`, you can run SCION directly using scion-orchestrator. Just follow the instructions in the README. This guide is about installing SCION on Windows manually for a more permanent setup that is currently not possible with scion-orchestrator.

However, scion-orchestrator is still the easiest way to obtain the configuration files needed by every SCION host. The minimum configuration requires a file called ``topology.json`` that contains the IP addresses of local SCION routers among other critical information, and a set of certificates known as the Trust Root Certificate (TRC). The scion-orchestrator performs a bootstrapping process to locate and download these files. If you run the orchestrator and everything goes well, you will find the ``topology.json`` file in its config directory ``config/`` and one or more ``*.trc`` files in ``config/certs``.

Instead of running the orchestrator you can also grab the host configuration from a SCION bootstrapping server manually. Ask an administrator for the address of the bootstrapping server in your network or look it up in the `scion-orchestrator-releases <https://github.com/netsys-lab/scion-orchestrator-releases/tree/main/configs>`_ repository. Assuming the bootstrap server is at ``127.0.0.1:8000``, you can simply go to ``http://127.0.0.1:8000/topology`` in a webbrowser to obtain ``topology.json``. A list of available certificates is returned from ``http://127.0.0.1:8000/trcs``. Each certificate has an associated ISD, base number, and serial number. The certificates are available at ``http://127.0.0.1:8000/trcs/isd71-b1-s2`` for a hypothetical TRC for ISD 71, with base number 1, and serial number 2. Save the certificates to files named ``ISD71-B1-S2.trc`` and so on. In practice you don't have to download all TRCs offered, you only need to latest one from your ISD which is ISD71 in SCIERA's case. Here is a copy of :download:`SCIERA's current TRC <./ISD71-B1-S4.trc>` (base 1, serial 4).

It really doesn't matter how you obtain the ``topology.json`` and latest TRC, as long as you have them. TRCs are updated from time to time (maybe once a year in SCIERA) and you will have to download the newest one when the old certificates expire. scion-orchestrator can automate this process, but this is not fully supported on Windows yet.

Example of how ``topology.json`` can look like:

.. code-block:: json

  {
    "attributes": [],
    "isd_as": "1-ff00:0:111",
    "mtu": 1472,
    "test_dispatcher": true,
    "dispatched_ports": "31000-32767",
    "control_service": {
      "cs1-ff00_0_111-1": {
        "addr": "127.0.0.18:31006"
      }
    },
    "discovery_service": {
      "cs1-ff00_0_111-1": {
        "addr": "127.0.0.18:31006"
      }
    },
    "border_routers": {
      "br1-ff00_0_111-1": {
        "internal_addr": "127.0.0.17:31008",
        "interfaces": {
          "41": {
            "underlay": {
              "local": "127.0.0.5:50000",
              "remote": "127.0.0.4:50000"
            },
            "isd_as": "1-ff00:0:110",
            "link_to": "parent",
            "mtu": 1280
          }
        }
      }
    }
  }

Compile SCION
-------------
Since we're running a SCION end host on Windows and not a router, there are only two executables needed, ``daemon.exe`` and ``scion.exe``. "daemon" is a background process that manages SCION paths. We'll install it as a service later in this guide. "scion" is a command line tool that offers ping and traceroute functionality among other functions. SCION hosts usually also run a "dispatcher", but this component relies on exotic socket features and does not compile on Windows. Fortunately the dispatcher is no longer strictly required and is only kept to provide backwards compatibility. It also enables SCION hosts to respond to SCION ping messages, but this is also not strictly necessary to run applications.

In order to compile SCION you need to have a recent version of `Go <https://go.dev/>`_ installed. You also need git to clone the SCION repository. Check out a recent commit from the SCION repository and build the executables

.. code-block:: powershell

  go build -o ..\bin\daemon.exe .\daemon\cmd\daemon\
  go build -o ..\bin\scion.exe .\scion\cmd\scion\

Note that SCION release **v0.12.0 is not recent enough** for this guide. Use the current master branch or v0.13.0 when it releases.

Here is a PowerShell script that automates the process:

.. code-block:: powershell

  param (
      [string]$Repository = "https://github.com/scionproto/scion.git",
      [string]$Tag = "master",
      [string]$Path = "scion"
  )
  Set-StrictMode -Version Latest

  function Clone-Scion {
      param (
          [string]$Repository = "https://github.com/scionproto/scion.git",
          [string]$Tag = "master",
          [string]$Path
      )
      git clone --branch $Tag $Repository $Path
  }

  function Build-Scion {
      param (
          [string]$Path
      )
      $out = "$(Get-Location)\bin"
      $env:CGO_ENABLED = 0
      Push-Location $Path
      try {
          New-Item -Path $out -ItemType "directory" -Force | Out-Null
          # go build -o $out .\control\cmd\control\
          # Check-Error
          go build -o $out .\daemon\cmd\daemon\
          Check-Error
          # Dispatcher doesn't build on Windows
          #go build -o $out .\dispatcher\cmd\dispatcher\
          #Check-Error
          # go build -o $out .\router\cmd\router\
          # Check-Error
          go build -o $out .\scion\cmd\scion\
          Check-Error
          # go build -o $out .\scion-pki\cmd\scion-pki\
          # Check-Error
      } finally {
          Remove-Item Env:\CGO_ENABLED
          Pop-Location
      }
  }

  function Check-Error {
      if ($lastexitcode -ne 0) {
          throw "Go build failed"
      }
  }

  try {
      if (-Not (Test-Path -Path $Path)) {
          Clone-Scion -Path $Path -Repository $Repository -Tag $Tag
      }
      Build-SCION -Path $Path
      Copy-Item -Path bin -Destination .\integration -Recurse -Force
  } catch {
      Write-Host "Error building SCION:"
      Write-Host $_
  }

Install the resulting executables in ``%ProgramFiles%\scion`` and add this directory to the PATH environment variable.

Install the SCION Configuration
-------------------------------
The recommended location of the SCION configuration in ``%ProgramData%\scion``. Create a directory structure the look as follows:

.. code-block::

  scion
  ├── config
  │   ├── certs
  │   │   ├── ISD64-B1-S8.trc
  │   │   ├── ISD71-B1-S1.trc
  │   │   ├── ISD71-B1-S2.trc
  │   │   ├── ISD71-B1-S3.trc
  │   │   ├── ISD71-B1-S4.trc
  │   │   └── ...
  │   ├── sd.toml
  │   └── topology.json
  ├── log
  │   └── sd.log
  └── run
      ├── sd.path.db
      └── sd.trust.db

The files in ``scion/log/`` and ``scion/run/`` will be created by the daemon once it runs, just create empty directories for now.

``sd.toml`` is the SCION daemon's configuration file. At minimum it should look like this:

.. code-block:: toml

  [general]
  id = "sd"
  config_dir = "C:/ProgramData/scion/config"

  [path_db]
  connection = "C:/ProgramData/scion/run/sd.path.db"

  [trust_db]
  connection = "C:/ProgramData/scion/run/sd.trust.db"

  [log.console]
  level = "info"

It is important that ``topology.json`` and ``certs/`` are both in ``config_dir``, beyond that you can change the paths if you like.

Install the SCION Daemon as a Service
-------------------------------------
We'll install the daemon as a service that can run automatically on system boot.

SCION services make use of the Windows event log through Go's `eventlog package <https://pkg.go.dev/golang.org/x/sys@v0.29.0/windows/svc/eventlog>`_. Log messages written by the daemon have the event source ``SCION Daemon`` using ``%SystemRoot%\System32\EventCreate.exe`` as the `event message file <https://learn.microsoft.com/en-us/windows/win32/eventlog/message-files>`_. To add the SCION Daemon to the Application log, you can import the following registry file or use the program provided at the end of this section:

.. code-block::

  [HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\EventLog\Application\SCION Daemon]
  "CustomSource"=dword:00000001
  "EventMessageFile"=hex(2):25,00,53,00,79,00,73,00,74,00,65,00,6d,00,52,00,6f,\
    00,6f,00,74,00,25,00,5c,00,53,00,79,00,73,00,74,00,65,00,6d,00,33,00,32,00,\
    5c,00,45,00,76,00,65,00,6e,00,74,00,43,00,72,00,65,00,61,00,74,00,65,00,2e,\
    00,65,00,78,00,65,00,00,00
  "TypesSupported"=dword:00000007


Installing the daemon as a service can be done through the ``sc.exe`` command line tool or the PowerShell cmdlet ``New-Service``. There is also a Go program at the end of the section performing the same action. Remember to execute the commands from an elevated command prompt.

``sc.exe`` example (replace *demand* with *auto* for start on boot):

.. code-block:: powershell

  $COMMAND="${env:ProgramFiles}\scion\daemon.exe"
  $CONFIG="${env:ProgramData}\scion\config\sd.toml"
  $LOGFILE="${env:ProgramData}\scion\log\sd.log"
  sc.exe create "SCION Daemon" start=demand binPath="$COMMAND --config $CONFIG --logfile $LOGFILE"

``New-Service`` example (replace *Manual* with *Automatic* for start on boot):

.. code-block:: powershell

  $COMMAND="${env:ProgramFiles}\scion\daemon.exe"
  $CONFIG="${env:ProgramData}\scion\config\sd.toml"
  $LOGFILE="${env:ProgramData}\scion\log\sd.log"
  $DESC="Facilitates communication of SCION applications with the control service."
  New-Service -Name "SCION Daemon" -StartupType "Manual" -BinaryPathName "$COMMAND --config $CONFIG --logfile $LOGFILE" -Description "$DESC"


Once the service is installed you can also manage it from the service control panel. You can remove the service again using ``sc.exe`` or the cmdlet ``Remove-Service`` (requires at least PowerShell 6.0):

.. code-block:: powershell

  sc.exe delete "SCION Daemon"
  Remove-Service -Name "SCION Daemon"

Automatic Service Installation
""""""""""""""""""""""""""""""
The steps above are also carried out by the following :download:`Go program <./svcctrl.go>`:

.. code-block:: go

  package main

  import (
      "fmt"
      "os"
      "syscall"

      "golang.org/x/sys/windows/svc"
      "golang.org/x/sys/windows/svc/eventlog"
      "golang.org/x/sys/windows/svc/mgr"
  )

  const ERROR_INVALID_NAME = 123

  type Service struct {
      Name    string
      Exepath string
      Desc    string
      Config  string
      Logfile string
  }

  var services = []Service{
      {
          Name:    "SCION Daemon",
          Exepath: os.ExpandEnv("$ProgramFiles\\scion\\daemon.exe"),
          Desc:    "Facilitates communication of SCION applications with the control service.",
          Config:  os.ExpandEnv("$ProgramData\\scion\\config\\sd.toml"),
          Logfile: os.ExpandEnv("$ProgramData\\scion\\log\\sd.log"),
      },
  }

  func main() {
      var err error
      if len(os.Args) != 2 {
          fmt.Println("invalid number of arguments")
          return
      }
      if os.Args[1] == "install" {
          fmt.Println("Installing service")
          err = installServices()
      } else if os.Args[1] == "remove" {
          fmt.Println("Removing service")
          err = removeServices()
      } else if os.Args[1] == "start" {
          fmt.Println("Starting SCION")
          err = startServices()
      } else if os.Args[1] == "stop" {
          fmt.Println("Stopping SCION")
          err = stopServices()
      }
      if err != nil {
          fmt.Printf("error: %v\n", err)
      }
  }

  func installServices() error {
      m, err := mgr.Connect()
      if err != nil {
          return err
      }
      defer m.Disconnect()

      for _, service := range services {
          c := mgr.Config{
              StartType:   mgr.StartManual,
              Description: service.Desc,
          }
          s, err := m.CreateService(service.Name, service.Exepath, c,
              "--config", service.Config,
              "--logfile", service.Logfile)
          if err != nil {
              if errno, ok := err.(syscall.Errno); ok {
                  if uintptr(errno) == ERROR_INVALID_NAME {
                      fmt.Printf("service %s already exists\n", service.Name)
                      continue
                  }
              }
              return err
          }
          defer s.Close()

          err = eventlog.InstallAsEventCreate(service.Name, eventlog.Error|eventlog.Warning|eventlog.Info)
          if err != nil {
              s.Delete()
              return fmt.Errorf("installing event log source failed: %v", err)
          }
      }

      return nil
  }

  func removeServices() error {
      return applyToServices(func(s *mgr.Service) error {
          err := eventlog.Remove(s.Name)
          if err != nil {
              return fmt.Errorf("removing event log source failed: %v", err)
          }
          err = s.Delete()
          if err != nil {
              return err
          }
          return nil
      })
  }

  func startServices() error {
      return applyToServices(func(s *mgr.Service) error {
          return s.Start()
      })
  }

  func stopServices() error {
      return applyToServices(func(s *mgr.Service) error {
          _, err := s.Control(svc.Stop)
          return err
      })
  }

  func applyToServices(f func(*mgr.Service) error) error {
      m, err := mgr.Connect()
      if err != nil {
          return err
      }
      defer m.Disconnect()

      for _, service := range services {
          s, err := m.OpenService(service.Name)
          if err != nil {
              fmt.Printf("service %s not found\n", service.Name)
              continue
          }
          defer s.Close()
          err = f(s)
          if err != nil {
              return err
          }
      }
      return nil
  }

Use as follows:

.. code-block:: powershell

  go build -o ./svcctrl.exe
  .\svcctrl.exe install # install
  .\svcctrl.exe start   # start service
  .\svcctrl.exe stop    # stop service
  .\svcctrl.exe remove  # uninstall


Running the Service
"""""""""""""""""""
After you have installed the daemon service, start it from the command line or the control panel. In the event viewer (``eventvwr``) you should see a message in the Application protocol from ``SCION Daemon`` that says "Service started". If not check the logs to see what went wrong.

As a next step you can try the ``scion`` program. Running ``scion address`` in a terminal window prints your local SCION address. ``scion showpaths <ISD-ASN>`` lists available paths to a remote AS. The output should look like this:

.. code-block::

  > scion showpaths 71-2:0:48
  Available paths to 71-2:0:48
  4 Hops:
  [0] Hops: [71-2:0:4a 2>31 71-20965 7>3 71-2:0:35 6>1 71-2:0:48] MTU: 1452 NextHop: 141.44.25.151:30001 Status: alive LocalIP: 10.44.25.3
  [1] Hops: [71-2:0:4a 2>31 71-20965 8>9 71-2:0:35 6>1 71-2:0:48] MTU: 1452 NextHop: 141.44.25.151:30001 Status: alive LocalIP: 10.44.25.3
  [2] Hops: [71-2:0:4a 2>31 71-20965 9>8 71-2:0:35 6>1 71-2:0:48] MTU: 1452 NextHop: 141.44.25.151:30001 Status: alive LocalIP: 10.44.25.3
  [3] Hops: [71-2:0:4a 2>31 71-20965 12>11 71-2:0:35 6>1 71-2:0:48] MTU: 1452 NextHop: 141.44.25.151:30001 Status: alive LocalIP: 10.44.25.3
  [4] Hops: [71-2:0:4a 2>31 71-20965 17>7 71-2:0:35 6>1 71-2:0:48] MTU: 1452 NextHop: 141.44.25.151:30001 Status: alive LocalIP: 10.44.25.3
  [5] Hops: [71-2:0:4a 2>31 71-20965 18>10 71-2:0:35 6>1 71-2:0:48] MTU: 1452 NextHop: 141.44.25.151:30001 Status: alive LocalIP: 10.44.25.3

If you get a warning from the Windows Firewall about ``scion.exe`` accessing the newtwork, just click ``Abort``. The ``scion`` command only needs outgoing connetions to make it through the Firewall, which should already be the case in the default Firewall configuration. You only need to add a firewall exception for a SCION program if you whish to run it as a server reachable to other hosts on the network. That being said, if ``scion showpaths`` gives you timeouts it might be worth checking the Firewall configuration. A useful tool for analyzing which packets make it through is `Wireshark <https://www.wireshark.org/>`_ for which a `SCION plugin <https://github.com/scionproto/scion/tree/master/tools/wireshark>`_ is available. If you want the interact with SCION packets on a deeper level, there is also a `SCION layer for Scapy <https://github.com/lschulz/scapy-scion-int>`_.
