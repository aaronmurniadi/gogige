|  ![img-33.jpeg](img-33.jpeg)CAN |   | ![img-34.jpeg](img-34.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

return hNewIface;
}

#### 3.8.5 OpenFirstDevice

Retrieve first Device Handle.

DEV_HANDLE OpenFirstDevice(hIF)
{
    IFUpdateDeviceList(hIF);
    IFGetNumDevices(hTL, NumDevices);
    if (NumDevices > 0)
    {
    // First query the buffer size
    IFGetDeviceID(hIF, 0, DeviceID, &bufferSize);

    // Open interface with index 0
    IFOpenDevice(hIF, DeviceID, hNewDevice);
    return hNewDevice;
    }
}

#### 3.8.6 OpenFirstDataStream

Retrieve first data Stream.

DS_HANDLE OpenFirstDataStream(hDev)
{
    // Retrieve the number of Data Stream
    DevGetNumDataStreams(hDev, NumStreams);

    if (NumStreams > 0)
    {
    // Get ID of first stream using
    DevGetDataStreamID(hdev, 0, StreamID, buffersize);
    // Instantiate Data Stream
    DevOpenDataStream(hDev, StreamID, hNewStream);
    }
}

#### 3.8.7 CloseDataStream

Close Data Stream.

void CloseDataStream ( hStream ) {