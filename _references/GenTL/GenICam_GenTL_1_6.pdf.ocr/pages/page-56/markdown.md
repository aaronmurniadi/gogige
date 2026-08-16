|  ![img-78.jpeg](img-78.jpeg)CAN |   | ![img-79.jpeg](img-79.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

- Allow transferring independent data parts into discrete target memory locations defined by the user or GenTL Consumer to optimize the possibly use case specific data handling (for example to route some data to GPU memory).

To achieve the first goal (parallel transfer), no special cooperation from GenTL Consumer is required. If acquisition into contiguous buffers is satisfactory, the GenTL Consumer can announce them as usual using DSAllocAndAnnounceBuffer or DSAnnounceBuffer. The GenTL Producer is responsible to configure the multi-flow acquisition under the hood (according to requirements of its respective transport layer technology) and transfer them into suitable locations of the target buffer. The split of the buffer to per-flow segments happens in this case transparently to GenTL Consumer which receives the data in single buffer, for example as a linear GenDC Container (5.7.3).

To achieve the second goal (discrete target locations), the GenTL Consumer must announce composite buffers (DSAnnounceCompositeBuffer, see 5.7.1) with structure matching the current flow configuration. In particular, each composite buffer should have not less segments than currently expected flows, each of the segments should not be smaller than the corresponding flow.

When announcing buffers for acquisition, the GenTL Consumer would first query the flow table using DSGetNumFlows and DSGetFlowInfo (FLOW_INFO_SIZE). Alternatively, it can query it in GenDC format using DSGetInfo (STREAM_INFO_FLOW_TABLE). Knowing the flow table, it would allocate and announce composite buffers with structure matching the reported flow table, so that each flow could transfer its data to the corresponding buffer segment. In presence of flows, the GenTL Producer must always be able to provide the flow table, otherwise the GenTL Consumer has no way to allocate its composite buffers.

It is important to keep in mind that the flow structure can depend on device configuration and it is thus important to query it after the configuration is finished, just before the acquisition start (refer also to TLParamsLocked feature definition in SFNC). When composite buffers are used for acquisition, the flow mapping table defines the buffer allocation requirements in a similar way as PayloadSize feature defines it for simple contiguous buffers (as discussed in 5.2.1).

The following figure illustrates acquisition through data stream flows for cases when consumer announces contiguous or composite buffers.