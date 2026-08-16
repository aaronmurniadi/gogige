|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|   | CXP25_X2 CXP1_X4 CXP2_X4 CXP3_X4 CXP5_X4 CXP6_X4 CXP10_X4 CXP12_X4 CXP25_X4 CXP1_X8 CXP2_X8 CXP3_X8 CXP5_X8 CXP6_X8 CXP10_X8 CXP12_X8 CXP25_X8 Device-specific  |
| --- | --- |

This feature allows specifying the Link configuration for the communication between the Receiver and Transmitter Device. In most cases this feature does not need to be written because automatic discovery will set configuration correctly to the value returned by CxpLinkConfigurationPreferred. Note that the currently active configuration of the Link can be read using CxpLinkConfigurationStatus.

The Link configuration is specified as the combination of the Connection speed and the number of active Connections using the following format "CXPm_Xn", where m is the Connection speed and n the number of active Connections. Selecting Auto sets the Link to normal, automatic, discovery, as described in the CoaXPress standard. The Receiver Device will automatically discover any Transmitter Device connected from then on.

Possible values are:

- Auto: Sets Automatic discovery for the Link Configuration.
- CXP1_X1: Force the Link to 1 Connection operating at CXP-1 speed (1.25 Gbps).
- CXP2_X1: Force the Link to 1 Connection operating at CXP-2 speed (2.50 Gbps).
- CXP3_X1: Force the Link to 1 Connection operating at CXP-3 speed (3.125 Gbps).
- CXP5_X1: Force the Link to 1 Connection operating at CXP-5 speed (5.00 Gbps).
- CXP6_X1: Force the Link to 1 Connection operating at CXP-6 speed (6.25 Gbps).
- CXP10_X1: Force the Link to 1 Connection operating at CXP-10 speed (10.00 Gbps).
- CXP12_X1: Force the Link to 1 Connection operating at CXP-12 speed (12.50 Gbps).