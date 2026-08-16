|  ![img-41.jpeg](img-41.jpeg) CAM |   | ![img-42.jpeg](img-42.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

• xml for uncompressed XML description files
• zip for zip-compressed XML description files

#### 4.1.3 Example

This sampel shows how to retrieve the Port module xmls.

{
    // Retrieve the number of available URLs
    GCGetNumPortURLs( hModule, NumURLs );
    for( i=0; i<NumURLs; i++ )
    {
    URLSize = 0;
    GCGetPortURLInfo( hModule, i, URL_INFO_URL, 0, 0, &URLSize );

    // Retrieve an string buffer to store the URL
    GCGetPortURLInfo( hModule, i, URL_INFO_URL, 0, pURL, &URLSize );

    if ( ParseURLLocation( pURL ) == local )
    {
    // Retrieve the address within the module register map from the URL
    Addr = ParseURLLocalAddress( pURL );
    Length = ParseURLLocalLength( pURL );
    // Retrieve an XMLBuffer to store the XML with the size Length
    ...
    // Load xml from local register map into memory
    GCReadPort( hModule, Addr, XMLBuffer, Length );
    }
    }
}

### 4.2 Signaling

The Signaling is used to notify the GenTL Consumer on asynchronous events. Usually all the communication is initiated by the GenTL Consumer. With an event the GenTL Consumer can get notified on specific GenTL Producer operations. This mechanism is an implementation of the observer pattern where the calling GenTL Consumer is the observer and the GenTL Producer is being observed.

The reason why an event object approach was chosen rather than callback functions is mainly thread priority problems. A callback function to signal the arrival of a new buffer is normally executed in the thread context of the acquisition engine. Thus all processing in this callback function is done also with its priority. If no additional precautions are taken the acquisition engine is blocked as long the callback function does processing.