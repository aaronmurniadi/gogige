|  ![img-29.jpeg](img-29.jpeg) CAM |   | ![img-30.jpeg](img-30.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GenTL Consumer through the C interface nor the nodemap, until the Consumer explicitly requests to update the list. Similarly, if the device gets physically disconnected (or powered off), it will not be removed from the list published to the GenTL Consumer (IFGetNumDevices/DeviceSelector) until the next list update is executed.

- List of devices discovered on a given interface at the time of the last request to update the device list (IFUpdateDeviceList function or DeviceUpdateList command in the nodemap) and published to the GenTL Consumer through the C interface (IFGetNumDevices) and the nodemap (DeviceSelector). While the GenTL Producer maintains just a single list and publishes it identically through both interfaces the two views might still temporarily differ from the GenTL Consumer's viewpoint. If the list is updated from the nodemap (using DeviceUpdateList command), it is reflected by the nodemap directly through the C interface. If the list is updated from the C interface (IFUpdateDeviceList function), it might not be reflected by the nodemap directly due to GenApi caching effects. Finally, both views (C interface and nodemaps) might be used by the GenTL Consumer independently. It might be querying information through the C interface about one device, while the user selected (DeviceSelector) a different one in the nodemap.

- Currently opened local device modules, e.g., modules for which the GenTL Consumer owns valid handles (IFOpenDevice). This is typically a subset of the list published through the C interface and the nodemap. However, the specification requires that instantiated handles are not affected by any list update requests. This means that if a device is physically disconnected at runtime (while the consumer owns a valid handle for it), the handle remains valid, until explicitly closed (DevClose) – even if most operations upon that handle would simply fail. A request to update the device list would, remove such a device from the list published by the parent interface. A module handle becomes implicitly invalid whenever its parent (or grandparent) module is closed. Please note that the specification allows to open the device (IFOpenDevice, similarly for interfaces) directly using a known device ID (the ID's should be unique and must not change between sessions) without calling IFUpdateDeviceList first. In this case the GenTL Producer might need to (re)execute the device discovery process on its own to connect to the device, providing the handle to the GenTL Consumer while the published device list remains unchanged (possibly even empty) until next list-update request.

### 3.8 Example

This sample code shows how to instantiate the first Data Stream of the first Device connected to the first Interface. Error checking is omitted for clarity reasons.

#### 3.8.1 Basic Device Access

Functions used in this section are listed in subsequent sections.

{
    InitLib( );
    TL_HANDLE hTL = OpenTL( );