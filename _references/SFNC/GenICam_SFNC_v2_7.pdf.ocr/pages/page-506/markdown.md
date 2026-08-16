|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Controls if the selected parameters are locked during acquisition.

Locking certain parameters during acquisition includes respecting TLParamsLocked as well as preventing internal changes from modules like sequencer or multiple regions.

### 27.2.5 PayloadSize

|  Name | PayloadSize  |
| --- | --- |
|  Category | TransportLayerControl  |
|  Level | Recommended  |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | B  |
|  Visibility | Expert  |
|  Values | ≥0  |

Provides the number of bytes transferred for each data buffer or chunk on the stream channel. This includes any end-of-line, end-of-frame statistics or other stamp data. This is the total size of data payload for a data block.

For the devices supporting multiple Stream Channels, the DeviceStreamChannelSelector feature should be used to select PayloadSize. This permits to inquire the payload size of each Stream channel individually.

This feature is mainly used by the application software to determine size of the data buffers to allocate (largest buffer possible for current mode of operation).

For example, an image with no statistics, timestamp or other metadata data has typically a PayloadSize equals to (width x height x pixel size) in bytes. But it is strongly recommended to retrieve PayloadSize from the camera instead of relying on the above formula.

It is updated every time a feature affecting its value is changed and must be static after TLParamsLocked is asserted to guarantee that it accurately represents the maximum size of the data buffer that will be streamed out when AcquisitionStart is executed. This feature is generally mandatory for transmitters and transceivers of most Transport Layers.

### 27.2.6 GenDCStreamingMode

|  Name | GenDCStreamingMode  |
| --- | --- |
|  Category | TransportLayerControl  |