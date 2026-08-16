|  GENICAM |   | ![img-58.jpeg](img-58.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

The SmartFeature node can retrieve the address of a smart feature when it is given a global unique identifier (GUID) describing that feature in the <FeatureID> element. It also inherits the elements and attributes from the Register node. The following example retrieves the address of a smart feature with a GUID of {5590D58E-1B84-11D8-8447-00105A5BAE55}:

<SmartFeature Name="TimeStampAdr">
    <FeatureID>5590D58E - 1B84 - 11D8 - 8447 - 00105A5BAE55</FeatureID>
    <Address>0xfffff2f00010</Address>
    <pPort>Device</pPort>
</SmartFeature>

#### 2.8.16 Port

The Port object is just a proxy that forwards Read and Write calls to the transport layer. Note, however, that the proxy has all of the properties of a Node. For example, it can be “not present.” This will tell all dependent nodes that the transport layer driver is currently not open and as a result, all dependent features will automatically also be “not present.” Another example would be the implementation of a user set loader. If a user set is loaded from flash ROM inside the camera, all features inside the node graph must be invalidated. This can be achieved by simply invalidating the Port node, which in turn can be automated using a <pInvalidator> linked to the ReadUserSet feature node.</pInvalidator>

If the transport layer is restricted to a maximal chunk length or needs special alignment, e.g., quadlet-wise, the transport layer implementation must emulate the IPort interface by breaking down calls longer than the maximum chunk length into multiple calls and must pad calls not fitting the necessary alignment. In order to support certain types of quadlet based interface the <SwapEndianess>> element has been introduced: if it reads true the endianess of each quadlet must be swapped before exposing the data to GenICam via the IPort interface.

The Port node inherits the elements and attributes of the Node node. In addition, it can have a <ChunkID> element which is a hexadecimal number that identifies a chunk of data in a buffer. This chunk may be mapped to a virtual port that does not give access to a real device, but rather to the chunk of data residing in memory.</ChunkID>

<Port Name="Device" NameSpace="Standard">
    <ChunkID>4711</ChunkID>
</Port>

Instead of a <ChunkId> entry a <pChunkID> entry may be used to retrieve the ChunkID value from another node.</pChunkID>

Chunk ports implement two int64 pseudo registers: The CHUNK_BASE_ADDRESS_REGISTER at the address of INT64_MAX indicates the memory address of the start of the chunk. The CHUNK_LENGTH_REGISTER at the address of (INT64_MAX-15) indicates the length of the chunk excluding the trailer.

Chunk ports can deal with negative addresses which will be interpreted as offset from the back of the chunk. If a register node is mapped to a chunk port and a chunk is present the method GetAddress() will always return the address from start of the chunk.