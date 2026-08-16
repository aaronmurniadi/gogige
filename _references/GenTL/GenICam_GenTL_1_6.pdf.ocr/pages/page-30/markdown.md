|  ![img-35.jpeg](img-35.jpeg)CAN |   | ![img-36.jpeg](img-36.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

DSClose(hStream);

#### 3.8.8 CloseDevice

Close Device.

void CloseDevice(hDevice)
{
    DevClose(hDevice);
}

#### 3.8.9 CloseInterface

Close Interface.

void CloseInterface(hIface)
{
    IFClose(hIface);
}

#### 3.8.10 CloseTL

Close System module.

void CloseTL(hTL)
{
    TLClose(hTL);
}

#### 3.8.11 CloseLib

Shutdown GenTL Producer.

void CloseLib( void )
{
    GCCloseLib( );
}