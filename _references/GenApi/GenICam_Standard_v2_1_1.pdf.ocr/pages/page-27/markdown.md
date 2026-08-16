|  GENICAM |   | ![img-40.jpeg](img-40.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

#### 2.8.3 Register

The Register node maps to a contiguous array of bytes in the register space of the camera. The Register node implements the IRegister interface and inherits its elements and attributes from the Node node. It in turn leaves its elements and nodes to all specialized register access nodes, such as IntReg, StringReg, etc. A Register node, however, can also be instantiated on its own giving access to the raw binary data. Here is a simple example:

<Register Name="SensorTemperature">
    <Address>0xff00</Address>
    <Length>4</Length>
    <AccessMode>RO</AccessMode>
    <pPort>Device</pPort>
    <Cachable>No</Cachable>
    <PollingTime>10000</PollingTime>
</Register>

The example exposes the temperature of the camera's sensor. The temperature can change at any time and is therefore not cacheable. If displayed, it should be polled every 10000 ms.

The <address> element gives the address of the register in the camera's register space.</address>

The <length> element gives the length of the register in bytes. Alternatively the length can be read from another node using an <pLength> entry.</pLength></length>

The <AccessMode> element can have the values RW (read/write), RO (read only), or WO (write only) and indicates what the camera can deliver.</AccessMode>

The <pPort> element contains the name of a Port node that gives access to the camera's register space (for details see section 2.8.16).</pPort>

The <cacheable> element can have the values NoCache, WriteThrough, and WriteAround. WriteThrough means that a value written to the camera is written to the cache as well. WriteAround means that only read values are written to the cache. The latter behavior makes sense, for example, with an IFloat::Gain node where the user can write any value, but when reading back, will retrieve a value that has been rounded by the camera to a value the internal analog-to-digital converter is able to deliver. Note that caching is an optional feature of any implementation.</cacheable>

The <PollingTime> element denotes a recommended time interval [in ms] after which a node should be invalidated. Note that polling is an optional feature of any implementation and the polling time is a hint only.</PollingTime>

Instead of a single <address> entry, a register can have multiple entries for the <address>, <pAddress>, and/or <IntSwissKnife> types. The values of these entries are summed, yielding the address of the register node.</pAddress></address>

The <pAddress> element points to a node implementing an IInteger interface delivering a contribution to the final address.</pAddress>

The <IntSwissKnife> element can be used to compute an address contribution from multiple sources (for details see section 2.8.12). </IntSwissKnife>

The <pIndex Offset="12"> element points to a node implementing an IInteger interface delivering an index. The element has an attribute Offset. The product of index and Offset is