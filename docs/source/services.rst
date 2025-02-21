Services
=======================================

We provide the following tool to facilitate an easy setup of endhosts within SCIERA: The  `scion-orchestrator <https://github.com/netsys-lab/scion-orchestrator>`_ .

The scion-orchestrator is a tool that automates the setup of SCION on endhosts and infratstructure nodes and is recommended to be used to setup the open-source SCION stack in SCIERA.

In addition to this tool, there is also the possibility to run the bootstrap-server and the endhost bootstrapper manually. This is described in the following sections.

SCION Endhost Bootstrapping
----------------

When running a SCION as via the scion-orchestrator, a bootstrap server is already included. This server is used to serve the SCION configuration to the endhosts that are part of the AS.

The boostrap-server requires an IP address that is reachable by all potential SCION hosts within the AS. This is typically an IP address of the official IP range of the institution. Of course we can help setting finding this IP and setting up the server instance, please contact us in case any questions occur: marten.gartner@ovgu.de.

Per default the bootstrap server of the scion-orchestrator is listening on 127.0.0.1, to configure the IP properly, check `scion orchestrator bootstrap config <https://github.com/netsys-lab/scion-orchestrator/tree/master/doc/bootstrap>`_

The alternative way to setup the bootstrapping server without the scion-orchestrator is to use the `boostrap-server repository <https://github.com/netsys-lab/bootstrap-server>`_ and run this setup on one of the machines that runs your SCION AS. This ensures that all configurations that are served to SCION endhosts match the current AS configuration.

Once the bootstrap-server is set up, please follow the guide on SCION Endhost Installers to bootstrap SCION on your endhosts.

.. _scion-endhost-installer:
SCION Endhost Installer
----------------

The scion-orchestrator facilitates the setup of SCION on pure endhosts, to achieve native SCION connectivity.

* Installation of the SCION Binaries onto the machine.
* Connect to the given bootstrap-server to fetch a SCION configuration.
* Run all required SCION endhost services.

So far we provide preliminary support to run SCION on the following operating systems:

* Linux (amd64)
* Windows (amd64)


To bootstrap into a SCIERA AS that has a bootstrap server running, choose from the following list of options to find your AS and operating system.

If your AS is not in the list, please contact us.

.. raw:: html

    <div id="config-selector">
        <style>
            /* Basic styling for select elements */
            #select1, #select2 {
                padding: 10px;
                font-size: 14px;
                border: 1px solid #ccc;
                border-radius: 5px;
                margin-right: 10px;
                width: 150px;
            }

            /* Styling for the button */
            #navigate-button {
                padding: 10px 20px;
                font-size: 14px;
                color: white;
                background-color: #007BFF;
                border: none;
                border-radius: 5px;
                cursor: pointer;
                transition: background-color 0.3s ease;
            }

            /* Hover effect for the button */
            #navigate-button:hover {
                background-color: #0056b3;
            }

            /* Add some spacing and alignment */
            label {
                margin-right: 5px;
                font-size: 14px;
                font-weight: bold;
                margin-top: 25px;
            }

            #config-selector {
                display: flex;
                align-items: center;
                gap: 10px;
                flex-wrap: wrap;
                flex-flow: column;
            }
        </style>

        <label for="select1">SCIERA Autonomous System</label>
        <select id="select1" style="margin-right: 10px;">
            <option value="uva">71-225 (UVA)</option>
            <option value="ovgu">71-2:0:4a (Ovgu)</option>
        </select>

        <label for="select2">Operating System</label>
        <select id="select2" style="margin-right: 10px;">
            <option value="linux_amd64">Linux (amd64)</option>
            <option value="windows_amd64">Windows (amd64)</option>
        </select>

        <button id="navigate-button">Get SCION!</button>
        <br/>
    </div>

    <script>
        document.getElementById("navigate-button").addEventListener("click", function () {
            // Get selected values
            const select1Value = document.getElementById("select1").value;
            const select2Value = document.getElementById("select2").value;

            // Define configuration mappings
            const configLinks = {
                "uva_linux_amd64": "https://github.com/netsys-lab/scion-orchestrator-releases/releases/download/v0.0.1-uva/scion_linux_amd64.zip",
                "uva_windows_amd64": "https://github.com/netsys-lab/scion-orchestrator-releases/releases/download/v0.0.1-uva/scion_windows_amd64.zip",
                "ovgu_linux_amd64": "https://github.com/netsys-lab/scion-orchestrator-releases/releases/download/v0.0.1-ovgu/scion_linux_amd64.zip",
                "ovgu_windows_amd64": "https://github.com/netsys-lab/scion-orchestrator-releases/releases/download/v0.0.1-ovgu/scion_windows_amd64.zip"
            };

            // Build the key for the selected combination
            const selectedConfig = `${select1Value}_${select2Value}`;

            // Redirect to the corresponding link
            if (configLinks[selectedConfig]) {
                // window.location.href = configLinks[selectedConfig];
                window.open(configLinks[selectedConfig], '_blank');
            } else {
                alert("Configuration not found.");
            }
        });
    </script>

To run SCION on your endhost, please follow the following steps as depicted `here <https://github.com/netsys-lab/scion-orchestrator-releases>`_.

To manually run a scion bootstrapper without the orchestrator, refer to the bootstrapper repositories: `Code <https://github.com/netsec-ethz/bootstrapper>`_, `Design <https://github.com/scionproto/scion/blob/master/doc/dev/design/endhost-bootstrap.rst>`_
