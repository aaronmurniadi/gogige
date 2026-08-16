|  ![img-237.jpeg](img-237.jpeg) CAM |   | ![img-238.jpeg](img-238.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  PART_DATATYPE_JPEG | 10 | JPEG compressed data in the format described in GEV 2.0.  |
|  PART_DATATYPE_JPEG2000 | 11 | JPEG 2000 compressed data in the format described in GEV 2.0.  |
|  PART_DATATYPE_CUSTOM_ID | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific.  |

##### 6.4.4.11 FLOW_INFO_CMD

enum FLOW_INFO_CMD

This enumeration defines commands to retrieve information with the DSGetFlowInfo function on a data stream handle.

The column labeled “Impl” in the following table lists if the implementation of a given command is mandatory (M), optional (O) or conditional mandatory (CM). Mandatory means that a GenTL Producer must implement the listed command. Optional means that it is up to the implementor if a given command is implemented or not. Conditional Mandatory means that command is to be implemented if possible.

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  FLOW_INFO_SIZE | M | 0 | Size of the flow in bytes.This information is essential for GenTL Consumer to be able to allocate suitable buffers for flow acquisition usingDSAnnounceCompositeBuffer.In case of GenDC streaming this directly corresponds to the information from GenDC mapping table.Note: querying the flow structure through FLOW_INFO_SIZE and throughSTREAM_INFO_FLOW_TABLE must yield same results.Data type: SIZET  |
|  FLOW_INFO_CUSTOM_ID | O | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific.  |