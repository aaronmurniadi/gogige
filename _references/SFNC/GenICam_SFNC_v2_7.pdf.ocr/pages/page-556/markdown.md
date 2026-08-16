|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|   | CXP1_X2 CXP2_X2 CXP3_X2 CXP5_X2 CXP6_X2 CXP10_X2 CXP12_X2 CXP25_X2 CXP1_X4 CXP2_X4 CXP3_X4 CXP5_X4 CXP6_X4 CXP10_X4 CXP12_X4 CXP25_X4 CXP1_X8 CXP2_X8 CXP3_X8 CXP5_X8 CXP6_X8 CXP10_X8 CXP12_X8 CXP25_X8 Device-specific  |
| --- | --- |

This feature indicates the current and active Link configuration used by the Device.

When the Link is active, this feature returns the Link configuration as the combination of the Connection speed and the number of active Connections using the following format "CXPm_Xn", where m is the Connection speed and n the number of active Connections. For example "CXP6_X4" means 4 connections are operating at CXP-6 speed (6.25 Gbps) so the total speed on the virtual single link is 25 Gbps.

Possible values are:

- None: The Link configuration of the Device is unknown. Either the configuration operation has failed or there is nothing connected.
- Pending: The Device is in the process of configuring the Link. The Link cannot be used yet.
- CXP1_X1: 1 Connection operating at CXP-1 speed (1.25 Gbps).
- CXP2_X1: 1 Connection operating at CXP-2 speed (2.50 Gbps).
- CXP3_X1: 1 Connection operating at CXP-3 speed (3.125 Gbps).