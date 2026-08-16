|  ![img-31.jpeg](img-31.jpeg) CAM |   | ![img-32.jpeg](img-32.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

IF_HANDLE hIface = OpenFirstInterface( hTL );
DEV_HANDLE hDevice = OpenFirstDevice( hIface );
DS_HANDLE hStream = OpenFirstDataStream( hDevice );

// At this point we have successfully created a data stream on the first
// device connected to the first interface. Now we could start to
// capture data...
CloseDataStream( hStream );
CloseDevice( hDevice );
CloseInterface( hIface );
CloseTL( hTL );
CloseLib( );
}

#### 3.8.2 InitLib

Initialize GenTL Producer.

void InitLib( void )
{
    GCInitLib( );
}

#### 3.8.3 OpenTL

Retrieve TL Handle.

TL_HANDLE OpenTL( void )
{
    TLOpen( hTL );
}

#### 3.8.4 OpenFirstInterface

Retrieve first Interface Handle.

IF_HANDLE OpenFirstInterface( hTL )
{
    TLUpdateInterfaceList( hTL );
    TLGetNumInterfaces( hTL, NumInterfaces );
    if ( NumInterfaces > 0 )
    {
    // First query the buffer size
    TLGetInterfaceID( hTL, 0, IfaceID, &bufferSize );

    // Open interface with index 0
    TLOpenInterface( hTL, IfaceID, hNewIface );