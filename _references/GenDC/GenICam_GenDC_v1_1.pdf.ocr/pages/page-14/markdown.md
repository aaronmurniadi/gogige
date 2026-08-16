|  2 | 10 | **Flags** Flags specifying the characteristics and format of the Container. See section 4 for more information.  |   |   |
| --- | --- | --- | --- | --- |
|   |   |  Width (bits) | Bit offset (lsb << x) | Description  |
|   |   |  1 | 0 | **TimestampPTP** If true, the timestamps of the Components are relative to the epoch January 1, 1970 00:00:00 (TAI) like in the PTP IEEE-1588 format.  |
|   |   |  1 | 1 | **ComponentInvalid** If True, Components in the Container might be invalid. The **Invalid** flag of each Component must be checked before using it. An example use case is a device with a static ComponentCount but leaving out Components on-the-fly if they cannot be generated.  |
|   |   |  14 | 2 | **Reserved** (set to 0)  |
|   |   |  |   |   |
|  4 | 12 | **HeaderSize** = Size of the Container Header. Size of the Container Header in bytes including the variable sized ComponentOffset array.  |   |   |
|  8 | 16 | **Id** Container identifier. Strictly monotonically incrementing by 1 with each transmitted Container. Note: Start value and reset condition can be specified by the Transport Layer Protocols.  |   |   |